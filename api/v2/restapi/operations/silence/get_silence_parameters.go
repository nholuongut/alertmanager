// Code generated by go-swagger; DO NOT EDIT.

// Copyright Nho luong DevOps
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package silence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetSilenceParams creates a new GetSilenceParams object
//
// There are no default values defined in the spec.
func NewGetSilenceParams() GetSilenceParams {

	return GetSilenceParams{}
}

// GetSilenceParams contains all the bound params for the get silence operation
// typically these are obtained from a http.Request
//
// swagger:parameters getSilence
type GetSilenceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of the silence to get
	  Required: true
	  In: path
	*/
	SilenceID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSilenceParams() beforehand.
func (o *GetSilenceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSilenceID, rhkSilenceID, _ := route.Params.GetOK("silenceID")
	if err := o.bindSilenceID(rSilenceID, rhkSilenceID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSilenceID binds and validates parameter SilenceID from path.
func (o *GetSilenceParams) bindSilenceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("silenceID", "path", "strfmt.UUID", raw)
	}
	o.SilenceID = *(value.(*strfmt.UUID))

	if err := o.validateSilenceID(formats); err != nil {
		return err
	}

	return nil
}

// validateSilenceID carries on validations for parameter SilenceID
func (o *GetSilenceParams) validateSilenceID(formats strfmt.Registry) error {

	if err := validate.FormatOf("silenceID", "path", "uuid", o.SilenceID.String(), formats); err != nil {
		return err
	}
	return nil
}
