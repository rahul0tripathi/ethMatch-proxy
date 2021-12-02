package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

const (
	TicketStatusPending int = iota + 1
	TicketStatusMatched
	TicketStatusCancelled
)

type Ticket struct {
	OperatorAddress common.Address `json:"operator_address"`
	Player          common.Address `json:"player"`
	EntryFee        float64         `json:"entry_fee"`
	OperatorsShare  float64        `json:"operators_share"`
	Status          int
	Rank            uint64
	Id              uuid.UUID `json:"id,omitempty"`
}

type Lobby struct {
	PlayerTickets     map[common.Address]uuid.UUID `json:"players"`
	PlayerSignatures  LobbySignatures              `json:"-"`
	Ranks             map[common.Address]uint64    `json:"ranks"`
	Pool              float64                       `json:"pool"`
	OperatorsShare    float64                      `json:"operators_share"`
	EntryFee          float64                      `json:"entry_fee"`
	Id                string                       `json:"id"`
	ExpireAt          int64                        `json:"expire_at"`
	OperatorSignature common.Hash                  `json:"operator_signature"`
	OperatorAddress   common.Address               `json:"operator_address"`
}

type LobbySignatures map[common.Address]common.Hash
