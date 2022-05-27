package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	usecase "github.com/ymatzki/go-clean-architecture/usecase/mock"
)

func TestHello(t *testing.T) {
	type fields struct {
		usecase func(u *usecase.MockUsecase)
		writer  *httptest.ResponseRecorder
		request *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"hello",
			fields{
				func(u *usecase.MockUsecase) {
					u.EXPECT().Hello().Return("Hello, world!\n")
				},
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodGet, "/", nil),
			},
			"Hello, world!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			u := usecase.NewMockUsecase(ctrl)
			tt.fields.usecase(u)
			h := NewHandler(u)

			// call
			h.Get(tt.fields.writer, tt.fields.request)
			res := tt.fields.writer.Result()
			defer res.Body.Close()

			// check
			got, _ := ioutil.ReadAll(res.Body)
			if string(got) != tt.want {
				t.Errorf("got %s: want %s", got, tt.want)
			}
		})
	}
}
