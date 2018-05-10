// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// APIBasedIntegrationAllOf1 api based integration all of1
// swagger:model apiBasedIntegrationAllOf1
type APIBasedIntegrationAllOf1 struct {

	// This is the unique key used to route events to this integration when received via the PagerDuty Events API.
	IntegrationKey string `json:"integration_key,omitempty"`
}

// Validate validates this api based integration all of1
func (m *APIBasedIntegrationAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIBasedIntegrationAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIBasedIntegrationAllOf1) UnmarshalBinary(b []byte) error {
	var res APIBasedIntegrationAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
