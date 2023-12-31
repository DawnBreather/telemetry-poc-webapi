// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/DawnBreather/go-commons/app/poc/telemetry/client_and_tests/models"
)

// GetUserUserIDReader is a Reader for the GetUserUserID structure.
type GetUserUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserUserIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /user/{userID}] GetUserUserID", response, response.Code())
	}
}

// NewGetUserUserIDOK creates a GetUserUserIDOK with default headers values
func NewGetUserUserIDOK() *GetUserUserIDOK {
	return &GetUserUserIDOK{}
}

/*
GetUserUserIDOK describes a response with status code 200, with default header values.

Detailed user information
*/
type GetUserUserIDOK struct {
	Payload *models.User
}

// IsSuccess returns true when this get user user Id o k response has a 2xx status code
func (o *GetUserUserIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user user Id o k response has a 3xx status code
func (o *GetUserUserIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user user Id o k response has a 4xx status code
func (o *GetUserUserIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user user Id o k response has a 5xx status code
func (o *GetUserUserIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user user Id o k response a status code equal to that given
func (o *GetUserUserIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user user Id o k response
func (o *GetUserUserIDOK) Code() int {
	return 200
}

func (o *GetUserUserIDOK) Error() string {
	return fmt.Sprintf("[GET /user/{userID}][%d] getUserUserIdOK  %+v", 200, o.Payload)
}

func (o *GetUserUserIDOK) String() string {
	return fmt.Sprintf("[GET /user/{userID}][%d] getUserUserIdOK  %+v", 200, o.Payload)
}

func (o *GetUserUserIDOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetUserUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserUserIDBadRequest creates a GetUserUserIDBadRequest with default headers values
func NewGetUserUserIDBadRequest() *GetUserUserIDBadRequest {
	return &GetUserUserIDBadRequest{}
}

/*
GetUserUserIDBadRequest describes a response with status code 400, with default header values.

Bad request or invalid user ID
*/
type GetUserUserIDBadRequest struct {
}

// IsSuccess returns true when this get user user Id bad request response has a 2xx status code
func (o *GetUserUserIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user user Id bad request response has a 3xx status code
func (o *GetUserUserIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user user Id bad request response has a 4xx status code
func (o *GetUserUserIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user user Id bad request response has a 5xx status code
func (o *GetUserUserIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user user Id bad request response a status code equal to that given
func (o *GetUserUserIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get user user Id bad request response
func (o *GetUserUserIDBadRequest) Code() int {
	return 400
}

func (o *GetUserUserIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /user/{userID}][%d] getUserUserIdBadRequest ", 400)
}

func (o *GetUserUserIDBadRequest) String() string {
	return fmt.Sprintf("[GET /user/{userID}][%d] getUserUserIdBadRequest ", 400)
}

func (o *GetUserUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserUserIDNotFound creates a GetUserUserIDNotFound with default headers values
func NewGetUserUserIDNotFound() *GetUserUserIDNotFound {
	return &GetUserUserIDNotFound{}
}

/*
GetUserUserIDNotFound describes a response with status code 404, with default header values.

User not found
*/
type GetUserUserIDNotFound struct {
}

// IsSuccess returns true when this get user user Id not found response has a 2xx status code
func (o *GetUserUserIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user user Id not found response has a 3xx status code
func (o *GetUserUserIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user user Id not found response has a 4xx status code
func (o *GetUserUserIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user user Id not found response has a 5xx status code
func (o *GetUserUserIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user user Id not found response a status code equal to that given
func (o *GetUserUserIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get user user Id not found response
func (o *GetUserUserIDNotFound) Code() int {
	return 404
}

func (o *GetUserUserIDNotFound) Error() string {
	return fmt.Sprintf("[GET /user/{userID}][%d] getUserUserIdNotFound ", 404)
}

func (o *GetUserUserIDNotFound) String() string {
	return fmt.Sprintf("[GET /user/{userID}][%d] getUserUserIdNotFound ", 404)
}

func (o *GetUserUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserUserIDInternalServerError creates a GetUserUserIDInternalServerError with default headers values
func NewGetUserUserIDInternalServerError() *GetUserUserIDInternalServerError {
	return &GetUserUserIDInternalServerError{}
}

/*
GetUserUserIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error or error retrieving user details
*/
type GetUserUserIDInternalServerError struct {
}

// IsSuccess returns true when this get user user Id internal server error response has a 2xx status code
func (o *GetUserUserIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user user Id internal server error response has a 3xx status code
func (o *GetUserUserIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user user Id internal server error response has a 4xx status code
func (o *GetUserUserIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user user Id internal server error response has a 5xx status code
func (o *GetUserUserIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user user Id internal server error response a status code equal to that given
func (o *GetUserUserIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get user user Id internal server error response
func (o *GetUserUserIDInternalServerError) Code() int {
	return 500
}

func (o *GetUserUserIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user/{userID}][%d] getUserUserIdInternalServerError ", 500)
}

func (o *GetUserUserIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /user/{userID}][%d] getUserUserIdInternalServerError ", 500)
}

func (o *GetUserUserIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
