// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExtensionSchemaAllOf1 extension schema all of1
// swagger:model extensionSchemaAllOf1
type ExtensionSchemaAllOf1 struct {

	// The long description for the Extension
	Description string `json:"description,omitempty"`

	// A link to the extension's support guide
	// Read Only: true
	GuideURL string `json:"guide_url,omitempty"`

	// A small logo, 18-by-18 pixels.
	// Read Only: true
	IconURL string `json:"icon_url,omitempty"`

	// Machine friendly display label
	// Read Only: true
	Key string `json:"key,omitempty"`

	// Human friendly display label
	// Read Only: true
	Label string `json:"label,omitempty"`

	// A large logo, 75 pixels high and no more than 300 pixels wide.
	// Read Only: true
	LogoURL string `json:"logo_url,omitempty"`

	// The types of PagerDuty incident events that will activate this Extension
	// Unique: true
	SendTypes []string `json:"send_types"`

	// The url that the webhook payload will be sent to for this Extension.
	// Read Only: true
	URL string `json:"url,omitempty"`
}

// Validate validates this extension schema all of1
func (m *ExtensionSchemaAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSendTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var extensionSchemaAllOf1SendTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["trigger","acknowledge","resolve","delegate","escalate","unacknowledge","assign","custom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		extensionSchemaAllOf1SendTypesItemsEnum = append(extensionSchemaAllOf1SendTypesItemsEnum, v)
	}
}

func (m *ExtensionSchemaAllOf1) validateSendTypesItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, extensionSchemaAllOf1SendTypesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *ExtensionSchemaAllOf1) validateSendTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.SendTypes) { // not required
		return nil
	}

	if err := validate.UniqueItems("send_types", "body", m.SendTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.SendTypes); i++ {

		// value enum
		if err := m.validateSendTypesItemsEnum("send_types"+"."+strconv.Itoa(i), "body", m.SendTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionSchemaAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionSchemaAllOf1) UnmarshalBinary(b []byte) error {
	var res ExtensionSchemaAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
