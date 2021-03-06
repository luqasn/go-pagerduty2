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

// CreateServiceCreatedBody create service created body
// swagger:model createServiceCreatedBody
type CreateServiceCreatedBody struct {

	// service
	// Required: true
	Service *Service `json:"service"`
}

// Validate validates this create service created body
func (m *CreateServiceCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateServiceCreatedBody) validateService(formats strfmt.Registry) error {

	if err := validate.Required("service", "body", m.Service); err != nil {
		return err
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateServiceCreatedBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateServiceCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateServiceCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
