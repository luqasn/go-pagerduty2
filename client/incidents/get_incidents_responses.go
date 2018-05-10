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

// GetIncidentsReader is a Reader for the GetIncidents structure.
type GetIncidentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIncidentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIncidentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetIncidentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetIncidentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewGetIncidentsPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetIncidentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetIncidentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIncidentsOK creates a GetIncidentsOK with default headers values
func NewGetIncidentsOK() *GetIncidentsOK {
	return &GetIncidentsOK{}
}

/*GetIncidentsOK handles this case with default header values.

A paginated array of incidents.
*/
type GetIncidentsOK struct {
	Payload *models.GetIncidentsOKBody
}

func (o *GetIncidentsOK) Error() string {
	return fmt.Sprintf("[GET /incidents][%d] getIncidentsOK  %+v", 200, o.Payload)
}

func (o *GetIncidentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIncidentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncidentsBadRequest creates a GetIncidentsBadRequest with default headers values
func NewGetIncidentsBadRequest() *GetIncidentsBadRequest {
	return &GetIncidentsBadRequest{}
}

/*GetIncidentsBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type GetIncidentsBadRequest struct {
	Payload *models.Error
}

func (o *GetIncidentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /incidents][%d] getIncidentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIncidentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncidentsUnauthorized creates a GetIncidentsUnauthorized with default headers values
func NewGetIncidentsUnauthorized() *GetIncidentsUnauthorized {
	return &GetIncidentsUnauthorized{}
}

/*GetIncidentsUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type GetIncidentsUnauthorized struct {
	Payload *models.Error
}

func (o *GetIncidentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /incidents][%d] getIncidentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIncidentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncidentsPaymentRequired creates a GetIncidentsPaymentRequired with default headers values
func NewGetIncidentsPaymentRequired() *GetIncidentsPaymentRequired {
	return &GetIncidentsPaymentRequired{}
}

/*GetIncidentsPaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type GetIncidentsPaymentRequired struct {
	Payload *models.Error
}

func (o *GetIncidentsPaymentRequired) Error() string {
	return fmt.Sprintf("[GET /incidents][%d] getIncidentsPaymentRequired  %+v", 402, o.Payload)
}

func (o *GetIncidentsPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncidentsForbidden creates a GetIncidentsForbidden with default headers values
func NewGetIncidentsForbidden() *GetIncidentsForbidden {
	return &GetIncidentsForbidden{}
}

/*GetIncidentsForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type GetIncidentsForbidden struct {
	Payload *models.Error
}

func (o *GetIncidentsForbidden) Error() string {
	return fmt.Sprintf("[GET /incidents][%d] getIncidentsForbidden  %+v", 403, o.Payload)
}

func (o *GetIncidentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncidentsTooManyRequests creates a GetIncidentsTooManyRequests with default headers values
func NewGetIncidentsTooManyRequests() *GetIncidentsTooManyRequests {
	return &GetIncidentsTooManyRequests{}
}

/*GetIncidentsTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type GetIncidentsTooManyRequests struct {
	Payload *models.Error
}

func (o *GetIncidentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /incidents][%d] getIncidentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIncidentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
