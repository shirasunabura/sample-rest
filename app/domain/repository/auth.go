// Package repository ...
package repository

import (
	"context"

	"app/domain/model"
)

// AuthRepository ...
type AuthRepository interface {
	Login(context.Context, string, string) (*model.User, error)
}
