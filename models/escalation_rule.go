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

// EscalationRule escalation rule
// swagger:model EscalationRule
type EscalationRule struct {

	// The number of minutes before an unacknowledged incident escalates away from this rule.
	// Required: true
	EscalationDelayInMinutes *int64 `json:"escalation_delay_in_minutes"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The targets an incident should be assigned to upon reaching this rule.
	// Required: true
	// Max Items: 10
	// Min Items: 1
	Targets []*EscalationTarget `json:"targets"`
}

// Validate validates this escalation rule
func (m *EscalationRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEscalationDelayInMinutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EscalationRule) validateEscalationDelayInMinutes(formats strfmt.Registry) error {

	if err := validate.Required("escalation_delay_in_minutes", "body", m.EscalationDelayInMinutes); err != nil {
		return err
	}

	return nil
}

func (m *EscalationRule) validateTargets(formats strfmt.Registry) error {

	if err := validate.Required("targets", "body", m.Targets); err != nil {
		return err
	}

	iTargetsSize := int64(len(m.Targets))

	if err := validate.MinItems("targets", "body", iTargetsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("targets", "body", iTargetsSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m.Targets); i++ {
		if swag.IsZero(m.Targets[i]) { // not required
			continue
		}

		if m.Targets[i] != nil {
			if err := m.Targets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EscalationRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EscalationRule) UnmarshalBinary(b []byte) error {
	var res EscalationRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}