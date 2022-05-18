package usecase

// echo hello string
func (u usecase) Hello() string {
	return u.repo.Greet()
}
