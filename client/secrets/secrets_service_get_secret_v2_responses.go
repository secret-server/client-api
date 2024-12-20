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

// SecretsServiceGetSecretV2Reader is a Reader for the SecretsServiceGetSecretV2 structure.
type SecretsServiceGetSecretV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsServiceGetSecretV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsServiceGetSecretV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecretsServiceGetSecretV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecretsServiceGetSecretV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretsServiceGetSecretV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/secrets/{id}] SecretsService_GetSecretV2", response, response.Code())
	}
}

// NewSecretsServiceGetSecretV2OK creates a SecretsServiceGetSecretV2OK with default headers values
func NewSecretsServiceGetSecretV2OK() *SecretsServiceGetSecretV2OK {
	return &SecretsServiceGetSecretV2OK{}
}

/*
SecretsServiceGetSecretV2OK describes a response with status code 200, with default header values.

Secret object
*/
type SecretsServiceGetSecretV2OK struct {
	Payload *models.SecretModelV2
}

// IsSuccess returns true when this secrets service get secret v2 o k response has a 2xx status code
func (o *SecretsServiceGetSecretV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secrets service get secret v2 o k response has a 3xx status code
func (o *SecretsServiceGetSecretV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get secret v2 o k response has a 4xx status code
func (o *SecretsServiceGetSecretV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service get secret v2 o k response has a 5xx status code
func (o *SecretsServiceGetSecretV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service get secret v2 o k response a status code equal to that given
func (o *SecretsServiceGetSecretV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the secrets service get secret v2 o k response
func (o *SecretsServiceGetSecretV2OK) Code() int {
	return 200
}

func (o *SecretsServiceGetSecretV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/secrets/{id}][%d] secretsServiceGetSecretV2OK  %+v", 200, o.Payload)
}

func (o *SecretsServiceGetSecretV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/secrets/{id}][%d] secretsServiceGetSecretV2OK  %+v", 200, o.Payload)
}

func (o *SecretsServiceGetSecretV2OK) GetPayload() *models.SecretModelV2 {
	return o.Payload
}

func (o *SecretsServiceGetSecretV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretModelV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceGetSecretV2BadRequest creates a SecretsServiceGetSecretV2BadRequest with default headers values
func NewSecretsServiceGetSecretV2BadRequest() *SecretsServiceGetSecretV2BadRequest {
	return &SecretsServiceGetSecretV2BadRequest{}
}

/*
SecretsServiceGetSecretV2BadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SecretsServiceGetSecretV2BadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this secrets service get secret v2 bad request response has a 2xx status code
func (o *SecretsServiceGetSecretV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service get secret v2 bad request response has a 3xx status code
func (o *SecretsServiceGetSecretV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get secret v2 bad request response has a 4xx status code
func (o *SecretsServiceGetSecretV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service get secret v2 bad request response has a 5xx status code
func (o *SecretsServiceGetSecretV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service get secret v2 bad request response a status code equal to that given
func (o *SecretsServiceGetSecretV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the secrets service get secret v2 bad request response
func (o *SecretsServiceGetSecretV2BadRequest) Code() int {
	return 400
}

func (o *SecretsServiceGetSecretV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/secrets/{id}][%d] secretsServiceGetSecretV2BadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceGetSecretV2BadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/secrets/{id}][%d] secretsServiceGetSecretV2BadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceGetSecretV2BadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *SecretsServiceGetSecretV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceGetSecretV2Forbidden creates a SecretsServiceGetSecretV2Forbidden with default headers values
func NewSecretsServiceGetSecretV2Forbidden() *SecretsServiceGetSecretV2Forbidden {
	return &SecretsServiceGetSecretV2Forbidden{}
}

/*
SecretsServiceGetSecretV2Forbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type SecretsServiceGetSecretV2Forbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this secrets service get secret v2 forbidden response has a 2xx status code
func (o *SecretsServiceGetSecretV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service get secret v2 forbidden response has a 3xx status code
func (o *SecretsServiceGetSecretV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get secret v2 forbidden response has a 4xx status code
func (o *SecretsServiceGetSecretV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service get secret v2 forbidden response has a 5xx status code
func (o *SecretsServiceGetSecretV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service get secret v2 forbidden response a status code equal to that given
func (o *SecretsServiceGetSecretV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the secrets service get secret v2 forbidden response
func (o *SecretsServiceGetSecretV2Forbidden) Code() int {
	return 403
}

func (o *SecretsServiceGetSecretV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/secrets/{id}][%d] secretsServiceGetSecretV2Forbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceGetSecretV2Forbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/secrets/{id}][%d] secretsServiceGetSecretV2Forbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceGetSecretV2Forbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *SecretsServiceGetSecretV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceGetSecretV2InternalServerError creates a SecretsServiceGetSecretV2InternalServerError with default headers values
func NewSecretsServiceGetSecretV2InternalServerError() *SecretsServiceGetSecretV2InternalServerError {
	return &SecretsServiceGetSecretV2InternalServerError{}
}

/*
SecretsServiceGetSecretV2InternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretsServiceGetSecretV2InternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this secrets service get secret v2 internal server error response has a 2xx status code
func (o *SecretsServiceGetSecretV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service get secret v2 internal server error response has a 3xx status code
func (o *SecretsServiceGetSecretV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get secret v2 internal server error response has a 4xx status code
func (o *SecretsServiceGetSecretV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service get secret v2 internal server error response has a 5xx status code
func (o *SecretsServiceGetSecretV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secrets service get secret v2 internal server error response a status code equal to that given
func (o *SecretsServiceGetSecretV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the secrets service get secret v2 internal server error response
func (o *SecretsServiceGetSecretV2InternalServerError) Code() int {
	return 500
}

func (o *SecretsServiceGetSecretV2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/secrets/{id}][%d] secretsServiceGetSecretV2InternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceGetSecretV2InternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/secrets/{id}][%d] secretsServiceGetSecretV2InternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceGetSecretV2InternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *SecretsServiceGetSecretV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
