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

// SnoozeIncidentParamsBody snooze incident params body
// swagger:model snoozeIncidentParamsBody
type SnoozeIncidentParamsBody struct {

	// The number of seconds to snooze the incident for. After this number of seconds has elapsed, the incident will return to the "triggered" state.
	// Required: true
	Duration *int64 `json:"duration"`
}

// Validate validates this snooze incident params body
func (m *SnoozeIncidentParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnoozeIncidentParamsBody) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnoozeIncidentParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnoozeIncidentParamsBody) UnmarshalBinary(b []byte) error {
	var res SnoozeIncidentParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
