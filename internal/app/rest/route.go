// Package routing
package rest

import (
	"context"
	"net/http"
	"test-be-IMP/gen/models"
	"test-be-IMP/gen/restapi/operations"
	"test-be-IMP/gen/restapi/operations/auth"
	"test-be-IMP/gen/restapi/operations/coolection"
	"test-be-IMP/gen/restapi/operations/health"
	"test-be-IMP/gen/restapi/operations/user"
	"test-be-IMP/internal/app/handlers"
	"test-be-IMP/internal/pkg/errors"

	"github.com/go-openapi/runtime/middleware"
)

func Route(api *operations.TestBeIMPServerAPI, handler handlers.Handler) {
	//  health
	api.HealthHealthHandler = health.HealthHandlerFunc(func(hp health.HealthParams) middleware.Responder {
		return health.NewHealthOK().WithPayload(&models.Success{
			Code:    200,
			Message: "Server UP!!",
		})
	})

	// login
	api.AuthLoginHandler = auth.LoginHandlerFunc(func(lp auth.LoginParams) middleware.Responder {
		token, expiredAt, err := handler.Login(context.Background(), &lp)
		if err != nil {
			errRes := errors.GetError(err)
			return auth.NewLoginDefault(int(errRes.Code())).WithPayload(&models.Error{
				Code:    int64(errRes.Code()),
				Message: errRes.Error(),
			})
		}

		return auth.NewLoginOK().WithPayload(&auth.LoginOKBody{
			Code:      http.StatusOK,
			ExpiredAt: expiredAt,
			Message:   "Success login",
			Token:     token,
		})
	})

	// signup
	api.AuthSignupHandler = auth.SignupHandlerFunc(func(sp auth.SignupParams) middleware.Responder {
		err := handler.Signup(context.Background(), sp)
		if err != nil {
			errRes := errors.GetError(err)
			return auth.NewSignupDefault(int(errRes.Code())).WithPayload(&models.Error{
				Code:    int64(errRes.Code()),
				Message: errRes.Error(),
			})
		}

		return auth.NewSignupOK().WithPayload(&auth.SignupOKBody{
			Code:    http.StatusOK,
			Message: "Success signup",
		})
	})

	// get list user
	api.UserListUserHandler = user.ListUserHandlerFunc(func(lup user.ListUserParams, p *models.Principal) middleware.Responder {
		resp, err := handler.GetListUserWithPagination(context.Background(), lup)
		if err != nil {
			errRes := errors.GetError(err)
			return user.NewListUserDefault(int(errRes.Code())).WithPayload(&models.Error{
				Code:    int64(errRes.Code()),
				Message: errRes.Error(),
			})
		}

		return user.NewListUserOK().WithPayload(resp)
	})

	api.CoolectionListCollectionHandler = coolection.ListCollectionHandlerFunc(func(lcp coolection.ListCollectionParams, p *models.Principal) middleware.Responder {
		resp, err := handler.GetListCollections(context.Background())
		if err != nil {
			errRes := errors.GetError(err)
			return coolection.NewListCollectionDefault(int(errRes.Code())).WithPayload(&models.Error{
				Code:    int64(errRes.Code()),
				Message: errRes.Error(),
			})
		}

		return coolection.NewListCollectionOK().WithPayload(resp.Payload)
	})
}
