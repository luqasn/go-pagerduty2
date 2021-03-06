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

// GetExtensionReader is a Reader for the GetExtension structure.
type GetExtensionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExtensionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetExtensionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetExtensionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetExtensionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetExtensionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetExtensionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetExtensionOK creates a GetExtensionOK with default headers values
func NewGetExtensionOK() *GetExtensionOK {
	return &GetExtensionOK{}
}

/*GetExtensionOK handles this case with default header values.

The extension that was requested.
*/
type GetExtensionOK struct {
	Payload *models.GetExtensionOKBody
}

func (o *GetExtensionOK) Error() string {
	return fmt.Sprintf("[GET /extensions/{id}][%d] getExtensionOK  %+v", 200, o.Payload)
}

func (o *GetExtensionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetExtensionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionBadRequest creates a GetExtensionBadRequest with default headers values
func NewGetExtensionBadRequest() *GetExtensionBadRequest {
	return &GetExtensionBadRequest{}
}

/*GetExtensionBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type GetExtensionBadRequest struct {
	Payload *models.Error
}

func (o *GetExtensionBadRequest) Error() string {
	return fmt.Sprintf("[GET /extensions/{id}][%d] getExtensionBadRequest  %+v", 400, o.Payload)
}

func (o *GetExtensionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionUnauthorized creates a GetExtensionUnauthorized with default headers values
func NewGetExtensionUnauthorized() *GetExtensionUnauthorized {
	return &GetExtensionUnauthorized{}
}

/*GetExtensionUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type GetExtensionUnauthorized struct {
	Payload *models.Error
}

func (o *GetExtensionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /extensions/{id}][%d] getExtensionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExtensionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionForbidden creates a GetExtensionForbidden with default headers values
func NewGetExtensionForbidden() *GetExtensionForbidden {
	return &GetExtensionForbidden{}
}

/*GetExtensionForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type GetExtensionForbidden struct {
	Payload *models.Error
}

func (o *GetExtensionForbidden) Error() string {
	return fmt.Sprintf("[GET /extensions/{id}][%d] getExtensionForbidden  %+v", 403, o.Payload)
}

func (o *GetExtensionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionNotFound creates a GetExtensionNotFound with default headers values
func NewGetExtensionNotFound() *GetExtensionNotFound {
	return &GetExtensionNotFound{}
}

/*GetExtensionNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetExtensionNotFound struct {
	Payload *models.Error
}

func (o *GetExtensionNotFound) Error() string {
	return fmt.Sprintf("[GET /extensions/{id}][%d] getExtensionNotFound  %+v", 404, o.Payload)
}

func (o *GetExtensionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
