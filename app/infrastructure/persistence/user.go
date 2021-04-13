// Package persistence ...
package persistence

import (
	"context"
	"log"
	"time"

	"app/domain/model"
	"app/domain/repository"
	"app/infrastructure/dao"
)

type userPersistence struct {
	Conn *dao.Connection
}

// NewUserPersistence ...
func NewUserPersistence(conn *dao.Connection) repository.UserRepository {
	return &userPersistence{
		Conn: conn,
	}
}

// FindUser : IDからユーザー取得
func (us userPersistence) FindUser(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	u := us.Conn.Read().Table("users").
		Select("*").
		Where("enabled = ?", 1).
		Where("id = ?", id)
	u.Scan(&user)
	return &user, nil
}

// FindUser : IDからユーザー取得
func (us userPersistence) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	u := us.Conn.Read().Table("users").
		Select("*").
		Where("enabled = ?", 1).
		Where("mail = ?", email)
	u.Scan(&user)
	return &user, nil
}

// CreateUser ...
func (us userPersistence) CreateUser(ctx context.Context, uc *model.UserCreater) (*model.User, error) {

	user := model.User{}
	user.Mail = uc.Mail
	user.Password = uc.Password
	user.FamilyName = uc.FamilyName
	user.FirstName = uc.FirstName
	user.FamilyNameKana = uc.FamilyNameKana
	user.FirstNameKana = uc.FirstNameKana
	user.LastLoginAt = time.Now()

	tx := us.Conn.Write().Begin()
	//tx.SingularTable(false)
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &user, nil
}

// ModifyUser ...
func (us userPersistence) ModifyUser(ctx context.Context, uid int64, uc *model.UserCreater) error {

	user := model.User{}
	user.ID = uid
	user.Mail = uc.Mail
	user.Password = uc.Password
	user.FamilyName = uc.FamilyName
	user.FirstName = uc.FirstName
	user.FamilyNameKana = uc.FamilyNameKana
	user.FirstNameKana = uc.FirstNameKana
	user.LastLoginAt = time.Now()

	tx := us.Conn.Write().Begin()
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// DeleteUser ...
func (us userPersistence) DeleteUser(ctx context.Context, id int64) error {
	log.Println(id)
	if err := us.Conn.Write().Table("users").Where("id = ?", id).Update("enabled", "0").Error; err != nil {
		return err
	}
	return nil
}
