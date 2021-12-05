package main

import (
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/orgs/ethMatch/p2pmatch/api"
	"github.com/orgs/ethMatch/p2pmatch/common"
	"github.com/orgs/ethMatch/p2pmatch/matchmaker"
	"github.com/orgs/ethMatch/p2pmatch/storage"
	"go.uber.org/zap"
)

func main() {
	go api.RunHTTPServer()
	go matchmaker.CommonMatchMaker.StartMatchMaker()
	for lobby := range matchmaker.CommonMatchMaker.GetLobbyStream() {
		if lobby.Id != (ethcommon.Hash{}).String() {
			err := storage.CommonStorage.NewProposal(context.Background(), lobby)
			if err != nil {
				common.Logger.Error("failed store new proposal", zap.Error(err))
			}
			for player := range lobby.PlayerTickets {
				err := api.Ws.Write(player, api.ProposalEvent{Propsal: lobby})
				if err != nil {
					common.Logger.Error("failed write to socket", zap.Error(err))
				}
			}
		}

	}
}
