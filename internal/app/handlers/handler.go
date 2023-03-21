package handlers

import (
	"context"
	"test-be-IMP/gen/restapi/operations/auth"
	"test-be-IMP/gen/restapi/operations/user"

	"gorm.io/gorm"
)

type (
	Handler interface {
		IAuth
		IUser
	}

	IAuth interface {
		Login(ctx context.Context, params *auth.LoginParams) (token, expiredAt string, err error)
		Signup(ctx context.Context, params auth.SignupParams) (err error)
	}

	IUser interface {
		GetListUserWithPagination(ctx context.Context, params user.ListUserParams) (*user.ListUserOKBody, error)
	}
)

func NewHandler(db *gorm.DB) Handler {
	return &handler{
		db,
	}
}

type (
	handler struct {
		db *gorm.DB
	}
)
