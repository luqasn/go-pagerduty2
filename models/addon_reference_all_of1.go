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

// AddonReferenceAllOf1 addon reference all of1
// swagger:model addonReferenceAllOf1
type AddonReferenceAllOf1 struct {

	// type
	// Enum: [full_page_addon full_page_addon_reference incident_show_addon incident_show_addon_reference]
	Type string `json:"type,omitempty"`
}

// Validate validates this addon reference all of1
func (m *AddonReferenceAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addonReferenceAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["full_page_addon","full_page_addon_reference","incident_show_addon","incident_show_addon_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addonReferenceAllOf1TypeTypePropEnum = append(addonReferenceAllOf1TypeTypePropEnum, v)
	}
}

const (

	// AddonReferenceAllOf1TypeFullPageAddon captures enum value "full_page_addon"
	AddonReferenceAllOf1TypeFullPageAddon string = "full_page_addon"

	// AddonReferenceAllOf1TypeFullPageAddonReference captures enum value "full_page_addon_reference"
	AddonReferenceAllOf1TypeFullPageAddonReference string = "full_page_addon_reference"

	// AddonReferenceAllOf1TypeIncidentShowAddon captures enum value "incident_show_addon"
	AddonReferenceAllOf1TypeIncidentShowAddon string = "incident_show_addon"

	// AddonReferenceAllOf1TypeIncidentShowAddonReference captures enum value "incident_show_addon_reference"
	AddonReferenceAllOf1TypeIncidentShowAddonReference string = "incident_show_addon_reference"
)

// prop value enum
func (m *AddonReferenceAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addonReferenceAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AddonReferenceAllOf1) validateType(formats strfmt.Registry) error {

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
func (m *AddonReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddonReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res AddonReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}