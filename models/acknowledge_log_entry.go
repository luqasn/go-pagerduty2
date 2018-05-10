// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AcknowledgeLogEntry acknowledge log entry
// swagger:model AcknowledgeLogEntry
type AcknowledgeLogEntry struct {
	LogEntry

	AcknowledgeLogEntryReference

	AcknowledgeLogEntryAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AcknowledgeLogEntry) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 LogEntry
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.LogEntry = aO0

	// AO1
	var aO1 AcknowledgeLogEntryReference
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.AcknowledgeLogEntryReference = aO1

	// AO2
	var aO2 AcknowledgeLogEntryAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.AcknowledgeLogEntryAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AcknowledgeLogEntry) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.LogEntry)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.AcknowledgeLogEntryReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.AcknowledgeLogEntryAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this acknowledge log entry
func (m *AcknowledgeLogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with LogEntry
	if err := m.LogEntry.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with AcknowledgeLogEntryReference
	if err := m.AcknowledgeLogEntryReference.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with AcknowledgeLogEntryAllOf2
	if err := m.AcknowledgeLogEntryAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AcknowledgeLogEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcknowledgeLogEntry) UnmarshalBinary(b []byte) error {
	var res AcknowledgeLogEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
