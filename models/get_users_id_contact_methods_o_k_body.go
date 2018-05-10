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

// GetUsersIDContactMethodsOKBody get users Id contact methods o k body
// swagger:model getUsersIdContactMethodsOKBody
type GetUsersIDContactMethodsOKBody struct {

	// contact methods
	// Required: true
	ContactMethods []*ContactMethod `json:"contact_methods"`
}

// Validate validates this get users Id contact methods o k body
func (m *GetUsersIDContactMethodsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactMethods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUsersIDContactMethodsOKBody) validateContactMethods(formats strfmt.Registry) error {

	if err := validate.Required("contact_methods", "body", m.ContactMethods); err != nil {
		return err
	}

	for i := 0; i < len(m.ContactMethods); i++ {
		if swag.IsZero(m.ContactMethods[i]) { // not required
			continue
		}

		if m.ContactMethods[i] != nil {
			if err := m.ContactMethods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contact_methods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUsersIDContactMethodsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUsersIDContactMethodsOKBody) UnmarshalBinary(b []byte) error {
	var res GetUsersIDContactMethodsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
