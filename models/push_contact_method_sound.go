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

// PushContactMethodSound push contact method sound
// swagger:model PushContactMethodSound
type PushContactMethodSound struct {

	// The sound file name.
	File string `json:"file,omitempty"`

	// The type of sound.
	// Enum: [alert_high_urgency alert_low_urgency]
	Type string `json:"type,omitempty"`
}

// Validate validates this push contact method sound
func (m *PushContactMethodSound) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pushContactMethodSoundTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alert_high_urgency","alert_low_urgency"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pushContactMethodSoundTypeTypePropEnum = append(pushContactMethodSoundTypeTypePropEnum, v)
	}
}

const (

	// PushContactMethodSoundTypeAlertHighUrgency captures enum value "alert_high_urgency"
	PushContactMethodSoundTypeAlertHighUrgency string = "alert_high_urgency"

	// PushContactMethodSoundTypeAlertLowUrgency captures enum value "alert_low_urgency"
	PushContactMethodSoundTypeAlertLowUrgency string = "alert_low_urgency"
)

// prop value enum
func (m *PushContactMethodSound) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pushContactMethodSoundTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PushContactMethodSound) validateType(formats strfmt.Registry) error {

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
func (m *PushContactMethodSound) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PushContactMethodSound) UnmarshalBinary(b []byte) error {
	var res PushContactMethodSound
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}