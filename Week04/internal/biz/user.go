package biz

type User struct {
	ID     int32
	Name   string
	Avatar string
}

type UserRepo interface {
	GetUser(id int32) User
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetUser(id int32) User {
	return u.repo.GetUser(id)
}
