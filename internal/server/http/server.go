package internalhttp

import (
	"context"
	"fmt"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/storage/entity"
	"net/http"
	"time"
)

type Logger interface {
	Fatal(string)
	Error(string)
	Warn(string)
	Info(string)
	Debug(string)
	Trace(string)
}

type Balance interface {
	BalanceTransfer(ctx context.Context, tb entity.TransferredBalance) error
	GetUserBalance(ctx context.Context, userID int) (int, error)
	ReplenishBalance(ctx context.Context, rb entity.ReplenishedBalance) error
	ReserveBalance(ctx context.Context, rb entity.ReservedBalance) error
	ConfirmReservedBalance(ctx context.Context, cb entity.ConfirmedBalance) error
	CancelReservedBalance(ctx context.Context, cb entity.CancelledBalance) error
	GetConfirmedBalanceReportLink(ctx context.Context, month, year int) (string, error)
}

const (
	urlGetUserBalance                = "/select"
	urlReplenishBalance              = "/replenish"
	urlBalanceTransfer               = "/transfer"
	urlReserveBalance                = "/reserve"
	urlCancelReservedBalance         = "/reserve/cancel"
	urlConfirmReservedBalance        = "/reserve/confirm"
	urlGetConfirmedBalanceReportLink = "/report/link"
	urlGetConfirmedBalanceReportFile = "/report/file"
)

type Server struct {
	app    Balance
	logger Logger

	srv *http.Server
}

func NewServer(logger Logger, app Balance, hostPort string) *Server {
	handler := newHandler(logger, app, hostPort)

	mux := http.NewServeMux()
	mux.HandleFunc(urlGetUserBalance, handler.GetUserBalance)
	mux.HandleFunc(urlReplenishBalance, handler.ReplenishBalance)
	mux.HandleFunc(urlBalanceTransfer, handler.BalanceTransfer)
	mux.HandleFunc(urlReserveBalance, handler.ReserveBalance)
	mux.HandleFunc(urlCancelReservedBalance, handler.CancelReservedBalance)
	mux.HandleFunc(urlConfirmReservedBalance, handler.ConfirmReservedBalance)
	mux.HandleFunc(urlGetConfirmedBalanceReportLink, handler.GetConfirmedBalanceReportLink)
	mux.HandleFunc(urlGetConfirmedBalanceReportFile, handler.GetConfirmedBalanceReportFile)

	middleware := newMiddleware(logger, mux)
	middleware.logging()

	return &Server{
		logger: logger,
		app:    app,
		srv: &http.Server{
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 2 * time.Second,
			IdleTimeout:  2 * time.Second,
			Addr:         hostPort,
			Handler:      middleware.Handler,
		},
	}
}

func (s *Server) Start() error {
	s.logger.Info(fmt.Sprintf("http server listening: %s", s.srv.Addr))

	if err := s.srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	if err := s.srv.Close(); err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}
