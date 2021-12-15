package storage

import (
	"context"
	"github.com/ethMatch/proxy/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type Storage interface {
	AddTicket(context.Context, types.Ticket) error
	RemoveTicket(context.Context, uuid.UUID)
	GetTicketById(context.Context, uuid.UUID) (types.Ticket, error)
	GetUserTickets(context.Context, ethcommon.Address) (types.Ticket, error)
	GetUserLobbyProposals(context.Context, ethcommon.Address) ([]types.Proposal, error)
	AddSignature(context.Context, types.Proposal, ethcommon.Address, string) error
	GetSignature(context.Context, types.Proposal, ethcommon.Address) (string, error)
	GetProposalSignatures(context.Context, string) (types.LobbySignatures, error)
	GetProposal(context.Context, string) (types.Proposal, error)
	NewProposal(context.Context, types.Proposal) error
}
