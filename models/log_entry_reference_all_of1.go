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

// LogEntryReferenceAllOf1 log entry reference all of1
// swagger:model logEntryReferenceAllOf1
type LogEntryReferenceAllOf1 struct {

	// type
	// Enum: [acknowledge_log_entry acknowledge_log_entry_reference annotate_log_entry annotate_log_entry_reference assign_log_entry assign_log_entry_reference escalate_log_entry escalate_log_entry_reference exhaust_escalation_path_log_entry exhaust_escalation_path_log_entry_reference notify_log_entry notify_log_entry_reference reach_trigger_limit_log_entry reach_trigger_limit_log_entry_reference repeat_escalation_path_log_entry repeat_escalation_path_log_entry_reference resolve_log_entry resolve_log_entry_reference snooze_log_entry snooze_log_entry_reference trigger_log_entry trigger_log_entry_reference unacknowledge_log_entry unacknowledge_log_entry_reference]
	Type string `json:"type,omitempty"`
}

// Validate validates this log entry reference all of1
func (m *LogEntryReferenceAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var logEntryReferenceAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["acknowledge_log_entry","acknowledge_log_entry_reference","annotate_log_entry","annotate_log_entry_reference","assign_log_entry","assign_log_entry_reference","escalate_log_entry","escalate_log_entry_reference","exhaust_escalation_path_log_entry","exhaust_escalation_path_log_entry_reference","notify_log_entry","notify_log_entry_reference","reach_trigger_limit_log_entry","reach_trigger_limit_log_entry_reference","repeat_escalation_path_log_entry","repeat_escalation_path_log_entry_reference","resolve_log_entry","resolve_log_entry_reference","snooze_log_entry","snooze_log_entry_reference","trigger_log_entry","trigger_log_entry_reference","unacknowledge_log_entry","unacknowledge_log_entry_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logEntryReferenceAllOf1TypeTypePropEnum = append(logEntryReferenceAllOf1TypeTypePropEnum, v)
	}
}

const (

	// LogEntryReferenceAllOf1TypeAcknowledgeLogEntry captures enum value "acknowledge_log_entry"
	LogEntryReferenceAllOf1TypeAcknowledgeLogEntry string = "acknowledge_log_entry"

	// LogEntryReferenceAllOf1TypeAcknowledgeLogEntryReference captures enum value "acknowledge_log_entry_reference"
	LogEntryReferenceAllOf1TypeAcknowledgeLogEntryReference string = "acknowledge_log_entry_reference"

	// LogEntryReferenceAllOf1TypeAnnotateLogEntry captures enum value "annotate_log_entry"
	LogEntryReferenceAllOf1TypeAnnotateLogEntry string = "annotate_log_entry"

	// LogEntryReferenceAllOf1TypeAnnotateLogEntryReference captures enum value "annotate_log_entry_reference"
	LogEntryReferenceAllOf1TypeAnnotateLogEntryReference string = "annotate_log_entry_reference"

	// LogEntryReferenceAllOf1TypeAssignLogEntry captures enum value "assign_log_entry"
	LogEntryReferenceAllOf1TypeAssignLogEntry string = "assign_log_entry"

	// LogEntryReferenceAllOf1TypeAssignLogEntryReference captures enum value "assign_log_entry_reference"
	LogEntryReferenceAllOf1TypeAssignLogEntryReference string = "assign_log_entry_reference"

	// LogEntryReferenceAllOf1TypeEscalateLogEntry captures enum value "escalate_log_entry"
	LogEntryReferenceAllOf1TypeEscalateLogEntry string = "escalate_log_entry"

	// LogEntryReferenceAllOf1TypeEscalateLogEntryReference captures enum value "escalate_log_entry_reference"
	LogEntryReferenceAllOf1TypeEscalateLogEntryReference string = "escalate_log_entry_reference"

	// LogEntryReferenceAllOf1TypeExhaustEscalationPathLogEntry captures enum value "exhaust_escalation_path_log_entry"
	LogEntryReferenceAllOf1TypeExhaustEscalationPathLogEntry string = "exhaust_escalation_path_log_entry"

	// LogEntryReferenceAllOf1TypeExhaustEscalationPathLogEntryReference captures enum value "exhaust_escalation_path_log_entry_reference"
	LogEntryReferenceAllOf1TypeExhaustEscalationPathLogEntryReference string = "exhaust_escalation_path_log_entry_reference"

	// LogEntryReferenceAllOf1TypeNotifyLogEntry captures enum value "notify_log_entry"
	LogEntryReferenceAllOf1TypeNotifyLogEntry string = "notify_log_entry"

	// LogEntryReferenceAllOf1TypeNotifyLogEntryReference captures enum value "notify_log_entry_reference"
	LogEntryReferenceAllOf1TypeNotifyLogEntryReference string = "notify_log_entry_reference"

	// LogEntryReferenceAllOf1TypeReachTriggerLimitLogEntry captures enum value "reach_trigger_limit_log_entry"
	LogEntryReferenceAllOf1TypeReachTriggerLimitLogEntry string = "reach_trigger_limit_log_entry"

	// LogEntryReferenceAllOf1TypeReachTriggerLimitLogEntryReference captures enum value "reach_trigger_limit_log_entry_reference"
	LogEntryReferenceAllOf1TypeReachTriggerLimitLogEntryReference string = "reach_trigger_limit_log_entry_reference"

	// LogEntryReferenceAllOf1TypeRepeatEscalationPathLogEntry captures enum value "repeat_escalation_path_log_entry"
	LogEntryReferenceAllOf1TypeRepeatEscalationPathLogEntry string = "repeat_escalation_path_log_entry"

	// LogEntryReferenceAllOf1TypeRepeatEscalationPathLogEntryReference captures enum value "repeat_escalation_path_log_entry_reference"
	LogEntryReferenceAllOf1TypeRepeatEscalationPathLogEntryReference string = "repeat_escalation_path_log_entry_reference"

	// LogEntryReferenceAllOf1TypeResolveLogEntry captures enum value "resolve_log_entry"
	LogEntryReferenceAllOf1TypeResolveLogEntry string = "resolve_log_entry"

	// LogEntryReferenceAllOf1TypeResolveLogEntryReference captures enum value "resolve_log_entry_reference"
	LogEntryReferenceAllOf1TypeResolveLogEntryReference string = "resolve_log_entry_reference"

	// LogEntryReferenceAllOf1TypeSnoozeLogEntry captures enum value "snooze_log_entry"
	LogEntryReferenceAllOf1TypeSnoozeLogEntry string = "snooze_log_entry"

	// LogEntryReferenceAllOf1TypeSnoozeLogEntryReference captures enum value "snooze_log_entry_reference"
	LogEntryReferenceAllOf1TypeSnoozeLogEntryReference string = "snooze_log_entry_reference"

	// LogEntryReferenceAllOf1TypeTriggerLogEntry captures enum value "trigger_log_entry"
	LogEntryReferenceAllOf1TypeTriggerLogEntry string = "trigger_log_entry"

	// LogEntryReferenceAllOf1TypeTriggerLogEntryReference captures enum value "trigger_log_entry_reference"
	LogEntryReferenceAllOf1TypeTriggerLogEntryReference string = "trigger_log_entry_reference"

	// LogEntryReferenceAllOf1TypeUnacknowledgeLogEntry captures enum value "unacknowledge_log_entry"
	LogEntryReferenceAllOf1TypeUnacknowledgeLogEntry string = "unacknowledge_log_entry"

	// LogEntryReferenceAllOf1TypeUnacknowledgeLogEntryReference captures enum value "unacknowledge_log_entry_reference"
	LogEntryReferenceAllOf1TypeUnacknowledgeLogEntryReference string = "unacknowledge_log_entry_reference"
)

// prop value enum
func (m *LogEntryReferenceAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, logEntryReferenceAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LogEntryReferenceAllOf1) validateType(formats strfmt.Registry) error {

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
func (m *LogEntryReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogEntryReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res LogEntryReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}