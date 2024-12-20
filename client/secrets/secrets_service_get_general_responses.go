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

// SecretsServiceGetGeneralReader is a Reader for the SecretsServiceGetGeneral structure.
type SecretsServiceGetGeneralReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsServiceGetGeneralReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsServiceGetGeneralOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecretsServiceGetGeneralBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecretsServiceGetGeneralForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretsServiceGetGeneralInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/secrets/secret-detail/{id}/general] SecretsService_GetGeneral", response, response.Code())
	}
}

// NewSecretsServiceGetGeneralOK creates a SecretsServiceGetGeneralOK with default headers values
func NewSecretsServiceGetGeneralOK() *SecretsServiceGetGeneralOK {
	return &SecretsServiceGetGeneralOK{}
}

/*
SecretsServiceGetGeneralOK describes a response with status code 200, with default header values.

Secret Detail State View Model
*/
type SecretsServiceGetGeneralOK struct {
	Payload *models.SecretDetailGeneralModel
}

// IsSuccess returns true when this secrets service get general o k response has a 2xx status code
func (o *SecretsServiceGetGeneralOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secrets service get general o k response has a 3xx status code
func (o *SecretsServiceGetGeneralOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get general o k response has a 4xx status code
func (o *SecretsServiceGetGeneralOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service get general o k response has a 5xx status code
func (o *SecretsServiceGetGeneralOK) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service get general o k response a status code equal to that given
func (o *SecretsServiceGetGeneralOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the secrets service get general o k response
func (o *SecretsServiceGetGeneralOK) Code() int {
	return 200
}

func (o *SecretsServiceGetGeneralOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/secrets/secret-detail/{id}/general][%d] secretsServiceGetGeneralOK  %+v", 200, o.Payload)
}

func (o *SecretsServiceGetGeneralOK) String() string {
	return fmt.Sprintf("[GET /api/v1/secrets/secret-detail/{id}/general][%d] secretsServiceGetGeneralOK  %+v", 200, o.Payload)
}

func (o *SecretsServiceGetGeneralOK) GetPayload() *models.SecretDetailGeneralModel {
	return o.Payload
}

func (o *SecretsServiceGetGeneralOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretDetailGeneralModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceGetGeneralBadRequest creates a SecretsServiceGetGeneralBadRequest with default headers values
func NewSecretsServiceGetGeneralBadRequest() *SecretsServiceGetGeneralBadRequest {
	return &SecretsServiceGetGeneralBadRequest{}
}

/*
SecretsServiceGetGeneralBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SecretsServiceGetGeneralBadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this secrets service get general bad request response has a 2xx status code
func (o *SecretsServiceGetGeneralBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service get general bad request response has a 3xx status code
func (o *SecretsServiceGetGeneralBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get general bad request response has a 4xx status code
func (o *SecretsServiceGetGeneralBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service get general bad request response has a 5xx status code
func (o *SecretsServiceGetGeneralBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service get general bad request response a status code equal to that given
func (o *SecretsServiceGetGeneralBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the secrets service get general bad request response
func (o *SecretsServiceGetGeneralBadRequest) Code() int {
	return 400
}

func (o *SecretsServiceGetGeneralBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/secrets/secret-detail/{id}/general][%d] secretsServiceGetGeneralBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceGetGeneralBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/secrets/secret-detail/{id}/general][%d] secretsServiceGetGeneralBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceGetGeneralBadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *SecretsServiceGetGeneralBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceGetGeneralForbidden creates a SecretsServiceGetGeneralForbidden with default headers values
func NewSecretsServiceGetGeneralForbidden() *SecretsServiceGetGeneralForbidden {
	return &SecretsServiceGetGeneralForbidden{}
}

/*
SecretsServiceGetGeneralForbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type SecretsServiceGetGeneralForbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this secrets service get general forbidden response has a 2xx status code
func (o *SecretsServiceGetGeneralForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service get general forbidden response has a 3xx status code
func (o *SecretsServiceGetGeneralForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get general forbidden response has a 4xx status code
func (o *SecretsServiceGetGeneralForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service get general forbidden response has a 5xx status code
func (o *SecretsServiceGetGeneralForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service get general forbidden response a status code equal to that given
func (o *SecretsServiceGetGeneralForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the secrets service get general forbidden response
func (o *SecretsServiceGetGeneralForbidden) Code() int {
	return 403
}

func (o *SecretsServiceGetGeneralForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/secrets/secret-detail/{id}/general][%d] secretsServiceGetGeneralForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceGetGeneralForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/secrets/secret-detail/{id}/general][%d] secretsServiceGetGeneralForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceGetGeneralForbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *SecretsServiceGetGeneralForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceGetGeneralInternalServerError creates a SecretsServiceGetGeneralInternalServerError with default headers values
func NewSecretsServiceGetGeneralInternalServerError() *SecretsServiceGetGeneralInternalServerError {
	return &SecretsServiceGetGeneralInternalServerError{}
}

/*
SecretsServiceGetGeneralInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretsServiceGetGeneralInternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this secrets service get general internal server error response has a 2xx status code
func (o *SecretsServiceGetGeneralInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service get general internal server error response has a 3xx status code
func (o *SecretsServiceGetGeneralInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get general internal server error response has a 4xx status code
func (o *SecretsServiceGetGeneralInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service get general internal server error response has a 5xx status code
func (o *SecretsServiceGetGeneralInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secrets service get general internal server error response a status code equal to that given
func (o *SecretsServiceGetGeneralInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the secrets service get general internal server error response
func (o *SecretsServiceGetGeneralInternalServerError) Code() int {
	return 500
}

func (o *SecretsServiceGetGeneralInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/secrets/secret-detail/{id}/general][%d] secretsServiceGetGeneralInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceGetGeneralInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/secrets/secret-detail/{id}/general][%d] secretsServiceGetGeneralInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceGetGeneralInternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *SecretsServiceGetGeneralInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
