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

// PostUsersIDNotificationRulesReader is a Reader for the PostUsersIDNotificationRules structure.
type PostUsersIDNotificationRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersIDNotificationRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostUsersIDNotificationRulesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostUsersIDNotificationRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostUsersIDNotificationRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewPostUsersIDNotificationRulesPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostUsersIDNotificationRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostUsersIDNotificationRulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUsersIDNotificationRulesCreated creates a PostUsersIDNotificationRulesCreated with default headers values
func NewPostUsersIDNotificationRulesCreated() *PostUsersIDNotificationRulesCreated {
	return &PostUsersIDNotificationRulesCreated{}
}

/*PostUsersIDNotificationRulesCreated handles this case with default header values.

The notification rule that was created.
*/
type PostUsersIDNotificationRulesCreated struct {
	Payload *models.PostUsersIDNotificationRulesCreatedBody
}

func (o *PostUsersIDNotificationRulesCreated) Error() string {
	return fmt.Sprintf("[POST /users/{id}/notification_rules][%d] postUsersIdNotificationRulesCreated  %+v", 201, o.Payload)
}

func (o *PostUsersIDNotificationRulesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostUsersIDNotificationRulesCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersIDNotificationRulesBadRequest creates a PostUsersIDNotificationRulesBadRequest with default headers values
func NewPostUsersIDNotificationRulesBadRequest() *PostUsersIDNotificationRulesBadRequest {
	return &PostUsersIDNotificationRulesBadRequest{}
}

/*PostUsersIDNotificationRulesBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type PostUsersIDNotificationRulesBadRequest struct {
	Payload *models.Error
}

func (o *PostUsersIDNotificationRulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{id}/notification_rules][%d] postUsersIdNotificationRulesBadRequest  %+v", 400, o.Payload)
}

func (o *PostUsersIDNotificationRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersIDNotificationRulesUnauthorized creates a PostUsersIDNotificationRulesUnauthorized with default headers values
func NewPostUsersIDNotificationRulesUnauthorized() *PostUsersIDNotificationRulesUnauthorized {
	return &PostUsersIDNotificationRulesUnauthorized{}
}

/*PostUsersIDNotificationRulesUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type PostUsersIDNotificationRulesUnauthorized struct {
	Payload *models.Error
}

func (o *PostUsersIDNotificationRulesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users/{id}/notification_rules][%d] postUsersIdNotificationRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUsersIDNotificationRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersIDNotificationRulesPaymentRequired creates a PostUsersIDNotificationRulesPaymentRequired with default headers values
func NewPostUsersIDNotificationRulesPaymentRequired() *PostUsersIDNotificationRulesPaymentRequired {
	return &PostUsersIDNotificationRulesPaymentRequired{}
}

/*PostUsersIDNotificationRulesPaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type PostUsersIDNotificationRulesPaymentRequired struct {
	Payload *models.Error
}

func (o *PostUsersIDNotificationRulesPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /users/{id}/notification_rules][%d] postUsersIdNotificationRulesPaymentRequired  %+v", 402, o.Payload)
}

func (o *PostUsersIDNotificationRulesPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersIDNotificationRulesForbidden creates a PostUsersIDNotificationRulesForbidden with default headers values
func NewPostUsersIDNotificationRulesForbidden() *PostUsersIDNotificationRulesForbidden {
	return &PostUsersIDNotificationRulesForbidden{}
}

/*PostUsersIDNotificationRulesForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type PostUsersIDNotificationRulesForbidden struct {
	Payload *models.Error
}

func (o *PostUsersIDNotificationRulesForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{id}/notification_rules][%d] postUsersIdNotificationRulesForbidden  %+v", 403, o.Payload)
}

func (o *PostUsersIDNotificationRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersIDNotificationRulesTooManyRequests creates a PostUsersIDNotificationRulesTooManyRequests with default headers values
func NewPostUsersIDNotificationRulesTooManyRequests() *PostUsersIDNotificationRulesTooManyRequests {
	return &PostUsersIDNotificationRulesTooManyRequests{}
}

/*PostUsersIDNotificationRulesTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type PostUsersIDNotificationRulesTooManyRequests struct {
	Payload *models.Error
}

func (o *PostUsersIDNotificationRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /users/{id}/notification_rules][%d] postUsersIdNotificationRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUsersIDNotificationRulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}