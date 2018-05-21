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

// NewListNotificationRulesParams creates a new ListNotificationRulesParams object
// with the default values initialized.
func NewListNotificationRulesParams() *ListNotificationRulesParams {
	var ()
	return &ListNotificationRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListNotificationRulesParamsWithTimeout creates a new ListNotificationRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListNotificationRulesParamsWithTimeout(timeout time.Duration) *ListNotificationRulesParams {
	var ()
	return &ListNotificationRulesParams{

		timeout: timeout,
	}
}

// NewListNotificationRulesParamsWithContext creates a new ListNotificationRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListNotificationRulesParamsWithContext(ctx context.Context) *ListNotificationRulesParams {
	var ()
	return &ListNotificationRulesParams{

		Context: ctx,
	}
}

// NewListNotificationRulesParamsWithHTTPClient creates a new ListNotificationRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListNotificationRulesParamsWithHTTPClient(client *http.Client) *ListNotificationRulesParams {
	var ()
	return &ListNotificationRulesParams{
		HTTPClient: client,
	}
}

/*ListNotificationRulesParams contains all the parameters to send to the API endpoint
for the list notification rules operation typically these are written to a http.Request
*/
type ListNotificationRulesParams struct {

	/*ID*/
	ID string
	/*Include
	  Array of additional details to include.

	*/
	Include []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list notification rules params
func (o *ListNotificationRulesParams) WithTimeout(timeout time.Duration) *ListNotificationRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list notification rules params
func (o *ListNotificationRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list notification rules params
func (o *ListNotificationRulesParams) WithContext(ctx context.Context) *ListNotificationRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list notification rules params
func (o *ListNotificationRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list notification rules params
func (o *ListNotificationRulesParams) WithHTTPClient(client *http.Client) *ListNotificationRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list notification rules params
func (o *ListNotificationRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list notification rules params
func (o *ListNotificationRulesParams) WithID(id string) *ListNotificationRulesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list notification rules params
func (o *ListNotificationRulesParams) SetID(id string) {
	o.ID = id
}

// WithInclude adds the include to the list notification rules params
func (o *ListNotificationRulesParams) WithInclude(include []string) *ListNotificationRulesParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the list notification rules params
func (o *ListNotificationRulesParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *ListNotificationRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	valuesInclude := o.Include

	joinedInclude := swag.JoinByFormat(valuesInclude, "multi")
	// query array param include[]
	if err := r.SetQueryParam("include[]", joinedInclude...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
