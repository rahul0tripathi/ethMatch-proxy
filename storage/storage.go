package storage

import (
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/orgs/ethMatch/p2pmatch/types"
)

type Storage interface {
	AddTicket(context.Context, types.Ticket) error
	RemoveTicket(context.Context, uuid.UUID)
	GetTicketById(context.Context, uuid.UUID) (types.Ticket, error)
	GetUserTickets(context.Context, ethcommon.Address) ([]types.Ticket, error)
	GetUserLobbyProposals(context.Context, ethcommon.Address) ([]types.Lobby, error)
	AddSignature(context.Context, types.Lobby, ethcommon.Address, ethcommon.Hash) error
	GetProposalSignatures(context.Context, string) (types.LobbySignatures, error)
	GetProposal(context.Context, string) (types.Lobby, error)
	NewProposal(context.Context, types.Lobby) error
}
