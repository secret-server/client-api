// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUsersServicePatchUserParams creates a new UsersServicePatchUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersServicePatchUserParams() *UsersServicePatchUserParams {
	return &UsersServicePatchUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersServicePatchUserParamsWithTimeout creates a new UsersServicePatchUserParams object
// with the ability to set a timeout on a request.
func NewUsersServicePatchUserParamsWithTimeout(timeout time.Duration) *UsersServicePatchUserParams {
	return &UsersServicePatchUserParams{
		timeout: timeout,
	}
}

// NewUsersServicePatchUserParamsWithContext creates a new UsersServicePatchUserParams object
// with the ability to set a context for a request.
func NewUsersServicePatchUserParamsWithContext(ctx context.Context) *UsersServicePatchUserParams {
	return &UsersServicePatchUserParams{
		Context: ctx,
	}
}

// NewUsersServicePatchUserParamsWithHTTPClient creates a new UsersServicePatchUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersServicePatchUserParamsWithHTTPClient(client *http.Client) *UsersServicePatchUserParams {
	return &UsersServicePatchUserParams{
		HTTPClient: client,
	}
}

/*
UsersServicePatchUserParams contains all the parameters to send to the API endpoint

	for the users service patch user operation.

	Typically these are written to a http.Request.
*/
type UsersServicePatchUserParams struct {

	/* ID.

	   id

	   Format: int32
	*/
	ID int32

	/* PatchModel.

	   patchModel
	*/
	PatchModel *models.PatchUserModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users service patch user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersServicePatchUserParams) WithDefaults() *UsersServicePatchUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users service patch user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersServicePatchUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users service patch user params
func (o *UsersServicePatchUserParams) WithTimeout(timeout time.Duration) *UsersServicePatchUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users service patch user params
func (o *UsersServicePatchUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users service patch user params
func (o *UsersServicePatchUserParams) WithContext(ctx context.Context) *UsersServicePatchUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users service patch user params
func (o *UsersServicePatchUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users service patch user params
func (o *UsersServicePatchUserParams) WithHTTPClient(client *http.Client) *UsersServicePatchUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users service patch user params
func (o *UsersServicePatchUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users service patch user params
func (o *UsersServicePatchUserParams) WithID(id int32) *UsersServicePatchUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users service patch user params
func (o *UsersServicePatchUserParams) SetID(id int32) {
	o.ID = id
}

// WithPatchModel adds the patchModel to the users service patch user params
func (o *UsersServicePatchUserParams) WithPatchModel(patchModel *models.PatchUserModel) *UsersServicePatchUserParams {
	o.SetPatchModel(patchModel)
	return o
}

// SetPatchModel adds the patchModel to the users service patch user params
func (o *UsersServicePatchUserParams) SetPatchModel(patchModel *models.PatchUserModel) {
	o.PatchModel = patchModel
}

// WriteToRequest writes these params to a swagger request
func (o *UsersServicePatchUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}
	if o.PatchModel != nil {
		if err := r.SetBodyParam(o.PatchModel); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
