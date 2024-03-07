package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/abaykerimov/test_kmf/docs"
	"github.com/abaykerimov/test_kmf/internal/domain/entity"
	"github.com/abaykerimov/test_kmf/internal/infrastructure/dependencies"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

type Service interface {
	GetByDate(ctx context.Context, date, code string) ([]*entity.Rate, error)
	Create(ctx context.Context, date string) error
}

type Handler struct {
	DI *dependencies.Container
}

// getRateAndSaveHandler godoc
// @Summary По указанной дате сохранять курсы валют в БД
// @Tags kmf
// @Produce  json
// @Param id path string true "date"
// @Success 200 {object} saveResponse "response"
// @Failure 400
// @Failure 500
// @Router /currency/save/{date} [get]
func (h *Handler) getRateAndSaveHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	dateStr, ok := vars["date"]
	if !ok {
		sendJsonResponse(resp, http.StatusBadRequest, nil)
		return
	}

	err := h.DI.Service.Create(req.Context(), dateStr)
	if err != nil {
		sendJsonResponse(resp, http.StatusInternalServerError, err)
		return
	}

	sendJsonResponse(resp, http.StatusOK, saveResponse{Success: true})
}

// getRateHandler godoc
// @Summary По указанной дате показать курсы валют из БД
// @Tags kmf
// @Produce  json
// @Param id path string true "date"
// @Success 200 {object} rateResponse "response"
// @Failure 400
// @Failure 500
// @Router /currency/save/{date}/{code} [get]
func (h *Handler) getRateHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	dateStr, ok := vars["date"]
	if !ok {
		sendJsonResponse(resp, http.StatusBadRequest, nil)
		return
	}
	codeStr, _ := vars["code"]

	data, err := h.DI.Service.GetByDate(req.Context(), dateStr, codeStr)
	if err != nil {
		sendJsonResponse(resp, http.StatusInternalServerError, err)
		return
	}

	sendJsonResponse(resp, http.StatusOK, GetRateResponses(data))
}

func sendJsonResponse(resp http.ResponseWriter, status int, body interface{}) bool {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)
	var err error

	switch v := body.(type) {
	case nil:
		log.Println("response with nil body")
	case string:
		if v != "" {
			_, err = fmt.Fprint(resp, v)
		}
	case error:
		err = json.NewEncoder(resp).Encode(v.Error())
	case []byte:
		if bt, ok := body.([]byte); ok {
			_, err = resp.Write(bt)
		}
	default:
		err = json.NewEncoder(resp).Encode(body)
	}

	if err != nil {
		log.Fatal("writing json response failed")

		return false
	}

	return true
}

func (h *Handler) Close() {
	h.DI.DB.Close()
}

func (h *Handler) LoadRoutes(router *mux.Router) {
	router.HandleFunc("/currency/save/{date}", h.getRateAndSaveHandler).Methods(http.MethodGet)
	router.HandleFunc("/currency/{date}", h.getRateHandler).Methods(http.MethodGet)
	router.HandleFunc("/currency/{date}/{code:[A-Z]+?}", h.getRateHandler).Methods(http.MethodGet)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
