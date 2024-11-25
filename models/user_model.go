// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserModel User
//
// swagger:model UserModel
type UserModel struct {

	// Active Directory account expiration time
	// Format: date-time
	AdAccountExpires strfmt.DateTime `json:"adAccountExpires,omitempty"`

	// Active Directory unique identifier
	AdGUID string `json:"adGuid,omitempty"`

	// User creation time
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// DateOptionId
	DateOptionID int32 `json:"dateOptionId,omitempty"`

	// Display name
	DisplayName string `json:"displayName,omitempty"`

	// Active Directory domain ID
	DomainID int32 `json:"domainId,omitempty"`

	// Whether Duo two-factor authentication is enabled
	DuoTwoFactor bool `json:"duoTwoFactor,omitempty"`

	// Email address
	EmailAddress string `json:"emailAddress,omitempty"`

	// Whether the user account is enabled
	Enabled bool `json:"enabled,omitempty"`

	// Whether FIDO2 two-factor authentication is enabled
	Fido2TwoFactor bool `json:"fido2TwoFactor,omitempty"`

	// User ID
	ID int32 `json:"id,omitempty"`

	// Array of IP Address Restrictions for the user.
	IPAddressRestrictions []*UserIPAddressRestrictionModel `json:"ipAddressRestrictions"`

	// IsApplicationAccount
	IsApplicationAccount bool `json:"isApplicationAccount,omitempty"`

	// Whether the email address is derived from the Active Directory account
	IsEmailCopiedFromAD bool `json:"isEmailCopiedFromAD,omitempty"`

	// Whether the email address has been verified
	IsEmailVerified bool `json:"isEmailVerified,omitempty"`

	// Whether the user is locked out
	IsLockedOut bool `json:"isLockedOut,omitempty"`

	// Time of last login
	// Format: date-time
	LastLogin strfmt.DateTime `json:"lastLogin,omitempty"`

	// Time of last session activity
	// Format: date-time
	LastSessionActivity *strfmt.DateTime `json:"lastSessionActivity,omitempty"`

	// The reason for the lock out
	LockOutReason *string `json:"lockOutReason,omitempty"`

	// An optional description of the reason for the lock out
	LockOutReasonDescription string `json:"lockOutReasonDescription,omitempty"`

	// Number of login failures
	LoginFailures int32 `json:"loginFailures,omitempty"`

	// Whether the user must verify their email address
	MustVerifyEmail bool `json:"mustVerifyEmail,omitempty"`

	// Whether OATH two-factor authentication is enabled
	OathTwoFactor bool `json:"oathTwoFactor,omitempty"`

	// Whether OATH has been verified
	OathVerified bool `json:"oathVerified,omitempty"`

	// Time when the password was last changed
	// Format: date-time
	PasswordLastChanged strfmt.DateTime `json:"passwordLastChanged,omitempty"`

	// The personal group ID for this user.  Each user has a personal group that is a group that only contains that user.
	PersonalGroupID int32 `json:"personalGroupId,omitempty"`

	// Whether RADIUS two-factor authentication is enabled
	RadiusTwoFactor bool `json:"radiusTwoFactor,omitempty"`

	// RADIUS username
	RadiusUserName string `json:"radiusUserName,omitempty"`

	// ResetSessionStarted
	// Format: date-time
	ResetSessionStarted strfmt.DateTime `json:"resetSessionStarted,omitempty"`

	// Slack ID of the user
	SlackID string `json:"slackId,omitempty"`

	// TimeOptionId
	TimeOptionID int32 `json:"timeOptionId,omitempty"`

	// Whether two-factor authentication is enabled
	TwoFactor bool `json:"twoFactor,omitempty"`

	// UserLcid
	UserLcid int32 `json:"userLcid,omitempty"`

	// User name
	UserName string `json:"userName,omitempty"`

	// Time when the verification email was sent
	// Format: date-time
	VerifyEmailSentDate strfmt.DateTime `json:"verifyEmailSentDate,omitempty"`
}

// Validate validates this user model
func (m *UserModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdAccountExpires(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddressRestrictions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastLogin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSessionActivity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordLastChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResetSessionStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerifyEmailSentDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserModel) validateAdAccountExpires(formats strfmt.Registry) error {
	if swag.IsZero(m.AdAccountExpires) { // not required
		return nil
	}

	if err := validate.FormatOf("adAccountExpires", "body", "date-time", m.AdAccountExpires.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserModel) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserModel) validateIPAddressRestrictions(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAddressRestrictions) { // not required
		return nil
	}

	for i := 0; i < len(m.IPAddressRestrictions); i++ {
		if swag.IsZero(m.IPAddressRestrictions[i]) { // not required
			continue
		}

		if m.IPAddressRestrictions[i] != nil {
			if err := m.IPAddressRestrictions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipAddressRestrictions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipAddressRestrictions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserModel) validateLastLogin(formats strfmt.Registry) error {
	if swag.IsZero(m.LastLogin) { // not required
		return nil
	}

	if err := validate.FormatOf("lastLogin", "body", "date-time", m.LastLogin.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserModel) validateLastSessionActivity(formats strfmt.Registry) error {
	if swag.IsZero(m.LastSessionActivity) { // not required
		return nil
	}

	if err := validate.FormatOf("lastSessionActivity", "body", "date-time", m.LastSessionActivity.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserModel) validatePasswordLastChanged(formats strfmt.Registry) error {
	if swag.IsZero(m.PasswordLastChanged) { // not required
		return nil
	}

	if err := validate.FormatOf("passwordLastChanged", "body", "date-time", m.PasswordLastChanged.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserModel) validateResetSessionStarted(formats strfmt.Registry) error {
	if swag.IsZero(m.ResetSessionStarted) { // not required
		return nil
	}

	if err := validate.FormatOf("resetSessionStarted", "body", "date-time", m.ResetSessionStarted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserModel) validateVerifyEmailSentDate(formats strfmt.Registry) error {
	if swag.IsZero(m.VerifyEmailSentDate) { // not required
		return nil
	}

	if err := validate.FormatOf("verifyEmailSentDate", "body", "date-time", m.VerifyEmailSentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user model based on the context it is used
func (m *UserModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPAddressRestrictions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserModel) contextValidateIPAddressRestrictions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPAddressRestrictions); i++ {

		if m.IPAddressRestrictions[i] != nil {

			if swag.IsZero(m.IPAddressRestrictions[i]) { // not required
				return nil
			}

			if err := m.IPAddressRestrictions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipAddressRestrictions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipAddressRestrictions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserModel) UnmarshalBinary(b []byte) error {
	var res UserModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
