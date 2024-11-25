// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SecretDetailExpirationType Expiration Type
//
// swagger:model SecretDetailExpirationType
type SecretDetailExpirationType string

func NewSecretDetailExpirationType(value SecretDetailExpirationType) *SecretDetailExpirationType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SecretDetailExpirationType.
func (m SecretDetailExpirationType) Pointer() *SecretDetailExpirationType {
	return &m
}

const (

	// SecretDetailExpirationTypeTemplate captures enum value "Template"
	SecretDetailExpirationTypeTemplate SecretDetailExpirationType = "Template"

	// SecretDetailExpirationTypeDayInterval captures enum value "DayInterval"
	SecretDetailExpirationTypeDayInterval SecretDetailExpirationType = "DayInterval"

	// SecretDetailExpirationTypeSpecificDate captures enum value "SpecificDate"
	SecretDetailExpirationTypeSpecificDate SecretDetailExpirationType = "SpecificDate"

	// SecretDetailExpirationTypeDisabled captures enum value "Disabled"
	SecretDetailExpirationTypeDisabled SecretDetailExpirationType = "Disabled"
)

// for schema
var secretDetailExpirationTypeEnum []interface{}

func init() {
	var res []SecretDetailExpirationType
	if err := json.Unmarshal([]byte(`["Template","DayInterval","SpecificDate","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		secretDetailExpirationTypeEnum = append(secretDetailExpirationTypeEnum, v)
	}
}

func (m SecretDetailExpirationType) validateSecretDetailExpirationTypeEnum(path, location string, value SecretDetailExpirationType) error {
	if err := validate.EnumCase(path, location, value, secretDetailExpirationTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this secret detail expiration type
func (m SecretDetailExpirationType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSecretDetailExpirationTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this secret detail expiration type based on context it is used
func (m SecretDetailExpirationType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}