// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RoleModel Role
//
// swagger:model RoleModel
type RoleModel struct {

	// Created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// Role Id
	ID int32 `json:"id,omitempty"`

	// Is System Role
	IsSystem bool `json:"isSystem,omitempty"`

	// Name
	Name string `json:"name,omitempty"`
}

// Validate validates this role model
func (m *RoleModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleModel) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this role model based on context it is used
func (m *RoleModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoleModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleModel) UnmarshalBinary(b []byte) error {
	var res RoleModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
