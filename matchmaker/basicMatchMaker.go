package matchmaker

import (
	"context"
	"errors"
	"github.com/ethMatch/proxy/common"
	"github.com/ethMatch/proxy/config"
	"github.com/ethMatch/proxy/eth"
	"github.com/ethMatch/proxy/storage"
	"github.com/ethMatch/proxy/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"sync"
	"time"
)

var (
	CommonMatchMaker   *BasicMatchMaker
	proposalExpireTime time.Duration
)

type BasicMatchMaker struct {
	MMTime      time.Duration
	io          sync.RWMutex
	storage     storage.Storage
	stream      chan types.Proposal
	config      map[string]interface{}
	NextRunTime time.Time
}

func InitMM() {
	d, err := time.ParseDuration(config.ENV.MMTime)
	if err != nil {
		zap.Error(err)
		return
	}
	CommonMatchMaker = NewBasicMatchMaker(d, config.ENV.MinPlayers, config.ENV.MaxPlayers)
	proposalExpireTime, err = time.ParseDuration(config.ENV.ProposalExpireDuration)
	if err != nil {
		zap.Error(err)
	}
}
func NewBasicMatchMaker(ticker time.Duration, minPlayers int, maxPlayers int) *BasicMatchMaker {
	return &BasicMatchMaker{
		MMTime:  ticker,
		io:      sync.RWMutex{},
		storage: storage.NewEthMatchStorage(),
		stream:  make(chan types.Proposal),
		config: map[string]interface{}{
			"minPlayers": minPlayers,
			"maxPlayers": maxPlayers,
		},
	}
}
func (mm *BasicMatchMaker) StartMatchMaker() {
	for now := range time.Tick(mm.MMTime) {
		mm.NextRunTime = time.Now().Add(mm.MMTime)
		now := now
		go func() {
			err := mm.MMF(now)
			if err != nil {
				common.Logger.Error("MMF error", zap.Error(err))
			}
		}()
	}
}
func (mm *BasicMatchMaker) GetLobbyStream() <-chan types.Proposal {
	return mm.stream
}

func (mm *BasicMatchMaker) MMF(timestamp time.Time) error {
	common.Logger.Debug("New MMF Invoked", zap.Time("timestamp", timestamp))
	if _, ok := mm.storage.(storage.EthMatchStorage); ok {
		tickets, err := mm.storage.(storage.EthMatchStorage).GetAllTickets(context.Background())
		if err != nil {
			return err
		}
		i := 0
		max := mm.config["maxPlayers"].(int)
		common.Logger.Debug("maxPlayers", zap.Int("count", max))
		for {
			if i > len(tickets) {
				break
			}
			var tmpTicketGroup []types.Ticket
			if i+max > len(tickets) {
				tmpTicketGroup = tickets[i:]

			} else {
				tmpTicketGroup = tickets[i : i+max]
			}
			common.Logger.Debug("new ticket group formed", zap.Any("tmpTicketGroup", tmpTicketGroup))
			common.Logger.Debug("evaluating expn", zap.Bool("is valid group", len(tmpTicketGroup) >= mm.config["minPlayers"].(int) && len(tmpTicketGroup) <= mm.config["maxPlayers"].(int)))
			if len(tmpTicketGroup) >= mm.config["minPlayers"].(int) && len(tmpTicketGroup) <= mm.config["maxPlayers"].(int) {

				tmpTickets := map[ethcommon.Address]uuid.UUID{}
				lobbyId := crypto.Keccak256Hash([]byte(time.Now().String()))
				for _, ticket := range tmpTicketGroup {
					tmpTickets[ticket.Player] = ticket.Id
					lobbyId = crypto.Keccak256Hash(lobbyId.Bytes(), []byte(ticket.Id.String()))
				}
				common.Logger.Debug("lobby properties", zap.Any("tickets", tmpTickets), zap.Any("lobbyId", lobbyId.String()))
				mm.stream <- types.Proposal{
					PlayerTickets:    tmpTickets,
					PlayerSignatures: nil,
					Pool:             tmpTicketGroup[0].EntryFee * uint64(len(tmpTicketGroup)),
					OperatorsShare:   tmpTicketGroup[0].OperatorsShare,
					EntryFee:         tmpTicketGroup[0].EntryFee,
					Id:               lobbyId.String(),
					ExpireAt:         time.Now().Add(proposalExpireTime).Unix(),
					OperatorAddress:  eth.OperatorPublicAddress,
				}
				common.Logger.Debug("published new lobby", zap.String("id", lobbyId.String()))
			} else {
				break
			}
			i += max
		}
	}
	return nil
}

func (mm *BasicMatchMaker) AddTicketToQueue(ticket types.Ticket) (ticketId uuid.UUID, err error) {
	ticket.Id = uuid.New()
	var existing types.Ticket
	existing, err = mm.storage.GetUserTickets(context.Background(), ticket.Player)
	if existing.EntryFee != 0 {
		err = errors.New("The ticket is already in queue")
		return
	}
	err = mm.storage.AddTicket(context.Background(), ticket)
	return ticket.Id, err
}

func (mm *BasicMatchMaker) RemoveTicketFromQueue(ticketId uuid.UUID) (removed bool) {
	mm.storage.RemoveTicket(context.Background(), ticketId)
	return true
}

func (mm *BasicMatchMaker) PushLobbyToChain(id string) {
	sig, err := mm.storage.GetProposalSignatures(context.Background(), id)
	if err != nil {
		return
	}
	p, err := mm.storage.GetProposal(context.Background(), id)
	if err != nil {
		return
	}
	p.PlayerSignatures = sig
	signatureCount := 0
	for range p.PlayerSignatures {
		signatureCount += 1
	}
	if signatureCount >= mm.config["minPlayers"].(int) {
		err = eth.SubmitNewLobby(p)
		if err != nil {
			common.Logger.Error("failed to push proposal to chain ", zap.Error(err))
		}
	}

}
