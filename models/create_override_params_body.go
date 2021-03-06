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

// CreateOverrideParamsBody create override params body
// swagger:model createOverrideParamsBody
type CreateOverrideParamsBody struct {

	// override
	// Required: true
	Override *Override `json:"override"`
}

// Validate validates this create override params body
func (m *CreateOverrideParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOverride(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateOverrideParamsBody) validateOverride(formats strfmt.Registry) error {

	if err := validate.Required("override", "body", m.Override); err != nil {
		return err
	}

	if m.Override != nil {
		if err := m.Override.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("override")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateOverrideParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateOverrideParamsBody) UnmarshalBinary(b []byte) error {
	var res CreateOverrideParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
