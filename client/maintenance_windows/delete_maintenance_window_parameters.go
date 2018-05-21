// Code generated by go-swagger; DO NOT EDIT.

package maintenance_windows

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

// NewDeleteMaintenanceWindowParams creates a new DeleteMaintenanceWindowParams object
// with the default values initialized.
func NewDeleteMaintenanceWindowParams() *DeleteMaintenanceWindowParams {
	var ()
	return &DeleteMaintenanceWindowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMaintenanceWindowParamsWithTimeout creates a new DeleteMaintenanceWindowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMaintenanceWindowParamsWithTimeout(timeout time.Duration) *DeleteMaintenanceWindowParams {
	var ()
	return &DeleteMaintenanceWindowParams{

		timeout: timeout,
	}
}

// NewDeleteMaintenanceWindowParamsWithContext creates a new DeleteMaintenanceWindowParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMaintenanceWindowParamsWithContext(ctx context.Context) *DeleteMaintenanceWindowParams {
	var ()
	return &DeleteMaintenanceWindowParams{

		Context: ctx,
	}
}

// NewDeleteMaintenanceWindowParamsWithHTTPClient creates a new DeleteMaintenanceWindowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMaintenanceWindowParamsWithHTTPClient(client *http.Client) *DeleteMaintenanceWindowParams {
	var ()
	return &DeleteMaintenanceWindowParams{
		HTTPClient: client,
	}
}

/*DeleteMaintenanceWindowParams contains all the parameters to send to the API endpoint
for the delete maintenance window operation typically these are written to a http.Request
*/
type DeleteMaintenanceWindowParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete maintenance window params
func (o *DeleteMaintenanceWindowParams) WithTimeout(timeout time.Duration) *DeleteMaintenanceWindowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete maintenance window params
func (o *DeleteMaintenanceWindowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete maintenance window params
func (o *DeleteMaintenanceWindowParams) WithContext(ctx context.Context) *DeleteMaintenanceWindowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete maintenance window params
func (o *DeleteMaintenanceWindowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete maintenance window params
func (o *DeleteMaintenanceWindowParams) WithHTTPClient(client *http.Client) *DeleteMaintenanceWindowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete maintenance window params
func (o *DeleteMaintenanceWindowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete maintenance window params
func (o *DeleteMaintenanceWindowParams) WithID(id string) *DeleteMaintenanceWindowParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete maintenance window params
func (o *DeleteMaintenanceWindowParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMaintenanceWindowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
