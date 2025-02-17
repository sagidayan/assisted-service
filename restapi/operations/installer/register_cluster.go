// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RegisterClusterHandlerFunc turns a function with the right signature into a register cluster handler
type RegisterClusterHandlerFunc func(RegisterClusterParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterClusterHandlerFunc) Handle(params RegisterClusterParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// RegisterClusterHandler interface for that can handle valid register cluster params
type RegisterClusterHandler interface {
	Handle(RegisterClusterParams, interface{}) middleware.Responder
}

// NewRegisterCluster creates a new http.Handler for the register cluster operation
func NewRegisterCluster(ctx *middleware.Context, handler RegisterClusterHandler) *RegisterCluster {
	return &RegisterCluster{Context: ctx, Handler: handler}
}

/* RegisterCluster swagger:route POST /v1/clusters installer registerCluster

Creates a new OpenShift cluster definition.

*/
type RegisterCluster struct {
	Context *middleware.Context
	Handler RegisterClusterHandler
}

func (o *RegisterCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRegisterClusterParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
