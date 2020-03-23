// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewEnableHostParams creates a new EnableHostParams object
// with the default values initialized.
func NewEnableHostParams() *EnableHostParams {
	var ()
	return &EnableHostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnableHostParamsWithTimeout creates a new EnableHostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnableHostParamsWithTimeout(timeout time.Duration) *EnableHostParams {
	var ()
	return &EnableHostParams{

		timeout: timeout,
	}
}

// NewEnableHostParamsWithContext creates a new EnableHostParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnableHostParamsWithContext(ctx context.Context) *EnableHostParams {
	var ()
	return &EnableHostParams{

		Context: ctx,
	}
}

// NewEnableHostParamsWithHTTPClient creates a new EnableHostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnableHostParamsWithHTTPClient(client *http.Client) *EnableHostParams {
	var ()
	return &EnableHostParams{
		HTTPClient: client,
	}
}

/*EnableHostParams contains all the parameters to send to the API endpoint
for the enable host operation typically these are written to a http.Request
*/
type EnableHostParams struct {

	/*HostID
	  The ID of the host to enable

	*/
	HostID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enable host params
func (o *EnableHostParams) WithTimeout(timeout time.Duration) *EnableHostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable host params
func (o *EnableHostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable host params
func (o *EnableHostParams) WithContext(ctx context.Context) *EnableHostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable host params
func (o *EnableHostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable host params
func (o *EnableHostParams) WithHTTPClient(client *http.Client) *EnableHostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable host params
func (o *EnableHostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHostID adds the hostID to the enable host params
func (o *EnableHostParams) WithHostID(hostID strfmt.UUID) *EnableHostParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the enable host params
func (o *EnableHostParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WriteToRequest writes these params to a swagger request
func (o *EnableHostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
