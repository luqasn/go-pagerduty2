// Code generated by go-swagger; DO NOT EDIT.

package log_entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new log entries API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for log entries API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetLogEntry gets a log entry

Get details for a specific incident log entry. This method provides additional information you can use to get at raw event data.
*/
func (a *Client) GetLogEntry(params *GetLogEntryParams, authInfo runtime.ClientAuthInfoWriter) (*GetLogEntryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLogEntryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLogEntry",
		Method:             "GET",
		PathPattern:        "/log_entries/{id}",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLogEntryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLogEntryOK), nil

}

/*
ListAllLogEntries lists log entries

List all of the incident log entries across the entire account.
*/
func (a *Client) ListAllLogEntries(params *ListAllLogEntriesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllLogEntriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllLogEntriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllLogEntries",
		Method:             "GET",
		PathPattern:        "/log_entries",
		ProducesMediaTypes: []string{"application/vnd.pagerduty+json;version=2"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAllLogEntriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAllLogEntriesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
