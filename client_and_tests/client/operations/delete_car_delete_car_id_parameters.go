// Code generated by go-swagger; DO NOT EDIT.

package operations

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
	"github.com/go-openapi/swag"
)

// NewDeleteCarDeleteCarIDParams creates a new DeleteCarDeleteCarIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCarDeleteCarIDParams() *DeleteCarDeleteCarIDParams {
	return &DeleteCarDeleteCarIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCarDeleteCarIDParamsWithTimeout creates a new DeleteCarDeleteCarIDParams object
// with the ability to set a timeout on a request.
func NewDeleteCarDeleteCarIDParamsWithTimeout(timeout time.Duration) *DeleteCarDeleteCarIDParams {
	return &DeleteCarDeleteCarIDParams{
		timeout: timeout,
	}
}

// NewDeleteCarDeleteCarIDParamsWithContext creates a new DeleteCarDeleteCarIDParams object
// with the ability to set a context for a request.
func NewDeleteCarDeleteCarIDParamsWithContext(ctx context.Context) *DeleteCarDeleteCarIDParams {
	return &DeleteCarDeleteCarIDParams{
		Context: ctx,
	}
}

// NewDeleteCarDeleteCarIDParamsWithHTTPClient creates a new DeleteCarDeleteCarIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCarDeleteCarIDParamsWithHTTPClient(client *http.Client) *DeleteCarDeleteCarIDParams {
	return &DeleteCarDeleteCarIDParams{
		HTTPClient: client,
	}
}

/*
DeleteCarDeleteCarIDParams contains all the parameters to send to the API endpoint

	for the delete car delete car ID operation.

	Typically these are written to a http.Request.
*/
type DeleteCarDeleteCarIDParams struct {

	/* CarID.

	   Unique identifier of the car to be deleted
	*/
	CarID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete car delete car ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCarDeleteCarIDParams) WithDefaults() *DeleteCarDeleteCarIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete car delete car ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCarDeleteCarIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete car delete car ID params
func (o *DeleteCarDeleteCarIDParams) WithTimeout(timeout time.Duration) *DeleteCarDeleteCarIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete car delete car ID params
func (o *DeleteCarDeleteCarIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete car delete car ID params
func (o *DeleteCarDeleteCarIDParams) WithContext(ctx context.Context) *DeleteCarDeleteCarIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete car delete car ID params
func (o *DeleteCarDeleteCarIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete car delete car ID params
func (o *DeleteCarDeleteCarIDParams) WithHTTPClient(client *http.Client) *DeleteCarDeleteCarIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete car delete car ID params
func (o *DeleteCarDeleteCarIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCarID adds the carID to the delete car delete car ID params
func (o *DeleteCarDeleteCarIDParams) WithCarID(carID int64) *DeleteCarDeleteCarIDParams {
	o.SetCarID(carID)
	return o
}

// SetCarID adds the carId to the delete car delete car ID params
func (o *DeleteCarDeleteCarIDParams) SetCarID(carID int64) {
	o.CarID = carID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCarDeleteCarIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param carID
	if err := r.SetPathParam("carID", swag.FormatInt64(o.CarID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
