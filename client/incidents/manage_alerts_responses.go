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

// ManageAlertsReader is a Reader for the ManageAlerts structure.
type ManageAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ManageAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewManageAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewManageAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewManageAlertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewManageAlertsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewManageAlertsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewManageAlertsOK creates a ManageAlertsOK with default headers values
func NewManageAlertsOK() *ManageAlertsOK {
	return &ManageAlertsOK{}
}

/*ManageAlertsOK handles this case with default header values.

All of the updates succeeded.
*/
type ManageAlertsOK struct {
	Payload *models.ManageAlertsOKBody
}

func (o *ManageAlertsOK) Error() string {
	return fmt.Sprintf("[PUT /incidents/{id}/alerts][%d] manageAlertsOK  %+v", 200, o.Payload)
}

func (o *ManageAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManageAlertsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManageAlertsBadRequest creates a ManageAlertsBadRequest with default headers values
func NewManageAlertsBadRequest() *ManageAlertsBadRequest {
	return &ManageAlertsBadRequest{}
}

/*ManageAlertsBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type ManageAlertsBadRequest struct {
	Payload *models.Error
}

func (o *ManageAlertsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /incidents/{id}/alerts][%d] manageAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *ManageAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManageAlertsUnauthorized creates a ManageAlertsUnauthorized with default headers values
func NewManageAlertsUnauthorized() *ManageAlertsUnauthorized {
	return &ManageAlertsUnauthorized{}
}

/*ManageAlertsUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type ManageAlertsUnauthorized struct {
	Payload *models.Error
}

func (o *ManageAlertsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /incidents/{id}/alerts][%d] manageAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *ManageAlertsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManageAlertsForbidden creates a ManageAlertsForbidden with default headers values
func NewManageAlertsForbidden() *ManageAlertsForbidden {
	return &ManageAlertsForbidden{}
}

/*ManageAlertsForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type ManageAlertsForbidden struct {
	Payload *models.Error
}

func (o *ManageAlertsForbidden) Error() string {
	return fmt.Sprintf("[PUT /incidents/{id}/alerts][%d] manageAlertsForbidden  %+v", 403, o.Payload)
}

func (o *ManageAlertsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManageAlertsTooManyRequests creates a ManageAlertsTooManyRequests with default headers values
func NewManageAlertsTooManyRequests() *ManageAlertsTooManyRequests {
	return &ManageAlertsTooManyRequests{}
}

/*ManageAlertsTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type ManageAlertsTooManyRequests struct {
	Payload *models.Error
}

func (o *ManageAlertsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /incidents/{id}/alerts][%d] manageAlertsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ManageAlertsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}