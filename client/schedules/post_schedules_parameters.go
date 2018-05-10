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

// NewPostSchedulesParams creates a new PostSchedulesParams object
// with the default values initialized.
func NewPostSchedulesParams() *PostSchedulesParams {
	var (
		overflowDefault = bool(false)
	)
	return &PostSchedulesParams{
		Overflow: &overflowDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSchedulesParamsWithTimeout creates a new PostSchedulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSchedulesParamsWithTimeout(timeout time.Duration) *PostSchedulesParams {
	var (
		overflowDefault = bool(false)
	)
	return &PostSchedulesParams{
		Overflow: &overflowDefault,

		timeout: timeout,
	}
}

// NewPostSchedulesParamsWithContext creates a new PostSchedulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSchedulesParamsWithContext(ctx context.Context) *PostSchedulesParams {
	var (
		overflowDefault = bool(false)
	)
	return &PostSchedulesParams{
		Overflow: &overflowDefault,

		Context: ctx,
	}
}

// NewPostSchedulesParamsWithHTTPClient creates a new PostSchedulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSchedulesParamsWithHTTPClient(client *http.Client) *PostSchedulesParams {
	var (
		overflowDefault = bool(false)
	)
	return &PostSchedulesParams{
		Overflow:   &overflowDefault,
		HTTPClient: client,
	}
}

/*PostSchedulesParams contains all the parameters to send to the API endpoint
for the post schedules operation typically these are written to a http.Request
*/
type PostSchedulesParams struct {

	/*Overflow
	  Any on-call schedule entries that pass the date range bounds will be truncated at the bounds, unless the parameter `overflow=true` is passed. This parameter defaults to false.
	For instance, if your schedule is a rotation that changes daily at midnight UTC, and your date range is from `2011-06-01T10:00:00Z` to `2011-06-01T14:00:00Z`:

	- If you don't pass the `overflow=true` parameter, you will get one schedule entry returned with a start of `2011-06-01T10:00:00Z` and end of `2011-06-01T14:00:00Z`.
	- If you do pass the `overflow=true` parameter, you will get one schedule entry returned with a start of `2011-06-01T00:00:00Z` and end of `2011-06-02T00:00:00Z`.


	*/
	Overflow *bool
	/*Schedule
	  The schedule to be created.

	*/
	Schedule *models.PostSchedulesParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post schedules params
func (o *PostSchedulesParams) WithTimeout(timeout time.Duration) *PostSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post schedules params
func (o *PostSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post schedules params
func (o *PostSchedulesParams) WithContext(ctx context.Context) *PostSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post schedules params
func (o *PostSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post schedules params
func (o *PostSchedulesParams) WithHTTPClient(client *http.Client) *PostSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post schedules params
func (o *PostSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOverflow adds the overflow to the post schedules params
func (o *PostSchedulesParams) WithOverflow(overflow *bool) *PostSchedulesParams {
	o.SetOverflow(overflow)
	return o
}

// SetOverflow adds the overflow to the post schedules params
func (o *PostSchedulesParams) SetOverflow(overflow *bool) {
	o.Overflow = overflow
}

// WithSchedule adds the schedule to the post schedules params
func (o *PostSchedulesParams) WithSchedule(schedule *models.PostSchedulesParamsBody) *PostSchedulesParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the post schedules params
func (o *PostSchedulesParams) SetSchedule(schedule *models.PostSchedulesParamsBody) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *PostSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}