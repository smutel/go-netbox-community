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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimModuleTypesBulkDeleteReader is a Reader for the DcimModuleTypesBulkDelete structure.
type DcimModuleTypesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleTypesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimModuleTypesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleTypesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleTypesBulkDeleteNoContent creates a DcimModuleTypesBulkDeleteNoContent with default headers values
func NewDcimModuleTypesBulkDeleteNoContent() *DcimModuleTypesBulkDeleteNoContent {
	return &DcimModuleTypesBulkDeleteNoContent{}
}

/*
DcimModuleTypesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimModuleTypesBulkDeleteNoContent dcim module types bulk delete no content
*/
type DcimModuleTypesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim module types bulk delete no content response has a 2xx status code
func (o *DcimModuleTypesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module types bulk delete no content response has a 3xx status code
func (o *DcimModuleTypesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module types bulk delete no content response has a 4xx status code
func (o *DcimModuleTypesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module types bulk delete no content response has a 5xx status code
func (o *DcimModuleTypesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module types bulk delete no content response a status code equal to that given
func (o *DcimModuleTypesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimModuleTypesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-types/][%d] dcimModuleTypesBulkDeleteNoContent ", 204)
}

func (o *DcimModuleTypesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-types/][%d] dcimModuleTypesBulkDeleteNoContent ", 204)
}

func (o *DcimModuleTypesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimModuleTypesBulkDeleteDefault creates a DcimModuleTypesBulkDeleteDefault with default headers values
func NewDcimModuleTypesBulkDeleteDefault(code int) *DcimModuleTypesBulkDeleteDefault {
	return &DcimModuleTypesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimModuleTypesBulkDeleteDefault describes a response with status code -1, with default header values.

DcimModuleTypesBulkDeleteDefault dcim module types bulk delete default
*/
type DcimModuleTypesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module types bulk delete default response
func (o *DcimModuleTypesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module types bulk delete default response has a 2xx status code
func (o *DcimModuleTypesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module types bulk delete default response has a 3xx status code
func (o *DcimModuleTypesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module types bulk delete default response has a 4xx status code
func (o *DcimModuleTypesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module types bulk delete default response has a 5xx status code
func (o *DcimModuleTypesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module types bulk delete default response a status code equal to that given
func (o *DcimModuleTypesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleTypesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-types/][%d] dcim_module-types_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleTypesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-types/][%d] dcim_module-types_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleTypesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleTypesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}