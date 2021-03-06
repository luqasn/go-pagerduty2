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

// ListAlertsOKBodyAllOf1 list alerts o k body all of1
// swagger:model listAlertsOKBodyAllOf1
type ListAlertsOKBodyAllOf1 struct {

	// alerts
	// Required: true
	Alerts []*Alert `json:"alerts"`
}

// Validate validates this list alerts o k body all of1
func (m *ListAlertsOKBodyAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlerts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListAlertsOKBodyAllOf1) validateAlerts(formats strfmt.Registry) error {

	if err := validate.Required("alerts", "body", m.Alerts); err != nil {
		return err
	}

	for i := 0; i < len(m.Alerts); i++ {
		if swag.IsZero(m.Alerts[i]) { // not required
			continue
		}

		if m.Alerts[i] != nil {
			if err := m.Alerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListAlertsOKBodyAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAlertsOKBodyAllOf1) UnmarshalBinary(b []byte) error {
	var res ListAlertsOKBodyAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
