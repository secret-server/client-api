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

// SecretsServiceUndeleteSecretV2Reader is a Reader for the SecretsServiceUndeleteSecretV2 structure.
type SecretsServiceUndeleteSecretV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsServiceUndeleteSecretV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsServiceUndeleteSecretV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecretsServiceUndeleteSecretV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecretsServiceUndeleteSecretV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretsServiceUndeleteSecretV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v2/secrets/{id}/activate] SecretsService_UndeleteSecretV2", response, response.Code())
	}
}

// NewSecretsServiceUndeleteSecretV2OK creates a SecretsServiceUndeleteSecretV2OK with default headers values
func NewSecretsServiceUndeleteSecretV2OK() *SecretsServiceUndeleteSecretV2OK {
	return &SecretsServiceUndeleteSecretV2OK{}
}

/*
SecretsServiceUndeleteSecretV2OK describes a response with status code 200, with default header values.

Secret
*/
type SecretsServiceUndeleteSecretV2OK struct {
	Payload *models.SecretDetailGeneralModel
}

// IsSuccess returns true when this secrets service undelete secret v2 o k response has a 2xx status code
func (o *SecretsServiceUndeleteSecretV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secrets service undelete secret v2 o k response has a 3xx status code
func (o *SecretsServiceUndeleteSecretV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service undelete secret v2 o k response has a 4xx status code
func (o *SecretsServiceUndeleteSecretV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service undelete secret v2 o k response has a 5xx status code
func (o *SecretsServiceUndeleteSecretV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service undelete secret v2 o k response a status code equal to that given
func (o *SecretsServiceUndeleteSecretV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the secrets service undelete secret v2 o k response
func (o *SecretsServiceUndeleteSecretV2OK) Code() int {
	return 200
}

func (o *SecretsServiceUndeleteSecretV2OK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/secrets/{id}/activate][%d] secretsServiceUndeleteSecretV2OK  %+v", 200, o.Payload)
}

func (o *SecretsServiceUndeleteSecretV2OK) String() string {
	return fmt.Sprintf("[PUT /api/v2/secrets/{id}/activate][%d] secretsServiceUndeleteSecretV2OK  %+v", 200, o.Payload)
}

func (o *SecretsServiceUndeleteSecretV2OK) GetPayload() *models.SecretDetailGeneralModel {
	return o.Payload
}

func (o *SecretsServiceUndeleteSecretV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretDetailGeneralModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceUndeleteSecretV2BadRequest creates a SecretsServiceUndeleteSecretV2BadRequest with default headers values
func NewSecretsServiceUndeleteSecretV2BadRequest() *SecretsServiceUndeleteSecretV2BadRequest {
	return &SecretsServiceUndeleteSecretV2BadRequest{}
}

/*
SecretsServiceUndeleteSecretV2BadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SecretsServiceUndeleteSecretV2BadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this secrets service undelete secret v2 bad request response has a 2xx status code
func (o *SecretsServiceUndeleteSecretV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service undelete secret v2 bad request response has a 3xx status code
func (o *SecretsServiceUndeleteSecretV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service undelete secret v2 bad request response has a 4xx status code
func (o *SecretsServiceUndeleteSecretV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service undelete secret v2 bad request response has a 5xx status code
func (o *SecretsServiceUndeleteSecretV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service undelete secret v2 bad request response a status code equal to that given
func (o *SecretsServiceUndeleteSecretV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the secrets service undelete secret v2 bad request response
func (o *SecretsServiceUndeleteSecretV2BadRequest) Code() int {
	return 400
}

func (o *SecretsServiceUndeleteSecretV2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/secrets/{id}/activate][%d] secretsServiceUndeleteSecretV2BadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceUndeleteSecretV2BadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/secrets/{id}/activate][%d] secretsServiceUndeleteSecretV2BadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceUndeleteSecretV2BadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *SecretsServiceUndeleteSecretV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceUndeleteSecretV2Forbidden creates a SecretsServiceUndeleteSecretV2Forbidden with default headers values
func NewSecretsServiceUndeleteSecretV2Forbidden() *SecretsServiceUndeleteSecretV2Forbidden {
	return &SecretsServiceUndeleteSecretV2Forbidden{}
}

/*
SecretsServiceUndeleteSecretV2Forbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type SecretsServiceUndeleteSecretV2Forbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this secrets service undelete secret v2 forbidden response has a 2xx status code
func (o *SecretsServiceUndeleteSecretV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service undelete secret v2 forbidden response has a 3xx status code
func (o *SecretsServiceUndeleteSecretV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service undelete secret v2 forbidden response has a 4xx status code
func (o *SecretsServiceUndeleteSecretV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service undelete secret v2 forbidden response has a 5xx status code
func (o *SecretsServiceUndeleteSecretV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service undelete secret v2 forbidden response a status code equal to that given
func (o *SecretsServiceUndeleteSecretV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the secrets service undelete secret v2 forbidden response
func (o *SecretsServiceUndeleteSecretV2Forbidden) Code() int {
	return 403
}

func (o *SecretsServiceUndeleteSecretV2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/secrets/{id}/activate][%d] secretsServiceUndeleteSecretV2Forbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceUndeleteSecretV2Forbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/secrets/{id}/activate][%d] secretsServiceUndeleteSecretV2Forbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceUndeleteSecretV2Forbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *SecretsServiceUndeleteSecretV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceUndeleteSecretV2InternalServerError creates a SecretsServiceUndeleteSecretV2InternalServerError with default headers values
func NewSecretsServiceUndeleteSecretV2InternalServerError() *SecretsServiceUndeleteSecretV2InternalServerError {
	return &SecretsServiceUndeleteSecretV2InternalServerError{}
}

/*
SecretsServiceUndeleteSecretV2InternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretsServiceUndeleteSecretV2InternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this secrets service undelete secret v2 internal server error response has a 2xx status code
func (o *SecretsServiceUndeleteSecretV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service undelete secret v2 internal server error response has a 3xx status code
func (o *SecretsServiceUndeleteSecretV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service undelete secret v2 internal server error response has a 4xx status code
func (o *SecretsServiceUndeleteSecretV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service undelete secret v2 internal server error response has a 5xx status code
func (o *SecretsServiceUndeleteSecretV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secrets service undelete secret v2 internal server error response a status code equal to that given
func (o *SecretsServiceUndeleteSecretV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the secrets service undelete secret v2 internal server error response
func (o *SecretsServiceUndeleteSecretV2InternalServerError) Code() int {
	return 500
}

func (o *SecretsServiceUndeleteSecretV2InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/secrets/{id}/activate][%d] secretsServiceUndeleteSecretV2InternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceUndeleteSecretV2InternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/secrets/{id}/activate][%d] secretsServiceUndeleteSecretV2InternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceUndeleteSecretV2InternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *SecretsServiceUndeleteSecretV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}