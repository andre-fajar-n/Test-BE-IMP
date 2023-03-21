// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"test-be-IMP/gen/models"
)

// HealthOKCode is the HTTP code returned for type HealthOK
const HealthOKCode int = 200

/*
HealthOK Health Check

swagger:response healthOK
*/
type HealthOK struct {

	/*
	  In: Body
	*/
	Payload *models.Success `json:"body,omitempty"`
}

// NewHealthOK creates HealthOK with default headers values
func NewHealthOK() *HealthOK {

	return &HealthOK{}
}

// WithPayload adds the payload to the health o k response
func (o *HealthOK) WithPayload(payload *models.Success) *HealthOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the health o k response
func (o *HealthOK) SetPayload(payload *models.Success) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HealthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HealthInternalServerErrorCode is the HTTP code returned for type HealthInternalServerError
const HealthInternalServerErrorCode int = 500

/*
HealthInternalServerError Server Error

swagger:response healthInternalServerError
*/
type HealthInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewHealthInternalServerError creates HealthInternalServerError with default headers values
func NewHealthInternalServerError() *HealthInternalServerError {

	return &HealthInternalServerError{}
}

// WithPayload adds the payload to the health internal server error response
func (o *HealthInternalServerError) WithPayload(payload *models.Error) *HealthInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the health internal server error response
func (o *HealthInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HealthInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}