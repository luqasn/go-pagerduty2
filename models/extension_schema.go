// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ExtensionSchema extension schema
// swagger:model ExtensionSchema
type ExtensionSchema struct {
	ExtensionSchemaReference

	ExtensionSchemaAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ExtensionSchema) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ExtensionSchemaReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ExtensionSchemaReference = aO0

	// AO1
	var aO1 ExtensionSchemaAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ExtensionSchemaAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ExtensionSchema) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ExtensionSchemaReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ExtensionSchemaAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this extension schema
func (m *ExtensionSchema) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ExtensionSchemaReference
	if err := m.ExtensionSchemaReference.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ExtensionSchemaAllOf1
	if err := m.ExtensionSchemaAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionSchema) UnmarshalBinary(b []byte) error {
	var res ExtensionSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
