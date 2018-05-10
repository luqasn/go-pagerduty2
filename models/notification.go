// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Notification notification
// swagger:model Notification
type Notification struct {

	// The address where the notification was sent. This will be null for notification type `push_notification`.
	// Read Only: true
	Address string `json:"address,omitempty"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The time at which the notification was sent
	// Read Only: true
	// Format: date-time
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`

	// The type of notification.
	// Read Only: true
	// Enum: [sms_notification email_notification phone_notification push_notification]
	Type string `json:"type,omitempty"`

	// The user the notification was sent to.
	// Read Only: true
	User *UserReference `json:"user,omitempty"`
}

// Validate validates this notification
func (m *Notification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Notification) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var notificationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sms_notification","email_notification","phone_notification","push_notification"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		notificationTypeTypePropEnum = append(notificationTypeTypePropEnum, v)
	}
}

const (

	// NotificationTypeSmsNotification captures enum value "sms_notification"
	NotificationTypeSmsNotification string = "sms_notification"

	// NotificationTypeEmailNotification captures enum value "email_notification"
	NotificationTypeEmailNotification string = "email_notification"

	// NotificationTypePhoneNotification captures enum value "phone_notification"
	NotificationTypePhoneNotification string = "phone_notification"

	// NotificationTypePushNotification captures enum value "push_notification"
	NotificationTypePushNotification string = "push_notification"
)

// prop value enum
func (m *Notification) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, notificationTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Notification) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Notification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notification) UnmarshalBinary(b []byte) error {
	var res Notification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}