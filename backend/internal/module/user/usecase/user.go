package usecase

import "database/sql"

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserUsecase interface {
	GetUsers() ([]User, error)
}

type userUsecase struct {
	db *sql.DB
}

func NewUserUsecase(db *sql.DB) UserUsecase {
	return &userUsecase{db: db}
}

func (u *userUsecase) GetUsers() ([]User, error) {
	rows, err := u.db.Query(`
		SELECT id, email, role
		FROM users
		ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
