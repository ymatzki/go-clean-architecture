package postgresql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(mock sqlmock.Sqlmock)
		want    string
		err     string
	}{
		{
			"pass",
			func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"word"}).AddRow("hello")
				mock.ExpectQuery("SELECT word FROM greeting").WillReturnRows(rows)
			},
			"hello",
			"",
		},
		{
			"failed - get some error",
			func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT word FROM greeting").WillReturnError(errors.New("some error"))
			},
			"",
			": some error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrenge
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not ecpected when opening a stub database connection", err)
			}
			defer db.Close()

			tt.prepare(mock)
			r := NewRepository(db)

			// act
			got, err := r.Greet()

			// assert
			if got != tt.want {
				t.Errorf("got %s: want %s", got, tt.want)
			}
			if tt.err != "" {
				if err == nil {
					t.Errorf("got %+v: want %+v", err, tt.err)
				}
				if err.Error() != tt.err {
					t.Errorf("got %s: want %s", err, tt.err)
				}
			}
		})
	}
}
