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

// GetCarCarIDReader is a Reader for the GetCarCarID structure.
type GetCarCarIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCarCarIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCarCarIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCarCarIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCarCarIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCarCarIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCarCarIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /car/{carID}] GetCarCarID", response, response.Code())
	}
}

// NewGetCarCarIDOK creates a GetCarCarIDOK with default headers values
func NewGetCarCarIDOK() *GetCarCarIDOK {
	return &GetCarCarIDOK{}
}

/*
GetCarCarIDOK describes a response with status code 200, with default header values.

Car details
*/
type GetCarCarIDOK struct {
	Payload *models.Car
}

// IsSuccess returns true when this get car car Id o k response has a 2xx status code
func (o *GetCarCarIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get car car Id o k response has a 3xx status code
func (o *GetCarCarIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get car car Id o k response has a 4xx status code
func (o *GetCarCarIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get car car Id o k response has a 5xx status code
func (o *GetCarCarIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get car car Id o k response a status code equal to that given
func (o *GetCarCarIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get car car Id o k response
func (o *GetCarCarIDOK) Code() int {
	return 200
}

func (o *GetCarCarIDOK) Error() string {
	return fmt.Sprintf("[GET /car/{carID}][%d] getCarCarIdOK  %+v", 200, o.Payload)
}

func (o *GetCarCarIDOK) String() string {
	return fmt.Sprintf("[GET /car/{carID}][%d] getCarCarIdOK  %+v", 200, o.Payload)
}

func (o *GetCarCarIDOK) GetPayload() *models.Car {
	return o.Payload
}

func (o *GetCarCarIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Car)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCarCarIDBadRequest creates a GetCarCarIDBadRequest with default headers values
func NewGetCarCarIDBadRequest() *GetCarCarIDBadRequest {
	return &GetCarCarIDBadRequest{}
}

/*
GetCarCarIDBadRequest describes a response with status code 400, with default header values.

Bad request or invalid car ID
*/
type GetCarCarIDBadRequest struct {
}

// IsSuccess returns true when this get car car Id bad request response has a 2xx status code
func (o *GetCarCarIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get car car Id bad request response has a 3xx status code
func (o *GetCarCarIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get car car Id bad request response has a 4xx status code
func (o *GetCarCarIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get car car Id bad request response has a 5xx status code
func (o *GetCarCarIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get car car Id bad request response a status code equal to that given
func (o *GetCarCarIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get car car Id bad request response
func (o *GetCarCarIDBadRequest) Code() int {
	return 400
}

func (o *GetCarCarIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /car/{carID}][%d] getCarCarIdBadRequest ", 400)
}

func (o *GetCarCarIDBadRequest) String() string {
	return fmt.Sprintf("[GET /car/{carID}][%d] getCarCarIdBadRequest ", 400)
}

func (o *GetCarCarIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCarCarIDUnauthorized creates a GetCarCarIDUnauthorized with default headers values
func NewGetCarCarIDUnauthorized() *GetCarCarIDUnauthorized {
	return &GetCarCarIDUnauthorized{}
}

/*
GetCarCarIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized attempt to access car details
*/
type GetCarCarIDUnauthorized struct {
}

// IsSuccess returns true when this get car car Id unauthorized response has a 2xx status code
func (o *GetCarCarIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get car car Id unauthorized response has a 3xx status code
func (o *GetCarCarIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get car car Id unauthorized response has a 4xx status code
func (o *GetCarCarIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get car car Id unauthorized response has a 5xx status code
func (o *GetCarCarIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get car car Id unauthorized response a status code equal to that given
func (o *GetCarCarIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get car car Id unauthorized response
func (o *GetCarCarIDUnauthorized) Code() int {
	return 401
}

func (o *GetCarCarIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /car/{carID}][%d] getCarCarIdUnauthorized ", 401)
}

func (o *GetCarCarIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /car/{carID}][%d] getCarCarIdUnauthorized ", 401)
}

func (o *GetCarCarIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCarCarIDNotFound creates a GetCarCarIDNotFound with default headers values
func NewGetCarCarIDNotFound() *GetCarCarIDNotFound {
	return &GetCarCarIDNotFound{}
}

/*
GetCarCarIDNotFound describes a response with status code 404, with default header values.

Car not found
*/
type GetCarCarIDNotFound struct {
}

// IsSuccess returns true when this get car car Id not found response has a 2xx status code
func (o *GetCarCarIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get car car Id not found response has a 3xx status code
func (o *GetCarCarIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get car car Id not found response has a 4xx status code
func (o *GetCarCarIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get car car Id not found response has a 5xx status code
func (o *GetCarCarIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get car car Id not found response a status code equal to that given
func (o *GetCarCarIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get car car Id not found response
func (o *GetCarCarIDNotFound) Code() int {
	return 404
}

func (o *GetCarCarIDNotFound) Error() string {
	return fmt.Sprintf("[GET /car/{carID}][%d] getCarCarIdNotFound ", 404)
}

func (o *GetCarCarIDNotFound) String() string {
	return fmt.Sprintf("[GET /car/{carID}][%d] getCarCarIdNotFound ", 404)
}

func (o *GetCarCarIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCarCarIDInternalServerError creates a GetCarCarIDInternalServerError with default headers values
func NewGetCarCarIDInternalServerError() *GetCarCarIDInternalServerError {
	return &GetCarCarIDInternalServerError{}
}

/*
GetCarCarIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error or error retrieving car from database
*/
type GetCarCarIDInternalServerError struct {
}

// IsSuccess returns true when this get car car Id internal server error response has a 2xx status code
func (o *GetCarCarIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get car car Id internal server error response has a 3xx status code
func (o *GetCarCarIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get car car Id internal server error response has a 4xx status code
func (o *GetCarCarIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get car car Id internal server error response has a 5xx status code
func (o *GetCarCarIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get car car Id internal server error response a status code equal to that given
func (o *GetCarCarIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get car car Id internal server error response
func (o *GetCarCarIDInternalServerError) Code() int {
	return 500
}

func (o *GetCarCarIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /car/{carID}][%d] getCarCarIdInternalServerError ", 500)
}

func (o *GetCarCarIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /car/{carID}][%d] getCarCarIdInternalServerError ", 500)
}

func (o *GetCarCarIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
