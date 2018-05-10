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

// ScheduleReferenceAllOf1 schedule reference all of1
// swagger:model scheduleReferenceAllOf1
type ScheduleReferenceAllOf1 struct {

	// type
	// Enum: [schedule schedule_reference]
	Type string `json:"type,omitempty"`
}

// Validate validates this schedule reference all of1
func (m *ScheduleReferenceAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var scheduleReferenceAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["schedule","schedule_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduleReferenceAllOf1TypeTypePropEnum = append(scheduleReferenceAllOf1TypeTypePropEnum, v)
	}
}

const (

	// ScheduleReferenceAllOf1TypeSchedule captures enum value "schedule"
	ScheduleReferenceAllOf1TypeSchedule string = "schedule"

	// ScheduleReferenceAllOf1TypeScheduleReference captures enum value "schedule_reference"
	ScheduleReferenceAllOf1TypeScheduleReference string = "schedule_reference"
)

// prop value enum
func (m *ScheduleReferenceAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, scheduleReferenceAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ScheduleReferenceAllOf1) validateType(formats strfmt.Registry) error {

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
func (m *ScheduleReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res ScheduleReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}