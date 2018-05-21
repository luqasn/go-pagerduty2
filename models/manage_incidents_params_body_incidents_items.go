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

// ManageIncidentsParamsBodyIncidentsItems manage incidents params body incidents items
// swagger:model manageIncidentsParamsBodyIncidentsItems
type ManageIncidentsParamsBodyIncidentsItems struct {

	// Assign the incident to these assignees.
	Assignments []*ManageIncidentsParamsBodyIncidentsItemsAssignmentsItems `json:"assignments"`

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

// Validate validates this manage incidents params body incidents items
func (m *ManageIncidentsParamsBodyIncidentsItems) Validate(formats strfmt.Registry) error {
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

func (m *ManageIncidentsParamsBodyIncidentsItems) validateAssignments(formats strfmt.Registry) error {

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

func (m *ManageIncidentsParamsBodyIncidentsItems) validateEscalationPolicy(formats strfmt.Registry) error {

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

func (m *ManageIncidentsParamsBodyIncidentsItems) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ManageIncidentsParamsBodyIncidentsItems) validatePriority(formats strfmt.Registry) error {

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

var manageIncidentsParamsBodyIncidentsItemsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["resolved","acknowledged"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		manageIncidentsParamsBodyIncidentsItemsTypeStatusPropEnum = append(manageIncidentsParamsBodyIncidentsItemsTypeStatusPropEnum, v)
	}
}

const (

	// ManageIncidentsParamsBodyIncidentsItemsStatusResolved captures enum value "resolved"
	ManageIncidentsParamsBodyIncidentsItemsStatusResolved string = "resolved"

	// ManageIncidentsParamsBodyIncidentsItemsStatusAcknowledged captures enum value "acknowledged"
	ManageIncidentsParamsBodyIncidentsItemsStatusAcknowledged string = "acknowledged"
)

// prop value enum
func (m *ManageIncidentsParamsBodyIncidentsItems) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, manageIncidentsParamsBodyIncidentsItemsTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ManageIncidentsParamsBodyIncidentsItems) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var manageIncidentsParamsBodyIncidentsItemsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["incident","incident_reference"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		manageIncidentsParamsBodyIncidentsItemsTypeTypePropEnum = append(manageIncidentsParamsBodyIncidentsItemsTypeTypePropEnum, v)
	}
}

const (

	// ManageIncidentsParamsBodyIncidentsItemsTypeIncident captures enum value "incident"
	ManageIncidentsParamsBodyIncidentsItemsTypeIncident string = "incident"

	// ManageIncidentsParamsBodyIncidentsItemsTypeIncidentReference captures enum value "incident_reference"
	ManageIncidentsParamsBodyIncidentsItemsTypeIncidentReference string = "incident_reference"
)

// prop value enum
func (m *ManageIncidentsParamsBodyIncidentsItems) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, manageIncidentsParamsBodyIncidentsItemsTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ManageIncidentsParamsBodyIncidentsItems) validateType(formats strfmt.Registry) error {

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
func (m *ManageIncidentsParamsBodyIncidentsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManageIncidentsParamsBodyIncidentsItems) UnmarshalBinary(b []byte) error {
	var res ManageIncidentsParamsBodyIncidentsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}