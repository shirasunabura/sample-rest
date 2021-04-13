// Package usecase ...
package usecase

import (
	"context"
	"fmt"
	"log"

	"app/domain/model"
	"app/domain/repository"
)

// UserUseCase ...
type UserUseCase interface {
	FindUser(context.Context, int64) (*model.User, int64, error)
	CreateUser(context.Context, *model.UserCreater) (*model.User, error)
	DeleteUser(context.Context, int64) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

// NewUserUseCase
func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

// FindUser ユーザー取得
func (uc userUseCase) FindUser(ctx context.Context, id int64) (*model.User, int64, error) {
	u, err := uc.userRepository.FindUser(ctx, id)
	if err != nil {
		return nil, 0, err
	}
	return u, 1, nil
}

// CreateUser ユーザー作成
func (uc userUseCase) CreateUser(ctx context.Context, u *model.UserCreater) (*model.User, error) {
	// 既存ユーザー確認
	us, err := uc.userRepository.FindUserByEmail(ctx, u.Mail)
	if err != nil {
		return nil, err
	}
	if 0 < us.ID {
		log.Printf("User already Exist : %v", us)
		return nil, fmt.Errorf("User already Exist : %v", us)
	}
	cu, err := uc.userRepository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return cu, nil
}

// DeleteUser ユーザー無効化
func (uc userUseCase) DeleteUser(ctx context.Context, id int64) error {
	err := uc.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
