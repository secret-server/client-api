// Code generated by go-swagger; DO NOT EDIT.

package authentication

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
)

// NewOAuth2ServiceAuthorizeParams creates a new OAuth2ServiceAuthorizeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOAuth2ServiceAuthorizeParams() *OAuth2ServiceAuthorizeParams {
	return &OAuth2ServiceAuthorizeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOAuth2ServiceAuthorizeParamsWithTimeout creates a new OAuth2ServiceAuthorizeParams object
// with the ability to set a timeout on a request.
func NewOAuth2ServiceAuthorizeParamsWithTimeout(timeout time.Duration) *OAuth2ServiceAuthorizeParams {
	return &OAuth2ServiceAuthorizeParams{
		timeout: timeout,
	}
}

// NewOAuth2ServiceAuthorizeParamsWithContext creates a new OAuth2ServiceAuthorizeParams object
// with the ability to set a context for a request.
func NewOAuth2ServiceAuthorizeParamsWithContext(ctx context.Context) *OAuth2ServiceAuthorizeParams {
	return &OAuth2ServiceAuthorizeParams{
		Context: ctx,
	}
}

// NewOAuth2ServiceAuthorizeParamsWithHTTPClient creates a new OAuth2ServiceAuthorizeParams object
// with the ability to set a custom HTTPClient for a request.
func NewOAuth2ServiceAuthorizeParamsWithHTTPClient(client *http.Client) *OAuth2ServiceAuthorizeParams {
	return &OAuth2ServiceAuthorizeParams{
		HTTPClient: client,
	}
}

/*
OAuth2ServiceAuthorizeParams contains all the parameters to send to the API endpoint

	for the o auth2 service authorize operation.

	Typically these are written to a http.Request.
*/
type OAuth2ServiceAuthorizeParams struct {

	/* GrantType.

	   Authentication grant type.  Use 'password' when authenticating, and 'refresh_token' when refreshing a token.
	*/
	GrantType string

	/* Password.

	   Secret Server authentication password.  Required when authenticating.
	*/
	Password *string

	/* RefreshToken.

	   The refresh token.  Required when refreshing a token.
	*/
	RefreshToken *string

	/* Username.

	   Secret Server authentication username.  Required when authenticating.
	*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the o auth2 service authorize params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OAuth2ServiceAuthorizeParams) WithDefaults() *OAuth2ServiceAuthorizeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the o auth2 service authorize params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OAuth2ServiceAuthorizeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) WithTimeout(timeout time.Duration) *OAuth2ServiceAuthorizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) WithContext(ctx context.Context) *OAuth2ServiceAuthorizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) WithHTTPClient(client *http.Client) *OAuth2ServiceAuthorizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGrantType adds the grantType to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) WithGrantType(grantType string) *OAuth2ServiceAuthorizeParams {
	o.SetGrantType(grantType)
	return o
}

// SetGrantType adds the grantType to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) SetGrantType(grantType string) {
	o.GrantType = grantType
}

// WithPassword adds the password to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) WithPassword(password *string) *OAuth2ServiceAuthorizeParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) SetPassword(password *string) {
	o.Password = password
}

// WithRefreshToken adds the refreshToken to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) WithRefreshToken(refreshToken *string) *OAuth2ServiceAuthorizeParams {
	o.SetRefreshToken(refreshToken)
	return o
}

// SetRefreshToken adds the refreshToken to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) SetRefreshToken(refreshToken *string) {
	o.RefreshToken = refreshToken
}

// WithUsername adds the username to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) WithUsername(username *string) *OAuth2ServiceAuthorizeParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the o auth2 service authorize params
func (o *OAuth2ServiceAuthorizeParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *OAuth2ServiceAuthorizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param grant_type
	frGrantType := o.GrantType
	fGrantType := frGrantType
	if fGrantType != "" {
		if err := r.SetFormParam("grant_type", fGrantType); err != nil {
			return err
		}
	}

	if o.Password != nil {

		// form param password
		var frPassword string
		if o.Password != nil {
			frPassword = *o.Password
		}
		fPassword := frPassword
		if fPassword != "" {
			if err := r.SetFormParam("password", fPassword); err != nil {
				return err
			}
		}
	}

	if o.RefreshToken != nil {

		// form param refresh_token
		var frRefreshToken string
		if o.RefreshToken != nil {
			frRefreshToken = *o.RefreshToken
		}
		fRefreshToken := frRefreshToken
		if fRefreshToken != "" {
			if err := r.SetFormParam("refresh_token", fRefreshToken); err != nil {
				return err
			}
		}
	}

	if o.Username != nil {

		// form param username
		var frUsername string
		if o.Username != nil {
			frUsername = *o.Username
		}
		fUsername := frUsername
		if fUsername != "" {
			if err := r.SetFormParam("username", fUsername); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
