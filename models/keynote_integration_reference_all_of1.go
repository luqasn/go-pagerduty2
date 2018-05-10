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

// KeynoteIntegrationReferenceAllOf1 keynote integration reference all of1
// swagger:model keynoteIntegrationReferenceAllOf1
type KeynoteIntegrationReferenceAllOf1 struct {

	// type
	// Enum: [keynote_inbound_integration keynote_inbound_integration_reference]
	Type string `json:"type,omitempty"`
}

// Validate validates this keynote integration reference all of1
func (m *KeynoteIntegrationReferenceAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var keynoteIntegrationReferenceAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["keynote_inbound_integration","keynote_inbound_integration_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		keynoteIntegrationReferenceAllOf1TypeTypePropEnum = append(keynoteIntegrationReferenceAllOf1TypeTypePropEnum, v)
	}
}

const (

	// KeynoteIntegrationReferenceAllOf1TypeKeynoteInboundIntegration captures enum value "keynote_inbound_integration"
	KeynoteIntegrationReferenceAllOf1TypeKeynoteInboundIntegration string = "keynote_inbound_integration"

	// KeynoteIntegrationReferenceAllOf1TypeKeynoteInboundIntegrationReference captures enum value "keynote_inbound_integration_reference"
	KeynoteIntegrationReferenceAllOf1TypeKeynoteInboundIntegrationReference string = "keynote_inbound_integration_reference"
)

// prop value enum
func (m *KeynoteIntegrationReferenceAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, keynoteIntegrationReferenceAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *KeynoteIntegrationReferenceAllOf1) validateType(formats strfmt.Registry) error {

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
func (m *KeynoteIntegrationReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeynoteIntegrationReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res KeynoteIntegrationReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}