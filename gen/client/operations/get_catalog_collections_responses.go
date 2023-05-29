// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"test-be-IMP/gen/models"
)

// GetCatalogCollectionsReader is a Reader for the GetCatalogCollections structure.
type GetCatalogCollectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogCollectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCatalogCollectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCatalogCollectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCatalogCollectionsOK creates a GetCatalogCollectionsOK with default headers values
func NewGetCatalogCollectionsOK() *GetCatalogCollectionsOK {
	return &GetCatalogCollectionsOK{}
}

/*
GetCatalogCollectionsOK describes a response with status code 200, with default header values.

success get data
*/
type GetCatalogCollectionsOK struct {
	Payload *GetCatalogCollectionsOKBody
}

// IsSuccess returns true when this get catalog collections o k response has a 2xx status code
func (o *GetCatalogCollectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get catalog collections o k response has a 3xx status code
func (o *GetCatalogCollectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get catalog collections o k response has a 4xx status code
func (o *GetCatalogCollectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get catalog collections o k response has a 5xx status code
func (o *GetCatalogCollectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get catalog collections o k response a status code equal to that given
func (o *GetCatalogCollectionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCatalogCollectionsOK) Error() string {
	return fmt.Sprintf("[GET /catalog/collections][%d] getCatalogCollectionsOK  %+v", 200, o.Payload)
}

func (o *GetCatalogCollectionsOK) String() string {
	return fmt.Sprintf("[GET /catalog/collections][%d] getCatalogCollectionsOK  %+v", 200, o.Payload)
}

func (o *GetCatalogCollectionsOK) GetPayload() *GetCatalogCollectionsOKBody {
	return o.Payload
}

func (o *GetCatalogCollectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCatalogCollectionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogCollectionsDefault creates a GetCatalogCollectionsDefault with default headers values
func NewGetCatalogCollectionsDefault(code int) *GetCatalogCollectionsDefault {
	return &GetCatalogCollectionsDefault{
		_statusCode: code,
	}
}

/*
GetCatalogCollectionsDefault describes a response with status code -1, with default header values.

Request Error
*/
type GetCatalogCollectionsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get catalog collections default response
func (o *GetCatalogCollectionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get catalog collections default response has a 2xx status code
func (o *GetCatalogCollectionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get catalog collections default response has a 3xx status code
func (o *GetCatalogCollectionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get catalog collections default response has a 4xx status code
func (o *GetCatalogCollectionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get catalog collections default response has a 5xx status code
func (o *GetCatalogCollectionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get catalog collections default response a status code equal to that given
func (o *GetCatalogCollectionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetCatalogCollectionsDefault) Error() string {
	return fmt.Sprintf("[GET /catalog/collections][%d] GetCatalogCollections default  %+v", o._statusCode, o.Payload)
}

func (o *GetCatalogCollectionsDefault) String() string {
	return fmt.Sprintf("[GET /catalog/collections][%d] GetCatalogCollections default  %+v", o._statusCode, o.Payload)
}

func (o *GetCatalogCollectionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCatalogCollectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCatalogCollectionsOKBody get catalog collections o k body
swagger:model GetCatalogCollectionsOKBody
*/
type GetCatalogCollectionsOKBody struct {

	// collections
	Collections []*GetCatalogCollectionsOKBodyCollectionsItems0 `json:"collections"`

	// status
	Status string `json:"status,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this get catalog collections o k body
func (o *GetCatalogCollectionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCollections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCatalogCollectionsOKBody) validateCollections(formats strfmt.Registry) error {
	if swag.IsZero(o.Collections) { // not required
		return nil
	}

	for i := 0; i < len(o.Collections); i++ {
		if swag.IsZero(o.Collections[i]) { // not required
			continue
		}

		if o.Collections[i] != nil {
			if err := o.Collections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getCatalogCollectionsOK" + "." + "collections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getCatalogCollectionsOK" + "." + "collections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get catalog collections o k body based on the context it is used
func (o *GetCatalogCollectionsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCollections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCatalogCollectionsOKBody) contextValidateCollections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Collections); i++ {

		if o.Collections[i] != nil {
			if err := o.Collections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getCatalogCollectionsOK" + "." + "collections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getCatalogCollectionsOK" + "." + "collections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCatalogCollectionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCatalogCollectionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCatalogCollectionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetCatalogCollectionsOKBodyCollectionsItems0 get catalog collections o k body collections items0
swagger:model GetCatalogCollectionsOKBodyCollectionsItems0
*/
type GetCatalogCollectionsOKBodyCollectionsItems0 struct {

	// id
	ID int64 `json:"id,omitempty"`

	// long text
	LongText string `json:"long_text,omitempty"`

	// repositoryid
	Repositoryid string `json:"repositoryid,omitempty"`

	// short text
	ShortText string `json:"short_text,omitempty"`

	// thumbnail
	Thumbnail string `json:"thumbnail,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this get catalog collections o k body collections items0
func (o *GetCatalogCollectionsOKBodyCollectionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get catalog collections o k body collections items0 based on context it is used
func (o *GetCatalogCollectionsOKBodyCollectionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCatalogCollectionsOKBodyCollectionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCatalogCollectionsOKBodyCollectionsItems0) UnmarshalBinary(b []byte) error {
	var res GetCatalogCollectionsOKBodyCollectionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}