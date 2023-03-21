package rest

import (
	"net/http"
	"test-be-IMP/gen/models"
	"test-be-IMP/gen/restapi/operations"
	"test-be-IMP/internal/pkg/errors"
	utilsJwt "test-be-IMP/internal/utils/jwt"
	"time"
)

func parseToken(token string) (*models.Principal, error) {
	secret := utilsJwt.Cfg.JwtSecret
	maker, err := utilsJwt.NewJWTMaker(secret)
	if err != nil {
		return nil, err
	}

	payload, err := maker.VerifyToken(token)
	if err != nil {
		return nil, errors.SetError(http.StatusUnauthorized, "Unauthorized: invalid API key token: %v", err)
	}

	return &models.Principal{
		ExpiredAt: payload.ExpiredAt.Format(time.RFC3339),
		UserID:    payload.UserID,
		Username:  payload.Username,
	}, nil
}

func Authorization(api *operations.TestBeIMPServerAPI) {
	api.HasTokenAuth = func(token string) (*models.Principal, error) {
		return parseToken(token)
	}
}
