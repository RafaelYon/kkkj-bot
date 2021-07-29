package repository

import (
	"fmt"

	"github.com/RafaelYon/kkkj-bot/domain"
)

type DevSaver interface {
	Save(dev *domain.Dev) error
}

type DevRetriver interface {
	Get(name string) (*domain.Dev, error)
	GetAll() []*domain.Dev
}

type DevSaveRetriver interface {
	DevSaver
	DevRetriver
}

type Dev struct {
	devs map[string]*domain.Dev
}

func (d *Dev) Get(name string) (*domain.Dev, error) {
	if dev, ok := d.devs[name]; !ok {
		return nil, fmt.Errorf("DEV \"%s\" não encontrado", name)
	} else {
		return dev, nil
	}
}

func (d *Dev) GetAll() (result []*domain.Dev) {
	for _, v := range d.devs {
		result = append(result, v)
	}

	return
}

func (d *Dev) Save(dev *domain.Dev) error {
	if d.devs == nil {
		d.devs = make(map[string]*domain.Dev)
	}

	d.devs[dev.Name] = dev
	return nil
}
