package domain

type Dev struct {
	Name    string
	Ammount int
}

func (dev *Dev) NewDev(name string) *Dev {
	return &Dev{
		Name:    name,
		Ammount: 0,
	}
}

func (dev *Dev) Add(ammount int) {
	dev.Ammount += ammount
}
