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

// NewInstallAddonParams creates a new InstallAddonParams object
// with the default values initialized.
func NewInstallAddonParams() *InstallAddonParams {
	var ()
	return &InstallAddonParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInstallAddonParamsWithTimeout creates a new InstallAddonParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstallAddonParamsWithTimeout(timeout time.Duration) *InstallAddonParams {
	var ()
	return &InstallAddonParams{

		timeout: timeout,
	}
}

// NewInstallAddonParamsWithContext creates a new InstallAddonParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstallAddonParamsWithContext(ctx context.Context) *InstallAddonParams {
	var ()
	return &InstallAddonParams{

		Context: ctx,
	}
}

// NewInstallAddonParamsWithHTTPClient creates a new InstallAddonParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstallAddonParamsWithHTTPClient(client *http.Client) *InstallAddonParams {
	var ()
	return &InstallAddonParams{
		HTTPClient: client,
	}
}

/*InstallAddonParams contains all the parameters to send to the API endpoint
for the install addon operation typically these are written to a http.Request
*/
type InstallAddonParams struct {

	/*Addon
	  The add-on to be installed.

	*/
	Addon *models.InstallAddonParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the install addon params
func (o *InstallAddonParams) WithTimeout(timeout time.Duration) *InstallAddonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the install addon params
func (o *InstallAddonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the install addon params
func (o *InstallAddonParams) WithContext(ctx context.Context) *InstallAddonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the install addon params
func (o *InstallAddonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the install addon params
func (o *InstallAddonParams) WithHTTPClient(client *http.Client) *InstallAddonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the install addon params
func (o *InstallAddonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddon adds the addon to the install addon params
func (o *InstallAddonParams) WithAddon(addon *models.InstallAddonParamsBody) *InstallAddonParams {
	o.SetAddon(addon)
	return o
}

// SetAddon adds the addon to the install addon params
func (o *InstallAddonParams) SetAddon(addon *models.InstallAddonParamsBody) {
	o.Addon = addon
}

// WriteToRequest writes these params to a swagger request
func (o *InstallAddonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
