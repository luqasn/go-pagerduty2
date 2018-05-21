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

// CreateTeamReader is a Reader for the CreateTeam structure.
type CreateTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateTeamCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateTeamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateTeamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewCreateTeamPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewCreateTeamTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateTeamCreated creates a CreateTeamCreated with default headers values
func NewCreateTeamCreated() *CreateTeamCreated {
	return &CreateTeamCreated{}
}

/*CreateTeamCreated handles this case with default header values.

The team that was created.
*/
type CreateTeamCreated struct {
	Payload *models.CreateTeamCreatedBody
}

func (o *CreateTeamCreated) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamCreated  %+v", 201, o.Payload)
}

func (o *CreateTeamCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateTeamCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamBadRequest creates a CreateTeamBadRequest with default headers values
func NewCreateTeamBadRequest() *CreateTeamBadRequest {
	return &CreateTeamBadRequest{}
}

/*CreateTeamBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type CreateTeamBadRequest struct {
	Payload *models.Error
}

func (o *CreateTeamBadRequest) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTeamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamUnauthorized creates a CreateTeamUnauthorized with default headers values
func NewCreateTeamUnauthorized() *CreateTeamUnauthorized {
	return &CreateTeamUnauthorized{}
}

/*CreateTeamUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type CreateTeamUnauthorized struct {
	Payload *models.Error
}

func (o *CreateTeamUnauthorized) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateTeamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamPaymentRequired creates a CreateTeamPaymentRequired with default headers values
func NewCreateTeamPaymentRequired() *CreateTeamPaymentRequired {
	return &CreateTeamPaymentRequired{}
}

/*CreateTeamPaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type CreateTeamPaymentRequired struct {
	Payload *models.Error
}

func (o *CreateTeamPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamPaymentRequired  %+v", 402, o.Payload)
}

func (o *CreateTeamPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamForbidden creates a CreateTeamForbidden with default headers values
func NewCreateTeamForbidden() *CreateTeamForbidden {
	return &CreateTeamForbidden{}
}

/*CreateTeamForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type CreateTeamForbidden struct {
	Payload *models.Error
}

func (o *CreateTeamForbidden) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamForbidden  %+v", 403, o.Payload)
}

func (o *CreateTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamTooManyRequests creates a CreateTeamTooManyRequests with default headers values
func NewCreateTeamTooManyRequests() *CreateTeamTooManyRequests {
	return &CreateTeamTooManyRequests{}
}

/*CreateTeamTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type CreateTeamTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateTeamTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateTeamTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}