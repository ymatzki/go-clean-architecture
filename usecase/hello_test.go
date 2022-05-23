package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	repository "github.com/ymatzki/go-clean-architecture/domain/repository/mock"
	"github.com/ymatzki/go-clean-architecture/usecase"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(r *repository.MockRepository)
		want    string
	}{
		{
			"hello",
			func(r *repository.MockRepository) {
				r.EXPECT().Greet().Return("hello")
			},
			"hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			r := repository.NewMockRepository(ctrl)
			tt.prepare(r)

			u := usecase.NewUsecase(r)
			got := u.Hello()
			if got != tt.want {
				t.Errorf("got %s: want %s", got, tt.want)
			}
		})
	}
}
