// Code generated by go-swagger; DO NOT EDIT.

package schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// PostSchedulesReader is a Reader for the PostSchedules structure.
type PostSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostSchedulesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostSchedulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostSchedulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostSchedulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostSchedulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSchedulesCreated creates a PostSchedulesCreated with default headers values
func NewPostSchedulesCreated() *PostSchedulesCreated {
	return &PostSchedulesCreated{}
}

/*PostSchedulesCreated handles this case with default header values.

The schedule object created.
*/
type PostSchedulesCreated struct {
	Payload *models.PostSchedulesCreatedBody
}

func (o *PostSchedulesCreated) Error() string {
	return fmt.Sprintf("[POST /schedules][%d] postSchedulesCreated  %+v", 201, o.Payload)
}

func (o *PostSchedulesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostSchedulesCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSchedulesBadRequest creates a PostSchedulesBadRequest with default headers values
func NewPostSchedulesBadRequest() *PostSchedulesBadRequest {
	return &PostSchedulesBadRequest{}
}

/*PostSchedulesBadRequest handles this case with default header values.

Caller provided invalid arguments. Please review the response for error details. Retrying with the same arguments will *not* work.
*/
type PostSchedulesBadRequest struct {
	Payload *models.Error
}

func (o *PostSchedulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /schedules][%d] postSchedulesBadRequest  %+v", 400, o.Payload)
}

func (o *PostSchedulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSchedulesUnauthorized creates a PostSchedulesUnauthorized with default headers values
func NewPostSchedulesUnauthorized() *PostSchedulesUnauthorized {
	return &PostSchedulesUnauthorized{}
}

/*PostSchedulesUnauthorized handles this case with default header values.

Caller did not supply credentials or did not provide the correct credentials.
If you are using an API key, it may be invalid or your Authorization header may be malformed.

*/
type PostSchedulesUnauthorized struct {
	Payload *models.Error
}

func (o *PostSchedulesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /schedules][%d] postSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSchedulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSchedulesForbidden creates a PostSchedulesForbidden with default headers values
func NewPostSchedulesForbidden() *PostSchedulesForbidden {
	return &PostSchedulesForbidden{}
}

/*PostSchedulesForbidden handles this case with default header values.

Caller is not authorized to view the requested resource.
While your authentication is valid, the authenticated user or token does not have permission to perform this action.

*/
type PostSchedulesForbidden struct {
	Payload *models.Error
}

func (o *PostSchedulesForbidden) Error() string {
	return fmt.Sprintf("[POST /schedules][%d] postSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *PostSchedulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSchedulesTooManyRequests creates a PostSchedulesTooManyRequests with default headers values
func NewPostSchedulesTooManyRequests() *PostSchedulesTooManyRequests {
	return &PostSchedulesTooManyRequests{}
}

/*PostSchedulesTooManyRequests handles this case with default header values.

Too many requests have been made, the rate limit has been reached.
*/
type PostSchedulesTooManyRequests struct {
	Payload *models.Error
}

func (o *PostSchedulesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /schedules][%d] postSchedulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostSchedulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}