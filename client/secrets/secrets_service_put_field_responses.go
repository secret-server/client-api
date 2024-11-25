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

// SecretsServicePutFieldReader is a Reader for the SecretsServicePutField structure.
type SecretsServicePutFieldReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsServicePutFieldReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsServicePutFieldOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecretsServicePutFieldBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecretsServicePutFieldForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretsServicePutFieldInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v1/secrets/{id}/fields/{slug}] SecretsService_PutField", response, response.Code())
	}
}

// NewSecretsServicePutFieldOK creates a SecretsServicePutFieldOK with default headers values
func NewSecretsServicePutFieldOK() *SecretsServicePutFieldOK {
	return &SecretsServicePutFieldOK{}
}

/*
SecretsServicePutFieldOK describes a response with status code 200, with default header values.

The updated value, or 'true' if the field is a file attachment
*/
type SecretsServicePutFieldOK struct {
	Payload string
}

// IsSuccess returns true when this secrets service put field o k response has a 2xx status code
func (o *SecretsServicePutFieldOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secrets service put field o k response has a 3xx status code
func (o *SecretsServicePutFieldOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service put field o k response has a 4xx status code
func (o *SecretsServicePutFieldOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service put field o k response has a 5xx status code
func (o *SecretsServicePutFieldOK) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service put field o k response a status code equal to that given
func (o *SecretsServicePutFieldOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the secrets service put field o k response
func (o *SecretsServicePutFieldOK) Code() int {
	return 200
}

func (o *SecretsServicePutFieldOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/fields/{slug}][%d] secretsServicePutFieldOK  %+v", 200, o.Payload)
}

func (o *SecretsServicePutFieldOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/fields/{slug}][%d] secretsServicePutFieldOK  %+v", 200, o.Payload)
}

func (o *SecretsServicePutFieldOK) GetPayload() string {
	return o.Payload
}

func (o *SecretsServicePutFieldOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServicePutFieldBadRequest creates a SecretsServicePutFieldBadRequest with default headers values
func NewSecretsServicePutFieldBadRequest() *SecretsServicePutFieldBadRequest {
	return &SecretsServicePutFieldBadRequest{}
}

/*
SecretsServicePutFieldBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SecretsServicePutFieldBadRequest struct {
	Payload *models.BadRequestResponse
}

// IsSuccess returns true when this secrets service put field bad request response has a 2xx status code
func (o *SecretsServicePutFieldBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service put field bad request response has a 3xx status code
func (o *SecretsServicePutFieldBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service put field bad request response has a 4xx status code
func (o *SecretsServicePutFieldBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service put field bad request response has a 5xx status code
func (o *SecretsServicePutFieldBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service put field bad request response a status code equal to that given
func (o *SecretsServicePutFieldBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the secrets service put field bad request response
func (o *SecretsServicePutFieldBadRequest) Code() int {
	return 400
}

func (o *SecretsServicePutFieldBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/fields/{slug}][%d] secretsServicePutFieldBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServicePutFieldBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/fields/{slug}][%d] secretsServicePutFieldBadRequest  %+v", 400, o.Payload)
}

func (o *SecretsServicePutFieldBadRequest) GetPayload() *models.BadRequestResponse {
	return o.Payload
}

func (o *SecretsServicePutFieldBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServicePutFieldForbidden creates a SecretsServicePutFieldForbidden with default headers values
func NewSecretsServicePutFieldForbidden() *SecretsServicePutFieldForbidden {
	return &SecretsServicePutFieldForbidden{}
}

/*
SecretsServicePutFieldForbidden describes a response with status code 403, with default header values.

Authentication failed
*/
type SecretsServicePutFieldForbidden struct {
	Payload *models.AuthenticationFailedResponse
}

// IsSuccess returns true when this secrets service put field forbidden response has a 2xx status code
func (o *SecretsServicePutFieldForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service put field forbidden response has a 3xx status code
func (o *SecretsServicePutFieldForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service put field forbidden response has a 4xx status code
func (o *SecretsServicePutFieldForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this secrets service put field forbidden response has a 5xx status code
func (o *SecretsServicePutFieldForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this secrets service put field forbidden response a status code equal to that given
func (o *SecretsServicePutFieldForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the secrets service put field forbidden response
func (o *SecretsServicePutFieldForbidden) Code() int {
	return 403
}

func (o *SecretsServicePutFieldForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/fields/{slug}][%d] secretsServicePutFieldForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServicePutFieldForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/fields/{slug}][%d] secretsServicePutFieldForbidden  %+v", 403, o.Payload)
}

func (o *SecretsServicePutFieldForbidden) GetPayload() *models.AuthenticationFailedResponse {
	return o.Payload
}

func (o *SecretsServicePutFieldForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationFailedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsServicePutFieldInternalServerError creates a SecretsServicePutFieldInternalServerError with default headers values
func NewSecretsServicePutFieldInternalServerError() *SecretsServicePutFieldInternalServerError {
	return &SecretsServicePutFieldInternalServerError{}
}

/*
SecretsServicePutFieldInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretsServicePutFieldInternalServerError struct {
	Payload *models.InternalServerErrorResponse
}

// IsSuccess returns true when this secrets service put field internal server error response has a 2xx status code
func (o *SecretsServicePutFieldInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secrets service put field internal server error response has a 3xx status code
func (o *SecretsServicePutFieldInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secrets service put field internal server error response has a 4xx status code
func (o *SecretsServicePutFieldInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secrets service put field internal server error response has a 5xx status code
func (o *SecretsServicePutFieldInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secrets service put field internal server error response a status code equal to that given
func (o *SecretsServicePutFieldInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the secrets service put field internal server error response
func (o *SecretsServicePutFieldInternalServerError) Code() int {
	return 500
}

func (o *SecretsServicePutFieldInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/fields/{slug}][%d] secretsServicePutFieldInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServicePutFieldInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/secrets/{id}/fields/{slug}][%d] secretsServicePutFieldInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretsServicePutFieldInternalServerError) GetPayload() *models.InternalServerErrorResponse {
	return o.Payload
}

func (o *SecretsServicePutFieldInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}