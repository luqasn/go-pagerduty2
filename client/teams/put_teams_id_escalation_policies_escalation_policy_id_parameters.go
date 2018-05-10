// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewPutTeamsIDEscalationPoliciesEscalationPolicyIDParams creates a new PutTeamsIDEscalationPoliciesEscalationPolicyIDParams object
// with the default values initialized.
func NewPutTeamsIDEscalationPoliciesEscalationPolicyIDParams() *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	var ()
	return &PutTeamsIDEscalationPoliciesEscalationPolicyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTeamsIDEscalationPoliciesEscalationPolicyIDParamsWithTimeout creates a new PutTeamsIDEscalationPoliciesEscalationPolicyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTeamsIDEscalationPoliciesEscalationPolicyIDParamsWithTimeout(timeout time.Duration) *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	var ()
	return &PutTeamsIDEscalationPoliciesEscalationPolicyIDParams{

		timeout: timeout,
	}
}

// NewPutTeamsIDEscalationPoliciesEscalationPolicyIDParamsWithContext creates a new PutTeamsIDEscalationPoliciesEscalationPolicyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTeamsIDEscalationPoliciesEscalationPolicyIDParamsWithContext(ctx context.Context) *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	var ()
	return &PutTeamsIDEscalationPoliciesEscalationPolicyIDParams{

		Context: ctx,
	}
}

// NewPutTeamsIDEscalationPoliciesEscalationPolicyIDParamsWithHTTPClient creates a new PutTeamsIDEscalationPoliciesEscalationPolicyIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTeamsIDEscalationPoliciesEscalationPolicyIDParamsWithHTTPClient(client *http.Client) *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	var ()
	return &PutTeamsIDEscalationPoliciesEscalationPolicyIDParams{
		HTTPClient: client,
	}
}

/*PutTeamsIDEscalationPoliciesEscalationPolicyIDParams contains all the parameters to send to the API endpoint
for the put teams ID escalation policies escalation policy ID operation typically these are written to a http.Request
*/
type PutTeamsIDEscalationPoliciesEscalationPolicyIDParams struct {

	/*EscalationPolicyID
	  The escalation policy ID on the team.

	*/
	EscalationPolicyID string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put teams ID escalation policies escalation policy ID params
func (o *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams) WithTimeout(timeout time.Duration) *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put teams ID escalation policies escalation policy ID params
func (o *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put teams ID escalation policies escalation policy ID params
func (o *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams) WithContext(ctx context.Context) *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put teams ID escalation policies escalation policy ID params
func (o *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put teams ID escalation policies escalation policy ID params
func (o *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams) WithHTTPClient(client *http.Client) *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put teams ID escalation policies escalation policy ID params
func (o *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEscalationPolicyID adds the escalationPolicyID to the put teams ID escalation policies escalation policy ID params
func (o *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams) WithEscalationPolicyID(escalationPolicyID string) *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	o.SetEscalationPolicyID(escalationPolicyID)
	return o
}

// SetEscalationPolicyID adds the escalationPolicyId to the put teams ID escalation policies escalation policy ID params
func (o *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams) SetEscalationPolicyID(escalationPolicyID string) {
	o.EscalationPolicyID = escalationPolicyID
}

// WithID adds the id to the put teams ID escalation policies escalation policy ID params
func (o *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams) WithID(id string) *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put teams ID escalation policies escalation policy ID params
func (o *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutTeamsIDEscalationPoliciesEscalationPolicyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param escalation_policy_id
	if err := r.SetPathParam("escalation_policy_id", o.EscalationPolicyID); err != nil {
		return err
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
