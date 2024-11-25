// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/secret-server/client-api/models"
)

// NewRolesServiceUpdatePermissionsParams creates a new RolesServiceUpdatePermissionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRolesServiceUpdatePermissionsParams() *RolesServiceUpdatePermissionsParams {
	return &RolesServiceUpdatePermissionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRolesServiceUpdatePermissionsParamsWithTimeout creates a new RolesServiceUpdatePermissionsParams object
// with the ability to set a timeout on a request.
func NewRolesServiceUpdatePermissionsParamsWithTimeout(timeout time.Duration) *RolesServiceUpdatePermissionsParams {
	return &RolesServiceUpdatePermissionsParams{
		timeout: timeout,
	}
}

// NewRolesServiceUpdatePermissionsParamsWithContext creates a new RolesServiceUpdatePermissionsParams object
// with the ability to set a context for a request.
func NewRolesServiceUpdatePermissionsParamsWithContext(ctx context.Context) *RolesServiceUpdatePermissionsParams {
	return &RolesServiceUpdatePermissionsParams{
		Context: ctx,
	}
}

// NewRolesServiceUpdatePermissionsParamsWithHTTPClient creates a new RolesServiceUpdatePermissionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRolesServiceUpdatePermissionsParamsWithHTTPClient(client *http.Client) *RolesServiceUpdatePermissionsParams {
	return &RolesServiceUpdatePermissionsParams{
		HTTPClient: client,
	}
}

/*
RolesServiceUpdatePermissionsParams contains all the parameters to send to the API endpoint

	for the roles service update permissions operation.

	Typically these are written to a http.Request.
*/
type RolesServiceUpdatePermissionsParams struct {

	/* Args.

	   args
	*/
	Args *models.RolePermissionsAssignmentRequest

	/* ID.

	   id

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the roles service update permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolesServiceUpdatePermissionsParams) WithDefaults() *RolesServiceUpdatePermissionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the roles service update permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolesServiceUpdatePermissionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the roles service update permissions params
func (o *RolesServiceUpdatePermissionsParams) WithTimeout(timeout time.Duration) *RolesServiceUpdatePermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the roles service update permissions params
func (o *RolesServiceUpdatePermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the roles service update permissions params
func (o *RolesServiceUpdatePermissionsParams) WithContext(ctx context.Context) *RolesServiceUpdatePermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the roles service update permissions params
func (o *RolesServiceUpdatePermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the roles service update permissions params
func (o *RolesServiceUpdatePermissionsParams) WithHTTPClient(client *http.Client) *RolesServiceUpdatePermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the roles service update permissions params
func (o *RolesServiceUpdatePermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the roles service update permissions params
func (o *RolesServiceUpdatePermissionsParams) WithArgs(args *models.RolePermissionsAssignmentRequest) *RolesServiceUpdatePermissionsParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the roles service update permissions params
func (o *RolesServiceUpdatePermissionsParams) SetArgs(args *models.RolePermissionsAssignmentRequest) {
	o.Args = args
}

// WithID adds the id to the roles service update permissions params
func (o *RolesServiceUpdatePermissionsParams) WithID(id int32) *RolesServiceUpdatePermissionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the roles service update permissions params
func (o *RolesServiceUpdatePermissionsParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RolesServiceUpdatePermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Args != nil {
		if err := r.SetBodyParam(o.Args); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
