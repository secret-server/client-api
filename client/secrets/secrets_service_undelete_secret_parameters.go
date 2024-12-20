// Code generated by go-swagger; DO NOT EDIT.

package secrets

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

// NewSecretsServiceUndeleteSecretParams creates a new SecretsServiceUndeleteSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretsServiceUndeleteSecretParams() *SecretsServiceUndeleteSecretParams {
	return &SecretsServiceUndeleteSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsServiceUndeleteSecretParamsWithTimeout creates a new SecretsServiceUndeleteSecretParams object
// with the ability to set a timeout on a request.
func NewSecretsServiceUndeleteSecretParamsWithTimeout(timeout time.Duration) *SecretsServiceUndeleteSecretParams {
	return &SecretsServiceUndeleteSecretParams{
		timeout: timeout,
	}
}

// NewSecretsServiceUndeleteSecretParamsWithContext creates a new SecretsServiceUndeleteSecretParams object
// with the ability to set a context for a request.
func NewSecretsServiceUndeleteSecretParamsWithContext(ctx context.Context) *SecretsServiceUndeleteSecretParams {
	return &SecretsServiceUndeleteSecretParams{
		Context: ctx,
	}
}

// NewSecretsServiceUndeleteSecretParamsWithHTTPClient creates a new SecretsServiceUndeleteSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretsServiceUndeleteSecretParamsWithHTTPClient(client *http.Client) *SecretsServiceUndeleteSecretParams {
	return &SecretsServiceUndeleteSecretParams{
		HTTPClient: client,
	}
}

/*
SecretsServiceUndeleteSecretParams contains all the parameters to send to the API endpoint

	for the secrets service undelete secret operation.

	Typically these are written to a http.Request.
*/
type SecretsServiceUndeleteSecretParams struct {

	/* ID.

	   id

	   Format: int32
	*/
	ID int32

	/* SecretPath.

	   A full path including folder and secret name can be passed as a query string parameter when the secret ID is set to 0.  This will lookup the secret ID by path.
	*/
	SecretPath *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secrets service undelete secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretsServiceUndeleteSecretParams) WithDefaults() *SecretsServiceUndeleteSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secrets service undelete secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretsServiceUndeleteSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secrets service undelete secret params
func (o *SecretsServiceUndeleteSecretParams) WithTimeout(timeout time.Duration) *SecretsServiceUndeleteSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets service undelete secret params
func (o *SecretsServiceUndeleteSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets service undelete secret params
func (o *SecretsServiceUndeleteSecretParams) WithContext(ctx context.Context) *SecretsServiceUndeleteSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets service undelete secret params
func (o *SecretsServiceUndeleteSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets service undelete secret params
func (o *SecretsServiceUndeleteSecretParams) WithHTTPClient(client *http.Client) *SecretsServiceUndeleteSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets service undelete secret params
func (o *SecretsServiceUndeleteSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the secrets service undelete secret params
func (o *SecretsServiceUndeleteSecretParams) WithID(id int32) *SecretsServiceUndeleteSecretParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the secrets service undelete secret params
func (o *SecretsServiceUndeleteSecretParams) SetID(id int32) {
	o.ID = id
}

// WithSecretPath adds the secretPath to the secrets service undelete secret params
func (o *SecretsServiceUndeleteSecretParams) WithSecretPath(secretPath *string) *SecretsServiceUndeleteSecretParams {
	o.SetSecretPath(secretPath)
	return o
}

// SetSecretPath adds the secretPath to the secrets service undelete secret params
func (o *SecretsServiceUndeleteSecretParams) SetSecretPath(secretPath *string) {
	o.SecretPath = secretPath
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsServiceUndeleteSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.SecretPath != nil {

		// query param secretPath
		var qrSecretPath string

		if o.SecretPath != nil {
			qrSecretPath = *o.SecretPath
		}
		qSecretPath := qrSecretPath
		if qSecretPath != "" {

			if err := r.SetQueryParam("secretPath", qSecretPath); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
