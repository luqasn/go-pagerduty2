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

// PostServicesIDIntegrationsReader is a Reader for the PostServicesIDIntegrations structure.
type PostServicesIDIntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostServicesIDIntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostServicesIDIntegrationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostServicesIDIntegrationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostServicesIDIntegrationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostServicesIDIntegrationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostServicesIDIntegrationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostServicesIDIntegrationsCreated creates a PostServicesIDIntegrationsCreated with default headers values
func NewPostServicesIDIntegrationsCreated() *PostServicesIDIntegrationsCreated {
	return &PostServicesIDIntegrationsCreated{}
}

/*PostServicesIDIntegrationsCreated handles this case with default header values.

The integration that was created.
*/
type PostServicesIDIntegrationsCreated struct {
	Payload *models.PostServicesIDIntegrationsCreatedBody
}

func (o *PostServicesIDIntegrationsCreated) Error() string {
	return fmt.Sprintf("[POST /services/{id}/integrations][%d] postServicesIdIntegrationsCreated  %+v", 201, o.Payload)
}

func (o *PostServicesIDIntegrationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostServicesIDIntegrationsCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServicesIDIntegrationsBadRequest creates a PostServicesIDIntegrationsBadRequest with default headers values
func NewPostServicesIDIntegrationsBadRequest() *PostServicesIDIntegrationsBadRequest {
	return &PostServicesIDIntegrationsBadRequest{}
}

/*PostServicesIDIntegrationsBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type PostServicesIDIntegrationsBadRequest struct {
	Payload *models.Error
}

func (o *PostServicesIDIntegrationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/{id}/integrations][%d] postServicesIdIntegrationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostServicesIDIntegrationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServicesIDIntegrationsUnauthorized creates a PostServicesIDIntegrationsUnauthorized with default headers values
func NewPostServicesIDIntegrationsUnauthorized() *PostServicesIDIntegrationsUnauthorized {
	return &PostServicesIDIntegrationsUnauthorized{}
}

/*PostServicesIDIntegrationsUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type PostServicesIDIntegrationsUnauthorized struct {
	Payload *models.Error
}

func (o *PostServicesIDIntegrationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /services/{id}/integrations][%d] postServicesIdIntegrationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostServicesIDIntegrationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServicesIDIntegrationsForbidden creates a PostServicesIDIntegrationsForbidden with default headers values
func NewPostServicesIDIntegrationsForbidden() *PostServicesIDIntegrationsForbidden {
	return &PostServicesIDIntegrationsForbidden{}
}

/*PostServicesIDIntegrationsForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type PostServicesIDIntegrationsForbidden struct {
	Payload *models.Error
}

func (o *PostServicesIDIntegrationsForbidden) Error() string {
	return fmt.Sprintf("[POST /services/{id}/integrations][%d] postServicesIdIntegrationsForbidden  %+v", 403, o.Payload)
}

func (o *PostServicesIDIntegrationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServicesIDIntegrationsTooManyRequests creates a PostServicesIDIntegrationsTooManyRequests with default headers values
func NewPostServicesIDIntegrationsTooManyRequests() *PostServicesIDIntegrationsTooManyRequests {
	return &PostServicesIDIntegrationsTooManyRequests{}
}

/*PostServicesIDIntegrationsTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type PostServicesIDIntegrationsTooManyRequests struct {
	Payload *models.Error
}

func (o *PostServicesIDIntegrationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /services/{id}/integrations][%d] postServicesIdIntegrationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostServicesIDIntegrationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}