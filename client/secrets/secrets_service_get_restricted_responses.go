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

// SecretsServiceGetRestrictedReader is a Reader for the SecretsServiceGetRestricted structure.
type SecretsServiceGetRestrictedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsServiceGetRestrictedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsServiceGetRestrictedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecretsServiceGetRestrictedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecretsServiceGetRestrictedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretsServiceGetRestrictedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/secrets/{id}/restricted] SecretsService_GetRestricted", response, response.Code())
	}
}

// NewSecretsServiceGetRestrictedOK creates a SecretsServiceGetRestrictedOK with default headers values
func NewSecretsServiceGetRestrictedOK() *SecretsServiceGetRestrictedOK {
	return &SecretsServiceGetRestrictedOK{}
}

/*
SecretsServiceGetRestrictedOK describes a response with status code 200, with default header values.

Secret object
*/
type SecretsServiceGetRestrictedOK struct {
	Payload *models.SecretModel
}

// IsSuccess returns true when this secrets service get restricted o k response has a 2xx status code
func (o *SecretsServiceGetRestrictedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secrets service get restricted o k response has a 3xx status code
func (o *SecretsServiceGetRestrictedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get restricted o k response has a 4xx status code
func (o *SecretsServiceGetRestrictedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service get restricted o k response has a 5xx status code
func (o *SecretsServiceGetRestrictedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service get restricted o k response a status code equal to that given
func (o *SecretsServiceGetRestrictedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the secrets service get restricted o k response
func (o *SecretsServiceGetRestrictedOK) Code() int {
	return 200
}

func (o *SecretsServiceGetRestrictedOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/secrets/{id}/restricted][%d] secretsServiceGetRestrictedOK  %+v", 200, o.Payload)
}

func (o *SecretsServiceGetRestrictedOK) String() string {
	return fmt.Sprintf("[POST /api/v1/secrets/{id}/restricted][%d] secretsServiceGetRestrictedOK  %+v", 200, o.Payload)
}

func (o *SecretsServiceGetRestrictedOK) GetPayload() *models.SecretModel {
	return o.Payload
}

func (o *SecretsServiceGetRestrictedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceGetRestrictedBadRequest creates a SecretsServiceGetRestrictedBadRequest with default headers values
func NewSecretsServiceGetRestrictedBadRequest() *SecretsServiceGetRestrictedBadRequest {
	return &SecretsServiceGetRestrictedBadRequest{}
}

/*
SecretsServiceGetRestrictedBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SecretsServiceGetRestrictedBadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this secrets service get restricted bad request response has a 2xx status code
func (o *SecretsServiceGetRestrictedBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service get restricted bad request response has a 3xx status code
func (o *SecretsServiceGetRestrictedBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get restricted bad request response has a 4xx status code
func (o *SecretsServiceGetRestrictedBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service get restricted bad request response has a 5xx status code
func (o *SecretsServiceGetRestrictedBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service get restricted bad request response a status code equal to that given
func (o *SecretsServiceGetRestrictedBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the secrets service get restricted bad request response
func (o *SecretsServiceGetRestrictedBadRequest) Code() int {
	return 400
}

func (o *SecretsServiceGetRestrictedBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/secrets/{id}/restricted][%d] secretsServiceGetRestrictedBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceGetRestrictedBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/secrets/{id}/restricted][%d] secretsServiceGetRestrictedBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceGetRestrictedBadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *SecretsServiceGetRestrictedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceGetRestrictedForbidden creates a SecretsServiceGetRestrictedForbidden with default headers values
func NewSecretsServiceGetRestrictedForbidden() *SecretsServiceGetRestrictedForbidden {
	return &SecretsServiceGetRestrictedForbidden{}
}

/*
SecretsServiceGetRestrictedForbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type SecretsServiceGetRestrictedForbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this secrets service get restricted forbidden response has a 2xx status code
func (o *SecretsServiceGetRestrictedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service get restricted forbidden response has a 3xx status code
func (o *SecretsServiceGetRestrictedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get restricted forbidden response has a 4xx status code
func (o *SecretsServiceGetRestrictedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service get restricted forbidden response has a 5xx status code
func (o *SecretsServiceGetRestrictedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service get restricted forbidden response a status code equal to that given
func (o *SecretsServiceGetRestrictedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the secrets service get restricted forbidden response
func (o *SecretsServiceGetRestrictedForbidden) Code() int {
	return 403
}

func (o *SecretsServiceGetRestrictedForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/secrets/{id}/restricted][%d] secretsServiceGetRestrictedForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceGetRestrictedForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/secrets/{id}/restricted][%d] secretsServiceGetRestrictedForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceGetRestrictedForbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *SecretsServiceGetRestrictedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceGetRestrictedInternalServerError creates a SecretsServiceGetRestrictedInternalServerError with default headers values
func NewSecretsServiceGetRestrictedInternalServerError() *SecretsServiceGetRestrictedInternalServerError {
	return &SecretsServiceGetRestrictedInternalServerError{}
}

/*
SecretsServiceGetRestrictedInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretsServiceGetRestrictedInternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this secrets service get restricted internal server error response has a 2xx status code
func (o *SecretsServiceGetRestrictedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service get restricted internal server error response has a 3xx status code
func (o *SecretsServiceGetRestrictedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get restricted internal server error response has a 4xx status code
func (o *SecretsServiceGetRestrictedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service get restricted internal server error response has a 5xx status code
func (o *SecretsServiceGetRestrictedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secrets service get restricted internal server error response a status code equal to that given
func (o *SecretsServiceGetRestrictedInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the secrets service get restricted internal server error response
func (o *SecretsServiceGetRestrictedInternalServerError) Code() int {
	return 500
}

func (o *SecretsServiceGetRestrictedInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/secrets/{id}/restricted][%d] secretsServiceGetRestrictedInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceGetRestrictedInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/secrets/{id}/restricted][%d] secretsServiceGetRestrictedInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceGetRestrictedInternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *SecretsServiceGetRestrictedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}