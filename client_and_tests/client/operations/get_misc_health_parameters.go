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
)

// NewGetMiscHealthParams creates a new GetMiscHealthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMiscHealthParams() *GetMiscHealthParams {
	return &GetMiscHealthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMiscHealthParamsWithTimeout creates a new GetMiscHealthParams object
// with the ability to set a timeout on a request.
func NewGetMiscHealthParamsWithTimeout(timeout time.Duration) *GetMiscHealthParams {
	return &GetMiscHealthParams{
		timeout: timeout,
	}
}

// NewGetMiscHealthParamsWithContext creates a new GetMiscHealthParams object
// with the ability to set a context for a request.
func NewGetMiscHealthParamsWithContext(ctx context.Context) *GetMiscHealthParams {
	return &GetMiscHealthParams{
		Context: ctx,
	}
}

// NewGetMiscHealthParamsWithHTTPClient creates a new GetMiscHealthParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMiscHealthParamsWithHTTPClient(client *http.Client) *GetMiscHealthParams {
	return &GetMiscHealthParams{
		HTTPClient: client,
	}
}

/*
GetMiscHealthParams contains all the parameters to send to the API endpoint

	for the get misc health operation.

	Typically these are written to a http.Request.
*/
type GetMiscHealthParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get misc health params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMiscHealthParams) WithDefaults() *GetMiscHealthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get misc health params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMiscHealthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get misc health params
func (o *GetMiscHealthParams) WithTimeout(timeout time.Duration) *GetMiscHealthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get misc health params
func (o *GetMiscHealthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get misc health params
func (o *GetMiscHealthParams) WithContext(ctx context.Context) *GetMiscHealthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get misc health params
func (o *GetMiscHealthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get misc health params
func (o *GetMiscHealthParams) WithHTTPClient(client *http.Client) *GetMiscHealthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get misc health params
func (o *GetMiscHealthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMiscHealthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
