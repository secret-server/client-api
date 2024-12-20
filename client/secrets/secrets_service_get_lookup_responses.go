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

// SecretsServiceGetLookupReader is a Reader for the SecretsServiceGetLookup structure.
type SecretsServiceGetLookupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsServiceGetLookupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsServiceGetLookupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecretsServiceGetLookupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecretsServiceGetLookupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretsServiceGetLookupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/secrets/lookup/{id}] SecretsService_GetLookup", response, response.Code())
	}
}

// NewSecretsServiceGetLookupOK creates a SecretsServiceGetLookupOK with default headers values
func NewSecretsServiceGetLookupOK() *SecretsServiceGetLookupOK {
	return &SecretsServiceGetLookupOK{}
}

/*
SecretsServiceGetLookupOK describes a response with status code 200, with default header values.

Secret lookup result object
*/
type SecretsServiceGetLookupOK struct {
	Payload *models.SecretLookup
}

// IsSuccess returns true when this secrets service get lookup o k response has a 2xx status code
func (o *SecretsServiceGetLookupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secrets service get lookup o k response has a 3xx status code
func (o *SecretsServiceGetLookupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get lookup o k response has a 4xx status code
func (o *SecretsServiceGetLookupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service get lookup o k response has a 5xx status code
func (o *SecretsServiceGetLookupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service get lookup o k response a status code equal to that given
func (o *SecretsServiceGetLookupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the secrets service get lookup o k response
func (o *SecretsServiceGetLookupOK) Code() int {
	return 200
}

func (o *SecretsServiceGetLookupOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/secrets/lookup/{id}][%d] secretsServiceGetLookupOK  %+v", 200, o.Payload)
}

func (o *SecretsServiceGetLookupOK) String() string {
	return fmt.Sprintf("[GET /api/v1/secrets/lookup/{id}][%d] secretsServiceGetLookupOK  %+v", 200, o.Payload)
}

func (o *SecretsServiceGetLookupOK) GetPayload() *models.SecretLookup {
	return o.Payload
}

func (o *SecretsServiceGetLookupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretLookup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceGetLookupBadRequest creates a SecretsServiceGetLookupBadRequest with default headers values
func NewSecretsServiceGetLookupBadRequest() *SecretsServiceGetLookupBadRequest {
	return &SecretsServiceGetLookupBadRequest{}
}

/*
SecretsServiceGetLookupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SecretsServiceGetLookupBadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this secrets service get lookup bad request response has a 2xx status code
func (o *SecretsServiceGetLookupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service get lookup bad request response has a 3xx status code
func (o *SecretsServiceGetLookupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get lookup bad request response has a 4xx status code
func (o *SecretsServiceGetLookupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service get lookup bad request response has a 5xx status code
func (o *SecretsServiceGetLookupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service get lookup bad request response a status code equal to that given
func (o *SecretsServiceGetLookupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the secrets service get lookup bad request response
func (o *SecretsServiceGetLookupBadRequest) Code() int {
	return 400
}

func (o *SecretsServiceGetLookupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/secrets/lookup/{id}][%d] secretsServiceGetLookupBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceGetLookupBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/secrets/lookup/{id}][%d] secretsServiceGetLookupBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServiceGetLookupBadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *SecretsServiceGetLookupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceGetLookupForbidden creates a SecretsServiceGetLookupForbidden with default headers values
func NewSecretsServiceGetLookupForbidden() *SecretsServiceGetLookupForbidden {
	return &SecretsServiceGetLookupForbidden{}
}

/*
SecretsServiceGetLookupForbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type SecretsServiceGetLookupForbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this secrets service get lookup forbidden response has a 2xx status code
func (o *SecretsServiceGetLookupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service get lookup forbidden response has a 3xx status code
func (o *SecretsServiceGetLookupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get lookup forbidden response has a 4xx status code
func (o *SecretsServiceGetLookupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service get lookup forbidden response has a 5xx status code
func (o *SecretsServiceGetLookupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service get lookup forbidden response a status code equal to that given
func (o *SecretsServiceGetLookupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the secrets service get lookup forbidden response
func (o *SecretsServiceGetLookupForbidden) Code() int {
	return 403
}

func (o *SecretsServiceGetLookupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/secrets/lookup/{id}][%d] secretsServiceGetLookupForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceGetLookupForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/secrets/lookup/{id}][%d] secretsServiceGetLookupForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServiceGetLookupForbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *SecretsServiceGetLookupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServiceGetLookupInternalServerError creates a SecretsServiceGetLookupInternalServerError with default headers values
func NewSecretsServiceGetLookupInternalServerError() *SecretsServiceGetLookupInternalServerError {
	return &SecretsServiceGetLookupInternalServerError{}
}

/*
SecretsServiceGetLookupInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretsServiceGetLookupInternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this secrets service get lookup internal server error response has a 2xx status code
func (o *SecretsServiceGetLookupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service get lookup internal server error response has a 3xx status code
func (o *SecretsServiceGetLookupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service get lookup internal server error response has a 4xx status code
func (o *SecretsServiceGetLookupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service get lookup internal server error response has a 5xx status code
func (o *SecretsServiceGetLookupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secrets service get lookup internal server error response a status code equal to that given
func (o *SecretsServiceGetLookupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the secrets service get lookup internal server error response
func (o *SecretsServiceGetLookupInternalServerError) Code() int {
	return 500
}

func (o *SecretsServiceGetLookupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/secrets/lookup/{id}][%d] secretsServiceGetLookupInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceGetLookupInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/secrets/lookup/{id}][%d] secretsServiceGetLookupInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServiceGetLookupInternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *SecretsServiceGetLookupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
