// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// ListTeamsReader is a Reader for the ListTeams structure.
type ListTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListTeamsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListTeamsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewListTeamsPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListTeamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewListTeamsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListTeamsOK creates a ListTeamsOK with default headers values
func NewListTeamsOK() *ListTeamsOK {
	return &ListTeamsOK{}
}

/*ListTeamsOK handles this case with default header values.

A paginated array of teams.
*/
type ListTeamsOK struct {
	Payload *models.ListTeamsOKBody
}

func (o *ListTeamsOK) Error() string {
	return fmt.Sprintf("[GET /teams][%d] listTeamsOK  %+v", 200, o.Payload)
}

func (o *ListTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListTeamsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamsBadRequest creates a ListTeamsBadRequest with default headers values
func NewListTeamsBadRequest() *ListTeamsBadRequest {
	return &ListTeamsBadRequest{}
}

/*ListTeamsBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type ListTeamsBadRequest struct {
	Payload *models.Error
}

func (o *ListTeamsBadRequest) Error() string {
	return fmt.Sprintf("[GET /teams][%d] listTeamsBadRequest  %+v", 400, o.Payload)
}

func (o *ListTeamsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamsUnauthorized creates a ListTeamsUnauthorized with default headers values
func NewListTeamsUnauthorized() *ListTeamsUnauthorized {
	return &ListTeamsUnauthorized{}
}

/*ListTeamsUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type ListTeamsUnauthorized struct {
	Payload *models.Error
}

func (o *ListTeamsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /teams][%d] listTeamsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListTeamsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamsPaymentRequired creates a ListTeamsPaymentRequired with default headers values
func NewListTeamsPaymentRequired() *ListTeamsPaymentRequired {
	return &ListTeamsPaymentRequired{}
}

/*ListTeamsPaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type ListTeamsPaymentRequired struct {
	Payload *models.Error
}

func (o *ListTeamsPaymentRequired) Error() string {
	return fmt.Sprintf("[GET /teams][%d] listTeamsPaymentRequired  %+v", 402, o.Payload)
}

func (o *ListTeamsPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamsForbidden creates a ListTeamsForbidden with default headers values
func NewListTeamsForbidden() *ListTeamsForbidden {
	return &ListTeamsForbidden{}
}

/*ListTeamsForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type ListTeamsForbidden struct {
	Payload *models.Error
}

func (o *ListTeamsForbidden) Error() string {
	return fmt.Sprintf("[GET /teams][%d] listTeamsForbidden  %+v", 403, o.Payload)
}

func (o *ListTeamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamsTooManyRequests creates a ListTeamsTooManyRequests with default headers values
func NewListTeamsTooManyRequests() *ListTeamsTooManyRequests {
	return &ListTeamsTooManyRequests{}
}

/*ListTeamsTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type ListTeamsTooManyRequests struct {
	Payload *models.Error
}

func (o *ListTeamsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /teams][%d] listTeamsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListTeamsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}