package AuthRepository

import (
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/models"
	"github.com/jackc/pgx"
)

type RepositoryAuth struct {
	DB *pgx.ConnPool
}

func NewRepositoryAuth(db *pgx.ConnPool) *RepositoryAuth {
	return &RepositoryAuth{DB: db}
}

func (r *RepositoryAuth) GetUserById(id models.UserId) (*models.UserDB, error) {
	UserResponse := &models.UserDB{}
	err := r.DB.QueryRow(
		`select "id", "username", "password" from "users" where "id" = $1`,
		id,
	).Scan(&UserResponse.Id, &UserResponse.Username, &UserResponse.Password)
	if err != nil {
		return nil, err
	}
	return UserResponse, nil
}

func (r *RepositoryAuth) GetUser(UserRequest *models.UserDB) (*models.UserDB, error) {
	UserResponse := &models.UserDB{}
	err := r.DB.QueryRow(
		`select "id", "username", "password" from "users" where "username" = $1 limit 1`,
		UserRequest.Username,
	).Scan(&UserResponse.Id, &UserResponse.Username, &UserResponse.Password)
	if err != nil {
		return nil, err
	}
	return UserResponse, nil
}

func (r *RepositoryAuth) CreateUser(UserRequest *models.UserDB) (*models.UserDB, error) {
	UserResponse := &models.UserDB{}
	err := r.DB.QueryRow(
		`INSERT INTO "users" ("username","password") VALUES ($1,$2) RETURNING "id", "username", "password"`,
		UserRequest.Username, UserRequest.Password,
	).Scan(&UserResponse.Id, &UserResponse.Username, &UserResponse.Password)
	if err != nil {
		return nil, err
	}
	return UserResponse, nil
}

func (r *RepositoryAuth) Login(UserRequest *models.UserDB) (*models.UserDB, error) {
	UserResponse := &models.UserDB{}
	err := r.DB.QueryRow(
		`SELECT * FROM "users" WHERE "username" = $1 and "password" = $2`,
		UserRequest.Username, UserRequest.Password,
	).Scan(&UserResponse.Id, &UserResponse.Username, &UserResponse.Password)
	if err != nil {
		return nil, err
	}
	return UserResponse, nil
}
