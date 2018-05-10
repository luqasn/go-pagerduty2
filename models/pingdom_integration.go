// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PingdomIntegration pingdom integration
// swagger:model PingdomIntegration
type PingdomIntegration struct {
	EmailBasedIntegration

	PingdomIntegrationReference
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PingdomIntegration) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EmailBasedIntegration
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EmailBasedIntegration = aO0

	// AO1
	var aO1 PingdomIntegrationReference
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PingdomIntegrationReference = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PingdomIntegration) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EmailBasedIntegration)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PingdomIntegrationReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pingdom integration
func (m *PingdomIntegration) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EmailBasedIntegration
	if err := m.EmailBasedIntegration.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PingdomIntegrationReference
	if err := m.PingdomIntegrationReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PingdomIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PingdomIntegration) UnmarshalBinary(b []byte) error {
	var res PingdomIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}