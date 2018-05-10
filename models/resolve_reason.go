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

// ResolveReason resolve reason
// swagger:model ResolveReason
type ResolveReason struct {

	// incident
	Incident *IncidentReference `json:"incident,omitempty"`

	// The reason the incident was resolved. The only reason currently supported is merge.
	// Enum: [merge_resolve_reason]
	Type *string `json:"type,omitempty"`
}

// Validate validates this resolve reason
func (m *ResolveReason) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncident(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResolveReason) validateIncident(formats strfmt.Registry) error {

	if swag.IsZero(m.Incident) { // not required
		return nil
	}

	if m.Incident != nil {
		if err := m.Incident.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incident")
			}
			return err
		}
	}

	return nil
}

var resolveReasonTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["merge_resolve_reason"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resolveReasonTypeTypePropEnum = append(resolveReasonTypeTypePropEnum, v)
	}
}

const (

	// ResolveReasonTypeMergeResolveReason captures enum value "merge_resolve_reason"
	ResolveReasonTypeMergeResolveReason string = "merge_resolve_reason"
)

// prop value enum
func (m *ResolveReason) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resolveReasonTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResolveReason) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResolveReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResolveReason) UnmarshalBinary(b []byte) error {
	var res ResolveReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
