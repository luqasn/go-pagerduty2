// Code generated by go-swagger; DO NOT EDIT.

package on_calls

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

// NewListOnCallsParams creates a new ListOnCallsParams object
// with the default values initialized.
func NewListOnCallsParams() *ListOnCallsParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &ListOnCallsParams{
		TimeZone: &timeZoneDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListOnCallsParamsWithTimeout creates a new ListOnCallsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOnCallsParamsWithTimeout(timeout time.Duration) *ListOnCallsParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &ListOnCallsParams{
		TimeZone: &timeZoneDefault,

		timeout: timeout,
	}
}

// NewListOnCallsParamsWithContext creates a new ListOnCallsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOnCallsParamsWithContext(ctx context.Context) *ListOnCallsParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &ListOnCallsParams{
		TimeZone: &timeZoneDefault,

		Context: ctx,
	}
}

// NewListOnCallsParamsWithHTTPClient creates a new ListOnCallsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOnCallsParamsWithHTTPClient(client *http.Client) *ListOnCallsParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &ListOnCallsParams{
		TimeZone:   &timeZoneDefault,
		HTTPClient: client,
	}
}

/*ListOnCallsParams contains all the parameters to send to the API endpoint
for the list on calls operation typically these are written to a http.Request
*/
type ListOnCallsParams struct {

	/*Earliest
	  This will filter on-calls such that only the earliest on-call for each combination of escalation policy, escalation level, and user is returned. This is useful for determining when the "next" on-calls are for a given set of filters.

	*/
	Earliest *bool
	/*EscalationPolicyIds
	  Filters the results, showing only on-calls for the specified escalation policy IDs.

	*/
	EscalationPolicyIds []string
	/*Include
	  Array of additional details to include.

	*/
	Include []string
	/*ScheduleIds
	  Filters the results, showing only on-calls for the specified schedule IDs. If `null` is provided in the array, it includes permanent on-calls due to direct user escalation targets.

	*/
	ScheduleIds []string
	/*Since
	  The start of the time range over which you want to search. If an on-call period overlaps with the range, it will be included in the result. Defaults to current time. The search range cannot exceed 3 months.

	*/
	Since *strfmt.DateTime
	/*TimeZone
	  Time zone in which dates in the result will be rendered.

	*/
	TimeZone *string
	/*Until
	  The end of the time range over which you want to search. If an on-call period overlaps with the range, it will be included in the result. Defaults to current time. The search range cannot exceed 3 months, and the `until` time cannot be before the `since` time.

	*/
	Until *strfmt.DateTime
	/*UserIds
	  Filters the results, showing only on-calls for the specified user IDs.

	*/
	UserIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list on calls params
func (o *ListOnCallsParams) WithTimeout(timeout time.Duration) *ListOnCallsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list on calls params
func (o *ListOnCallsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list on calls params
func (o *ListOnCallsParams) WithContext(ctx context.Context) *ListOnCallsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list on calls params
func (o *ListOnCallsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list on calls params
func (o *ListOnCallsParams) WithHTTPClient(client *http.Client) *ListOnCallsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list on calls params
func (o *ListOnCallsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEarliest adds the earliest to the list on calls params
func (o *ListOnCallsParams) WithEarliest(earliest *bool) *ListOnCallsParams {
	o.SetEarliest(earliest)
	return o
}

// SetEarliest adds the earliest to the list on calls params
func (o *ListOnCallsParams) SetEarliest(earliest *bool) {
	o.Earliest = earliest
}

// WithEscalationPolicyIds adds the escalationPolicyIds to the list on calls params
func (o *ListOnCallsParams) WithEscalationPolicyIds(escalationPolicyIds []string) *ListOnCallsParams {
	o.SetEscalationPolicyIds(escalationPolicyIds)
	return o
}

// SetEscalationPolicyIds adds the escalationPolicyIds to the list on calls params
func (o *ListOnCallsParams) SetEscalationPolicyIds(escalationPolicyIds []string) {
	o.EscalationPolicyIds = escalationPolicyIds
}

// WithInclude adds the include to the list on calls params
func (o *ListOnCallsParams) WithInclude(include []string) *ListOnCallsParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the list on calls params
func (o *ListOnCallsParams) SetInclude(include []string) {
	o.Include = include
}

// WithScheduleIds adds the scheduleIds to the list on calls params
func (o *ListOnCallsParams) WithScheduleIds(scheduleIds []string) *ListOnCallsParams {
	o.SetScheduleIds(scheduleIds)
	return o
}

// SetScheduleIds adds the scheduleIds to the list on calls params
func (o *ListOnCallsParams) SetScheduleIds(scheduleIds []string) {
	o.ScheduleIds = scheduleIds
}

// WithSince adds the since to the list on calls params
func (o *ListOnCallsParams) WithSince(since *strfmt.DateTime) *ListOnCallsParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the list on calls params
func (o *ListOnCallsParams) SetSince(since *strfmt.DateTime) {
	o.Since = since
}

// WithTimeZone adds the timeZone to the list on calls params
func (o *ListOnCallsParams) WithTimeZone(timeZone *string) *ListOnCallsParams {
	o.SetTimeZone(timeZone)
	return o
}

// SetTimeZone adds the timeZone to the list on calls params
func (o *ListOnCallsParams) SetTimeZone(timeZone *string) {
	o.TimeZone = timeZone
}

// WithUntil adds the until to the list on calls params
func (o *ListOnCallsParams) WithUntil(until *strfmt.DateTime) *ListOnCallsParams {
	o.SetUntil(until)
	return o
}

// SetUntil adds the until to the list on calls params
func (o *ListOnCallsParams) SetUntil(until *strfmt.DateTime) {
	o.Until = until
}

// WithUserIds adds the userIds to the list on calls params
func (o *ListOnCallsParams) WithUserIds(userIds []string) *ListOnCallsParams {
	o.SetUserIds(userIds)
	return o
}

// SetUserIds adds the userIds to the list on calls params
func (o *ListOnCallsParams) SetUserIds(userIds []string) {
	o.UserIds = userIds
}

// WriteToRequest writes these params to a swagger request
func (o *ListOnCallsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Earliest != nil {

		// query param earliest
		var qrEarliest bool
		if o.Earliest != nil {
			qrEarliest = *o.Earliest
		}
		qEarliest := swag.FormatBool(qrEarliest)
		if qEarliest != "" {
			if err := r.SetQueryParam("earliest", qEarliest); err != nil {
				return err
			}
		}

	}

	valuesEscalationPolicyIds := o.EscalationPolicyIds

	joinedEscalationPolicyIds := swag.JoinByFormat(valuesEscalationPolicyIds, "multi")
	// query array param escalation_policy_ids[]
	if err := r.SetQueryParam("escalation_policy_ids[]", joinedEscalationPolicyIds...); err != nil {
		return err
	}

	valuesInclude := o.Include

	joinedInclude := swag.JoinByFormat(valuesInclude, "multi")
	// query array param include[]
	if err := r.SetQueryParam("include[]", joinedInclude...); err != nil {
		return err
	}

	valuesScheduleIds := o.ScheduleIds

	joinedScheduleIds := swag.JoinByFormat(valuesScheduleIds, "multi")
	// query array param schedule_ids[]
	if err := r.SetQueryParam("schedule_ids[]", joinedScheduleIds...); err != nil {
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