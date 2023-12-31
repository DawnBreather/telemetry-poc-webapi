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

// NewGetUserUserIDParams creates a new GetUserUserIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserUserIDParams() *GetUserUserIDParams {
	return &GetUserUserIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserUserIDParamsWithTimeout creates a new GetUserUserIDParams object
// with the ability to set a timeout on a request.
func NewGetUserUserIDParamsWithTimeout(timeout time.Duration) *GetUserUserIDParams {
	return &GetUserUserIDParams{
		timeout: timeout,
	}
}

// NewGetUserUserIDParamsWithContext creates a new GetUserUserIDParams object
// with the ability to set a context for a request.
func NewGetUserUserIDParamsWithContext(ctx context.Context) *GetUserUserIDParams {
	return &GetUserUserIDParams{
		Context: ctx,
	}
}

// NewGetUserUserIDParamsWithHTTPClient creates a new GetUserUserIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserUserIDParamsWithHTTPClient(client *http.Client) *GetUserUserIDParams {
	return &GetUserUserIDParams{
		HTTPClient: client,
	}
}

/*
GetUserUserIDParams contains all the parameters to send to the API endpoint

	for the get user user ID operation.

	Typically these are written to a http.Request.
*/
type GetUserUserIDParams struct {

	/* UserID.

	   Unique identifier of the user
	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user user ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserUserIDParams) WithDefaults() *GetUserUserIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user user ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserUserIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user user ID params
func (o *GetUserUserIDParams) WithTimeout(timeout time.Duration) *GetUserUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user user ID params
func (o *GetUserUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user user ID params
func (o *GetUserUserIDParams) WithContext(ctx context.Context) *GetUserUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user user ID params
func (o *GetUserUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user user ID params
func (o *GetUserUserIDParams) WithHTTPClient(client *http.Client) *GetUserUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user user ID params
func (o *GetUserUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get user user ID params
func (o *GetUserUserIDParams) WithUserID(userID int64) *GetUserUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user user ID params
func (o *GetUserUserIDParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userID
	if err := r.SetPathParam("userID", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
