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

package alertgroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nholuongut/alertmanager/api/v2/models"
)

// GetAlertGroupsReader is a Reader for the GetAlertGroups structure.
type GetAlertGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /alerts/groups] getAlertGroups", response, response.Code())
	}
}

// NewGetAlertGroupsOK creates a GetAlertGroupsOK with default headers values
func NewGetAlertGroupsOK() *GetAlertGroupsOK {
	return &GetAlertGroupsOK{}
}

/*
GetAlertGroupsOK describes a response with status code 200, with default header values.

Get alert groups response
*/
type GetAlertGroupsOK struct {
	Payload models.AlertGroups
}

// IsSuccess returns true when this get alert groups o k response has a 2xx status code
func (o *GetAlertGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert groups o k response has a 3xx status code
func (o *GetAlertGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert groups o k response has a 4xx status code
func (o *GetAlertGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert groups o k response has a 5xx status code
func (o *GetAlertGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert groups o k response a status code equal to that given
func (o *GetAlertGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alert groups o k response
func (o *GetAlertGroupsOK) Code() int {
	return 200
}

func (o *GetAlertGroupsOK) Error() string {
	return fmt.Sprintf("[GET /alerts/groups][%d] getAlertGroupsOK  %+v", 200, o.Payload)
}

func (o *GetAlertGroupsOK) String() string {
	return fmt.Sprintf("[GET /alerts/groups][%d] getAlertGroupsOK  %+v", 200, o.Payload)
}

func (o *GetAlertGroupsOK) GetPayload() models.AlertGroups {
	return o.Payload
}

func (o *GetAlertGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertGroupsBadRequest creates a GetAlertGroupsBadRequest with default headers values
func NewGetAlertGroupsBadRequest() *GetAlertGroupsBadRequest {
	return &GetAlertGroupsBadRequest{}
}

/*
GetAlertGroupsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAlertGroupsBadRequest struct {
	Payload string
}

// IsSuccess returns true when this get alert groups bad request response has a 2xx status code
func (o *GetAlertGroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert groups bad request response has a 3xx status code
func (o *GetAlertGroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert groups bad request response has a 4xx status code
func (o *GetAlertGroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert groups bad request response has a 5xx status code
func (o *GetAlertGroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert groups bad request response a status code equal to that given
func (o *GetAlertGroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get alert groups bad request response
func (o *GetAlertGroupsBadRequest) Code() int {
	return 400
}

func (o *GetAlertGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /alerts/groups][%d] getAlertGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertGroupsBadRequest) String() string {
	return fmt.Sprintf("[GET /alerts/groups][%d] getAlertGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertGroupsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetAlertGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertGroupsInternalServerError creates a GetAlertGroupsInternalServerError with default headers values
func NewGetAlertGroupsInternalServerError() *GetAlertGroupsInternalServerError {
	return &GetAlertGroupsInternalServerError{}
}

/*
GetAlertGroupsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAlertGroupsInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get alert groups internal server error response has a 2xx status code
func (o *GetAlertGroupsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert groups internal server error response has a 3xx status code
func (o *GetAlertGroupsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert groups internal server error response has a 4xx status code
func (o *GetAlertGroupsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert groups internal server error response has a 5xx status code
func (o *GetAlertGroupsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get alert groups internal server error response a status code equal to that given
func (o *GetAlertGroupsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get alert groups internal server error response
func (o *GetAlertGroupsInternalServerError) Code() int {
	return 500
}

func (o *GetAlertGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alerts/groups][%d] getAlertGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertGroupsInternalServerError) String() string {
	return fmt.Sprintf("[GET /alerts/groups][%d] getAlertGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertGroupsInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetAlertGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
