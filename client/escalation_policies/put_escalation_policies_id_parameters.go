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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// NewPutEscalationPoliciesIDParams creates a new PutEscalationPoliciesIDParams object
// with the default values initialized.
func NewPutEscalationPoliciesIDParams() *PutEscalationPoliciesIDParams {
	var ()
	return &PutEscalationPoliciesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutEscalationPoliciesIDParamsWithTimeout creates a new PutEscalationPoliciesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutEscalationPoliciesIDParamsWithTimeout(timeout time.Duration) *PutEscalationPoliciesIDParams {
	var ()
	return &PutEscalationPoliciesIDParams{

		timeout: timeout,
	}
}

// NewPutEscalationPoliciesIDParamsWithContext creates a new PutEscalationPoliciesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutEscalationPoliciesIDParamsWithContext(ctx context.Context) *PutEscalationPoliciesIDParams {
	var ()
	return &PutEscalationPoliciesIDParams{

		Context: ctx,
	}
}

// NewPutEscalationPoliciesIDParamsWithHTTPClient creates a new PutEscalationPoliciesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutEscalationPoliciesIDParamsWithHTTPClient(client *http.Client) *PutEscalationPoliciesIDParams {
	var ()
	return &PutEscalationPoliciesIDParams{
		HTTPClient: client,
	}
}

/*PutEscalationPoliciesIDParams contains all the parameters to send to the API endpoint
for the put escalation policies ID operation typically these are written to a http.Request
*/
type PutEscalationPoliciesIDParams struct {

	/*EscalationPolicy
	  The escalation policy to be updated.

	*/
	EscalationPolicy *models.PutEscalationPoliciesIDParamsBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put escalation policies ID params
func (o *PutEscalationPoliciesIDParams) WithTimeout(timeout time.Duration) *PutEscalationPoliciesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put escalation policies ID params
func (o *PutEscalationPoliciesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put escalation policies ID params
func (o *PutEscalationPoliciesIDParams) WithContext(ctx context.Context) *PutEscalationPoliciesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put escalation policies ID params
func (o *PutEscalationPoliciesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put escalation policies ID params
func (o *PutEscalationPoliciesIDParams) WithHTTPClient(client *http.Client) *PutEscalationPoliciesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put escalation policies ID params
func (o *PutEscalationPoliciesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEscalationPolicy adds the escalationPolicy to the put escalation policies ID params
func (o *PutEscalationPoliciesIDParams) WithEscalationPolicy(escalationPolicy *models.PutEscalationPoliciesIDParamsBody) *PutEscalationPoliciesIDParams {
	o.SetEscalationPolicy(escalationPolicy)
	return o
}

// SetEscalationPolicy adds the escalationPolicy to the put escalation policies ID params
func (o *PutEscalationPoliciesIDParams) SetEscalationPolicy(escalationPolicy *models.PutEscalationPoliciesIDParamsBody) {
	o.EscalationPolicy = escalationPolicy
}

// WithID adds the id to the put escalation policies ID params
func (o *PutEscalationPoliciesIDParams) WithID(id string) *PutEscalationPoliciesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put escalation policies ID params
func (o *PutEscalationPoliciesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutEscalationPoliciesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EscalationPolicy != nil {
		if err := r.SetBodyParam(o.EscalationPolicy); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
