package entity

import "time"

type CancelledBalance interface {
	GetUserID() int
	GetOrderID() int
	GetServiceID() int
	GetCreateDate() time.Time
	GetAmount() int
}

type cancelledBalance struct {
	userID     int
	orderID    int
	serviceID  int
	createDate time.Time
	amount     int
}

func NewCancelledBalance(userID, orderID, serviceID, amount int, createDate time.Time) *cancelledBalance {
	return &cancelledBalance{
		userID:     userID,
		orderID:    orderID,
		serviceID:  serviceID,
		createDate: createDate,
		amount:     amount,
	}
}

func (c *cancelledBalance) GetUserID() int {
	return c.userID
}

func (c *cancelledBalance) GetOrderID() int {
	return c.orderID
}

func (c *cancelledBalance) GetServiceID() int {
	return c.serviceID
}

func (c *cancelledBalance) GetCreateDate() time.Time {
	return c.createDate
}

func (c *cancelledBalance) GetAmount() int {
	return c.amount
}
