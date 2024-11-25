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

// SecretDetailActionType Secret Detail Action Type
//
// swagger:model SecretDetailActionType
type SecretDetailActionType string

func NewSecretDetailActionType(value SecretDetailActionType) *SecretDetailActionType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SecretDetailActionType.
func (m SecretDetailActionType) Pointer() *SecretDetailActionType {
	return &m
}

const (

	// SecretDetailActionTypeChangePasswordNow captures enum value "ChangePasswordNow"
	SecretDetailActionTypeChangePasswordNow SecretDetailActionType = "ChangePasswordNow"

	// SecretDetailActionTypeConvertTemplate captures enum value "ConvertTemplate"
	SecretDetailActionTypeConvertTemplate SecretDetailActionType = "ConvertTemplate"

	// SecretDetailActionTypeCopy captures enum value "Copy"
	SecretDetailActionTypeCopy SecretDetailActionType = "Copy"

	// SecretDetailActionTypeDelete captures enum value "Delete"
	SecretDetailActionTypeDelete SecretDetailActionType = "Delete"

	// SecretDetailActionTypeEdit captures enum value "Edit"
	SecretDetailActionTypeEdit SecretDetailActionType = "Edit"

	// SecretDetailActionTypeEditExpiration captures enum value "EditExpiration"
	SecretDetailActionTypeEditExpiration SecretDetailActionType = "EditExpiration"

	// SecretDetailActionTypeEditRPC captures enum value "EditRpc"
	SecretDetailActionTypeEditRPC SecretDetailActionType = "EditRpc"

	// SecretDetailActionTypeEditSecurity captures enum value "EditSecurity"
	SecretDetailActionTypeEditSecurity SecretDetailActionType = "EditSecurity"

	// SecretDetailActionTypeExpire captures enum value "Expire"
	SecretDetailActionTypeExpire SecretDetailActionType = "Expire"

	// SecretDetailActionTypeHeartbeat captures enum value "Heartbeat"
	SecretDetailActionTypeHeartbeat SecretDetailActionType = "Heartbeat"

	// SecretDetailActionTypeEditShare captures enum value "EditShare"
	SecretDetailActionTypeEditShare SecretDetailActionType = "EditShare"

	// SecretDetailActionTypeShowSSHProxyCredentials captures enum value "ShowSshProxyCredentials"
	SecretDetailActionTypeShowSSHProxyCredentials SecretDetailActionType = "ShowSshProxyCredentials"

	// SecretDetailActionTypeStopChangePasswordNow captures enum value "StopChangePasswordNow"
	SecretDetailActionTypeStopChangePasswordNow SecretDetailActionType = "StopChangePasswordNow"

	// SecretDetailActionTypeViewAudit captures enum value "ViewAudit"
	SecretDetailActionTypeViewAudit SecretDetailActionType = "ViewAudit"

	// SecretDetailActionTypeViewDependencies captures enum value "ViewDependencies"
	SecretDetailActionTypeViewDependencies SecretDetailActionType = "ViewDependencies"

	// SecretDetailActionTypeViewLaunchers captures enum value "ViewLaunchers"
	SecretDetailActionTypeViewLaunchers SecretDetailActionType = "ViewLaunchers"

	// SecretDetailActionTypeViewExpiration captures enum value "ViewExpiration"
	SecretDetailActionTypeViewExpiration SecretDetailActionType = "ViewExpiration"

	// SecretDetailActionTypeViewHooks captures enum value "ViewHooks"
	SecretDetailActionTypeViewHooks SecretDetailActionType = "ViewHooks"

	// SecretDetailActionTypeViewRPC captures enum value "ViewRpc"
	SecretDetailActionTypeViewRPC SecretDetailActionType = "ViewRpc"

	// SecretDetailActionTypeViewSecurity captures enum value "ViewSecurity"
	SecretDetailActionTypeViewSecurity SecretDetailActionType = "ViewSecurity"

	// SecretDetailActionTypeViewSettings captures enum value "ViewSettings"
	SecretDetailActionTypeViewSettings SecretDetailActionType = "ViewSettings"

	// SecretDetailActionTypeUndelete captures enum value "Undelete"
	SecretDetailActionTypeUndelete SecretDetailActionType = "Undelete"

	// SecretDetailActionTypeForceCheckin captures enum value "ForceCheckin"
	SecretDetailActionTypeForceCheckin SecretDetailActionType = "ForceCheckin"

	// SecretDetailActionTypeViewShare captures enum value "ViewShare"
	SecretDetailActionTypeViewShare SecretDetailActionType = "ViewShare"

	// SecretDetailActionTypeEditHooks captures enum value "EditHooks"
	SecretDetailActionTypeEditHooks SecretDetailActionType = "EditHooks"

	// SecretDetailActionTypeEditDependencies captures enum value "EditDependencies"
	SecretDetailActionTypeEditDependencies SecretDetailActionType = "EditDependencies"

	// SecretDetailActionTypeViewGeneralDetails captures enum value "ViewGeneralDetails"
	SecretDetailActionTypeViewGeneralDetails SecretDetailActionType = "ViewGeneralDetails"

	// SecretDetailActionTypeViewHeartbeatStatus captures enum value "ViewHeartbeatStatus"
	SecretDetailActionTypeViewHeartbeatStatus SecretDetailActionType = "ViewHeartbeatStatus"

	// SecretDetailActionTypeCheckin captures enum value "Checkin"
	SecretDetailActionTypeCheckin SecretDetailActionType = "Checkin"

	// SecretDetailActionTypeCheckout captures enum value "Checkout"
	SecretDetailActionTypeCheckout SecretDetailActionType = "Checkout"

	// SecretDetailActionTypeGenerateOneTimePassword captures enum value "GenerateOneTimePassword"
	SecretDetailActionTypeGenerateOneTimePassword SecretDetailActionType = "GenerateOneTimePassword"

	// SecretDetailActionTypeShowSSHTerminalDetails captures enum value "ShowSshTerminalDetails"
	SecretDetailActionTypeShowSSHTerminalDetails SecretDetailActionType = "ShowSshTerminalDetails"

	// SecretDetailActionTypeShowRdpProxyCredentials captures enum value "ShowRdpProxyCredentials"
	SecretDetailActionTypeShowRdpProxyCredentials SecretDetailActionType = "ShowRdpProxyCredentials"

	// SecretDetailActionTypeViewMetadata captures enum value "ViewMetadata"
	SecretDetailActionTypeViewMetadata SecretDetailActionType = "ViewMetadata"

	// SecretDetailActionTypeViewSecretExposure captures enum value "ViewSecretExposure"
	SecretDetailActionTypeViewSecretExposure SecretDetailActionType = "ViewSecretExposure"

	// SecretDetailActionTypeViewCheckOut captures enum value "ViewCheckOut"
	SecretDetailActionTypeViewCheckOut SecretDetailActionType = "ViewCheckOut"

	// SecretDetailActionTypeViewApproval captures enum value "ViewApproval"
	SecretDetailActionTypeViewApproval SecretDetailActionType = "ViewApproval"

	// SecretDetailActionTypeViewListFields captures enum value "ViewListFields"
	SecretDetailActionTypeViewListFields SecretDetailActionType = "ViewListFields"

	// SecretDetailActionTypeErase captures enum value "Erase"
	SecretDetailActionTypeErase SecretDetailActionType = "Erase"

	// SecretDetailActionTypeViewMfa captures enum value "ViewMfa"
	SecretDetailActionTypeViewMfa SecretDetailActionType = "ViewMfa"
)

// for schema
var secretDetailActionTypeEnum []interface{}

func init() {
	var res []SecretDetailActionType
	if err := json.Unmarshal([]byte(`["ChangePasswordNow","ConvertTemplate","Copy","Delete","Edit","EditExpiration","EditRpc","EditSecurity","Expire","Heartbeat","EditShare","ShowSshProxyCredentials","StopChangePasswordNow","ViewAudit","ViewDependencies","ViewLaunchers","ViewExpiration","ViewHooks","ViewRpc","ViewSecurity","ViewSettings","Undelete","ForceCheckin","ViewShare","EditHooks","EditDependencies","ViewGeneralDetails","ViewHeartbeatStatus","Checkin","Checkout","GenerateOneTimePassword","ShowSshTerminalDetails","ShowRdpProxyCredentials","ViewMetadata","ViewSecretExposure","ViewCheckOut","ViewApproval","ViewListFields","Erase","ViewMfa"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		secretDetailActionTypeEnum = append(secretDetailActionTypeEnum, v)
	}
}

func (m SecretDetailActionType) validateSecretDetailActionTypeEnum(path, location string, value SecretDetailActionType) error {
	if err := validate.EnumCase(path, location, value, secretDetailActionTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this secret detail action type
func (m SecretDetailActionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSecretDetailActionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this secret detail action type based on context it is used
func (m SecretDetailActionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
