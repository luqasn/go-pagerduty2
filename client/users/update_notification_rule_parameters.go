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

// NewUpdateNotificationRuleParams creates a new UpdateNotificationRuleParams object
// with the default values initialized.
func NewUpdateNotificationRuleParams() *UpdateNotificationRuleParams {
	var ()
	return &UpdateNotificationRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNotificationRuleParamsWithTimeout creates a new UpdateNotificationRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNotificationRuleParamsWithTimeout(timeout time.Duration) *UpdateNotificationRuleParams {
	var ()
	return &UpdateNotificationRuleParams{

		timeout: timeout,
	}
}

// NewUpdateNotificationRuleParamsWithContext creates a new UpdateNotificationRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNotificationRuleParamsWithContext(ctx context.Context) *UpdateNotificationRuleParams {
	var ()
	return &UpdateNotificationRuleParams{

		Context: ctx,
	}
}

// NewUpdateNotificationRuleParamsWithHTTPClient creates a new UpdateNotificationRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNotificationRuleParamsWithHTTPClient(client *http.Client) *UpdateNotificationRuleParams {
	var ()
	return &UpdateNotificationRuleParams{
		HTTPClient: client,
	}
}

/*UpdateNotificationRuleParams contains all the parameters to send to the API endpoint
for the update notification rule operation typically these are written to a http.Request
*/
type UpdateNotificationRuleParams struct {

	/*ID*/
	ID string
	/*NotificationRule
	  The user's notification rule to be updated.

	*/
	NotificationRule *models.UpdateNotificationRuleParamsBody
	/*NotificationRuleID
	  The notification rule ID on the user.

	*/
	NotificationRuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update notification rule params
func (o *UpdateNotificationRuleParams) WithTimeout(timeout time.Duration) *UpdateNotificationRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update notification rule params
func (o *UpdateNotificationRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update notification rule params
func (o *UpdateNotificationRuleParams) WithContext(ctx context.Context) *UpdateNotificationRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update notification rule params
func (o *UpdateNotificationRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update notification rule params
func (o *UpdateNotificationRuleParams) WithHTTPClient(client *http.Client) *UpdateNotificationRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update notification rule params
func (o *UpdateNotificationRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update notification rule params
func (o *UpdateNotificationRuleParams) WithID(id string) *UpdateNotificationRuleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update notification rule params
func (o *UpdateNotificationRuleParams) SetID(id string) {
	o.ID = id
}

// WithNotificationRule adds the notificationRule to the update notification rule params
func (o *UpdateNotificationRuleParams) WithNotificationRule(notificationRule *models.UpdateNotificationRuleParamsBody) *UpdateNotificationRuleParams {
	o.SetNotificationRule(notificationRule)
	return o
}

// SetNotificationRule adds the notificationRule to the update notification rule params
func (o *UpdateNotificationRuleParams) SetNotificationRule(notificationRule *models.UpdateNotificationRuleParamsBody) {
	o.NotificationRule = notificationRule
}

// WithNotificationRuleID adds the notificationRuleID to the update notification rule params
func (o *UpdateNotificationRuleParams) WithNotificationRuleID(notificationRuleID string) *UpdateNotificationRuleParams {
	o.SetNotificationRuleID(notificationRuleID)
	return o
}

// SetNotificationRuleID adds the notificationRuleId to the update notification rule params
func (o *UpdateNotificationRuleParams) SetNotificationRuleID(notificationRuleID string) {
	o.NotificationRuleID = notificationRuleID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNotificationRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.NotificationRule != nil {
		if err := r.SetBodyParam(o.NotificationRule); err != nil {
			return err
		}
	}

	// path param notification_rule_id
	if err := r.SetPathParam("notification_rule_id", o.NotificationRuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}