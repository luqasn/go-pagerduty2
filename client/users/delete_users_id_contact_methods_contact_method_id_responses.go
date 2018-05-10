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

// DeleteUsersIDContactMethodsContactMethodIDReader is a Reader for the DeleteUsersIDContactMethodsContactMethodID structure.
type DeleteUsersIDContactMethodsContactMethodIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsersIDContactMethodsContactMethodIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUsersIDContactMethodsContactMethodIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteUsersIDContactMethodsContactMethodIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteUsersIDContactMethodsContactMethodIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUsersIDContactMethodsContactMethodIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewDeleteUsersIDContactMethodsContactMethodIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUsersIDContactMethodsContactMethodIDNoContent creates a DeleteUsersIDContactMethodsContactMethodIDNoContent with default headers values
func NewDeleteUsersIDContactMethodsContactMethodIDNoContent() *DeleteUsersIDContactMethodsContactMethodIDNoContent {
	return &DeleteUsersIDContactMethodsContactMethodIDNoContent{}
}

/*DeleteUsersIDContactMethodsContactMethodIDNoContent handles this case with default header values.

The contact method was deleted successfully.
*/
type DeleteUsersIDContactMethodsContactMethodIDNoContent struct {
}

func (o *DeleteUsersIDContactMethodsContactMethodIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{id}/contact_methods/{contact_method_id}][%d] deleteUsersIdContactMethodsContactMethodIdNoContent ", 204)
}

func (o *DeleteUsersIDContactMethodsContactMethodIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsersIDContactMethodsContactMethodIDUnauthorized creates a DeleteUsersIDContactMethodsContactMethodIDUnauthorized with default headers values
func NewDeleteUsersIDContactMethodsContactMethodIDUnauthorized() *DeleteUsersIDContactMethodsContactMethodIDUnauthorized {
	return &DeleteUsersIDContactMethodsContactMethodIDUnauthorized{}
}

/*DeleteUsersIDContactMethodsContactMethodIDUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type DeleteUsersIDContactMethodsContactMethodIDUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteUsersIDContactMethodsContactMethodIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /users/{id}/contact_methods/{contact_method_id}][%d] deleteUsersIdContactMethodsContactMethodIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUsersIDContactMethodsContactMethodIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUsersIDContactMethodsContactMethodIDForbidden creates a DeleteUsersIDContactMethodsContactMethodIDForbidden with default headers values
func NewDeleteUsersIDContactMethodsContactMethodIDForbidden() *DeleteUsersIDContactMethodsContactMethodIDForbidden {
	return &DeleteUsersIDContactMethodsContactMethodIDForbidden{}
}

/*DeleteUsersIDContactMethodsContactMethodIDForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type DeleteUsersIDContactMethodsContactMethodIDForbidden struct {
	Payload *models.Error
}

func (o *DeleteUsersIDContactMethodsContactMethodIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /users/{id}/contact_methods/{contact_method_id}][%d] deleteUsersIdContactMethodsContactMethodIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUsersIDContactMethodsContactMethodIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUsersIDContactMethodsContactMethodIDNotFound creates a DeleteUsersIDContactMethodsContactMethodIDNotFound with default headers values
func NewDeleteUsersIDContactMethodsContactMethodIDNotFound() *DeleteUsersIDContactMethodsContactMethodIDNotFound {
	return &DeleteUsersIDContactMethodsContactMethodIDNotFound{}
}

/*DeleteUsersIDContactMethodsContactMethodIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteUsersIDContactMethodsContactMethodIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteUsersIDContactMethodsContactMethodIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{id}/contact_methods/{contact_method_id}][%d] deleteUsersIdContactMethodsContactMethodIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUsersIDContactMethodsContactMethodIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUsersIDContactMethodsContactMethodIDTooManyRequests creates a DeleteUsersIDContactMethodsContactMethodIDTooManyRequests with default headers values
func NewDeleteUsersIDContactMethodsContactMethodIDTooManyRequests() *DeleteUsersIDContactMethodsContactMethodIDTooManyRequests {
	return &DeleteUsersIDContactMethodsContactMethodIDTooManyRequests{}
}

/*DeleteUsersIDContactMethodsContactMethodIDTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type DeleteUsersIDContactMethodsContactMethodIDTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteUsersIDContactMethodsContactMethodIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /users/{id}/contact_methods/{contact_method_id}][%d] deleteUsersIdContactMethodsContactMethodIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteUsersIDContactMethodsContactMethodIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
