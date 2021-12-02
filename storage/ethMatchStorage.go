package storage

import (
	"context"
	"encoding/json"
	"flag"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/orgs/ethMatch/p2pmatch/common"
	"github.com/orgs/ethMatch/p2pmatch/types"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

const (
	ticketPrefix          = "TICKET"
	userPrefix            = "USER"
	proposalPrefix        = "LOBBYPROPOSAL"
	proposedLobbyPrefix   = "PROPOSEDLOBBY"
	lobbySignatures       = "LOBBYSIGNATURES"
	proposalExpireTime    = time.Minute * 3
	ticketPool            = "TICKET_POOL"
	maxTicketTTL          = time.Hour
	maxProposalFetchCount = 100
)

var redisAddress = flag.String("redisAddr", "localhost:6379", "address of redis server [IP]:[PORT]")
var redisPassword = flag.String("redisPass", "", "redis server password")
var redisDb = flag.Int("redisDb", 0, "redis database to use")

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
	client := redis.NewClient(&redis.Options{
		Addr:     *redisAddress,
		Password: *redisPassword,
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
	err = em.redisClient.HSet(ctx, genUserKey(ticket.Player), ticket.Id.String(), time.Now().Add(maxTicketTTL).Unix()).Err()
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
func (em EthMatchStorage) GetUserTickets(ctx context.Context, user ethcommon.Address) (tickets []types.Ticket, err error) {
	allTicketsCmd := em.redisClient.HGetAll(ctx, genUserKey(user))
	if err = allTicketsCmd.Err(); err != nil {
		return
	}
	for ticketId, ttlStr := range allTicketsCmd.Val() {
		ttl, _ := strconv.ParseInt(ttlStr, 10, 64)
		if time.Now().Unix() < ttl {
			ticket, _ := em.GetTicketById(ctx, uuid.MustParse(ticketId))
			if (ticket.Player != ethcommon.Address{}) {
				tickets = append(tickets, ticket)
			}
		} else {
			_ = em.redisClient.HDel(ctx, ticketId)
		}
	}
	return
}
func (em EthMatchStorage) GetUserLobbyProposals(ctx context.Context, user ethcommon.Address) (lobbies []types.Lobby, err error) {
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
		parsedLobby := types.Lobby{}
		if err = json.Unmarshal([]byte(lobby.Val()), &parsedLobby); err != nil {
			return
		}
		if parsedLobby.Id != "" {
			lobbies = append(lobbies, parsedLobby)
		}
	}
	return
}

func (em EthMatchStorage) NewProposal(ctx context.Context, lobby types.Lobby) error {
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
func (em EthMatchStorage) AddSignature(ctx context.Context, lobby types.Lobby, user ethcommon.Address, signature ethcommon.Hash) (err error) {
	err = em.redisClient.HSet(ctx, genLobbySignatureMap(lobby.Id), user.String(), signature).Err()
	return
}
func (em EthMatchStorage) GetProposalSignatures(ctx context.Context, lobbyId string) (signatures types.LobbySignatures, err error) {
	signaturesCmd := em.redisClient.HGetAll(ctx, genLobbySignatureMap(lobbyId))
	if err = signaturesCmd.Err(); err != nil {
		return
	}
	for addr, sig := range signaturesCmd.Val() {
		signatures[ethcommon.HexToAddress(addr)] = ethcommon.HexToHash(sig)
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
