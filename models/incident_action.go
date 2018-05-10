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
	"github.com/go-openapi/validate"
)

// IncidentAction An incident action is a pending change to an incident that will automatically happen at some future time.
// swagger:discriminator IncidentAction type
type IncidentAction interface {
	runtime.Validatable

	// at
	// Required: true
	// Format: date-time
	At() *strfmt.DateTime
	SetAt(*strfmt.DateTime)

	// type
	// Required: true
	// Enum: [unacknowledge escalate resolve urgency_change]
	Type() string
	SetType(string)
}

type incidentAction struct {
	atField *strfmt.DateTime

	typeField string
}

// At gets the at of this polymorphic type
func (m *incidentAction) At() *strfmt.DateTime {
	return m.atField
}

// SetAt sets the at of this polymorphic type
func (m *incidentAction) SetAt(val *strfmt.DateTime) {
	m.atField = val
}

// Type gets the type of this polymorphic type
func (m *incidentAction) Type() string {
	return "IncidentAction"
}

// SetType sets the type of this polymorphic type
func (m *incidentAction) SetType(val string) {

}

// UnmarshalIncidentActionSlice unmarshals polymorphic slices of IncidentAction
func UnmarshalIncidentActionSlice(reader io.Reader, consumer runtime.Consumer) ([]IncidentAction, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []IncidentAction
	for _, element := range elements {
		obj, err := unmarshalIncidentAction(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalIncidentAction unmarshals polymorphic IncidentAction
func UnmarshalIncidentAction(reader io.Reader, consumer runtime.Consumer) (IncidentAction, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalIncidentAction(data, consumer)
}

func unmarshalIncidentAction(data []byte, consumer runtime.Consumer) (IncidentAction, error) {
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
	case "IncidentAction":
		var result incidentAction
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "UrgencyChangeIncidentAction":
		var result UrgencyChangeIncidentAction
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)

}

// Validate validates this incident action
func (m *incidentAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *incidentAction) validateAt(formats strfmt.Registry) error {

	if err := validate.Required("at", "body", m.At()); err != nil {
		return err
	}

	if err := validate.FormatOf("at", "body", "date-time", m.At().String(), formats); err != nil {
		return err
	}

	return nil
}

var incidentActionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unacknowledge","escalate","resolve","urgency_change"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		incidentActionTypeTypePropEnum = append(incidentActionTypeTypePropEnum, v)
	}
}

const (

	// IncidentActionTypeUnacknowledge captures enum value "unacknowledge"
	IncidentActionTypeUnacknowledge string = "unacknowledge"

	// IncidentActionTypeEscalate captures enum value "escalate"
	IncidentActionTypeEscalate string = "escalate"

	// IncidentActionTypeResolve captures enum value "resolve"
	IncidentActionTypeResolve string = "resolve"

	// IncidentActionTypeUrgencyChange captures enum value "urgency_change"
	IncidentActionTypeUrgencyChange string = "urgency_change"
)

// prop value enum
func (m *incidentAction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, incidentActionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}
