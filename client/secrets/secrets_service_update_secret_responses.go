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

// SecretsServiceUpdateSecretReader is a Reader for the SecretsServiceUpdateSecret structure.
type SecretsServiceUpdateSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsServiceUpdateSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsServiceUpdateSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecretsServiceUpdateSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecretsServiceUpdateSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretsServiceUpdateSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v1/secrets/{id}] SecretsService_UpdateSecret", response, response.Code())
	}
}

// NewSecretsServiceUpdateSecretOK creates a SecretsServiceUpdateSecretOK with default headers values
func NewSecretsServiceUpdateSecretOK() *SecretsServiceUpdateSecretOK {
	return &SecretsServiceUpdateSecretOK{}
}

/*
SecretsServiceUpdateSecretOK describes a response with status code 200, with default header values.

Secret object
*/
type SecretsServiceUpdateSecretOK struct {
	Payload *models.SecretModel
}

// IsSuccess returns true when this secrets service update secret o k response has a 2xx status code
func (o *SecretsServiceUpdateSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secrets service update secret o k response has a 3xx status code
func (o *SecretsServiceUpdateSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service update secret o k response has a 4xx status code
func (o *SecretsServiceUpdateSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service update secret o k response has a 5xx status code
func (o *SecretsServiceUpdateSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service update secret o k response a status code equal to that given
func (o *SecretsServiceUpdateSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the secrets service update secret o k response
func (o *SecretsServiceUpdateSecretOK) Code() int {
	return 200
}

func (o *SecretsServiceUpdateSecretOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}][%d] secretsServiceUpdateSecretOK  %+v", 200, o.Payload)
}

func (o *SecretsServiceUpdateSecretOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}][%d] secretsServiceUpdateSecretOK  %+v", 200, o.Payload)
}

func (o *SecretsServiceUpdateSecretOK) GetPayload() *models.SecretModel {
	return o.Payload
}

func (o *SecretsServiceUpdateSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceUpdateSecretBadRequest creates a SecretsServiceUpdateSecretBadRequest with default headers values
func NewSecretsServiceUpdateSecretBadRequest() *SecretsServiceUpdateSecretBadRequest {
	return &SecretsServiceUpdateSecretBadRequest{}
}

/*
SecretsServiceUpdateSecretBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SecretsServiceUpdateSecretBadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this secrets service update secret bad request response has a 2xx status code
func (o *SecretsServiceUpdateSecretBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service update secret bad request response has a 3xx status code
func (o *SecretsServiceUpdateSecretBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service update secret bad request response has a 4xx status code
func (o *SecretsServiceUpdateSecretBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service update secret bad request response has a 5xx status code
func (o *SecretsServiceUpdateSecretBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service update secret bad request response a status code equal to that given
func (o *SecretsServiceUpdateSecretBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the secrets service update secret bad request response
func (o *SecretsServiceUpdateSecretBadRequest) Code() int {
	return 400
}

func (o *SecretsServiceUpdateSecretBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}][%d] secretsServiceUpdateSecretBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceUpdateSecretBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}][%d] secretsServiceUpdateSecretBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceUpdateSecretBadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *SecretsServiceUpdateSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceUpdateSecretForbidden creates a SecretsServiceUpdateSecretForbidden with default headers values
func NewSecretsServiceUpdateSecretForbidden() *SecretsServiceUpdateSecretForbidden {
	return &SecretsServiceUpdateSecretForbidden{}
}

/*
SecretsServiceUpdateSecretForbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type SecretsServiceUpdateSecretForbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this secrets service update secret forbidden response has a 2xx status code
func (o *SecretsServiceUpdateSecretForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service update secret forbidden response has a 3xx status code
func (o *SecretsServiceUpdateSecretForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service update secret forbidden response has a 4xx status code
func (o *SecretsServiceUpdateSecretForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service update secret forbidden response has a 5xx status code
func (o *SecretsServiceUpdateSecretForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service update secret forbidden response a status code equal to that given
func (o *SecretsServiceUpdateSecretForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the secrets service update secret forbidden response
func (o *SecretsServiceUpdateSecretForbidden) Code() int {
	return 403
}

func (o *SecretsServiceUpdateSecretForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}][%d] secretsServiceUpdateSecretForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceUpdateSecretForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}][%d] secretsServiceUpdateSecretForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceUpdateSecretForbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *SecretsServiceUpdateSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceUpdateSecretInternalServerError creates a SecretsServiceUpdateSecretInternalServerError with default headers values
func NewSecretsServiceUpdateSecretInternalServerError() *SecretsServiceUpdateSecretInternalServerError {
	return &SecretsServiceUpdateSecretInternalServerError{}
}

/*
SecretsServiceUpdateSecretInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretsServiceUpdateSecretInternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this secrets service update secret internal server error response has a 2xx status code
func (o *SecretsServiceUpdateSecretInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service update secret internal server error response has a 3xx status code
func (o *SecretsServiceUpdateSecretInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service update secret internal server error response has a 4xx status code
func (o *SecretsServiceUpdateSecretInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service update secret internal server error response has a 5xx status code
func (o *SecretsServiceUpdateSecretInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secrets service update secret internal server error response a status code equal to that given
func (o *SecretsServiceUpdateSecretInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the secrets service update secret internal server error response
func (o *SecretsServiceUpdateSecretInternalServerError) Code() int {
	return 500
}

func (o *SecretsServiceUpdateSecretInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}][%d] secretsServiceUpdateSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceUpdateSecretInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}][%d] secretsServiceUpdateSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceUpdateSecretInternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *SecretsServiceUpdateSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
