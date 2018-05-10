// Code generated by go-swagger; DO NOT EDIT.

package extensions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// DeleteExtensionsIDReader is a Reader for the DeleteExtensionsID structure.
type DeleteExtensionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExtensionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteExtensionsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteExtensionsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteExtensionsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteExtensionsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteExtensionsIDNoContent creates a DeleteExtensionsIDNoContent with default headers values
func NewDeleteExtensionsIDNoContent() *DeleteExtensionsIDNoContent {
	return &DeleteExtensionsIDNoContent{}
}

/*DeleteExtensionsIDNoContent handles this case with default header values.

The extension was deleted successfully.
*/
type DeleteExtensionsIDNoContent struct {
}

func (o *DeleteExtensionsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extensions/{id}][%d] deleteExtensionsIdNoContent ", 204)
}

func (o *DeleteExtensionsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExtensionsIDUnauthorized creates a DeleteExtensionsIDUnauthorized with default headers values
func NewDeleteExtensionsIDUnauthorized() *DeleteExtensionsIDUnauthorized {
	return &DeleteExtensionsIDUnauthorized{}
}

/*DeleteExtensionsIDUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type DeleteExtensionsIDUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteExtensionsIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /extensions/{id}][%d] deleteExtensionsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExtensionsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExtensionsIDForbidden creates a DeleteExtensionsIDForbidden with default headers values
func NewDeleteExtensionsIDForbidden() *DeleteExtensionsIDForbidden {
	return &DeleteExtensionsIDForbidden{}
}

/*DeleteExtensionsIDForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type DeleteExtensionsIDForbidden struct {
	Payload *models.Error
}

func (o *DeleteExtensionsIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /extensions/{id}][%d] deleteExtensionsIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExtensionsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExtensionsIDNotFound creates a DeleteExtensionsIDNotFound with default headers values
func NewDeleteExtensionsIDNotFound() *DeleteExtensionsIDNotFound {
	return &DeleteExtensionsIDNotFound{}
}

/*DeleteExtensionsIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteExtensionsIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteExtensionsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /extensions/{id}][%d] deleteExtensionsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExtensionsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}