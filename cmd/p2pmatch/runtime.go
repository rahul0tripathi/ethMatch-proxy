package main

import (
	"context"
	"github.com/ethMatch/proxy/api"
	"github.com/ethMatch/proxy/common"
	"github.com/ethMatch/proxy/matchmaker"
	"github.com/ethMatch/proxy/storage"
	ethcommon "github.com/ethereum/go-ethereum/common"
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
