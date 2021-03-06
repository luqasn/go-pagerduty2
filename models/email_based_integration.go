// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EmailBasedIntegration email based integration
// swagger:model EmailBasedIntegration
type EmailBasedIntegration struct {
	Integration

	EmailBasedIntegrationAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *EmailBasedIntegration) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Integration
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Integration = aO0

	// AO1
	var aO1 EmailBasedIntegrationAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.EmailBasedIntegrationAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m EmailBasedIntegration) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Integration)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.EmailBasedIntegrationAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this email based integration
func (m *EmailBasedIntegration) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Integration
	if err := m.Integration.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with EmailBasedIntegrationAllOf1
	if err := m.EmailBasedIntegrationAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *EmailBasedIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailBasedIntegration) UnmarshalBinary(b []byte) error {
	var res EmailBasedIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
