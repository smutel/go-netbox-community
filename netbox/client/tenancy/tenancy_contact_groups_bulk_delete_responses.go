// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyContactGroupsBulkDeleteReader is a Reader for the TenancyContactGroupsBulkDelete structure.
type TenancyContactGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyContactGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactGroupsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactGroupsBulkDeleteNoContent creates a TenancyContactGroupsBulkDeleteNoContent with default headers values
func NewTenancyContactGroupsBulkDeleteNoContent() *TenancyContactGroupsBulkDeleteNoContent {
	return &TenancyContactGroupsBulkDeleteNoContent{}
}

/*
TenancyContactGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

TenancyContactGroupsBulkDeleteNoContent tenancy contact groups bulk delete no content
*/
type TenancyContactGroupsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this tenancy contact groups bulk delete no content response has a 2xx status code
func (o *TenancyContactGroupsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact groups bulk delete no content response has a 3xx status code
func (o *TenancyContactGroupsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact groups bulk delete no content response has a 4xx status code
func (o *TenancyContactGroupsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact groups bulk delete no content response has a 5xx status code
func (o *TenancyContactGroupsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact groups bulk delete no content response a status code equal to that given
func (o *TenancyContactGroupsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *TenancyContactGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-groups/][%d] tenancyContactGroupsBulkDeleteNoContent ", 204)
}

func (o *TenancyContactGroupsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-groups/][%d] tenancyContactGroupsBulkDeleteNoContent ", 204)
}

func (o *TenancyContactGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTenancyContactGroupsBulkDeleteDefault creates a TenancyContactGroupsBulkDeleteDefault with default headers values
func NewTenancyContactGroupsBulkDeleteDefault(code int) *TenancyContactGroupsBulkDeleteDefault {
	return &TenancyContactGroupsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
TenancyContactGroupsBulkDeleteDefault describes a response with status code -1, with default header values.

TenancyContactGroupsBulkDeleteDefault tenancy contact groups bulk delete default
*/
type TenancyContactGroupsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contact groups bulk delete default response
func (o *TenancyContactGroupsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy contact groups bulk delete default response has a 2xx status code
func (o *TenancyContactGroupsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact groups bulk delete default response has a 3xx status code
func (o *TenancyContactGroupsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact groups bulk delete default response has a 4xx status code
func (o *TenancyContactGroupsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact groups bulk delete default response has a 5xx status code
func (o *TenancyContactGroupsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact groups bulk delete default response a status code equal to that given
func (o *TenancyContactGroupsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyContactGroupsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-groups/][%d] tenancy_contact-groups_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactGroupsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-groups/][%d] tenancy_contact-groups_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactGroupsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactGroupsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}