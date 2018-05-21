// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateTeamCreatedBody create team created body
// swagger:model createTeamCreatedBody
type CreateTeamCreatedBody struct {

	// team
	// Required: true
	Team *Team `json:"team"`
}

// Validate validates this create team created body
func (m *CreateTeamCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTeamCreatedBody) validateTeam(formats strfmt.Registry) error {

	if err := validate.Required("team", "body", m.Team); err != nil {
		return err
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateTeamCreatedBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTeamCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateTeamCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
