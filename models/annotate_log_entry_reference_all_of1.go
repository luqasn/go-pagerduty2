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

// AnnotateLogEntryReferenceAllOf1 annotate log entry reference all of1
// swagger:model annotateLogEntryReferenceAllOf1
type AnnotateLogEntryReferenceAllOf1 struct {

	// type
	// Enum: [annotate_log_entry annotate_log_entry_reference]
	Type string `json:"type,omitempty"`
}

// Validate validates this annotate log entry reference all of1
func (m *AnnotateLogEntryReferenceAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var annotateLogEntryReferenceAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["annotate_log_entry","annotate_log_entry_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		annotateLogEntryReferenceAllOf1TypeTypePropEnum = append(annotateLogEntryReferenceAllOf1TypeTypePropEnum, v)
	}
}

const (

	// AnnotateLogEntryReferenceAllOf1TypeAnnotateLogEntry captures enum value "annotate_log_entry"
	AnnotateLogEntryReferenceAllOf1TypeAnnotateLogEntry string = "annotate_log_entry"

	// AnnotateLogEntryReferenceAllOf1TypeAnnotateLogEntryReference captures enum value "annotate_log_entry_reference"
	AnnotateLogEntryReferenceAllOf1TypeAnnotateLogEntryReference string = "annotate_log_entry_reference"
)

// prop value enum
func (m *AnnotateLogEntryReferenceAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, annotateLogEntryReferenceAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AnnotateLogEntryReferenceAllOf1) validateType(formats strfmt.Registry) error {

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
func (m *AnnotateLogEntryReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnnotateLogEntryReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res AnnotateLogEntryReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}