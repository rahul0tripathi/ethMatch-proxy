package controllers

import (
	"encoding/json"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/orgs/ethMatch/p2pmatch/common"
	"github.com/orgs/ethMatch/p2pmatch/matchmaker"
	"github.com/orgs/ethMatch/p2pmatch/storage"
	"github.com/orgs/ethMatch/p2pmatch/types"
	"net/http"
	"time"
)

const (
	PlayerAddr = "playerAddr"
)

type (
	AddTicketRequest struct {
		EntryFee       uint64 `json:"entry_fee"`
		OperatorsShare uint64 `json:"operators_share"`
	}
	AddTicketResponse struct {
		Id       string        `json:"id"`
		WaitTime time.Duration `json:"wait_time"`
	}
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

func AddNewTicket(w http.ResponseWriter, r *http.Request) {
	if addr := chi.URLParam(r, PlayerAddr); addr != "" {
		var body AddTicketRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil && body.OperatorsShare > 0 && body.EntryFee > 0 {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, common.NewResponse(http.StatusBadRequest, "invalid payload", struct{}{}))
		}
		id, err := matchmaker.CommonMatchMaker.AddTicketToQueue(types.Ticket{
			OperatorAddress: ethcommon.Address{},
			Player:          ethcommon.HexToAddress(addr),
			EntryFee:        body.EntryFee,
			OperatorsShare:  body.OperatorsShare,
			Status:          types.TicketStatusPending,
			Rank:            10,
		})
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, common.NewResponse(http.StatusInternalServerError, err.Error(), struct{}{}))
		} else {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, common.NewResponse(http.StatusOK, http.StatusText(http.StatusOK), common.NewResponse(http.StatusOK, "successfully added ticket to the queue", AddTicketResponse{
				Id:       id.String(),
				WaitTime: matchmaker.CommonMatchMaker.NextRunTime.Sub(time.Now()),
			})))
		}
	}
}
