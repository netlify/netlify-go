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

// ListPaymentMethodsForUserReader is a Reader for the ListPaymentMethodsForUser structure.
type ListPaymentMethodsForUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPaymentMethodsForUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListPaymentMethodsForUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListPaymentMethodsForUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPaymentMethodsForUserOK creates a ListPaymentMethodsForUserOK with default headers values
func NewListPaymentMethodsForUserOK() *ListPaymentMethodsForUserOK {
	return &ListPaymentMethodsForUserOK{}
}

/*ListPaymentMethodsForUserOK handles this case with default header values.

OK
*/
type ListPaymentMethodsForUserOK struct {
	Payload []*models.PaymentMethod
}

func (o *ListPaymentMethodsForUserOK) Error() string {
	return fmt.Sprintf("[GET /billing/payment_methods][%d] listPaymentMethodsForUserOK  %+v", 200, o.Payload)
}

func (o *ListPaymentMethodsForUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPaymentMethodsForUserDefault creates a ListPaymentMethodsForUserDefault with default headers values
func NewListPaymentMethodsForUserDefault(code int) *ListPaymentMethodsForUserDefault {
	return &ListPaymentMethodsForUserDefault{
		_statusCode: code,
	}
}

/*ListPaymentMethodsForUserDefault handles this case with default header values.

error
*/
type ListPaymentMethodsForUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list payment methods for user default response
func (o *ListPaymentMethodsForUserDefault) Code() int {
	return o._statusCode
}

func (o *ListPaymentMethodsForUserDefault) Error() string {
	return fmt.Sprintf("[GET /billing/payment_methods][%d] listPaymentMethodsForUser default  %+v", o._statusCode, o.Payload)
}

func (o *ListPaymentMethodsForUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
