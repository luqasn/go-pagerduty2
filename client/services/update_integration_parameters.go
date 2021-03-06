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

// NewUpdateIntegrationParams creates a new UpdateIntegrationParams object
// with the default values initialized.
func NewUpdateIntegrationParams() *UpdateIntegrationParams {
	var ()
	return &UpdateIntegrationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIntegrationParamsWithTimeout creates a new UpdateIntegrationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateIntegrationParamsWithTimeout(timeout time.Duration) *UpdateIntegrationParams {
	var ()
	return &UpdateIntegrationParams{

		timeout: timeout,
	}
}

// NewUpdateIntegrationParamsWithContext creates a new UpdateIntegrationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateIntegrationParamsWithContext(ctx context.Context) *UpdateIntegrationParams {
	var ()
	return &UpdateIntegrationParams{

		Context: ctx,
	}
}

// NewUpdateIntegrationParamsWithHTTPClient creates a new UpdateIntegrationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateIntegrationParamsWithHTTPClient(client *http.Client) *UpdateIntegrationParams {
	var ()
	return &UpdateIntegrationParams{
		HTTPClient: client,
	}
}

/*UpdateIntegrationParams contains all the parameters to send to the API endpoint
for the update integration operation typically these are written to a http.Request
*/
type UpdateIntegrationParams struct {

	/*ID*/
	ID string
	/*Integration
	  The integration to be updated

	*/
	Integration *models.UpdateIntegrationParamsBody
	/*IntegrationID
	  The integration ID on the service.

	*/
	IntegrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update integration params
func (o *UpdateIntegrationParams) WithTimeout(timeout time.Duration) *UpdateIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update integration params
func (o *UpdateIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update integration params
func (o *UpdateIntegrationParams) WithContext(ctx context.Context) *UpdateIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update integration params
func (o *UpdateIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update integration params
func (o *UpdateIntegrationParams) WithHTTPClient(client *http.Client) *UpdateIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update integration params
func (o *UpdateIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update integration params
func (o *UpdateIntegrationParams) WithID(id string) *UpdateIntegrationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update integration params
func (o *UpdateIntegrationParams) SetID(id string) {
	o.ID = id
}

// WithIntegration adds the integration to the update integration params
func (o *UpdateIntegrationParams) WithIntegration(integration *models.UpdateIntegrationParamsBody) *UpdateIntegrationParams {
	o.SetIntegration(integration)
	return o
}

// SetIntegration adds the integration to the update integration params
func (o *UpdateIntegrationParams) SetIntegration(integration *models.UpdateIntegrationParamsBody) {
	o.Integration = integration
}

// WithIntegrationID adds the integrationID to the update integration params
func (o *UpdateIntegrationParams) WithIntegrationID(integrationID string) *UpdateIntegrationParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the update integration params
func (o *UpdateIntegrationParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Integration != nil {
		if err := r.SetBodyParam(o.Integration); err != nil {
			return err
		}
	}

	// path param integration_id
	if err := r.SetPathParam("integration_id", o.IntegrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
