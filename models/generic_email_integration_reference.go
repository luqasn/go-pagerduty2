// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GenericEmailIntegrationReference generic email integration reference
// swagger:model GenericEmailIntegrationReference
type GenericEmailIntegrationReference struct {
	IntegrationReference

	GenericEmailIntegrationReferenceAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GenericEmailIntegrationReference) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 IntegrationReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.IntegrationReference = aO0

	// AO1
	var aO1 GenericEmailIntegrationReferenceAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.GenericEmailIntegrationReferenceAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GenericEmailIntegrationReference) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.IntegrationReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.GenericEmailIntegrationReferenceAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this generic email integration reference
func (m *GenericEmailIntegrationReference) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with IntegrationReference
	if err := m.IntegrationReference.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with GenericEmailIntegrationReferenceAllOf1
	if err := m.GenericEmailIntegrationReferenceAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GenericEmailIntegrationReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericEmailIntegrationReference) UnmarshalBinary(b []byte) error {
	var res GenericEmailIntegrationReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
