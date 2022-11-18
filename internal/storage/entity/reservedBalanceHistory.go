package entity

import "time"

type ReservedBalanceHistory interface {
	GetUserID() int
	GetOrderID() int
	GetServiceID() int
	GetCreateDate() time.Time
	GetAmount() int
}

type reservedBalanceHistory struct {
	userID     int
	orderID    int
	serviceID  int
	createDate time.Time
	amount     int
}

func NewReservedBalanceHistory(userID, orderID, serviceID, amount int, createDate time.Time) *reservedBalanceHistory {
	return &reservedBalanceHistory{
		userID:     userID,
		orderID:    orderID,
		serviceID:  serviceID,
		createDate: createDate,
		amount:     amount,
	}
}

func (r *reservedBalanceHistory) GetUserID() int {
	return r.userID
}

func (r *reservedBalanceHistory) GetOrderID() int {
	return r.orderID
}

func (r *reservedBalanceHistory) GetServiceID() int {
	return r.serviceID
}

func (r *reservedBalanceHistory) GetCreateDate() time.Time {
	return r.createDate
}

func (r *reservedBalanceHistory) GetAmount() int {
	return r.amount
}
