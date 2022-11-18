//go:build integration

package storage

import (
	"context"
	"path"
	"reflect"
	"testing"
	"time"

	databaseConfig "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/config/database"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/storage/entity"
	"github.com/stretchr/testify/require"
)

var configPath = path.Join("..", "..", "..", "configs", "test", "config.toml")

func TestUser(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	sqlStorage := New(dbConfig.GetConnectionString())
	ctx := context.Background()
	elements := []entity.User{
		entity.NewUser(1, 100),
		entity.NewUser(2, 200),
	}

	require.Nil(t, sqlStorage.Connect(ctx))

	for _, element := range elements {
		exist, err := sqlStorage.ExistsUser(ctx, element.GetID())
		require.Nil(t, err)
		require.False(t, exist)

		require.Nil(t, sqlStorage.CreateUser(ctx, element))

		exist, err = sqlStorage.ExistsUser(ctx, element.GetID())
		require.Nil(t, err)
		require.True(t, exist)

		createdElement, err := sqlStorage.SelectUser(ctx, element.GetID())
		require.Nil(t, err)
		require.True(t, reflect.DeepEqual(createdElement, element))

		require.Nil(t, sqlStorage.DeleteUser(ctx, element.GetID()))

		exist, err = sqlStorage.ExistsUser(ctx, element.GetID())
		require.Nil(t, err)
		require.False(t, exist)
	}
}

func TestReservedBalance(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	sqlStorage := New(dbConfig.GetConnectionString())
	ctx := context.Background()

	elements := []entity.ReservedBalance{
		entity.NewReservedBalance(1, 2, 3, 4, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC)),
		entity.NewReservedBalance(5, 6, 7, 8, time.Date(2022, 3, 3, 3, 3, 3, 0, time.UTC)),
	}

	require.Nil(t, sqlStorage.Connect(ctx))

	for _, element := range elements {
		exist, err := sqlStorage.ExistsReservedBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.False(t, exist)

		require.Nil(t, sqlStorage.CreateReservedBalance(ctx, element))

		exist, err = sqlStorage.ExistsReservedBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.True(t, exist)

		createdElement, err := sqlStorage.SelectReservedBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.True(t, reflect.DeepEqual(createdElement, element))

		err = sqlStorage.DeleteReservedBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)

		exist, err = sqlStorage.ExistsReservedBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.False(t, exist)
	}
}

func TestReservedBalanceHistory(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	sqlStorage := New(dbConfig.GetConnectionString())
	ctx := context.Background()

	elements := []entity.ReservedBalanceHistory{
		entity.NewReservedBalanceHistory(1, 2, 3, 4, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC)),
		entity.NewReservedBalanceHistory(5, 6, 7, 8, time.Date(2022, 3, 3, 3, 3, 3, 0, time.UTC)),
	}

	err = sqlStorage.Connect(ctx)
	require.Nil(t, err)

	for _, element := range elements {
		exist, err := sqlStorage.ExistsReservedBalanceHistory(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.False(t, exist)

		err = sqlStorage.CreateReservedBalanceHistory(ctx, element)
		require.Nil(t, err)

		exist, err = sqlStorage.ExistsReservedBalanceHistory(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.True(t, exist)

		createdElement, err := sqlStorage.SelectReservedBalanceHistory(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.True(t, reflect.DeepEqual(createdElement, element))

		err = sqlStorage.DeleteReservedBalanceHistory(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)

		exist, err = sqlStorage.ExistsReservedBalanceHistory(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.False(t, exist)
	}
}

func TestCancelledBalance(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	sqlStorage := New(dbConfig.GetConnectionString())
	ctx := context.Background()

	elements := []entity.CancelledBalance{
		entity.NewCancelledBalance(1, 2, 3, 4, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC)),
		entity.NewCancelledBalance(5, 6, 7, 8, time.Date(2022, 3, 3, 3, 3, 3, 0, time.UTC)),
	}

	err = sqlStorage.Connect(ctx)
	require.Nil(t, err)

	for _, element := range elements {
		exist, err := sqlStorage.ExistsCancelledBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.False(t, exist)

		err = sqlStorage.CreateCancelledBalance(ctx, element)
		require.Nil(t, err)

		exist, err = sqlStorage.ExistsCancelledBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.True(t, exist)

		createdElement, err := sqlStorage.SelectCancelledBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.True(t, reflect.DeepEqual(createdElement, element))

		err = sqlStorage.DeleteCancelledBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)

		exist, err = sqlStorage.ExistsCancelledBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.False(t, exist)
	}
}

func TestConfirmedBalance(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	sqlStorage := New(dbConfig.GetConnectionString())
	ctx := context.Background()

	elements := []entity.ConfirmedBalance{
		entity.NewConfirmedBalance(1, 2, 3, 4, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC)),
		entity.NewConfirmedBalance(5, 6, 7, 8, time.Date(2022, 3, 3, 3, 3, 3, 0, time.UTC)),
	}

	err = sqlStorage.Connect(ctx)
	require.Nil(t, err)

	for _, element := range elements {
		exist, err := sqlStorage.ExistsConfirmedBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.False(t, exist)

		err = sqlStorage.CreateConfirmedBalance(ctx, element)
		require.Nil(t, err)

		exist, err = sqlStorage.ExistsConfirmedBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.True(t, exist)

		createdElement, err := sqlStorage.SelectConfirmedBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.True(t, reflect.DeepEqual(createdElement, element))

		err = sqlStorage.DeleteConfirmedBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)

		exist, err = sqlStorage.ExistsConfirmedBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
		require.False(t, exist)
	}
}

func TestReplenishedBalance(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	sqlStorage := New(dbConfig.GetConnectionString())
	ctx := context.Background()

	elements := []entity.ReplenishedBalance{
		entity.NewReplenishedBalance(1, 2, time.Date(2022, 2, 2, 1, 2, 2, 0, time.UTC)),
		entity.NewReplenishedBalance(1, 5, time.Date(2022, 2, 2, 6, 2, 2, 0, time.UTC)),
		entity.NewReplenishedBalance(1, 6, time.Date(2022, 2, 2, 5, 2, 2, 0, time.UTC)),
		entity.NewReplenishedBalance(1, 4, time.Date(2022, 3, 3, 1, 3, 3, 0, time.UTC)),
	}

	err = sqlStorage.Connect(ctx)
	require.Nil(t, err)

	for _, element := range elements {
		err = sqlStorage.CreateReplenishedBalance(ctx, element)
		require.Nil(t, err)
	}

	replenishes, err := sqlStorage.SelectReplenishedBalance(ctx, elements[0].GetUserID())
	require.Nil(t, err)

	require.True(t, reflect.DeepEqual(elements, replenishes))

	for _, element := range replenishes {
		err = sqlStorage.DeleteReplenishedBalance(ctx, element.GetUserID(), element.GetCreateDate())
		require.Nil(t, err)
	}
}

func TestTransferredBalance(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	sqlStorage := New(dbConfig.GetConnectionString())
	ctx := context.Background()

	elements := []entity.TransferredBalance{
		entity.NewTransferredBalance(1, 2, 30, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC)),
		entity.NewTransferredBalance(1, 2, 70, time.Date(2022, 3, 3, 3, 3, 3, 0, time.UTC)),
		entity.NewTransferredBalance(1, 2, 60, time.Date(2022, 3, 3, 3, 3, 3, 0, time.UTC)),
		entity.NewTransferredBalance(1, 2, 80, time.Date(2022, 3, 3, 3, 3, 3, 0, time.UTC)),
		entity.NewTransferredBalance(1, 2, 10, time.Date(2022, 3, 3, 3, 3, 3, 0, time.UTC)),
	}

	err = sqlStorage.Connect(ctx)
	require.Nil(t, err)

	for _, element := range elements {
		err = sqlStorage.CreateTransferredBalance(ctx, element)
		require.Nil(t, err)
	}

	replenishes, err := sqlStorage.SelectTransferredBalance(ctx, elements[0].GetSrcUserID())
	require.Nil(t, err)

	require.True(t, reflect.DeepEqual(elements, replenishes))

	for _, element := range elements {
		err = sqlStorage.DeleteTransferredBalance(ctx, element.GetSrcUserID(), element.GetDstUserID(), element.GetCreateDate())
		require.Nil(t, err)
	}
}

func TestGetConfirmedBalanceReport(t *testing.T) {
	dbConfig, err := databaseConfig.New(configPath)
	require.Nil(t, err)

	sqlStorage := New(dbConfig.GetConnectionString())
	ctx := context.Background()

	elements := []entity.ConfirmedBalance{
		entity.NewConfirmedBalance(1, 7, 60, 100, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC)),
		entity.NewConfirmedBalance(1, 7, 60, 100, time.Date(2022, 3, 3, 3, 3, 3, 0, time.UTC)),
		entity.NewConfirmedBalance(1, 2, 60, 100, time.Date(2022, 3, 3, 3, 3, 3, 0, time.UTC)),
		entity.NewConfirmedBalance(1, 7, 10, 100, time.Date(2022, 3, 3, 3, 3, 3, 0, time.UTC)),
		entity.NewConfirmedBalance(1, 7, 10, 100, time.Date(2022, 3, 3, 3, 3, 3, 0, time.UTC)),
	}

	err = sqlStorage.Connect(ctx)
	require.Nil(t, err)

	for _, element := range elements {
		err = sqlStorage.CreateConfirmedBalance(ctx, element)
		require.Nil(t, err)
	}

	reports := []entity.Report{
		entity.NewReport(60, 200),
		entity.NewReport(10, 200),
	}

	selectedReports, err := sqlStorage.GetConfirmedBalanceReport(ctx, 3, 2022)
	require.Nil(t, err)

	require.True(t, reflect.DeepEqual(reports[0], selectedReports[0]) || reflect.DeepEqual(reports[1], selectedReports[0]))
	require.True(t, reflect.DeepEqual(reports[0], selectedReports[1]) || reflect.DeepEqual(reports[1], selectedReports[1]))

	for _, element := range elements {
		err = sqlStorage.DeleteConfirmedBalance(ctx, element.GetUserID(), element.GetOrderID(), element.GetServiceID())
		require.Nil(t, err)
	}
}
