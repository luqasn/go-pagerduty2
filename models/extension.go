// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Extension extension
// swagger:model Extension
type Extension struct {
	ExtensionReference

	ExtensionAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Extension) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ExtensionReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ExtensionReference = aO0

	// AO1
	var aO1 ExtensionAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ExtensionAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Extension) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ExtensionReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ExtensionAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this extension
func (m *Extension) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ExtensionReference
	if err := m.ExtensionReference.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ExtensionAllOf1
	if err := m.ExtensionAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Extension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Extension) UnmarshalBinary(b []byte) error {
	var res Extension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
