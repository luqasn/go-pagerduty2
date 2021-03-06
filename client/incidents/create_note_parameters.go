// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// NewCreateNoteParams creates a new CreateNoteParams object
// with the default values initialized.
func NewCreateNoteParams() *CreateNoteParams {
	var ()
	return &CreateNoteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNoteParamsWithTimeout creates a new CreateNoteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateNoteParamsWithTimeout(timeout time.Duration) *CreateNoteParams {
	var ()
	return &CreateNoteParams{

		timeout: timeout,
	}
}

// NewCreateNoteParamsWithContext creates a new CreateNoteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateNoteParamsWithContext(ctx context.Context) *CreateNoteParams {
	var ()
	return &CreateNoteParams{

		Context: ctx,
	}
}

// NewCreateNoteParamsWithHTTPClient creates a new CreateNoteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateNoteParamsWithHTTPClient(client *http.Client) *CreateNoteParams {
	var ()
	return &CreateNoteParams{
		HTTPClient: client,
	}
}

/*CreateNoteParams contains all the parameters to send to the API endpoint
for the create note operation typically these are written to a http.Request
*/
type CreateNoteParams struct {

	/*From
	  The email address of the user making the request.

	*/
	From strfmt.Email
	/*ID*/
	ID string
	/*Payload*/
	Payload *models.CreateNoteParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create note params
func (o *CreateNoteParams) WithTimeout(timeout time.Duration) *CreateNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create note params
func (o *CreateNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create note params
func (o *CreateNoteParams) WithContext(ctx context.Context) *CreateNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create note params
func (o *CreateNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create note params
func (o *CreateNoteParams) WithHTTPClient(client *http.Client) *CreateNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create note params
func (o *CreateNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the create note params
func (o *CreateNoteParams) WithFrom(from strfmt.Email) *CreateNoteParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the create note params
func (o *CreateNoteParams) SetFrom(from strfmt.Email) {
	o.From = from
}

// WithID adds the id to the create note params
func (o *CreateNoteParams) WithID(id string) *CreateNoteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create note params
func (o *CreateNoteParams) SetID(id string) {
	o.ID = id
}

// WithPayload adds the payload to the create note params
func (o *CreateNoteParams) WithPayload(payload *models.CreateNoteParamsBody) *CreateNoteParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the create note params
func (o *CreateNoteParams) SetPayload(payload *models.CreateNoteParamsBody) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param From
	if err := r.SetHeaderParam("From", o.From.String()); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
