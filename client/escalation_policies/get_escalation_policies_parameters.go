// Code generated by go-swagger; DO NOT EDIT.

package escalation_policies

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

// NewGetEscalationPoliciesParams creates a new GetEscalationPoliciesParams object
// with the default values initialized.
func NewGetEscalationPoliciesParams() *GetEscalationPoliciesParams {
	var (
		sortByDefault = string("name")
	)
	return &GetEscalationPoliciesParams{
		SortBy: &sortByDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEscalationPoliciesParamsWithTimeout creates a new GetEscalationPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEscalationPoliciesParamsWithTimeout(timeout time.Duration) *GetEscalationPoliciesParams {
	var (
		sortByDefault = string("name")
	)
	return &GetEscalationPoliciesParams{
		SortBy: &sortByDefault,

		timeout: timeout,
	}
}

// NewGetEscalationPoliciesParamsWithContext creates a new GetEscalationPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEscalationPoliciesParamsWithContext(ctx context.Context) *GetEscalationPoliciesParams {
	var (
		sortByDefault = string("name")
	)
	return &GetEscalationPoliciesParams{
		SortBy: &sortByDefault,

		Context: ctx,
	}
}

// NewGetEscalationPoliciesParamsWithHTTPClient creates a new GetEscalationPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEscalationPoliciesParamsWithHTTPClient(client *http.Client) *GetEscalationPoliciesParams {
	var (
		sortByDefault = string("name")
	)
	return &GetEscalationPoliciesParams{
		SortBy:     &sortByDefault,
		HTTPClient: client,
	}
}

/*GetEscalationPoliciesParams contains all the parameters to send to the API endpoint
for the get escalation policies operation typically these are written to a http.Request
*/
type GetEscalationPoliciesParams struct {

	/*Include
	  Array of additional details to include.

	*/
	Include []string
	/*Query
	  Filters the results, showing only the escalation policies whose names contain the query.

	*/
	Query *string
	/*SortBy
	  Used to specify the field you wish to sort the results on.

	*/
	SortBy *string
	/*TeamIds
	  An array of team IDs. Only results related to these teams will be returned. Account must have the `teams` ability to use this parameter.

	*/
	TeamIds []string
	/*UserIds
	  Filters the results, showing only escalation policies on which any of the users is a target.

	*/
	UserIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get escalation policies params
func (o *GetEscalationPoliciesParams) WithTimeout(timeout time.Duration) *GetEscalationPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get escalation policies params
func (o *GetEscalationPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get escalation policies params
func (o *GetEscalationPoliciesParams) WithContext(ctx context.Context) *GetEscalationPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get escalation policies params
func (o *GetEscalationPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get escalation policies params
func (o *GetEscalationPoliciesParams) WithHTTPClient(client *http.Client) *GetEscalationPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get escalation policies params
func (o *GetEscalationPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInclude adds the include to the get escalation policies params
func (o *GetEscalationPoliciesParams) WithInclude(include []string) *GetEscalationPoliciesParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get escalation policies params
func (o *GetEscalationPoliciesParams) SetInclude(include []string) {
	o.Include = include
}

// WithQuery adds the query to the get escalation policies params
func (o *GetEscalationPoliciesParams) WithQuery(query *string) *GetEscalationPoliciesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get escalation policies params
func (o *GetEscalationPoliciesParams) SetQuery(query *string) {
	o.Query = query
}

// WithSortBy adds the sortBy to the get escalation policies params
func (o *GetEscalationPoliciesParams) WithSortBy(sortBy *string) *GetEscalationPoliciesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get escalation policies params
func (o *GetEscalationPoliciesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithTeamIds adds the teamIds to the get escalation policies params
func (o *GetEscalationPoliciesParams) WithTeamIds(teamIds []string) *GetEscalationPoliciesParams {
	o.SetTeamIds(teamIds)
	return o
}

// SetTeamIds adds the teamIds to the get escalation policies params
func (o *GetEscalationPoliciesParams) SetTeamIds(teamIds []string) {
	o.TeamIds = teamIds
}

// WithUserIds adds the userIds to the get escalation policies params
func (o *GetEscalationPoliciesParams) WithUserIds(userIds []string) *GetEscalationPoliciesParams {
	o.SetUserIds(userIds)
	return o
}

// SetUserIds adds the userIds to the get escalation policies params
func (o *GetEscalationPoliciesParams) SetUserIds(userIds []string) {
	o.UserIds = userIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetEscalationPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SortBy != nil {

		// query param sort_by
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sort_by", qSortBy); err != nil {
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

	valuesUserIds := o.UserIds

	joinedUserIds := swag.JoinByFormat(valuesUserIds, "multi")
	// query array param user_ids[]
	if err := r.SetQueryParam("user_ids[]", joinedUserIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}