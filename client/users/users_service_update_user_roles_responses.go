// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/secret-server/client-api/models"
)

// UsersServiceUpdateUserRolesReader is a Reader for the UsersServiceUpdateUserRoles structure.
type UsersServiceUpdateUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersServiceUpdateUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersServiceUpdateUserRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersServiceUpdateUserRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersServiceUpdateUserRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersServiceUpdateUserRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v1/users/{id}/roles] UsersService_UpdateUserRoles", response, response.Code())
	}
}

// NewUsersServiceUpdateUserRolesOK creates a UsersServiceUpdateUserRolesOK with default headers values
func NewUsersServiceUpdateUserRolesOK() *UsersServiceUpdateUserRolesOK {
	return &UsersServiceUpdateUserRolesOK{}
}

/*
UsersServiceUpdateUserRolesOK describes a response with status code 200, with default header values.

Success / Fail
*/
type UsersServiceUpdateUserRolesOK struct {
	Payload *models.RoleChangeStatusModel
}

// IsSuccess returns true when this users service update user roles o k response has a 2xx status code
func (o *UsersServiceUpdateUserRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users service update user roles o k response has a 3xx status code
func (o *UsersServiceUpdateUserRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users service update user roles o k response has a 4xx status code
func (o *UsersServiceUpdateUserRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users service update user roles o k response has a 5xx status code
func (o *UsersServiceUpdateUserRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users service update user roles o k response a status code equal to that given
func (o *UsersServiceUpdateUserRolesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users service update user roles o k response
func (o *UsersServiceUpdateUserRolesOK) Code() int {
	return 200
}

func (o *UsersServiceUpdateUserRolesOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/users/{id}/roles][%d] usersServiceUpdateUserRolesOK  %+v", 200, o.Payload)
}

func (o *UsersServiceUpdateUserRolesOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/users/{id}/roles][%d] usersServiceUpdateUserRolesOK  %+v", 200, o.Payload)
}

func (o *UsersServiceUpdateUserRolesOK) GetPayload() *models.RoleChangeStatusModel {
	return o.Payload
}

func (o *UsersServiceUpdateUserRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleChangeStatusModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersServiceUpdateUserRolesBadRequest creates a UsersServiceUpdateUserRolesBadRequest with default headers values
func NewUsersServiceUpdateUserRolesBadRequest() *UsersServiceUpdateUserRolesBadRequest {
	return &UsersServiceUpdateUserRolesBadRequest{}
}

/*
UsersServiceUpdateUserRolesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UsersServiceUpdateUserRolesBadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this users service update user roles bad request response has a 2xx status code
func (o *UsersServiceUpdateUserRolesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users service update user roles bad request response has a 3xx status code
func (o *UsersServiceUpdateUserRolesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users service update user roles bad request response has a 4xx status code
func (o *UsersServiceUpdateUserRolesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users service update user roles bad request response has a 5xx status code
func (o *UsersServiceUpdateUserRolesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users service update user roles bad request response a status code equal to that given
func (o *UsersServiceUpdateUserRolesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the users service update user roles bad request response
func (o *UsersServiceUpdateUserRolesBadRequest) Code() int {
	return 400
}

func (o *UsersServiceUpdateUserRolesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/users/{id}/roles][%d] usersServiceUpdateUserRolesBadRequest  %+v", 400, o.Payload)
}

func (o *UsersServiceUpdateUserRolesBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/users/{id}/roles][%d] usersServiceUpdateUserRolesBadRequest  %+v", 400, o.Payload)
}

func (o *UsersServiceUpdateUserRolesBadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *UsersServiceUpdateUserRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersServiceUpdateUserRolesForbidden creates a UsersServiceUpdateUserRolesForbidden with default headers values
func NewUsersServiceUpdateUserRolesForbidden() *UsersServiceUpdateUserRolesForbidden {
	return &UsersServiceUpdateUserRolesForbidden{}
}

/*
UsersServiceUpdateUserRolesForbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type UsersServiceUpdateUserRolesForbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this users service update user roles forbidden response has a 2xx status code
func (o *UsersServiceUpdateUserRolesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users service update user roles forbidden response has a 3xx status code
func (o *UsersServiceUpdateUserRolesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users service update user roles forbidden response has a 4xx status code
func (o *UsersServiceUpdateUserRolesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users service update user roles forbidden response has a 5xx status code
func (o *UsersServiceUpdateUserRolesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users service update user roles forbidden response a status code equal to that given
func (o *UsersServiceUpdateUserRolesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the users service update user roles forbidden response
func (o *UsersServiceUpdateUserRolesForbidden) Code() int {
	return 403
}

func (o *UsersServiceUpdateUserRolesForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/users/{id}/roles][%d] usersServiceUpdateUserRolesForbidden  %+v", 403, o.Payload)
}

func (o *UsersServiceUpdateUserRolesForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/users/{id}/roles][%d] usersServiceUpdateUserRolesForbidden  %+v", 403, o.Payload)
}

func (o *UsersServiceUpdateUserRolesForbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *UsersServiceUpdateUserRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersServiceUpdateUserRolesInternalServerError creates a UsersServiceUpdateUserRolesInternalServerError with default headers values
func NewUsersServiceUpdateUserRolesInternalServerError() *UsersServiceUpdateUserRolesInternalServerError {
	return &UsersServiceUpdateUserRolesInternalServerError{}
}

/*
UsersServiceUpdateUserRolesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UsersServiceUpdateUserRolesInternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this users service update user roles internal server error response has a 2xx status code
func (o *UsersServiceUpdateUserRolesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users service update user roles internal server error response has a 3xx status code
func (o *UsersServiceUpdateUserRolesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users service update user roles internal server error response has a 4xx status code
func (o *UsersServiceUpdateUserRolesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users service update user roles internal server error response has a 5xx status code
func (o *UsersServiceUpdateUserRolesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users service update user roles internal server error response a status code equal to that given
func (o *UsersServiceUpdateUserRolesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the users service update user roles internal server error response
func (o *UsersServiceUpdateUserRolesInternalServerError) Code() int {
	return 500
}

func (o *UsersServiceUpdateUserRolesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/users/{id}/roles][%d] usersServiceUpdateUserRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *UsersServiceUpdateUserRolesInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/users/{id}/roles][%d] usersServiceUpdateUserRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *UsersServiceUpdateUserRolesInternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *UsersServiceUpdateUserRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
