package matchmaker

import (
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/orgs/ethMatch/p2pmatch/common"
	"github.com/orgs/ethMatch/p2pmatch/storage"
	"github.com/orgs/ethMatch/p2pmatch/types"
	"go.uber.org/zap"
	"sync"
	"time"
)

type BasicMatchMaker struct {
	MMTime  time.Duration
	io      sync.RWMutex
	storage storage.Storage
	stream  chan types.Lobby
	config  map[string]interface{}
}

const (
	proposalExpireTime = time.Minute * 3
)

func NewBasicMatchMaker(ticker time.Duration, minPlayers int, maxPlayers int) BasicMatchMaker {
	return BasicMatchMaker{
		MMTime:  ticker,
		io:      sync.RWMutex{},
		storage: storage.NewEthMatchStorage(),
		stream:  make(chan types.Lobby),
		config: map[string]interface{}{
			"minPlayers": minPlayers,
			"maxPlayers": maxPlayers,
		},
	}
}
func (mm *BasicMatchMaker) StartMatchMaker() {
	for now := range time.Tick(mm.MMTime) {
		now := now
		go func() {
			err := mm.MMF(now)
			if err != nil {
				common.Logger.Error("MMF error", zap.Error(err))
			}
		}()
	}
}
func (mm *BasicMatchMaker) GetLobbyStream() <-chan types.Lobby {
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
			var tmpTicketGroup []types.Ticket
			if i+max > len(tickets) {
				tmpTicketGroup = tickets[i:]

			} else {
				tmpTicketGroup = tickets[i : i+max]
			}
			common.Logger.Debug("new ticket group formed", zap.Any("tmpTicketGroup", tmpTicketGroup))
			common.Logger.Debug("evaluating expn", zap.Bool("is valid group", len(tmpTicketGroup) >= mm.config["minPlayers"].(int) && len(tmpTicketGroup) <= mm.config["maxPlayers"].(int)))
			if len(tmpTicketGroup) >= mm.config["minPlayers"].(int) && len(tmpTicketGroup) <= mm.config["maxPlayers"].(int) {
				tmpRanks := map[ethcommon.Address]uint64{}
				tmpTickets := map[ethcommon.Address]uuid.UUID{}
				lobbyId := crypto.Keccak256Hash([]byte(time.Now().String()))
				for _, ticket := range tmpTicketGroup {
					tmpRanks[ticket.Player] = ticket.Rank
					tmpTickets[ticket.Player] = ticket.Id
					lobbyId = crypto.Keccak256Hash(lobbyId.Bytes(), []byte(ticket.Id.String()))
				}
				common.Logger.Debug("lobby properties", zap.Any("ranks", tmpRanks), zap.Any("tickets", tmpTickets), zap.Any("lobbyId", lobbyId.String()))
				mm.stream <- types.Lobby{
					PlayerTickets:     tmpTickets,
					PlayerSignatures:  nil,
					Ranks:             tmpRanks,
					Pool:              tmpTicketGroup[0].EntryFee * uint64(len(tmpTicketGroup)),
					OperatorsShare:    tmpTicketGroup[0].OperatorsShare,
					EntryFee:          tmpTicketGroup[0].EntryFee,
					Id:                lobbyId.String(),
					ExpireAt:          time.Now().Add(proposalExpireTime).Unix(),
					OperatorSignature: ethcommon.Hash{},
					OperatorAddress:   ethcommon.Address{},
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
	err = mm.storage.AddTicket(context.Background(), ticket)
	return ticket.Id, err
}

func (mm *BasicMatchMaker) RemoveTicketFromQueue(ticketId uuid.UUID) (removed bool) {
	mm.storage.RemoveTicket(context.Background(), ticketId)
	return true
}
