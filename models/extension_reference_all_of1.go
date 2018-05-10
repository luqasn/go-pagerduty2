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

// ExtensionReferenceAllOf1 extension reference all of1
// swagger:model extensionReferenceAllOf1
type ExtensionReferenceAllOf1 struct {

	// type
	// Enum: [extension extension_reference]
	Type string `json:"type,omitempty"`
}

// Validate validates this extension reference all of1
func (m *ExtensionReferenceAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var extensionReferenceAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extension","extension_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		extensionReferenceAllOf1TypeTypePropEnum = append(extensionReferenceAllOf1TypeTypePropEnum, v)
	}
}

const (

	// ExtensionReferenceAllOf1TypeExtension captures enum value "extension"
	ExtensionReferenceAllOf1TypeExtension string = "extension"

	// ExtensionReferenceAllOf1TypeExtensionReference captures enum value "extension_reference"
	ExtensionReferenceAllOf1TypeExtensionReference string = "extension_reference"
)

// prop value enum
func (m *ExtensionReferenceAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, extensionReferenceAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ExtensionReferenceAllOf1) validateType(formats strfmt.Registry) error {

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
func (m *ExtensionReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res ExtensionReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}