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
)

// NewRolesServiceGetAllRolePermissionsByTypeParams creates a new RolesServiceGetAllRolePermissionsByTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRolesServiceGetAllRolePermissionsByTypeParams() *RolesServiceGetAllRolePermissionsByTypeParams {
	return &RolesServiceGetAllRolePermissionsByTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRolesServiceGetAllRolePermissionsByTypeParamsWithTimeout creates a new RolesServiceGetAllRolePermissionsByTypeParams object
// with the ability to set a timeout on a request.
func NewRolesServiceGetAllRolePermissionsByTypeParamsWithTimeout(timeout time.Duration) *RolesServiceGetAllRolePermissionsByTypeParams {
	return &RolesServiceGetAllRolePermissionsByTypeParams{
		timeout: timeout,
	}
}

// NewRolesServiceGetAllRolePermissionsByTypeParamsWithContext creates a new RolesServiceGetAllRolePermissionsByTypeParams object
// with the ability to set a context for a request.
func NewRolesServiceGetAllRolePermissionsByTypeParamsWithContext(ctx context.Context) *RolesServiceGetAllRolePermissionsByTypeParams {
	return &RolesServiceGetAllRolePermissionsByTypeParams{
		Context: ctx,
	}
}

// NewRolesServiceGetAllRolePermissionsByTypeParamsWithHTTPClient creates a new RolesServiceGetAllRolePermissionsByTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewRolesServiceGetAllRolePermissionsByTypeParamsWithHTTPClient(client *http.Client) *RolesServiceGetAllRolePermissionsByTypeParams {
	return &RolesServiceGetAllRolePermissionsByTypeParams{
		HTTPClient: client,
	}
}

/*
RolesServiceGetAllRolePermissionsByTypeParams contains all the parameters to send to the API endpoint

	for the roles service get all role permissions by type operation.

	Typically these are written to a http.Request.
*/
type RolesServiceGetAllRolePermissionsByTypeParams struct {

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

// WithDefaults hydrates default values in the roles service get all role permissions by type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolesServiceGetAllRolePermissionsByTypeParams) WithDefaults() *RolesServiceGetAllRolePermissionsByTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the roles service get all role permissions by type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolesServiceGetAllRolePermissionsByTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) WithTimeout(timeout time.Duration) *RolesServiceGetAllRolePermissionsByTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) WithContext(ctx context.Context) *RolesServiceGetAllRolePermissionsByTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) WithHTTPClient(client *http.Client) *RolesServiceGetAllRolePermissionsByTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) WithID(id int32) *RolesServiceGetAllRolePermissionsByTypeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) SetID(id int32) {
	o.ID = id
}

// WithSkip adds the skip to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) WithSkip(skip *int32) *RolesServiceGetAllRolePermissionsByTypeParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithSortBy0Direction adds the sortBy0Direction to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) WithSortBy0Direction(sortBy0Direction *string) *RolesServiceGetAllRolePermissionsByTypeParams {
	o.SetSortBy0Direction(sortBy0Direction)
	return o
}

// SetSortBy0Direction adds the sortBy0Direction to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) SetSortBy0Direction(sortBy0Direction *string) {
	o.SortBy0Direction = sortBy0Direction
}

// WithSortBy0Name adds the sortBy0Name to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) WithSortBy0Name(sortBy0Name *string) *RolesServiceGetAllRolePermissionsByTypeParams {
	o.SetSortBy0Name(sortBy0Name)
	return o
}

// SetSortBy0Name adds the sortBy0Name to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) SetSortBy0Name(sortBy0Name *string) {
	o.SortBy0Name = sortBy0Name
}

// WithSortBy0Priority adds the sortBy0Priority to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) WithSortBy0Priority(sortBy0Priority *int32) *RolesServiceGetAllRolePermissionsByTypeParams {
	o.SetSortBy0Priority(sortBy0Priority)
	return o
}

// SetSortBy0Priority adds the sortBy0Priority to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) SetSortBy0Priority(sortBy0Priority *int32) {
	o.SortBy0Priority = sortBy0Priority
}

// WithTake adds the take to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) WithTake(take *int32) *RolesServiceGetAllRolePermissionsByTypeParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the roles service get all role permissions by type params
func (o *RolesServiceGetAllRolePermissionsByTypeParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *RolesServiceGetAllRolePermissionsByTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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