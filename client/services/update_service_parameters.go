// Code generated by go-swagger; DO NOT EDIT.

package services

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

// NewUpdateServiceParams creates a new UpdateServiceParams object
// with the default values initialized.
func NewUpdateServiceParams() *UpdateServiceParams {
	var ()
	return &UpdateServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceParamsWithTimeout creates a new UpdateServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateServiceParamsWithTimeout(timeout time.Duration) *UpdateServiceParams {
	var ()
	return &UpdateServiceParams{

		timeout: timeout,
	}
}

// NewUpdateServiceParamsWithContext creates a new UpdateServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateServiceParamsWithContext(ctx context.Context) *UpdateServiceParams {
	var ()
	return &UpdateServiceParams{

		Context: ctx,
	}
}

// NewUpdateServiceParamsWithHTTPClient creates a new UpdateServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateServiceParamsWithHTTPClient(client *http.Client) *UpdateServiceParams {
	var ()
	return &UpdateServiceParams{
		HTTPClient: client,
	}
}

/*UpdateServiceParams contains all the parameters to send to the API endpoint
for the update service operation typically these are written to a http.Request
*/
type UpdateServiceParams struct {

	/*ID*/
	ID string
	/*Service
	  The service to be updated.

	*/
	Service *models.UpdateServiceParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update service params
func (o *UpdateServiceParams) WithTimeout(timeout time.Duration) *UpdateServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service params
func (o *UpdateServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service params
func (o *UpdateServiceParams) WithContext(ctx context.Context) *UpdateServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service params
func (o *UpdateServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service params
func (o *UpdateServiceParams) WithHTTPClient(client *http.Client) *UpdateServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service params
func (o *UpdateServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update service params
func (o *UpdateServiceParams) WithID(id string) *UpdateServiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update service params
func (o *UpdateServiceParams) SetID(id string) {
	o.ID = id
}

// WithService adds the service to the update service params
func (o *UpdateServiceParams) WithService(service *models.UpdateServiceParamsBody) *UpdateServiceParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the update service params
func (o *UpdateServiceParams) SetService(service *models.UpdateServiceParamsBody) {
	o.Service = service
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Service != nil {
		if err := r.SetBodyParam(o.Service); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
