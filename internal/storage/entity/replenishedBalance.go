package entity

import "time"

type ReplenishedBalance interface {
	GetUserID() int
	GetCreateDate() time.Time
	GetAmount() int
}

type replenishedBalance struct {
	userID     int
	createDate time.Time
	amount     int
}

func NewReplenishedBalance(userID, amount int, createDate time.Time) *replenishedBalance {
	return &replenishedBalance{
		userID:     userID,
		createDate: createDate,
		amount:     amount,
	}
}

func (r *replenishedBalance) GetUserID() int {
	return r.userID
}

func (r *replenishedBalance) GetCreateDate() time.Time {
	return r.createDate
}

func (r *replenishedBalance) GetAmount() int {
	return r.amount
}
