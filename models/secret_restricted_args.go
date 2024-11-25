// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecretRestrictedArgs Restricted secret update options
//
// swagger:model SecretRestrictedArgs
type SecretRestrictedArgs struct {

	// If the secret requires a comment to view or requires approval to view, a reason for accessing the secret must be provided.
	Comment string `json:"comment,omitempty"`

	// If the secret is DoubleLocked, this is the DoubleLock password needed to access the secret.
	DoubleLockPassword string `json:"doubleLockPassword,omitempty"`

	// Force the secret to be checked in, even if checked out by someone else. The user must have the "Force Check In" permission.
	ForceCheckIn bool `json:"forceCheckIn,omitempty"`

	// If the secret is deactivated, this must be set to true in order to access the secret. The user must also have the "View Inactive Secrets" permission.
	IncludeInactive bool `json:"includeInactive,omitempty"`

	// New secret password.
	NewPassword string `json:"newPassword,omitempty"`

	// Don't check out the secret automatically.
	NoAutoCheckout bool `json:"noAutoCheckout,omitempty"`

	// If the secret requires a comment to view or requires approval and a user must provide a help desk a ticket number, this is the ticket number to the help desk request.
	TicketNumber string `json:"ticketNumber,omitempty"`

	// If the secret requires a comment to view or requires approval and a user must provide a help desk a ticket number, this is the id of the help desk system configured in Secret Server that should be used to validate the ticket number.
	TicketSystemID *int32 `json:"ticketSystemId,omitempty"`
}

// Validate validates this secret restricted args
func (m *SecretRestrictedArgs) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret restricted args based on context it is used
func (m *SecretRestrictedArgs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecretRestrictedArgs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretRestrictedArgs) UnmarshalBinary(b []byte) error {
	var res SecretRestrictedArgs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}