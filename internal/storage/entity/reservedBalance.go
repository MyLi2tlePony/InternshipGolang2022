package entity

import "time"

type ReservedBalance interface {
	GetUserID() int
	GetOrderID() int
	GetServiceID() int
	GetCreateDate() time.Time
	GetAmount() int
}

type reservedBalance struct {
	userID     int
	orderID    int
	serviceID  int
	createDate time.Time
	amount     int
}

func NewReservedBalance(userID, orderID, serviceID, amount int, createDate time.Time) *reservedBalance {
	return &reservedBalance{
		userID:     userID,
		orderID:    orderID,
		serviceID:  serviceID,
		createDate: createDate,
		amount:     amount,
	}
}

func (r *reservedBalance) GetUserID() int {
	return r.userID
}

func (r *reservedBalance) GetOrderID() int {
	return r.orderID
}

func (r *reservedBalance) GetServiceID() int {
	return r.serviceID
}

func (r *reservedBalance) GetCreateDate() time.Time {
	return r.createDate
}

func (r *reservedBalance) GetAmount() int {
	return r.amount
}
