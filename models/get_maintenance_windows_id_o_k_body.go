// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetMaintenanceWindowsIDOKBody get maintenance windows Id o k body
// swagger:model getMaintenanceWindowsIdOKBody
type GetMaintenanceWindowsIDOKBody struct {

	// maintenance window
	// Required: true
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window"`
}

// Validate validates this get maintenance windows Id o k body
func (m *GetMaintenanceWindowsIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaintenanceWindow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetMaintenanceWindowsIDOKBody) validateMaintenanceWindow(formats strfmt.Registry) error {

	if err := validate.Required("maintenance_window", "body", m.MaintenanceWindow); err != nil {
		return err
	}

	if m.MaintenanceWindow != nil {
		if err := m.MaintenanceWindow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance_window")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetMaintenanceWindowsIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMaintenanceWindowsIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetMaintenanceWindowsIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
