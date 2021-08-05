package repository

type Staging struct {
	Amount int
}

func NewStagingRepository() *Staging {
	return &Staging{
		Amount: 0,
	}
}

func (s *Staging) AddAmount() {
	s.Amount += 1
}
