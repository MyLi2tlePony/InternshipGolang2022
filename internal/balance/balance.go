package balance

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/storage/entity"
)

type Storage interface {
	Connect(ctx context.Context) (err error)

	ExistsUser(ctx context.Context, userID int) (bool, error)
	ExistsReservedBalance(ctx context.Context, userID, orderID, serviceID int) (bool, error)
	ExistsReservedBalanceHistory(ctx context.Context, userID, orderID, serviceID int) (bool, error)
	ExistsCancelledBalance(ctx context.Context, userID, orderID, serviceID int) (bool, error)
	ExistsConfirmedBalance(ctx context.Context, userID, orderID, serviceID int) (bool, error)

	CreateUser(ctx context.Context, user entity.User) error
	CreateReservedBalance(ctx context.Context, e entity.ReservedBalance) error
	CreateReservedBalanceHistory(ctx context.Context, e entity.ReservedBalanceHistory) error
	CreateCancelledBalance(ctx context.Context, e entity.CancelledBalance) error
	CreateConfirmedBalance(ctx context.Context, e entity.ConfirmedBalance) error
	CreateReplenishedBalance(ctx context.Context, e entity.ReplenishedBalance) error
	CreateTransferredBalance(ctx context.Context, e entity.TransferredBalance) error

	GetConfirmedBalanceReport(ctx context.Context, month, year int) ([]entity.Report, error)
	UpdateUserBalance(ctx context.Context, userID, amount int) error

	DeleteUser(ctx context.Context, userID int) error
	DeleteReservedBalance(ctx context.Context, userID, orderID, serviceID int) error
	DeleteReservedBalanceHistory(ctx context.Context, userID, orderID, serviceID int) error
	DeleteCancelledBalance(ctx context.Context, userID, orderID, serviceID int) error
	DeleteConfirmedBalance(ctx context.Context, userID, orderID, serviceID int) error
	DeleteReplenishedBalance(ctx context.Context, userID int, createDate time.Time) error
	DeleteTransferredBalance(ctx context.Context, srcUserID, dstUserID int, CreateDate time.Time) error

	SelectUser(ctx context.Context, userID int) (entity.User, error)
	SelectReservedBalance(ctx context.Context, userID, orderID, serviceID int) (entity.ReservedBalance, error)
	SelectReservedBalanceHistory(ctx context.Context, userID, orderID, serviceID int) (entity.ReservedBalanceHistory, error)
	SelectCancelledBalance(ctx context.Context, userID, orderID, serviceID int) (entity.CancelledBalance, error)
	SelectConfirmedBalance(ctx context.Context, userID, orderID, serviceID int) (entity.ConfirmedBalance, error)
	SelectReplenishedBalance(ctx context.Context, userID int) ([]entity.ReplenishedBalance, error)
	SelectTransferredBalance(ctx context.Context, userID int) ([]entity.TransferredBalance, error)
}

var (
	ErrNegativeID     = errors.New("negative id error")
	ErrNegativeAmount = errors.New("negative id error")

	ErrNotExistUser     = errors.New("user does not exist error")
	ErrInvalidDstUser   = errors.New("invalid destination id error")
	ErrBalanceNotEnough = errors.New("the balance is not enough for the transfer")

	ErrNotExistReservedBalance  = errors.New("reserved balance does not exist error")
	ErrDuplicateReservedBalance = errors.New("reserved balance is duplicated")
)

type balance struct {
	storage Storage
	mutex   sync.Mutex
}

func New(storage Storage) *balance {
	return &balance{
		storage: storage,
	}
}

func (b *balance) BalanceTransfer(ctx context.Context, tb entity.TransferredBalance) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if tb.GetSrcUserID() < 0 || tb.GetDstUserID() < 0 {
		return ErrNegativeID
	}
	if tb.GetAmount() < 0 {
		return ErrNegativeAmount
	}
	if tb.GetSrcUserID() == tb.GetDstUserID() {
		return ErrInvalidDstUser
	}

	exist, err := b.storage.ExistsUser(ctx, tb.GetSrcUserID())
	if err != nil {
		return err
	}

	if !exist {
		return ErrNotExistUser
	}

	exist, err = b.storage.ExistsUser(ctx, tb.GetDstUserID())
	if err != nil {
		return err
	}

	if !exist {
		return ErrNotExistUser
	}

	selectedSrcUser, err := b.storage.SelectUser(ctx, tb.GetSrcUserID())
	if err != nil {
		return err
	}

	if selectedSrcUser.GetBalance() < tb.GetAmount() {
		return ErrBalanceNotEnough
	}

	return b.balanceTransfer(ctx, tb)
}

func (b *balance) balanceTransfer(ctx context.Context, tb entity.TransferredBalance) (err error) {
	err = b.storage.UpdateUserBalance(ctx, tb.GetSrcUserID(), -tb.GetAmount())
	if err != nil {
		return err
	}

	err = b.storage.UpdateUserBalance(ctx, tb.GetDstUserID(), tb.GetAmount())
	if err != nil {
		return err
	}

	return b.storage.CreateTransferredBalance(ctx, tb)
}

func (b *balance) GetUserBalance(ctx context.Context, userID int) (int, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if userID < 0 {
		return 0, ErrNegativeID
	}

	exist, err := b.storage.ExistsUser(ctx, userID)
	if err != nil {
		return 0, err
	}

	if !exist {
		return 0, ErrNotExistUser
	}

	return b.getUserBalance(ctx, userID)
}

func (b *balance) getUserBalance(ctx context.Context, userID int) (int, error) {
	user, err := b.storage.SelectUser(ctx, userID)
	if err != nil {
		return 0, err
	}

	return user.GetBalance(), nil
}

func (b *balance) ReplenishBalance(ctx context.Context, rb entity.ReplenishedBalance) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if rb.GetUserID() < 0 {
		return ErrNegativeID
	}

	if rb.GetAmount() < 0 {
		return ErrNegativeAmount
	}

	exist, err := b.storage.ExistsUser(ctx, rb.GetUserID())
	if err != nil {
		return err
	}

	if exist {
		err = b.storage.UpdateUserBalance(ctx, rb.GetUserID(), rb.GetAmount())
		if err != nil {
			return err
		}
	} else {
		err = b.storage.CreateUser(ctx, entity.NewUser(rb.GetUserID(), rb.GetAmount()))
		if err != nil {
			return err
		}
	}

	err = b.storage.CreateReplenishedBalance(ctx, rb)
	if err != nil {
		return err
	}

	return nil
}

func (b *balance) ReserveBalance(ctx context.Context, rb entity.ReservedBalance) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if rb.GetUserID() < 0 {
		return ErrNegativeID
	}

	if rb.GetAmount() < 0 {
		return ErrNegativeAmount
	}

	exist, err := b.storage.ExistsUser(ctx, rb.GetUserID())
	if err != nil {
		return err
	}

	if !exist {
		return ErrNotExistUser
	}

	exist, err = b.storage.ExistsReservedBalance(ctx, rb.GetUserID(), rb.GetOrderID(), rb.GetServiceID())
	if err != nil {
		return err
	}

	if exist {
		return ErrDuplicateReservedBalance
	}

	return b.reserveBalance(ctx, rb)
}

func (b *balance) reserveBalance(ctx context.Context, rb entity.ReservedBalance) error {
	err := b.storage.UpdateUserBalance(ctx, rb.GetUserID(), -rb.GetAmount())
	if err != nil {
		return err
	}

	err = b.storage.CreateReservedBalance(ctx, rb)
	if err != nil {
		return err
	}

	return nil
}

func (b *balance) ConfirmReservedBalance(ctx context.Context, cb entity.ConfirmedBalance) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if cb.GetUserID() < 0 {
		return ErrNegativeID
	}

	if cb.GetAmount() < 0 {
		return ErrNegativeAmount
	}

	exist, err := b.storage.ExistsUser(ctx, cb.GetUserID())
	if err != nil {
		return err
	}

	if !exist {
		return ErrNotExistUser
	}

	exist, err = b.storage.ExistsReservedBalance(ctx, cb.GetUserID(), cb.GetOrderID(), cb.GetServiceID())
	if err != nil {
		return err
	}

	if !exist {
		return ErrNotExistReservedBalance
	}

	return b.confirmReservedBalance(ctx, cb)
}

func (b *balance) confirmReservedBalance(ctx context.Context, cb entity.ConfirmedBalance) error {
	rb, err := b.storage.SelectReservedBalance(ctx, cb.GetUserID(), cb.GetOrderID(), cb.GetServiceID())
	if err != nil {
		return err
	}

	err = b.storage.CreateReservedBalanceHistory(ctx, rb)
	if err != nil {
		return err
	}

	err = b.storage.DeleteReservedBalance(ctx, cb.GetUserID(), cb.GetOrderID(), cb.GetServiceID())
	if err != nil {
		return err
	}

	err = b.storage.CreateConfirmedBalance(ctx, cb)
	if err != nil {
		return err
	}

	return nil
}

func (b *balance) CancelReservedBalance(ctx context.Context, cb entity.CancelledBalance) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if cb.GetUserID() < 0 {
		return ErrNegativeID
	}

	if cb.GetAmount() < 0 {
		return ErrNegativeAmount
	}

	exist, err := b.storage.ExistsUser(ctx, cb.GetUserID())
	if err != nil {
		return err
	}

	if !exist {
		return ErrNotExistUser
	}

	exist, err = b.storage.ExistsReservedBalance(ctx, cb.GetUserID(), cb.GetOrderID(), cb.GetServiceID())
	if err != nil {
		return err
	}

	if !exist {
		return ErrNotExistReservedBalance
	}

	return b.cancelReservedBalance(ctx, cb)
}

func (b *balance) cancelReservedBalance(ctx context.Context, cb entity.CancelledBalance) error {
	rb, err := b.storage.SelectReservedBalance(ctx, cb.GetUserID(), cb.GetOrderID(), cb.GetServiceID())
	if err != nil {
		return err
	}

	err = b.storage.CreateReservedBalanceHistory(ctx, rb)
	if err != nil {
		return err
	}

	err = b.storage.DeleteReservedBalance(ctx, cb.GetUserID(), cb.GetOrderID(), cb.GetServiceID())
	if err != nil {
		return err
	}

	err = b.storage.CreateCancelledBalance(ctx, cb)
	if err != nil {
		return err
	}

	err = b.storage.UpdateUserBalance(ctx, cb.GetUserID(), cb.GetAmount())
	if err != nil {
		return err
	}

	return nil
}

func (b *balance) GetConfirmedBalanceReportLink(ctx context.Context, month, year int) (string, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if month < 0 {
		return "", ErrNegativeID
	}

	if year < 0 {
		return "", ErrNegativeID
	}

	reports, err := b.storage.GetConfirmedBalanceReport(ctx, month, year)
	if err != nil {
		return "", err
	}

	file, err := os.Create(fmt.Sprintf("./reports/%04d_%02d.csv", year, month))
	if err != nil {
		return "", err
	}

	defer func() {
		if err := file.Close(); err != nil {
			return
		}
	}()

	for _, report := range reports {
		_, err = file.WriteString(fmt.Sprintf("%d;%d", report.GetServiceID(), report.GetBalance()) + "\n")
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("?file=%04d_%02d", year, month), nil
}
