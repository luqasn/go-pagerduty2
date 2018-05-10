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

// GetLogEntriesIDReader is a Reader for the GetLogEntriesID structure.
type GetLogEntriesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogEntriesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLogEntriesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetLogEntriesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetLogEntriesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetLogEntriesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetLogEntriesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetLogEntriesIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLogEntriesIDOK creates a GetLogEntriesIDOK with default headers values
func NewGetLogEntriesIDOK() *GetLogEntriesIDOK {
	return &GetLogEntriesIDOK{}
}

/*GetLogEntriesIDOK handles this case with default header values.

A single log entry.
*/
type GetLogEntriesIDOK struct {
	Payload *models.GetLogEntriesIDOKBody
}

func (o *GetLogEntriesIDOK) Error() string {
	return fmt.Sprintf("[GET /log_entries/{id}][%d] getLogEntriesIdOK  %+v", 200, o.Payload)
}

func (o *GetLogEntriesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogEntriesIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntriesIDBadRequest creates a GetLogEntriesIDBadRequest with default headers values
func NewGetLogEntriesIDBadRequest() *GetLogEntriesIDBadRequest {
	return &GetLogEntriesIDBadRequest{}
}

/*GetLogEntriesIDBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type GetLogEntriesIDBadRequest struct {
	Payload *models.Error
}

func (o *GetLogEntriesIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /log_entries/{id}][%d] getLogEntriesIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLogEntriesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntriesIDUnauthorized creates a GetLogEntriesIDUnauthorized with default headers values
func NewGetLogEntriesIDUnauthorized() *GetLogEntriesIDUnauthorized {
	return &GetLogEntriesIDUnauthorized{}
}

/*GetLogEntriesIDUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type GetLogEntriesIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetLogEntriesIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /log_entries/{id}][%d] getLogEntriesIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLogEntriesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntriesIDForbidden creates a GetLogEntriesIDForbidden with default headers values
func NewGetLogEntriesIDForbidden() *GetLogEntriesIDForbidden {
	return &GetLogEntriesIDForbidden{}
}

/*GetLogEntriesIDForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type GetLogEntriesIDForbidden struct {
	Payload *models.Error
}

func (o *GetLogEntriesIDForbidden) Error() string {
	return fmt.Sprintf("[GET /log_entries/{id}][%d] getLogEntriesIdForbidden  %+v", 403, o.Payload)
}

func (o *GetLogEntriesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntriesIDNotFound creates a GetLogEntriesIDNotFound with default headers values
func NewGetLogEntriesIDNotFound() *GetLogEntriesIDNotFound {
	return &GetLogEntriesIDNotFound{}
}

/*GetLogEntriesIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLogEntriesIDNotFound struct {
	Payload *models.Error
}

func (o *GetLogEntriesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /log_entries/{id}][%d] getLogEntriesIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLogEntriesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntriesIDTooManyRequests creates a GetLogEntriesIDTooManyRequests with default headers values
func NewGetLogEntriesIDTooManyRequests() *GetLogEntriesIDTooManyRequests {
	return &GetLogEntriesIDTooManyRequests{}
}

/*GetLogEntriesIDTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type GetLogEntriesIDTooManyRequests struct {
	Payload *models.Error
}

func (o *GetLogEntriesIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /log_entries/{id}][%d] getLogEntriesIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLogEntriesIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
