package matchmaker

import (
	"github.com/google/uuid"
	"github.com/orgs/ethMatch/p2pmatch/types"
)

type MatchMaker interface {
	GetLobbyStream() <-chan types.Lobby
	AddTicketToQueue(types.Ticket) (uuid.UUID, error)
	RemoveTicketFromQueue(uuid.UUID) bool
	StartMatchMaker()
}
