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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetServicesIDIntegrationsIntegrationIDParams creates a new GetServicesIDIntegrationsIntegrationIDParams object
// with the default values initialized.
func NewGetServicesIDIntegrationsIntegrationIDParams() *GetServicesIDIntegrationsIntegrationIDParams {
	var ()
	return &GetServicesIDIntegrationsIntegrationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServicesIDIntegrationsIntegrationIDParamsWithTimeout creates a new GetServicesIDIntegrationsIntegrationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServicesIDIntegrationsIntegrationIDParamsWithTimeout(timeout time.Duration) *GetServicesIDIntegrationsIntegrationIDParams {
	var ()
	return &GetServicesIDIntegrationsIntegrationIDParams{

		timeout: timeout,
	}
}

// NewGetServicesIDIntegrationsIntegrationIDParamsWithContext creates a new GetServicesIDIntegrationsIntegrationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServicesIDIntegrationsIntegrationIDParamsWithContext(ctx context.Context) *GetServicesIDIntegrationsIntegrationIDParams {
	var ()
	return &GetServicesIDIntegrationsIntegrationIDParams{

		Context: ctx,
	}
}

// NewGetServicesIDIntegrationsIntegrationIDParamsWithHTTPClient creates a new GetServicesIDIntegrationsIntegrationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServicesIDIntegrationsIntegrationIDParamsWithHTTPClient(client *http.Client) *GetServicesIDIntegrationsIntegrationIDParams {
	var ()
	return &GetServicesIDIntegrationsIntegrationIDParams{
		HTTPClient: client,
	}
}

/*GetServicesIDIntegrationsIntegrationIDParams contains all the parameters to send to the API endpoint
for the get services ID integrations integration ID operation typically these are written to a http.Request
*/
type GetServicesIDIntegrationsIntegrationIDParams struct {

	/*ID*/
	ID string
	/*Include
	  Array of additional details to include.

	*/
	Include []string
	/*IntegrationID
	  The integration ID on the service.

	*/
	IntegrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get services ID integrations integration ID params
func (o *GetServicesIDIntegrationsIntegrationIDParams) WithTimeout(timeout time.Duration) *GetServicesIDIntegrationsIntegrationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get services ID integrations integration ID params
func (o *GetServicesIDIntegrationsIntegrationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get services ID integrations integration ID params
func (o *GetServicesIDIntegrationsIntegrationIDParams) WithContext(ctx context.Context) *GetServicesIDIntegrationsIntegrationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get services ID integrations integration ID params
func (o *GetServicesIDIntegrationsIntegrationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get services ID integrations integration ID params
func (o *GetServicesIDIntegrationsIntegrationIDParams) WithHTTPClient(client *http.Client) *GetServicesIDIntegrationsIntegrationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get services ID integrations integration ID params
func (o *GetServicesIDIntegrationsIntegrationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get services ID integrations integration ID params
func (o *GetServicesIDIntegrationsIntegrationIDParams) WithID(id string) *GetServicesIDIntegrationsIntegrationIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get services ID integrations integration ID params
func (o *GetServicesIDIntegrationsIntegrationIDParams) SetID(id string) {
	o.ID = id
}

// WithInclude adds the include to the get services ID integrations integration ID params
func (o *GetServicesIDIntegrationsIntegrationIDParams) WithInclude(include []string) *GetServicesIDIntegrationsIntegrationIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get services ID integrations integration ID params
func (o *GetServicesIDIntegrationsIntegrationIDParams) SetInclude(include []string) {
	o.Include = include
}

// WithIntegrationID adds the integrationID to the get services ID integrations integration ID params
func (o *GetServicesIDIntegrationsIntegrationIDParams) WithIntegrationID(integrationID string) *GetServicesIDIntegrationsIntegrationIDParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the get services ID integrations integration ID params
func (o *GetServicesIDIntegrationsIntegrationIDParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServicesIDIntegrationsIntegrationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	valuesInclude := o.Include

	joinedInclude := swag.JoinByFormat(valuesInclude, "multi")
	// query array param include[]
	if err := r.SetQueryParam("include[]", joinedInclude...); err != nil {
		return err
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
