package api

import (
	"fmt"
	"github.com/ethMatch/proxy/common"
	"github.com/ethMatch/proxy/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"log"
	"net/http"
	"sync"
	"time"
)

type WsHandler struct {
	players       map[string]*websocket.Conn
	io            sync.Mutex
	flushInterval time.Duration
}
type ProposalEvent struct {
	Propsal types.Lobby `json:"proposal"`
}

func WsUpgradeHandler(upgrader websocket.Upgrader, handler *WsHandler, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	addr := r.Header.Get("address")
	if addr != "" && ethcommon.HexToAddress(addr) != (ethcommon.Address{}) {
		handler.AddPlayer(ethcommon.HexToAddress(addr), conn)
	}
}

func (h *WsHandler) FlushConnections() {
	ticker := time.Tick(h.flushInterval)
	for range ticker {

	}
}
func (h *WsHandler) AddPlayer(player ethcommon.Address, conn *websocket.Conn) {
	defer h.io.Unlock()
	h.io.Lock()
	common.Logger.Debug("adding user conn", zap.String("address", player.String()))
	if player != (ethcommon.Address{}) {
		h.players[player.String()] = conn
	}
}
func (h *WsHandler) GetPlayer(player ethcommon.Address) (conn *websocket.Conn) {
	defer h.io.Unlock()
	h.io.Lock()
	if player != (ethcommon.Address{}) {
		conn = h.players[player.String()]
	}
	return
}
func (h *WsHandler) RemovePlayer(player ethcommon.Address) {
	defer h.io.Unlock()
	h.io.Lock()
	if player != (ethcommon.Address{}) {
		if _, ok := h.players[player.String()]; ok {
			delete(h.players, player.String())
		}
	}
}

func (h *WsHandler) Write(player ethcommon.Address, data interface{}) (err error) {
	playerConn := h.GetPlayer(player)
	fmt.Println(playerConn)
	if playerConn == nil {
		return fmt.Errorf("no conn found for user %s", player.String())
	}
	err = playerConn.WriteJSON(data)
	if err != nil {
		return
	}
	return
}
