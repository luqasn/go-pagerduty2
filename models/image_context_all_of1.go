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

// ImageContextAllOf1 image context all of1
// swagger:model imageContextAllOf1
type ImageContextAllOf1 struct {

	// Optional alt text for the image
	Alt string `json:"alt,omitempty"`

	// Optional link for the image
	Href string `json:"href,omitempty"`

	// The source of the image being attached to the incident
	// Required: true
	Src *string `json:"src"`
}

// Validate validates this image context all of1
func (m *ImageContextAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSrc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageContextAllOf1) validateSrc(formats strfmt.Registry) error {

	if err := validate.Required("src", "body", m.Src); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageContextAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageContextAllOf1) UnmarshalBinary(b []byte) error {
	var res ImageContextAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
