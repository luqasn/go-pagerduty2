// Code generated by go-swagger; DO NOT EDIT.

package escalation_policies

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

// NewGetEscalationPolicyParams creates a new GetEscalationPolicyParams object
// with the default values initialized.
func NewGetEscalationPolicyParams() *GetEscalationPolicyParams {
	var ()
	return &GetEscalationPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEscalationPolicyParamsWithTimeout creates a new GetEscalationPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEscalationPolicyParamsWithTimeout(timeout time.Duration) *GetEscalationPolicyParams {
	var ()
	return &GetEscalationPolicyParams{

		timeout: timeout,
	}
}

// NewGetEscalationPolicyParamsWithContext creates a new GetEscalationPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEscalationPolicyParamsWithContext(ctx context.Context) *GetEscalationPolicyParams {
	var ()
	return &GetEscalationPolicyParams{

		Context: ctx,
	}
}

// NewGetEscalationPolicyParamsWithHTTPClient creates a new GetEscalationPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEscalationPolicyParamsWithHTTPClient(client *http.Client) *GetEscalationPolicyParams {
	var ()
	return &GetEscalationPolicyParams{
		HTTPClient: client,
	}
}

/*GetEscalationPolicyParams contains all the parameters to send to the API endpoint
for the get escalation policy operation typically these are written to a http.Request
*/
type GetEscalationPolicyParams struct {

	/*ID*/
	ID string
	/*Include
	  Array of additional details to include.

	*/
	Include []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get escalation policy params
func (o *GetEscalationPolicyParams) WithTimeout(timeout time.Duration) *GetEscalationPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get escalation policy params
func (o *GetEscalationPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get escalation policy params
func (o *GetEscalationPolicyParams) WithContext(ctx context.Context) *GetEscalationPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get escalation policy params
func (o *GetEscalationPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get escalation policy params
func (o *GetEscalationPolicyParams) WithHTTPClient(client *http.Client) *GetEscalationPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get escalation policy params
func (o *GetEscalationPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get escalation policy params
func (o *GetEscalationPolicyParams) WithID(id string) *GetEscalationPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get escalation policy params
func (o *GetEscalationPolicyParams) SetID(id string) {
	o.ID = id
}

// WithInclude adds the include to the get escalation policy params
func (o *GetEscalationPolicyParams) WithInclude(include []string) *GetEscalationPolicyParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get escalation policy params
func (o *GetEscalationPolicyParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *GetEscalationPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
