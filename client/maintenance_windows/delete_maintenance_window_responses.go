// Code generated by go-swagger; DO NOT EDIT.

package maintenance_windows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// DeleteMaintenanceWindowReader is a Reader for the DeleteMaintenanceWindow structure.
type DeleteMaintenanceWindowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMaintenanceWindowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteMaintenanceWindowNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteMaintenanceWindowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteMaintenanceWindowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteMaintenanceWindowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewDeleteMaintenanceWindowMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewDeleteMaintenanceWindowTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteMaintenanceWindowNoContent creates a DeleteMaintenanceWindowNoContent with default headers values
func NewDeleteMaintenanceWindowNoContent() *DeleteMaintenanceWindowNoContent {
	return &DeleteMaintenanceWindowNoContent{}
}

/*DeleteMaintenanceWindowNoContent handles this case with default header values.

The maintenance window was deleted successfully.
*/
type DeleteMaintenanceWindowNoContent struct {
}

func (o *DeleteMaintenanceWindowNoContent) Error() string {
	return fmt.Sprintf("[DELETE /maintenance_windows/{id}][%d] deleteMaintenanceWindowNoContent ", 204)
}

func (o *DeleteMaintenanceWindowNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMaintenanceWindowUnauthorized creates a DeleteMaintenanceWindowUnauthorized with default headers values
func NewDeleteMaintenanceWindowUnauthorized() *DeleteMaintenanceWindowUnauthorized {
	return &DeleteMaintenanceWindowUnauthorized{}
}

/*DeleteMaintenanceWindowUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type DeleteMaintenanceWindowUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteMaintenanceWindowUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /maintenance_windows/{id}][%d] deleteMaintenanceWindowUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteMaintenanceWindowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMaintenanceWindowForbidden creates a DeleteMaintenanceWindowForbidden with default headers values
func NewDeleteMaintenanceWindowForbidden() *DeleteMaintenanceWindowForbidden {
	return &DeleteMaintenanceWindowForbidden{}
}

/*DeleteMaintenanceWindowForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type DeleteMaintenanceWindowForbidden struct {
	Payload *models.Error
}

func (o *DeleteMaintenanceWindowForbidden) Error() string {
	return fmt.Sprintf("[DELETE /maintenance_windows/{id}][%d] deleteMaintenanceWindowForbidden  %+v", 403, o.Payload)
}

func (o *DeleteMaintenanceWindowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMaintenanceWindowNotFound creates a DeleteMaintenanceWindowNotFound with default headers values
func NewDeleteMaintenanceWindowNotFound() *DeleteMaintenanceWindowNotFound {
	return &DeleteMaintenanceWindowNotFound{}
}

/*DeleteMaintenanceWindowNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteMaintenanceWindowNotFound struct {
	Payload *models.Error
}

func (o *DeleteMaintenanceWindowNotFound) Error() string {
	return fmt.Sprintf("[DELETE /maintenance_windows/{id}][%d] deleteMaintenanceWindowNotFound  %+v", 404, o.Payload)
}

func (o *DeleteMaintenanceWindowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMaintenanceWindowMethodNotAllowed creates a DeleteMaintenanceWindowMethodNotAllowed with default headers values
func NewDeleteMaintenanceWindowMethodNotAllowed() *DeleteMaintenanceWindowMethodNotAllowed {
	return &DeleteMaintenanceWindowMethodNotAllowed{}
}

/*DeleteMaintenanceWindowMethodNotAllowed handles this case with default header values.

The maintenance window can't be deleted because it has already ended.
*/
type DeleteMaintenanceWindowMethodNotAllowed struct {
}

func (o *DeleteMaintenanceWindowMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /maintenance_windows/{id}][%d] deleteMaintenanceWindowMethodNotAllowed ", 405)
}

func (o *DeleteMaintenanceWindowMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMaintenanceWindowTooManyRequests creates a DeleteMaintenanceWindowTooManyRequests with default headers values
func NewDeleteMaintenanceWindowTooManyRequests() *DeleteMaintenanceWindowTooManyRequests {
	return &DeleteMaintenanceWindowTooManyRequests{}
}

/*DeleteMaintenanceWindowTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type DeleteMaintenanceWindowTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteMaintenanceWindowTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /maintenance_windows/{id}][%d] deleteMaintenanceWindowTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteMaintenanceWindowTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
