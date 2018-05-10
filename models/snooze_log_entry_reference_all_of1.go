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

// SnoozeLogEntryReferenceAllOf1 snooze log entry reference all of1
// swagger:model snoozeLogEntryReferenceAllOf1
type SnoozeLogEntryReferenceAllOf1 struct {

	// type
	// Enum: [snooze_log_entry snooze_log_entry_reference]
	Type string `json:"type,omitempty"`
}

// Validate validates this snooze log entry reference all of1
func (m *SnoozeLogEntryReferenceAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var snoozeLogEntryReferenceAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["snooze_log_entry","snooze_log_entry_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snoozeLogEntryReferenceAllOf1TypeTypePropEnum = append(snoozeLogEntryReferenceAllOf1TypeTypePropEnum, v)
	}
}

const (

	// SnoozeLogEntryReferenceAllOf1TypeSnoozeLogEntry captures enum value "snooze_log_entry"
	SnoozeLogEntryReferenceAllOf1TypeSnoozeLogEntry string = "snooze_log_entry"

	// SnoozeLogEntryReferenceAllOf1TypeSnoozeLogEntryReference captures enum value "snooze_log_entry_reference"
	SnoozeLogEntryReferenceAllOf1TypeSnoozeLogEntryReference string = "snooze_log_entry_reference"
)

// prop value enum
func (m *SnoozeLogEntryReferenceAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, snoozeLogEntryReferenceAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SnoozeLogEntryReferenceAllOf1) validateType(formats strfmt.Registry) error {

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
func (m *SnoozeLogEntryReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnoozeLogEntryReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res SnoozeLogEntryReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}