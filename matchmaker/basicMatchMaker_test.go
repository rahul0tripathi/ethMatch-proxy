package matchmaker

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/orgs/ethMatch/p2pmatch/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var ticketId uuid.UUID

func TestNewBasicMatchMaker(t *testing.T) {
	var err error
	req := require.New(t)
	mm := NewBasicMatchMaker(time.Minute, 2, 4)
	ticketId, err = mm.AddTicketToQueue(types.Ticket{
		OperatorAddress: ethcommon.Address{},
		Player:          ethcommon.Address{},
		EntryFee:        20,
		OperatorsShare:  1,
		Status:          0,
		Rank:            100,
	})
	req.Nil(err)
	req.NotEmpty(ticketId)
}

func TestRedisMatchMaker_GetTicketById(t *testing.T) {

}
