//go:generate mockgen -source ${GOFILE} -destination ./mock/${GOFILE} -package ${GOPACKAGE}
package usecase

import "github.com/ymatzki/go-clean-architecture/domain/repository"

type Usecase interface {
	Hello() string
}

type usecase struct {
	repo repository.Repository
}

// get hello usecase
func NewUsecase(repo repository.Repository) Usecase {
	return usecase{repo: repo}
}
