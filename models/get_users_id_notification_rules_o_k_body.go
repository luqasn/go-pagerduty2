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

// GetUsersIDNotificationRulesOKBody get users Id notification rules o k body
// swagger:model getUsersIdNotificationRulesOKBody
type GetUsersIDNotificationRulesOKBody struct {

	// notification rules
	// Required: true
	NotificationRules []*NotificationRule `json:"notification_rules"`
}

// Validate validates this get users Id notification rules o k body
func (m *GetUsersIDNotificationRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotificationRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUsersIDNotificationRulesOKBody) validateNotificationRules(formats strfmt.Registry) error {

	if err := validate.Required("notification_rules", "body", m.NotificationRules); err != nil {
		return err
	}

	for i := 0; i < len(m.NotificationRules); i++ {
		if swag.IsZero(m.NotificationRules[i]) { // not required
			continue
		}

		if m.NotificationRules[i] != nil {
			if err := m.NotificationRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notification_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUsersIDNotificationRulesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUsersIDNotificationRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetUsersIDNotificationRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}