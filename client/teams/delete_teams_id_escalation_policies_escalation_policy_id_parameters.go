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

// NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams creates a new DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams object
// with the default values initialized.
func NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams() *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	var ()
	return &DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDParamsWithTimeout creates a new DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDParamsWithTimeout(timeout time.Duration) *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	var ()
	return &DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams{

		timeout: timeout,
	}
}

// NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDParamsWithContext creates a new DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDParamsWithContext(ctx context.Context) *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	var ()
	return &DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams{

		Context: ctx,
	}
}

// NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDParamsWithHTTPClient creates a new DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTeamsIDEscalationPoliciesEscalationPolicyIDParamsWithHTTPClient(client *http.Client) *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	var ()
	return &DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams{
		HTTPClient: client,
	}
}

/*DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams contains all the parameters to send to the API endpoint
for the delete teams ID escalation policies escalation policy ID operation typically these are written to a http.Request
*/
type DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams struct {

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

// WithTimeout adds the timeout to the delete teams ID escalation policies escalation policy ID params
func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams) WithTimeout(timeout time.Duration) *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete teams ID escalation policies escalation policy ID params
func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete teams ID escalation policies escalation policy ID params
func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams) WithContext(ctx context.Context) *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete teams ID escalation policies escalation policy ID params
func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete teams ID escalation policies escalation policy ID params
func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams) WithHTTPClient(client *http.Client) *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete teams ID escalation policies escalation policy ID params
func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEscalationPolicyID adds the escalationPolicyID to the delete teams ID escalation policies escalation policy ID params
func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams) WithEscalationPolicyID(escalationPolicyID string) *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	o.SetEscalationPolicyID(escalationPolicyID)
	return o
}

// SetEscalationPolicyID adds the escalationPolicyId to the delete teams ID escalation policies escalation policy ID params
func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams) SetEscalationPolicyID(escalationPolicyID string) {
	o.EscalationPolicyID = escalationPolicyID
}

// WithID adds the id to the delete teams ID escalation policies escalation policy ID params
func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams) WithID(id string) *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete teams ID escalation policies escalation policy ID params
func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTeamsIDEscalationPoliciesEscalationPolicyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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