package domain

type Dev struct {
	Name   string
	Amount int
}

func NewDev(name string) *Dev {
	return &Dev{
		Name:   name,
		Amount: 0,
	}
}

func (dev *Dev) Add(amount int) {
	dev.Amount += amount
}
