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

// ScheduledActionAt Represents when scheduled action will occur.
// swagger:model scheduledActionAt
type ScheduledActionAt struct {

	// Designates either the start or the end of support hours.
	// Required: true
	// Enum: [support_hours_start support_hours_end]
	Name *string `json:"name"`

	// Must be set to named_time.
	// Required: true
	// Enum: [named_time]
	Type *string `json:"type"`
}

// Validate validates this scheduled action at
func (m *ScheduledActionAt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
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

var scheduledActionAtTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["support_hours_start","support_hours_end"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduledActionAtTypeNamePropEnum = append(scheduledActionAtTypeNamePropEnum, v)
	}
}

const (

	// ScheduledActionAtNameSupportHoursStart captures enum value "support_hours_start"
	ScheduledActionAtNameSupportHoursStart string = "support_hours_start"

	// ScheduledActionAtNameSupportHoursEnd captures enum value "support_hours_end"
	ScheduledActionAtNameSupportHoursEnd string = "support_hours_end"
)

// prop value enum
func (m *ScheduledActionAt) validateNameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, scheduledActionAtTypeNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ScheduledActionAt) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	// value enum
	if err := m.validateNameEnum("name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

var scheduledActionAtTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["named_time"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduledActionAtTypeTypePropEnum = append(scheduledActionAtTypeTypePropEnum, v)
	}
}

const (

	// ScheduledActionAtTypeNamedTime captures enum value "named_time"
	ScheduledActionAtTypeNamedTime string = "named_time"
)

// prop value enum
func (m *ScheduledActionAt) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, scheduledActionAtTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ScheduledActionAt) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduledActionAt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduledActionAt) UnmarshalBinary(b []byte) error {
	var res ScheduledActionAt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
