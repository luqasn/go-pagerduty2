// Code generated by go-swagger; DO NOT EDIT.

package add_ons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// UpdateAddonReader is a Reader for the UpdateAddon structure.
type UpdateAddonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAddonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateAddonOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateAddonBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateAddonUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateAddonForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateAddonNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewUpdateAddonTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAddonOK creates a UpdateAddonOK with default headers values
func NewUpdateAddonOK() *UpdateAddonOK {
	return &UpdateAddonOK{}
}

/*UpdateAddonOK handles this case with default header values.

The add-on that was updated.
*/
type UpdateAddonOK struct {
	Payload *models.UpdateAddonOKBody
}

func (o *UpdateAddonOK) Error() string {
	return fmt.Sprintf("[PUT /addons/{id}][%d] updateAddonOK  %+v", 200, o.Payload)
}

func (o *UpdateAddonOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateAddonOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAddonBadRequest creates a UpdateAddonBadRequest with default headers values
func NewUpdateAddonBadRequest() *UpdateAddonBadRequest {
	return &UpdateAddonBadRequest{}
}

/*UpdateAddonBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type UpdateAddonBadRequest struct {
	Payload *models.Error
}

func (o *UpdateAddonBadRequest) Error() string {
	return fmt.Sprintf("[PUT /addons/{id}][%d] updateAddonBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAddonBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAddonUnauthorized creates a UpdateAddonUnauthorized with default headers values
func NewUpdateAddonUnauthorized() *UpdateAddonUnauthorized {
	return &UpdateAddonUnauthorized{}
}

/*UpdateAddonUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type UpdateAddonUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateAddonUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /addons/{id}][%d] updateAddonUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateAddonUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAddonForbidden creates a UpdateAddonForbidden with default headers values
func NewUpdateAddonForbidden() *UpdateAddonForbidden {
	return &UpdateAddonForbidden{}
}

/*UpdateAddonForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type UpdateAddonForbidden struct {
	Payload *models.Error
}

func (o *UpdateAddonForbidden) Error() string {
	return fmt.Sprintf("[PUT /addons/{id}][%d] updateAddonForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAddonForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAddonNotFound creates a UpdateAddonNotFound with default headers values
func NewUpdateAddonNotFound() *UpdateAddonNotFound {
	return &UpdateAddonNotFound{}
}

/*UpdateAddonNotFound handles this case with default header values.

The requested resource was not found.
*/
type UpdateAddonNotFound struct {
	Payload *models.Error
}

func (o *UpdateAddonNotFound) Error() string {
	return fmt.Sprintf("[PUT /addons/{id}][%d] updateAddonNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAddonNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAddonTooManyRequests creates a UpdateAddonTooManyRequests with default headers values
func NewUpdateAddonTooManyRequests() *UpdateAddonTooManyRequests {
	return &UpdateAddonTooManyRequests{}
}

/*UpdateAddonTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type UpdateAddonTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateAddonTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /addons/{id}][%d] updateAddonTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateAddonTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}