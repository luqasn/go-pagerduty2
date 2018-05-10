// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PushContactMethodAllOf1 push contact method all of1
// swagger:model pushContactMethodAllOf1
type PushContactMethodAllOf1 struct {

	// If true, this phone has been blacklisted by PagerDuty and no messages will be sent to it.
	// Read Only: true
	Blacklisted *bool `json:"blacklisted,omitempty"`

	// Time at which the contact method was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The type of device.
	// Required: true
	// Read Only: true
	// Enum: [android ios]
	DeviceType string `json:"device_type"`

	// sounds
	Sounds []*PushContactMethodSound `json:"sounds"`

	// type
	// Enum: [push_notification_contact_method]
	Type string `json:"type,omitempty"`
}

// Validate validates this push contact method all of1
func (m *PushContactMethodAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSounds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PushContactMethodAllOf1) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var pushContactMethodAllOf1TypeDeviceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["android","ios"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pushContactMethodAllOf1TypeDeviceTypePropEnum = append(pushContactMethodAllOf1TypeDeviceTypePropEnum, v)
	}
}

const (

	// PushContactMethodAllOf1DeviceTypeAndroid captures enum value "android"
	PushContactMethodAllOf1DeviceTypeAndroid string = "android"

	// PushContactMethodAllOf1DeviceTypeIos captures enum value "ios"
	PushContactMethodAllOf1DeviceTypeIos string = "ios"
)

// prop value enum
func (m *PushContactMethodAllOf1) validateDeviceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pushContactMethodAllOf1TypeDeviceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PushContactMethodAllOf1) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.RequiredString("device_type", "body", string(m.DeviceType)); err != nil {
		return err
	}

	// value enum
	if err := m.validateDeviceTypeEnum("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *PushContactMethodAllOf1) validateSounds(formats strfmt.Registry) error {

	if swag.IsZero(m.Sounds) { // not required
		return nil
	}

	for i := 0; i < len(m.Sounds); i++ {
		if swag.IsZero(m.Sounds[i]) { // not required
			continue
		}

		if m.Sounds[i] != nil {
			if err := m.Sounds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sounds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var pushContactMethodAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["push_notification_contact_method"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pushContactMethodAllOf1TypeTypePropEnum = append(pushContactMethodAllOf1TypeTypePropEnum, v)
	}
}

const (

	// PushContactMethodAllOf1TypePushNotificationContactMethod captures enum value "push_notification_contact_method"
	PushContactMethodAllOf1TypePushNotificationContactMethod string = "push_notification_contact_method"
)

// prop value enum
func (m *PushContactMethodAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pushContactMethodAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PushContactMethodAllOf1) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PushContactMethodAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PushContactMethodAllOf1) UnmarshalBinary(b []byte) error {
	var res PushContactMethodAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
