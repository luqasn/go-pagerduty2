// Code generated by go-swagger; DO NOT EDIT.

package incidents

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

// NewListIncidentsParams creates a new ListIncidentsParams object
// with the default values initialized.
func NewListIncidentsParams() *ListIncidentsParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &ListIncidentsParams{
		TimeZone: &timeZoneDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListIncidentsParamsWithTimeout creates a new ListIncidentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListIncidentsParamsWithTimeout(timeout time.Duration) *ListIncidentsParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &ListIncidentsParams{
		TimeZone: &timeZoneDefault,

		timeout: timeout,
	}
}

// NewListIncidentsParamsWithContext creates a new ListIncidentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListIncidentsParamsWithContext(ctx context.Context) *ListIncidentsParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &ListIncidentsParams{
		TimeZone: &timeZoneDefault,

		Context: ctx,
	}
}

// NewListIncidentsParamsWithHTTPClient creates a new ListIncidentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListIncidentsParamsWithHTTPClient(client *http.Client) *ListIncidentsParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &ListIncidentsParams{
		TimeZone:   &timeZoneDefault,
		HTTPClient: client,
	}
}

/*ListIncidentsParams contains all the parameters to send to the API endpoint
for the list incidents operation typically these are written to a http.Request
*/
type ListIncidentsParams struct {

	/*DateRange
	  When set to all, the since and until parameters and defaults are ignored.

	*/
	DateRange *string
	/*IncidentKey
	  Incident de-duplication key. Incidents with child alerts do not have an incident key; querying by incident key will return incidents whose alerts have alert_key matching the given incident key.

	*/
	IncidentKey *string
	/*Include
	  Array of additional details to include.

	*/
	Include []string
	/*ServiceIds
	  Returns only the incidents associated with the passed service(s). This expects one or more service IDs.

	*/
	ServiceIds []string
	/*Since
	  The start of the date range over which you want to search.

	*/
	Since *strfmt.DateTime
	/*SortBy
	  Used to specify both the field you wish to sort the results on (incident_number/created_at/resolved_at/urgency), as well as the direction (asc/desc) of the results. The sort_by field and direction should be separated by a colon. A maximum of two fields can be included, separated by a comma. Sort direction defaults to ascending. The account must have the `urgencies` ability to sort by the urgency.

	*/
	SortBy []string
	/*Statuses
	  Return only incidents with the given statuses. (More status codes may be introduced in the future.)

	*/
	Statuses []string
	/*TeamIds
	  An array of team IDs. Only results related to these teams will be returned. Account must have the `teams` ability to use this parameter.

	*/
	TeamIds []string
	/*TimeZone
	  Time zone in which dates in the result will be rendered.

	*/
	TimeZone *string
	/*Until
	  The end of the date range over which you want to search.

	*/
	Until *strfmt.DateTime
	/*Urgencies
	  Array of the urgencies of the incidents to be returned. Defaults to all urgencies. Account must have the `urgencies` ability to do this.

	*/
	Urgencies []string
	/*UserIds
	  Returns only the incidents currently assigned to the passed user(s). This expects one or more user IDs. Note: When using the assigned_to_user filter, you will only receive incidents with statuses of triggered or acknowledged. This is because resolved incidents are not assigned to any user.

	*/
	UserIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list incidents params
func (o *ListIncidentsParams) WithTimeout(timeout time.Duration) *ListIncidentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list incidents params
func (o *ListIncidentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list incidents params
func (o *ListIncidentsParams) WithContext(ctx context.Context) *ListIncidentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list incidents params
func (o *ListIncidentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list incidents params
func (o *ListIncidentsParams) WithHTTPClient(client *http.Client) *ListIncidentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list incidents params
func (o *ListIncidentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateRange adds the dateRange to the list incidents params
func (o *ListIncidentsParams) WithDateRange(dateRange *string) *ListIncidentsParams {
	o.SetDateRange(dateRange)
	return o
}

// SetDateRange adds the dateRange to the list incidents params
func (o *ListIncidentsParams) SetDateRange(dateRange *string) {
	o.DateRange = dateRange
}

// WithIncidentKey adds the incidentKey to the list incidents params
func (o *ListIncidentsParams) WithIncidentKey(incidentKey *string) *ListIncidentsParams {
	o.SetIncidentKey(incidentKey)
	return o
}

// SetIncidentKey adds the incidentKey to the list incidents params
func (o *ListIncidentsParams) SetIncidentKey(incidentKey *string) {
	o.IncidentKey = incidentKey
}

// WithInclude adds the include to the list incidents params
func (o *ListIncidentsParams) WithInclude(include []string) *ListIncidentsParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the list incidents params
func (o *ListIncidentsParams) SetInclude(include []string) {
	o.Include = include
}

// WithServiceIds adds the serviceIds to the list incidents params
func (o *ListIncidentsParams) WithServiceIds(serviceIds []string) *ListIncidentsParams {
	o.SetServiceIds(serviceIds)
	return o
}

// SetServiceIds adds the serviceIds to the list incidents params
func (o *ListIncidentsParams) SetServiceIds(serviceIds []string) {
	o.ServiceIds = serviceIds
}

// WithSince adds the since to the list incidents params
func (o *ListIncidentsParams) WithSince(since *strfmt.DateTime) *ListIncidentsParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the list incidents params
func (o *ListIncidentsParams) SetSince(since *strfmt.DateTime) {
	o.Since = since
}

// WithSortBy adds the sortBy to the list incidents params
func (o *ListIncidentsParams) WithSortBy(sortBy []string) *ListIncidentsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the list incidents params
func (o *ListIncidentsParams) SetSortBy(sortBy []string) {
	o.SortBy = sortBy
}

// WithStatuses adds the statuses to the list incidents params
func (o *ListIncidentsParams) WithStatuses(statuses []string) *ListIncidentsParams {
	o.SetStatuses(statuses)
	return o
}

// SetStatuses adds the statuses to the list incidents params
func (o *ListIncidentsParams) SetStatuses(statuses []string) {
	o.Statuses = statuses
}

// WithTeamIds adds the teamIds to the list incidents params
func (o *ListIncidentsParams) WithTeamIds(teamIds []string) *ListIncidentsParams {
	o.SetTeamIds(teamIds)
	return o
}

// SetTeamIds adds the teamIds to the list incidents params
func (o *ListIncidentsParams) SetTeamIds(teamIds []string) {
	o.TeamIds = teamIds
}

// WithTimeZone adds the timeZone to the list incidents params
func (o *ListIncidentsParams) WithTimeZone(timeZone *string) *ListIncidentsParams {
	o.SetTimeZone(timeZone)
	return o
}

// SetTimeZone adds the timeZone to the list incidents params
func (o *ListIncidentsParams) SetTimeZone(timeZone *string) {
	o.TimeZone = timeZone
}

// WithUntil adds the until to the list incidents params
func (o *ListIncidentsParams) WithUntil(until *strfmt.DateTime) *ListIncidentsParams {
	o.SetUntil(until)
	return o
}

// SetUntil adds the until to the list incidents params
func (o *ListIncidentsParams) SetUntil(until *strfmt.DateTime) {
	o.Until = until
}

// WithUrgencies adds the urgencies to the list incidents params
func (o *ListIncidentsParams) WithUrgencies(urgencies []string) *ListIncidentsParams {
	o.SetUrgencies(urgencies)
	return o
}

// SetUrgencies adds the urgencies to the list incidents params
func (o *ListIncidentsParams) SetUrgencies(urgencies []string) {
	o.Urgencies = urgencies
}

// WithUserIds adds the userIds to the list incidents params
func (o *ListIncidentsParams) WithUserIds(userIds []string) *ListIncidentsParams {
	o.SetUserIds(userIds)
	return o
}

// SetUserIds adds the userIds to the list incidents params
func (o *ListIncidentsParams) SetUserIds(userIds []string) {
	o.UserIds = userIds
}

// WriteToRequest writes these params to a swagger request
func (o *ListIncidentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DateRange != nil {

		// query param date_range
		var qrDateRange string
		if o.DateRange != nil {
			qrDateRange = *o.DateRange
		}
		qDateRange := qrDateRange
		if qDateRange != "" {
			if err := r.SetQueryParam("date_range", qDateRange); err != nil {
				return err
			}
		}

	}

	if o.IncidentKey != nil {

		// query param incident_key
		var qrIncidentKey string
		if o.IncidentKey != nil {
			qrIncidentKey = *o.IncidentKey
		}
		qIncidentKey := qrIncidentKey
		if qIncidentKey != "" {
			if err := r.SetQueryParam("incident_key", qIncidentKey); err != nil {
				return err
			}
		}

	}

	valuesInclude := o.Include

	joinedInclude := swag.JoinByFormat(valuesInclude, "multi")
	// query array param include[]
	if err := r.SetQueryParam("include[]", joinedInclude...); err != nil {
		return err
	}

	valuesServiceIds := o.ServiceIds

	joinedServiceIds := swag.JoinByFormat(valuesServiceIds, "multi")
	// query array param service_ids[]
	if err := r.SetQueryParam("service_ids[]", joinedServiceIds...); err != nil {
		return err
	}

	if o.Since != nil {

		// query param since
		var qrSince strfmt.DateTime
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince.String()
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	valuesSortBy := o.SortBy

	joinedSortBy := swag.JoinByFormat(valuesSortBy, "csv")
	// query array param sort_by
	if err := r.SetQueryParam("sort_by", joinedSortBy...); err != nil {
		return err
	}

	valuesStatuses := o.Statuses

	joinedStatuses := swag.JoinByFormat(valuesStatuses, "multi")
	// query array param statuses[]
	if err := r.SetQueryParam("statuses[]", joinedStatuses...); err != nil {
		return err
	}

	valuesTeamIds := o.TeamIds

	joinedTeamIds := swag.JoinByFormat(valuesTeamIds, "multi")
	// query array param team_ids[]
	if err := r.SetQueryParam("team_ids[]", joinedTeamIds...); err != nil {
		return err
	}

	if o.TimeZone != nil {

		// query param time_zone
		var qrTimeZone string
		if o.TimeZone != nil {
			qrTimeZone = *o.TimeZone
		}
		qTimeZone := qrTimeZone
		if qTimeZone != "" {
			if err := r.SetQueryParam("time_zone", qTimeZone); err != nil {
				return err
			}
		}

	}

	if o.Until != nil {

		// query param until
		var qrUntil strfmt.DateTime
		if o.Until != nil {
			qrUntil = *o.Until
		}
		qUntil := qrUntil.String()
		if qUntil != "" {
			if err := r.SetQueryParam("until", qUntil); err != nil {
				return err
			}
		}

	}

	valuesUrgencies := o.Urgencies

	joinedUrgencies := swag.JoinByFormat(valuesUrgencies, "multi")
	// query array param urgencies[]
	if err := r.SetQueryParam("urgencies[]", joinedUrgencies...); err != nil {
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