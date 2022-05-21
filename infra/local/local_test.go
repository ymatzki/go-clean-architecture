package local_test

import (
	"testing"

	"github.com/golang/mock/gomock"

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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			r := local.NewRepository()

			got := r.Greet()
			if got != tt.want {
				t.Errorf("got %s: want %s", got, tt.want)
			}
		})
	}
}
