// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetMaintenanceWindowsOKBody get maintenance windows o k body
// swagger:model getMaintenanceWindowsOKBody
type GetMaintenanceWindowsOKBody struct {
	Pagination

	GetMaintenanceWindowsOKBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetMaintenanceWindowsOKBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Pagination
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Pagination = aO0

	// AO1
	var aO1 GetMaintenanceWindowsOKBodyAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.GetMaintenanceWindowsOKBodyAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetMaintenanceWindowsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Pagination)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.GetMaintenanceWindowsOKBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get maintenance windows o k body
func (m *GetMaintenanceWindowsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Pagination
	if err := m.Pagination.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with GetMaintenanceWindowsOKBodyAllOf1
	if err := m.GetMaintenanceWindowsOKBodyAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetMaintenanceWindowsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMaintenanceWindowsOKBody) UnmarshalBinary(b []byte) error {
	var res GetMaintenanceWindowsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
