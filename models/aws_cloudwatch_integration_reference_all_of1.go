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

// AwsCloudwatchIntegrationReferenceAllOf1 aws cloudwatch integration reference all of1
// swagger:model awsCloudwatchIntegrationReferenceAllOf1
type AwsCloudwatchIntegrationReferenceAllOf1 struct {

	// type
	// Enum: [aws_cloudwatch_inbound_integration aws_cloudwatch_inbound_integration_reference]
	Type string `json:"type,omitempty"`
}

// Validate validates this aws cloudwatch integration reference all of1
func (m *AwsCloudwatchIntegrationReferenceAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var awsCloudwatchIntegrationReferenceAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aws_cloudwatch_inbound_integration","aws_cloudwatch_inbound_integration_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsCloudwatchIntegrationReferenceAllOf1TypeTypePropEnum = append(awsCloudwatchIntegrationReferenceAllOf1TypeTypePropEnum, v)
	}
}

const (

	// AwsCloudwatchIntegrationReferenceAllOf1TypeAwsCloudwatchInboundIntegration captures enum value "aws_cloudwatch_inbound_integration"
	AwsCloudwatchIntegrationReferenceAllOf1TypeAwsCloudwatchInboundIntegration string = "aws_cloudwatch_inbound_integration"

	// AwsCloudwatchIntegrationReferenceAllOf1TypeAwsCloudwatchInboundIntegrationReference captures enum value "aws_cloudwatch_inbound_integration_reference"
	AwsCloudwatchIntegrationReferenceAllOf1TypeAwsCloudwatchInboundIntegrationReference string = "aws_cloudwatch_inbound_integration_reference"
)

// prop value enum
func (m *AwsCloudwatchIntegrationReferenceAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, awsCloudwatchIntegrationReferenceAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AwsCloudwatchIntegrationReferenceAllOf1) validateType(formats strfmt.Registry) error {

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
func (m *AwsCloudwatchIntegrationReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsCloudwatchIntegrationReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res AwsCloudwatchIntegrationReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}