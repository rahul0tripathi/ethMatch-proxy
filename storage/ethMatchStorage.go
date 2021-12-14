package storage

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/ethMatch/proxy/common"
	"github.com/ethMatch/proxy/config"
	"github.com/ethMatch/proxy/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"strings"
	"time"
)

const (
	ticketPrefix        = "TICKET"
	userPrefix          = "USER"
	proposalPrefix      = "LOBBYPROPOSAL"
	proposedLobbyPrefix = "PROPOSEDLOBBY"
	lobbySignatures     = "LOBBYSIGNATURES"

	ticketPool            = "TICKET_POOL"
	maxProposalFetchCount = 100
)

var (
	proposalExpireTime time.Duration
	maxTicketTTL       time.Duration
	redisAddress       = flag.String("redisAddr", "localhost:6379", "address of redis server [IP]:[PORT]")
	redisPassword      = flag.String("redisPass", "", "redis server password")
	redisDb            = flag.Int("redisDb", 0, "redis database to use")
	CommonStorage      = NewEthMatchStorage()
)

func genTicketKey(ticketId string) string {
	return ticketPrefix + ":" + ticketId
}
func genUserKey(addr ethcommon.Address) string {
	return userPrefix + ":" + addr.String()
}
func genNewProposalKey(lobbyId string, user ethcommon.Address) string {
	return proposalPrefix + ":" + user.String() + ":" + lobbyId
}
func genNewLobbyKey(lobbyId string) string {
	return proposedLobbyPrefix + ":" + lobbyId
}
func newProposalMatchString(user ethcommon.Address) string {
	return proposalPrefix + ":" + user.String() + "*"
}
func genLobbySignatureMap(lobbyId string) string {
	return lobbySignatures + ":" + lobbyId
}
func getLobbyIdFromProposal(key string) string {
	splitKeys := strings.Split(key, ":")
	if len(splitKeys) != 3 {
		return ""
	} else {
		return splitKeys[2]
	}
}

type EthMatchStorage struct {
	redisClient *redis.Client
}

func NewEthMatchStorage() EthMatchStorage {
	var err error
	proposalExpireTime, err = time.ParseDuration(config.ENV.ProposalExpireDuration)
	if err != nil {
		zap.Error(err)
	}
	maxTicketTTL, err = time.ParseDuration(config.ENV.MaxTicketTTL)
	if err != nil {
		zap.Error(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     config.ENV.Redis.Address,
		Password: config.ENV.Redis.Password,
		DB:       *redisDb,
	})
	return EthMatchStorage{
		redisClient: client,
	}
}

func (em EthMatchStorage) AddTicket(ctx context.Context, ticket types.Ticket) (err error) {
	var ticketJSON []byte
	ticketJSON, err = json.Marshal(ticket)
	if err != nil {
		return
	}
	err = em.redisClient.Set(ctx, genTicketKey(ticket.Id.String()), ticketJSON, maxTicketTTL).Err()
	if err != nil {
		return
	}
	err = em.redisClient.ZAdd(ctx, ticketPool, &redis.Z{
		Score:  float64(ticket.Rank),
		Member: ticket.Id.String(),
	}).Err()
	err = em.redisClient.Set(ctx, genUserKey(ticket.Player), ticket.Id.String(), maxTicketTTL).Err()
	return
}

func (em EthMatchStorage) RemoveTicket(ctx context.Context, ticketId uuid.UUID) {
	em.redisClient.Del(ctx, genTicketKey(ticketId.String()))
	em.redisClient.ZRem(ctx, ticketPool, ticketId.String())
}

func (em EthMatchStorage) GetTicketById(ctx context.Context, ticketId uuid.UUID) (ticket types.Ticket, err error) {
	replyCmd := em.redisClient.Get(ctx, genTicketKey(ticketId.String()))
	if replyCmd.Err() != nil {
		err = replyCmd.Err()
		return
	}
	var ticketJSON []byte
	ticketJSON, err = replyCmd.Bytes()
	if err != nil {
		return
	}
	err = json.Unmarshal(ticketJSON, &ticket)
	return
}
func (em EthMatchStorage) GetUserTickets(ctx context.Context, user ethcommon.Address) (ticket types.Ticket, err error) {
	currentTicket := em.redisClient.Get(ctx, genUserKey(user))
	err = currentTicket.Err()
	if err != nil {
		return
	}
	ticket, err = em.GetTicketById(ctx, uuid.MustParse(currentTicket.String()))
	return
}
func (em EthMatchStorage) GetUserLobbyProposals(ctx context.Context, user ethcommon.Address) (lobbies []types.Proposal, err error) {
	common.Logger.Debug("proposal string", zap.String("key", newProposalMatchString(user)))
	scanCmd := em.redisClient.Scan(ctx, 0, newProposalMatchString(user), maxProposalFetchCount)
	if err = scanCmd.Err(); err != nil {
		return
	}
	scannedKeys, _ := scanCmd.Val()
	for _, val := range scannedKeys {
		lobby := em.redisClient.Get(ctx, genNewLobbyKey(getLobbyIdFromProposal(val)))
		if err = lobby.Err(); err != nil {
			return
		}
		parsedLobby := types.Proposal{}
		if err = json.Unmarshal([]byte(lobby.Val()), &parsedLobby); err != nil {
			return
		}
		if parsedLobby.Id != "" {
			lobbies = append(lobbies, parsedLobby)
		}
	}
	return
}
func (em EthMatchStorage) GetProposal(ctx context.Context, lobbyId string) (lobby types.Proposal, err error) {
	proposalCmd := em.redisClient.Get(ctx, genNewLobbyKey(lobbyId))
	if err = proposalCmd.Err(); err != nil {
		return
	}
	var data []byte
	data, err = proposalCmd.Bytes()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &lobby)
	return
}
func (em EthMatchStorage) NewProposal(ctx context.Context, lobby types.Proposal) error {
	lobbyBytes, err := json.Marshal(lobby)
	if err != nil {
		return err
	}
	em.redisClient.Set(ctx, genNewLobbyKey(lobby.Id), lobbyBytes, proposalExpireTime)
	for playerAddr := range lobby.PlayerTickets {
		err = em.redisClient.Set(ctx, genNewProposalKey(lobby.Id, playerAddr), lobby.ExpireAt, proposalExpireTime).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
func (em EthMatchStorage) AddSignature(ctx context.Context, lobby types.Proposal, user ethcommon.Address, signature string) (err error) {
	err = em.redisClient.HSet(ctx, genLobbySignatureMap(lobby.Id), user.String(), signature).Err()
	return
}
func (em EthMatchStorage) GetSignature(ctx context.Context, lobby types.Proposal, user ethcommon.Address) (addr string, err error) {
	respCmd := em.redisClient.HGet(ctx, genLobbySignatureMap(lobby.Id), user.String())
	if err = respCmd.Err(); err == nil {
		addr = respCmd.Val()
	}
	return
}
func (em EthMatchStorage) GetProposalSignatures(ctx context.Context, lobbyId string) (signatures types.LobbySignatures, err error) {
	signaturesCmd := em.redisClient.HGetAll(ctx, genLobbySignatureMap(lobbyId))
	signatures = map[ethcommon.Address]string{}
	if err = signaturesCmd.Err(); err != nil {
		return
	}
	for addr, sig := range signaturesCmd.Val() {
		signatures[ethcommon.HexToAddress(addr)] = sig
	}
	return
}
func (em EthMatchStorage) GetAllTickets(ctx context.Context) (tickets []types.Ticket, err error) {
	countCmd := em.redisClient.ZCard(ctx, ticketPool)
	if err = countCmd.Err(); err != nil {
		return
	}
	common.Logger.Debug("popped ticket counts", zap.Any("count", countCmd.Val()))
	ticketsCmd := em.redisClient.ZPopMax(ctx, ticketPool, countCmd.Val())
	if err = ticketsCmd.Err(); err != nil {
		return
	}
	for _, val := range ticketsCmd.Val() {
		ticketCmd, _ := em.redisClient.Get(ctx, genTicketKey(val.Member.(string))).Bytes()
		tmpTicket := types.Ticket{}
		err = json.Unmarshal(ticketCmd, &tmpTicket)
		if err != nil || tmpTicket.Id.String() == "" {
			return
		}
		tmpTicket.Rank = uint64(val.Score)
		tickets = append(tickets, tmpTicket)
	}
	common.Logger.Debug("tickets with rank", zap.Any("tickets", tickets))
	return
}
