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

// PostServicesIDIntegrationsCreatedBody post services Id integrations created body
// swagger:model postServicesIdIntegrationsCreatedBody
type PostServicesIDIntegrationsCreatedBody struct {

	// integration
	// Required: true
	Integration *Integration `json:"integration"`
}

// Validate validates this post services Id integrations created body
func (m *PostServicesIDIntegrationsCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostServicesIDIntegrationsCreatedBody) validateIntegration(formats strfmt.Registry) error {

	if err := validate.Required("integration", "body", m.Integration); err != nil {
		return err
	}

	if m.Integration != nil {
		if err := m.Integration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostServicesIDIntegrationsCreatedBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostServicesIDIntegrationsCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostServicesIDIntegrationsCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
