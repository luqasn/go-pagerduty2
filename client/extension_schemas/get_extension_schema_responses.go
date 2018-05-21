// Code generated by go-swagger; DO NOT EDIT.

package extension_schemas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// GetExtensionSchemaReader is a Reader for the GetExtensionSchema structure.
type GetExtensionSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExtensionSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetExtensionSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetExtensionSchemaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetExtensionSchemaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetExtensionSchemaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetExtensionSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetExtensionSchemaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetExtensionSchemaOK creates a GetExtensionSchemaOK with default headers values
func NewGetExtensionSchemaOK() *GetExtensionSchemaOK {
	return &GetExtensionSchemaOK{}
}

/*GetExtensionSchemaOK handles this case with default header values.

The extension vendor requested
*/
type GetExtensionSchemaOK struct {
	Payload *models.GetExtensionSchemaOKBody
}

func (o *GetExtensionSchemaOK) Error() string {
	return fmt.Sprintf("[GET /extension_schemas/{id}][%d] getExtensionSchemaOK  %+v", 200, o.Payload)
}

func (o *GetExtensionSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetExtensionSchemaOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionSchemaBadRequest creates a GetExtensionSchemaBadRequest with default headers values
func NewGetExtensionSchemaBadRequest() *GetExtensionSchemaBadRequest {
	return &GetExtensionSchemaBadRequest{}
}

/*GetExtensionSchemaBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type GetExtensionSchemaBadRequest struct {
	Payload *models.Error
}

func (o *GetExtensionSchemaBadRequest) Error() string {
	return fmt.Sprintf("[GET /extension_schemas/{id}][%d] getExtensionSchemaBadRequest  %+v", 400, o.Payload)
}

func (o *GetExtensionSchemaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionSchemaUnauthorized creates a GetExtensionSchemaUnauthorized with default headers values
func NewGetExtensionSchemaUnauthorized() *GetExtensionSchemaUnauthorized {
	return &GetExtensionSchemaUnauthorized{}
}

/*GetExtensionSchemaUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type GetExtensionSchemaUnauthorized struct {
	Payload *models.Error
}

func (o *GetExtensionSchemaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /extension_schemas/{id}][%d] getExtensionSchemaUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExtensionSchemaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionSchemaForbidden creates a GetExtensionSchemaForbidden with default headers values
func NewGetExtensionSchemaForbidden() *GetExtensionSchemaForbidden {
	return &GetExtensionSchemaForbidden{}
}

/*GetExtensionSchemaForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type GetExtensionSchemaForbidden struct {
	Payload *models.Error
}

func (o *GetExtensionSchemaForbidden) Error() string {
	return fmt.Sprintf("[GET /extension_schemas/{id}][%d] getExtensionSchemaForbidden  %+v", 403, o.Payload)
}

func (o *GetExtensionSchemaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionSchemaNotFound creates a GetExtensionSchemaNotFound with default headers values
func NewGetExtensionSchemaNotFound() *GetExtensionSchemaNotFound {
	return &GetExtensionSchemaNotFound{}
}

/*GetExtensionSchemaNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetExtensionSchemaNotFound struct {
	Payload *models.Error
}

func (o *GetExtensionSchemaNotFound) Error() string {
	return fmt.Sprintf("[GET /extension_schemas/{id}][%d] getExtensionSchemaNotFound  %+v", 404, o.Payload)
}

func (o *GetExtensionSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionSchemaTooManyRequests creates a GetExtensionSchemaTooManyRequests with default headers values
func NewGetExtensionSchemaTooManyRequests() *GetExtensionSchemaTooManyRequests {
	return &GetExtensionSchemaTooManyRequests{}
}

/*GetExtensionSchemaTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type GetExtensionSchemaTooManyRequests struct {
	Payload *models.Error
}

func (o *GetExtensionSchemaTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /extension_schemas/{id}][%d] getExtensionSchemaTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExtensionSchemaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}