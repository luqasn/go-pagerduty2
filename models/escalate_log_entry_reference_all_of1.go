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

// EscalateLogEntryReferenceAllOf1 escalate log entry reference all of1
// swagger:model escalateLogEntryReferenceAllOf1
type EscalateLogEntryReferenceAllOf1 struct {

	// type
	// Enum: [escalate_log_entry escalate_log_entry_reference]
	Type string `json:"type,omitempty"`
}

// Validate validates this escalate log entry reference all of1
func (m *EscalateLogEntryReferenceAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var escalateLogEntryReferenceAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["escalate_log_entry","escalate_log_entry_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		escalateLogEntryReferenceAllOf1TypeTypePropEnum = append(escalateLogEntryReferenceAllOf1TypeTypePropEnum, v)
	}
}

const (

	// EscalateLogEntryReferenceAllOf1TypeEscalateLogEntry captures enum value "escalate_log_entry"
	EscalateLogEntryReferenceAllOf1TypeEscalateLogEntry string = "escalate_log_entry"

	// EscalateLogEntryReferenceAllOf1TypeEscalateLogEntryReference captures enum value "escalate_log_entry_reference"
	EscalateLogEntryReferenceAllOf1TypeEscalateLogEntryReference string = "escalate_log_entry_reference"
)

// prop value enum
func (m *EscalateLogEntryReferenceAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, escalateLogEntryReferenceAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EscalateLogEntryReferenceAllOf1) validateType(formats strfmt.Registry) error {

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
func (m *EscalateLogEntryReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EscalateLogEntryReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res EscalateLogEntryReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
