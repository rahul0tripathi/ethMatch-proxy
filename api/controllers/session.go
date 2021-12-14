package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ethMatch/proxy/common"
	"github.com/ethMatch/proxy/db"
	"github.com/ethMatch/proxy/session"
	"github.com/ethMatch/proxy/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type (
	AddPlayerToSessionRequest struct {
		Player ethcommon.Address `json:"player"`
		Key    string            `json:"key"`
		Id     string            `json:"id"`
	}
	SessionResultRequest struct {
		GameData string      `json:"game_data"`
		Winner   ethcommon.Address `json:"winner"`
		Id       string            `json:"id"`
	}
)

const (
	Session = "sessionId"
)

func AddPlayerToSession(w http.ResponseWriter, r *http.Request) {
	var body AddPlayerToSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, common.NewResponse(http.StatusBadRequest, "invalid payload", struct{}{}))
	}
	gameSession, err := session.AddPlayerToSession(body.Player, body.Key, body.Id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponse(http.StatusInternalServerError, err.Error(), struct{}{}))
	} else {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, common.NewResponse(http.StatusOK, "added player to gameSession", *gameSession))
	}
}

func SubmitSessionResult(w http.ResponseWriter, r *http.Request) {
	var body SessionResultRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, common.NewResponse(http.StatusBadRequest, err.Error(), struct{}{}))
		return
	}
	fmt.Println(body.GameData)
	err := session.SubmitResults(body.Id, body.Winner, body.GameData)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponse(http.StatusInternalServerError, err.Error(), struct{}{}))
	} else {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, common.NewResponse(http.StatusOK, "successfully submitted result", struct {
		}{}))
	}
}

func GetPlayerSessions(w http.ResponseWriter, r *http.Request) {
	if addr := chi.URLParam(r, PlayerAddr); addr != "" {
		s, err := db.GetPlayerLobbies(ethcommon.HexToAddress(addr))
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, common.NewResponse(http.StatusInternalServerError, err.Error(), struct{}{}))
		} else {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, common.NewResponse(http.StatusOK, "fetched gameSessions", struct {
				sessions []types.GameSession
			}{
				sessions: s,
			}))
		}
	}
}
func GetGameSession(w http.ResponseWriter, r *http.Request) {
	if s := chi.URLParam(r, Session); s != "" {
		lobby := session.GetSession(s)
		if lobby != nil {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, common.NewResponse(http.StatusOK, "fetched gameSessions", *lobby))
		} else {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, common.NewResponse(http.StatusNotFound, "lobby not found", struct{}{}))
		}

	}
}
