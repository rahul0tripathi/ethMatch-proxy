package db

import (
	"encoding/json"
	"fmt"
	"github.com/ethMatch/proxy/config"
	"github.com/ethMatch/proxy/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	db *leveldb.DB
)

func PLAYERLOBBIES(id ethcommon.Address) []byte {
	return []byte(fmt.Sprintf("PLAYER:%s", id.String()))
}
func SESSION(id string) []byte {
	return []byte(fmt.Sprintf("LOBBY:%s", id))
}
func InitDB() (err error) {
	db, err = leveldb.OpenFile(config.ENV.Db, nil)
	if err != nil {
		return
	}
	return
}
func Close() {
	db.Close()
}
func GetSession(session string) (s types.GameSession, err error) {
	var val []byte
	val, err = db.Get(SESSION(session), nil)
	err = json.Unmarshal(val, &s)
	return
}

func SnapshotSession(session *types.GameSession) (err error) {
	var val []byte
	val, err = json.Marshal(*session)
	if err != nil {
		return
	}
	err = db.Put(SESSION(session.LobbyId), val, nil)
	return
}

func GetPlayerLobbies(id ethcommon.Address) (s types.PlayerLobbies, err error) {
	var val []byte
	val, err = db.Get(PLAYERLOBBIES(id), nil)
	err = json.Unmarshal(val, &s)
	return
}

func AddPlayerLobby(id string, player ethcommon.Address) (err error) {
	var val []byte
	var l types.PlayerLobbies
	val, err = db.Get(PLAYERLOBBIES(player), nil)
	err = json.Unmarshal(val, &l)
	//if err != nil {
	//	return
	//}

	l.Lobbies = append(l.Lobbies, id)
	val, err = json.Marshal(l)
	if err != nil {
		return
	}
	err = db.Put(PLAYERLOBBIES(player), val, nil)
	return
}
