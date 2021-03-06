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

// RemoveEscalationPolicyReader is a Reader for the RemoveEscalationPolicy structure.
type RemoveEscalationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveEscalationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRemoveEscalationPolicyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveEscalationPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveEscalationPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewRemoveEscalationPolicyPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveEscalationPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveEscalationPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewRemoveEscalationPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveEscalationPolicyNoContent creates a RemoveEscalationPolicyNoContent with default headers values
func NewRemoveEscalationPolicyNoContent() *RemoveEscalationPolicyNoContent {
	return &RemoveEscalationPolicyNoContent{}
}

/*RemoveEscalationPolicyNoContent handles this case with default header values.

The escalation policy was removed from the team.
*/
type RemoveEscalationPolicyNoContent struct {
}

func (o *RemoveEscalationPolicyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] removeEscalationPolicyNoContent ", 204)
}

func (o *RemoveEscalationPolicyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveEscalationPolicyBadRequest creates a RemoveEscalationPolicyBadRequest with default headers values
func NewRemoveEscalationPolicyBadRequest() *RemoveEscalationPolicyBadRequest {
	return &RemoveEscalationPolicyBadRequest{}
}

/*RemoveEscalationPolicyBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type RemoveEscalationPolicyBadRequest struct {
	Payload *models.Error
}

func (o *RemoveEscalationPolicyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] removeEscalationPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveEscalationPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveEscalationPolicyUnauthorized creates a RemoveEscalationPolicyUnauthorized with default headers values
func NewRemoveEscalationPolicyUnauthorized() *RemoveEscalationPolicyUnauthorized {
	return &RemoveEscalationPolicyUnauthorized{}
}

/*RemoveEscalationPolicyUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type RemoveEscalationPolicyUnauthorized struct {
	Payload *models.Error
}

func (o *RemoveEscalationPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] removeEscalationPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveEscalationPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveEscalationPolicyPaymentRequired creates a RemoveEscalationPolicyPaymentRequired with default headers values
func NewRemoveEscalationPolicyPaymentRequired() *RemoveEscalationPolicyPaymentRequired {
	return &RemoveEscalationPolicyPaymentRequired{}
}

/*RemoveEscalationPolicyPaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type RemoveEscalationPolicyPaymentRequired struct {
	Payload *models.Error
}

func (o *RemoveEscalationPolicyPaymentRequired) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] removeEscalationPolicyPaymentRequired  %+v", 402, o.Payload)
}

func (o *RemoveEscalationPolicyPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveEscalationPolicyForbidden creates a RemoveEscalationPolicyForbidden with default headers values
func NewRemoveEscalationPolicyForbidden() *RemoveEscalationPolicyForbidden {
	return &RemoveEscalationPolicyForbidden{}
}

/*RemoveEscalationPolicyForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type RemoveEscalationPolicyForbidden struct {
	Payload *models.Error
}

func (o *RemoveEscalationPolicyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] removeEscalationPolicyForbidden  %+v", 403, o.Payload)
}

func (o *RemoveEscalationPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveEscalationPolicyNotFound creates a RemoveEscalationPolicyNotFound with default headers values
func NewRemoveEscalationPolicyNotFound() *RemoveEscalationPolicyNotFound {
	return &RemoveEscalationPolicyNotFound{}
}

/*RemoveEscalationPolicyNotFound handles this case with default header values.

The requested resource was not found.
*/
type RemoveEscalationPolicyNotFound struct {
	Payload *models.Error
}

func (o *RemoveEscalationPolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] removeEscalationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *RemoveEscalationPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveEscalationPolicyTooManyRequests creates a RemoveEscalationPolicyTooManyRequests with default headers values
func NewRemoveEscalationPolicyTooManyRequests() *RemoveEscalationPolicyTooManyRequests {
	return &RemoveEscalationPolicyTooManyRequests{}
}

/*RemoveEscalationPolicyTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type RemoveEscalationPolicyTooManyRequests struct {
	Payload *models.Error
}

func (o *RemoveEscalationPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] removeEscalationPolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *RemoveEscalationPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
