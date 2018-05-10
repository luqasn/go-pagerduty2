// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Restriction restriction
// swagger:discriminator Restriction type
type Restriction interface {
	runtime.Validatable

	// The duration of the restriction in seconds.
	// Required: true
	DurationSeconds() *int64
	SetDurationSeconds(*int64)

	// Only required for use with a `weekly_restriction` restriction type. The first day of the weekly rotation schedule as [ISO 8601 day](https://en.wikipedia.org/wiki/ISO_week_date) (1 is Monday, etc.)
	// Maximum: 7
	// Minimum: 1
	StartDayOfWeek() int64
	SetStartDayOfWeek(int64)

	// The start time in HH:mm:ss format.
	// Required: true
	StartTimeOfDay() *string
	SetStartTimeOfDay(*string)

	// Specify the types of `restriction`.
	// Required: true
	// Enum: [daily_restriction weekly_restriction]
	Type() string
	SetType(string)
}

type restriction struct {
	durationSecondsField *int64

	startDayOfWeekField int64

	startTimeOfDayField *string

	typeField string
}

// DurationSeconds gets the duration seconds of this polymorphic type
func (m *restriction) DurationSeconds() *int64 {
	return m.durationSecondsField
}

// SetDurationSeconds sets the duration seconds of this polymorphic type
func (m *restriction) SetDurationSeconds(val *int64) {
	m.durationSecondsField = val
}

// StartDayOfWeek gets the start day of week of this polymorphic type
func (m *restriction) StartDayOfWeek() int64 {
	return m.startDayOfWeekField
}

// SetStartDayOfWeek sets the start day of week of this polymorphic type
func (m *restriction) SetStartDayOfWeek(val int64) {
	m.startDayOfWeekField = val
}

// StartTimeOfDay gets the start time of day of this polymorphic type
func (m *restriction) StartTimeOfDay() *string {
	return m.startTimeOfDayField
}

// SetStartTimeOfDay sets the start time of day of this polymorphic type
func (m *restriction) SetStartTimeOfDay(val *string) {
	m.startTimeOfDayField = val
}

// Type gets the type of this polymorphic type
func (m *restriction) Type() string {
	return "Restriction"
}

// SetType sets the type of this polymorphic type
func (m *restriction) SetType(val string) {

}

// UnmarshalRestrictionSlice unmarshals polymorphic slices of Restriction
func UnmarshalRestrictionSlice(reader io.Reader, consumer runtime.Consumer) ([]Restriction, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Restriction
	for _, element := range elements {
		obj, err := unmarshalRestriction(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalRestriction unmarshals polymorphic Restriction
func UnmarshalRestriction(reader io.Reader, consumer runtime.Consumer) (Restriction, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalRestriction(data, consumer)
}

func unmarshalRestriction(data []byte, consumer runtime.Consumer) (Restriction, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "Restriction":
		var result restriction
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "WeeklyRestriction":
		var result WeeklyRestriction
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)

}

// Validate validates this restriction
func (m *restriction) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *restriction) validateDurationSeconds(formats strfmt.Registry) error {

	if err := validate.Required("duration_seconds", "body", m.DurationSeconds()); err != nil {
		return err
	}

	return nil
}

func (m *restriction) validateStartDayOfWeek(formats strfmt.Registry) error {

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

func (m *restriction) validateStartTimeOfDay(formats strfmt.Registry) error {

	if err := validate.Required("start_time_of_day", "body", m.StartTimeOfDay()); err != nil {
		return err
	}

	return nil
}

var restrictionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["daily_restriction","weekly_restriction"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		restrictionTypeTypePropEnum = append(restrictionTypeTypePropEnum, v)
	}
}

const (

	// RestrictionTypeDailyRestriction captures enum value "daily_restriction"
	RestrictionTypeDailyRestriction string = "daily_restriction"

	// RestrictionTypeWeeklyRestriction captures enum value "weekly_restriction"
	RestrictionTypeWeeklyRestriction string = "weekly_restriction"
)

// prop value enum
func (m *restriction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, restrictionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}