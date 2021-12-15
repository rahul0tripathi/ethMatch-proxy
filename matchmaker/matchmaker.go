package matchmaker

import (
	"github.com/ethMatch/proxy/types"
	"github.com/google/uuid"
)

type MatchMaker interface {
	GetLobbyStream() <-chan types.Proposal
	AddTicketToQueue(types.Ticket) (uuid.UUID, error)
	RemoveTicketFromQueue(uuid.UUID) bool
	StartMatchMaker()
}
