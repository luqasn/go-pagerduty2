// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WeeklyRestriction weekly restriction
// swagger:model WeeklyRestriction
type WeeklyRestriction struct {
	durationSecondsField *int64

	startDayOfWeekField int64

	startTimeOfDayField *string

	WeeklyRestrictionAllOf1
}

// DurationSeconds gets the duration seconds of this subtype
func (m *WeeklyRestriction) DurationSeconds() *int64 {
	return m.durationSecondsField
}

// SetDurationSeconds sets the duration seconds of this subtype
func (m *WeeklyRestriction) SetDurationSeconds(val *int64) {
	m.durationSecondsField = val
}

// StartDayOfWeek gets the start day of week of this subtype
func (m *WeeklyRestriction) StartDayOfWeek() int64 {
	return m.startDayOfWeekField
}

// SetStartDayOfWeek sets the start day of week of this subtype
func (m *WeeklyRestriction) SetStartDayOfWeek(val int64) {
	m.startDayOfWeekField = val
}

// StartTimeOfDay gets the start time of day of this subtype
func (m *WeeklyRestriction) StartTimeOfDay() *string {
	return m.startTimeOfDayField
}

// SetStartTimeOfDay sets the start time of day of this subtype
func (m *WeeklyRestriction) SetStartTimeOfDay(val *string) {
	m.startTimeOfDayField = val
}

// Type gets the type of this subtype
func (m *WeeklyRestriction) Type() string {
	return "WeeklyRestriction"
}

// SetType sets the type of this subtype
func (m *WeeklyRestriction) SetType(val string) {

}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *WeeklyRestriction) UnmarshalJSON(raw []byte) error {
	var data struct {
		WeeklyRestrictionAllOf1
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		DurationSeconds *int64 `json:"duration_seconds"`

		StartDayOfWeek int64 `json:"start_day_of_week,omitempty"`

		StartTimeOfDay *string `json:"start_time_of_day"`

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result WeeklyRestriction

	result.durationSecondsField = base.DurationSeconds

	result.startDayOfWeekField = base.StartDayOfWeek

	result.startTimeOfDayField = base.StartTimeOfDay

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.WeeklyRestrictionAllOf1 = data.WeeklyRestrictionAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m WeeklyRestriction) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		WeeklyRestrictionAllOf1
	}{

		WeeklyRestrictionAllOf1: m.WeeklyRestrictionAllOf1,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		DurationSeconds *int64 `json:"duration_seconds"`

		StartDayOfWeek int64 `json:"start_day_of_week,omitempty"`

		StartTimeOfDay *string `json:"start_time_of_day"`

		Type string `json:"type"`
	}{

		DurationSeconds: m.DurationSeconds(),

		StartDayOfWeek: m.StartDayOfWeek(),

		StartTimeOfDay: m.StartTimeOfDay(),

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this weekly restriction
func (m *WeeklyRestriction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDurationSeconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDayOfWeek(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTimeOfDay(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with WeeklyRestrictionAllOf1
	if err := m.WeeklyRestrictionAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WeeklyRestriction) validateDurationSeconds(formats strfmt.Registry) error {

	if err := validate.Required("duration_seconds", "body", m.DurationSeconds()); err != nil {
		return err
	}

	return nil
}

func (m *WeeklyRestriction) validateStartDayOfWeek(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDayOfWeek()) { // not required
		return nil
	}

	if err := validate.MinimumInt("start_day_of_week", "body", int64(m.StartDayOfWeek()), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("start_day_of_week", "body", int64(m.StartDayOfWeek()), 7, false); err != nil {
		return err
	}

	return nil
}

func (m *WeeklyRestriction) validateStartTimeOfDay(formats strfmt.Registry) error {

	if err := validate.Required("start_time_of_day", "body", m.StartTimeOfDay()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WeeklyRestriction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WeeklyRestriction) UnmarshalBinary(b []byte) error {
	var res WeeklyRestriction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}