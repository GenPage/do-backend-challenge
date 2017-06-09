package manager

import (
	"errors"
	"sync"

	"github.com/genpage/do-backend-challenge/lib/model"
)

var (
	errAlreadyExists    = errors.New("Already Exists\n")
	errDoesNotExist     = errors.New("Does Not Exist\n")
	errMissingDepedency = errors.New("Missing Depedency\n")
	errDepedencyFound   = errors.New("Depedency Found\n")
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

func (m *manager) Index(p *model.Package) error {
	m.packagesMx.Lock()
	defer m.packagesMx.Unlock()

	//Check if package already in map
	if _, ok := m.packages[p.Name]; ok {
		return errAlreadyExists
	}

	//Make sure package has depedencies in map
	if len(p.Depedencies) > 0 {
		for _, dep := range p.Depedencies {
			if _, ok := m.packages[dep]; !ok {
				return errMissingDepedency
			}
		}
	}

	//Insert package into map
	m.packages[p.Name] = p
	return nil
}

func (m *manager) Remove(p *model.Package) error {
	m.packagesMx.Lock()
	defer m.packagesMx.Unlock()

	//Check if package already in map
	if _, ok := m.packages[p.Name]; !ok {
		return nil
	}

	//Check for any remaining dependencies
	for _, pack := range m.packages {
		for _, dep := range pack.Depedencies {
			if dep == p.Name {
				return errDepedencyFound
			}
		}
	}

	delete(m.packages, p.Name)
	return nil
}

func (m *manager) Query(p *model.Package) error {
	m.packagesMx.Lock()
	defer m.packagesMx.Unlock()

	if _, ok := m.packages[p.Name]; !ok {
		return errDoesNotExist
	}

	return nil
}
