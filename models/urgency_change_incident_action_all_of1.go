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

// UrgencyChangeIncidentActionAllOf1 urgency change incident action all of1
// swagger:model urgencyChangeIncidentActionAllOf1
type UrgencyChangeIncidentActionAllOf1 struct {

	// Urgency that the incident will change to when the 'at' time is reached.
	// Enum: [high low]
	To string `json:"to,omitempty"`
}

// Validate validates this urgency change incident action all of1
func (m *UrgencyChangeIncidentActionAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var urgencyChangeIncidentActionAllOf1TypeToPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["high","low"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		urgencyChangeIncidentActionAllOf1TypeToPropEnum = append(urgencyChangeIncidentActionAllOf1TypeToPropEnum, v)
	}
}

const (

	// UrgencyChangeIncidentActionAllOf1ToHigh captures enum value "high"
	UrgencyChangeIncidentActionAllOf1ToHigh string = "high"

	// UrgencyChangeIncidentActionAllOf1ToLow captures enum value "low"
	UrgencyChangeIncidentActionAllOf1ToLow string = "low"
)

// prop value enum
func (m *UrgencyChangeIncidentActionAllOf1) validateToEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, urgencyChangeIncidentActionAllOf1TypeToPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UrgencyChangeIncidentActionAllOf1) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	// value enum
	if err := m.validateToEnum("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UrgencyChangeIncidentActionAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UrgencyChangeIncidentActionAllOf1) UnmarshalBinary(b []byte) error {
	var res UrgencyChangeIncidentActionAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
