// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IncidentUrgencyRuleAllOf1 incident urgency rule all of1
// swagger:model incidentUrgencyRuleAllOf1
type IncidentUrgencyRuleAllOf1 struct {

	// Incidents' urgency during support hours
	DuringSupportHours *IncidentUrgencyType `json:"during_support_hours,omitempty"`

	// Incidents' urgency outside support hours
	OutsideSupportHours *IncidentUrgencyType `json:"outside_support_hours,omitempty"`
}

// Validate validates this incident urgency rule all of1
func (m *IncidentUrgencyRuleAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuringSupportHours(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutsideSupportHours(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IncidentUrgencyRuleAllOf1) validateDuringSupportHours(formats strfmt.Registry) error {

	if swag.IsZero(m.DuringSupportHours) { // not required
		return nil
	}

	if m.DuringSupportHours != nil {
		if err := m.DuringSupportHours.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("during_support_hours")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentUrgencyRuleAllOf1) validateOutsideSupportHours(formats strfmt.Registry) error {

	if swag.IsZero(m.OutsideSupportHours) { // not required
		return nil
	}

	if m.OutsideSupportHours != nil {
		if err := m.OutsideSupportHours.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outside_support_hours")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IncidentUrgencyRuleAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncidentUrgencyRuleAllOf1) UnmarshalBinary(b []byte) error {
	var res IncidentUrgencyRuleAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}