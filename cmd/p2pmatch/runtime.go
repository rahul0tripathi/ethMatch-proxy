package main

import (
	"context"
	"flag"
	"github.com/ethMatch/proxy/api"
	"github.com/ethMatch/proxy/common"
	"github.com/ethMatch/proxy/config"
	"github.com/ethMatch/proxy/db"
	"github.com/ethMatch/proxy/eth"
	"github.com/ethMatch/proxy/matchmaker"
	"github.com/ethMatch/proxy/session"
	"github.com/ethMatch/proxy/storage"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

var (
	configPath  = flag.String("config", "", "path of the config file")
	sessionChan chan string
	resultChan  chan string
)

func main() {
	sessionChan = make(chan string)
	resultChan = make(chan string)
	err := config.InitConfig(*(configPath))
	if err != nil {
		common.Logger.Error("failed to init config", zap.String("path", *(configPath)), zap.Error(err))
	}
	common.Logger.Debug("current ENV", zap.Any("env", *config.ENV))
	err = db.InitDB()
	defer db.Close()
	if err != nil {
		common.Logger.Error("failed to init DB", zap.Error(err))
	}
	go api.RunHTTPServer()
	go session.NewResultRoutine(resultChan)
	go session.NewSessionRoutine(sessionChan)
	eth.InitETHClient()
	go eth.ContractWatcher(sessionChan, resultChan)
	matchmaker.InitMM()
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
