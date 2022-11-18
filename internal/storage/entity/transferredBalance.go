package entity

import "time"

type TransferredBalance interface {
	GetSrcUserID() int
	GetDstUserID() int
	GetCreateDate() time.Time
	GetAmount() int
}

type transferredBalance struct {
	srcUserID  int
	dstUserID  int
	createDate time.Time
	amount     int
}

func NewTransferredBalance(srcUserID, dstUserID, amount int, createDate time.Time) *transferredBalance {
	return &transferredBalance{
		srcUserID:  srcUserID,
		dstUserID:  dstUserID,
		createDate: createDate,
		amount:     amount,
	}
}

func (t *transferredBalance) GetSrcUserID() int {
	return t.srcUserID
}

func (t *transferredBalance) GetDstUserID() int {
	return t.dstUserID
}

func (t *transferredBalance) GetCreateDate() time.Time {
	return t.createDate
}

func (t *transferredBalance) GetAmount() int {
	return t.amount
}
