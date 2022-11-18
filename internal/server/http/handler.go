package internalhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type handler struct {
	logger   Logger
	app      Balance
	hostPort string
}

func newHandler(logger Logger, app Balance, hostPort string) *handler {
	return &handler{
		logger:   logger,
		app:      app,
		hostPort: hostPort,
	}
}

func (h *handler) BalanceTransfer(w http.ResponseWriter, r *http.Request) {
	var tb transferredBalance
	err := readFromBody(r, &tb)
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tb.CreateDate = time.Now()

	if err := h.app.BalanceTransfer(context.Background(), &tb); err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("status", fmt.Sprint(http.StatusOK))
	w.WriteHeader(http.StatusOK)
}
func (h *handler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	var u user
	err := readFromBody(r, &u)
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.Balance, err = h.app.GetUserBalance(context.Background(), u.GetID())
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = writeToBody(w, u)
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *handler) ReplenishBalance(w http.ResponseWriter, r *http.Request) {
	var tb replenishedBalance
	err := readFromBody(r, &tb)
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tb.CreateDate = time.Now()

	if err := h.app.ReplenishBalance(context.Background(), &tb); err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("status", fmt.Sprint(http.StatusOK))
	w.WriteHeader(http.StatusOK)
}

func (h *handler) ReserveBalance(w http.ResponseWriter, r *http.Request) {
	var tb reservedBalance
	err := readFromBody(r, &tb)
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tb.CreateDate = time.Now()

	if err := h.app.ReserveBalance(context.Background(), &tb); err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("status", fmt.Sprint(http.StatusOK))
	w.WriteHeader(http.StatusOK)
}

func (h *handler) ConfirmReservedBalance(w http.ResponseWriter, r *http.Request) {
	var tb confirmedBalance
	err := readFromBody(r, &tb)
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tb.CreateDate = time.Now()

	if err := h.app.ConfirmReservedBalance(context.Background(), &tb); err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("status", fmt.Sprint(http.StatusOK))
	w.WriteHeader(http.StatusOK)
}

func (h *handler) CancelReservedBalance(w http.ResponseWriter, r *http.Request) {
	var tb cancelledBalance
	err := readFromBody(r, &tb)
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tb.CreateDate = time.Now()

	if err := h.app.CancelReservedBalance(context.Background(), &tb); err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("status", fmt.Sprint(http.StatusOK))
	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetConfirmedBalanceReportLink(w http.ResponseWriter, r *http.Request) {
	var tb confirmedBalanceRecord
	err := readFromBody(r, &tb)
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	report, err := h.app.GetConfirmedBalanceReportLink(context.Background(), tb.GetMonth(), tb.GetYear())
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = writeToBody(w, "http://"+h.hostPort+"/report/file"+report)
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *handler) GetConfirmedBalanceReportFile(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fileName := query.Get("file") + ".csv"

	stream, err := ioutil.ReadFile("./reports/" + fileName)
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b := bytes.NewBuffer(stream)
	w.Header().Set("Content-type", "application/csv")
	w.Header().Set("content-disposition", "attachment; filename=\""+fileName+"\"")

	if _, err = b.WriteTo(w); err != nil {
		_, err = fmt.Fprintf(w, "%s", err)
		if err != nil {
			h.logger.Error(err.Error())
			w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	_, err = w.Write([]byte(""))
	if err != nil {
		h.logger.Error(err.Error())
		w.Header().Add("status", fmt.Sprint(http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func readFromBody(r *http.Request, entity any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, entity); err != nil {
		return err
	}

	return nil
}

func writeToBody(w http.ResponseWriter, entity any) error {
	marshal, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	_, err = w.Write(marshal)
	return err
}
