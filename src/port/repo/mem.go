package repo

import (
	"portdb.io/src/errors"

	"portdb.io/src/models"
)

var _ models.PortDomainRepository = &mem{}

type mem struct {
	ports map[string]*models.Port
}

func (m *mem) Store(p *models.Port) error {
	m.ports[p.Code] = p
	return nil
}

func (m *mem) Update(p *models.Port) error {
	if _, ok := m.ports[p.Code]; ok {
		m.ports[p.Code] = p
	}
	return errors.ErrNotFound
}

func (m *mem) Fetch(code string) (models.Port, error) {
	if v, ok := m.ports[code]; ok {
		return *v, nil
	}
	return models.Port{}, errors.ErrNotFound
}

func New() *mem {
	return &mem{ports: map[string]*models.Port{}}
}
