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
)

// EscalationPolicyReference escalation policy reference
// swagger:model EscalationPolicyReference
type EscalationPolicyReference struct {
	htmlUrlField string

	idField string

	selfField string

	summaryField string

	EscalationPolicyReferenceAllOf1
}

// HTMLURL gets the html url of this subtype
func (m *EscalationPolicyReference) HTMLURL() string {
	return m.htmlUrlField
}

// SetHTMLURL sets the html url of this subtype
func (m *EscalationPolicyReference) SetHTMLURL(val string) {
	m.htmlUrlField = val
}

// ID gets the id of this subtype
func (m *EscalationPolicyReference) ID() string {
	return m.idField
}

// SetID sets the id of this subtype
func (m *EscalationPolicyReference) SetID(val string) {
	m.idField = val
}

// Self gets the self of this subtype
func (m *EscalationPolicyReference) Self() string {
	return m.selfField
}

// SetSelf sets the self of this subtype
func (m *EscalationPolicyReference) SetSelf(val string) {
	m.selfField = val
}

// Summary gets the summary of this subtype
func (m *EscalationPolicyReference) Summary() string {
	return m.summaryField
}

// SetSummary sets the summary of this subtype
func (m *EscalationPolicyReference) SetSummary(val string) {
	m.summaryField = val
}

// Type gets the type of this subtype
func (m *EscalationPolicyReference) Type() string {
	return "EscalationPolicyReference"
}

// SetType sets the type of this subtype
func (m *EscalationPolicyReference) SetType(val string) {

}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *EscalationPolicyReference) UnmarshalJSON(raw []byte) error {
	var data struct {
		EscalationPolicyReferenceAllOf1
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		HTMLURL string `json:"html_url,omitempty"`

		ID string `json:"id,omitempty"`

		Self string `json:"self,omitempty"`

		Summary string `json:"summary,omitempty"`

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result EscalationPolicyReference

	result.htmlUrlField = base.HTMLURL

	result.idField = base.ID

	result.selfField = base.Self

	result.summaryField = base.Summary

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.EscalationPolicyReferenceAllOf1 = data.EscalationPolicyReferenceAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m EscalationPolicyReference) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		EscalationPolicyReferenceAllOf1
	}{

		EscalationPolicyReferenceAllOf1: m.EscalationPolicyReferenceAllOf1,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		HTMLURL string `json:"html_url,omitempty"`

		ID string `json:"id,omitempty"`

		Self string `json:"self,omitempty"`

		Summary string `json:"summary,omitempty"`

		Type string `json:"type"`
	}{

		HTMLURL: m.HTMLURL(),

		ID: m.ID(),

		Self: m.Self(),

		Summary: m.Summary(),

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this escalation policy reference
func (m *EscalationPolicyReference) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EscalationPolicyReferenceAllOf1
	if err := m.EscalationPolicyReferenceAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *EscalationPolicyReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EscalationPolicyReference) UnmarshalBinary(b []byte) error {
	var res EscalationPolicyReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}