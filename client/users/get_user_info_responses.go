//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// GetUserInfoReader is a Reader for the GetUserInfo structure.
type GetUserInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserInfoOK creates a GetUserInfoOK with default headers values
func NewGetUserInfoOK() *GetUserInfoOK {
	return &GetUserInfoOK{}
}

/*
GetUserInfoOK describes a response with status code 200, with default header values.

Info about the user
*/
type GetUserInfoOK struct {
	Payload *models.UserInfo
}

// IsSuccess returns true when this get user info o k response has a 2xx status code
func (o *GetUserInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user info o k response has a 3xx status code
func (o *GetUserInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user info o k response has a 4xx status code
func (o *GetUserInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user info o k response has a 5xx status code
func (o *GetUserInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user info o k response a status code equal to that given
func (o *GetUserInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user info o k response
func (o *GetUserInfoOK) Code() int {
	return 200
}

func (o *GetUserInfoOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserInfoOK  %+v", 200, o.Payload)
}

func (o *GetUserInfoOK) String() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserInfoOK  %+v", 200, o.Payload)
}

func (o *GetUserInfoOK) GetPayload() *models.UserInfo {
	return o.Payload
}

func (o *GetUserInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserInfoUnauthorized creates a GetUserInfoUnauthorized with default headers values
func NewGetUserInfoUnauthorized() *GetUserInfoUnauthorized {
	return &GetUserInfoUnauthorized{}
}

/*
GetUserInfoUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type GetUserInfoUnauthorized struct {
}

// IsSuccess returns true when this get user info unauthorized response has a 2xx status code
func (o *GetUserInfoUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user info unauthorized response has a 3xx status code
func (o *GetUserInfoUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user info unauthorized response has a 4xx status code
func (o *GetUserInfoUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user info unauthorized response has a 5xx status code
func (o *GetUserInfoUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user info unauthorized response a status code equal to that given
func (o *GetUserInfoUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get user info unauthorized response
func (o *GetUserInfoUnauthorized) Code() int {
	return 401
}

func (o *GetUserInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserInfoUnauthorized ", 401)
}

func (o *GetUserInfoUnauthorized) String() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserInfoUnauthorized ", 401)
}

func (o *GetUserInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserInfoForbidden creates a GetUserInfoForbidden with default headers values
func NewGetUserInfoForbidden() *GetUserInfoForbidden {
	return &GetUserInfoForbidden{}
}

/*
GetUserInfoForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUserInfoForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get user info forbidden response has a 2xx status code
func (o *GetUserInfoForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user info forbidden response has a 3xx status code
func (o *GetUserInfoForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user info forbidden response has a 4xx status code
func (o *GetUserInfoForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user info forbidden response has a 5xx status code
func (o *GetUserInfoForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user info forbidden response a status code equal to that given
func (o *GetUserInfoForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get user info forbidden response
func (o *GetUserInfoForbidden) Code() int {
	return 403
}

func (o *GetUserInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserInfoForbidden  %+v", 403, o.Payload)
}

func (o *GetUserInfoForbidden) String() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserInfoForbidden  %+v", 403, o.Payload)
}

func (o *GetUserInfoForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetUserInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserInfoNotFound creates a GetUserInfoNotFound with default headers values
func NewGetUserInfoNotFound() *GetUserInfoNotFound {
	return &GetUserInfoNotFound{}
}

/*
GetUserInfoNotFound describes a response with status code 404, with default header values.

user not found
*/
type GetUserInfoNotFound struct {
}

// IsSuccess returns true when this get user info not found response has a 2xx status code
func (o *GetUserInfoNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user info not found response has a 3xx status code
func (o *GetUserInfoNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user info not found response has a 4xx status code
func (o *GetUserInfoNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user info not found response has a 5xx status code
func (o *GetUserInfoNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user info not found response a status code equal to that given
func (o *GetUserInfoNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get user info not found response
func (o *GetUserInfoNotFound) Code() int {
	return 404
}

func (o *GetUserInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserInfoNotFound ", 404)
}

func (o *GetUserInfoNotFound) String() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserInfoNotFound ", 404)
}

func (o *GetUserInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserInfoInternalServerError creates a GetUserInfoInternalServerError with default headers values
func NewGetUserInfoInternalServerError() *GetUserInfoInternalServerError {
	return &GetUserInfoInternalServerError{}
}

/*
GetUserInfoInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type GetUserInfoInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get user info internal server error response has a 2xx status code
func (o *GetUserInfoInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user info internal server error response has a 3xx status code
func (o *GetUserInfoInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user info internal server error response has a 4xx status code
func (o *GetUserInfoInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user info internal server error response has a 5xx status code
func (o *GetUserInfoInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user info internal server error response a status code equal to that given
func (o *GetUserInfoInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get user info internal server error response
func (o *GetUserInfoInternalServerError) Code() int {
	return 500
}

func (o *GetUserInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserInfoInternalServerError) String() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserInfoInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetUserInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
