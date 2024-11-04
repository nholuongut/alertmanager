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

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nholuongut/alertmanager/api/v2/models"
)

// GetStatusOKCode is the HTTP code returned for type GetStatusOK
const GetStatusOKCode int = 200

/*
GetStatusOK Get status response

swagger:response getStatusOK
*/
type GetStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.AlertmanagerStatus `json:"body,omitempty"`
}

// NewGetStatusOK creates GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {

	return &GetStatusOK{}
}

// WithPayload adds the payload to the get status o k response
func (o *GetStatusOK) WithPayload(payload *models.AlertmanagerStatus) *GetStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get status o k response
func (o *GetStatusOK) SetPayload(payload *models.AlertmanagerStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
