// Code generated by go-swagger; DO NOT EDIT.

package vendors

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

// NewGetVendorsIDParams creates a new GetVendorsIDParams object
// with the default values initialized.
func NewGetVendorsIDParams() *GetVendorsIDParams {
	var ()
	return &GetVendorsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVendorsIDParamsWithTimeout creates a new GetVendorsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVendorsIDParamsWithTimeout(timeout time.Duration) *GetVendorsIDParams {
	var ()
	return &GetVendorsIDParams{

		timeout: timeout,
	}
}

// NewGetVendorsIDParamsWithContext creates a new GetVendorsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVendorsIDParamsWithContext(ctx context.Context) *GetVendorsIDParams {
	var ()
	return &GetVendorsIDParams{

		Context: ctx,
	}
}

// NewGetVendorsIDParamsWithHTTPClient creates a new GetVendorsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVendorsIDParamsWithHTTPClient(client *http.Client) *GetVendorsIDParams {
	var ()
	return &GetVendorsIDParams{
		HTTPClient: client,
	}
}

/*GetVendorsIDParams contains all the parameters to send to the API endpoint
for the get vendors ID operation typically these are written to a http.Request
*/
type GetVendorsIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vendors ID params
func (o *GetVendorsIDParams) WithTimeout(timeout time.Duration) *GetVendorsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vendors ID params
func (o *GetVendorsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vendors ID params
func (o *GetVendorsIDParams) WithContext(ctx context.Context) *GetVendorsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vendors ID params
func (o *GetVendorsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vendors ID params
func (o *GetVendorsIDParams) WithHTTPClient(client *http.Client) *GetVendorsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vendors ID params
func (o *GetVendorsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get vendors ID params
func (o *GetVendorsIDParams) WithID(id string) *GetVendorsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vendors ID params
func (o *GetVendorsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVendorsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}