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

// V2ListFeatureSupportLevelsReader is a Reader for the V2ListFeatureSupportLevels structure.
type V2ListFeatureSupportLevelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ListFeatureSupportLevelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2ListFeatureSupportLevelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2ListFeatureSupportLevelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2ListFeatureSupportLevelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewV2ListFeatureSupportLevelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2ListFeatureSupportLevelsOK creates a V2ListFeatureSupportLevelsOK with default headers values
func NewV2ListFeatureSupportLevelsOK() *V2ListFeatureSupportLevelsOK {
	return &V2ListFeatureSupportLevelsOK{}
}

/*V2ListFeatureSupportLevelsOK handles this case with default header values.

Success.
*/
type V2ListFeatureSupportLevelsOK struct {
	Payload models.FeatureSupportLevels
}

func (o *V2ListFeatureSupportLevelsOK) Error() string {
	return fmt.Sprintf("[GET /v2/feature-support-levels][%d] v2ListFeatureSupportLevelsOK  %+v", 200, o.Payload)
}

func (o *V2ListFeatureSupportLevelsOK) GetPayload() models.FeatureSupportLevels {
	return o.Payload
}

func (o *V2ListFeatureSupportLevelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListFeatureSupportLevelsUnauthorized creates a V2ListFeatureSupportLevelsUnauthorized with default headers values
func NewV2ListFeatureSupportLevelsUnauthorized() *V2ListFeatureSupportLevelsUnauthorized {
	return &V2ListFeatureSupportLevelsUnauthorized{}
}

/*V2ListFeatureSupportLevelsUnauthorized handles this case with default header values.

Unauthorized.
*/
type V2ListFeatureSupportLevelsUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2ListFeatureSupportLevelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/feature-support-levels][%d] v2ListFeatureSupportLevelsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ListFeatureSupportLevelsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListFeatureSupportLevelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListFeatureSupportLevelsForbidden creates a V2ListFeatureSupportLevelsForbidden with default headers values
func NewV2ListFeatureSupportLevelsForbidden() *V2ListFeatureSupportLevelsForbidden {
	return &V2ListFeatureSupportLevelsForbidden{}
}

/*V2ListFeatureSupportLevelsForbidden handles this case with default header values.

Forbidden.
*/
type V2ListFeatureSupportLevelsForbidden struct {
	Payload *models.InfraError
}

func (o *V2ListFeatureSupportLevelsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/feature-support-levels][%d] v2ListFeatureSupportLevelsForbidden  %+v", 403, o.Payload)
}

func (o *V2ListFeatureSupportLevelsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListFeatureSupportLevelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListFeatureSupportLevelsServiceUnavailable creates a V2ListFeatureSupportLevelsServiceUnavailable with default headers values
func NewV2ListFeatureSupportLevelsServiceUnavailable() *V2ListFeatureSupportLevelsServiceUnavailable {
	return &V2ListFeatureSupportLevelsServiceUnavailable{}
}

/*V2ListFeatureSupportLevelsServiceUnavailable handles this case with default header values.

Unavailable.
*/
type V2ListFeatureSupportLevelsServiceUnavailable struct {
	Payload *models.Error
}

func (o *V2ListFeatureSupportLevelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/feature-support-levels][%d] v2ListFeatureSupportLevelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2ListFeatureSupportLevelsServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListFeatureSupportLevelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
