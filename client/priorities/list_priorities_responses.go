// Code generated by go-swagger; DO NOT EDIT.

package priorities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// ListPrioritiesReader is a Reader for the ListPriorities structure.
type ListPrioritiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPrioritiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListPrioritiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListPrioritiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListPrioritiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewListPrioritiesPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListPrioritiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewListPrioritiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListPrioritiesOK creates a ListPrioritiesOK with default headers values
func NewListPrioritiesOK() *ListPrioritiesOK {
	return &ListPrioritiesOK{}
}

/*ListPrioritiesOK handles this case with default header values.

A paginated array of priorities.
*/
type ListPrioritiesOK struct {
	Payload *models.ListPrioritiesOKBody
}

func (o *ListPrioritiesOK) Error() string {
	return fmt.Sprintf("[GET /priorities][%d] listPrioritiesOK  %+v", 200, o.Payload)
}

func (o *ListPrioritiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListPrioritiesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPrioritiesBadRequest creates a ListPrioritiesBadRequest with default headers values
func NewListPrioritiesBadRequest() *ListPrioritiesBadRequest {
	return &ListPrioritiesBadRequest{}
}

/*ListPrioritiesBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type ListPrioritiesBadRequest struct {
	Payload *models.Error
}

func (o *ListPrioritiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /priorities][%d] listPrioritiesBadRequest  %+v", 400, o.Payload)
}

func (o *ListPrioritiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPrioritiesUnauthorized creates a ListPrioritiesUnauthorized with default headers values
func NewListPrioritiesUnauthorized() *ListPrioritiesUnauthorized {
	return &ListPrioritiesUnauthorized{}
}

/*ListPrioritiesUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type ListPrioritiesUnauthorized struct {
	Payload *models.Error
}

func (o *ListPrioritiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /priorities][%d] listPrioritiesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListPrioritiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPrioritiesPaymentRequired creates a ListPrioritiesPaymentRequired with default headers values
func NewListPrioritiesPaymentRequired() *ListPrioritiesPaymentRequired {
	return &ListPrioritiesPaymentRequired{}
}

/*ListPrioritiesPaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type ListPrioritiesPaymentRequired struct {
	Payload *models.Error
}

func (o *ListPrioritiesPaymentRequired) Error() string {
	return fmt.Sprintf("[GET /priorities][%d] listPrioritiesPaymentRequired  %+v", 402, o.Payload)
}

func (o *ListPrioritiesPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPrioritiesForbidden creates a ListPrioritiesForbidden with default headers values
func NewListPrioritiesForbidden() *ListPrioritiesForbidden {
	return &ListPrioritiesForbidden{}
}

/*ListPrioritiesForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type ListPrioritiesForbidden struct {
	Payload *models.Error
}

func (o *ListPrioritiesForbidden) Error() string {
	return fmt.Sprintf("[GET /priorities][%d] listPrioritiesForbidden  %+v", 403, o.Payload)
}

func (o *ListPrioritiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPrioritiesTooManyRequests creates a ListPrioritiesTooManyRequests with default headers values
func NewListPrioritiesTooManyRequests() *ListPrioritiesTooManyRequests {
	return &ListPrioritiesTooManyRequests{}
}

/*ListPrioritiesTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type ListPrioritiesTooManyRequests struct {
	Payload *models.Error
}

func (o *ListPrioritiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /priorities][%d] listPrioritiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListPrioritiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
