// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EscalationPolicyReferenceAllOf1 escalation policy reference all of1
// swagger:model escalationPolicyReferenceAllOf1
type EscalationPolicyReferenceAllOf1 struct {

	// type
	// Enum: [escalation_policy escalation_policy_reference]
	Type string `json:"type,omitempty"`
}

// Validate validates this escalation policy reference all of1
func (m *EscalationPolicyReferenceAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var escalationPolicyReferenceAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["escalation_policy","escalation_policy_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		escalationPolicyReferenceAllOf1TypeTypePropEnum = append(escalationPolicyReferenceAllOf1TypeTypePropEnum, v)
	}
}

const (

	// EscalationPolicyReferenceAllOf1TypeEscalationPolicy captures enum value "escalation_policy"
	EscalationPolicyReferenceAllOf1TypeEscalationPolicy string = "escalation_policy"

	// EscalationPolicyReferenceAllOf1TypeEscalationPolicyReference captures enum value "escalation_policy_reference"
	EscalationPolicyReferenceAllOf1TypeEscalationPolicyReference string = "escalation_policy_reference"
)

// prop value enum
func (m *EscalationPolicyReferenceAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, escalationPolicyReferenceAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EscalationPolicyReferenceAllOf1) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EscalationPolicyReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EscalationPolicyReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res EscalationPolicyReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
