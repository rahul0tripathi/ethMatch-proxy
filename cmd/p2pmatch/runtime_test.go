package main

import (
	"context"
	"github.com/ethMatch/proxy/api"
	"github.com/ethMatch/proxy/matchmaker"
	"github.com/ethMatch/proxy/storage"
	"github.com/ethMatch/proxy/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"testing"
	"time"
)

func TestMatchMakingLoop(t *testing.T) {
	t.Log("initializing newBasicMatchMaker with config", "ticker", time.Second*10, "minPlayers", 1, "maxPlayer", 1)
	basicMatchMaker := matchmaker.NewBasicMatchMaker(time.Second*10, 1, 1)
	t.Log("Adding sample ticket")
	id, err := basicMatchMaker.AddTicketToQueue(types.Ticket{
		OperatorAddress: ethcommon.Address{},
		Player:          ethcommon.HexToAddress(""),
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
			for player := range lobby.PlayerTickets {
				err := api.Ws.Write(player, api.ProposalEvent{Propsal: lobby})
				if err != nil {
					t.Log(err, "failed to write lobby to socket")
				}
			}
		}
		t.Log("Getting user Proposals")
		proposal, _ := store.GetUserLobbyProposals(context.Background(), ethcommon.Address{})
		t.Log("fetched currentProposals", proposal)
		break
	}
}
