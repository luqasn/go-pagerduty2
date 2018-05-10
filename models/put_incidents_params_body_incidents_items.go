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

// PutIncidentsParamsBodyIncidentsItems put incidents params body incidents items
// swagger:model putIncidentsParamsBodyIncidentsItems
type PutIncidentsParamsBodyIncidentsItems struct {

	// Assign the incident to these assignees.
	Assignments []*PutIncidentsParamsBodyIncidentsItemsAssignmentsItems `json:"assignments"`

	// Escalate the incident to this level in the escalation policy.
	EscalationLevel int64 `json:"escalation_level,omitempty"`

	// Delegate this incident to the specified escalation policy. This restarts the incident's escalation following the new policy.
	EscalationPolicy *EscalationPolicyReference `json:"escalation_policy,omitempty"`

	// The id of the incident to update.
	// Required: true
	ID *string `json:"id"`

	// The priority of the incident.
	Priority *PriorityReference `json:"priority,omitempty"`

	// The resolution for this incident if status is set to resolved.
	Resolution string `json:"resolution,omitempty"`

	// The new status of the incident.
	// Enum: [resolved acknowledged]
	Status string `json:"status,omitempty"`

	// A succinct description of the nature, symptoms, cause, or effect of the incident.
	Title string `json:"title,omitempty"`

	// The incident type.
	// Required: true
	// Enum: [incident incident_reference]
	Type *string `json:"type"`
}

// Validate validates this put incidents params body incidents items
func (m *PutIncidentsParamsBodyIncidentsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEscalationPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutIncidentsParamsBodyIncidentsItems) validateAssignments(formats strfmt.Registry) error {

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

func (m *PutIncidentsParamsBodyIncidentsItems) validateEscalationPolicy(formats strfmt.Registry) error {

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

func (m *PutIncidentsParamsBodyIncidentsItems) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PutIncidentsParamsBodyIncidentsItems) validatePriority(formats strfmt.Registry) error {

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

var putIncidentsParamsBodyIncidentsItemsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["resolved","acknowledged"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		putIncidentsParamsBodyIncidentsItemsTypeStatusPropEnum = append(putIncidentsParamsBodyIncidentsItemsTypeStatusPropEnum, v)
	}
}

const (

	// PutIncidentsParamsBodyIncidentsItemsStatusResolved captures enum value "resolved"
	PutIncidentsParamsBodyIncidentsItemsStatusResolved string = "resolved"

	// PutIncidentsParamsBodyIncidentsItemsStatusAcknowledged captures enum value "acknowledged"
	PutIncidentsParamsBodyIncidentsItemsStatusAcknowledged string = "acknowledged"
)

// prop value enum
func (m *PutIncidentsParamsBodyIncidentsItems) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, putIncidentsParamsBodyIncidentsItemsTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PutIncidentsParamsBodyIncidentsItems) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var putIncidentsParamsBodyIncidentsItemsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["incident","incident_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		putIncidentsParamsBodyIncidentsItemsTypeTypePropEnum = append(putIncidentsParamsBodyIncidentsItemsTypeTypePropEnum, v)
	}
}

const (

	// PutIncidentsParamsBodyIncidentsItemsTypeIncident captures enum value "incident"
	PutIncidentsParamsBodyIncidentsItemsTypeIncident string = "incident"

	// PutIncidentsParamsBodyIncidentsItemsTypeIncidentReference captures enum value "incident_reference"
	PutIncidentsParamsBodyIncidentsItemsTypeIncidentReference string = "incident_reference"
)

// prop value enum
func (m *PutIncidentsParamsBodyIncidentsItems) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, putIncidentsParamsBodyIncidentsItemsTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PutIncidentsParamsBodyIncidentsItems) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutIncidentsParamsBodyIncidentsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutIncidentsParamsBodyIncidentsItems) UnmarshalBinary(b []byte) error {
	var res PutIncidentsParamsBodyIncidentsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}