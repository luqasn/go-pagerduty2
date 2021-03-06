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

// AcknowledgeLogEntryReferenceAllOf1 acknowledge log entry reference all of1
// swagger:model acknowledgeLogEntryReferenceAllOf1
type AcknowledgeLogEntryReferenceAllOf1 struct {

	// type
	// Enum: [acknowledge_log_entry acknowledge_log_entry_reference]
	Type string `json:"type,omitempty"`
}

// Validate validates this acknowledge log entry reference all of1
func (m *AcknowledgeLogEntryReferenceAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var acknowledgeLogEntryReferenceAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["acknowledge_log_entry","acknowledge_log_entry_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		acknowledgeLogEntryReferenceAllOf1TypeTypePropEnum = append(acknowledgeLogEntryReferenceAllOf1TypeTypePropEnum, v)
	}
}

const (

	// AcknowledgeLogEntryReferenceAllOf1TypeAcknowledgeLogEntry captures enum value "acknowledge_log_entry"
	AcknowledgeLogEntryReferenceAllOf1TypeAcknowledgeLogEntry string = "acknowledge_log_entry"

	// AcknowledgeLogEntryReferenceAllOf1TypeAcknowledgeLogEntryReference captures enum value "acknowledge_log_entry_reference"
	AcknowledgeLogEntryReferenceAllOf1TypeAcknowledgeLogEntryReference string = "acknowledge_log_entry_reference"
)

// prop value enum
func (m *AcknowledgeLogEntryReferenceAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, acknowledgeLogEntryReferenceAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AcknowledgeLogEntryReferenceAllOf1) validateType(formats strfmt.Registry) error {

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
func (m *AcknowledgeLogEntryReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcknowledgeLogEntryReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res AcknowledgeLogEntryReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
