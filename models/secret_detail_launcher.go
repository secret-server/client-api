// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecretDetailLauncher Launchers
//
// swagger:model SecretDetailLauncher
type SecretDetailLauncher struct {

	// HasProxyCredentials
	HasProxyCredentials bool `json:"hasProxyCredentials,omitempty"`

	// ImagePath
	ImagePath string `json:"imagePath,omitempty"`

	// IsRecorded
	IsRecorded bool `json:"isRecorded,omitempty"`

	// Name
	Name string `json:"name,omitempty"`

	// ShowLauncher
	ShowLauncher bool `json:"showLauncher,omitempty"`

	// TypeId
	TypeID int32 `json:"typeId,omitempty"`
}

// Validate validates this secret detail launcher
func (m *SecretDetailLauncher) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret detail launcher based on context it is used
func (m *SecretDetailLauncher) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecretDetailLauncher) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretDetailLauncher) UnmarshalBinary(b []byte) error {
	var res SecretDetailLauncher
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}