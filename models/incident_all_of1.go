// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IncidentAllOf1 incident all of1
// swagger:model incidentAllOf1
type IncidentAllOf1 struct {

	// List of all acknowledgements for this incident.
	Acknowledgements []*Acknowledgement `json:"acknowledgements"`

	// A summary of the number of alerts by status.
	// Read Only: true
	AlertCounts *AlertCount `json:"alert_counts,omitempty"`

	// List of all assignments for this incident.
	Assignments []*Assignment `json:"assignments"`

	// A JSON object containing data describing the incident.
	// Read Only: true
	Body *IncidentBody `json:"body,omitempty"`

	// The date/time the incident was first triggered.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The escalation policy that the incident is currently following.
	EscalationPolicy *EscalationPolicyReference `json:"escalation_policy,omitempty"`

	// The first trigger log entry for the incident.
	// Read Only: true
	FirstTriggerLogEntry *LogEntryReference `json:"first_trigger_log_entry,omitempty"`

	// The incident's de-duplication key.
	// Read Only: true
	IncidentKey string `json:"incident_key,omitempty"`

	// The number of the incident. This is unique across your account.
	// Read Only: true
	IncidentNumber int64 `json:"incident_number,omitempty"`

	// The time at which the status of the incident last changed.
	// Read Only: true
	// Format: date-time
	LastStatusChangeAt strfmt.DateTime `json:"last_status_change_at,omitempty"`

	// The user or service which is responsible for the incident's last status change. If the incident is in the acknowledged or resolved status, this will be the user that took the first acknowledged or resolved action. If the incident was automatically resolved (say through the Events API), or if the incident is in the triggered state, this will be the incident's service.
	// Read Only: true
	LastStatusChangeBy *Agent `json:"last_status_change_by,omitempty"`

	pendingActionsField []IncidentAction

	// The priority of the incident.
	Priority *Priority `json:"priority,omitempty"`

	// The reason the incident was resolved. Currently the only valid values are `nil` and `merged` with plans to introduce additional reasons in the future.
	// Read Only: true
	ResolveReason *ResolveReason `json:"resolve_reason,omitempty"`

	// The PagerDuty service that the incident belongs to.
	Service *ServiceReference `json:"service,omitempty"`

	// The current status of the incident.
	// Enum: [triggered acknowledged resolved]
	Status string `json:"status,omitempty"`

	// The teams involved in the incident’s lifecycle.
	Teams []*TeamReference `json:"teams"`

	// A succinct description of the nature, symptoms, cause, or effect of the incident.
	Title string `json:"title,omitempty"`

	// The current urgency of the incident.
	// Enum: [high low]
	Urgency string `json:"urgency,omitempty"`
}

// PendingActions gets the pending actions of this base type
func (m *IncidentAllOf1) PendingActions() []IncidentAction {
	return m.pendingActionsField
}

// SetPendingActions sets the pending actions of this base type
func (m *IncidentAllOf1) SetPendingActions(val []IncidentAction) {
	m.pendingActionsField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *IncidentAllOf1) UnmarshalJSON(raw []byte) error {
	var data struct {
		Acknowledgements []*Acknowledgement `json:"acknowledgements,omitempty"`

		AlertCounts *AlertCount `json:"alert_counts,omitempty"`

		Assignments []*Assignment `json:"assignments,omitempty"`

		Body *IncidentBody `json:"body,omitempty"`

		CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

		EscalationPolicy *EscalationPolicyReference `json:"escalation_policy,omitempty"`

		FirstTriggerLogEntry *LogEntryReference `json:"first_trigger_log_entry,omitempty"`

		IncidentKey string `json:"incident_key,omitempty"`

		IncidentNumber int64 `json:"incident_number,omitempty"`

		LastStatusChangeAt strfmt.DateTime `json:"last_status_change_at,omitempty"`

		LastStatusChangeBy *Agent `json:"last_status_change_by,omitempty"`

		PendingActions json.RawMessage `json:"pending_actions,omitempty"`

		Priority *Priority `json:"priority,omitempty"`

		ResolveReason *ResolveReason `json:"resolve_reason,omitempty"`

		Service *ServiceReference `json:"service,omitempty"`

		Status string `json:"status,omitempty"`

		Teams []*TeamReference `json:"teams,omitempty"`

		Title string `json:"title,omitempty"`

		Urgency string `json:"urgency,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	pendingActions, err := UnmarshalIncidentActionSlice(bytes.NewBuffer(data.PendingActions), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result IncidentAllOf1

	// acknowledgements
	result.Acknowledgements = data.Acknowledgements

	// alert_counts
	result.AlertCounts = data.AlertCounts

	// assignments
	result.Assignments = data.Assignments

	// body
	result.Body = data.Body

	// created_at
	result.CreatedAt = data.CreatedAt

	// escalation_policy
	result.EscalationPolicy = data.EscalationPolicy

	// first_trigger_log_entry
	result.FirstTriggerLogEntry = data.FirstTriggerLogEntry

	// incident_key
	result.IncidentKey = data.IncidentKey

	// incident_number
	result.IncidentNumber = data.IncidentNumber

	// last_status_change_at
	result.LastStatusChangeAt = data.LastStatusChangeAt

	// last_status_change_by
	result.LastStatusChangeBy = data.LastStatusChangeBy

	// pending_actions
	result.pendingActionsField = pendingActions

	// priority
	result.Priority = data.Priority

	// resolve_reason
	result.ResolveReason = data.ResolveReason

	// service
	result.Service = data.Service

	// status
	result.Status = data.Status

	// teams
	result.Teams = data.Teams

	// title
	result.Title = data.Title

	// urgency
	result.Urgency = data.Urgency

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m IncidentAllOf1) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Acknowledgements []*Acknowledgement `json:"acknowledgements,omitempty"`

		AlertCounts *AlertCount `json:"alert_counts,omitempty"`

		Assignments []*Assignment `json:"assignments,omitempty"`

		Body *IncidentBody `json:"body,omitempty"`

		CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

		EscalationPolicy *EscalationPolicyReference `json:"escalation_policy,omitempty"`

		FirstTriggerLogEntry *LogEntryReference `json:"first_trigger_log_entry,omitempty"`

		IncidentKey string `json:"incident_key,omitempty"`

		IncidentNumber int64 `json:"incident_number,omitempty"`

		LastStatusChangeAt strfmt.DateTime `json:"last_status_change_at,omitempty"`

		LastStatusChangeBy *Agent `json:"last_status_change_by,omitempty"`

		Priority *Priority `json:"priority,omitempty"`

		ResolveReason *ResolveReason `json:"resolve_reason,omitempty"`

		Service *ServiceReference `json:"service,omitempty"`

		Status string `json:"status,omitempty"`

		Teams []*TeamReference `json:"teams,omitempty"`

		Title string `json:"title,omitempty"`

		Urgency string `json:"urgency,omitempty"`
	}{

		Acknowledgements: m.Acknowledgements,

		AlertCounts: m.AlertCounts,

		Assignments: m.Assignments,

		Body: m.Body,

		CreatedAt: m.CreatedAt,

		EscalationPolicy: m.EscalationPolicy,

		FirstTriggerLogEntry: m.FirstTriggerLogEntry,

		IncidentKey: m.IncidentKey,

		IncidentNumber: m.IncidentNumber,

		LastStatusChangeAt: m.LastStatusChangeAt,

		LastStatusChangeBy: m.LastStatusChangeBy,

		Priority: m.Priority,

		ResolveReason: m.ResolveReason,

		Service: m.Service,

		Status: m.Status,

		Teams: m.Teams,

		Title: m.Title,

		Urgency: m.Urgency,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		PendingActions []IncidentAction `json:"pending_actions,omitempty"`
	}{

		PendingActions: m.pendingActionsField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this incident all of1
func (m *IncidentAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcknowledgements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlertCounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEscalationPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstTriggerLogEntry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastStatusChangeAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastStatusChangeBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePendingActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolveReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUrgency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IncidentAllOf1) validateAcknowledgements(formats strfmt.Registry) error {

	if swag.IsZero(m.Acknowledgements) { // not required
		return nil
	}

	for i := 0; i < len(m.Acknowledgements); i++ {
		if swag.IsZero(m.Acknowledgements[i]) { // not required
			continue
		}

		if m.Acknowledgements[i] != nil {
			if err := m.Acknowledgements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acknowledgements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentAllOf1) validateAlertCounts(formats strfmt.Registry) error {

	if swag.IsZero(m.AlertCounts) { // not required
		return nil
	}

	if m.AlertCounts != nil {
		if err := m.AlertCounts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alert_counts")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentAllOf1) validateAssignments(formats strfmt.Registry) error {

	if swag.IsZero(m.Assignments) { // not required
		return nil
	}

	for i := 0; i < len(m.Assignments); i++ {
		if swag.IsZero(m.Assignments[i]) { // not required
			continue
		}

		if m.Assignments[i] != nil {
			if err := m.Assignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentAllOf1) validateBody(formats strfmt.Registry) error {

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

func (m *IncidentAllOf1) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IncidentAllOf1) validateEscalationPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.EscalationPolicy) { // not required
		return nil
	}

	if m.EscalationPolicy != nil {
		if err := m.EscalationPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("escalation_policy")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentAllOf1) validateFirstTriggerLogEntry(formats strfmt.Registry) error {

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

func (m *IncidentAllOf1) validateLastStatusChangeAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastStatusChangeAt) { // not required
		return nil
	}

	if err := validate.FormatOf("last_status_change_at", "body", "date-time", m.LastStatusChangeAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IncidentAllOf1) validateLastStatusChangeBy(formats strfmt.Registry) error {

	if swag.IsZero(m.LastStatusChangeBy) { // not required
		return nil
	}

	if m.LastStatusChangeBy != nil {
		if err := m.LastStatusChangeBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_status_change_by")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentAllOf1) validatePendingActions(formats strfmt.Registry) error {

	if swag.IsZero(m.PendingActions()) { // not required
		return nil
	}

	for i := 0; i < len(m.PendingActions()); i++ {

		if err := m.pendingActionsField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pending_actions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *IncidentAllOf1) validatePriority(formats strfmt.Registry) error {

	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	if m.Priority != nil {
		if err := m.Priority.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priority")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentAllOf1) validateResolveReason(formats strfmt.Registry) error {

	if swag.IsZero(m.ResolveReason) { // not required
		return nil
	}

	if m.ResolveReason != nil {
		if err := m.ResolveReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resolve_reason")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentAllOf1) validateService(formats strfmt.Registry) error {

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

var incidentAllOf1TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["triggered","acknowledged","resolved"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		incidentAllOf1TypeStatusPropEnum = append(incidentAllOf1TypeStatusPropEnum, v)
	}
}

const (

	// IncidentAllOf1StatusTriggered captures enum value "triggered"
	IncidentAllOf1StatusTriggered string = "triggered"

	// IncidentAllOf1StatusAcknowledged captures enum value "acknowledged"
	IncidentAllOf1StatusAcknowledged string = "acknowledged"

	// IncidentAllOf1StatusResolved captures enum value "resolved"
	IncidentAllOf1StatusResolved string = "resolved"
)

// prop value enum
func (m *IncidentAllOf1) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, incidentAllOf1TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IncidentAllOf1) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *IncidentAllOf1) validateTeams(formats strfmt.Registry) error {

	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	for i := 0; i < len(m.Teams); i++ {
		if swag.IsZero(m.Teams[i]) { // not required
			continue
		}

		if m.Teams[i] != nil {
			if err := m.Teams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var incidentAllOf1TypeUrgencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["high","low"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		incidentAllOf1TypeUrgencyPropEnum = append(incidentAllOf1TypeUrgencyPropEnum, v)
	}
}

const (

	// IncidentAllOf1UrgencyHigh captures enum value "high"
	IncidentAllOf1UrgencyHigh string = "high"

	// IncidentAllOf1UrgencyLow captures enum value "low"
	IncidentAllOf1UrgencyLow string = "low"
)

// prop value enum
func (m *IncidentAllOf1) validateUrgencyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, incidentAllOf1TypeUrgencyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IncidentAllOf1) validateUrgency(formats strfmt.Registry) error {

	if swag.IsZero(m.Urgency) { // not required
		return nil
	}

	// value enum
	if err := m.validateUrgencyEnum("urgency", "body", m.Urgency); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IncidentAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncidentAllOf1) UnmarshalBinary(b []byte) error {
	var res IncidentAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
