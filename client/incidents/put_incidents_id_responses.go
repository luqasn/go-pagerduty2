// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// PutIncidentsIDReader is a Reader for the PutIncidentsID structure.
type PutIncidentsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIncidentsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutIncidentsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPutIncidentsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutIncidentsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutIncidentsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPutIncidentsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutIncidentsIDOK creates a PutIncidentsIDOK with default headers values
func NewPutIncidentsIDOK() *PutIncidentsIDOK {
	return &PutIncidentsIDOK{}
}

/*PutIncidentsIDOK handles this case with default header values.

The incident was updated.
*/
type PutIncidentsIDOK struct {
	Payload *models.PutIncidentsIDOKBody
}

func (o *PutIncidentsIDOK) Error() string {
	return fmt.Sprintf("[PUT /incidents/{id}][%d] putIncidentsIdOK  %+v", 200, o.Payload)
}

func (o *PutIncidentsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutIncidentsIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIncidentsIDUnauthorized creates a PutIncidentsIDUnauthorized with default headers values
func NewPutIncidentsIDUnauthorized() *PutIncidentsIDUnauthorized {
	return &PutIncidentsIDUnauthorized{}
}

/*PutIncidentsIDUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type PutIncidentsIDUnauthorized struct {
	Payload *models.Error
}

func (o *PutIncidentsIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /incidents/{id}][%d] putIncidentsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIncidentsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIncidentsIDForbidden creates a PutIncidentsIDForbidden with default headers values
func NewPutIncidentsIDForbidden() *PutIncidentsIDForbidden {
	return &PutIncidentsIDForbidden{}
}

/*PutIncidentsIDForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type PutIncidentsIDForbidden struct {
	Payload *models.Error
}

func (o *PutIncidentsIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /incidents/{id}][%d] putIncidentsIdForbidden  %+v", 403, o.Payload)
}

func (o *PutIncidentsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIncidentsIDNotFound creates a PutIncidentsIDNotFound with default headers values
func NewPutIncidentsIDNotFound() *PutIncidentsIDNotFound {
	return &PutIncidentsIDNotFound{}
}

/*PutIncidentsIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutIncidentsIDNotFound struct {
	Payload *models.Error
}

func (o *PutIncidentsIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /incidents/{id}][%d] putIncidentsIdNotFound  %+v", 404, o.Payload)
}

func (o *PutIncidentsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIncidentsIDTooManyRequests creates a PutIncidentsIDTooManyRequests with default headers values
func NewPutIncidentsIDTooManyRequests() *PutIncidentsIDTooManyRequests {
	return &PutIncidentsIDTooManyRequests{}
}

/*PutIncidentsIDTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type PutIncidentsIDTooManyRequests struct {
	Payload *models.Error
}

func (o *PutIncidentsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /incidents/{id}][%d] putIncidentsIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIncidentsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
