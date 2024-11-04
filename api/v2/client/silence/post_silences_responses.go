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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostSilencesReader is a Reader for the PostSilences structure.
type PostSilencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSilencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSilencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSilencesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostSilencesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /silences] postSilences", response, response.Code())
	}
}

// NewPostSilencesOK creates a PostSilencesOK with default headers values
func NewPostSilencesOK() *PostSilencesOK {
	return &PostSilencesOK{}
}

/*
PostSilencesOK describes a response with status code 200, with default header values.

Create / update silence response
*/
type PostSilencesOK struct {
	Payload *PostSilencesOKBody
}

// IsSuccess returns true when this post silences o k response has a 2xx status code
func (o *PostSilencesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post silences o k response has a 3xx status code
func (o *PostSilencesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post silences o k response has a 4xx status code
func (o *PostSilencesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post silences o k response has a 5xx status code
func (o *PostSilencesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post silences o k response a status code equal to that given
func (o *PostSilencesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post silences o k response
func (o *PostSilencesOK) Code() int {
	return 200
}

func (o *PostSilencesOK) Error() string {
	return fmt.Sprintf("[POST /silences][%d] postSilencesOK  %+v", 200, o.Payload)
}

func (o *PostSilencesOK) String() string {
	return fmt.Sprintf("[POST /silences][%d] postSilencesOK  %+v", 200, o.Payload)
}

func (o *PostSilencesOK) GetPayload() *PostSilencesOKBody {
	return o.Payload
}

func (o *PostSilencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostSilencesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSilencesBadRequest creates a PostSilencesBadRequest with default headers values
func NewPostSilencesBadRequest() *PostSilencesBadRequest {
	return &PostSilencesBadRequest{}
}

/*
PostSilencesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostSilencesBadRequest struct {
	Payload string
}

// IsSuccess returns true when this post silences bad request response has a 2xx status code
func (o *PostSilencesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post silences bad request response has a 3xx status code
func (o *PostSilencesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post silences bad request response has a 4xx status code
func (o *PostSilencesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post silences bad request response has a 5xx status code
func (o *PostSilencesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post silences bad request response a status code equal to that given
func (o *PostSilencesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post silences bad request response
func (o *PostSilencesBadRequest) Code() int {
	return 400
}

func (o *PostSilencesBadRequest) Error() string {
	return fmt.Sprintf("[POST /silences][%d] postSilencesBadRequest  %+v", 400, o.Payload)
}

func (o *PostSilencesBadRequest) String() string {
	return fmt.Sprintf("[POST /silences][%d] postSilencesBadRequest  %+v", 400, o.Payload)
}

func (o *PostSilencesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PostSilencesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSilencesNotFound creates a PostSilencesNotFound with default headers values
func NewPostSilencesNotFound() *PostSilencesNotFound {
	return &PostSilencesNotFound{}
}

/*
PostSilencesNotFound describes a response with status code 404, with default header values.

A silence with the specified ID was not found
*/
type PostSilencesNotFound struct {
	Payload string
}

// IsSuccess returns true when this post silences not found response has a 2xx status code
func (o *PostSilencesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post silences not found response has a 3xx status code
func (o *PostSilencesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post silences not found response has a 4xx status code
func (o *PostSilencesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post silences not found response has a 5xx status code
func (o *PostSilencesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post silences not found response a status code equal to that given
func (o *PostSilencesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post silences not found response
func (o *PostSilencesNotFound) Code() int {
	return 404
}

func (o *PostSilencesNotFound) Error() string {
	return fmt.Sprintf("[POST /silences][%d] postSilencesNotFound  %+v", 404, o.Payload)
}

func (o *PostSilencesNotFound) String() string {
	return fmt.Sprintf("[POST /silences][%d] postSilencesNotFound  %+v", 404, o.Payload)
}

func (o *PostSilencesNotFound) GetPayload() string {
	return o.Payload
}

func (o *PostSilencesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostSilencesOKBody post silences o k body
swagger:model PostSilencesOKBody
*/
type PostSilencesOKBody struct {

	// silence ID
	SilenceID string `json:"silenceID,omitempty"`
}

// Validate validates this post silences o k body
func (o *PostSilencesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post silences o k body based on context it is used
func (o *PostSilencesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostSilencesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostSilencesOKBody) UnmarshalBinary(b []byte) error {
	var res PostSilencesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
