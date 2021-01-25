package port

import "portdb.io/src/models"

var _ models.PortDomainRepository = &testRepo{}

type testRepo struct {
	port *models.Port
	err  error
}

func (t *testRepo) Store(p *models.Port) error {
	t.port = p
	return t.err
}

func (t *testRepo) Update(p *models.Port) error {
	if t.port != nil && t.port.Code == p.Code {
		t.port = p
		return nil
	}
	return t.err
}

func (t *testRepo) Fetch(code string) (models.Port, error) {
	if t.port != nil && t.port.Code == code {
		return *t.port, nil
	}
	return models.Port{}, t.err
}
