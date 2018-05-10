// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUsersParams creates a new GetUsersParams object
// with the default values initialized.
func NewGetUsersParams() *GetUsersParams {
	var ()
	return &GetUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersParamsWithTimeout creates a new GetUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersParamsWithTimeout(timeout time.Duration) *GetUsersParams {
	var ()
	return &GetUsersParams{

		timeout: timeout,
	}
}

// NewGetUsersParamsWithContext creates a new GetUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersParamsWithContext(ctx context.Context) *GetUsersParams {
	var ()
	return &GetUsersParams{

		Context: ctx,
	}
}

// NewGetUsersParamsWithHTTPClient creates a new GetUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersParamsWithHTTPClient(client *http.Client) *GetUsersParams {
	var ()
	return &GetUsersParams{
		HTTPClient: client,
	}
}

/*GetUsersParams contains all the parameters to send to the API endpoint
for the get users operation typically these are written to a http.Request
*/
type GetUsersParams struct {

	/*Include
	  Array of additional details to include.

	*/
	Include []string
	/*Query
	  Filters the result, showing only the users whose names or email addresses match the query.

	*/
	Query *string
	/*TeamIds
	  An array of team IDs. Only results related to these teams will be returned. Account must have the `teams` ability to use this parameter.

	*/
	TeamIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users params
func (o *GetUsersParams) WithTimeout(timeout time.Duration) *GetUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users params
func (o *GetUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users params
func (o *GetUsersParams) WithContext(ctx context.Context) *GetUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users params
func (o *GetUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) WithHTTPClient(client *http.Client) *GetUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInclude adds the include to the get users params
func (o *GetUsersParams) WithInclude(include []string) *GetUsersParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get users params
func (o *GetUsersParams) SetInclude(include []string) {
	o.Include = include
}

// WithQuery adds the query to the get users params
func (o *GetUsersParams) WithQuery(query *string) *GetUsersParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get users params
func (o *GetUsersParams) SetQuery(query *string) {
	o.Query = query
}

// WithTeamIds adds the teamIds to the get users params
func (o *GetUsersParams) WithTeamIds(teamIds []string) *GetUsersParams {
	o.SetTeamIds(teamIds)
	return o
}

// SetTeamIds adds the teamIds to the get users params
func (o *GetUsersParams) SetTeamIds(teamIds []string) {
	o.TeamIds = teamIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesInclude := o.Include

	joinedInclude := swag.JoinByFormat(valuesInclude, "multi")
	// query array param include[]
	if err := r.SetQueryParam("include[]", joinedInclude...); err != nil {
		return err
	}

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

	valuesTeamIds := o.TeamIds

	joinedTeamIds := swag.JoinByFormat(valuesTeamIds, "multi")
	// query array param team_ids[]
	if err := r.SetQueryParam("team_ids[]", joinedTeamIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
