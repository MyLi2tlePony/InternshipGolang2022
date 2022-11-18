//go:build integration

package internalhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"reflect"
	"testing"
	"time"

	"github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/balance"
	databaseConfig "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/config/database"
	loggerConfig "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/config/logger"
	serverConfig "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/config/server"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/logger"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/storage/entity"
	storage "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/storage/postgres"
	"github.com/stretchr/testify/require"
)

var configPath = path.Join("..", "..", "..", "configs", "test", "config.toml")

func TestBalanceTransfer(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	dbConnectionString := dbConfig.GetConnectionString()
	sqlStorage := balance.Storage(storage.New(dbConnectionString))

	logConfig, err := loggerConfig.New(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	srvConfig, err := serverConfig.New(configPath)
	if err != nil {
		return
	}

	srv := NewServer(logger.New(logConfig), balance.New(sqlStorage), srvConfig.GetHostPort())
	go func() {
		require.Equal(t, http.ErrServerClosed, srv.Start())
	}()

	defer func() {
		require.Nil(t, srv.Stop())
	}()

	httpClient := &http.Client{}

	ctx := context.Background()
	require.Nil(t, sqlStorage.Connect(ctx))

	srcUserBefore := entity.NewUser(110, 1000)
	require.Nil(t, sqlStorage.CreateUser(ctx, srcUserBefore))

	dstUserBefore := entity.NewUser(111, 0)
	require.Nil(t, sqlStorage.CreateUser(ctx, dstUserBefore))

	tb := transferredBalance{
		SrcUserID:  srcUserBefore.GetID(),
		DstUserID:  dstUserBefore.GetID(),
		CreateDate: time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC),
		Amount:     1000,
	}
	jsonUser, err := json.Marshal(tb)
	require.Nil(t, err)

	request, err := http.NewRequest(http.MethodPost, "http://"+srvConfig.GetHostPort()+urlBalanceTransfer, bytes.NewBuffer(jsonUser))
	require.Nil(t, err)
	request = request.WithContext(ctx)

	resp, err := httpClient.Do(request)
	require.Nil(t, err)

	require.Nil(t, resp.Body.Close())

	srcUserAfter := entity.NewUser(110, 0)
	selectedSrcUser, err := sqlStorage.SelectUser(ctx, srcUserAfter.GetID())
	require.Nil(t, err)
	require.True(t, reflect.DeepEqual(srcUserAfter, selectedSrcUser))

	dstUserAfter := entity.NewUser(111, 1000)
	selectedDstUser, err := sqlStorage.SelectUser(ctx, dstUserAfter.GetID())
	require.Nil(t, err)
	require.True(t, reflect.DeepEqual(dstUserAfter, selectedDstUser))

	selectedTb, err := sqlStorage.SelectTransferredBalance(ctx, srcUserAfter.GetID())
	require.Nil(t, err)
	require.True(t, len(selectedTb) > 0)
	require.True(t, selectedTb[0].GetSrcUserID() == tb.GetSrcUserID())
	require.True(t, selectedTb[0].GetAmount() == tb.GetAmount())
	require.True(t, selectedTb[0].GetDstUserID() == tb.GetDstUserID())
	require.True(t, selectedTb[0].GetCreateDate() == tb.GetCreateDate())

	require.Nil(t, sqlStorage.DeleteUser(ctx, selectedSrcUser.GetID()))
	require.Nil(t, sqlStorage.DeleteUser(ctx, selectedDstUser.GetID()))
	require.Nil(t, sqlStorage.DeleteTransferredBalance(ctx, tb.GetSrcUserID(), tb.GetDstUserID(), tb.GetCreateDate()))
}

func TestGetUserBalance(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	dbConnectionString := dbConfig.GetConnectionString()
	sqlStorage := balance.Storage(storage.New(dbConnectionString))

	logConfig, err := loggerConfig.New(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	srvConfig, err := serverConfig.New(configPath)
	if err != nil {
		return
	}

	srv := NewServer(logger.New(logConfig), balance.New(sqlStorage), srvConfig.GetHostPort())
	go func() {
		require.Equal(t, http.ErrServerClosed, srv.Start())
	}()

	defer func() {
		require.Nil(t, srv.Stop())
	}()

	httpClient := &http.Client{}

	ctx := context.Background()
	require.Nil(t, sqlStorage.Connect(ctx))

	u := user{
		ID:      113,
		Balance: 1000,
	}
	require.Nil(t, sqlStorage.CreateUser(ctx, u))

	jsonUser, err := json.Marshal(u)
	require.Nil(t, err)

	request, err := http.NewRequest(http.MethodPost, "http://"+srvConfig.GetHostPort()+urlGetUserBalance, bytes.NewBuffer(jsonUser))
	require.Nil(t, err)
	request = request.WithContext(ctx)

	resp, err := httpClient.Do(request)
	require.Nil(t, err)

	var selectedUser user
	require.Nil(t, json.NewDecoder(resp.Body).Decode(&selectedUser))
	require.Nil(t, resp.Body.Close())

	require.True(t, reflect.DeepEqual(selectedUser.GetBalance(), u.GetBalance()))
	require.True(t, reflect.DeepEqual(selectedUser.GetBalance(), u.GetBalance()))

	require.Nil(t, sqlStorage.DeleteUser(ctx, u.GetID()))
}

func TestReplenishBalance(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	dbConnectionString := dbConfig.GetConnectionString()
	sqlStorage := balance.Storage(storage.New(dbConnectionString))

	logConfig, err := loggerConfig.New(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	srvConfig, err := serverConfig.New(configPath)
	if err != nil {
		return
	}

	srv := NewServer(logger.New(logConfig), balance.New(sqlStorage), srvConfig.GetHostPort())
	go func() {
		require.Equal(t, http.ErrServerClosed, srv.Start())
	}()

	defer func() {
		require.Nil(t, srv.Stop())
	}()

	httpClient := &http.Client{}

	ctx := context.Background()
	require.Nil(t, sqlStorage.Connect(ctx))

	rb := replenishedBalance{
		UserID:     115,
		Amount:     100,
		CreateDate: time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC),
	}

	jsonUser, err := json.Marshal(rb)
	require.Nil(t, err)

	request, err := http.NewRequest(http.MethodPost, "http://"+srvConfig.GetHostPort()+urlReplenishBalance, bytes.NewBuffer(jsonUser))
	require.Nil(t, err)
	request = request.WithContext(ctx)

	resp, err := httpClient.Do(request)
	require.Nil(t, err)

	require.Nil(t, resp.Body.Close())

	exist, err := sqlStorage.ExistsUser(ctx, rb.GetUserID())
	require.Nil(t, err)
	require.True(t, exist)

	require.Nil(t, sqlStorage.DeleteUser(ctx, rb.GetUserID()))

	exist, err = sqlStorage.ExistsUser(ctx, rb.GetUserID())
	require.Nil(t, err)
	require.False(t, exist)

	require.Nil(t, sqlStorage.CreateUser(ctx, entity.NewUser(rb.GetUserID(), 0)))

	jsonUser, err = json.Marshal(rb)
	require.Nil(t, err)

	request, err = http.NewRequest(http.MethodPost, "http://"+srvConfig.GetHostPort()+urlReplenishBalance, bytes.NewBuffer(jsonUser))
	require.Nil(t, err)
	request = request.WithContext(ctx)

	resp, err = httpClient.Do(request)
	require.Nil(t, err)

	require.Nil(t, resp.Body.Close())

	exist, err = sqlStorage.ExistsUser(ctx, rb.GetUserID())
	require.Nil(t, err)
	require.True(t, exist)

	require.Nil(t, sqlStorage.DeleteUser(ctx, rb.GetUserID()))

	exist, err = sqlStorage.ExistsUser(ctx, rb.GetUserID())
	require.Nil(t, err)
	require.False(t, exist)
}

func TestReserveBalance(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	dbConnectionString := dbConfig.GetConnectionString()
	sqlStorage := balance.Storage(storage.New(dbConnectionString))

	logConfig, err := loggerConfig.New(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	srvConfig, err := serverConfig.New(configPath)
	if err != nil {
		return
	}

	srv := NewServer(logger.New(logConfig), balance.New(sqlStorage), srvConfig.GetHostPort())
	go func() {
		require.Equal(t, http.ErrServerClosed, srv.Start())
	}()

	defer func() {
		require.Nil(t, srv.Stop())
	}()

	httpClient := &http.Client{}

	ctx := context.Background()
	require.Nil(t, sqlStorage.Connect(ctx))

	user := entity.NewUser(120, 2000)
	require.Nil(t, sqlStorage.CreateUser(ctx, user))

	rb := reservedBalance{
		UserID:     user.GetID(),
		OrderID:    20,
		ServiceID:  20,
		Amount:     2000,
		CreateDate: time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC),
	}

	jsonUser, err := json.Marshal(rb)
	require.Nil(t, err)

	request, err := http.NewRequest(http.MethodPost, "http://"+srvConfig.GetHostPort()+urlReserveBalance, bytes.NewBuffer(jsonUser))
	require.Nil(t, err)
	request = request.WithContext(ctx)

	resp, err := httpClient.Do(request)
	require.Nil(t, err)

	require.Nil(t, resp.Body.Close())

	selectUser, err := sqlStorage.SelectUser(ctx, user.GetID())
	require.Nil(t, err)
	require.True(t, selectUser.GetBalance() == user.GetBalance()-rb.GetAmount())

	reservedBalance, err := sqlStorage.SelectReservedBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID())
	require.Nil(t, err)
	require.True(t, reservedBalance.GetUserID() == rb.GetUserID())
	require.True(t, reservedBalance.GetOrderID() == rb.GetOrderID())
	require.True(t, reservedBalance.GetServiceID() == rb.GetServiceID())
	require.True(t, reservedBalance.GetAmount() == rb.GetAmount())
	require.True(t, reservedBalance.GetCreateDate() == rb.GetCreateDate())

	require.Nil(t, sqlStorage.DeleteReservedBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID()))
	require.Nil(t, sqlStorage.DeleteUser(ctx, user.GetID()))
}

func TestConfirmReservedBalance(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	dbConnectionString := dbConfig.GetConnectionString()
	sqlStorage := balance.Storage(storage.New(dbConnectionString))

	logConfig, err := loggerConfig.New(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	srvConfig, err := serverConfig.New(configPath)
	if err != nil {
		return
	}

	srv := NewServer(logger.New(logConfig), balance.New(sqlStorage), srvConfig.GetHostPort())
	go func() {
		require.Equal(t, http.ErrServerClosed, srv.Start())
	}()

	defer func() {
		require.Nil(t, srv.Stop())
	}()

	httpClient := &http.Client{}

	ctx := context.Background()
	require.Nil(t, sqlStorage.Connect(ctx))

	user := entity.NewUser(122, 2000)
	require.Nil(t, sqlStorage.CreateUser(ctx, user))

	rb := reservedBalance{
		UserID:     user.GetID(),
		OrderID:    20,
		ServiceID:  20,
		Amount:     2000,
		CreateDate: time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC),
	}

	require.Nil(t, sqlStorage.CreateReservedBalance(ctx, &rb))

	jsonUser, err := json.Marshal(rb)
	require.Nil(t, err)

	request, err := http.NewRequest(http.MethodPost, "http://"+srvConfig.GetHostPort()+urlConfirmReservedBalance, bytes.NewBuffer(jsonUser))
	require.Nil(t, err)
	request = request.WithContext(ctx)

	resp, err := httpClient.Do(request)
	require.Nil(t, err)

	require.Nil(t, resp.Body.Close())

	exist, err := sqlStorage.ExistsReservedBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID())
	require.Nil(t, err)
	require.False(t, exist)

	exist, err = sqlStorage.ExistsConfirmedBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID())
	require.Nil(t, err)
	require.True(t, exist)

	exist, err = sqlStorage.ExistsReservedBalanceHistory(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID())
	require.Nil(t, err)
	require.True(t, exist)

	require.Nil(t, sqlStorage.DeleteUser(ctx, user.GetID()))
	require.Nil(t, sqlStorage.DeleteConfirmedBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID()))
	require.Nil(t, sqlStorage.DeleteReservedBalanceHistory(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID()))
}

func TestCancelReservedBalance(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	dbConnectionString := dbConfig.GetConnectionString()
	sqlStorage := balance.Storage(storage.New(dbConnectionString))

	logConfig, err := loggerConfig.New(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	srvConfig, err := serverConfig.New(configPath)
	if err != nil {
		return
	}

	srv := NewServer(logger.New(logConfig), balance.New(sqlStorage), srvConfig.GetHostPort())
	go func() {
		require.Equal(t, http.ErrServerClosed, srv.Start())
	}()

	defer func() {
		require.Nil(t, srv.Stop())
	}()

	httpClient := &http.Client{}

	ctx := context.Background()
	require.Nil(t, sqlStorage.Connect(ctx))

	user := entity.NewUser(123, 2000)
	require.Nil(t, sqlStorage.CreateUser(ctx, user))

	rb := reservedBalance{
		UserID:     user.GetID(),
		OrderID:    20,
		Amount:     2000,
		CreateDate: time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC),
	}
	require.Nil(t, sqlStorage.CreateReservedBalance(ctx, &rb))

	jsonUser, err := json.Marshal(rb)
	require.Nil(t, err)

	request, err := http.NewRequest(http.MethodPost, "http://"+srvConfig.GetHostPort()+urlCancelReservedBalance, bytes.NewBuffer(jsonUser))
	require.Nil(t, err)
	request = request.WithContext(ctx)

	resp, err := httpClient.Do(request)
	require.Nil(t, err)

	require.Nil(t, resp.Body.Close())

	exist, err := sqlStorage.ExistsReservedBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID())
	require.Nil(t, err)
	require.False(t, exist)

	exist, err = sqlStorage.ExistsCancelledBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID())
	require.Nil(t, err)
	require.True(t, exist)

	exist, err = sqlStorage.ExistsReservedBalanceHistory(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID())
	require.Nil(t, err)
	require.True(t, exist)

	require.Nil(t, sqlStorage.DeleteUser(ctx, user.GetID()))
	require.Nil(t, sqlStorage.DeleteCancelledBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID()))
	require.Nil(t, sqlStorage.DeleteReservedBalanceHistory(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID()))
}
