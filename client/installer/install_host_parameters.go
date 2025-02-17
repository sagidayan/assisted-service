// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewInstallHostParams creates a new InstallHostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstallHostParams() *InstallHostParams {
	return &InstallHostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstallHostParamsWithTimeout creates a new InstallHostParams object
// with the ability to set a timeout on a request.
func NewInstallHostParamsWithTimeout(timeout time.Duration) *InstallHostParams {
	return &InstallHostParams{
		timeout: timeout,
	}
}

// NewInstallHostParamsWithContext creates a new InstallHostParams object
// with the ability to set a context for a request.
func NewInstallHostParamsWithContext(ctx context.Context) *InstallHostParams {
	return &InstallHostParams{
		Context: ctx,
	}
}

// NewInstallHostParamsWithHTTPClient creates a new InstallHostParams object
// with the ability to set a custom HTTPClient for a request.
func NewInstallHostParamsWithHTTPClient(client *http.Client) *InstallHostParams {
	return &InstallHostParams{
		HTTPClient: client,
	}
}

/* InstallHostParams contains all the parameters to send to the API endpoint
   for the install host operation.

   Typically these are written to a http.Request.
*/
type InstallHostParams struct {

	/* ClusterID.

	   The cluster of the host that is being installed.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	/* HostID.

	   The host that is being installed.

	   Format: uuid
	*/
	HostID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the install host params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstallHostParams) WithDefaults() *InstallHostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the install host params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstallHostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the install host params
func (o *InstallHostParams) WithTimeout(timeout time.Duration) *InstallHostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the install host params
func (o *InstallHostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the install host params
func (o *InstallHostParams) WithContext(ctx context.Context) *InstallHostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the install host params
func (o *InstallHostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the install host params
func (o *InstallHostParams) WithHTTPClient(client *http.Client) *InstallHostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the install host params
func (o *InstallHostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the install host params
func (o *InstallHostParams) WithClusterID(clusterID strfmt.UUID) *InstallHostParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the install host params
func (o *InstallHostParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithHostID adds the hostID to the install host params
func (o *InstallHostParams) WithHostID(hostID strfmt.UUID) *InstallHostParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the install host params
func (o *InstallHostParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WriteToRequest writes these params to a swagger request
func (o *InstallHostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
