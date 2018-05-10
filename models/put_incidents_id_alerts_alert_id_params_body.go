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

// PutIncidentsIDAlertsAlertIDParamsBody put incidents Id alerts alert Id params body
// swagger:model putIncidentsIdAlertsAlertIdParamsBody
type PutIncidentsIDAlertsAlertIDParamsBody struct {

	// alert
	// Required: true
	Alert *Alert `json:"alert"`
}

// Validate validates this put incidents Id alerts alert Id params body
func (m *PutIncidentsIDAlertsAlertIDParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlert(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutIncidentsIDAlertsAlertIDParamsBody) validateAlert(formats strfmt.Registry) error {

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
func (m *PutIncidentsIDAlertsAlertIDParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutIncidentsIDAlertsAlertIDParamsBody) UnmarshalBinary(b []byte) error {
	var res PutIncidentsIDAlertsAlertIDParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}