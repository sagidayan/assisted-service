// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DisableHostNoContentCode is the HTTP code returned for type DisableHostNoContent
const DisableHostNoContentCode int = 204

/*DisableHostNoContent Disabled

swagger:response disableHostNoContent
*/
type DisableHostNoContent struct {
}

// NewDisableHostNoContent creates DisableHostNoContent with default headers values
func NewDisableHostNoContent() *DisableHostNoContent {

	return &DisableHostNoContent{}
}

// WriteResponse to the client
func (o *DisableHostNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DisableHostNotFoundCode is the HTTP code returned for type DisableHostNotFound
const DisableHostNotFoundCode int = 404

/*DisableHostNotFound Host not found

swagger:response disableHostNotFound
*/
type DisableHostNotFound struct {
}

// NewDisableHostNotFound creates DisableHostNotFound with default headers values
func NewDisableHostNotFound() *DisableHostNotFound {

	return &DisableHostNotFound{}
}

// WriteResponse to the client
func (o *DisableHostNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DisableHostInternalServerErrorCode is the HTTP code returned for type DisableHostInternalServerError
const DisableHostInternalServerErrorCode int = 500

/*DisableHostInternalServerError Internal server error

swagger:response disableHostInternalServerError
*/
type DisableHostInternalServerError struct {
}

// NewDisableHostInternalServerError creates DisableHostInternalServerError with default headers values
func NewDisableHostInternalServerError() *DisableHostInternalServerError {

	return &DisableHostInternalServerError{}
}

// WriteResponse to the client
func (o *DisableHostInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
