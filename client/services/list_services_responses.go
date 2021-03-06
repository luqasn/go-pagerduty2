// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// ListServicesReader is a Reader for the ListServices structure.
type ListServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListServicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListServicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListServicesOK creates a ListServicesOK with default headers values
func NewListServicesOK() *ListServicesOK {
	return &ListServicesOK{}
}

/*ListServicesOK handles this case with default header values.

A paginated array of services.
*/
type ListServicesOK struct {
	Payload *models.ListServicesOKBody
}

func (o *ListServicesOK) Error() string {
	return fmt.Sprintf("[GET /services][%d] listServicesOK  %+v", 200, o.Payload)
}

func (o *ListServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListServicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesBadRequest creates a ListServicesBadRequest with default headers values
func NewListServicesBadRequest() *ListServicesBadRequest {
	return &ListServicesBadRequest{}
}

/*ListServicesBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type ListServicesBadRequest struct {
	Payload *models.Error
}

func (o *ListServicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /services][%d] listServicesBadRequest  %+v", 400, o.Payload)
}

func (o *ListServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesUnauthorized creates a ListServicesUnauthorized with default headers values
func NewListServicesUnauthorized() *ListServicesUnauthorized {
	return &ListServicesUnauthorized{}
}

/*ListServicesUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type ListServicesUnauthorized struct {
	Payload *models.Error
}

func (o *ListServicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /services][%d] listServicesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListServicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesForbidden creates a ListServicesForbidden with default headers values
func NewListServicesForbidden() *ListServicesForbidden {
	return &ListServicesForbidden{}
}

/*ListServicesForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type ListServicesForbidden struct {
	Payload *models.Error
}

func (o *ListServicesForbidden) Error() string {
	return fmt.Sprintf("[GET /services][%d] listServicesForbidden  %+v", 403, o.Payload)
}

func (o *ListServicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
