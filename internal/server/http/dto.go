package internalhttp

import "time"

type cancelledBalance struct {
	UserID     int
	OrderID    int
	ServiceID  int
	CreateDate time.Time
	Amount     int
}

func (c *cancelledBalance) GetUserID() int {
	return c.UserID
}

func (c *cancelledBalance) GetOrderID() int {
	return c.OrderID
}

func (c *cancelledBalance) GetServiceID() int {
	return c.ServiceID
}

func (c *cancelledBalance) GetCreateDate() time.Time {
	return c.CreateDate
}

func (c *cancelledBalance) GetAmount() int {
	return c.Amount
}

type confirmedBalance struct {
	UserID     int
	OrderID    int
	ServiceID  int
	CreateDate time.Time
	Amount     int
}

func (c *confirmedBalance) GetUserID() int {
	return c.UserID
}

func (c *confirmedBalance) GetOrderID() int {
	return c.OrderID
}

func (c *confirmedBalance) GetServiceID() int {
	return c.ServiceID
}

func (c *confirmedBalance) GetCreateDate() time.Time {
	return c.CreateDate
}

func (c *confirmedBalance) GetAmount() int {
	return c.Amount
}

type replenishedBalance struct {
	UserID     int
	CreateDate time.Time
	Amount     int
}

func (r *replenishedBalance) GetUserID() int {
	return r.UserID
}

func (r *replenishedBalance) GetCreateDate() time.Time {
	return r.CreateDate
}

func (r *replenishedBalance) GetAmount() int {
	return r.Amount
}

type reservedBalance struct {
	UserID     int
	OrderID    int
	ServiceID  int
	CreateDate time.Time
	Amount     int
}

func (r *reservedBalance) GetUserID() int {
	return r.UserID
}

func (r *reservedBalance) GetOrderID() int {
	return r.OrderID
}

func (r *reservedBalance) GetServiceID() int {
	return r.ServiceID
}

func (r *reservedBalance) GetCreateDate() time.Time {
	return r.CreateDate
}

func (r *reservedBalance) GetAmount() int {
	return r.Amount
}

type transferredBalance struct {
	SrcUserID  int
	DstUserID  int
	CreateDate time.Time
	Amount     int
}

func (t *transferredBalance) GetSrcUserID() int {
	return t.SrcUserID
}

func (t *transferredBalance) GetDstUserID() int {
	return t.DstUserID
}

func (t *transferredBalance) GetCreateDate() time.Time {
	return t.CreateDate
}

func (t *transferredBalance) GetAmount() int {
	return t.Amount
}

type user struct {
	ID      int
	Balance int
}

func (u user) GetID() int {
	return u.ID
}

func (u user) GetBalance() int {
	return u.Balance
}

type confirmedBalanceRecord struct {
	Month int
	Year  int
}

func (u confirmedBalanceRecord) GetMonth() int {
	return u.Month
}

func (u confirmedBalanceRecord) GetYear() int {
	return u.Year
}
