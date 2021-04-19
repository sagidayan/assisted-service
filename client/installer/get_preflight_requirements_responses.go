// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// GetPreflightRequirementsReader is a Reader for the GetPreflightRequirements structure.
type GetPreflightRequirementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPreflightRequirementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPreflightRequirementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPreflightRequirementsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPreflightRequirementsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPreflightRequirementsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetPreflightRequirementsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPreflightRequirementsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPreflightRequirementsOK creates a GetPreflightRequirementsOK with default headers values
func NewGetPreflightRequirementsOK() *GetPreflightRequirementsOK {
	return &GetPreflightRequirementsOK{}
}

/*GetPreflightRequirementsOK handles this case with default header values.

Success.
*/
type GetPreflightRequirementsOK struct {
	Payload *models.PreflightHardwareRequirements
}

func (o *GetPreflightRequirementsOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/preflight-requirements][%d] getPreflightRequirementsOK  %+v", 200, o.Payload)
}

func (o *GetPreflightRequirementsOK) GetPayload() *models.PreflightHardwareRequirements {
	return o.Payload
}

func (o *GetPreflightRequirementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PreflightHardwareRequirements)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreflightRequirementsUnauthorized creates a GetPreflightRequirementsUnauthorized with default headers values
func NewGetPreflightRequirementsUnauthorized() *GetPreflightRequirementsUnauthorized {
	return &GetPreflightRequirementsUnauthorized{}
}

/*GetPreflightRequirementsUnauthorized handles this case with default header values.

Unauthorized.
*/
type GetPreflightRequirementsUnauthorized struct {
	Payload *models.InfraError
}

func (o *GetPreflightRequirementsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/preflight-requirements][%d] getPreflightRequirementsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPreflightRequirementsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetPreflightRequirementsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreflightRequirementsForbidden creates a GetPreflightRequirementsForbidden with default headers values
func NewGetPreflightRequirementsForbidden() *GetPreflightRequirementsForbidden {
	return &GetPreflightRequirementsForbidden{}
}

/*GetPreflightRequirementsForbidden handles this case with default header values.

Forbidden.
*/
type GetPreflightRequirementsForbidden struct {
	Payload *models.InfraError
}

func (o *GetPreflightRequirementsForbidden) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/preflight-requirements][%d] getPreflightRequirementsForbidden  %+v", 403, o.Payload)
}

func (o *GetPreflightRequirementsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetPreflightRequirementsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreflightRequirementsNotFound creates a GetPreflightRequirementsNotFound with default headers values
func NewGetPreflightRequirementsNotFound() *GetPreflightRequirementsNotFound {
	return &GetPreflightRequirementsNotFound{}
}

/*GetPreflightRequirementsNotFound handles this case with default header values.

Error.
*/
type GetPreflightRequirementsNotFound struct {
	Payload *models.Error
}

func (o *GetPreflightRequirementsNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/preflight-requirements][%d] getPreflightRequirementsNotFound  %+v", 404, o.Payload)
}

func (o *GetPreflightRequirementsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPreflightRequirementsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreflightRequirementsMethodNotAllowed creates a GetPreflightRequirementsMethodNotAllowed with default headers values
func NewGetPreflightRequirementsMethodNotAllowed() *GetPreflightRequirementsMethodNotAllowed {
	return &GetPreflightRequirementsMethodNotAllowed{}
}

/*GetPreflightRequirementsMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type GetPreflightRequirementsMethodNotAllowed struct {
	Payload *models.Error
}

func (o *GetPreflightRequirementsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/preflight-requirements][%d] getPreflightRequirementsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetPreflightRequirementsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPreflightRequirementsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreflightRequirementsInternalServerError creates a GetPreflightRequirementsInternalServerError with default headers values
func NewGetPreflightRequirementsInternalServerError() *GetPreflightRequirementsInternalServerError {
	return &GetPreflightRequirementsInternalServerError{}
}

/*GetPreflightRequirementsInternalServerError handles this case with default header values.

Error.
*/
type GetPreflightRequirementsInternalServerError struct {
	Payload *models.Error
}

func (o *GetPreflightRequirementsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/preflight-requirements][%d] getPreflightRequirementsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPreflightRequirementsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPreflightRequirementsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
