// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetIncidentsOKBodyAllOf1 get incidents o k body all of1
// swagger:model getIncidentsOKBodyAllOf1
type GetIncidentsOKBodyAllOf1 struct {

	// incidents
	// Required: true
	Incidents []*Incident `json:"incidents"`
}

// Validate validates this get incidents o k body all of1
func (m *GetIncidentsOKBodyAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncidents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetIncidentsOKBodyAllOf1) validateIncidents(formats strfmt.Registry) error {

	if err := validate.Required("incidents", "body", m.Incidents); err != nil {
		return err
	}

	for i := 0; i < len(m.Incidents); i++ {
		if swag.IsZero(m.Incidents[i]) { // not required
			continue
		}

		if m.Incidents[i] != nil {
			if err := m.Incidents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetIncidentsOKBodyAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetIncidentsOKBodyAllOf1) UnmarshalBinary(b []byte) error {
	var res GetIncidentsOKBodyAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
