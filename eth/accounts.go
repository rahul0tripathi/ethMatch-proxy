package eth

import (
	"github.com/ethMatch/proxy/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

func GetAccountBalance(address ethcommon.Address) *big.Int {
	balances, err := gameContract.AccountBalances(nil, address)
	if err != nil {
		return big.NewInt(0)
	}
	return balances
}

func NewSessionFromChain(id string) (session types.GameSession, err error) {
	lobby, err := gameContract.Lobbies(nil, id)
	if err != nil {
		return
	}
	players, err := gameContract.GetLobbyPlayers(nil, id)
	if err != nil {
		return
	}
	//if !lobby.IsCompleted {
	session = types.GameSession{
		LobbyId:    id,
		MaxPlayers: len(players),
		LobbyReady: false,
		Timeout:    time.Unix(lobby.ExpireAt.Int64()/1000, lobby.ExpireAt.Int64()%1000),
		GameLog:    "",
		JoinedPlayers: map[ethcommon.Address]struct {
			Sig      []byte
			IsWinner bool
		}{},
		AllowedPlayers: players,
	}
	//}
	return
}
