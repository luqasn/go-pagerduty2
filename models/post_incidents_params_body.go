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

// PostIncidentsParamsBody post incidents params body
// swagger:model postIncidentsParamsBody
type PostIncidentsParamsBody struct {

	// incident
	// Required: true
	Incident *PostIncidentsParamsBodyIncident `json:"incident"`
}

// Validate validates this post incidents params body
func (m *PostIncidentsParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncident(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostIncidentsParamsBody) validateIncident(formats strfmt.Registry) error {

	if err := validate.Required("incident", "body", m.Incident); err != nil {
		return err
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

// MarshalBinary interface implementation
func (m *PostIncidentsParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostIncidentsParamsBody) UnmarshalBinary(b []byte) error {
	var res PostIncidentsParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
