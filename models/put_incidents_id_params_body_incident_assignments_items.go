// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PutIncidentsIDParamsBodyIncidentAssignmentsItems put incidents Id params body incident assignments items
// swagger:model putIncidentsIdParamsBodyIncidentAssignmentsItems
type PutIncidentsIDParamsBodyIncidentAssignmentsItems struct {

	// User that was assigned.
	Assignee *UserReference `json:"assignee,omitempty"`
}

// Validate validates this put incidents Id params body incident assignments items
func (m *PutIncidentsIDParamsBodyIncidentAssignmentsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignee(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutIncidentsIDParamsBodyIncidentAssignmentsItems) validateAssignee(formats strfmt.Registry) error {

	if swag.IsZero(m.Assignee) { // not required
		return nil
	}

	if m.Assignee != nil {
		if err := m.Assignee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignee")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutIncidentsIDParamsBodyIncidentAssignmentsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutIncidentsIDParamsBodyIncidentAssignmentsItems) UnmarshalBinary(b []byte) error {
	var res PutIncidentsIDParamsBodyIncidentAssignmentsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}