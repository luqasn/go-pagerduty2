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

// ListNotificationRulesReader is a Reader for the ListNotificationRules structure.
type ListNotificationRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNotificationRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListNotificationRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListNotificationRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListNotificationRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListNotificationRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewListNotificationRulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListNotificationRulesOK creates a ListNotificationRulesOK with default headers values
func NewListNotificationRulesOK() *ListNotificationRulesOK {
	return &ListNotificationRulesOK{}
}

/*ListNotificationRulesOK handles this case with default header values.

A list of notification rules.
*/
type ListNotificationRulesOK struct {
	Payload *models.ListNotificationRulesOKBody
}

func (o *ListNotificationRulesOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}/notification_rules][%d] listNotificationRulesOK  %+v", 200, o.Payload)
}

func (o *ListNotificationRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListNotificationRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNotificationRulesBadRequest creates a ListNotificationRulesBadRequest with default headers values
func NewListNotificationRulesBadRequest() *ListNotificationRulesBadRequest {
	return &ListNotificationRulesBadRequest{}
}

/*ListNotificationRulesBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type ListNotificationRulesBadRequest struct {
	Payload *models.Error
}

func (o *ListNotificationRulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{id}/notification_rules][%d] listNotificationRulesBadRequest  %+v", 400, o.Payload)
}

func (o *ListNotificationRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNotificationRulesUnauthorized creates a ListNotificationRulesUnauthorized with default headers values
func NewListNotificationRulesUnauthorized() *ListNotificationRulesUnauthorized {
	return &ListNotificationRulesUnauthorized{}
}

/*ListNotificationRulesUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type ListNotificationRulesUnauthorized struct {
	Payload *models.Error
}

func (o *ListNotificationRulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{id}/notification_rules][%d] listNotificationRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListNotificationRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNotificationRulesForbidden creates a ListNotificationRulesForbidden with default headers values
func NewListNotificationRulesForbidden() *ListNotificationRulesForbidden {
	return &ListNotificationRulesForbidden{}
}

/*ListNotificationRulesForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type ListNotificationRulesForbidden struct {
	Payload *models.Error
}

func (o *ListNotificationRulesForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{id}/notification_rules][%d] listNotificationRulesForbidden  %+v", 403, o.Payload)
}

func (o *ListNotificationRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNotificationRulesTooManyRequests creates a ListNotificationRulesTooManyRequests with default headers values
func NewListNotificationRulesTooManyRequests() *ListNotificationRulesTooManyRequests {
	return &ListNotificationRulesTooManyRequests{}
}

/*ListNotificationRulesTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type ListNotificationRulesTooManyRequests struct {
	Payload *models.Error
}

func (o *ListNotificationRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /users/{id}/notification_rules][%d] listNotificationRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListNotificationRulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
