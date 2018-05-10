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

// EscalationPolicyAllOf1 escalation policy all of1
// swagger:model escalationPolicyAllOf1
type EscalationPolicyAllOf1 struct {

	// Escalation policy description.
	Description string `json:"description,omitempty"`

	// escalation rules
	// Required: true
	EscalationRules []*EscalationRule `json:"escalation_rules"`

	// The name of the escalation policy.
	// Required: true
	Name *string `json:"name"`

	// The number of times the escalation policy will repeat after reaching the end of its escalation.
	// Minimum: 0
	NumLoops *int64 `json:"num_loops,omitempty"`

	// Whether or not to allow this policy to repeat its escalation rules after the last rule is finished.
	RepeatEnabled *bool `json:"repeat_enabled,omitempty"`

	// services
	// Min Length: 0
	Services []*ServiceReference `json:"services"`

	// Teams associated with the policy. Account must have the `teams` ability to use this parameter.
	// Min Length: 0
	Teams []*TeamReference `json:"teams"`
}

// Validate validates this escalation policy all of1
func (m *EscalationPolicyAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEscalationRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumLoops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EscalationPolicyAllOf1) validateEscalationRules(formats strfmt.Registry) error {

	if err := validate.Required("escalation_rules", "body", m.EscalationRules); err != nil {
		return err
	}

	for i := 0; i < len(m.EscalationRules); i++ {
		if swag.IsZero(m.EscalationRules[i]) { // not required
			continue
		}

		if m.EscalationRules[i] != nil {
			if err := m.EscalationRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("escalation_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EscalationPolicyAllOf1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EscalationPolicyAllOf1) validateNumLoops(formats strfmt.Registry) error {

	if swag.IsZero(m.NumLoops) { // not required
		return nil
	}

	if err := validate.MinimumInt("num_loops", "body", int64(*m.NumLoops), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *EscalationPolicyAllOf1) validateServices(formats strfmt.Registry) error {

	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {
		if swag.IsZero(m.Services[i]) { // not required
			continue
		}

		if m.Services[i] != nil {
			if err := m.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EscalationPolicyAllOf1) validateTeams(formats strfmt.Registry) error {

	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	for i := 0; i < len(m.Teams); i++ {
		if swag.IsZero(m.Teams[i]) { // not required
			continue
		}

		if m.Teams[i] != nil {
			if err := m.Teams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EscalationPolicyAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EscalationPolicyAllOf1) UnmarshalBinary(b []byte) error {
	var res EscalationPolicyAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
