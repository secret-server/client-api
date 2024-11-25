// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new roles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for roles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	RolesServiceCreate(params *RolesServiceCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceCreateOK, error)

	RolesServiceGet(params *RolesServiceGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceGetOK, error)

	RolesServiceGetAll(params *RolesServiceGetAllParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceGetAllOK, error)

	RolesServiceGetAllRolePermissionsByType(params *RolesServiceGetAllRolePermissionsByTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceGetAllRolePermissionsByTypeOK, error)

	RolesServiceGetRoleGroups(params *RolesServiceGetRoleGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceGetRoleGroupsOK, error)

	RolesServiceGetRolePermissions(params *RolesServiceGetRolePermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceGetRolePermissionsOK, error)

	RolesServicePatchGroups(params *RolesServicePatchGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServicePatchGroupsOK, error)

	RolesServiceStub(params *RolesServiceStubParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceStubOK, error)

	RolesServiceUpdate(params *RolesServiceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceUpdateOK, error)

	RolesServiceUpdatePermissions(params *RolesServiceUpdatePermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceUpdatePermissionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
RolesServiceCreate creates role

Create a new Role
*/
func (a *Client) RolesServiceCreate(params *RolesServiceCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRolesServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RolesService_Create",
		Method:             "POST",
		PathPattern:        "/api/v1/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RolesServiceCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RolesServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RolesService_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RolesServiceGet gets role

Get Role by Role ID
*/
func (a *Client) RolesServiceGet(params *RolesServiceGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRolesServiceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RolesService_Get",
		Method:             "GET",
		PathPattern:        "/api/v1/roles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RolesServiceGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RolesServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RolesService_Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RolesServiceGetAll searches roles

Search, filter, sort, and page Roles
*/
func (a *Client) RolesServiceGetAll(params *RolesServiceGetAllParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceGetAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRolesServiceGetAllParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RolesService_GetAll",
		Method:             "GET",
		PathPattern:        "/api/v1/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RolesServiceGetAllReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RolesServiceGetAllOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RolesService_GetAll: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RolesServiceGetAllRolePermissionsByType gets unassigned role permissions

Get unassigned Role Permissions matching the type of a specific Role by Role ID
*/
func (a *Client) RolesServiceGetAllRolePermissionsByType(params *RolesServiceGetAllRolePermissionsByTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceGetAllRolePermissionsByTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRolesServiceGetAllRolePermissionsByTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RolesService_GetAllRolePermissionsByType",
		Method:             "GET",
		PathPattern:        "/api/v1/roles/{id}/permissions/unassigned",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RolesServiceGetAllRolePermissionsByTypeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RolesServiceGetAllRolePermissionsByTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RolesService_GetAllRolePermissionsByType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RolesServiceGetRoleGroups gets role groups

Get assigned Groups by RoleId
*/
func (a *Client) RolesServiceGetRoleGroups(params *RolesServiceGetRoleGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceGetRoleGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRolesServiceGetRoleGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RolesService_GetRoleGroups",
		Method:             "GET",
		PathPattern:        "/api/v1/roles/{id}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RolesServiceGetRoleGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RolesServiceGetRoleGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RolesService_GetRoleGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RolesServiceGetRolePermissions gets assigned role permissions

Get Permissions assigned to a single Role by Role ID
*/
func (a *Client) RolesServiceGetRolePermissions(params *RolesServiceGetRolePermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceGetRolePermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRolesServiceGetRolePermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RolesService_GetRolePermissions",
		Method:             "GET",
		PathPattern:        "/api/v1/roles/{id}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RolesServiceGetRolePermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RolesServiceGetRolePermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RolesService_GetRolePermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RolesServicePatchGroups patches role group assignments

Update Groups assigned to a Role by sending list(s) of Group IDs to add/remove
*/
func (a *Client) RolesServicePatchGroups(params *RolesServicePatchGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServicePatchGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRolesServicePatchGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RolesService_PatchGroups",
		Method:             "PATCH",
		PathPattern:        "/api/v1/roles/{roleId}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RolesServicePatchGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RolesServicePatchGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RolesService_PatchGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RolesServiceStub gets role stub

Return the default values for a new Role
*/
func (a *Client) RolesServiceStub(params *RolesServiceStubParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceStubOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRolesServiceStubParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RolesService_Stub",
		Method:             "GET",
		PathPattern:        "/api/v1/roles/stub",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RolesServiceStubReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RolesServiceStubOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RolesService_Stub: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RolesServiceUpdate updates role

Update a single Role by ID
*/
func (a *Client) RolesServiceUpdate(params *RolesServiceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRolesServiceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RolesService_Update",
		Method:             "PATCH",
		PathPattern:        "/api/v1/roles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RolesServiceUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RolesServiceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RolesService_Update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RolesServiceUpdatePermissions updates role permission assignments

Update all Permissions assigned to Role
*/
func (a *Client) RolesServiceUpdatePermissions(params *RolesServiceUpdatePermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RolesServiceUpdatePermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRolesServiceUpdatePermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RolesService_UpdatePermissions",
		Method:             "PUT",
		PathPattern:        "/api/v1/roles/{id}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RolesServiceUpdatePermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RolesServiceUpdatePermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RolesService_UpdatePermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
