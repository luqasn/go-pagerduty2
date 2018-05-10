// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SQLMonitorIntegration Sql monitor integration
// swagger:model SqlMonitorIntegration
type SQLMonitorIntegration struct {
	EmailBasedIntegration

	SQLMonitorIntegrationReference

	SQLMonitorIntegrationAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SQLMonitorIntegration) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EmailBasedIntegration
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EmailBasedIntegration = aO0

	// AO1
	var aO1 SQLMonitorIntegrationReference
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.SQLMonitorIntegrationReference = aO1

	// AO2
	var aO2 SQLMonitorIntegrationAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.SQLMonitorIntegrationAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SQLMonitorIntegration) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.EmailBasedIntegration)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.SQLMonitorIntegrationReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.SQLMonitorIntegrationAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this Sql monitor integration
func (m *SQLMonitorIntegration) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EmailBasedIntegration
	if err := m.EmailBasedIntegration.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SQLMonitorIntegrationReference
	if err := m.SQLMonitorIntegrationReference.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SQLMonitorIntegrationAllOf2
	if err := m.SQLMonitorIntegrationAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SQLMonitorIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLMonitorIntegration) UnmarshalBinary(b []byte) error {
	var res SQLMonitorIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}