package postgresql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(mock sqlmock.Sqlmock)
		want    string
	}{
		{
			"hello",
			func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"word"}).AddRow("hello")
				mock.ExpectQuery("SELECT word FROM greeting").WillReturnRows(rows)
			},
			"hello",
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
			got := r.Greet()

			// assert
			if got != tt.want {
				t.Errorf("got %s: want %s", got, tt.want)
			}
		})
	}
}
