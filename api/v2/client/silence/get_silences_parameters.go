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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetSilencesParams creates a new GetSilencesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSilencesParams() *GetSilencesParams {
	return &GetSilencesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSilencesParamsWithTimeout creates a new GetSilencesParams object
// with the ability to set a timeout on a request.
func NewGetSilencesParamsWithTimeout(timeout time.Duration) *GetSilencesParams {
	return &GetSilencesParams{
		timeout: timeout,
	}
}

// NewGetSilencesParamsWithContext creates a new GetSilencesParams object
// with the ability to set a context for a request.
func NewGetSilencesParamsWithContext(ctx context.Context) *GetSilencesParams {
	return &GetSilencesParams{
		Context: ctx,
	}
}

// NewGetSilencesParamsWithHTTPClient creates a new GetSilencesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSilencesParamsWithHTTPClient(client *http.Client) *GetSilencesParams {
	return &GetSilencesParams{
		HTTPClient: client,
	}
}

/*
GetSilencesParams contains all the parameters to send to the API endpoint

	for the get silences operation.

	Typically these are written to a http.Request.
*/
type GetSilencesParams struct {

	/* Filter.

	   A list of matchers to filter silences by
	*/
	Filter []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get silences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSilencesParams) WithDefaults() *GetSilencesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get silences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSilencesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get silences params
func (o *GetSilencesParams) WithTimeout(timeout time.Duration) *GetSilencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get silences params
func (o *GetSilencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get silences params
func (o *GetSilencesParams) WithContext(ctx context.Context) *GetSilencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get silences params
func (o *GetSilencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get silences params
func (o *GetSilencesParams) WithHTTPClient(client *http.Client) *GetSilencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get silences params
func (o *GetSilencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get silences params
func (o *GetSilencesParams) WithFilter(filter []string) *GetSilencesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get silences params
func (o *GetSilencesParams) SetFilter(filter []string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *GetSilencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// binding items for filter
		joinedFilter := o.bindParamFilter(reg)

		// query array param filter
		if err := r.SetQueryParam("filter", joinedFilter...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetSilences binds the parameter filter
func (o *GetSilencesParams) bindParamFilter(formats strfmt.Registry) []string {
	filterIR := o.Filter

	var filterIC []string
	for _, filterIIR := range filterIR { // explode []string

		filterIIV := filterIIR // string as string
		filterIC = append(filterIC, filterIIV)
	}

	// items.CollectionFormat: "multi"
	filterIS := swag.JoinByFormat(filterIC, "multi")

	return filterIS
}
