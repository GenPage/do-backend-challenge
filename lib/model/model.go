package model

import "fmt"

type Package struct {
	Name        string
	Command     string
	Depedencies []string
}

func (p *Package) String() string {
	return fmt.Sprintf("name: '%v', command: '%v', dependencies: '%v'", p.Name, p.Command, p.Depedencies)
}
