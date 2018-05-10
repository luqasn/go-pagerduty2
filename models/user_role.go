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

// UserRole user role
// swagger:model UserRole
type UserRole struct {
	resourcesField []Reference

	// The user's role for a set of resources.
	//
	// Enum: [observer responder manager]
	Role string `json:"role,omitempty"`
}

// Resources gets the resources of this base type
func (m *UserRole) Resources() []Reference {
	return m.resourcesField
}

// SetResources sets the resources of this base type
func (m *UserRole) SetResources(val []Reference) {
	m.resourcesField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *UserRole) UnmarshalJSON(raw []byte) error {
	var data struct {
		Resources json.RawMessage `json:"resources,omitempty"`

		Role string `json:"role,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	resources, err := UnmarshalReferenceSlice(bytes.NewBuffer(data.Resources), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result UserRole

	// resources
	result.resourcesField = resources

	// role
	result.Role = data.Role

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m UserRole) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Role string `json:"role,omitempty"`
	}{

		Role: m.Role,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Resources []Reference `json:"resources,omitempty"`
	}{

		Resources: m.resourcesField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this user role
func (m *UserRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserRole) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources()) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources()); i++ {

		if err := m.resourcesField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

var userRoleTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["observer","responder","manager"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userRoleTypeRolePropEnum = append(userRoleTypeRolePropEnum, v)
	}
}

const (

	// UserRoleRoleObserver captures enum value "observer"
	UserRoleRoleObserver string = "observer"

	// UserRoleRoleResponder captures enum value "responder"
	UserRoleRoleResponder string = "responder"

	// UserRoleRoleManager captures enum value "manager"
	UserRoleRoleManager string = "manager"
)

// prop value enum
func (m *UserRole) validateRoleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userRoleTypeRolePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UserRole) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRole) UnmarshalBinary(b []byte) error {
	var res UserRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}