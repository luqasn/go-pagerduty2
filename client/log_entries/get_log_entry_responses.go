// Code generated by go-swagger; DO NOT EDIT.

package log_entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// GetLogEntryReader is a Reader for the GetLogEntry structure.
type GetLogEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLogEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetLogEntryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetLogEntryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetLogEntryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetLogEntryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetLogEntryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLogEntryOK creates a GetLogEntryOK with default headers values
func NewGetLogEntryOK() *GetLogEntryOK {
	return &GetLogEntryOK{}
}

/*GetLogEntryOK handles this case with default header values.

A single log entry.
*/
type GetLogEntryOK struct {
	Payload *models.GetLogEntryOKBody
}

func (o *GetLogEntryOK) Error() string {
	return fmt.Sprintf("[GET /log_entries/{id}][%d] getLogEntryOK  %+v", 200, o.Payload)
}

func (o *GetLogEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogEntryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntryBadRequest creates a GetLogEntryBadRequest with default headers values
func NewGetLogEntryBadRequest() *GetLogEntryBadRequest {
	return &GetLogEntryBadRequest{}
}

/*GetLogEntryBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type GetLogEntryBadRequest struct {
	Payload *models.Error
}

func (o *GetLogEntryBadRequest) Error() string {
	return fmt.Sprintf("[GET /log_entries/{id}][%d] getLogEntryBadRequest  %+v", 400, o.Payload)
}

func (o *GetLogEntryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntryUnauthorized creates a GetLogEntryUnauthorized with default headers values
func NewGetLogEntryUnauthorized() *GetLogEntryUnauthorized {
	return &GetLogEntryUnauthorized{}
}

/*GetLogEntryUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type GetLogEntryUnauthorized struct {
	Payload *models.Error
}

func (o *GetLogEntryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /log_entries/{id}][%d] getLogEntryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLogEntryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntryForbidden creates a GetLogEntryForbidden with default headers values
func NewGetLogEntryForbidden() *GetLogEntryForbidden {
	return &GetLogEntryForbidden{}
}

/*GetLogEntryForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type GetLogEntryForbidden struct {
	Payload *models.Error
}

func (o *GetLogEntryForbidden) Error() string {
	return fmt.Sprintf("[GET /log_entries/{id}][%d] getLogEntryForbidden  %+v", 403, o.Payload)
}

func (o *GetLogEntryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntryNotFound creates a GetLogEntryNotFound with default headers values
func NewGetLogEntryNotFound() *GetLogEntryNotFound {
	return &GetLogEntryNotFound{}
}

/*GetLogEntryNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLogEntryNotFound struct {
	Payload *models.Error
}

func (o *GetLogEntryNotFound) Error() string {
	return fmt.Sprintf("[GET /log_entries/{id}][%d] getLogEntryNotFound  %+v", 404, o.Payload)
}

func (o *GetLogEntryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntryTooManyRequests creates a GetLogEntryTooManyRequests with default headers values
func NewGetLogEntryTooManyRequests() *GetLogEntryTooManyRequests {
	return &GetLogEntryTooManyRequests{}
}

/*GetLogEntryTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type GetLogEntryTooManyRequests struct {
	Payload *models.Error
}

func (o *GetLogEntryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /log_entries/{id}][%d] getLogEntryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLogEntryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
