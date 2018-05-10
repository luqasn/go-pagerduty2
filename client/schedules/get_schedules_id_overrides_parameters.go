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
)

// NewGetSchedulesIDOverridesParams creates a new GetSchedulesIDOverridesParams object
// with the default values initialized.
func NewGetSchedulesIDOverridesParams() *GetSchedulesIDOverridesParams {
	var ()
	return &GetSchedulesIDOverridesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSchedulesIDOverridesParamsWithTimeout creates a new GetSchedulesIDOverridesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSchedulesIDOverridesParamsWithTimeout(timeout time.Duration) *GetSchedulesIDOverridesParams {
	var ()
	return &GetSchedulesIDOverridesParams{

		timeout: timeout,
	}
}

// NewGetSchedulesIDOverridesParamsWithContext creates a new GetSchedulesIDOverridesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSchedulesIDOverridesParamsWithContext(ctx context.Context) *GetSchedulesIDOverridesParams {
	var ()
	return &GetSchedulesIDOverridesParams{

		Context: ctx,
	}
}

// NewGetSchedulesIDOverridesParamsWithHTTPClient creates a new GetSchedulesIDOverridesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSchedulesIDOverridesParamsWithHTTPClient(client *http.Client) *GetSchedulesIDOverridesParams {
	var ()
	return &GetSchedulesIDOverridesParams{
		HTTPClient: client,
	}
}

/*GetSchedulesIDOverridesParams contains all the parameters to send to the API endpoint
for the get schedules ID overrides operation typically these are written to a http.Request
*/
type GetSchedulesIDOverridesParams struct {

	/*Editable
	  When this parameter is present, only editable overrides will be returned. The result will only include the id of the override if this parameter is present. Only future overrides are editable.

	*/
	Editable *bool
	/*ID*/
	ID string
	/*Overflow
	  Any on-call schedule entries that pass the date range bounds will be truncated at the bounds, unless the parameter overflow=true is passed. This parameter defaults to false.

	*/
	Overflow *bool
	/*Since
	  The start of the date range over which you want to search.

	*/
	Since strfmt.DateTime
	/*Until
	  The end of the date range over which you want to search.

	*/
	Until strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) WithTimeout(timeout time.Duration) *GetSchedulesIDOverridesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) WithContext(ctx context.Context) *GetSchedulesIDOverridesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) WithHTTPClient(client *http.Client) *GetSchedulesIDOverridesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEditable adds the editable to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) WithEditable(editable *bool) *GetSchedulesIDOverridesParams {
	o.SetEditable(editable)
	return o
}

// SetEditable adds the editable to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) SetEditable(editable *bool) {
	o.Editable = editable
}

// WithID adds the id to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) WithID(id string) *GetSchedulesIDOverridesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) SetID(id string) {
	o.ID = id
}

// WithOverflow adds the overflow to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) WithOverflow(overflow *bool) *GetSchedulesIDOverridesParams {
	o.SetOverflow(overflow)
	return o
}

// SetOverflow adds the overflow to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) SetOverflow(overflow *bool) {
	o.Overflow = overflow
}

// WithSince adds the since to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) WithSince(since strfmt.DateTime) *GetSchedulesIDOverridesParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) SetSince(since strfmt.DateTime) {
	o.Since = since
}

// WithUntil adds the until to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) WithUntil(until strfmt.DateTime) *GetSchedulesIDOverridesParams {
	o.SetUntil(until)
	return o
}

// SetUntil adds the until to the get schedules ID overrides params
func (o *GetSchedulesIDOverridesParams) SetUntil(until strfmt.DateTime) {
	o.Until = until
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchedulesIDOverridesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Editable != nil {

		// query param editable
		var qrEditable bool
		if o.Editable != nil {
			qrEditable = *o.Editable
		}
		qEditable := swag.FormatBool(qrEditable)
		if qEditable != "" {
			if err := r.SetQueryParam("editable", qEditable); err != nil {
				return err
			}
		}

	}

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

	// query param since
	qrSince := o.Since
	qSince := qrSince.String()
	if qSince != "" {
		if err := r.SetQueryParam("since", qSince); err != nil {
			return err
		}
	}

	// query param until
	qrUntil := o.Until
	qUntil := qrUntil.String()
	if qUntil != "" {
		if err := r.SetQueryParam("until", qUntil); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}