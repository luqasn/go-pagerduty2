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

// ListExtensionsReader is a Reader for the ListExtensions structure.
type ListExtensionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListExtensionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListExtensionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListExtensionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListExtensionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListExtensionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListExtensionsOK creates a ListExtensionsOK with default headers values
func NewListExtensionsOK() *ListExtensionsOK {
	return &ListExtensionsOK{}
}

/*ListExtensionsOK handles this case with default header values.

A paginated array of extensions.
*/
type ListExtensionsOK struct {
	Payload *models.ListExtensionsOKBody
}

func (o *ListExtensionsOK) Error() string {
	return fmt.Sprintf("[GET /extensions][%d] listExtensionsOK  %+v", 200, o.Payload)
}

func (o *ListExtensionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListExtensionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExtensionsBadRequest creates a ListExtensionsBadRequest with default headers values
func NewListExtensionsBadRequest() *ListExtensionsBadRequest {
	return &ListExtensionsBadRequest{}
}

/*ListExtensionsBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type ListExtensionsBadRequest struct {
	Payload *models.Error
}

func (o *ListExtensionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /extensions][%d] listExtensionsBadRequest  %+v", 400, o.Payload)
}

func (o *ListExtensionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExtensionsUnauthorized creates a ListExtensionsUnauthorized with default headers values
func NewListExtensionsUnauthorized() *ListExtensionsUnauthorized {
	return &ListExtensionsUnauthorized{}
}

/*ListExtensionsUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type ListExtensionsUnauthorized struct {
	Payload *models.Error
}

func (o *ListExtensionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /extensions][%d] listExtensionsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListExtensionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExtensionsForbidden creates a ListExtensionsForbidden with default headers values
func NewListExtensionsForbidden() *ListExtensionsForbidden {
	return &ListExtensionsForbidden{}
}

/*ListExtensionsForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type ListExtensionsForbidden struct {
	Payload *models.Error
}

func (o *ListExtensionsForbidden) Error() string {
	return fmt.Sprintf("[GET /extensions][%d] listExtensionsForbidden  %+v", 403, o.Payload)
}

func (o *ListExtensionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
