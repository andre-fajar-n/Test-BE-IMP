package handlers

import (
	"context"
	"net/http"
	"test-be-IMP/gen/models"
	"test-be-IMP/gen/restapi/operations/auth"
	"test-be-IMP/internal/pkg/errors"
	utilsHash "test-be-IMP/internal/utils/hash"
	utilsJwt "test-be-IMP/internal/utils/jwt"
	"time"

	"gorm.io/gorm"
)

// Login handler contains business logic of login
func (h *handler) Login(ctx context.Context, params *auth.LoginParams) (token, expiredAt string, err error) {
	// find user by username
	var user models.User
	err = h.db.Model(user).WithContext(ctx).Where("username = ?", params.Data.Username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.SetError(http.StatusNotFound, "Username is Invalid")
		}
		return
	}

	hashedPassword := utilsHash.HashSha256(*params.Data.Password)
	if user.Password != hashedPassword {
		err = errors.SetError(http.StatusBadRequest, "Password is Invalid")
		return
	}

	secret := utilsJwt.Cfg.JwtSecret
	maker, err := utilsJwt.NewJWTMaker(secret)
	if err != nil {
		err = errors.SetError(http.StatusBadRequest, err.Error())
		return
	}

	exp := time.Duration(utilsJwt.Cfg.JwtExp) * time.Second
	token, payload, err := maker.CreateToken(user.ID, user.Username, exp)
	if err != nil {
		err = errors.SetError(http.StatusBadRequest, err.Error())
		return
	}

	expiredAt = payload.ExpiredAt.Format(time.RFC3339)

	return
}

// Signup handler contains bussiness logic to signup new user
func (h *handler) Signup(ctx context.Context, params auth.SignupParams) (err error) {
	var existingUser models.User
	err = h.db.Where("username", params.Data.Username).First(&existingUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if existingUser.ID > 0 {
		err = errors.SetError(http.StatusConflict, "username already exist")
		return
	}

	hashPassword := utilsHash.HashSha256(*params.Data.Password)

	req := models.User{
		Fullname:  *params.Data.Fullname,
		Password:  hashPassword,
		Username:  *params.Data.Username,
		CreatedAt: time.Now().UTC().Unix(),
	}
	err = h.db.Create(&req).Error
	if err != nil {
		return
	}

	return
}
