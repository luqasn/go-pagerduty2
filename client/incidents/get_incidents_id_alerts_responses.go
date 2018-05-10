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

// GetIncidentsIDAlertsReader is a Reader for the GetIncidentsIDAlerts structure.
type GetIncidentsIDAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIncidentsIDAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIncidentsIDAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetIncidentsIDAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetIncidentsIDAlertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetIncidentsIDAlertsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetIncidentsIDAlertsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIncidentsIDAlertsOK creates a GetIncidentsIDAlertsOK with default headers values
func NewGetIncidentsIDAlertsOK() *GetIncidentsIDAlertsOK {
	return &GetIncidentsIDAlertsOK{}
}

/*GetIncidentsIDAlertsOK handles this case with default header values.

A paginated array of the incident's alerts.
*/
type GetIncidentsIDAlertsOK struct {
	Payload *models.GetIncidentsIDAlertsOKBody
}

func (o *GetIncidentsIDAlertsOK) Error() string {
	return fmt.Sprintf("[GET /incidents/{id}/alerts][%d] getIncidentsIdAlertsOK  %+v", 200, o.Payload)
}

func (o *GetIncidentsIDAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIncidentsIDAlertsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncidentsIDAlertsBadRequest creates a GetIncidentsIDAlertsBadRequest with default headers values
func NewGetIncidentsIDAlertsBadRequest() *GetIncidentsIDAlertsBadRequest {
	return &GetIncidentsIDAlertsBadRequest{}
}

/*GetIncidentsIDAlertsBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type GetIncidentsIDAlertsBadRequest struct {
	Payload *models.Error
}

func (o *GetIncidentsIDAlertsBadRequest) Error() string {
	return fmt.Sprintf("[GET /incidents/{id}/alerts][%d] getIncidentsIdAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIncidentsIDAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncidentsIDAlertsUnauthorized creates a GetIncidentsIDAlertsUnauthorized with default headers values
func NewGetIncidentsIDAlertsUnauthorized() *GetIncidentsIDAlertsUnauthorized {
	return &GetIncidentsIDAlertsUnauthorized{}
}

/*GetIncidentsIDAlertsUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type GetIncidentsIDAlertsUnauthorized struct {
	Payload *models.Error
}

func (o *GetIncidentsIDAlertsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /incidents/{id}/alerts][%d] getIncidentsIdAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIncidentsIDAlertsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncidentsIDAlertsForbidden creates a GetIncidentsIDAlertsForbidden with default headers values
func NewGetIncidentsIDAlertsForbidden() *GetIncidentsIDAlertsForbidden {
	return &GetIncidentsIDAlertsForbidden{}
}

/*GetIncidentsIDAlertsForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type GetIncidentsIDAlertsForbidden struct {
	Payload *models.Error
}

func (o *GetIncidentsIDAlertsForbidden) Error() string {
	return fmt.Sprintf("[GET /incidents/{id}/alerts][%d] getIncidentsIdAlertsForbidden  %+v", 403, o.Payload)
}

func (o *GetIncidentsIDAlertsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncidentsIDAlertsTooManyRequests creates a GetIncidentsIDAlertsTooManyRequests with default headers values
func NewGetIncidentsIDAlertsTooManyRequests() *GetIncidentsIDAlertsTooManyRequests {
	return &GetIncidentsIDAlertsTooManyRequests{}
}

/*GetIncidentsIDAlertsTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type GetIncidentsIDAlertsTooManyRequests struct {
	Payload *models.Error
}

func (o *GetIncidentsIDAlertsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /incidents/{id}/alerts][%d] getIncidentsIdAlertsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIncidentsIDAlertsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
