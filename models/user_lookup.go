// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserLookup Simple user representation
//
// swagger:model UserLookup
type UserLookup struct {

	// User ID
	ID int32 `json:"id,omitempty"`

	// User name
	Value string `json:"value,omitempty"`
}

// Validate validates this user lookup
func (m *UserLookup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user lookup based on context it is used
func (m *UserLookup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserLookup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserLookup) UnmarshalBinary(b []byte) error {
	var res UserLookup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
