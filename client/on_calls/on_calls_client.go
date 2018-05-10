// Code generated by go-swagger; DO NOT EDIT.

package on_calls

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new on calls API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for on calls API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetOncalls lists all of the on calls

List the on-call entries during a given time range.
*/
func (a *Client) GetOncalls(params *GetOncallsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOncallsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOncallsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOncalls",
		Method:             "GET",
		PathPattern:        "/oncalls",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOncallsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOncallsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
