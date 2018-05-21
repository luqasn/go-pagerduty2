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

// ListIncidentsReader is a Reader for the ListIncidents structure.
type ListIncidentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIncidentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListIncidentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListIncidentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListIncidentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewListIncidentsPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListIncidentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewListIncidentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListIncidentsOK creates a ListIncidentsOK with default headers values
func NewListIncidentsOK() *ListIncidentsOK {
	return &ListIncidentsOK{}
}

/*ListIncidentsOK handles this case with default header values.

A paginated array of incidents.
*/
type ListIncidentsOK struct {
	Payload *models.ListIncidentsOKBody
}

func (o *ListIncidentsOK) Error() string {
	return fmt.Sprintf("[GET /incidents][%d] listIncidentsOK  %+v", 200, o.Payload)
}

func (o *ListIncidentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListIncidentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIncidentsBadRequest creates a ListIncidentsBadRequest with default headers values
func NewListIncidentsBadRequest() *ListIncidentsBadRequest {
	return &ListIncidentsBadRequest{}
}

/*ListIncidentsBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type ListIncidentsBadRequest struct {
	Payload *models.Error
}

func (o *ListIncidentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /incidents][%d] listIncidentsBadRequest  %+v", 400, o.Payload)
}

func (o *ListIncidentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIncidentsUnauthorized creates a ListIncidentsUnauthorized with default headers values
func NewListIncidentsUnauthorized() *ListIncidentsUnauthorized {
	return &ListIncidentsUnauthorized{}
}

/*ListIncidentsUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type ListIncidentsUnauthorized struct {
	Payload *models.Error
}

func (o *ListIncidentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /incidents][%d] listIncidentsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListIncidentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIncidentsPaymentRequired creates a ListIncidentsPaymentRequired with default headers values
func NewListIncidentsPaymentRequired() *ListIncidentsPaymentRequired {
	return &ListIncidentsPaymentRequired{}
}

/*ListIncidentsPaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type ListIncidentsPaymentRequired struct {
	Payload *models.Error
}

func (o *ListIncidentsPaymentRequired) Error() string {
	return fmt.Sprintf("[GET /incidents][%d] listIncidentsPaymentRequired  %+v", 402, o.Payload)
}

func (o *ListIncidentsPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIncidentsForbidden creates a ListIncidentsForbidden with default headers values
func NewListIncidentsForbidden() *ListIncidentsForbidden {
	return &ListIncidentsForbidden{}
}

/*ListIncidentsForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type ListIncidentsForbidden struct {
	Payload *models.Error
}

func (o *ListIncidentsForbidden) Error() string {
	return fmt.Sprintf("[GET /incidents][%d] listIncidentsForbidden  %+v", 403, o.Payload)
}

func (o *ListIncidentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIncidentsTooManyRequests creates a ListIncidentsTooManyRequests with default headers values
func NewListIncidentsTooManyRequests() *ListIncidentsTooManyRequests {
	return &ListIncidentsTooManyRequests{}
}

/*ListIncidentsTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type ListIncidentsTooManyRequests struct {
	Payload *models.Error
}

func (o *ListIncidentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /incidents][%d] listIncidentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListIncidentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}