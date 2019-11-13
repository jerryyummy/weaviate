//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ActionsPatchHandlerFunc turns a function with the right signature into a actions patch handler
type ActionsPatchHandlerFunc func(ActionsPatchParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ActionsPatchHandlerFunc) Handle(params ActionsPatchParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ActionsPatchHandler interface for that can handle valid actions patch params
type ActionsPatchHandler interface {
	Handle(ActionsPatchParams, *models.Principal) middleware.Responder
}

// NewActionsPatch creates a new http.Handler for the actions patch operation
func NewActionsPatch(ctx *middleware.Context, handler ActionsPatchHandler) *ActionsPatch {
	return &ActionsPatch{Context: ctx, Handler: handler}
}

/*ActionsPatch swagger:route PATCH /actions/{id} actions actionsPatch

Update an Action based on its UUID (using patch semantics).

Updates an Action. This method supports json-merge style patch semantics (RFC 7396). Provided meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.

*/
type ActionsPatch struct {
	Context *middleware.Context
	Handler ActionsPatchHandler
}

func (o *ActionsPatch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewActionsPatchParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
