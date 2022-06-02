package postgresql

import (
	"database/sql"

	domainRepository "github.com/ymatzki/go-clean-architecture/domain/repository"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domainRepository.Repository {
	return repository{
		db,
	}
}

func (r repository) Greet() string {
	// TODO: implement
	return "fix me and select data from database"
}
