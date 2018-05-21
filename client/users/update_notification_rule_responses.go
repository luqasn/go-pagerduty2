// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// UpdateNotificationRuleReader is a Reader for the UpdateNotificationRule structure.
type UpdateNotificationRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNotificationRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateNotificationRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateNotificationRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateNotificationRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewUpdateNotificationRulePaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateNotificationRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateNotificationRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewUpdateNotificationRuleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNotificationRuleOK creates a UpdateNotificationRuleOK with default headers values
func NewUpdateNotificationRuleOK() *UpdateNotificationRuleOK {
	return &UpdateNotificationRuleOK{}
}

/*UpdateNotificationRuleOK handles this case with default header values.

The user's notification rule that was updated.
*/
type UpdateNotificationRuleOK struct {
	Payload *models.UpdateNotificationRuleOKBody
}

func (o *UpdateNotificationRuleOK) Error() string {
	return fmt.Sprintf("[PUT /users/{id}/notification_rules/{notification_rule_id}][%d] updateNotificationRuleOK  %+v", 200, o.Payload)
}

func (o *UpdateNotificationRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateNotificationRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationRuleBadRequest creates a UpdateNotificationRuleBadRequest with default headers values
func NewUpdateNotificationRuleBadRequest() *UpdateNotificationRuleBadRequest {
	return &UpdateNotificationRuleBadRequest{}
}

/*UpdateNotificationRuleBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type UpdateNotificationRuleBadRequest struct {
	Payload *models.Error
}

func (o *UpdateNotificationRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/{id}/notification_rules/{notification_rule_id}][%d] updateNotificationRuleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateNotificationRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationRuleUnauthorized creates a UpdateNotificationRuleUnauthorized with default headers values
func NewUpdateNotificationRuleUnauthorized() *UpdateNotificationRuleUnauthorized {
	return &UpdateNotificationRuleUnauthorized{}
}

/*UpdateNotificationRuleUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type UpdateNotificationRuleUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateNotificationRuleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /users/{id}/notification_rules/{notification_rule_id}][%d] updateNotificationRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateNotificationRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationRulePaymentRequired creates a UpdateNotificationRulePaymentRequired with default headers values
func NewUpdateNotificationRulePaymentRequired() *UpdateNotificationRulePaymentRequired {
	return &UpdateNotificationRulePaymentRequired{}
}

/*UpdateNotificationRulePaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type UpdateNotificationRulePaymentRequired struct {
	Payload *models.Error
}

func (o *UpdateNotificationRulePaymentRequired) Error() string {
	return fmt.Sprintf("[PUT /users/{id}/notification_rules/{notification_rule_id}][%d] updateNotificationRulePaymentRequired  %+v", 402, o.Payload)
}

func (o *UpdateNotificationRulePaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationRuleForbidden creates a UpdateNotificationRuleForbidden with default headers values
func NewUpdateNotificationRuleForbidden() *UpdateNotificationRuleForbidden {
	return &UpdateNotificationRuleForbidden{}
}

/*UpdateNotificationRuleForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type UpdateNotificationRuleForbidden struct {
	Payload *models.Error
}

func (o *UpdateNotificationRuleForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{id}/notification_rules/{notification_rule_id}][%d] updateNotificationRuleForbidden  %+v", 403, o.Payload)
}

func (o *UpdateNotificationRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationRuleNotFound creates a UpdateNotificationRuleNotFound with default headers values
func NewUpdateNotificationRuleNotFound() *UpdateNotificationRuleNotFound {
	return &UpdateNotificationRuleNotFound{}
}

/*UpdateNotificationRuleNotFound handles this case with default header values.

The requested resource was not found.
*/
type UpdateNotificationRuleNotFound struct {
	Payload *models.Error
}

func (o *UpdateNotificationRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/{id}/notification_rules/{notification_rule_id}][%d] updateNotificationRuleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateNotificationRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationRuleTooManyRequests creates a UpdateNotificationRuleTooManyRequests with default headers values
func NewUpdateNotificationRuleTooManyRequests() *UpdateNotificationRuleTooManyRequests {
	return &UpdateNotificationRuleTooManyRequests{}
}

/*UpdateNotificationRuleTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type UpdateNotificationRuleTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateNotificationRuleTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /users/{id}/notification_rules/{notification_rule_id}][%d] updateNotificationRuleTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateNotificationRuleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}