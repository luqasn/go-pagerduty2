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

// AlertAllOf1 alert all of1
// swagger:model alertAllOf1
type AlertAllOf1 struct {

	// The alert's de-duplication key.
	// Read Only: true
	AlertKey string `json:"alert_key,omitempty"`

	// body
	Body *AlertAllOf1Body `json:"body,omitempty"`

	// The date/time the alert was first triggered.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The trigger log entry for the alert.
	// Read Only: true
	FirstTriggerLogEntry *LogEntryReference `json:"first_trigger_log_entry,omitempty"`

	// The parent incident of the alert, if any.
	Incident *IncidentReference `json:"incident,omitempty"`

	// The monitoring tool that created the alert due to an incoming event.
	Integration *Integration `json:"integration,omitempty"`

	// The PagerDuty service that the alert belongs to.
	// Read Only: true
	Service *ServiceReference `json:"service,omitempty"`

	// The magnitude of the problem as reported by the monitoring tool.
	// Read Only: true
	// Enum: [info warning error critical]
	Severity string `json:"severity,omitempty"`

	// The current status of the alert.
	// Enum: [triggered resolved]
	Status string `json:"status,omitempty"`

	// Whether or not an alert is suppressed. Suppressed alerts are not created with a parent incident.
	// Read Only: true
	Suppressed *bool `json:"suppressed,omitempty"`
}

// Validate validates this alert all of1
func (m *AlertAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstTriggerLogEntry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncident(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertAllOf1) validateBody(formats strfmt.Registry) error {

	if swag.IsZero(m.Body) { // not required
		return nil
	}

	if m.Body != nil {
		if err := m.Body.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body")
			}
			return err
		}
	}

	return nil
}

func (m *AlertAllOf1) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AlertAllOf1) validateFirstTriggerLogEntry(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstTriggerLogEntry) { // not required
		return nil
	}

	if m.FirstTriggerLogEntry != nil {
		if err := m.FirstTriggerLogEntry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("first_trigger_log_entry")
			}
			return err
		}
	}

	return nil
}

func (m *AlertAllOf1) validateIncident(formats strfmt.Registry) error {

	if swag.IsZero(m.Incident) { // not required
		return nil
	}

	if m.Incident != nil {
		if err := m.Incident.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incident")
			}
			return err
		}
	}

	return nil
}

func (m *AlertAllOf1) validateIntegration(formats strfmt.Registry) error {

	if swag.IsZero(m.Integration) { // not required
		return nil
	}

	if m.Integration != nil {
		if err := m.Integration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

func (m *AlertAllOf1) validateService(formats strfmt.Registry) error {

	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

var alertAllOf1TypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["info","warning","error","critical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertAllOf1TypeSeverityPropEnum = append(alertAllOf1TypeSeverityPropEnum, v)
	}
}

const (

	// AlertAllOf1SeverityInfo captures enum value "info"
	AlertAllOf1SeverityInfo string = "info"

	// AlertAllOf1SeverityWarning captures enum value "warning"
	AlertAllOf1SeverityWarning string = "warning"

	// AlertAllOf1SeverityError captures enum value "error"
	AlertAllOf1SeverityError string = "error"

	// AlertAllOf1SeverityCritical captures enum value "critical"
	AlertAllOf1SeverityCritical string = "critical"
)

// prop value enum
func (m *AlertAllOf1) validateSeverityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, alertAllOf1TypeSeverityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AlertAllOf1) validateSeverity(formats strfmt.Registry) error {

	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

var alertAllOf1TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["triggered","resolved"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertAllOf1TypeStatusPropEnum = append(alertAllOf1TypeStatusPropEnum, v)
	}
}

const (

	// AlertAllOf1StatusTriggered captures enum value "triggered"
	AlertAllOf1StatusTriggered string = "triggered"

	// AlertAllOf1StatusResolved captures enum value "resolved"
	AlertAllOf1StatusResolved string = "resolved"
)

// prop value enum
func (m *AlertAllOf1) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, alertAllOf1TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AlertAllOf1) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertAllOf1) UnmarshalBinary(b []byte) error {
	var res AlertAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
