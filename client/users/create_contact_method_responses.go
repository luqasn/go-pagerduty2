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

// CreateContactMethodReader is a Reader for the CreateContactMethod structure.
type CreateContactMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateContactMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateContactMethodCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateContactMethodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateContactMethodUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateContactMethodForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewCreateContactMethodTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateContactMethodCreated creates a CreateContactMethodCreated with default headers values
func NewCreateContactMethodCreated() *CreateContactMethodCreated {
	return &CreateContactMethodCreated{}
}

/*CreateContactMethodCreated handles this case with default header values.

The contact method that was created.
*/
type CreateContactMethodCreated struct {
	Payload *models.CreateContactMethodCreatedBody
}

func (o *CreateContactMethodCreated) Error() string {
	return fmt.Sprintf("[POST /users/{id}/contact_methods][%d] createContactMethodCreated  %+v", 201, o.Payload)
}

func (o *CreateContactMethodCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateContactMethodCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateContactMethodBadRequest creates a CreateContactMethodBadRequest with default headers values
func NewCreateContactMethodBadRequest() *CreateContactMethodBadRequest {
	return &CreateContactMethodBadRequest{}
}

/*CreateContactMethodBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type CreateContactMethodBadRequest struct {
	Payload *models.Error
}

func (o *CreateContactMethodBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{id}/contact_methods][%d] createContactMethodBadRequest  %+v", 400, o.Payload)
}

func (o *CreateContactMethodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateContactMethodUnauthorized creates a CreateContactMethodUnauthorized with default headers values
func NewCreateContactMethodUnauthorized() *CreateContactMethodUnauthorized {
	return &CreateContactMethodUnauthorized{}
}

/*CreateContactMethodUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type CreateContactMethodUnauthorized struct {
	Payload *models.Error
}

func (o *CreateContactMethodUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users/{id}/contact_methods][%d] createContactMethodUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateContactMethodUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateContactMethodForbidden creates a CreateContactMethodForbidden with default headers values
func NewCreateContactMethodForbidden() *CreateContactMethodForbidden {
	return &CreateContactMethodForbidden{}
}

/*CreateContactMethodForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type CreateContactMethodForbidden struct {
	Payload *models.Error
}

func (o *CreateContactMethodForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{id}/contact_methods][%d] createContactMethodForbidden  %+v", 403, o.Payload)
}

func (o *CreateContactMethodForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateContactMethodTooManyRequests creates a CreateContactMethodTooManyRequests with default headers values
func NewCreateContactMethodTooManyRequests() *CreateContactMethodTooManyRequests {
	return &CreateContactMethodTooManyRequests{}
}

/*CreateContactMethodTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type CreateContactMethodTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateContactMethodTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /users/{id}/contact_methods][%d] createContactMethodTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateContactMethodTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
