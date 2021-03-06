// Code generated by go-swagger; DO NOT EDIT.

package schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// CreateOverrideReader is a Reader for the CreateOverride structure.
type CreateOverrideReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOverrideReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateOverrideCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateOverrideBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateOverrideUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateOverrideForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateOverrideNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewCreateOverrideTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateOverrideCreated creates a CreateOverrideCreated with default headers values
func NewCreateOverrideCreated() *CreateOverrideCreated {
	return &CreateOverrideCreated{}
}

/*CreateOverrideCreated handles this case with default header values.

The override that was created.
*/
type CreateOverrideCreated struct {
	Payload *models.CreateOverrideCreatedBody
}

func (o *CreateOverrideCreated) Error() string {
	return fmt.Sprintf("[POST /schedules/{id}/overrides][%d] createOverrideCreated  %+v", 201, o.Payload)
}

func (o *CreateOverrideCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateOverrideCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOverrideBadRequest creates a CreateOverrideBadRequest with default headers values
func NewCreateOverrideBadRequest() *CreateOverrideBadRequest {
	return &CreateOverrideBadRequest{}
}

/*CreateOverrideBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type CreateOverrideBadRequest struct {
	Payload *models.Error
}

func (o *CreateOverrideBadRequest) Error() string {
	return fmt.Sprintf("[POST /schedules/{id}/overrides][%d] createOverrideBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOverrideBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOverrideUnauthorized creates a CreateOverrideUnauthorized with default headers values
func NewCreateOverrideUnauthorized() *CreateOverrideUnauthorized {
	return &CreateOverrideUnauthorized{}
}

/*CreateOverrideUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type CreateOverrideUnauthorized struct {
	Payload *models.Error
}

func (o *CreateOverrideUnauthorized) Error() string {
	return fmt.Sprintf("[POST /schedules/{id}/overrides][%d] createOverrideUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateOverrideUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOverrideForbidden creates a CreateOverrideForbidden with default headers values
func NewCreateOverrideForbidden() *CreateOverrideForbidden {
	return &CreateOverrideForbidden{}
}

/*CreateOverrideForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type CreateOverrideForbidden struct {
	Payload *models.Error
}

func (o *CreateOverrideForbidden) Error() string {
	return fmt.Sprintf("[POST /schedules/{id}/overrides][%d] createOverrideForbidden  %+v", 403, o.Payload)
}

func (o *CreateOverrideForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOverrideNotFound creates a CreateOverrideNotFound with default headers values
func NewCreateOverrideNotFound() *CreateOverrideNotFound {
	return &CreateOverrideNotFound{}
}

/*CreateOverrideNotFound handles this case with default header values.

The requested resource was not found.
*/
type CreateOverrideNotFound struct {
	Payload *models.Error
}

func (o *CreateOverrideNotFound) Error() string {
	return fmt.Sprintf("[POST /schedules/{id}/overrides][%d] createOverrideNotFound  %+v", 404, o.Payload)
}

func (o *CreateOverrideNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOverrideTooManyRequests creates a CreateOverrideTooManyRequests with default headers values
func NewCreateOverrideTooManyRequests() *CreateOverrideTooManyRequests {
	return &CreateOverrideTooManyRequests{}
}

/*CreateOverrideTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type CreateOverrideTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateOverrideTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /schedules/{id}/overrides][%d] createOverrideTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateOverrideTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
