// Code generated by go-swagger; DO NOT EDIT.

package abilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new abilities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for abilities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ListAbilities lists abilities

List all of your account's abilities, by name.
*/
func (a *Client) ListAbilities(params *ListAbilitiesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAbilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAbilitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAbilities",
		Method:             "GET",
		PathPattern:        "/abilities",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAbilitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAbilitiesOK), nil

}

/*
TestAbility tests an ability

Test whether your account has a given ability.
*/
func (a *Client) TestAbility(params *TestAbilityParams, authInfo runtime.ClientAuthInfoWriter) (*TestAbilityNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestAbilityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "testAbility",
		Method:             "GET",
		PathPattern:        "/abilities/{id}",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestAbilityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TestAbilityNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
