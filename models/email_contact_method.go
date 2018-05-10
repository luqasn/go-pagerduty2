// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EmailContactMethod email contact method
// swagger:model EmailContactMethod
type EmailContactMethod struct {
	ContactMethod

	EmailContactMethodAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *EmailContactMethod) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ContactMethod
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ContactMethod = aO0

	// AO1
	var aO1 EmailContactMethodAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.EmailContactMethodAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m EmailContactMethod) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ContactMethod)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.EmailContactMethodAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this email contact method
func (m *EmailContactMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ContactMethod
	if err := m.ContactMethod.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with EmailContactMethodAllOf1
	if err := m.EmailContactMethodAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *EmailContactMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailContactMethod) UnmarshalBinary(b []byte) error {
	var res EmailContactMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}