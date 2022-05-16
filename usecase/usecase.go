package usecase

type usecase struct{}

type Usecase interface {
	Hello() string
}

// get hello usecase
func NewUseCase() usecase {
	return usecase{}
}
