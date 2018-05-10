// Code generated by go-swagger; DO NOT EDIT.

package abilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// GetAbilitiesReader is a Reader for the GetAbilities structure.
type GetAbilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAbilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAbilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAbilitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAbilitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetAbilitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAbilitiesOK creates a GetAbilitiesOK with default headers values
func NewGetAbilitiesOK() *GetAbilitiesOK {
	return &GetAbilitiesOK{}
}

/*GetAbilitiesOK handles this case with default header values.

An array of ability names.
*/
type GetAbilitiesOK struct {
	Payload *models.GetAbilitiesOKBody
}

func (o *GetAbilitiesOK) Error() string {
	return fmt.Sprintf("[GET /abilities][%d] getAbilitiesOK  %+v", 200, o.Payload)
}

func (o *GetAbilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAbilitiesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAbilitiesUnauthorized creates a GetAbilitiesUnauthorized with default headers values
func NewGetAbilitiesUnauthorized() *GetAbilitiesUnauthorized {
	return &GetAbilitiesUnauthorized{}
}

/*GetAbilitiesUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type GetAbilitiesUnauthorized struct {
	Payload *models.Error
}

func (o *GetAbilitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /abilities][%d] getAbilitiesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAbilitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAbilitiesForbidden creates a GetAbilitiesForbidden with default headers values
func NewGetAbilitiesForbidden() *GetAbilitiesForbidden {
	return &GetAbilitiesForbidden{}
}

/*GetAbilitiesForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type GetAbilitiesForbidden struct {
	Payload *models.Error
}

func (o *GetAbilitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /abilities][%d] getAbilitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetAbilitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAbilitiesTooManyRequests creates a GetAbilitiesTooManyRequests with default headers values
func NewGetAbilitiesTooManyRequests() *GetAbilitiesTooManyRequests {
	return &GetAbilitiesTooManyRequests{}
}

/*GetAbilitiesTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type GetAbilitiesTooManyRequests struct {
	Payload *models.Error
}

func (o *GetAbilitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /abilities][%d] getAbilitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAbilitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
