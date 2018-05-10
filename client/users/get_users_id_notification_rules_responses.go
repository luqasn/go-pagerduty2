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

// GetUsersIDNotificationRulesReader is a Reader for the GetUsersIDNotificationRules structure.
type GetUsersIDNotificationRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersIDNotificationRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersIDNotificationRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetUsersIDNotificationRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetUsersIDNotificationRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUsersIDNotificationRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetUsersIDNotificationRulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersIDNotificationRulesOK creates a GetUsersIDNotificationRulesOK with default headers values
func NewGetUsersIDNotificationRulesOK() *GetUsersIDNotificationRulesOK {
	return &GetUsersIDNotificationRulesOK{}
}

/*GetUsersIDNotificationRulesOK handles this case with default header values.

A list of notification rules.
*/
type GetUsersIDNotificationRulesOK struct {
	Payload *models.GetUsersIDNotificationRulesOKBody
}

func (o *GetUsersIDNotificationRulesOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}/notification_rules][%d] getUsersIdNotificationRulesOK  %+v", 200, o.Payload)
}

func (o *GetUsersIDNotificationRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUsersIDNotificationRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersIDNotificationRulesBadRequest creates a GetUsersIDNotificationRulesBadRequest with default headers values
func NewGetUsersIDNotificationRulesBadRequest() *GetUsersIDNotificationRulesBadRequest {
	return &GetUsersIDNotificationRulesBadRequest{}
}

/*GetUsersIDNotificationRulesBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type GetUsersIDNotificationRulesBadRequest struct {
	Payload *models.Error
}

func (o *GetUsersIDNotificationRulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{id}/notification_rules][%d] getUsersIdNotificationRulesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsersIDNotificationRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersIDNotificationRulesUnauthorized creates a GetUsersIDNotificationRulesUnauthorized with default headers values
func NewGetUsersIDNotificationRulesUnauthorized() *GetUsersIDNotificationRulesUnauthorized {
	return &GetUsersIDNotificationRulesUnauthorized{}
}

/*GetUsersIDNotificationRulesUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type GetUsersIDNotificationRulesUnauthorized struct {
	Payload *models.Error
}

func (o *GetUsersIDNotificationRulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{id}/notification_rules][%d] getUsersIdNotificationRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUsersIDNotificationRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersIDNotificationRulesForbidden creates a GetUsersIDNotificationRulesForbidden with default headers values
func NewGetUsersIDNotificationRulesForbidden() *GetUsersIDNotificationRulesForbidden {
	return &GetUsersIDNotificationRulesForbidden{}
}

/*GetUsersIDNotificationRulesForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type GetUsersIDNotificationRulesForbidden struct {
	Payload *models.Error
}

func (o *GetUsersIDNotificationRulesForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{id}/notification_rules][%d] getUsersIdNotificationRulesForbidden  %+v", 403, o.Payload)
}

func (o *GetUsersIDNotificationRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersIDNotificationRulesTooManyRequests creates a GetUsersIDNotificationRulesTooManyRequests with default headers values
func NewGetUsersIDNotificationRulesTooManyRequests() *GetUsersIDNotificationRulesTooManyRequests {
	return &GetUsersIDNotificationRulesTooManyRequests{}
}

/*GetUsersIDNotificationRulesTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type GetUsersIDNotificationRulesTooManyRequests struct {
	Payload *models.Error
}

func (o *GetUsersIDNotificationRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /users/{id}/notification_rules][%d] getUsersIdNotificationRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUsersIDNotificationRulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}