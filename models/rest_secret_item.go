// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestSecretItem Secret data field item
//
// swagger:model RestSecretItem
type RestSecretItem struct {

	// Longer description of the secret field.
	FieldDescription string `json:"fieldDescription,omitempty"`

	// The id of the field definition from the secret template.
	FieldID *int32 `json:"fieldId,omitempty"`

	// The display name of the secret field.
	FieldName string `json:"fieldName,omitempty"`

	// If the field is a file attachment field, the id of the file attachment.
	FileAttachmentID *int32 `json:"fileAttachmentId,omitempty"`

	// If the field is a file attachment field, the name of the attached file.
	Filename string `json:"filename,omitempty"`

	// Whether the field is a file attachment.
	IsFile bool `json:"isFile,omitempty"`

	// Whether or not the secret field is a list.
	IsList bool `json:"isList,omitempty"`

	// Whether the field is represented as a multi-line text box. Used for long-form text fields.
	IsNotes bool `json:"isNotes,omitempty"`

	// Whether the field is a password. Password fields are hidden by default in the UI and their value is not returned in GET calls that return secrets. To retrieve a password field value, make a GET call to /api/secrets/{secretId}/fields/{slug}.
	IsPassword bool `json:"isPassword,omitempty"`

	// The id of the secret field item. Leave empty when creating a new secret.
	ItemID *int32 `json:"itemId,omitempty"`

	// The value of the secret field item. For list fields, this is a comma-delimited list of the list id guids that are assigned to this field.
	ItemValue string `json:"itemValue,omitempty"`

	// The type of list. Valid values are “None”, “Generic”, and “URL”.
	ListType SecretFieldListType `json:"listType,omitempty"`

	// A unique name for the secret field on the template. Slugs cannot contain spaces and are used in many places to easily refer to a secret field without having to know the field id.
	Slug string `json:"slug,omitempty"`
}

// Validate validates this rest secret item
func (m *RestSecretItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateListType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestSecretItem) validateListType(formats strfmt.Registry) error {
	if swag.IsZero(m.ListType) { // not required
		return nil
	}

	if err := m.ListType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("listType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("listType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this rest secret item based on the context it is used
func (m *RestSecretItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateListType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestSecretItem) contextValidateListType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.ListType) { // not required
		return nil
	}

	if err := m.ListType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("listType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("listType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestSecretItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestSecretItem) UnmarshalBinary(b []byte) error {
	var res RestSecretItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
