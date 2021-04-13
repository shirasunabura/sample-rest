// Package persistence ...
package persistence

import (
	"context"

	"app/domain/model"
	"app/domain/repository"
	"app/infrastructure/dao"
)

type authPersistence struct {
	Conn *dao.Connection
}

// NewAuthPersistence ...
func NewAuthPersistence(conn *dao.Connection) repository.AuthRepository {
	return &authPersistence{
		Conn: conn,
	}
}

// Login : mail,passwordユーザー取得
func (us authPersistence) Login(ctx context.Context, email, password string) (*model.User, error) {
	var user model.User
	u := us.Conn.Read().Table("users").
		Select("*").
		Where("enabled = ?", 1).
		Where("mail = ?", email).
		Where("password = ?", password)
	u.Scan(&user)
	return &user, nil
}
