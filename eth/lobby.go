package eth

import (
	"fmt"
	"github.com/ethMatch/proxy/common"
	"github.com/ethMatch/proxy/eth/contracts"
	"github.com/ethMatch/proxy/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/zap"
	"math/big"
)

func SubmitLobbyResult(session *types.GameSession, finalState string) error {
	var sigReq []contracts.GamesigReq
	lobby, err := gameContract.Lobbies(nil, session.LobbyId)
	if err != nil {
		return err
	}
	pool := lobby.Pool
	for player, meta := range session.JoinedPlayers {
		var winningAmount *big.Int
		if meta.IsWinner {
			winningAmount = pool
		} else {
			winningAmount = big.NewInt(int64(0))
		}
		sigReq = append(sigReq, contracts.GamesigReq{
			Player:        player,
			WinningAmount: winningAmount,
			InitalState:   meta.Sig,
			FinalState:    finalState,
		})
	}
	txn, err := gameContract.SubmitResults(transactor, sigReq, session.LobbyId, session.GameLog)

	if err != nil {
		return err
	}
	common.Logger.Info("new result submitted", zap.String("txnId", txn.Hash().String()), zap.Uint64("gas fees", txn.Gas()))
	return nil
}

func SubmitNewLobby(lobby types.Proposal) error {
	var tickets []contracts.Gameticket
	for t, v := range lobby.PlayerTickets {
		sig, err := hexutil.Decode(lobby.PlayerSignatures[t])
		fmt.Println(lobby.PlayerSignatures[t], sig, err)
		if err != nil {
			return err
		}
		tickets = append(tickets, contracts.Gameticket{
			Id:        v.String(),
			Player:    t,
			Amount:    big.NewInt(int64(lobby.EntryFee)),
			Signature: sig,
		})
	}
	createLobby, err := gameContract.CreateLobby(transactor, tickets, lobby.Id, big.NewInt(int64(lobby.OperatorsShare)), big.NewInt(int64(600)))
	if err != nil {
		return err
	}
	common.Logger.Info("new lobby generated", zap.String("txnId", createLobby.Hash().String()), zap.Uint64("gas", createLobby.Gas()))
	return nil
}
