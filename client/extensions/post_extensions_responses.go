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

// PostExtensionsReader is a Reader for the PostExtensions structure.
type PostExtensionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExtensionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostExtensionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostExtensionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostExtensionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewPostExtensionsPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostExtensionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostExtensionsCreated creates a PostExtensionsCreated with default headers values
func NewPostExtensionsCreated() *PostExtensionsCreated {
	return &PostExtensionsCreated{}
}

/*PostExtensionsCreated handles this case with default header values.

The extension that was created
*/
type PostExtensionsCreated struct {
	Payload *models.PostExtensionsCreatedBody
}

func (o *PostExtensionsCreated) Error() string {
	return fmt.Sprintf("[POST /extensions][%d] postExtensionsCreated  %+v", 201, o.Payload)
}

func (o *PostExtensionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostExtensionsCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExtensionsBadRequest creates a PostExtensionsBadRequest with default headers values
func NewPostExtensionsBadRequest() *PostExtensionsBadRequest {
	return &PostExtensionsBadRequest{}
}

/*PostExtensionsBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type PostExtensionsBadRequest struct {
	Payload *models.Error
}

func (o *PostExtensionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /extensions][%d] postExtensionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostExtensionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExtensionsUnauthorized creates a PostExtensionsUnauthorized with default headers values
func NewPostExtensionsUnauthorized() *PostExtensionsUnauthorized {
	return &PostExtensionsUnauthorized{}
}

/*PostExtensionsUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type PostExtensionsUnauthorized struct {
	Payload *models.Error
}

func (o *PostExtensionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /extensions][%d] postExtensionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExtensionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExtensionsPaymentRequired creates a PostExtensionsPaymentRequired with default headers values
func NewPostExtensionsPaymentRequired() *PostExtensionsPaymentRequired {
	return &PostExtensionsPaymentRequired{}
}

/*PostExtensionsPaymentRequired handles this case with default header values.

Account does not have the abilities to perform the action. Please review the response for the required abilities.
You can also use the [Abilities API](#resource_Abilities) to determine what features are available to your account.

*/
type PostExtensionsPaymentRequired struct {
	Payload *models.Error
}

func (o *PostExtensionsPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /extensions][%d] postExtensionsPaymentRequired  %+v", 402, o.Payload)
}

func (o *PostExtensionsPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExtensionsForbidden creates a PostExtensionsForbidden with default headers values
func NewPostExtensionsForbidden() *PostExtensionsForbidden {
	return &PostExtensionsForbidden{}
}

/*PostExtensionsForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type PostExtensionsForbidden struct {
	Payload *models.Error
}

func (o *PostExtensionsForbidden) Error() string {
	return fmt.Sprintf("[POST /extensions][%d] postExtensionsForbidden  %+v", 403, o.Payload)
}

func (o *PostExtensionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}