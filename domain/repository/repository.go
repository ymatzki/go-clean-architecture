//go:generate mockgen -source ${GOFILE} -destination ./mock/${GOFILE} -package ${GOPACKAGE}
package repository

type Repository interface {
	Greet() (string, error)
}
