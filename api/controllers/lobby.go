package controllers

import (
	"encoding/json"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"time"

	"github.com/ethMatch/proxy/common"
	"github.com/ethMatch/proxy/storage"
	"github.com/ethMatch/proxy/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type (
	SignProposalRequest struct {
		LobbyId   string `json:"lobby_id"`
		Signature string `json:"signature"`
	}
)

func getProposalHash(proposal types.Lobby, player ethcommon.Address) (hash []byte) {
	var err error
	hash, err = common.NewSignedDataV4(proposal.PlayerTickets[player].String(), proposal)
	if err != nil {
		common.Logger.Error("Failed to generate signed data", zap.Error(err))
	}
	return
}
func GetUserLobbyProposals(w http.ResponseWriter, r *http.Request) {
	if addr := chi.URLParam(r, PlayerAddr); addr != "" {
		proposals, err := storage.CommonStorage.GetUserLobbyProposals(r.Context(), ethcommon.HexToAddress(addr))
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, common.NewResponse(http.StatusInternalServerError, err.Error(), struct{}{}))
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, common.NewResponse(http.StatusOK, http.StatusText(http.StatusOK), proposals))
	}
}

func SignProposal(w http.ResponseWriter, r *http.Request) {
	if addr := chi.URLParam(r, PlayerAddr); addr != "" {
		player := ethcommon.HexToAddress(addr)
		var body SignProposalRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil && body.LobbyId != "" && body.Signature != "" {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, common.NewResponse(http.StatusBadRequest, "invalid payload", struct{}{}))
		}
		proposal, err := storage.CommonStorage.GetProposal(r.Context(), body.LobbyId)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, common.NewResponse(http.StatusInternalServerError, "failed to fetch proposal", struct{}{}))
		}
		if proposal.Id != "" {
			h := getProposalHash(proposal, player)
			verified := common.VerifySig(player, body.Signature, h)
			if !verified {
				render.Status(r, http.StatusOK)
				render.JSON(w, r, common.NewResponse(http.StatusForbidden, "failed to verify signature", struct{}{}))
			} else {
				var existingSig ethcommon.Hash
				existingSig, err = storage.CommonStorage.GetSignature(r.Context(), proposal, player)
				if existingSig == (ethcommon.Hash{}) {
					err = storage.CommonStorage.AddSignature(r.Context(), proposal, player, ethcommon.HexToHash(body.Signature))
					if err != nil {
						render.Status(r, http.StatusInternalServerError)
						render.JSON(w, r, common.NewResponse(http.StatusInternalServerError, err.Error(), struct{}{}))
					} else {
						render.Status(r, http.StatusOK)
						render.JSON(w, r, common.NewResponse(http.StatusOK, "added signature to proposal waiting for others to sign", struct {
							LobbyId   string    `json:"lobby_id"`
							TimeStamp time.Time `json:"timestamp"`
						}{
							LobbyId:   proposal.Id,
							TimeStamp: time.Now(),
						}))
					}
				} else {
					render.Status(r, http.StatusOK)
					render.JSON(w, r, common.NewResponse(http.StatusInternalServerError, "signature already exists", struct{}{}))
				}
			}
		}
	}
}
