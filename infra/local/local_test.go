package local_test

import (
	"testing"

	"github.com/ymatzki/go-clean-architecture/infra/local"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name, want string
	}{
		{"greet", "Hello, world!\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrenge
			r := local.NewRepository()

			// act
			got := r.Greet()

			// assert
			if got != tt.want {
				t.Errorf("got %s: want %s", got, tt.want)
			}
		})
	}
}
