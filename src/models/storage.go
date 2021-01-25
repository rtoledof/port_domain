package models

type PortDomainRepository interface {
	Store(p *Port) error
	Update(p *Port) error
	Fetch(code string) (Port, error)
}
