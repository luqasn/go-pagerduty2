// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ResolveLogEntry resolve log entry
// swagger:model ResolveLogEntry
type ResolveLogEntry struct {
	LogEntry

	ResolveLogEntryReference
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ResolveLogEntry) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 LogEntry
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.LogEntry = aO0

	// AO1
	var aO1 ResolveLogEntryReference
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ResolveLogEntryReference = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ResolveLogEntry) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.LogEntry)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ResolveLogEntryReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this resolve log entry
func (m *ResolveLogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with LogEntry
	if err := m.LogEntry.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ResolveLogEntryReference
	if err := m.ResolveLogEntryReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ResolveLogEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResolveLogEntry) UnmarshalBinary(b []byte) error {
	var res ResolveLogEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}