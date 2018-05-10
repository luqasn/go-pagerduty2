// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NotifyLogEntry notify log entry
// swagger:model NotifyLogEntry
type NotifyLogEntry struct {
	LogEntry

	NotifyLogEntryReference

	NotifyLogEntryAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NotifyLogEntry) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 LogEntry
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.LogEntry = aO0

	// AO1
	var aO1 NotifyLogEntryReference
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.NotifyLogEntryReference = aO1

	// AO2
	var aO2 NotifyLogEntryAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.NotifyLogEntryAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NotifyLogEntry) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.LogEntry)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.NotifyLogEntryReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.NotifyLogEntryAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this notify log entry
func (m *NotifyLogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with LogEntry
	if err := m.LogEntry.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with NotifyLogEntryReference
	if err := m.NotifyLogEntryReference.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with NotifyLogEntryAllOf2
	if err := m.NotifyLogEntryAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *NotifyLogEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotifyLogEntry) UnmarshalBinary(b []byte) error {
	var res NotifyLogEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}