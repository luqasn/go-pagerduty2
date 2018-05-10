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

// PutEscalationPoliciesIDOKBody put escalation policies Id o k body
// swagger:model putEscalationPoliciesIdOKBody
type PutEscalationPoliciesIDOKBody struct {

	// escalation policy
	// Required: true
	EscalationPolicy *EscalationPolicy `json:"escalation_policy"`
}

// Validate validates this put escalation policies Id o k body
func (m *PutEscalationPoliciesIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEscalationPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutEscalationPoliciesIDOKBody) validateEscalationPolicy(formats strfmt.Registry) error {

	if err := validate.Required("escalation_policy", "body", m.EscalationPolicy); err != nil {
		return err
	}

	if m.EscalationPolicy != nil {
		if err := m.EscalationPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("escalation_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutEscalationPoliciesIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutEscalationPoliciesIDOKBody) UnmarshalBinary(b []byte) error {
	var res PutEscalationPoliciesIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
