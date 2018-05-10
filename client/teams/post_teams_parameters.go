// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewPostTeamsParams creates a new PostTeamsParams object
// with the default values initialized.
func NewPostTeamsParams() *PostTeamsParams {
	var ()
	return &PostTeamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTeamsParamsWithTimeout creates a new PostTeamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTeamsParamsWithTimeout(timeout time.Duration) *PostTeamsParams {
	var ()
	return &PostTeamsParams{

		timeout: timeout,
	}
}

// NewPostTeamsParamsWithContext creates a new PostTeamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTeamsParamsWithContext(ctx context.Context) *PostTeamsParams {
	var ()
	return &PostTeamsParams{

		Context: ctx,
	}
}

// NewPostTeamsParamsWithHTTPClient creates a new PostTeamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTeamsParamsWithHTTPClient(client *http.Client) *PostTeamsParams {
	var ()
	return &PostTeamsParams{
		HTTPClient: client,
	}
}

/*PostTeamsParams contains all the parameters to send to the API endpoint
for the post teams operation typically these are written to a http.Request
*/
type PostTeamsParams struct {

	/*Team
	  The team to be created.

	*/
	Team *models.PostTeamsParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post teams params
func (o *PostTeamsParams) WithTimeout(timeout time.Duration) *PostTeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post teams params
func (o *PostTeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post teams params
func (o *PostTeamsParams) WithContext(ctx context.Context) *PostTeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post teams params
func (o *PostTeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post teams params
func (o *PostTeamsParams) WithHTTPClient(client *http.Client) *PostTeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post teams params
func (o *PostTeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeam adds the team to the post teams params
func (o *PostTeamsParams) WithTeam(team *models.PostTeamsParamsBody) *PostTeamsParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the post teams params
func (o *PostTeamsParams) SetTeam(team *models.PostTeamsParamsBody) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *PostTeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Team != nil {
		if err := r.SetBodyParam(o.Team); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
