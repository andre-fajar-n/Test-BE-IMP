// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"test-be-IMP/gen/models"
)

// ListUserOKCode is the HTTP code returned for type ListUserOK
const ListUserOKCode int = 200

/*
ListUserOK Success Get List User

swagger:response listUserOK
*/
type ListUserOK struct {

	/*
	  In: Body
	*/
	Payload *ListUserOKBody `json:"body,omitempty"`
}

// NewListUserOK creates ListUserOK with default headers values
func NewListUserOK() *ListUserOK {

	return &ListUserOK{}
}

// WithPayload adds the payload to the list user o k response
func (o *ListUserOK) WithPayload(payload *ListUserOKBody) *ListUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user o k response
func (o *ListUserOK) SetPayload(payload *ListUserOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ListUserDefault Request Error

swagger:response listUserDefault
*/
type ListUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListUserDefault creates ListUserDefault with default headers values
func NewListUserDefault(code int) *ListUserDefault {
	if code <= 0 {
		code = 500
	}

	return &ListUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list user default response
func (o *ListUserDefault) WithStatusCode(code int) *ListUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list user default response
func (o *ListUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list user default response
func (o *ListUserDefault) WithPayload(payload *models.Error) *ListUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user default response
func (o *ListUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
