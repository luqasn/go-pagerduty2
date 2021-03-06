// Code generated by go-swagger; DO NOT EDIT.

package escalation_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new escalation policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for escalation policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateEscalationPolicy creates an escalation policy

Creates a new escalation policy. There must be at least one existing escalation rule added to create a new escalation policy.
*/
func (a *Client) CreateEscalationPolicy(params *CreateEscalationPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateEscalationPolicyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEscalationPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createEscalationPolicy",
		Method:             "POST",
		PathPattern:        "/escalation_policies",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateEscalationPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateEscalationPolicyCreated), nil

}

/*
DeleteEscalationPolicy deletes an escalation policy

Deletes an existing escalation policy and rules. The escalation policy must not be in use by any services.
*/
func (a *Client) DeleteEscalationPolicy(params *DeleteEscalationPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteEscalationPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEscalationPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEscalationPolicy",
		Method:             "DELETE",
		PathPattern:        "/escalation_policies/{id}",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEscalationPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEscalationPolicyNoContent), nil

}

/*
GetEscalationPolicy gets an escalation policy

Get information about an existing escalation policy and its rules.
*/
func (a *Client) GetEscalationPolicy(params *GetEscalationPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*GetEscalationPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEscalationPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEscalationPolicy",
		Method:             "GET",
		PathPattern:        "/escalation_policies/{id}",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEscalationPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEscalationPolicyOK), nil

}

/*
ListEscalationPolicies lists escalation policies

List all of the existing escalation policies.
*/
func (a *Client) ListEscalationPolicies(params *ListEscalationPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*ListEscalationPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListEscalationPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listEscalationPolicies",
		Method:             "GET",
		PathPattern:        "/escalation_policies",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListEscalationPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListEscalationPoliciesOK), nil

}

/*
UpdateEscalationPolicy updates an escalation policy

Updates an existing escalation policy and rules.
*/
func (a *Client) UpdateEscalationPolicy(params *UpdateEscalationPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateEscalationPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEscalationPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateEscalationPolicy",
		Method:             "PUT",
		PathPattern:        "/escalation_policies/{id}",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateEscalationPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateEscalationPolicyOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
