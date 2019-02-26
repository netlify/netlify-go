// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netlify/go-client/models"
)

// ListHookTypesReader is a Reader for the ListHookTypes structure.
type ListHookTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListHookTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListHookTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListHookTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListHookTypesOK creates a ListHookTypesOK with default headers values
func NewListHookTypesOK() *ListHookTypesOK {
	return &ListHookTypesOK{}
}

/*ListHookTypesOK handles this case with default header values.

OK
*/
type ListHookTypesOK struct {
	Payload []*models.HookType
}

func (o *ListHookTypesOK) Error() string {
	return fmt.Sprintf("[GET /hooks/types][%d] listHookTypesOK  %+v", 200, o.Payload)
}

func (o *ListHookTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHookTypesDefault creates a ListHookTypesDefault with default headers values
func NewListHookTypesDefault(code int) *ListHookTypesDefault {
	return &ListHookTypesDefault{
		_statusCode: code,
	}
}

/*ListHookTypesDefault handles this case with default header values.

error
*/
type ListHookTypesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list hook types default response
func (o *ListHookTypesDefault) Code() int {
	return o._statusCode
}

func (o *ListHookTypesDefault) Error() string {
	return fmt.Sprintf("[GET /hooks/types][%d] listHookTypes default  %+v", o._statusCode, o.Payload)
}

func (o *ListHookTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}