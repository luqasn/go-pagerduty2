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

// LinkContextAllOf1 link context all of1
// swagger:model linkContextAllOf1
type LinkContextAllOf1 struct {

	// The link being attached to the incident
	// Required: true
	Href *string `json:"href"`

	// The text that will be displayed as the link.
	Text string `json:"text,omitempty"`
}

// Validate validates this link context all of1
func (m *LinkContextAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkContextAllOf1) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinkContextAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkContextAllOf1) UnmarshalBinary(b []byte) error {
	var res LinkContextAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}