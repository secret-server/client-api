// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/secret-server/client-api/models"
)

// RolesServiceCreateReader is a Reader for the RolesServiceCreate structure.
type RolesServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolesServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRolesServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRolesServiceCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRolesServiceCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRolesServiceCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/roles] RolesService_Create", response, response.Code())
	}
}

// NewRolesServiceCreateOK creates a RolesServiceCreateOK with default headers values
func NewRolesServiceCreateOK() *RolesServiceCreateOK {
	return &RolesServiceCreateOK{}
}

/*
RolesServiceCreateOK describes a response with status code 200, with default header values.

Role object
*/
type RolesServiceCreateOK struct {
	Payload *models.RoleModel
}

// IsSuccess returns true when this roles service create o k response has a 2xx status code
func (o *RolesServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this roles service create o k response has a 3xx status code
func (o *RolesServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles service create o k response has a 4xx status code
func (o *RolesServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this roles service create o k response has a 5xx status code
func (o *RolesServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this roles service create o k response a status code equal to that given
func (o *RolesServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the roles service create o k response
func (o *RolesServiceCreateOK) Code() int {
	return 200
}

func (o *RolesServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/roles][%d] rolesServiceCreateOK  %+v", 200, o.Payload)
}

func (o *RolesServiceCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v1/roles][%d] rolesServiceCreateOK  %+v", 200, o.Payload)
}

func (o *RolesServiceCreateOK) GetPayload() *models.RoleModel {
	return o.Payload
}

func (o *RolesServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRolesServiceCreateBadRequest creates a RolesServiceCreateBadRequest with default headers values
func NewRolesServiceCreateBadRequest() *RolesServiceCreateBadRequest {
	return &RolesServiceCreateBadRequest{}
}

/*
RolesServiceCreateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type RolesServiceCreateBadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this roles service create bad request response has a 2xx status code
func (o *RolesServiceCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this roles service create bad request response has a 3xx status code
func (o *RolesServiceCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles service create bad request response has a 4xx status code
func (o *RolesServiceCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this roles service create bad request response has a 5xx status code
func (o *RolesServiceCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this roles service create bad request response a status code equal to that given
func (o *RolesServiceCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the roles service create bad request response
func (o *RolesServiceCreateBadRequest) Code() int {
	return 400
}

func (o *RolesServiceCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/roles][%d] rolesServiceCreateBadRequest  %+v", 400, o.Payload)
}

func (o *RolesServiceCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/roles][%d] rolesServiceCreateBadRequest  %+v", 400, o.Payload)
}

func (o *RolesServiceCreateBadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *RolesServiceCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRolesServiceCreateForbidden creates a RolesServiceCreateForbidden with default headers values
func NewRolesServiceCreateForbidden() *RolesServiceCreateForbidden {
	return &RolesServiceCreateForbidden{}
}

/*
RolesServiceCreateForbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type RolesServiceCreateForbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this roles service create forbidden response has a 2xx status code
func (o *RolesServiceCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this roles service create forbidden response has a 3xx status code
func (o *RolesServiceCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles service create forbidden response has a 4xx status code
func (o *RolesServiceCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this roles service create forbidden response has a 5xx status code
func (o *RolesServiceCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this roles service create forbidden response a status code equal to that given
func (o *RolesServiceCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the roles service create forbidden response
func (o *RolesServiceCreateForbidden) Code() int {
	return 403
}

func (o *RolesServiceCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/roles][%d] rolesServiceCreateForbidden  %+v", 403, o.Payload)
}

func (o *RolesServiceCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/roles][%d] rolesServiceCreateForbidden  %+v", 403, o.Payload)
}

func (o *RolesServiceCreateForbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *RolesServiceCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRolesServiceCreateInternalServerError creates a RolesServiceCreateInternalServerError with default headers values
func NewRolesServiceCreateInternalServerError() *RolesServiceCreateInternalServerError {
	return &RolesServiceCreateInternalServerError{}
}

/*
RolesServiceCreateInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type RolesServiceCreateInternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this roles service create internal server error response has a 2xx status code
func (o *RolesServiceCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this roles service create internal server error response has a 3xx status code
func (o *RolesServiceCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles service create internal server error response has a 4xx status code
func (o *RolesServiceCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this roles service create internal server error response has a 5xx status code
func (o *RolesServiceCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this roles service create internal server error response a status code equal to that given
func (o *RolesServiceCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the roles service create internal server error response
func (o *RolesServiceCreateInternalServerError) Code() int {
	return 500
}

func (o *RolesServiceCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/roles][%d] rolesServiceCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *RolesServiceCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/roles][%d] rolesServiceCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *RolesServiceCreateInternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *RolesServiceCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
