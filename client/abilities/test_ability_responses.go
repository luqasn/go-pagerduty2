// Code generated by go-swagger; DO NOT EDIT.

package abilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// TestAbilityReader is a Reader for the TestAbility structure.
type TestAbilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestAbilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewTestAbilityNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewTestAbilityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewTestAbilityPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewTestAbilityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTestAbilityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewTestAbilityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestAbilityNoContent creates a TestAbilityNoContent with default headers values
func NewTestAbilityNoContent() *TestAbilityNoContent {
	return &TestAbilityNoContent{}
}

/*TestAbilityNoContent handles this case with default header values.

The account has the requested ability.
*/
type TestAbilityNoContent struct {
}

func (o *TestAbilityNoContent) Error() string {
	return fmt.Sprintf("[GET /abilities/{id}][%d] testAbilityNoContent ", 204)
}

func (o *TestAbilityNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTestAbilityUnauthorized creates a TestAbilityUnauthorized with default headers values
func NewTestAbilityUnauthorized() *TestAbilityUnauthorized {
	return &TestAbilityUnauthorized{}
}

/*TestAbilityUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type TestAbilityUnauthorized struct {
	Payload *models.Error
}

func (o *TestAbilityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /abilities/{id}][%d] testAbilityUnauthorized  %+v", 401, o.Payload)
}

func (o *TestAbilityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestAbilityPaymentRequired creates a TestAbilityPaymentRequired with default headers values
func NewTestAbilityPaymentRequired() *TestAbilityPaymentRequired {
	return &TestAbilityPaymentRequired{}
}

/*TestAbilityPaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type TestAbilityPaymentRequired struct {
	Payload *models.Error
}

func (o *TestAbilityPaymentRequired) Error() string {
	return fmt.Sprintf("[GET /abilities/{id}][%d] testAbilityPaymentRequired  %+v", 402, o.Payload)
}

func (o *TestAbilityPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestAbilityForbidden creates a TestAbilityForbidden with default headers values
func NewTestAbilityForbidden() *TestAbilityForbidden {
	return &TestAbilityForbidden{}
}

/*TestAbilityForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type TestAbilityForbidden struct {
	Payload *models.Error
}

func (o *TestAbilityForbidden) Error() string {
	return fmt.Sprintf("[GET /abilities/{id}][%d] testAbilityForbidden  %+v", 403, o.Payload)
}

func (o *TestAbilityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestAbilityNotFound creates a TestAbilityNotFound with default headers values
func NewTestAbilityNotFound() *TestAbilityNotFound {
	return &TestAbilityNotFound{}
}

/*TestAbilityNotFound handles this case with default header values.

The requested resource was not found.
*/
type TestAbilityNotFound struct {
	Payload *models.Error
}

func (o *TestAbilityNotFound) Error() string {
	return fmt.Sprintf("[GET /abilities/{id}][%d] testAbilityNotFound  %+v", 404, o.Payload)
}

func (o *TestAbilityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestAbilityTooManyRequests creates a TestAbilityTooManyRequests with default headers values
func NewTestAbilityTooManyRequests() *TestAbilityTooManyRequests {
	return &TestAbilityTooManyRequests{}
}

/*TestAbilityTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type TestAbilityTooManyRequests struct {
	Payload *models.Error
}

func (o *TestAbilityTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /abilities/{id}][%d] testAbilityTooManyRequests  %+v", 429, o.Payload)
}

func (o *TestAbilityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
