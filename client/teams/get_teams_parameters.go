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

// NewGetTeamsParams creates a new GetTeamsParams object
// with the default values initialized.
func NewGetTeamsParams() *GetTeamsParams {
	var ()
	return &GetTeamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeamsParamsWithTimeout creates a new GetTeamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTeamsParamsWithTimeout(timeout time.Duration) *GetTeamsParams {
	var ()
	return &GetTeamsParams{

		timeout: timeout,
	}
}

// NewGetTeamsParamsWithContext creates a new GetTeamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTeamsParamsWithContext(ctx context.Context) *GetTeamsParams {
	var ()
	return &GetTeamsParams{

		Context: ctx,
	}
}

// NewGetTeamsParamsWithHTTPClient creates a new GetTeamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTeamsParamsWithHTTPClient(client *http.Client) *GetTeamsParams {
	var ()
	return &GetTeamsParams{
		HTTPClient: client,
	}
}

/*GetTeamsParams contains all the parameters to send to the API endpoint
for the get teams operation typically these are written to a http.Request
*/
type GetTeamsParams struct {

	/*Query
	  Filters the result, showing only the teams whose names or email addresses match the query.

	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get teams params
func (o *GetTeamsParams) WithTimeout(timeout time.Duration) *GetTeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get teams params
func (o *GetTeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get teams params
func (o *GetTeamsParams) WithContext(ctx context.Context) *GetTeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get teams params
func (o *GetTeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get teams params
func (o *GetTeamsParams) WithHTTPClient(client *http.Client) *GetTeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get teams params
func (o *GetTeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the get teams params
func (o *GetTeamsParams) WithQuery(query *string) *GetTeamsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get teams params
func (o *GetTeamsParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
