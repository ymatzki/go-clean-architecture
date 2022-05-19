//go:generate mockgen -source ${GOFILE} -destination ./mock/${GOFILE} -package ${GOPACKAGE}
package usecase

import "github.com/ymatzki/go-clean-architecture/domain/repository"

type usecase struct {
	repo repository.Repository
}

type Usecase interface {
	Hello() string
}

// get hello usecase
func NewUseCase(repo repository.Repository) usecase {
	return usecase{repo: repo}
}
