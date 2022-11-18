package entity

import "time"

type ConfirmedBalance interface {
	GetUserID() int
	GetOrderID() int
	GetServiceID() int
	GetCreateDate() time.Time
	GetAmount() int
}

type confirmedBalance struct {
	userID     int
	orderID    int
	serviceID  int
	createDate time.Time
	amount     int
}

func NewConfirmedBalance(userID, orderID, serviceID, amount int, createDate time.Time) *confirmedBalance {
	return &confirmedBalance{
		userID:     userID,
		orderID:    orderID,
		serviceID:  serviceID,
		createDate: createDate,
		amount:     amount,
	}
}

func (c *confirmedBalance) GetUserID() int {
	return c.userID
}

func (c *confirmedBalance) GetOrderID() int {
	return c.orderID
}

func (c *confirmedBalance) GetServiceID() int {
	return c.serviceID
}

func (c *confirmedBalance) GetCreateDate() time.Time {
	return c.createDate
}

func (c *confirmedBalance) GetAmount() int {
	return c.amount
}
