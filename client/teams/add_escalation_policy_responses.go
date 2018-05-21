// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// AddEscalationPolicyReader is a Reader for the AddEscalationPolicy structure.
type AddEscalationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddEscalationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewAddEscalationPolicyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddEscalationPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddEscalationPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewAddEscalationPolicyPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddEscalationPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddEscalationPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewAddEscalationPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddEscalationPolicyNoContent creates a AddEscalationPolicyNoContent with default headers values
func NewAddEscalationPolicyNoContent() *AddEscalationPolicyNoContent {
	return &AddEscalationPolicyNoContent{}
}

/*AddEscalationPolicyNoContent handles this case with default header values.

The escalation policy was added to the team.
*/
type AddEscalationPolicyNoContent struct {
}

func (o *AddEscalationPolicyNoContent) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}/escalation_policies/{escalation_policy_id}][%d] addEscalationPolicyNoContent ", 204)
}

func (o *AddEscalationPolicyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddEscalationPolicyBadRequest creates a AddEscalationPolicyBadRequest with default headers values
func NewAddEscalationPolicyBadRequest() *AddEscalationPolicyBadRequest {
	return &AddEscalationPolicyBadRequest{}
}

/*AddEscalationPolicyBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type AddEscalationPolicyBadRequest struct {
	Payload *models.Error
}

func (o *AddEscalationPolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}/escalation_policies/{escalation_policy_id}][%d] addEscalationPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *AddEscalationPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEscalationPolicyUnauthorized creates a AddEscalationPolicyUnauthorized with default headers values
func NewAddEscalationPolicyUnauthorized() *AddEscalationPolicyUnauthorized {
	return &AddEscalationPolicyUnauthorized{}
}

/*AddEscalationPolicyUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type AddEscalationPolicyUnauthorized struct {
	Payload *models.Error
}

func (o *AddEscalationPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}/escalation_policies/{escalation_policy_id}][%d] addEscalationPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *AddEscalationPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEscalationPolicyPaymentRequired creates a AddEscalationPolicyPaymentRequired with default headers values
func NewAddEscalationPolicyPaymentRequired() *AddEscalationPolicyPaymentRequired {
	return &AddEscalationPolicyPaymentRequired{}
}

/*AddEscalationPolicyPaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type AddEscalationPolicyPaymentRequired struct {
	Payload *models.Error
}

func (o *AddEscalationPolicyPaymentRequired) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}/escalation_policies/{escalation_policy_id}][%d] addEscalationPolicyPaymentRequired  %+v", 402, o.Payload)
}

func (o *AddEscalationPolicyPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEscalationPolicyForbidden creates a AddEscalationPolicyForbidden with default headers values
func NewAddEscalationPolicyForbidden() *AddEscalationPolicyForbidden {
	return &AddEscalationPolicyForbidden{}
}

/*AddEscalationPolicyForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type AddEscalationPolicyForbidden struct {
	Payload *models.Error
}

func (o *AddEscalationPolicyForbidden) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}/escalation_policies/{escalation_policy_id}][%d] addEscalationPolicyForbidden  %+v", 403, o.Payload)
}

func (o *AddEscalationPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEscalationPolicyNotFound creates a AddEscalationPolicyNotFound with default headers values
func NewAddEscalationPolicyNotFound() *AddEscalationPolicyNotFound {
	return &AddEscalationPolicyNotFound{}
}

/*AddEscalationPolicyNotFound handles this case with default header values.

The requested resource was not found.
*/
type AddEscalationPolicyNotFound struct {
	Payload *models.Error
}

func (o *AddEscalationPolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}/escalation_policies/{escalation_policy_id}][%d] addEscalationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *AddEscalationPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEscalationPolicyTooManyRequests creates a AddEscalationPolicyTooManyRequests with default headers values
func NewAddEscalationPolicyTooManyRequests() *AddEscalationPolicyTooManyRequests {
	return &AddEscalationPolicyTooManyRequests{}
}

/*AddEscalationPolicyTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type AddEscalationPolicyTooManyRequests struct {
	Payload *models.Error
}

func (o *AddEscalationPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}/escalation_policies/{escalation_policy_id}][%d] addEscalationPolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddEscalationPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
