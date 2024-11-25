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

// NewUsersServiceGetRolesParams creates a new UsersServiceGetRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersServiceGetRolesParams() *UsersServiceGetRolesParams {
	return &UsersServiceGetRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersServiceGetRolesParamsWithTimeout creates a new UsersServiceGetRolesParams object
// with the ability to set a timeout on a request.
func NewUsersServiceGetRolesParamsWithTimeout(timeout time.Duration) *UsersServiceGetRolesParams {
	return &UsersServiceGetRolesParams{
		timeout: timeout,
	}
}

// NewUsersServiceGetRolesParamsWithContext creates a new UsersServiceGetRolesParams object
// with the ability to set a context for a request.
func NewUsersServiceGetRolesParamsWithContext(ctx context.Context) *UsersServiceGetRolesParams {
	return &UsersServiceGetRolesParams{
		Context: ctx,
	}
}

// NewUsersServiceGetRolesParamsWithHTTPClient creates a new UsersServiceGetRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersServiceGetRolesParamsWithHTTPClient(client *http.Client) *UsersServiceGetRolesParams {
	return &UsersServiceGetRolesParams{
		HTTPClient: client,
	}
}

/*
UsersServiceGetRolesParams contains all the parameters to send to the API endpoint

	for the users service get roles operation.

	Typically these are written to a http.Request.
*/
type UsersServiceGetRolesParams struct {

	/* ID.

	   id

	   Format: int32
	*/
	ID int32

	/* Skip.

	   Number of records to skip before taking results

	   Format: int32
	*/
	Skip *int32

	/* SortBy0Direction.

	   Sort direction
	*/
	SortBy0Direction *string

	/* SortBy0Name.

	   Sort field name
	*/
	SortBy0Name *string

	/* SortBy0Priority.

	   Priority index. Sorts with lower values are executed earlier

	   Format: int32
	*/
	SortBy0Priority *int32

	/* Take.

	   Maximum number of records to include in results

	   Format: int32
	*/
	Take *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users service get roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersServiceGetRolesParams) WithDefaults() *UsersServiceGetRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users service get roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersServiceGetRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users service get roles params
func (o *UsersServiceGetRolesParams) WithTimeout(timeout time.Duration) *UsersServiceGetRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users service get roles params
func (o *UsersServiceGetRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users service get roles params
func (o *UsersServiceGetRolesParams) WithContext(ctx context.Context) *UsersServiceGetRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users service get roles params
func (o *UsersServiceGetRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users service get roles params
func (o *UsersServiceGetRolesParams) WithHTTPClient(client *http.Client) *UsersServiceGetRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users service get roles params
func (o *UsersServiceGetRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users service get roles params
func (o *UsersServiceGetRolesParams) WithID(id int32) *UsersServiceGetRolesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users service get roles params
func (o *UsersServiceGetRolesParams) SetID(id int32) {
	o.ID = id
}

// WithSkip adds the skip to the users service get roles params
func (o *UsersServiceGetRolesParams) WithSkip(skip *int32) *UsersServiceGetRolesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the users service get roles params
func (o *UsersServiceGetRolesParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithSortBy0Direction adds the sortBy0Direction to the users service get roles params
func (o *UsersServiceGetRolesParams) WithSortBy0Direction(sortBy0Direction *string) *UsersServiceGetRolesParams {
	o.SetSortBy0Direction(sortBy0Direction)
	return o
}

// SetSortBy0Direction adds the sortBy0Direction to the users service get roles params
func (o *UsersServiceGetRolesParams) SetSortBy0Direction(sortBy0Direction *string) {
	o.SortBy0Direction = sortBy0Direction
}

// WithSortBy0Name adds the sortBy0Name to the users service get roles params
func (o *UsersServiceGetRolesParams) WithSortBy0Name(sortBy0Name *string) *UsersServiceGetRolesParams {
	o.SetSortBy0Name(sortBy0Name)
	return o
}

// SetSortBy0Name adds the sortBy0Name to the users service get roles params
func (o *UsersServiceGetRolesParams) SetSortBy0Name(sortBy0Name *string) {
	o.SortBy0Name = sortBy0Name
}

// WithSortBy0Priority adds the sortBy0Priority to the users service get roles params
func (o *UsersServiceGetRolesParams) WithSortBy0Priority(sortBy0Priority *int32) *UsersServiceGetRolesParams {
	o.SetSortBy0Priority(sortBy0Priority)
	return o
}

// SetSortBy0Priority adds the sortBy0Priority to the users service get roles params
func (o *UsersServiceGetRolesParams) SetSortBy0Priority(sortBy0Priority *int32) {
	o.SortBy0Priority = sortBy0Priority
}

// WithTake adds the take to the users service get roles params
func (o *UsersServiceGetRolesParams) WithTake(take *int32) *UsersServiceGetRolesParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the users service get roles params
func (o *UsersServiceGetRolesParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *UsersServiceGetRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Skip != nil {

		// query param skip
		var qrSkip int32

		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt32(qrSkip)
		if qSkip != "" {

			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}
	}

	if o.SortBy0Direction != nil {

		// query param sortBy[0].direction
		var qrSortBy0Direction string

		if o.SortBy0Direction != nil {
			qrSortBy0Direction = *o.SortBy0Direction
		}
		qSortBy0Direction := qrSortBy0Direction
		if qSortBy0Direction != "" {

			if err := r.SetQueryParam("sortBy[0].direction", qSortBy0Direction); err != nil {
				return err
			}
		}
	}

	if o.SortBy0Name != nil {

		// query param sortBy[0].name
		var qrSortBy0Name string

		if o.SortBy0Name != nil {
			qrSortBy0Name = *o.SortBy0Name
		}
		qSortBy0Name := qrSortBy0Name
		if qSortBy0Name != "" {

			if err := r.SetQueryParam("sortBy[0].name", qSortBy0Name); err != nil {
				return err
			}
		}
	}

	if o.SortBy0Priority != nil {

		// query param sortBy[0].priority
		var qrSortBy0Priority int32

		if o.SortBy0Priority != nil {
			qrSortBy0Priority = *o.SortBy0Priority
		}
		qSortBy0Priority := swag.FormatInt32(qrSortBy0Priority)
		if qSortBy0Priority != "" {

			if err := r.SetQueryParam("sortBy[0].priority", qSortBy0Priority); err != nil {
				return err
			}
		}
	}

	if o.Take != nil {

		// query param take
		var qrTake int32

		if o.Take != nil {
			qrTake = *o.Take
		}
		qTake := swag.FormatInt32(qrTake)
		if qTake != "" {

			if err := r.SetQueryParam("take", qTake); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
