package controllers

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/orgs/ethMatch/p2pmatch/common"
	"github.com/orgs/ethMatch/p2pmatch/storage"
	"net/http"
)

const (
	PlayerAddr = "playerAddr"
)

func GetUserTickets(w http.ResponseWriter, r *http.Request) {
	if addr := chi.URLParam(r, PlayerAddr); addr != "" {
		tickets, err := storage.CommonStorage.GetUserTickets(r.Context(), ethcommon.HexToAddress(addr))
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, common.NewResponse(http.StatusInternalServerError, err.Error(), struct{}{}))
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, common.NewResponse(http.StatusOK, http.StatusText(http.StatusOK), tickets))
	}
}
