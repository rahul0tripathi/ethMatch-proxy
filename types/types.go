package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"math/big"
	"time"
)

const (
	TicketStatusPending int = iota + 1
	TicketStatusMatched
	TicketStatusCancelled
)

type (
	Ticket struct {
		OperatorAddress common.Address `json:"operator_address"`
		Player          common.Address `json:"player"`
		EntryFee        uint64         `json:"entry_fee"`
		OperatorsShare  uint64         `json:"operators_share"`
		Status          int
		Rank            uint64
		Id              uuid.UUID `json:"id,omitempty"`
	}

	Proposal struct {
		PlayerTickets    map[common.Address]uuid.UUID `json:"players"`
		PlayerSignatures LobbySignatures              `json:"-"`
		Pool             uint64                       `json:"pool"`
		OperatorsShare   uint64                       `json:"operators_share"`
		EntryFee         uint64                       `json:"entry_fee"`
		Id               string                       `json:"id"`
		ExpireAt         int64                        `json:"expire_at"`
		OperatorAddress  common.Address               `json:"operator_address"`
	}

	LobbySignatures map[common.Address]string

	Config struct {
		OperatorPrivateKey string         `yaml:"private_key"`
		ContractAddress    common.Address `yaml:"contract_address"`
		EthNode            string         `yaml:"eth_node"`
		Name               string         `json:"name"`
		Db                 string         `json:"db"`
		Pinata             struct {
			ApiKey    string `yaml:"api_key"`
			ApiSecret string `yaml:"api_secret"`
		} `yaml:"pinata"`
		ProposalExpireDuration string `yaml:"proposal_expire_duration"`
		MaxTicketTTL           string `yaml:"max_ticket_ttl"`
		MMTime                 string `yaml:"mm_time"`
		MaxPlayers             int    `yaml:"max_players"`
		MinPlayers             int    `yaml:"min_players"`
		Redis                  struct {
			Address  string `yaml:"address"`
			Password string `yaml:"password"`
		} `yaml:"redis"`
		ChainId *big.Int
	}

	GameSession struct {
		LobbyId            string    `json:"lobby_id"`
		MaxPlayers         int       `json:"max_players"`
		JoinedPlayersCount int       `json:"-"`
		LobbyReady         bool      `json:"lobby_ready"`
		Timeout            time.Time `json:"timeout"`
		GameLog            string    `json:"game_log"`
		JoinedPlayers      map[common.Address]struct {
			Sig      []byte
			IsWinner bool
		} `json:"-"`
		AllowedPlayers []common.Address `json:"allowed_players"`
		Completed      bool             `json:"completed"`
	}

	PlayerLobbies struct {
		Lobbies []string `json:"lobbies"`
	}
)
