package postgresql

import (
	"database/sql"

	"github.com/pkg/errors"
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

func (r repository) Greet() (string, error) {
	rows, err := r.db.Query("SELECT word FROM greeting")
	if err != nil {
		return "", errors.Wrap(err, "")
	}
	defer rows.Close()

	var word string
	for rows.Next() {
		rows.Scan(&word)
	}
	if err = rows.Err(); err != nil {
		return "", errors.Wrap(err, "")
	}

	return word, nil
}
