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

// PutIncidentsIDAlertsAlertIDOKBody put incidents Id alerts alert Id o k body
// swagger:model putIncidentsIdAlertsAlertIdOKBody
type PutIncidentsIDAlertsAlertIDOKBody struct {

	// alert
	// Required: true
	Alert *Alert `json:"alert"`
}

// Validate validates this put incidents Id alerts alert Id o k body
func (m *PutIncidentsIDAlertsAlertIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlert(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutIncidentsIDAlertsAlertIDOKBody) validateAlert(formats strfmt.Registry) error {

	if err := validate.Required("alert", "body", m.Alert); err != nil {
		return err
	}

	if m.Alert != nil {
		if err := m.Alert.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alert")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutIncidentsIDAlertsAlertIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutIncidentsIDAlertsAlertIDOKBody) UnmarshalBinary(b []byte) error {
	var res PutIncidentsIDAlertsAlertIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
