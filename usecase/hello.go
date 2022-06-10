package usecase

// echo hello string
func (u usecase) Hello() string {
	greeting, err := u.repo.Greet()
	if err != nil {
		return "" // FIXME: handle error
	}
	return greeting
}
