// Code generated by go-swagger; DO NOT EDIT.

package priorities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new priorities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for priorities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPriorities lists priorities

List existing priorities, in order (most to least severe).
*/
func (a *Client) GetPriorities(params *GetPrioritiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPrioritiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrioritiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPriorities",
		Method:             "GET",
		PathPattern:        "/priorities",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPrioritiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrioritiesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}