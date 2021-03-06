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

// GetUserReader is a Reader for the GetUser structure.
type GetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserOK creates a GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {
	return &GetUserOK{}
}

/*GetUserOK handles this case with default header values.

The user requested.
*/
type GetUserOK struct {
	Payload *models.GetUserOKBody
}

func (o *GetUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserOK  %+v", 200, o.Payload)
}

func (o *GetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUserOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserBadRequest creates a GetUserBadRequest with default headers values
func NewGetUserBadRequest() *GetUserBadRequest {
	return &GetUserBadRequest{}
}

/*GetUserBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type GetUserBadRequest struct {
	Payload *models.Error
}

func (o *GetUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserUnauthorized creates a GetUserUnauthorized with default headers values
func NewGetUserUnauthorized() *GetUserUnauthorized {
	return &GetUserUnauthorized{}
}

/*GetUserUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type GetUserUnauthorized struct {
	Payload *models.Error
}

func (o *GetUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserForbidden creates a GetUserForbidden with default headers values
func NewGetUserForbidden() *GetUserForbidden {
	return &GetUserForbidden{}
}

/*GetUserForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type GetUserForbidden struct {
	Payload *models.Error
}

func (o *GetUserForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserForbidden  %+v", 403, o.Payload)
}

func (o *GetUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserNotFound creates a GetUserNotFound with default headers values
func NewGetUserNotFound() *GetUserNotFound {
	return &GetUserNotFound{}
}

/*GetUserNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetUserNotFound struct {
	Payload *models.Error
}

func (o *GetUserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserNotFound  %+v", 404, o.Payload)
}

func (o *GetUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTooManyRequests creates a GetUserTooManyRequests with default headers values
func NewGetUserTooManyRequests() *GetUserTooManyRequests {
	return &GetUserTooManyRequests{}
}

/*GetUserTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type GetUserTooManyRequests struct {
	Payload *models.Error
}

func (o *GetUserTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
