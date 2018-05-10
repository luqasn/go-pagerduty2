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

// DeleteTeamsIDEscalationPoliciesEscalationPolicyIDReader is a Reader for the DeleteTeamsIDEscalationPoliciesEscalationPolicyID structure.
type DeleteTeamsIDEscalationPoliciesEscalationPolicyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDNoContent creates a DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNoContent with default headers values
func NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDNoContent() *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNoContent {
	return &DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNoContent{}
}

/*DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNoContent handles this case with default header values.

The escalation policy was removed from the team.
*/
type DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNoContent struct {
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] deleteTeamsIdEscalationPoliciesEscalationPolicyIdNoContent ", 204)
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDBadRequest creates a DeleteTeamsIDEscalationPoliciesEscalationPolicyIDBadRequest with default headers values
func NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDBadRequest() *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDBadRequest {
	return &DeleteTeamsIDEscalationPoliciesEscalationPolicyIDBadRequest{}
}

/*DeleteTeamsIDEscalationPoliciesEscalationPolicyIDBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type DeleteTeamsIDEscalationPoliciesEscalationPolicyIDBadRequest struct {
	Payload *models.Error
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] deleteTeamsIdEscalationPoliciesEscalationPolicyIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDUnauthorized creates a DeleteTeamsIDEscalationPoliciesEscalationPolicyIDUnauthorized with default headers values
func NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDUnauthorized() *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDUnauthorized {
	return &DeleteTeamsIDEscalationPoliciesEscalationPolicyIDUnauthorized{}
}

/*DeleteTeamsIDEscalationPoliciesEscalationPolicyIDUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type DeleteTeamsIDEscalationPoliciesEscalationPolicyIDUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] deleteTeamsIdEscalationPoliciesEscalationPolicyIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDPaymentRequired creates a DeleteTeamsIDEscalationPoliciesEscalationPolicyIDPaymentRequired with default headers values
func NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDPaymentRequired() *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDPaymentRequired {
	return &DeleteTeamsIDEscalationPoliciesEscalationPolicyIDPaymentRequired{}
}

/*DeleteTeamsIDEscalationPoliciesEscalationPolicyIDPaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type DeleteTeamsIDEscalationPoliciesEscalationPolicyIDPaymentRequired struct {
	Payload *models.Error
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDPaymentRequired) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] deleteTeamsIdEscalationPoliciesEscalationPolicyIdPaymentRequired  %+v", 402, o.Payload)
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDForbidden creates a DeleteTeamsIDEscalationPoliciesEscalationPolicyIDForbidden with default headers values
func NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDForbidden() *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDForbidden {
	return &DeleteTeamsIDEscalationPoliciesEscalationPolicyIDForbidden{}
}

/*DeleteTeamsIDEscalationPoliciesEscalationPolicyIDForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type DeleteTeamsIDEscalationPoliciesEscalationPolicyIDForbidden struct {
	Payload *models.Error
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] deleteTeamsIdEscalationPoliciesEscalationPolicyIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDNotFound creates a DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNotFound with default headers values
func NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDNotFound() *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNotFound {
	return &DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNotFound{}
}

/*DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] deleteTeamsIdEscalationPoliciesEscalationPolicyIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDTooManyRequests creates a DeleteTeamsIDEscalationPoliciesEscalationPolicyIDTooManyRequests with default headers values
func NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDTooManyRequests() *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDTooManyRequests {
	return &DeleteTeamsIDEscalationPoliciesEscalationPolicyIDTooManyRequests{}
}

/*DeleteTeamsIDEscalationPoliciesEscalationPolicyIDTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type DeleteTeamsIDEscalationPoliciesEscalationPolicyIDTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}/escalation_policies/{escalation_policy_id}][%d] deleteTeamsIdEscalationPoliciesEscalationPolicyIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
