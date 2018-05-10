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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// NewPostSchedulesPreviewParams creates a new PostSchedulesPreviewParams object
// with the default values initialized.
func NewPostSchedulesPreviewParams() *PostSchedulesPreviewParams {
	var (
		overflowDefault = bool(false)
	)
	return &PostSchedulesPreviewParams{
		Overflow: &overflowDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSchedulesPreviewParamsWithTimeout creates a new PostSchedulesPreviewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSchedulesPreviewParamsWithTimeout(timeout time.Duration) *PostSchedulesPreviewParams {
	var (
		overflowDefault = bool(false)
	)
	return &PostSchedulesPreviewParams{
		Overflow: &overflowDefault,

		timeout: timeout,
	}
}

// NewPostSchedulesPreviewParamsWithContext creates a new PostSchedulesPreviewParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSchedulesPreviewParamsWithContext(ctx context.Context) *PostSchedulesPreviewParams {
	var (
		overflowDefault = bool(false)
	)
	return &PostSchedulesPreviewParams{
		Overflow: &overflowDefault,

		Context: ctx,
	}
}

// NewPostSchedulesPreviewParamsWithHTTPClient creates a new PostSchedulesPreviewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSchedulesPreviewParamsWithHTTPClient(client *http.Client) *PostSchedulesPreviewParams {
	var (
		overflowDefault = bool(false)
	)
	return &PostSchedulesPreviewParams{
		Overflow:   &overflowDefault,
		HTTPClient: client,
	}
}

/*PostSchedulesPreviewParams contains all the parameters to send to the API endpoint
for the post schedules preview operation typically these are written to a http.Request
*/
type PostSchedulesPreviewParams struct {

	/*Overflow
	  Any on-call schedule entries that pass the date range bounds will be truncated at the bounds, unless the parameter `overflow=true` is passed. This parameter defaults to false.
	For instance, if your schedule is a rotation that changes daily at midnight UTC, and your date range is from `2011-06-01T10:00:00Z` to `2011-06-01T14:00:00Z`:

	- If you don't pass the `overflow=true` parameter, you will get one schedule entry returned with a start of `2011-06-01T10:00:00Z` and end of `2011-06-01T14:00:00Z`.
	- If you do pass the `overflow=true` parameter, you will get one schedule entry returned with a start of `2011-06-01T00:00:00Z` and end of `2011-06-02T00:00:00Z`.


	*/
	Overflow *bool
	/*Schedule
	  The schedule to be previewed.

	*/
	Schedule *models.PostSchedulesPreviewParamsBody
	/*Since
	  The start of the date range over which you want to search.

	*/
	Since *strfmt.DateTime
	/*Until
	  The end of the date range over which you want to search.

	*/
	Until *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post schedules preview params
func (o *PostSchedulesPreviewParams) WithTimeout(timeout time.Duration) *PostSchedulesPreviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post schedules preview params
func (o *PostSchedulesPreviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post schedules preview params
func (o *PostSchedulesPreviewParams) WithContext(ctx context.Context) *PostSchedulesPreviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post schedules preview params
func (o *PostSchedulesPreviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post schedules preview params
func (o *PostSchedulesPreviewParams) WithHTTPClient(client *http.Client) *PostSchedulesPreviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post schedules preview params
func (o *PostSchedulesPreviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOverflow adds the overflow to the post schedules preview params
func (o *PostSchedulesPreviewParams) WithOverflow(overflow *bool) *PostSchedulesPreviewParams {
	o.SetOverflow(overflow)
	return o
}

// SetOverflow adds the overflow to the post schedules preview params
func (o *PostSchedulesPreviewParams) SetOverflow(overflow *bool) {
	o.Overflow = overflow
}

// WithSchedule adds the schedule to the post schedules preview params
func (o *PostSchedulesPreviewParams) WithSchedule(schedule *models.PostSchedulesPreviewParamsBody) *PostSchedulesPreviewParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the post schedules preview params
func (o *PostSchedulesPreviewParams) SetSchedule(schedule *models.PostSchedulesPreviewParamsBody) {
	o.Schedule = schedule
}

// WithSince adds the since to the post schedules preview params
func (o *PostSchedulesPreviewParams) WithSince(since *strfmt.DateTime) *PostSchedulesPreviewParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the post schedules preview params
func (o *PostSchedulesPreviewParams) SetSince(since *strfmt.DateTime) {
	o.Since = since
}

// WithUntil adds the until to the post schedules preview params
func (o *PostSchedulesPreviewParams) WithUntil(until *strfmt.DateTime) *PostSchedulesPreviewParams {
	o.SetUntil(until)
	return o
}

// SetUntil adds the until to the post schedules preview params
func (o *PostSchedulesPreviewParams) SetUntil(until *strfmt.DateTime) {
	o.Until = until
}

// WriteToRequest writes these params to a swagger request
func (o *PostSchedulesPreviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Overflow != nil {

		// query param overflow
		var qrOverflow bool
		if o.Overflow != nil {
			qrOverflow = *o.Overflow
		}
		qOverflow := swag.FormatBool(qrOverflow)
		if qOverflow != "" {
			if err := r.SetQueryParam("overflow", qOverflow); err != nil {
				return err
			}
		}

	}

	if o.Schedule != nil {
		if err := r.SetBodyParam(o.Schedule); err != nil {
			return err
		}
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
