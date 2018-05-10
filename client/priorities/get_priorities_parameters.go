// Code generated by go-swagger; DO NOT EDIT.

package priorities

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

// NewGetPrioritiesParams creates a new GetPrioritiesParams object
// with the default values initialized.
func NewGetPrioritiesParams() *GetPrioritiesParams {

	return &GetPrioritiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrioritiesParamsWithTimeout creates a new GetPrioritiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrioritiesParamsWithTimeout(timeout time.Duration) *GetPrioritiesParams {

	return &GetPrioritiesParams{

		timeout: timeout,
	}
}

// NewGetPrioritiesParamsWithContext creates a new GetPrioritiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrioritiesParamsWithContext(ctx context.Context) *GetPrioritiesParams {

	return &GetPrioritiesParams{

		Context: ctx,
	}
}

// NewGetPrioritiesParamsWithHTTPClient creates a new GetPrioritiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrioritiesParamsWithHTTPClient(client *http.Client) *GetPrioritiesParams {

	return &GetPrioritiesParams{
		HTTPClient: client,
	}
}

/*GetPrioritiesParams contains all the parameters to send to the API endpoint
for the get priorities operation typically these are written to a http.Request
*/
type GetPrioritiesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get priorities params
func (o *GetPrioritiesParams) WithTimeout(timeout time.Duration) *GetPrioritiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get priorities params
func (o *GetPrioritiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get priorities params
func (o *GetPrioritiesParams) WithContext(ctx context.Context) *GetPrioritiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get priorities params
func (o *GetPrioritiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get priorities params
func (o *GetPrioritiesParams) WithHTTPClient(client *http.Client) *GetPrioritiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get priorities params
func (o *GetPrioritiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrioritiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}