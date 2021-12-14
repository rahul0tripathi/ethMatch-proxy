package eth

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethMatch/proxy/common"
	"github.com/ethMatch/proxy/config"
	game "github.com/ethMatch/proxy/eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"math/big"
)

var (
	gameContract          *game.Contracts
	client                *ethclient.Client
	ChainId               *big.Int
	transactor            *bind.TransactOpts
	OperatorPublicAddress ethcommon.Address
)

func InitETHClient() {
	var err error
	client, err = ethclient.Dial(config.ENV.EthNode)
	if err != nil {
		common.Logger.Error("failed to create new ethclient", zap.Error(err))
	}
	gameContract, err = game.NewContracts(config.ENV.ContractAddress, client)
	if err != nil {
		common.Logger.Error("failed to generate contract binding", zap.Error(err))
	}
	ChainId, err = client.ChainID(context.Background())
	config.ENV.ChainId = ChainId
	if err != nil {
		common.Logger.Error("failed to get chainId", zap.Error(err))
	}
	//load private key
	var pk *ecdsa.PrivateKey
	pk, err = crypto.HexToECDSA(config.ENV.OperatorPrivateKey)
	if err != nil {
		common.Logger.Error("failed to load private key", zap.Error(err))
	}
	transactor, err = bind.NewKeyedTransactorWithChainID(pk, ChainId)
	if err != nil {
		common.Logger.Error("failed to generate transactor from private key", zap.Error(err))
	}
	OperatorPublicAddress = transactor.From
	transactor.GasFeeCap = big.NewInt(18000000000)
	transactor.GasTipCap = big.NewInt(18000000000)
}

func ContractWatcher(sessionChan chan string, result chan string) {
	lobbySink := make(chan *game.ContractsLobbyGenerated)
	resultSink := make(chan *game.ContractsLobbyResult)
	_, err := gameContract.WatchLobbyGenerated(nil, lobbySink)
	if err != nil {
		common.Logger.Error("failed to add lobby event subscription", zap.Error(err))
	}
	_, err = gameContract.WatchLobbyResult(nil, resultSink)
	if err != nil {
		common.Logger.Error("failed to add result event subscription", zap.Error(err))
	}
	for {
		select {
		case l := <-lobbySink:
			go func() {
				select {
				case sessionChan <- l.Id:
					common.Logger.Debug("new lobby event", zap.String("id", l.Id))
				default:
					common.Logger.Error("failed to publish lobby to sessionChan", zap.String("id", l.Id))
				}
			}()
		case r := <-resultSink:
			go func() {
				select {
				case result <- r.Id:
					common.Logger.Debug("new result event", zap.String("id", r.Id))
				default:
					common.Logger.Error("failed to publish result to resultChan", zap.String("id", r.Id))
				}
			}()
		}

	}

}
