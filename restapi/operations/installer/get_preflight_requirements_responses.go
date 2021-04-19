// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// GetPreflightRequirementsOKCode is the HTTP code returned for type GetPreflightRequirementsOK
const GetPreflightRequirementsOKCode int = 200

/*GetPreflightRequirementsOK Success.

swagger:response getPreflightRequirementsOK
*/
type GetPreflightRequirementsOK struct {

	/*
	  In: Body
	*/
	Payload *models.PreflightHardwareRequirements `json:"body,omitempty"`
}

// NewGetPreflightRequirementsOK creates GetPreflightRequirementsOK with default headers values
func NewGetPreflightRequirementsOK() *GetPreflightRequirementsOK {

	return &GetPreflightRequirementsOK{}
}

// WithPayload adds the payload to the get preflight requirements o k response
func (o *GetPreflightRequirementsOK) WithPayload(payload *models.PreflightHardwareRequirements) *GetPreflightRequirementsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get preflight requirements o k response
func (o *GetPreflightRequirementsOK) SetPayload(payload *models.PreflightHardwareRequirements) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPreflightRequirementsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPreflightRequirementsUnauthorizedCode is the HTTP code returned for type GetPreflightRequirementsUnauthorized
const GetPreflightRequirementsUnauthorizedCode int = 401

/*GetPreflightRequirementsUnauthorized Unauthorized.

swagger:response getPreflightRequirementsUnauthorized
*/
type GetPreflightRequirementsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewGetPreflightRequirementsUnauthorized creates GetPreflightRequirementsUnauthorized with default headers values
func NewGetPreflightRequirementsUnauthorized() *GetPreflightRequirementsUnauthorized {

	return &GetPreflightRequirementsUnauthorized{}
}

// WithPayload adds the payload to the get preflight requirements unauthorized response
func (o *GetPreflightRequirementsUnauthorized) WithPayload(payload *models.InfraError) *GetPreflightRequirementsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get preflight requirements unauthorized response
func (o *GetPreflightRequirementsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPreflightRequirementsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPreflightRequirementsForbiddenCode is the HTTP code returned for type GetPreflightRequirementsForbidden
const GetPreflightRequirementsForbiddenCode int = 403

/*GetPreflightRequirementsForbidden Forbidden.

swagger:response getPreflightRequirementsForbidden
*/
type GetPreflightRequirementsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewGetPreflightRequirementsForbidden creates GetPreflightRequirementsForbidden with default headers values
func NewGetPreflightRequirementsForbidden() *GetPreflightRequirementsForbidden {

	return &GetPreflightRequirementsForbidden{}
}

// WithPayload adds the payload to the get preflight requirements forbidden response
func (o *GetPreflightRequirementsForbidden) WithPayload(payload *models.InfraError) *GetPreflightRequirementsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get preflight requirements forbidden response
func (o *GetPreflightRequirementsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPreflightRequirementsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPreflightRequirementsNotFoundCode is the HTTP code returned for type GetPreflightRequirementsNotFound
const GetPreflightRequirementsNotFoundCode int = 404

/*GetPreflightRequirementsNotFound Error.

swagger:response getPreflightRequirementsNotFound
*/
type GetPreflightRequirementsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPreflightRequirementsNotFound creates GetPreflightRequirementsNotFound with default headers values
func NewGetPreflightRequirementsNotFound() *GetPreflightRequirementsNotFound {

	return &GetPreflightRequirementsNotFound{}
}

// WithPayload adds the payload to the get preflight requirements not found response
func (o *GetPreflightRequirementsNotFound) WithPayload(payload *models.Error) *GetPreflightRequirementsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get preflight requirements not found response
func (o *GetPreflightRequirementsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPreflightRequirementsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPreflightRequirementsMethodNotAllowedCode is the HTTP code returned for type GetPreflightRequirementsMethodNotAllowed
const GetPreflightRequirementsMethodNotAllowedCode int = 405

/*GetPreflightRequirementsMethodNotAllowed Method Not Allowed.

swagger:response getPreflightRequirementsMethodNotAllowed
*/
type GetPreflightRequirementsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPreflightRequirementsMethodNotAllowed creates GetPreflightRequirementsMethodNotAllowed with default headers values
func NewGetPreflightRequirementsMethodNotAllowed() *GetPreflightRequirementsMethodNotAllowed {

	return &GetPreflightRequirementsMethodNotAllowed{}
}

// WithPayload adds the payload to the get preflight requirements method not allowed response
func (o *GetPreflightRequirementsMethodNotAllowed) WithPayload(payload *models.Error) *GetPreflightRequirementsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get preflight requirements method not allowed response
func (o *GetPreflightRequirementsMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPreflightRequirementsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPreflightRequirementsInternalServerErrorCode is the HTTP code returned for type GetPreflightRequirementsInternalServerError
const GetPreflightRequirementsInternalServerErrorCode int = 500

/*GetPreflightRequirementsInternalServerError Error.

swagger:response getPreflightRequirementsInternalServerError
*/
type GetPreflightRequirementsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPreflightRequirementsInternalServerError creates GetPreflightRequirementsInternalServerError with default headers values
func NewGetPreflightRequirementsInternalServerError() *GetPreflightRequirementsInternalServerError {

	return &GetPreflightRequirementsInternalServerError{}
}

// WithPayload adds the payload to the get preflight requirements internal server error response
func (o *GetPreflightRequirementsInternalServerError) WithPayload(payload *models.Error) *GetPreflightRequirementsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get preflight requirements internal server error response
func (o *GetPreflightRequirementsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPreflightRequirementsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
