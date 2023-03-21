// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LoginHandlerFunc turns a function with the right signature into a login handler
type LoginHandlerFunc func(LoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LoginHandlerFunc) Handle(params LoginParams) middleware.Responder {
	return fn(params)
}

// LoginHandler interface for that can handle valid login params
type LoginHandler interface {
	Handle(LoginParams) middleware.Responder
}

// NewLogin creates a new http.Handler for the login operation
func NewLogin(ctx *middleware.Context, handler LoginHandler) *Login {
	return &Login{Context: ctx, Handler: handler}
}

/*
	Login swagger:route POST /auth/login Auth login

# Login User

Login User
*/
type Login struct {
	Context *middleware.Context
	Handler LoginHandler
}

func (o *Login) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewLoginParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// LoginBody login body
//
// swagger:model LoginBody
type LoginBody struct {

	// password
	// Required: true
	// Min Length: 5
	Password *string `json:"password"`

	// username
	// Required: true
	// Min Length: 2
	Username *string `json:"username"`
}

// Validate validates this login body
func (o *LoginBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	if err := validate.MinLength("data"+"."+"password", "body", *o.Password, 5); err != nil {
		return err
	}

	return nil
}

func (o *LoginBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	if err := validate.MinLength("data"+"."+"username", "body", *o.Username, 2); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this login body based on context it is used
func (o *LoginBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LoginBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginBody) UnmarshalBinary(b []byte) error {
	var res LoginBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// LoginOKBody login o k body
//
// swagger:model LoginOKBody
type LoginOKBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// expired at
	ExpiredAt string `json:"expired_at,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this login o k body
func (o *LoginOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this login o k body based on context it is used
func (o *LoginOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LoginOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginOKBody) UnmarshalBinary(b []byte) error {
	var res LoginOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
