package AuthUseCase

import (
	"fmt"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/auth"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"gitgub.com/diploma-mppr/backend_mppr/tools"
	"github.com/pkg/errors"
)

type UseCaseAuth struct {
	Repo auth.Repository
}

func NewUseCaseAuth(TaskRepo auth.Repository) *UseCaseAuth {
	return &UseCaseAuth{
		Repo: TaskRepo,
	}
}

// func (u *UseCaseAuth) Register(User *models.UserJson) (*models.ResponseUserJson, error) {
//	User1, err := u.Repo.GetUser(&models.UserDB{Username: User.Username, Password: tools.GetMD5Hash(User.Password)})
//	if err == nil && User1.Username == User.Username {
//		return nil, errors.Errorf("такой пользователь уже существует")
//	}
//
//	User2, err := u.Repo.CreateUser(&models.UserDB{Username: User.Username, Password: tools.GetMD5Hash(User.Password)})
//	if err != nil {
//		return nil, err
//	}
//
//	if User.Username != User2.Username || tools.GetMD5Hash(User.Password) != User2.Password {
//		return nil, errors.Errorf("некорректаня работа чего то там")
//	}
//
//	return &models.ResponseUserJson{Id: User2.Id, Username: User2.Username, ColdStart: User2.ColdStart}, nil
//}

func (u *UseCaseAuth) Register(User *models.UserJson) (*models.ResponseUserJson, error) {
	User1, err := u.Repo.GetUser(&models.UserDB{Username: User.Username, Password: tools.GetMD5Hash(User.Password)})
	if err == nil && User1.Username == User.Username {
		return nil, errors.Errorf("такой пользователь уже существует")
	}

	User2, err := u.Repo.CreateUser(&models.UserDB{Username: User.Username, Password: User.Password})
	if err != nil {
		return nil, err
	}

	if User.Username != User2.Username || User.Password != User2.Password {
		return nil, errors.Errorf("некорректаня работа чего то там")
	}

	return &models.ResponseUserJson{Id: User2.Id, Username: User2.Username}, nil
}

func (u *UseCaseAuth) Login(User *models.UserJson) (*models.ResponseUserJson, error) {
	User1, err := u.Repo.GetUser(&models.UserDB{Username: User.Username, Password: tools.GetMD5Hash(User.Password)})
	if err != nil {
		return nil, err
	}
	if User1.Username == User.Username && User1.Password == tools.GetMD5Hash(User.Password) {
		return &models.ResponseUserJson{Id: User1.Id, Username: User1.Username}, nil
	}
	return nil, errors.Errorf("username || password не верно")
}

func (u *UseCaseAuth) GetUserById(id models.UserId) (models.ResponseUserJson, error) {
	user, err := u.Repo.GetUserById(id)
	if err != nil {
		return models.ResponseUserJson{}, err
	}

	user1 := models.ResponseUserJson{Id: user.Id, Username: user.Username}
	fmt.Println(user)
	return user1, nil
}
