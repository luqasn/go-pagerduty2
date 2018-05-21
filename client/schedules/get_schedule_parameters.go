// Code generated by go-swagger; DO NOT EDIT.

package schedules

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

// NewGetScheduleParams creates a new GetScheduleParams object
// with the default values initialized.
func NewGetScheduleParams() *GetScheduleParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &GetScheduleParams{
		TimeZone: &timeZoneDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduleParamsWithTimeout creates a new GetScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScheduleParamsWithTimeout(timeout time.Duration) *GetScheduleParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &GetScheduleParams{
		TimeZone: &timeZoneDefault,

		timeout: timeout,
	}
}

// NewGetScheduleParamsWithContext creates a new GetScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScheduleParamsWithContext(ctx context.Context) *GetScheduleParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &GetScheduleParams{
		TimeZone: &timeZoneDefault,

		Context: ctx,
	}
}

// NewGetScheduleParamsWithHTTPClient creates a new GetScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScheduleParamsWithHTTPClient(client *http.Client) *GetScheduleParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &GetScheduleParams{
		TimeZone:   &timeZoneDefault,
		HTTPClient: client,
	}
}

/*GetScheduleParams contains all the parameters to send to the API endpoint
for the get schedule operation typically these are written to a http.Request
*/
type GetScheduleParams struct {

	/*ID*/
	ID string
	/*Since
	  The start of the date range over which you want to search.

	*/
	Since *strfmt.DateTime
	/*TimeZone
	  Time zone in which dates in the result will be rendered.

	*/
	TimeZone *string
	/*Until
	  The end of the date range over which you want to search.

	*/
	Until *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get schedule params
func (o *GetScheduleParams) WithTimeout(timeout time.Duration) *GetScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedule params
func (o *GetScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedule params
func (o *GetScheduleParams) WithContext(ctx context.Context) *GetScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedule params
func (o *GetScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedule params
func (o *GetScheduleParams) WithHTTPClient(client *http.Client) *GetScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedule params
func (o *GetScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get schedule params
func (o *GetScheduleParams) WithID(id string) *GetScheduleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get schedule params
func (o *GetScheduleParams) SetID(id string) {
	o.ID = id
}

// WithSince adds the since to the get schedule params
func (o *GetScheduleParams) WithSince(since *strfmt.DateTime) *GetScheduleParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the get schedule params
func (o *GetScheduleParams) SetSince(since *strfmt.DateTime) {
	o.Since = since
}

// WithTimeZone adds the timeZone to the get schedule params
func (o *GetScheduleParams) WithTimeZone(timeZone *string) *GetScheduleParams {
	o.SetTimeZone(timeZone)
	return o
}

// SetTimeZone adds the timeZone to the get schedule params
func (o *GetScheduleParams) SetTimeZone(timeZone *string) {
	o.TimeZone = timeZone
}

// WithUntil adds the until to the get schedule params
func (o *GetScheduleParams) WithUntil(until *strfmt.DateTime) *GetScheduleParams {
	o.SetUntil(until)
	return o
}

// SetUntil adds the until to the get schedule params
func (o *GetScheduleParams) SetUntil(until *strfmt.DateTime) {
	o.Until = until
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}