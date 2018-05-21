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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/luqasn/go-pagerduty2/models"
)

// NewUpdateContactMethodParams creates a new UpdateContactMethodParams object
// with the default values initialized.
func NewUpdateContactMethodParams() *UpdateContactMethodParams {
	var ()
	return &UpdateContactMethodParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateContactMethodParamsWithTimeout creates a new UpdateContactMethodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateContactMethodParamsWithTimeout(timeout time.Duration) *UpdateContactMethodParams {
	var ()
	return &UpdateContactMethodParams{

		timeout: timeout,
	}
}

// NewUpdateContactMethodParamsWithContext creates a new UpdateContactMethodParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateContactMethodParamsWithContext(ctx context.Context) *UpdateContactMethodParams {
	var ()
	return &UpdateContactMethodParams{

		Context: ctx,
	}
}

// NewUpdateContactMethodParamsWithHTTPClient creates a new UpdateContactMethodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateContactMethodParamsWithHTTPClient(client *http.Client) *UpdateContactMethodParams {
	var ()
	return &UpdateContactMethodParams{
		HTTPClient: client,
	}
}

/*UpdateContactMethodParams contains all the parameters to send to the API endpoint
for the update contact method operation typically these are written to a http.Request
*/
type UpdateContactMethodParams struct {

	/*ContactMethodID
	  The contact method ID on the user.

	*/
	ContactMethodID string
	/*ID*/
	ID string
	/*User
	  The user's contact method to be updated.

	*/
	User *models.UpdateContactMethodParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update contact method params
func (o *UpdateContactMethodParams) WithTimeout(timeout time.Duration) *UpdateContactMethodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update contact method params
func (o *UpdateContactMethodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update contact method params
func (o *UpdateContactMethodParams) WithContext(ctx context.Context) *UpdateContactMethodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update contact method params
func (o *UpdateContactMethodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update contact method params
func (o *UpdateContactMethodParams) WithHTTPClient(client *http.Client) *UpdateContactMethodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update contact method params
func (o *UpdateContactMethodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactMethodID adds the contactMethodID to the update contact method params
func (o *UpdateContactMethodParams) WithContactMethodID(contactMethodID string) *UpdateContactMethodParams {
	o.SetContactMethodID(contactMethodID)
	return o
}

// SetContactMethodID adds the contactMethodId to the update contact method params
func (o *UpdateContactMethodParams) SetContactMethodID(contactMethodID string) {
	o.ContactMethodID = contactMethodID
}

// WithID adds the id to the update contact method params
func (o *UpdateContactMethodParams) WithID(id string) *UpdateContactMethodParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update contact method params
func (o *UpdateContactMethodParams) SetID(id string) {
	o.ID = id
}

// WithUser adds the user to the update contact method params
func (o *UpdateContactMethodParams) WithUser(user *models.UpdateContactMethodParamsBody) *UpdateContactMethodParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the update contact method params
func (o *UpdateContactMethodParams) SetUser(user *models.UpdateContactMethodParamsBody) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateContactMethodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contact_method_id
	if err := r.SetPathParam("contact_method_id", o.ContactMethodID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.User != nil {
		if err := r.SetBodyParam(o.User); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}