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
GetAbilities lists abilities

List all of your account's abilities, by name.
*/
func (a *Client) GetAbilities(params *GetAbilitiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetAbilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAbilitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAbilities",
		Method:             "GET",
		PathPattern:        "/abilities",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAbilitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAbilitiesOK), nil

}

/*
GetAbilitiesID tests an ability

Test whether your account has a given ability.
*/
func (a *Client) GetAbilitiesID(params *GetAbilitiesIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetAbilitiesIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAbilitiesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAbilitiesID",
		Method:             "GET",
		PathPattern:        "/abilities/{id}",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAbilitiesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAbilitiesIDNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
