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

// NotifyLogEntryReferenceAllOf1 notify log entry reference all of1
// swagger:model notifyLogEntryReferenceAllOf1
type NotifyLogEntryReferenceAllOf1 struct {

	// type
	// Enum: [notify_log_entry notify_log_entry_reference]
	Type string `json:"type,omitempty"`
}

// Validate validates this notify log entry reference all of1
func (m *NotifyLogEntryReferenceAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var notifyLogEntryReferenceAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["notify_log_entry","notify_log_entry_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		notifyLogEntryReferenceAllOf1TypeTypePropEnum = append(notifyLogEntryReferenceAllOf1TypeTypePropEnum, v)
	}
}

const (

	// NotifyLogEntryReferenceAllOf1TypeNotifyLogEntry captures enum value "notify_log_entry"
	NotifyLogEntryReferenceAllOf1TypeNotifyLogEntry string = "notify_log_entry"

	// NotifyLogEntryReferenceAllOf1TypeNotifyLogEntryReference captures enum value "notify_log_entry_reference"
	NotifyLogEntryReferenceAllOf1TypeNotifyLogEntryReference string = "notify_log_entry_reference"
)

// prop value enum
func (m *NotifyLogEntryReferenceAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, notifyLogEntryReferenceAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NotifyLogEntryReferenceAllOf1) validateType(formats strfmt.Registry) error {

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
func (m *NotifyLogEntryReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotifyLogEntryReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res NotifyLogEntryReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
