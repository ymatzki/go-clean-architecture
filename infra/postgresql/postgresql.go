package postgresql

import domainRepository "github.com/ymatzki/go-clean-architecture/domain/repository"

type repository struct{}

func NewRepository() domainRepository.Repository {
	return repository{
		// TODO: implement
	}
}

func (r repository) Greet() string {
	// TODO: implement
	return ""
}
