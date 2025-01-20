package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// user repo
func (ur *UserRepo) GetInfoUser() string {
	return "Anh Duy User info"

}
