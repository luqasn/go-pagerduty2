// Code generated by go-swagger; DO NOT EDIT.

package users

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
)

// NewGetUsersIDContactMethodsContactMethodIDParams creates a new GetUsersIDContactMethodsContactMethodIDParams object
// with the default values initialized.
func NewGetUsersIDContactMethodsContactMethodIDParams() *GetUsersIDContactMethodsContactMethodIDParams {
	var ()
	return &GetUsersIDContactMethodsContactMethodIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersIDContactMethodsContactMethodIDParamsWithTimeout creates a new GetUsersIDContactMethodsContactMethodIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersIDContactMethodsContactMethodIDParamsWithTimeout(timeout time.Duration) *GetUsersIDContactMethodsContactMethodIDParams {
	var ()
	return &GetUsersIDContactMethodsContactMethodIDParams{

		timeout: timeout,
	}
}

// NewGetUsersIDContactMethodsContactMethodIDParamsWithContext creates a new GetUsersIDContactMethodsContactMethodIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersIDContactMethodsContactMethodIDParamsWithContext(ctx context.Context) *GetUsersIDContactMethodsContactMethodIDParams {
	var ()
	return &GetUsersIDContactMethodsContactMethodIDParams{

		Context: ctx,
	}
}

// NewGetUsersIDContactMethodsContactMethodIDParamsWithHTTPClient creates a new GetUsersIDContactMethodsContactMethodIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersIDContactMethodsContactMethodIDParamsWithHTTPClient(client *http.Client) *GetUsersIDContactMethodsContactMethodIDParams {
	var ()
	return &GetUsersIDContactMethodsContactMethodIDParams{
		HTTPClient: client,
	}
}

/*GetUsersIDContactMethodsContactMethodIDParams contains all the parameters to send to the API endpoint
for the get users ID contact methods contact method ID operation typically these are written to a http.Request
*/
type GetUsersIDContactMethodsContactMethodIDParams struct {

	/*ContactMethodID
	  The contact method ID on the user.

	*/
	ContactMethodID string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users ID contact methods contact method ID params
func (o *GetUsersIDContactMethodsContactMethodIDParams) WithTimeout(timeout time.Duration) *GetUsersIDContactMethodsContactMethodIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users ID contact methods contact method ID params
func (o *GetUsersIDContactMethodsContactMethodIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users ID contact methods contact method ID params
func (o *GetUsersIDContactMethodsContactMethodIDParams) WithContext(ctx context.Context) *GetUsersIDContactMethodsContactMethodIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users ID contact methods contact method ID params
func (o *GetUsersIDContactMethodsContactMethodIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users ID contact methods contact method ID params
func (o *GetUsersIDContactMethodsContactMethodIDParams) WithHTTPClient(client *http.Client) *GetUsersIDContactMethodsContactMethodIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users ID contact methods contact method ID params
func (o *GetUsersIDContactMethodsContactMethodIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactMethodID adds the contactMethodID to the get users ID contact methods contact method ID params
func (o *GetUsersIDContactMethodsContactMethodIDParams) WithContactMethodID(contactMethodID string) *GetUsersIDContactMethodsContactMethodIDParams {
	o.SetContactMethodID(contactMethodID)
	return o
}

// SetContactMethodID adds the contactMethodId to the get users ID contact methods contact method ID params
func (o *GetUsersIDContactMethodsContactMethodIDParams) SetContactMethodID(contactMethodID string) {
	o.ContactMethodID = contactMethodID
}

// WithID adds the id to the get users ID contact methods contact method ID params
func (o *GetUsersIDContactMethodsContactMethodIDParams) WithID(id string) *GetUsersIDContactMethodsContactMethodIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get users ID contact methods contact method ID params
func (o *GetUsersIDContactMethodsContactMethodIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersIDContactMethodsContactMethodIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contact_method_id
	if err := r.SetPathParam("contact_method_id", o.ContactMethodID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
