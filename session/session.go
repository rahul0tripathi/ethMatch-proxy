package session

import (
	"errors"
	"github.com/ethMatch/proxy/common"
	"github.com/ethMatch/proxy/db"
	"github.com/ethMatch/proxy/eth"
	"github.com/ethMatch/proxy/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/zap"
	"sync"
)

var (
	currentSessions map[string]*types.GameSession
	io              sync.RWMutex
)

func init() {
	currentSessions = map[string]*types.GameSession{}
	io = sync.RWMutex{}
}
func getSession(id string) *types.GameSession {
	io.Lock()
	defer io.Unlock()
	if session, ok := currentSessions[id]; ok {
		return session
	} else {
		return nil
	}
}
func removeSession(id string) {
	io.Lock()
	defer io.Unlock()
	if session, ok := currentSessions[id]; ok {
		session.Completed = true
		db.SnapshotSession(session)
		delete(currentSessions, id)
	}

}
func addPlayer(address ethcommon.Address, signature []byte, id string) {
	io.Lock()
	defer io.Unlock()
	currentSessions[id].JoinedPlayers[address] = struct {
		Sig      []byte
		IsWinner bool
	}{Sig: signature, IsWinner: false}
	currentSessions[id].JoinedPlayersCount += 1
	if currentSessions[id].MaxPlayers == currentSessions[id].JoinedPlayersCount {
		currentSessions[id].LobbyReady = true
	}
}
func newSession(id string) {
	session, err := eth.NewSessionFromChain(id)
	if err != nil || session.LobbyId == "" {
		common.Logger.Error("failed to get session", zap.Error(err))
	}
	io.Lock()
	defer io.Unlock()
	if _, ok := currentSessions[id]; !ok {
		currentSessions[id] = &session
	} else {
		common.Logger.Error("session already exists")
	}

}
func NewSessionRoutine(sessionChan chan string) {
	for id := range sessionChan {
		newSession(id)
		s := GetSession(id)
		for _, p := range s.AllowedPlayers {
			err := db.AddPlayerLobby(id, p)
			if err != nil {
				zap.Error(err)
			}
		}
	}
}
func NewResultRoutine(result chan string) {
	for id := range result {
		removeSession(id)
	}
}
func AddPlayerToSession(address ethcommon.Address, signature string, id string) (lobby *types.GameSession, err error) {
	session := getSession(id)
	if session == nil {
		newSession(id)
		session = getSession(id)
	}
	if session == nil {
		err = errors.New("session not found")
		return
	}
	allowed := false
	for _, p := range session.AllowedPlayers {
		if p == address {
			allowed = true
			break
		}
	}
	if _, alreadyExists := session.JoinedPlayers[address]; alreadyExists || !allowed {
		err = errors.New("player not allowed/already present in the session")
		return
	}
	var h []byte
	h, err = common.NewSignedJoinKeyV4(id)
	if err != nil {
		return
	}
	verified := common.VerifySig(address, signature, h)
	if !verified {
		err = errors.New("failed to verify signature")
		return
	}
	var sigBytes []byte
	sigBytes, err = hexutil.Decode(signature)
	if err != nil {
		return
	}
	addPlayer(address, sigBytes, id)
	lobby = session
	return

}
func GetSession(id string) (lobby *types.GameSession) {
	newSession(id)
	io.Lock()
	defer io.Unlock()
	lobby = currentSessions[id]
	if lobby == nil {
		newSession(id)
	}
	lobby = currentSessions[id]
	return
}
func SubmitResults(id string, winner ethcommon.Address, gameData interface{}) (err error) {
	if session := getSession(id); session != nil {
		//session.GameLog, err = ipfs.PinJson(gameData)
		//if err != nil {
		//	return
		//}
		//var d []byte
		//if d, err = json.Marshal(gameData); err != nil {
		//	return
		//}
		//h := crypto.Keccak256Hash(d)
		session.Completed = true
		session.JoinedPlayers[winner] = struct {
			Sig      []byte
			IsWinner bool
		}{Sig: session.JoinedPlayers[winner].Sig, IsWinner: true}
		err = db.SnapshotSession(session)
		if err != nil {
			return
		}
		err = eth.SubmitLobbyResult(session, gameData.(string))
		return
	} else {
		return errors.New("invalid session")
	}
}
