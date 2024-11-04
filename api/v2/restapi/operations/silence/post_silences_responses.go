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

	"github.com/go-openapi/runtime"
)

// PostSilencesOKCode is the HTTP code returned for type PostSilencesOK
const PostSilencesOKCode int = 200

/*
PostSilencesOK Create / update silence response

swagger:response postSilencesOK
*/
type PostSilencesOK struct {

	/*
	  In: Body
	*/
	Payload *PostSilencesOKBody `json:"body,omitempty"`
}

// NewPostSilencesOK creates PostSilencesOK with default headers values
func NewPostSilencesOK() *PostSilencesOK {

	return &PostSilencesOK{}
}

// WithPayload adds the payload to the post silences o k response
func (o *PostSilencesOK) WithPayload(payload *PostSilencesOKBody) *PostSilencesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post silences o k response
func (o *PostSilencesOK) SetPayload(payload *PostSilencesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSilencesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostSilencesBadRequestCode is the HTTP code returned for type PostSilencesBadRequest
const PostSilencesBadRequestCode int = 400

/*
PostSilencesBadRequest Bad request

swagger:response postSilencesBadRequest
*/
type PostSilencesBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostSilencesBadRequest creates PostSilencesBadRequest with default headers values
func NewPostSilencesBadRequest() *PostSilencesBadRequest {

	return &PostSilencesBadRequest{}
}

// WithPayload adds the payload to the post silences bad request response
func (o *PostSilencesBadRequest) WithPayload(payload string) *PostSilencesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post silences bad request response
func (o *PostSilencesBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSilencesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostSilencesNotFoundCode is the HTTP code returned for type PostSilencesNotFound
const PostSilencesNotFoundCode int = 404

/*
PostSilencesNotFound A silence with the specified ID was not found

swagger:response postSilencesNotFound
*/
type PostSilencesNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostSilencesNotFound creates PostSilencesNotFound with default headers values
func NewPostSilencesNotFound() *PostSilencesNotFound {

	return &PostSilencesNotFound{}
}

// WithPayload adds the payload to the post silences not found response
func (o *PostSilencesNotFound) WithPayload(payload string) *PostSilencesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post silences not found response
func (o *PostSilencesNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSilencesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
