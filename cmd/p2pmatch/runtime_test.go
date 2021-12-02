package main

import (
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/orgs/ethMatch/p2pmatch/matchmaker"
	"github.com/orgs/ethMatch/p2pmatch/storage"
	"github.com/orgs/ethMatch/p2pmatch/types"
	"testing"
	"time"
)

func TestMatchMakingLoop(t *testing.T) {
	t.Log("initializing newBasicMatchMaker with config", "ticker", time.Second*10, "minPlayers", 1, "maxPlayer", 1)
	basicMatchMaker := matchmaker.NewBasicMatchMaker(time.Second*10, 1, 1)
	t.Log("Adding sample ticket")
	id, err := basicMatchMaker.AddTicketToQueue(types.Ticket{
		OperatorAddress: ethcommon.Address{},
		Player:          ethcommon.Address{},
		EntryFee:        20,
		OperatorsShare:  1,
		Status:          0,
		Rank:            100,
	})
	if err != nil {
		t.Error("Failed to add ticket to pool", err)
	}
	store := storage.NewEthMatchStorage()
	t.Log("ticketId", id.String())
	ticket, err := store.GetTicketById(context.Background(), id)
	if err != nil {
		t.Error("failed to get ticket", err)
	}
	t.Log("ticket", ticket)
	tickets, err := store.GetUserTickets(context.Background(), ethcommon.Address{})
	if err != nil {
		t.Error("failed to get new users", err)
	}
	t.Log("tickets-user", tickets)
	t.Log("starting matchmaker")
	go basicMatchMaker.StartMatchMaker()
	for lobby := range basicMatchMaker.GetLobbyStream() {
		if lobby.Id != (ethcommon.Hash{}).String() {
			err := store.NewProposal(context.Background(), lobby)
			if err != nil {
				t.Error("failed store new proposal", err)
			}
		}
		t.Log("Getting user Proposals")
		proposal, _ := store.GetUserLobbyProposals(context.Background(), ethcommon.Address{})
		t.Log("fetched currentProposals", proposal)
		break
	}
}
