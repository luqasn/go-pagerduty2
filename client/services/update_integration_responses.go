// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// UpdateIntegrationReader is a Reader for the UpdateIntegration structure.
type UpdateIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateIntegrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateIntegrationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateIntegrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateIntegrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewUpdateIntegrationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateIntegrationOK creates a UpdateIntegrationOK with default headers values
func NewUpdateIntegrationOK() *UpdateIntegrationOK {
	return &UpdateIntegrationOK{}
}

/*UpdateIntegrationOK handles this case with default header values.

The integration that was updated.
*/
type UpdateIntegrationOK struct {
	Payload *models.UpdateIntegrationOKBody
}

func (o *UpdateIntegrationOK) Error() string {
	return fmt.Sprintf("[PUT /services/{id}/integrations/{integration_id}][%d] updateIntegrationOK  %+v", 200, o.Payload)
}

func (o *UpdateIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateIntegrationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIntegrationBadRequest creates a UpdateIntegrationBadRequest with default headers values
func NewUpdateIntegrationBadRequest() *UpdateIntegrationBadRequest {
	return &UpdateIntegrationBadRequest{}
}

/*UpdateIntegrationBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type UpdateIntegrationBadRequest struct {
	Payload *models.Error
}

func (o *UpdateIntegrationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/{id}/integrations/{integration_id}][%d] updateIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateIntegrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIntegrationUnauthorized creates a UpdateIntegrationUnauthorized with default headers values
func NewUpdateIntegrationUnauthorized() *UpdateIntegrationUnauthorized {
	return &UpdateIntegrationUnauthorized{}
}

/*UpdateIntegrationUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type UpdateIntegrationUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateIntegrationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /services/{id}/integrations/{integration_id}][%d] updateIntegrationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateIntegrationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIntegrationForbidden creates a UpdateIntegrationForbidden with default headers values
func NewUpdateIntegrationForbidden() *UpdateIntegrationForbidden {
	return &UpdateIntegrationForbidden{}
}

/*UpdateIntegrationForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type UpdateIntegrationForbidden struct {
	Payload *models.Error
}

func (o *UpdateIntegrationForbidden) Error() string {
	return fmt.Sprintf("[PUT /services/{id}/integrations/{integration_id}][%d] updateIntegrationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateIntegrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIntegrationNotFound creates a UpdateIntegrationNotFound with default headers values
func NewUpdateIntegrationNotFound() *UpdateIntegrationNotFound {
	return &UpdateIntegrationNotFound{}
}

/*UpdateIntegrationNotFound handles this case with default header values.

The requested resource was not found.
*/
type UpdateIntegrationNotFound struct {
	Payload *models.Error
}

func (o *UpdateIntegrationNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/{id}/integrations/{integration_id}][%d] updateIntegrationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateIntegrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIntegrationTooManyRequests creates a UpdateIntegrationTooManyRequests with default headers values
func NewUpdateIntegrationTooManyRequests() *UpdateIntegrationTooManyRequests {
	return &UpdateIntegrationTooManyRequests{}
}

/*UpdateIntegrationTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type UpdateIntegrationTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateIntegrationTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /services/{id}/integrations/{integration_id}][%d] updateIntegrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateIntegrationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
