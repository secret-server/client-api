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

// SecretsServiceCreateSecretReader is a Reader for the SecretsServiceCreateSecret structure.
type SecretsServiceCreateSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsServiceCreateSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsServiceCreateSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecretsServiceCreateSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecretsServiceCreateSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretsServiceCreateSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/secrets] SecretsService_CreateSecret", response, response.Code())
	}
}

// NewSecretsServiceCreateSecretOK creates a SecretsServiceCreateSecretOK with default headers values
func NewSecretsServiceCreateSecretOK() *SecretsServiceCreateSecretOK {
	return &SecretsServiceCreateSecretOK{}
}

/*
SecretsServiceCreateSecretOK describes a response with status code 200, with default header values.

Secret object
*/
type SecretsServiceCreateSecretOK struct {
	Payload *models.SecretModel
}

// IsSuccess returns true when this secrets service create secret o k response has a 2xx status code
func (o *SecretsServiceCreateSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secrets service create secret o k response has a 3xx status code
func (o *SecretsServiceCreateSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service create secret o k response has a 4xx status code
func (o *SecretsServiceCreateSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service create secret o k response has a 5xx status code
func (o *SecretsServiceCreateSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service create secret o k response a status code equal to that given
func (o *SecretsServiceCreateSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the secrets service create secret o k response
func (o *SecretsServiceCreateSecretOK) Code() int {
	return 200
}

func (o *SecretsServiceCreateSecretOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/secrets][%d] secretsServiceCreateSecretOK  %+v", 200, o.Payload)
}

func (o *SecretsServiceCreateSecretOK) String() string {
	return fmt.Sprintf("[POST /api/v1/secrets][%d] secretsServiceCreateSecretOK  %+v", 200, o.Payload)
}

func (o *SecretsServiceCreateSecretOK) GetPayload() *models.SecretModel {
	return o.Payload
}

func (o *SecretsServiceCreateSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceCreateSecretBadRequest creates a SecretsServiceCreateSecretBadRequest with default headers values
func NewSecretsServiceCreateSecretBadRequest() *SecretsServiceCreateSecretBadRequest {
	return &SecretsServiceCreateSecretBadRequest{}
}

/*
SecretsServiceCreateSecretBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SecretsServiceCreateSecretBadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this secrets service create secret bad request response has a 2xx status code
func (o *SecretsServiceCreateSecretBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service create secret bad request response has a 3xx status code
func (o *SecretsServiceCreateSecretBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service create secret bad request response has a 4xx status code
func (o *SecretsServiceCreateSecretBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service create secret bad request response has a 5xx status code
func (o *SecretsServiceCreateSecretBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service create secret bad request response a status code equal to that given
func (o *SecretsServiceCreateSecretBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the secrets service create secret bad request response
func (o *SecretsServiceCreateSecretBadRequest) Code() int {
	return 400
}

func (o *SecretsServiceCreateSecretBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/secrets][%d] secretsServiceCreateSecretBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceCreateSecretBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/secrets][%d] secretsServiceCreateSecretBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceCreateSecretBadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *SecretsServiceCreateSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceCreateSecretForbidden creates a SecretsServiceCreateSecretForbidden with default headers values
func NewSecretsServiceCreateSecretForbidden() *SecretsServiceCreateSecretForbidden {
	return &SecretsServiceCreateSecretForbidden{}
}

/*
SecretsServiceCreateSecretForbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type SecretsServiceCreateSecretForbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this secrets service create secret forbidden response has a 2xx status code
func (o *SecretsServiceCreateSecretForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service create secret forbidden response has a 3xx status code
func (o *SecretsServiceCreateSecretForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service create secret forbidden response has a 4xx status code
func (o *SecretsServiceCreateSecretForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service create secret forbidden response has a 5xx status code
func (o *SecretsServiceCreateSecretForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service create secret forbidden response a status code equal to that given
func (o *SecretsServiceCreateSecretForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the secrets service create secret forbidden response
func (o *SecretsServiceCreateSecretForbidden) Code() int {
	return 403
}

func (o *SecretsServiceCreateSecretForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/secrets][%d] secretsServiceCreateSecretForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceCreateSecretForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/secrets][%d] secretsServiceCreateSecretForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceCreateSecretForbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *SecretsServiceCreateSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceCreateSecretInternalServerError creates a SecretsServiceCreateSecretInternalServerError with default headers values
func NewSecretsServiceCreateSecretInternalServerError() *SecretsServiceCreateSecretInternalServerError {
	return &SecretsServiceCreateSecretInternalServerError{}
}

/*
SecretsServiceCreateSecretInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretsServiceCreateSecretInternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this secrets service create secret internal server error response has a 2xx status code
func (o *SecretsServiceCreateSecretInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service create secret internal server error response has a 3xx status code
func (o *SecretsServiceCreateSecretInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service create secret internal server error response has a 4xx status code
func (o *SecretsServiceCreateSecretInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service create secret internal server error response has a 5xx status code
func (o *SecretsServiceCreateSecretInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secrets service create secret internal server error response a status code equal to that given
func (o *SecretsServiceCreateSecretInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the secrets service create secret internal server error response
func (o *SecretsServiceCreateSecretInternalServerError) Code() int {
	return 500
}

func (o *SecretsServiceCreateSecretInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/secrets][%d] secretsServiceCreateSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceCreateSecretInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/secrets][%d] secretsServiceCreateSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceCreateSecretInternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *SecretsServiceCreateSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
