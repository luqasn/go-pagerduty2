// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetMaintenanceWindowsOKBodyAllOf1 get maintenance windows o k body all of1
// swagger:model getMaintenanceWindowsOKBodyAllOf1
type GetMaintenanceWindowsOKBodyAllOf1 struct {

	// maintenance windows
	// Required: true
	MaintenanceWindows []*MaintenanceWindow `json:"maintenance_windows"`
}

// Validate validates this get maintenance windows o k body all of1
func (m *GetMaintenanceWindowsOKBodyAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaintenanceWindows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetMaintenanceWindowsOKBodyAllOf1) validateMaintenanceWindows(formats strfmt.Registry) error {

	if err := validate.Required("maintenance_windows", "body", m.MaintenanceWindows); err != nil {
		return err
	}

	for i := 0; i < len(m.MaintenanceWindows); i++ {
		if swag.IsZero(m.MaintenanceWindows[i]) { // not required
			continue
		}

		if m.MaintenanceWindows[i] != nil {
			if err := m.MaintenanceWindows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("maintenance_windows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetMaintenanceWindowsOKBodyAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMaintenanceWindowsOKBodyAllOf1) UnmarshalBinary(b []byte) error {
	var res GetMaintenanceWindowsOKBodyAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
