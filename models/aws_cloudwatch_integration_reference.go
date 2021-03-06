// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AwsCloudwatchIntegrationReference aws cloudwatch integration reference
// swagger:model AwsCloudwatchIntegrationReference
type AwsCloudwatchIntegrationReference struct {
	IntegrationReference

	AwsCloudwatchIntegrationReferenceAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AwsCloudwatchIntegrationReference) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 IntegrationReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.IntegrationReference = aO0

	// AO1
	var aO1 AwsCloudwatchIntegrationReferenceAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.AwsCloudwatchIntegrationReferenceAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AwsCloudwatchIntegrationReference) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.IntegrationReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.AwsCloudwatchIntegrationReferenceAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this aws cloudwatch integration reference
func (m *AwsCloudwatchIntegrationReference) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with IntegrationReference
	if err := m.IntegrationReference.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with AwsCloudwatchIntegrationReferenceAllOf1
	if err := m.AwsCloudwatchIntegrationReferenceAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AwsCloudwatchIntegrationReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsCloudwatchIntegrationReference) UnmarshalBinary(b []byte) error {
	var res AwsCloudwatchIntegrationReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
