package model

import (
	"errors"
	"fmt"
	"strings"
)

type Package struct {
	Name        string
	Command     string
	Depedencies []string
}

func PackageFromString(input string) (*Package, error) {
	s := strings.Split(input, "|")
	p := &Package{
		Command:     s[0],
		Name:        s[1],
		Depedencies: strings.Split(s[2], ","),
	}

	err := p.validate()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Package) validate() error {
	if p.Name != "" && p.Command != "" {
		return nil
	}
	return errors.New("Invalid name or command")
}

func (p *Package) String() string {
	return fmt.Sprintf("name: '%v', command: '%v', dependencies: '%v'", p.Name, p.Command, p.Depedencies)
}
