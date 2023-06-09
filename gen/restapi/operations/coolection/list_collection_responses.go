// Code generated by go-swagger; DO NOT EDIT.

package coolection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"test-be-IMP/gen/models"
)

// ListCollectionOKCode is the HTTP code returned for type ListCollectionOK
const ListCollectionOKCode int = 200

/*
ListCollectionOK Success Get List Collections

swagger:response listCollectionOK
*/
type ListCollectionOK struct {

	/*
	  In: Body
	*/
	Payload []*ListCollectionOKBodyItems0 `json:"body,omitempty"`
}

// NewListCollectionOK creates ListCollectionOK with default headers values
func NewListCollectionOK() *ListCollectionOK {

	return &ListCollectionOK{}
}

// WithPayload adds the payload to the list collection o k response
func (o *ListCollectionOK) WithPayload(payload []*ListCollectionOKBodyItems0) *ListCollectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list collection o k response
func (o *ListCollectionOK) SetPayload(payload []*ListCollectionOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCollectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*ListCollectionOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
ListCollectionDefault Request Error

swagger:response listCollectionDefault
*/
type ListCollectionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListCollectionDefault creates ListCollectionDefault with default headers values
func NewListCollectionDefault(code int) *ListCollectionDefault {
	if code <= 0 {
		code = 500
	}

	return &ListCollectionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list collection default response
func (o *ListCollectionDefault) WithStatusCode(code int) *ListCollectionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list collection default response
func (o *ListCollectionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list collection default response
func (o *ListCollectionDefault) WithPayload(payload *models.Error) *ListCollectionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list collection default response
func (o *ListCollectionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCollectionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
