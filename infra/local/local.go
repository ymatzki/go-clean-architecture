package local

import (
	domainRepository "github.com/ymatzki/go-clean-architecture/domain/repository"
)

type repository struct{}

func NewRepository() domainRepository.Repository {
	return repository{}
}

func (r repository) Greet() string {
	return "Hello, world!\n"
}
