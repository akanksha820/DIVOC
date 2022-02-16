// Code generated by go-swagger; DO NOT EDIT.

package certification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/divoc/api/swagger_gen/models"
)

// RevokeCertificateHandlerFunc turns a function with the right signature into a revoke certificate handler
type RevokeCertificateHandlerFunc func(RevokeCertificateParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn RevokeCertificateHandlerFunc) Handle(params RevokeCertificateParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// RevokeCertificateHandler interface for that can handle valid revoke certificate params
type RevokeCertificateHandler interface {
	Handle(RevokeCertificateParams, *models.JWTClaimBody) middleware.Responder
}

// NewRevokeCertificate creates a new http.Handler for the revoke certificate operation
func NewRevokeCertificate(ctx *middleware.Context, handler RevokeCertificateHandler) *RevokeCertificate {
	return &RevokeCertificate{Context: ctx, Handler: handler}
}

/* RevokeCertificate swagger:route DELETE /v1/certificates/{preEnrollmentCode} certification revokeCertificate

Revoke certificates for given preEnrollmentCode and dose(s)

*/
type RevokeCertificate struct {
	Context *middleware.Context
	Handler RevokeCertificateHandler
}

func (o *RevokeCertificate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRevokeCertificateParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.JWTClaimBody
	if uprinc != nil {
		principal = uprinc.(*models.JWTClaimBody) // this is really a models.JWTClaimBody, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
