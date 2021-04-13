// Package usecase ...
package usecase

import (
	"context"

	"app/domain/model"
	"app/domain/repository"
)

// AuthUseCase ...
type AuthUseCase interface {
	Login(context.Context, string, string) (*model.User, error)
}

type authUseCase struct {
	authRepository repository.AuthRepository
}

// NewAuthUseCase ...
func NewAuthUseCase(br repository.AuthRepository) AuthUseCase {
	return &authUseCase{
		authRepository: br,
	}
}

// FindUser ユーザー取得
func (uc authUseCase) Login(ctx context.Context, email, password string) (user *model.User, err error) {
	user, err = uc.authRepository.Login(ctx, email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
