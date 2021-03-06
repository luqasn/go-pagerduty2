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

// NewUpdateEscalationPolicyParams creates a new UpdateEscalationPolicyParams object
// with the default values initialized.
func NewUpdateEscalationPolicyParams() *UpdateEscalationPolicyParams {
	var ()
	return &UpdateEscalationPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEscalationPolicyParamsWithTimeout creates a new UpdateEscalationPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateEscalationPolicyParamsWithTimeout(timeout time.Duration) *UpdateEscalationPolicyParams {
	var ()
	return &UpdateEscalationPolicyParams{

		timeout: timeout,
	}
}

// NewUpdateEscalationPolicyParamsWithContext creates a new UpdateEscalationPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateEscalationPolicyParamsWithContext(ctx context.Context) *UpdateEscalationPolicyParams {
	var ()
	return &UpdateEscalationPolicyParams{

		Context: ctx,
	}
}

// NewUpdateEscalationPolicyParamsWithHTTPClient creates a new UpdateEscalationPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateEscalationPolicyParamsWithHTTPClient(client *http.Client) *UpdateEscalationPolicyParams {
	var ()
	return &UpdateEscalationPolicyParams{
		HTTPClient: client,
	}
}

/*UpdateEscalationPolicyParams contains all the parameters to send to the API endpoint
for the update escalation policy operation typically these are written to a http.Request
*/
type UpdateEscalationPolicyParams struct {

	/*EscalationPolicy
	  The escalation policy to be updated.

	*/
	EscalationPolicy *models.UpdateEscalationPolicyParamsBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update escalation policy params
func (o *UpdateEscalationPolicyParams) WithTimeout(timeout time.Duration) *UpdateEscalationPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update escalation policy params
func (o *UpdateEscalationPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update escalation policy params
func (o *UpdateEscalationPolicyParams) WithContext(ctx context.Context) *UpdateEscalationPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update escalation policy params
func (o *UpdateEscalationPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update escalation policy params
func (o *UpdateEscalationPolicyParams) WithHTTPClient(client *http.Client) *UpdateEscalationPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update escalation policy params
func (o *UpdateEscalationPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEscalationPolicy adds the escalationPolicy to the update escalation policy params
func (o *UpdateEscalationPolicyParams) WithEscalationPolicy(escalationPolicy *models.UpdateEscalationPolicyParamsBody) *UpdateEscalationPolicyParams {
	o.SetEscalationPolicy(escalationPolicy)
	return o
}

// SetEscalationPolicy adds the escalationPolicy to the update escalation policy params
func (o *UpdateEscalationPolicyParams) SetEscalationPolicy(escalationPolicy *models.UpdateEscalationPolicyParamsBody) {
	o.EscalationPolicy = escalationPolicy
}

// WithID adds the id to the update escalation policy params
func (o *UpdateEscalationPolicyParams) WithID(id string) *UpdateEscalationPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update escalation policy params
func (o *UpdateEscalationPolicyParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEscalationPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
