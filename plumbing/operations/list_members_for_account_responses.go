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

// ListMembersForAccountReader is a Reader for the ListMembersForAccount structure.
type ListMembersForAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMembersForAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListMembersForAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListMembersForAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListMembersForAccountOK creates a ListMembersForAccountOK with default headers values
func NewListMembersForAccountOK() *ListMembersForAccountOK {
	return &ListMembersForAccountOK{}
}

/*ListMembersForAccountOK handles this case with default header values.

OK
*/
type ListMembersForAccountOK struct {
	Payload []*models.Member
}

func (o *ListMembersForAccountOK) Error() string {
	return fmt.Sprintf("[GET /{account_slug}/members][%d] listMembersForAccountOK  %+v", 200, o.Payload)
}

func (o *ListMembersForAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMembersForAccountDefault creates a ListMembersForAccountDefault with default headers values
func NewListMembersForAccountDefault(code int) *ListMembersForAccountDefault {
	return &ListMembersForAccountDefault{
		_statusCode: code,
	}
}

/*ListMembersForAccountDefault handles this case with default header values.

error
*/
type ListMembersForAccountDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list members for account default response
func (o *ListMembersForAccountDefault) Code() int {
	return o._statusCode
}

func (o *ListMembersForAccountDefault) Error() string {
	return fmt.Sprintf("[GET /{account_slug}/members][%d] listMembersForAccount default  %+v", o._statusCode, o.Payload)
}

func (o *ListMembersForAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
