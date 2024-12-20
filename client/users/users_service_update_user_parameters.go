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

// NewUsersServiceUpdateUserParams creates a new UsersServiceUpdateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersServiceUpdateUserParams() *UsersServiceUpdateUserParams {
	return &UsersServiceUpdateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersServiceUpdateUserParamsWithTimeout creates a new UsersServiceUpdateUserParams object
// with the ability to set a timeout on a request.
func NewUsersServiceUpdateUserParamsWithTimeout(timeout time.Duration) *UsersServiceUpdateUserParams {
	return &UsersServiceUpdateUserParams{
		timeout: timeout,
	}
}

// NewUsersServiceUpdateUserParamsWithContext creates a new UsersServiceUpdateUserParams object
// with the ability to set a context for a request.
func NewUsersServiceUpdateUserParamsWithContext(ctx context.Context) *UsersServiceUpdateUserParams {
	return &UsersServiceUpdateUserParams{
		Context: ctx,
	}
}

// NewUsersServiceUpdateUserParamsWithHTTPClient creates a new UsersServiceUpdateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersServiceUpdateUserParamsWithHTTPClient(client *http.Client) *UsersServiceUpdateUserParams {
	return &UsersServiceUpdateUserParams{
		HTTPClient: client,
	}
}

/*
UsersServiceUpdateUserParams contains all the parameters to send to the API endpoint

	for the users service update user operation.

	Typically these are written to a http.Request.
*/
type UsersServiceUpdateUserParams struct {

	/* Args.

	   User update options
	*/
	Args *models.UserUpdateArgs

	/* ID.

	   User ID

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users service update user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersServiceUpdateUserParams) WithDefaults() *UsersServiceUpdateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users service update user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersServiceUpdateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users service update user params
func (o *UsersServiceUpdateUserParams) WithTimeout(timeout time.Duration) *UsersServiceUpdateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users service update user params
func (o *UsersServiceUpdateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users service update user params
func (o *UsersServiceUpdateUserParams) WithContext(ctx context.Context) *UsersServiceUpdateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users service update user params
func (o *UsersServiceUpdateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users service update user params
func (o *UsersServiceUpdateUserParams) WithHTTPClient(client *http.Client) *UsersServiceUpdateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users service update user params
func (o *UsersServiceUpdateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the users service update user params
func (o *UsersServiceUpdateUserParams) WithArgs(args *models.UserUpdateArgs) *UsersServiceUpdateUserParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the users service update user params
func (o *UsersServiceUpdateUserParams) SetArgs(args *models.UserUpdateArgs) {
	o.Args = args
}

// WithID adds the id to the users service update user params
func (o *UsersServiceUpdateUserParams) WithID(id int32) *UsersServiceUpdateUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users service update user params
func (o *UsersServiceUpdateUserParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersServiceUpdateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
