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

// NewListVendorsParams creates a new ListVendorsParams object
// with the default values initialized.
func NewListVendorsParams() *ListVendorsParams {

	return &ListVendorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListVendorsParamsWithTimeout creates a new ListVendorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListVendorsParamsWithTimeout(timeout time.Duration) *ListVendorsParams {

	return &ListVendorsParams{

		timeout: timeout,
	}
}

// NewListVendorsParamsWithContext creates a new ListVendorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListVendorsParamsWithContext(ctx context.Context) *ListVendorsParams {

	return &ListVendorsParams{

		Context: ctx,
	}
}

// NewListVendorsParamsWithHTTPClient creates a new ListVendorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListVendorsParamsWithHTTPClient(client *http.Client) *ListVendorsParams {

	return &ListVendorsParams{
		HTTPClient: client,
	}
}

/*ListVendorsParams contains all the parameters to send to the API endpoint
for the list vendors operation typically these are written to a http.Request
*/
type ListVendorsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list vendors params
func (o *ListVendorsParams) WithTimeout(timeout time.Duration) *ListVendorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list vendors params
func (o *ListVendorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list vendors params
func (o *ListVendorsParams) WithContext(ctx context.Context) *ListVendorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list vendors params
func (o *ListVendorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list vendors params
func (o *ListVendorsParams) WithHTTPClient(client *http.Client) *ListVendorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list vendors params
func (o *ListVendorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListVendorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}