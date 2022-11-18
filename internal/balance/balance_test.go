//go:build integration

package balance

import (
	"context"
	"fmt"
	"path"
	"reflect"
	"testing"
	"time"

	databaseConfig "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/config/database"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/storage/entity"
	storage "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/storage/postgres"
	"github.com/stretchr/testify/require"
)

var configPath = path.Join("..", "..", "configs", "test", "config.toml")

func TestBalanceTransfer(t *testing.T) {
	dbConf, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		dbConf.Prefix, dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DatabaseName)

	sqlStorage := Storage(storage.New(connString))
	testBalance := New(sqlStorage)

	ctx := context.Background()
	require.Nil(t, sqlStorage.Connect(ctx))

	srcUserBefore := entity.NewUser(10, 1000)
	require.Nil(t, sqlStorage.CreateUser(ctx, srcUserBefore))

	dstUserBefore := entity.NewUser(11, 0)
	require.Nil(t, sqlStorage.CreateUser(ctx, dstUserBefore))

	tb := entity.NewTransferredBalance(srcUserBefore.GetID(), dstUserBefore.GetID(), 1000, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC))

	require.Nil(t, testBalance.BalanceTransfer(ctx, tb))

	srcUserAfter := entity.NewUser(10, 0)
	selectedSrcUser, err := sqlStorage.SelectUser(ctx, srcUserAfter.GetID())
	require.Nil(t, err)
	require.True(t, reflect.DeepEqual(srcUserAfter, selectedSrcUser))

	dstUserAfter := entity.NewUser(11, 1000)
	selectedDstUser, err := sqlStorage.SelectUser(ctx, dstUserAfter.GetID())
	require.Nil(t, err)
	require.True(t, reflect.DeepEqual(dstUserAfter, selectedDstUser))

	selectedTb, err := sqlStorage.SelectTransferredBalance(ctx, srcUserAfter.GetID())
	require.Nil(t, err)
	require.True(t, len(selectedTb) > 0)
	require.True(t, reflect.DeepEqual(selectedTb[0], tb))

	require.Nil(t, sqlStorage.DeleteUser(ctx, selectedSrcUser.GetID()))
	require.Nil(t, sqlStorage.DeleteUser(ctx, selectedDstUser.GetID()))
	require.Nil(t, sqlStorage.DeleteTransferredBalance(ctx, tb.GetSrcUserID(), tb.GetDstUserID(), tb.GetCreateDate()))
}

func TestGetUserBalance(t *testing.T) {
	dbConf, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		dbConf.Prefix, dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DatabaseName)

	sqlStorage := Storage(storage.New(connString))
	testBalance := New(sqlStorage)

	ctx := context.Background()
	require.Nil(t, sqlStorage.Connect(ctx))

	user := entity.NewUser(12, 1000)
	require.Nil(t, sqlStorage.CreateUser(ctx, user))

	userBalance, err := testBalance.GetUserBalance(ctx, user.GetID())
	require.Nil(t, err)
	require.True(t, userBalance == user.GetBalance())

	require.Nil(t, sqlStorage.DeleteUser(ctx, user.GetID()))
}

func TestReplenishBalance(t *testing.T) {
	dbConf, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		dbConf.Prefix, dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DatabaseName)

	sqlStorage := Storage(storage.New(connString))
	testBalance := New(sqlStorage)

	ctx := context.Background()
	require.Nil(t, sqlStorage.Connect(ctx))

	rb := entity.NewReplenishedBalance(15, 100, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC))
	require.Nil(t, testBalance.ReplenishBalance(ctx, rb))

	exist, err := sqlStorage.ExistsUser(ctx, rb.GetUserID())
	require.Nil(t, err)
	require.True(t, exist)

	require.Nil(t, sqlStorage.DeleteUser(ctx, rb.GetUserID()))

	exist, err = sqlStorage.ExistsUser(ctx, rb.GetUserID())
	require.Nil(t, err)
	require.False(t, exist)

	require.Nil(t, sqlStorage.CreateUser(ctx, entity.NewUser(rb.GetUserID(), 0)))
	require.Nil(t, testBalance.ReplenishBalance(ctx, rb))

	exist, err = sqlStorage.ExistsUser(ctx, rb.GetUserID())
	require.Nil(t, err)
	require.True(t, exist)

	require.Nil(t, sqlStorage.DeleteUser(ctx, rb.GetUserID()))

	exist, err = sqlStorage.ExistsUser(ctx, rb.GetUserID())
	require.Nil(t, err)
	require.False(t, exist)
}

func TestReserveBalance(t *testing.T) {
	dbConf, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		dbConf.Prefix, dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DatabaseName)

	sqlStorage := Storage(storage.New(connString))
	testBalance := New(sqlStorage)

	ctx := context.Background()
	require.Nil(t, sqlStorage.Connect(ctx))

	user := entity.NewUser(20, 2000)
	require.Nil(t, sqlStorage.CreateUser(ctx, user))

	rb := entity.NewReservedBalance(user.GetID(), 20, 20, 2000, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC))
	require.Nil(t, testBalance.ReserveBalance(ctx, rb))

	selectUser, err := sqlStorage.SelectUser(ctx, user.GetID())
	require.Nil(t, err)
	require.True(t, selectUser.GetBalance() == user.GetBalance()-rb.GetAmount())

	reservedBalance, err := sqlStorage.SelectReservedBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID())
	require.Nil(t, err)
	require.True(t, reflect.DeepEqual(reservedBalance, rb))

	require.Nil(t, sqlStorage.DeleteReservedBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID()))
	require.Nil(t, sqlStorage.DeleteUser(ctx, user.GetID()))
}

func TestConfirmReservedBalance(t *testing.T) {
	dbConf, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		dbConf.Prefix, dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DatabaseName)

	sqlStorage := Storage(storage.New(connString))
	testBalance := New(sqlStorage)

	ctx := context.Background()
	require.Nil(t, sqlStorage.Connect(ctx))

	user := entity.NewUser(30, 2000)
	require.Nil(t, sqlStorage.CreateUser(ctx, user))

	rb := entity.NewReservedBalance(user.GetID(), 20, 20, 2000, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC))
	require.Nil(t, sqlStorage.CreateReservedBalance(ctx, rb))

	require.Nil(t, testBalance.ConfirmReservedBalance(ctx, rb))
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
	dbConf, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		dbConf.Prefix, dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DatabaseName)

	sqlStorage := Storage(storage.New(connString))
	testBalance := New(sqlStorage)

	ctx := context.Background()
	require.Nil(t, sqlStorage.Connect(ctx))

	user := entity.NewUser(40, 2000)
	require.Nil(t, sqlStorage.CreateUser(ctx, user))

	rb := entity.NewReservedBalance(user.GetID(), 20, 20, 2000, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC))
	require.Nil(t, testBalance.ReserveBalance(ctx, rb))

	selectedUser, err := sqlStorage.SelectUser(ctx, user.GetID())
	require.Nil(t, err)
	require.True(t, selectedUser.GetBalance() == user.GetBalance()-rb.GetAmount())

	require.Nil(t, testBalance.CancelReservedBalance(ctx, rb))
	exist, err := sqlStorage.ExistsReservedBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID())
	require.Nil(t, err)
	require.False(t, exist)

	exist, err = sqlStorage.ExistsCancelledBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID())
	require.Nil(t, err)
	require.True(t, exist)

	exist, err = sqlStorage.ExistsReservedBalanceHistory(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID())
	require.Nil(t, err)
	require.True(t, exist)

	selectedUser, err = sqlStorage.SelectUser(ctx, user.GetID())
	require.Nil(t, err)
	require.True(t, selectedUser.GetBalance() == user.GetBalance())

	require.Nil(t, sqlStorage.DeleteUser(ctx, user.GetID()))
	require.Nil(t, sqlStorage.DeleteCancelledBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID()))
	require.Nil(t, sqlStorage.DeleteReservedBalanceHistory(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID()))
}
