// Code generated by go-swagger; DO NOT EDIT.

package add_ons

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

// NewPostAddonsParams creates a new PostAddonsParams object
// with the default values initialized.
func NewPostAddonsParams() *PostAddonsParams {
	var ()
	return &PostAddonsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAddonsParamsWithTimeout creates a new PostAddonsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAddonsParamsWithTimeout(timeout time.Duration) *PostAddonsParams {
	var ()
	return &PostAddonsParams{

		timeout: timeout,
	}
}

// NewPostAddonsParamsWithContext creates a new PostAddonsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAddonsParamsWithContext(ctx context.Context) *PostAddonsParams {
	var ()
	return &PostAddonsParams{

		Context: ctx,
	}
}

// NewPostAddonsParamsWithHTTPClient creates a new PostAddonsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAddonsParamsWithHTTPClient(client *http.Client) *PostAddonsParams {
	var ()
	return &PostAddonsParams{
		HTTPClient: client,
	}
}

/*PostAddonsParams contains all the parameters to send to the API endpoint
for the post addons operation typically these are written to a http.Request
*/
type PostAddonsParams struct {

	/*Addon
	  The add-on to be installed.

	*/
	Addon *models.PostAddonsParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post addons params
func (o *PostAddonsParams) WithTimeout(timeout time.Duration) *PostAddonsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post addons params
func (o *PostAddonsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post addons params
func (o *PostAddonsParams) WithContext(ctx context.Context) *PostAddonsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post addons params
func (o *PostAddonsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post addons params
func (o *PostAddonsParams) WithHTTPClient(client *http.Client) *PostAddonsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post addons params
func (o *PostAddonsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddon adds the addon to the post addons params
func (o *PostAddonsParams) WithAddon(addon *models.PostAddonsParamsBody) *PostAddonsParams {
	o.SetAddon(addon)
	return o
}

// SetAddon adds the addon to the post addons params
func (o *PostAddonsParams) SetAddon(addon *models.PostAddonsParamsBody) {
	o.Addon = addon
}

// WriteToRequest writes these params to a swagger request
func (o *PostAddonsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Addon != nil {
		if err := r.SetBodyParam(o.Addon); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
