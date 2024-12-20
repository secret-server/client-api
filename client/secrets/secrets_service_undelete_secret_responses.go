// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/secret-server/client-api/models"
)

// SecretsServiceUndeleteSecretReader is a Reader for the SecretsServiceUndeleteSecret structure.
type SecretsServiceUndeleteSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsServiceUndeleteSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsServiceUndeleteSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecretsServiceUndeleteSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecretsServiceUndeleteSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretsServiceUndeleteSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v1/secrets/{id}/activate] SecretsService_UndeleteSecret", response, response.Code())
	}
}

// NewSecretsServiceUndeleteSecretOK creates a SecretsServiceUndeleteSecretOK with default headers values
func NewSecretsServiceUndeleteSecretOK() *SecretsServiceUndeleteSecretOK {
	return &SecretsServiceUndeleteSecretOK{}
}

/*
SecretsServiceUndeleteSecretOK describes a response with status code 200, with default header values.

Secret
*/
type SecretsServiceUndeleteSecretOK struct {
	Payload *models.SecretDetailGeneralModel
}

// IsSuccess returns true when this secrets service undelete secret o k response has a 2xx status code
func (o *SecretsServiceUndeleteSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secrets service undelete secret o k response has a 3xx status code
func (o *SecretsServiceUndeleteSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service undelete secret o k response has a 4xx status code
func (o *SecretsServiceUndeleteSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service undelete secret o k response has a 5xx status code
func (o *SecretsServiceUndeleteSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service undelete secret o k response a status code equal to that given
func (o *SecretsServiceUndeleteSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the secrets service undelete secret o k response
func (o *SecretsServiceUndeleteSecretOK) Code() int {
	return 200
}

func (o *SecretsServiceUndeleteSecretOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/activate][%d] secretsServiceUndeleteSecretOK  %+v", 200, o.Payload)
}

func (o *SecretsServiceUndeleteSecretOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/activate][%d] secretsServiceUndeleteSecretOK  %+v", 200, o.Payload)
}

func (o *SecretsServiceUndeleteSecretOK) GetPayload() *models.SecretDetailGeneralModel {
	return o.Payload
}

func (o *SecretsServiceUndeleteSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretDetailGeneralModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceUndeleteSecretBadRequest creates a SecretsServiceUndeleteSecretBadRequest with default headers values
func NewSecretsServiceUndeleteSecretBadRequest() *SecretsServiceUndeleteSecretBadRequest {
	return &SecretsServiceUndeleteSecretBadRequest{}
}

/*
SecretsServiceUndeleteSecretBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SecretsServiceUndeleteSecretBadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this secrets service undelete secret bad request response has a 2xx status code
func (o *SecretsServiceUndeleteSecretBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service undelete secret bad request response has a 3xx status code
func (o *SecretsServiceUndeleteSecretBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service undelete secret bad request response has a 4xx status code
func (o *SecretsServiceUndeleteSecretBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service undelete secret bad request response has a 5xx status code
func (o *SecretsServiceUndeleteSecretBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service undelete secret bad request response a status code equal to that given
func (o *SecretsServiceUndeleteSecretBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the secrets service undelete secret bad request response
func (o *SecretsServiceUndeleteSecretBadRequest) Code() int {
	return 400
}

func (o *SecretsServiceUndeleteSecretBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/activate][%d] secretsServiceUndeleteSecretBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceUndeleteSecretBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/activate][%d] secretsServiceUndeleteSecretBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceUndeleteSecretBadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *SecretsServiceUndeleteSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceUndeleteSecretForbidden creates a SecretsServiceUndeleteSecretForbidden with default headers values
func NewSecretsServiceUndeleteSecretForbidden() *SecretsServiceUndeleteSecretForbidden {
	return &SecretsServiceUndeleteSecretForbidden{}
}

/*
SecretsServiceUndeleteSecretForbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type SecretsServiceUndeleteSecretForbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this secrets service undelete secret forbidden response has a 2xx status code
func (o *SecretsServiceUndeleteSecretForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service undelete secret forbidden response has a 3xx status code
func (o *SecretsServiceUndeleteSecretForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service undelete secret forbidden response has a 4xx status code
func (o *SecretsServiceUndeleteSecretForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service undelete secret forbidden response has a 5xx status code
func (o *SecretsServiceUndeleteSecretForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service undelete secret forbidden response a status code equal to that given
func (o *SecretsServiceUndeleteSecretForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the secrets service undelete secret forbidden response
func (o *SecretsServiceUndeleteSecretForbidden) Code() int {
	return 403
}

func (o *SecretsServiceUndeleteSecretForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/activate][%d] secretsServiceUndeleteSecretForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceUndeleteSecretForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/activate][%d] secretsServiceUndeleteSecretForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceUndeleteSecretForbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *SecretsServiceUndeleteSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceUndeleteSecretInternalServerError creates a SecretsServiceUndeleteSecretInternalServerError with default headers values
func NewSecretsServiceUndeleteSecretInternalServerError() *SecretsServiceUndeleteSecretInternalServerError {
	return &SecretsServiceUndeleteSecretInternalServerError{}
}

/*
SecretsServiceUndeleteSecretInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretsServiceUndeleteSecretInternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this secrets service undelete secret internal server error response has a 2xx status code
func (o *SecretsServiceUndeleteSecretInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service undelete secret internal server error response has a 3xx status code
func (o *SecretsServiceUndeleteSecretInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service undelete secret internal server error response has a 4xx status code
func (o *SecretsServiceUndeleteSecretInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service undelete secret internal server error response has a 5xx status code
func (o *SecretsServiceUndeleteSecretInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secrets service undelete secret internal server error response a status code equal to that given
func (o *SecretsServiceUndeleteSecretInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the secrets service undelete secret internal server error response
func (o *SecretsServiceUndeleteSecretInternalServerError) Code() int {
	return 500
}

func (o *SecretsServiceUndeleteSecretInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/activate][%d] secretsServiceUndeleteSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceUndeleteSecretInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/activate][%d] secretsServiceUndeleteSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceUndeleteSecretInternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *SecretsServiceUndeleteSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
