package storage

import (
	"context"
	"errors"
	"time"

	"github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/storage/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrNotExist = errors.New("the record does not exist")

type storage struct {
	connString string
	conn       *pgxpool.Pool
}

func New(connString string) *storage {
	return &storage{
		connString: connString,
	}
}

func (s *storage) Connect(ctx context.Context) (err error) {
	s.conn, err = pgxpool.Connect(ctx, s.connString)
	return err
}

func (s *storage) ExistsUser(ctx context.Context, userID int) (bool, error) {
	sql := "SELECT 1 FROM Users WHERE ID = $1;"

	result, err := s.conn.Query(ctx, sql, userID)
	if err != nil {
		return false, err
	}

	return result.Next(), nil
}

func (s *storage) ExistsReservedBalance(ctx context.Context, userID, orderID, serviceID int) (bool, error) {
	sql := "SELECT 1 FROM ReservedBalance WHERE UserID = $1 AND OrderID = $2 AND ServiceID = $3;"

	result, err := s.conn.Query(ctx, sql, userID, orderID, serviceID)
	if err != nil {
		return false, err
	}

	return result.Next(), nil
}

func (s *storage) ExistsReservedBalanceHistory(ctx context.Context, userID, orderID, serviceID int) (bool, error) {
	sql := "SELECT 1 FROM ReservedBalanceHistory WHERE UserID = $1 AND OrderID = $2 AND ServiceID = $3;"

	result, err := s.conn.Query(ctx, sql, userID, orderID, serviceID)
	if err != nil {
		return false, err
	}

	return result.Next(), nil
}

func (s *storage) ExistsCancelledBalance(ctx context.Context, userID, orderID, serviceID int) (bool, error) {
	sql := "SELECT * FROM CancelledBalance WHERE UserID = $1 AND OrderID = $2 AND ServiceID = $3;"

	result, err := s.conn.Query(ctx, sql, userID, orderID, serviceID)
	if err != nil {
		return false, err
	}

	return result.Next(), nil
}

func (s *storage) ExistsConfirmedBalance(ctx context.Context, userID, orderID, serviceID int) (bool, error) {
	sql := "SELECT * FROM ConfirmedBalance WHERE UserID = $1 AND OrderID = $2 AND ServiceID = $3;"

	result, err := s.conn.Query(ctx, sql, userID, orderID, serviceID)
	if err != nil {
		return false, err
	}

	return result.Next(), nil
}

func (s *storage) CreateUser(ctx context.Context, user entity.User) error {
	sql := "CALL CreateUser($1, $2);"
	_, err := s.conn.Exec(ctx, sql, user.GetID(), user.GetBalance())
	return err
}

func (s *storage) CreateReservedBalance(ctx context.Context, e entity.ReservedBalance) error {
	sql := "CALL CreateReservedBalance($1, $2, $3, $4, $5);"
	_, err := s.conn.Exec(ctx, sql, e.GetUserID(), e.GetOrderID(), e.GetServiceID(), e.GetCreateDate(), e.GetAmount())
	return err
}

func (s *storage) CreateReservedBalanceHistory(ctx context.Context, e entity.ReservedBalanceHistory) error {
	sql := "CALL CreateReservedBalanceHistory($1, $2, $3, $4, $5);"
	_, err := s.conn.Exec(ctx, sql, e.GetUserID(), e.GetOrderID(), e.GetServiceID(), e.GetCreateDate(), e.GetAmount())
	return err
}

func (s *storage) CreateCancelledBalance(ctx context.Context, e entity.CancelledBalance) error {
	sql := "CALL CreateCancelledBalance($1, $2, $3, $4, $5);"
	_, err := s.conn.Exec(ctx, sql, e.GetUserID(), e.GetOrderID(), e.GetServiceID(), e.GetCreateDate(), e.GetAmount())
	return err
}

func (s *storage) CreateConfirmedBalance(ctx context.Context, e entity.ConfirmedBalance) error {
	sql := "CALL CreateConfirmedBalance($1, $2, $3, $4, $5);"
	_, err := s.conn.Exec(ctx, sql, e.GetUserID(), e.GetOrderID(), e.GetServiceID(), e.GetCreateDate(), e.GetAmount())
	return err
}

func (s *storage) CreateReplenishedBalance(ctx context.Context, e entity.ReplenishedBalance) error {
	sql := "CALL CreateReplenishedBalance($1, $2, $3);"
	_, err := s.conn.Exec(ctx, sql, e.GetUserID(), e.GetCreateDate(), e.GetAmount())
	return err
}

func (s *storage) CreateTransferredBalance(ctx context.Context, e entity.TransferredBalance) error {
	sql := "CALL CreateTransferredBalance($1, $2, $3, $4);"
	_, err := s.conn.Exec(ctx, sql, e.GetSrcUserID(), e.GetDstUserID(), e.GetCreateDate(), e.GetAmount())
	return err
}

func (s *storage) UpdateUserBalance(ctx context.Context, userID, amount int) error {
	sql := "CALL UpdateUserBalance($1, $2);"
	_, err := s.conn.Exec(ctx, sql, userID, amount)
	return err
}

func (s *storage) DeleteUser(ctx context.Context, userID int) error {
	sql := "CALL DeleteUser($1);"
	_, err := s.conn.Exec(ctx, sql, userID)
	return err
}

func (s *storage) DeleteReservedBalance(ctx context.Context, userID, orderID, serviceID int) error {
	sql := "CALL DeleteReservedBalance($1, $2, $3);"
	_, err := s.conn.Exec(ctx, sql, userID, orderID, serviceID)
	return err
}

func (s *storage) DeleteReservedBalanceHistory(ctx context.Context, userID, orderID, serviceID int) error {
	sql := "CALL DeleteReservedBalanceHistory($1, $2, $3)"
	_, err := s.conn.Exec(ctx, sql, userID, orderID, serviceID)
	return err
}

func (s *storage) DeleteCancelledBalance(ctx context.Context, userID, orderID, serviceID int) error {
	sql := "CALL DeleteCancelledBalance($1, $2, $3);"
	_, err := s.conn.Exec(ctx, sql, userID, orderID, serviceID)
	return err
}

func (s *storage) DeleteConfirmedBalance(ctx context.Context, userID, orderID, serviceID int) error {
	sql := "CALL DeleteConfirmedBalance($1, $2, $3);"
	_, err := s.conn.Exec(ctx, sql, userID, orderID, serviceID)
	return err
}

func (s *storage) DeleteReplenishedBalance(ctx context.Context, userID int, createDate time.Time) error {
	sql := "CALL DeleteReplenishedBalance($1, $2);"
	_, err := s.conn.Exec(ctx, sql, userID, createDate)
	return err
}

func (s *storage) DeleteTransferredBalance(ctx context.Context, srcUserID, dstUserID int, CreateDate time.Time) error {
	sql := "CALL DeleteTransferredBalance($1, $2, $3);"
	_, err := s.conn.Exec(ctx, sql, srcUserID, dstUserID, CreateDate)
	return err
}

func (s *storage) SelectUser(ctx context.Context, userID int) (entity.User, error) {
	sql := "SELECT ID, Balance FROM Users WHERE ID = $1;"

	result, err := s.conn.Query(ctx, sql, userID)
	if err != nil {
		return nil, err
	}

	if !result.Next() {
		return nil, ErrNotExist
	}

	id, balance := 0, 0
	err = result.Scan(&id, &balance)
	if err != nil {
		return nil, err
	}

	return entity.NewUser(id, balance), nil
}

func (s *storage) SelectReservedBalance(ctx context.Context, userID, orderID, serviceID int) (entity.ReservedBalance, error) {
	sql := `
	SELECT 
		UserID, OrderID, ServiceID, CreateDate, Amount 
	FROM 
		ReservedBalance 
	WHERE 
		UserID = $1 AND OrderID = $2 AND ServiceID = $3;`

	result, err := s.conn.Query(ctx, sql, userID, orderID, serviceID)
	if err != nil {
		return nil, err
	}

	if !result.Next() {
		return nil, ErrNotExist
	}

	userID, orderID, serviceID, amount, createDate := 0, 0, 0, 0, time.Time{}
	err = result.Scan(&userID, &orderID, &serviceID, &createDate, &amount)
	if err != nil {
		return nil, err
	}

	return entity.NewReservedBalance(userID, orderID, serviceID, amount, createDate), nil
}

func (s *storage) SelectReservedBalanceHistory(ctx context.Context, userID, orderID, serviceID int) (entity.ReservedBalanceHistory, error) {
	sql := `
	SELECT 
		UserID, OrderID, ServiceID, CreateDate, Amount 
	FROM 
		ReservedBalanceHistory 
	WHERE 
		UserID = $1 AND OrderID = $2 AND ServiceID = $3;`

	result, err := s.conn.Query(ctx, sql, userID, orderID, serviceID)
	if err != nil {
		return nil, err
	}

	if !result.Next() {
		return nil, ErrNotExist
	}

	userID, orderID, serviceID, amount, createDate := 0, 0, 0, 0, time.Time{}
	err = result.Scan(&userID, &orderID, &serviceID, &createDate, &amount)
	if err != nil {
		return nil, err
	}

	return entity.NewReservedBalanceHistory(userID, orderID, serviceID, amount, createDate), nil
}

func (s *storage) SelectCancelledBalance(ctx context.Context, userID, orderID, serviceID int) (entity.CancelledBalance, error) {
	sql := `
	SELECT 
		UserID, OrderID, ServiceID, CreateDate, Amount 
	FROM 
		CancelledBalance
	WHERE 
		UserID = $1 AND OrderID = $2 AND ServiceID = $3;`

	result, err := s.conn.Query(ctx, sql, userID, orderID, serviceID)
	if err != nil {
		return nil, err
	}

	if !result.Next() {
		return nil, ErrNotExist
	}

	userID, orderID, serviceID, amount, createDate := 0, 0, 0, 0, time.Time{}
	err = result.Scan(&userID, &orderID, &serviceID, &createDate, &amount)
	if err != nil {
		return nil, err
	}

	return entity.NewCancelledBalance(userID, orderID, serviceID, amount, createDate), nil
}

func (s *storage) SelectConfirmedBalance(ctx context.Context, userID, orderID, serviceID int) (entity.ConfirmedBalance, error) {
	sql := `
	SELECT 
		UserID, OrderID, ServiceID, CreateDate, Amount 
	FROM 
		ConfirmedBalance
	WHERE 
		UserID = $1 AND OrderID = $2 AND ServiceID = $3;`

	result, err := s.conn.Query(ctx, sql, userID, orderID, serviceID)
	if err != nil {
		return nil, err
	}

	if !result.Next() {
		return nil, ErrNotExist
	}

	userID, orderID, serviceID, amount, createDate := 0, 0, 0, 0, time.Time{}
	err = result.Scan(&userID, &orderID, &serviceID, &createDate, &amount)
	if err != nil {
		return nil, err
	}

	return entity.NewConfirmedBalance(userID, orderID, serviceID, amount, createDate), nil
}

func (s *storage) SelectReplenishedBalance(ctx context.Context, userID int) ([]entity.ReplenishedBalance, error) {
	sql := `
	SELECT 
		UserID, CreateDate, Amount 
	FROM 
		ReplenishedBalance
	WHERE 
		UserID = $1;`

	result, err := s.conn.Query(ctx, sql, userID)
	if err != nil {
		return nil, err
	}

	replenishes := make([]entity.ReplenishedBalance, 0)

	for result.Next() {
		userID, amount, createDate := 0, 0, time.Time{}

		err = result.Scan(&userID, &createDate, &amount)
		if err != nil {
			return nil, err
		}

		replenishes = append(replenishes, entity.NewReplenishedBalance(userID, amount, createDate))
	}

	return replenishes, nil
}

func (s *storage) SelectTransferredBalance(ctx context.Context, userID int) ([]entity.TransferredBalance, error) {
	sql := `
	SELECT 
		SrcUserID, DstUserID, CreateDate, Amount 
	FROM 
		TransferredBalance
	WHERE 
		SrcUserID = $1 OR DstUserID = $1;`

	result, err := s.conn.Query(ctx, sql, userID)
	if err != nil {
		return nil, err
	}

	transfers := make([]entity.TransferredBalance, 0)

	for result.Next() {
		srcUserID, dstUserID, amount, createDate := 0, 0, 0, time.Time{}

		err = result.Scan(&srcUserID, &dstUserID, &createDate, &amount)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, entity.NewTransferredBalance(srcUserID, dstUserID, amount, createDate))
	}

	return transfers, nil
}

func (s *storage) GetConfirmedBalanceReport(ctx context.Context, month, year int) ([]entity.Report, error) {
	sql := `
		SELECT 
			ServiceID, sum(Amount)
		FROM 
			ConfirmedBalance
		WHERE 
			EXTRACT(MONTH FROM CreateDate) = $1 AND  EXTRACT(YEAR FROM CreateDate) = $2
		GROUP BY 
			ServiceID;`

	result, err := s.conn.Query(ctx, sql, month, year)
	if err != nil {
		return nil, err
	}

	transfers := make([]entity.Report, 0)

	for result.Next() {
		serviceID, balance := 0, 0

		err = result.Scan(&serviceID, &balance)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, entity.NewReport(serviceID, balance))
	}

	return transfers, nil
}
