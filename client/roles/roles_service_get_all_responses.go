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

// RolesServiceGetAllReader is a Reader for the RolesServiceGetAll structure.
type RolesServiceGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolesServiceGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRolesServiceGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRolesServiceGetAllBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRolesServiceGetAllForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRolesServiceGetAllInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/roles] RolesService_GetAll", response, response.Code())
	}
}

// NewRolesServiceGetAllOK creates a RolesServiceGetAllOK with default headers values
func NewRolesServiceGetAllOK() *RolesServiceGetAllOK {
	return &RolesServiceGetAllOK{}
}

/*
RolesServiceGetAllOK describes a response with status code 200, with default header values.

Role search result object
*/
type RolesServiceGetAllOK struct {
	Payload *models.PagingOfRoleModel
}

// IsSuccess returns true when this roles service get all o k response has a 2xx status code
func (o *RolesServiceGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this roles service get all o k response has a 3xx status code
func (o *RolesServiceGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles service get all o k response has a 4xx status code
func (o *RolesServiceGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this roles service get all o k response has a 5xx status code
func (o *RolesServiceGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this roles service get all o k response a status code equal to that given
func (o *RolesServiceGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the roles service get all o k response
func (o *RolesServiceGetAllOK) Code() int {
	return 200
}

func (o *RolesServiceGetAllOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] rolesServiceGetAllOK  %+v", 200, o.Payload)
}

func (o *RolesServiceGetAllOK) String() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] rolesServiceGetAllOK  %+v", 200, o.Payload)
}

func (o *RolesServiceGetAllOK) GetPayload() *models.PagingOfRoleModel {
	return o.Payload
}

func (o *RolesServiceGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PagingOfRoleModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRolesServiceGetAllBadRequest creates a RolesServiceGetAllBadRequest with default headers values
func NewRolesServiceGetAllBadRequest() *RolesServiceGetAllBadRequest {
	return &RolesServiceGetAllBadRequest{}
}

/*
RolesServiceGetAllBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type RolesServiceGetAllBadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this roles service get all bad request response has a 2xx status code
func (o *RolesServiceGetAllBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this roles service get all bad request response has a 3xx status code
func (o *RolesServiceGetAllBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles service get all bad request response has a 4xx status code
func (o *RolesServiceGetAllBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this roles service get all bad request response has a 5xx status code
func (o *RolesServiceGetAllBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this roles service get all bad request response a status code equal to that given
func (o *RolesServiceGetAllBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the roles service get all bad request response
func (o *RolesServiceGetAllBadRequest) Code() int {
	return 400
}

func (o *RolesServiceGetAllBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] rolesServiceGetAllBadRequest  %+v", 400, o.Payload)
}

func (o *RolesServiceGetAllBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] rolesServiceGetAllBadRequest  %+v", 400, o.Payload)
}

func (o *RolesServiceGetAllBadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *RolesServiceGetAllBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRolesServiceGetAllForbidden creates a RolesServiceGetAllForbidden with default headers values
func NewRolesServiceGetAllForbidden() *RolesServiceGetAllForbidden {
	return &RolesServiceGetAllForbidden{}
}

/*
RolesServiceGetAllForbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type RolesServiceGetAllForbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this roles service get all forbidden response has a 2xx status code
func (o *RolesServiceGetAllForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this roles service get all forbidden response has a 3xx status code
func (o *RolesServiceGetAllForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles service get all forbidden response has a 4xx status code
func (o *RolesServiceGetAllForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this roles service get all forbidden response has a 5xx status code
func (o *RolesServiceGetAllForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this roles service get all forbidden response a status code equal to that given
func (o *RolesServiceGetAllForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the roles service get all forbidden response
func (o *RolesServiceGetAllForbidden) Code() int {
	return 403
}

func (o *RolesServiceGetAllForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] rolesServiceGetAllForbidden  %+v", 403, o.Payload)
}

func (o *RolesServiceGetAllForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] rolesServiceGetAllForbidden  %+v", 403, o.Payload)
}

func (o *RolesServiceGetAllForbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *RolesServiceGetAllForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRolesServiceGetAllInternalServerError creates a RolesServiceGetAllInternalServerError with default headers values
func NewRolesServiceGetAllInternalServerError() *RolesServiceGetAllInternalServerError {
	return &RolesServiceGetAllInternalServerError{}
}

/*
RolesServiceGetAllInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type RolesServiceGetAllInternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this roles service get all internal server error response has a 2xx status code
func (o *RolesServiceGetAllInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this roles service get all internal server error response has a 3xx status code
func (o *RolesServiceGetAllInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles service get all internal server error response has a 4xx status code
func (o *RolesServiceGetAllInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this roles service get all internal server error response has a 5xx status code
func (o *RolesServiceGetAllInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this roles service get all internal server error response a status code equal to that given
func (o *RolesServiceGetAllInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the roles service get all internal server error response
func (o *RolesServiceGetAllInternalServerError) Code() int {
	return 500
}

func (o *RolesServiceGetAllInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] rolesServiceGetAllInternalServerError  %+v", 500, o.Payload)
}

func (o *RolesServiceGetAllInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/roles][%d] rolesServiceGetAllInternalServerError  %+v", 500, o.Payload)
}

func (o *RolesServiceGetAllInternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *RolesServiceGetAllInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
