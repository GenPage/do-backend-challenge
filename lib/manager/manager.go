package manager

import (
	"sync"

	"github.com/genpage/do-backend-challenge/lib/model"
)

type Manager interface {
	Index(p *model.Package) error
	Remove(p *model.Package) error
	Query(p *model.Package) error
}

type manager struct {
	packages   map[string]*model.Package
	packagesMx sync.Mutex
}

func NewManager() Manager {
	return &manager{
		packages: make(map[string]*model.Package),
	}
}

func (r *manager) Index(p *model.Package) error {
	return nil
}

func (r *manager) Remove(p *model.Package) error {
	return nil
}

func (r *manager) Query(p *model.Package) error {
	return nil
}
