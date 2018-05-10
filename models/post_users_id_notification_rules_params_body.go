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

// PostUsersIDNotificationRulesParamsBody post users Id notification rules params body
// swagger:model postUsersIdNotificationRulesParamsBody
type PostUsersIDNotificationRulesParamsBody struct {

	// notification rule
	// Required: true
	NotificationRule *NotificationRule `json:"notification_rule"`
}

// Validate validates this post users Id notification rules params body
func (m *PostUsersIDNotificationRulesParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotificationRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostUsersIDNotificationRulesParamsBody) validateNotificationRule(formats strfmt.Registry) error {

	if err := validate.Required("notification_rule", "body", m.NotificationRule); err != nil {
		return err
	}

	if m.NotificationRule != nil {
		if err := m.NotificationRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notification_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostUsersIDNotificationRulesParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostUsersIDNotificationRulesParamsBody) UnmarshalBinary(b []byte) error {
	var res PostUsersIDNotificationRulesParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
