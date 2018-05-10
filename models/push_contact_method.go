// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PushContactMethod push contact method
// swagger:model PushContactMethod
type PushContactMethod struct {
	ContactMethod

	PushContactMethodAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PushContactMethod) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ContactMethod
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ContactMethod = aO0

	// AO1
	var aO1 PushContactMethodAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PushContactMethodAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PushContactMethod) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ContactMethod)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PushContactMethodAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this push contact method
func (m *PushContactMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ContactMethod
	if err := m.ContactMethod.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PushContactMethodAllOf1
	if err := m.PushContactMethodAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PushContactMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PushContactMethod) UnmarshalBinary(b []byte) error {
	var res PushContactMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}