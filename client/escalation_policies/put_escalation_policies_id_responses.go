// Code generated by go-swagger; DO NOT EDIT.

package escalation_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// PutEscalationPoliciesIDReader is a Reader for the PutEscalationPoliciesID structure.
type PutEscalationPoliciesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutEscalationPoliciesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutEscalationPoliciesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPutEscalationPoliciesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewPutEscalationPoliciesIDPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutEscalationPoliciesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutEscalationPoliciesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPutEscalationPoliciesIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutEscalationPoliciesIDOK creates a PutEscalationPoliciesIDOK with default headers values
func NewPutEscalationPoliciesIDOK() *PutEscalationPoliciesIDOK {
	return &PutEscalationPoliciesIDOK{}
}

/*PutEscalationPoliciesIDOK handles this case with default header values.

The escalation policy that was updated.
*/
type PutEscalationPoliciesIDOK struct {
	Payload *models.PutEscalationPoliciesIDOKBody
}

func (o *PutEscalationPoliciesIDOK) Error() string {
	return fmt.Sprintf("[PUT /escalation_policies/{id}][%d] putEscalationPoliciesIdOK  %+v", 200, o.Payload)
}

func (o *PutEscalationPoliciesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutEscalationPoliciesIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutEscalationPoliciesIDUnauthorized creates a PutEscalationPoliciesIDUnauthorized with default headers values
func NewPutEscalationPoliciesIDUnauthorized() *PutEscalationPoliciesIDUnauthorized {
	return &PutEscalationPoliciesIDUnauthorized{}
}

/*PutEscalationPoliciesIDUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type PutEscalationPoliciesIDUnauthorized struct {
	Payload *models.Error
}

func (o *PutEscalationPoliciesIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /escalation_policies/{id}][%d] putEscalationPoliciesIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutEscalationPoliciesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutEscalationPoliciesIDPaymentRequired creates a PutEscalationPoliciesIDPaymentRequired with default headers values
func NewPutEscalationPoliciesIDPaymentRequired() *PutEscalationPoliciesIDPaymentRequired {
	return &PutEscalationPoliciesIDPaymentRequired{}
}

/*PutEscalationPoliciesIDPaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type PutEscalationPoliciesIDPaymentRequired struct {
	Payload *models.Error
}

func (o *PutEscalationPoliciesIDPaymentRequired) Error() string {
	return fmt.Sprintf("[PUT /escalation_policies/{id}][%d] putEscalationPoliciesIdPaymentRequired  %+v", 402, o.Payload)
}

func (o *PutEscalationPoliciesIDPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutEscalationPoliciesIDForbidden creates a PutEscalationPoliciesIDForbidden with default headers values
func NewPutEscalationPoliciesIDForbidden() *PutEscalationPoliciesIDForbidden {
	return &PutEscalationPoliciesIDForbidden{}
}

/*PutEscalationPoliciesIDForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type PutEscalationPoliciesIDForbidden struct {
	Payload *models.Error
}

func (o *PutEscalationPoliciesIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /escalation_policies/{id}][%d] putEscalationPoliciesIdForbidden  %+v", 403, o.Payload)
}

func (o *PutEscalationPoliciesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutEscalationPoliciesIDNotFound creates a PutEscalationPoliciesIDNotFound with default headers values
func NewPutEscalationPoliciesIDNotFound() *PutEscalationPoliciesIDNotFound {
	return &PutEscalationPoliciesIDNotFound{}
}

/*PutEscalationPoliciesIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutEscalationPoliciesIDNotFound struct {
	Payload *models.Error
}

func (o *PutEscalationPoliciesIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /escalation_policies/{id}][%d] putEscalationPoliciesIdNotFound  %+v", 404, o.Payload)
}

func (o *PutEscalationPoliciesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutEscalationPoliciesIDTooManyRequests creates a PutEscalationPoliciesIDTooManyRequests with default headers values
func NewPutEscalationPoliciesIDTooManyRequests() *PutEscalationPoliciesIDTooManyRequests {
	return &PutEscalationPoliciesIDTooManyRequests{}
}

/*PutEscalationPoliciesIDTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type PutEscalationPoliciesIDTooManyRequests struct {
	Payload *models.Error
}

func (o *PutEscalationPoliciesIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /escalation_policies/{id}][%d] putEscalationPoliciesIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutEscalationPoliciesIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
