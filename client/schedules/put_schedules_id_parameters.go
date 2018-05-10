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

// NewPutSchedulesIDParams creates a new PutSchedulesIDParams object
// with the default values initialized.
func NewPutSchedulesIDParams() *PutSchedulesIDParams {
	var (
		overflowDefault = bool(false)
	)
	return &PutSchedulesIDParams{
		Overflow: &overflowDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSchedulesIDParamsWithTimeout creates a new PutSchedulesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSchedulesIDParamsWithTimeout(timeout time.Duration) *PutSchedulesIDParams {
	var (
		overflowDefault = bool(false)
	)
	return &PutSchedulesIDParams{
		Overflow: &overflowDefault,

		timeout: timeout,
	}
}

// NewPutSchedulesIDParamsWithContext creates a new PutSchedulesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSchedulesIDParamsWithContext(ctx context.Context) *PutSchedulesIDParams {
	var (
		overflowDefault = bool(false)
	)
	return &PutSchedulesIDParams{
		Overflow: &overflowDefault,

		Context: ctx,
	}
}

// NewPutSchedulesIDParamsWithHTTPClient creates a new PutSchedulesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSchedulesIDParamsWithHTTPClient(client *http.Client) *PutSchedulesIDParams {
	var (
		overflowDefault = bool(false)
	)
	return &PutSchedulesIDParams{
		Overflow:   &overflowDefault,
		HTTPClient: client,
	}
}

/*PutSchedulesIDParams contains all the parameters to send to the API endpoint
for the put schedules ID operation typically these are written to a http.Request
*/
type PutSchedulesIDParams struct {

	/*ID*/
	ID string
	/*Overflow
	  Any on-call schedule entries that pass the date range bounds will be truncated at the bounds, unless the parameter `overflow=true` is passed. This parameter defaults to false.
	For instance, if your schedule is a rotation that changes daily at midnight UTC, and your date range is from `2011-06-01T10:00:00Z` to `2011-06-01T14:00:00Z`:

	- If you don't pass the `overflow=true` parameter, you will get one schedule entry returned with a start of `2011-06-01T10:00:00Z` and end of `2011-06-01T14:00:00Z`.
	- If you do pass the `overflow=true` parameter, you will get one schedule entry returned with a start of `2011-06-01T00:00:00Z` and end of `2011-06-02T00:00:00Z`.


	*/
	Overflow *bool
	/*Schedule
	  The schedule to be updated.

	*/
	Schedule *models.PutSchedulesIDParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put schedules ID params
func (o *PutSchedulesIDParams) WithTimeout(timeout time.Duration) *PutSchedulesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put schedules ID params
func (o *PutSchedulesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put schedules ID params
func (o *PutSchedulesIDParams) WithContext(ctx context.Context) *PutSchedulesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put schedules ID params
func (o *PutSchedulesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put schedules ID params
func (o *PutSchedulesIDParams) WithHTTPClient(client *http.Client) *PutSchedulesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put schedules ID params
func (o *PutSchedulesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put schedules ID params
func (o *PutSchedulesIDParams) WithID(id string) *PutSchedulesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put schedules ID params
func (o *PutSchedulesIDParams) SetID(id string) {
	o.ID = id
}

// WithOverflow adds the overflow to the put schedules ID params
func (o *PutSchedulesIDParams) WithOverflow(overflow *bool) *PutSchedulesIDParams {
	o.SetOverflow(overflow)
	return o
}

// SetOverflow adds the overflow to the put schedules ID params
func (o *PutSchedulesIDParams) SetOverflow(overflow *bool) {
	o.Overflow = overflow
}

// WithSchedule adds the schedule to the put schedules ID params
func (o *PutSchedulesIDParams) WithSchedule(schedule *models.PutSchedulesIDParamsBody) *PutSchedulesIDParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the put schedules ID params
func (o *PutSchedulesIDParams) SetSchedule(schedule *models.PutSchedulesIDParamsBody) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *PutSchedulesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

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