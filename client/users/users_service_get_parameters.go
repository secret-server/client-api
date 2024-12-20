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
)

// NewUsersServiceGetParams creates a new UsersServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersServiceGetParams() *UsersServiceGetParams {
	return &UsersServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersServiceGetParamsWithTimeout creates a new UsersServiceGetParams object
// with the ability to set a timeout on a request.
func NewUsersServiceGetParamsWithTimeout(timeout time.Duration) *UsersServiceGetParams {
	return &UsersServiceGetParams{
		timeout: timeout,
	}
}

// NewUsersServiceGetParamsWithContext creates a new UsersServiceGetParams object
// with the ability to set a context for a request.
func NewUsersServiceGetParamsWithContext(ctx context.Context) *UsersServiceGetParams {
	return &UsersServiceGetParams{
		Context: ctx,
	}
}

// NewUsersServiceGetParamsWithHTTPClient creates a new UsersServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersServiceGetParamsWithHTTPClient(client *http.Client) *UsersServiceGetParams {
	return &UsersServiceGetParams{
		HTTPClient: client,
	}
}

/*
UsersServiceGetParams contains all the parameters to send to the API endpoint

	for the users service get operation.

	Typically these are written to a http.Request.
*/
type UsersServiceGetParams struct {

	/* ID.

	   User ID

	   Format: int32
	*/
	ID int32

	/* IncludeInactive.

	   Whether to include inactive users in the results
	*/
	IncludeInactive *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersServiceGetParams) WithDefaults() *UsersServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users service get params
func (o *UsersServiceGetParams) WithTimeout(timeout time.Duration) *UsersServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users service get params
func (o *UsersServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users service get params
func (o *UsersServiceGetParams) WithContext(ctx context.Context) *UsersServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users service get params
func (o *UsersServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users service get params
func (o *UsersServiceGetParams) WithHTTPClient(client *http.Client) *UsersServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users service get params
func (o *UsersServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users service get params
func (o *UsersServiceGetParams) WithID(id int32) *UsersServiceGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users service get params
func (o *UsersServiceGetParams) SetID(id int32) {
	o.ID = id
}

// WithIncludeInactive adds the includeInactive to the users service get params
func (o *UsersServiceGetParams) WithIncludeInactive(includeInactive *bool) *UsersServiceGetParams {
	o.SetIncludeInactive(includeInactive)
	return o
}

// SetIncludeInactive adds the includeInactive to the users service get params
func (o *UsersServiceGetParams) SetIncludeInactive(includeInactive *bool) {
	o.IncludeInactive = includeInactive
}

// WriteToRequest writes these params to a swagger request
func (o *UsersServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.IncludeInactive != nil {

		// query param includeInactive
		var qrIncludeInactive bool

		if o.IncludeInactive != nil {
			qrIncludeInactive = *o.IncludeInactive
		}
		qIncludeInactive := swag.FormatBool(qrIncludeInactive)
		if qIncludeInactive != "" {

			if err := r.SetQueryParam("includeInactive", qIncludeInactive); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
