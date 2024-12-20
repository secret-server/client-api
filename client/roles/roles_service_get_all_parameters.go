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

// NewRolesServiceGetAllParams creates a new RolesServiceGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRolesServiceGetAllParams() *RolesServiceGetAllParams {
	return &RolesServiceGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRolesServiceGetAllParamsWithTimeout creates a new RolesServiceGetAllParams object
// with the ability to set a timeout on a request.
func NewRolesServiceGetAllParamsWithTimeout(timeout time.Duration) *RolesServiceGetAllParams {
	return &RolesServiceGetAllParams{
		timeout: timeout,
	}
}

// NewRolesServiceGetAllParamsWithContext creates a new RolesServiceGetAllParams object
// with the ability to set a context for a request.
func NewRolesServiceGetAllParamsWithContext(ctx context.Context) *RolesServiceGetAllParams {
	return &RolesServiceGetAllParams{
		Context: ctx,
	}
}

// NewRolesServiceGetAllParamsWithHTTPClient creates a new RolesServiceGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewRolesServiceGetAllParamsWithHTTPClient(client *http.Client) *RolesServiceGetAllParams {
	return &RolesServiceGetAllParams{
		HTTPClient: client,
	}
}

/*
RolesServiceGetAllParams contains all the parameters to send to the API endpoint

	for the roles service get all operation.

	Typically these are written to a http.Request.
*/
type RolesServiceGetAllParams struct {

	/* FilterGroupID.

	   Only return roles assigned to this group id.  Will be ignored if UserId is set

	   Format: int32
	*/
	FilterGroupID *int32

	/* FilterIncludeInactive.

	   Whether to include inactive Roles in the results
	*/
	FilterIncludeInactive *bool

	/* FilterUserID.

	   Only return roles assigned to this user id.  Will supercede GroupId if set

	   Format: int32
	*/
	FilterUserID *int32

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

// WithDefaults hydrates default values in the roles service get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolesServiceGetAllParams) WithDefaults() *RolesServiceGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the roles service get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolesServiceGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the roles service get all params
func (o *RolesServiceGetAllParams) WithTimeout(timeout time.Duration) *RolesServiceGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the roles service get all params
func (o *RolesServiceGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the roles service get all params
func (o *RolesServiceGetAllParams) WithContext(ctx context.Context) *RolesServiceGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the roles service get all params
func (o *RolesServiceGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the roles service get all params
func (o *RolesServiceGetAllParams) WithHTTPClient(client *http.Client) *RolesServiceGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the roles service get all params
func (o *RolesServiceGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterGroupID adds the filterGroupID to the roles service get all params
func (o *RolesServiceGetAllParams) WithFilterGroupID(filterGroupID *int32) *RolesServiceGetAllParams {
	o.SetFilterGroupID(filterGroupID)
	return o
}

// SetFilterGroupID adds the filterGroupId to the roles service get all params
func (o *RolesServiceGetAllParams) SetFilterGroupID(filterGroupID *int32) {
	o.FilterGroupID = filterGroupID
}

// WithFilterIncludeInactive adds the filterIncludeInactive to the roles service get all params
func (o *RolesServiceGetAllParams) WithFilterIncludeInactive(filterIncludeInactive *bool) *RolesServiceGetAllParams {
	o.SetFilterIncludeInactive(filterIncludeInactive)
	return o
}

// SetFilterIncludeInactive adds the filterIncludeInactive to the roles service get all params
func (o *RolesServiceGetAllParams) SetFilterIncludeInactive(filterIncludeInactive *bool) {
	o.FilterIncludeInactive = filterIncludeInactive
}

// WithFilterUserID adds the filterUserID to the roles service get all params
func (o *RolesServiceGetAllParams) WithFilterUserID(filterUserID *int32) *RolesServiceGetAllParams {
	o.SetFilterUserID(filterUserID)
	return o
}

// SetFilterUserID adds the filterUserId to the roles service get all params
func (o *RolesServiceGetAllParams) SetFilterUserID(filterUserID *int32) {
	o.FilterUserID = filterUserID
}

// WithSkip adds the skip to the roles service get all params
func (o *RolesServiceGetAllParams) WithSkip(skip *int32) *RolesServiceGetAllParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the roles service get all params
func (o *RolesServiceGetAllParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithSortBy0Direction adds the sortBy0Direction to the roles service get all params
func (o *RolesServiceGetAllParams) WithSortBy0Direction(sortBy0Direction *string) *RolesServiceGetAllParams {
	o.SetSortBy0Direction(sortBy0Direction)
	return o
}

// SetSortBy0Direction adds the sortBy0Direction to the roles service get all params
func (o *RolesServiceGetAllParams) SetSortBy0Direction(sortBy0Direction *string) {
	o.SortBy0Direction = sortBy0Direction
}

// WithSortBy0Name adds the sortBy0Name to the roles service get all params
func (o *RolesServiceGetAllParams) WithSortBy0Name(sortBy0Name *string) *RolesServiceGetAllParams {
	o.SetSortBy0Name(sortBy0Name)
	return o
}

// SetSortBy0Name adds the sortBy0Name to the roles service get all params
func (o *RolesServiceGetAllParams) SetSortBy0Name(sortBy0Name *string) {
	o.SortBy0Name = sortBy0Name
}

// WithSortBy0Priority adds the sortBy0Priority to the roles service get all params
func (o *RolesServiceGetAllParams) WithSortBy0Priority(sortBy0Priority *int32) *RolesServiceGetAllParams {
	o.SetSortBy0Priority(sortBy0Priority)
	return o
}

// SetSortBy0Priority adds the sortBy0Priority to the roles service get all params
func (o *RolesServiceGetAllParams) SetSortBy0Priority(sortBy0Priority *int32) {
	o.SortBy0Priority = sortBy0Priority
}

// WithTake adds the take to the roles service get all params
func (o *RolesServiceGetAllParams) WithTake(take *int32) *RolesServiceGetAllParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the roles service get all params
func (o *RolesServiceGetAllParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *RolesServiceGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterGroupID != nil {

		// query param filter.groupId
		var qrFilterGroupID int32

		if o.FilterGroupID != nil {
			qrFilterGroupID = *o.FilterGroupID
		}
		qFilterGroupID := swag.FormatInt32(qrFilterGroupID)
		if qFilterGroupID != "" {

			if err := r.SetQueryParam("filter.groupId", qFilterGroupID); err != nil {
				return err
			}
		}
	}

	if o.FilterIncludeInactive != nil {

		// query param filter.includeInactive
		var qrFilterIncludeInactive bool

		if o.FilterIncludeInactive != nil {
			qrFilterIncludeInactive = *o.FilterIncludeInactive
		}
		qFilterIncludeInactive := swag.FormatBool(qrFilterIncludeInactive)
		if qFilterIncludeInactive != "" {

			if err := r.SetQueryParam("filter.includeInactive", qFilterIncludeInactive); err != nil {
				return err
			}
		}
	}

	if o.FilterUserID != nil {

		// query param filter.userId
		var qrFilterUserID int32

		if o.FilterUserID != nil {
			qrFilterUserID = *o.FilterUserID
		}
		qFilterUserID := swag.FormatInt32(qrFilterUserID)
		if qFilterUserID != "" {

			if err := r.SetQueryParam("filter.userId", qFilterUserID); err != nil {
				return err
			}
		}
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
