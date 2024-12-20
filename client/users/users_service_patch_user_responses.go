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

// UsersServicePatchUserReader is a Reader for the UsersServicePatchUser structure.
type UsersServicePatchUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersServicePatchUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersServicePatchUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersServicePatchUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersServicePatchUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersServicePatchUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /api/v1/users/{id}] UsersService_PatchUser", response, response.Code())
	}
}

// NewUsersServicePatchUserOK creates a UsersServicePatchUserOK with default headers values
func NewUsersServicePatchUserOK() *UsersServicePatchUserOK {
	return &UsersServicePatchUserOK{}
}

/*
UsersServicePatchUserOK describes a response with status code 200, with default header values.

User Configuration
*/
type UsersServicePatchUserOK struct {
	Payload *models.UserModel
}

// IsSuccess returns true when this users service patch user o k response has a 2xx status code
func (o *UsersServicePatchUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users service patch user o k response has a 3xx status code
func (o *UsersServicePatchUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users service patch user o k response has a 4xx status code
func (o *UsersServicePatchUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users service patch user o k response has a 5xx status code
func (o *UsersServicePatchUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users service patch user o k response a status code equal to that given
func (o *UsersServicePatchUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users service patch user o k response
func (o *UsersServicePatchUserOK) Code() int {
	return 200
}

func (o *UsersServicePatchUserOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/users/{id}][%d] usersServicePatchUserOK  %+v", 200, o.Payload)
}

func (o *UsersServicePatchUserOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/users/{id}][%d] usersServicePatchUserOK  %+v", 200, o.Payload)
}

func (o *UsersServicePatchUserOK) GetPayload() *models.UserModel {
	return o.Payload
}

func (o *UsersServicePatchUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersServicePatchUserBadRequest creates a UsersServicePatchUserBadRequest with default headers values
func NewUsersServicePatchUserBadRequest() *UsersServicePatchUserBadRequest {
	return &UsersServicePatchUserBadRequest{}
}

/*
UsersServicePatchUserBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UsersServicePatchUserBadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this users service patch user bad request response has a 2xx status code
func (o *UsersServicePatchUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users service patch user bad request response has a 3xx status code
func (o *UsersServicePatchUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users service patch user bad request response has a 4xx status code
func (o *UsersServicePatchUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users service patch user bad request response has a 5xx status code
func (o *UsersServicePatchUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users service patch user bad request response a status code equal to that given
func (o *UsersServicePatchUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the users service patch user bad request response
func (o *UsersServicePatchUserBadRequest) Code() int {
	return 400
}

func (o *UsersServicePatchUserBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/users/{id}][%d] usersServicePatchUserBadRequest  %+v", 400, o.Payload)
}

func (o *UsersServicePatchUserBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v1/users/{id}][%d] usersServicePatchUserBadRequest  %+v", 400, o.Payload)
}

func (o *UsersServicePatchUserBadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *UsersServicePatchUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersServicePatchUserForbidden creates a UsersServicePatchUserForbidden with default headers values
func NewUsersServicePatchUserForbidden() *UsersServicePatchUserForbidden {
	return &UsersServicePatchUserForbidden{}
}

/*
UsersServicePatchUserForbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type UsersServicePatchUserForbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this users service patch user forbidden response has a 2xx status code
func (o *UsersServicePatchUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users service patch user forbidden response has a 3xx status code
func (o *UsersServicePatchUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users service patch user forbidden response has a 4xx status code
func (o *UsersServicePatchUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users service patch user forbidden response has a 5xx status code
func (o *UsersServicePatchUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users service patch user forbidden response a status code equal to that given
func (o *UsersServicePatchUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the users service patch user forbidden response
func (o *UsersServicePatchUserForbidden) Code() int {
	return 403
}

func (o *UsersServicePatchUserForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/users/{id}][%d] usersServicePatchUserForbidden  %+v", 403, o.Payload)
}

func (o *UsersServicePatchUserForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/users/{id}][%d] usersServicePatchUserForbidden  %+v", 403, o.Payload)
}

func (o *UsersServicePatchUserForbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *UsersServicePatchUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersServicePatchUserInternalServerError creates a UsersServicePatchUserInternalServerError with default headers values
func NewUsersServicePatchUserInternalServerError() *UsersServicePatchUserInternalServerError {
	return &UsersServicePatchUserInternalServerError{}
}

/*
UsersServicePatchUserInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UsersServicePatchUserInternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this users service patch user internal server error response has a 2xx status code
func (o *UsersServicePatchUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users service patch user internal server error response has a 3xx status code
func (o *UsersServicePatchUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users service patch user internal server error response has a 4xx status code
func (o *UsersServicePatchUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users service patch user internal server error response has a 5xx status code
func (o *UsersServicePatchUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users service patch user internal server error response a status code equal to that given
func (o *UsersServicePatchUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the users service patch user internal server error response
func (o *UsersServicePatchUserInternalServerError) Code() int {
	return 500
}

func (o *UsersServicePatchUserInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/users/{id}][%d] usersServicePatchUserInternalServerError  %+v", 500, o.Payload)
}

func (o *UsersServicePatchUserInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v1/users/{id}][%d] usersServicePatchUserInternalServerError  %+v", 500, o.Payload)
}

func (o *UsersServicePatchUserInternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *UsersServicePatchUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
