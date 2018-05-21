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

// ListNotificationsOKBodyAllOf1 list notifications o k body all of1
// swagger:model listNotificationsOKBodyAllOf1
type ListNotificationsOKBodyAllOf1 struct {

	// notifications
	// Required: true
	Notifications []*Notification `json:"notifications"`
}

// Validate validates this list notifications o k body all of1
func (m *ListNotificationsOKBodyAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListNotificationsOKBodyAllOf1) validateNotifications(formats strfmt.Registry) error {

	if err := validate.Required("notifications", "body", m.Notifications); err != nil {
		return err
	}

	for i := 0; i < len(m.Notifications); i++ {
		if swag.IsZero(m.Notifications[i]) { // not required
			continue
		}

		if m.Notifications[i] != nil {
			if err := m.Notifications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListNotificationsOKBodyAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListNotificationsOKBodyAllOf1) UnmarshalBinary(b []byte) error {
	var res ListNotificationsOKBodyAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}