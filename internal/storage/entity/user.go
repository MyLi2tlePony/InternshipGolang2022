package entity

type User interface {
	GetID() int
	GetBalance() int
}

type user struct {
	id      int
	balance int
}

func NewUser(id, balance int) *user {
	return &user{
		id:      id,
		balance: balance,
	}
}

func (u user) GetID() int {
	return u.id
}

func (u user) GetBalance() int {
	return u.balance
}
