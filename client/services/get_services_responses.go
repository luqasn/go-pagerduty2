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

// GetServicesReader is a Reader for the GetServices structure.
type GetServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetServicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetServicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServicesOK creates a GetServicesOK with default headers values
func NewGetServicesOK() *GetServicesOK {
	return &GetServicesOK{}
}

/*GetServicesOK handles this case with default header values.

A paginated array of services.
*/
type GetServicesOK struct {
	Payload *models.GetServicesOKBody
}

func (o *GetServicesOK) Error() string {
	return fmt.Sprintf("[GET /services][%d] getServicesOK  %+v", 200, o.Payload)
}

func (o *GetServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetServicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesBadRequest creates a GetServicesBadRequest with default headers values
func NewGetServicesBadRequest() *GetServicesBadRequest {
	return &GetServicesBadRequest{}
}

/*GetServicesBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type GetServicesBadRequest struct {
	Payload *models.Error
}

func (o *GetServicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /services][%d] getServicesBadRequest  %+v", 400, o.Payload)
}

func (o *GetServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesUnauthorized creates a GetServicesUnauthorized with default headers values
func NewGetServicesUnauthorized() *GetServicesUnauthorized {
	return &GetServicesUnauthorized{}
}

/*GetServicesUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type GetServicesUnauthorized struct {
	Payload *models.Error
}

func (o *GetServicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /services][%d] getServicesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetServicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesForbidden creates a GetServicesForbidden with default headers values
func NewGetServicesForbidden() *GetServicesForbidden {
	return &GetServicesForbidden{}
}

/*GetServicesForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type GetServicesForbidden struct {
	Payload *models.Error
}

func (o *GetServicesForbidden) Error() string {
	return fmt.Sprintf("[GET /services][%d] getServicesForbidden  %+v", 403, o.Payload)
}

func (o *GetServicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
