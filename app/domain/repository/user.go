// Package repository ...
package repository

import (
	"context"

	"app/domain/model"
)

// UserRepository ...
type UserRepository interface {
	FindUser(context.Context, int64) (*model.User, error)
	FindUserByEmail(context.Context, string) (*model.User, error)
	CreateUser(context.Context, *model.UserCreater) (*model.User, error)
	ModifyUser(context.Context, int64, *model.UserCreater) error
	DeleteUser(context.Context, int64) error
}
