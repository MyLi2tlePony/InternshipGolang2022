package entity

type Report interface {
	GetServiceID() int
	GetBalance() int
}

type report struct {
	serviceID int
	balance   int
}

func NewReport(serviceID, balance int) *report {
	return &report{
		serviceID: serviceID,
		balance:   balance,
	}
}

func (u report) GetServiceID() int {
	return u.serviceID
}

func (u report) GetBalance() int {
	return u.balance
}
