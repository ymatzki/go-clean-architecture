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

type greeting struct {
	word string
}

// FIXME: handle error
func (r repository) Greet() string {
	rows, err := r.db.Query("select word from greeting")
	if err != nil {
		return err.Error()
	}
	defer rows.Close()

	var word string
	for rows.Next() {
		rows.Scan(&word)
	}
	if err = rows.Err(); err != nil {
		return err.Error()
	}

	return word
}
