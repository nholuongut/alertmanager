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
)

// NewDeleteSilenceParams creates a new DeleteSilenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSilenceParams() *DeleteSilenceParams {
	return &DeleteSilenceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSilenceParamsWithTimeout creates a new DeleteSilenceParams object
// with the ability to set a timeout on a request.
func NewDeleteSilenceParamsWithTimeout(timeout time.Duration) *DeleteSilenceParams {
	return &DeleteSilenceParams{
		timeout: timeout,
	}
}

// NewDeleteSilenceParamsWithContext creates a new DeleteSilenceParams object
// with the ability to set a context for a request.
func NewDeleteSilenceParamsWithContext(ctx context.Context) *DeleteSilenceParams {
	return &DeleteSilenceParams{
		Context: ctx,
	}
}

// NewDeleteSilenceParamsWithHTTPClient creates a new DeleteSilenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSilenceParamsWithHTTPClient(client *http.Client) *DeleteSilenceParams {
	return &DeleteSilenceParams{
		HTTPClient: client,
	}
}

/*
DeleteSilenceParams contains all the parameters to send to the API endpoint

	for the delete silence operation.

	Typically these are written to a http.Request.
*/
type DeleteSilenceParams struct {

	/* SilenceID.

	   ID of the silence to get

	   Format: uuid
	*/
	SilenceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete silence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSilenceParams) WithDefaults() *DeleteSilenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete silence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSilenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete silence params
func (o *DeleteSilenceParams) WithTimeout(timeout time.Duration) *DeleteSilenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete silence params
func (o *DeleteSilenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete silence params
func (o *DeleteSilenceParams) WithContext(ctx context.Context) *DeleteSilenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete silence params
func (o *DeleteSilenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete silence params
func (o *DeleteSilenceParams) WithHTTPClient(client *http.Client) *DeleteSilenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete silence params
func (o *DeleteSilenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSilenceID adds the silenceID to the delete silence params
func (o *DeleteSilenceParams) WithSilenceID(silenceID strfmt.UUID) *DeleteSilenceParams {
	o.SetSilenceID(silenceID)
	return o
}

// SetSilenceID adds the silenceId to the delete silence params
func (o *DeleteSilenceParams) SetSilenceID(silenceID strfmt.UUID) {
	o.SilenceID = silenceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSilenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param silenceID
	if err := r.SetPathParam("silenceID", o.SilenceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
