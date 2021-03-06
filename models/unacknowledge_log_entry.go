// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UnacknowledgeLogEntry unacknowledge log entry
// swagger:model UnacknowledgeLogEntry
type UnacknowledgeLogEntry struct {
	LogEntry

	UnacknowledgeLogEntryReference
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UnacknowledgeLogEntry) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 LogEntry
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.LogEntry = aO0

	// AO1
	var aO1 UnacknowledgeLogEntryReference
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.UnacknowledgeLogEntryReference = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UnacknowledgeLogEntry) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.LogEntry)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.UnacknowledgeLogEntryReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this unacknowledge log entry
func (m *UnacknowledgeLogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with LogEntry
	if err := m.LogEntry.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with UnacknowledgeLogEntryReference
	if err := m.UnacknowledgeLogEntryReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UnacknowledgeLogEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnacknowledgeLogEntry) UnmarshalBinary(b []byte) error {
	var res UnacknowledgeLogEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
