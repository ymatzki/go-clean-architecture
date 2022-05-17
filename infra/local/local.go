package local

import (
	domainRepository "github.com/ymatzki/go-clean-architecture/domain/repository"
)

type repository struct{}

func NewLocalRepository() domainRepository.Repository {
	return repository{}
}

func (r repository) Greet() string {
	return "Hello, world!\n"
}
