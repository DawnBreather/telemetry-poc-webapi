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

	"github.com/DawnBreather/go-commons/app/poc/telemetry/client_and_tests/models"
)

// NewPutUserUpdateUserIDParams creates a new PutUserUpdateUserIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutUserUpdateUserIDParams() *PutUserUpdateUserIDParams {
	return &PutUserUpdateUserIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutUserUpdateUserIDParamsWithTimeout creates a new PutUserUpdateUserIDParams object
// with the ability to set a timeout on a request.
func NewPutUserUpdateUserIDParamsWithTimeout(timeout time.Duration) *PutUserUpdateUserIDParams {
	return &PutUserUpdateUserIDParams{
		timeout: timeout,
	}
}

// NewPutUserUpdateUserIDParamsWithContext creates a new PutUserUpdateUserIDParams object
// with the ability to set a context for a request.
func NewPutUserUpdateUserIDParamsWithContext(ctx context.Context) *PutUserUpdateUserIDParams {
	return &PutUserUpdateUserIDParams{
		Context: ctx,
	}
}

// NewPutUserUpdateUserIDParamsWithHTTPClient creates a new PutUserUpdateUserIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutUserUpdateUserIDParamsWithHTTPClient(client *http.Client) *PutUserUpdateUserIDParams {
	return &PutUserUpdateUserIDParams{
		HTTPClient: client,
	}
}

/*
PutUserUpdateUserIDParams contains all the parameters to send to the API endpoint

	for the put user update user ID operation.

	Typically these are written to a http.Request.
*/
type PutUserUpdateUserIDParams struct {

	/* Body.

	   User object with updated information
	*/
	Body *models.User

	/* UserID.

	   Unique identifier of the user
	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put user update user ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutUserUpdateUserIDParams) WithDefaults() *PutUserUpdateUserIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put user update user ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutUserUpdateUserIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put user update user ID params
func (o *PutUserUpdateUserIDParams) WithTimeout(timeout time.Duration) *PutUserUpdateUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put user update user ID params
func (o *PutUserUpdateUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put user update user ID params
func (o *PutUserUpdateUserIDParams) WithContext(ctx context.Context) *PutUserUpdateUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put user update user ID params
func (o *PutUserUpdateUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put user update user ID params
func (o *PutUserUpdateUserIDParams) WithHTTPClient(client *http.Client) *PutUserUpdateUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put user update user ID params
func (o *PutUserUpdateUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put user update user ID params
func (o *PutUserUpdateUserIDParams) WithBody(body *models.User) *PutUserUpdateUserIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put user update user ID params
func (o *PutUserUpdateUserIDParams) SetBody(body *models.User) {
	o.Body = body
}

// WithUserID adds the userID to the put user update user ID params
func (o *PutUserUpdateUserIDParams) WithUserID(userID int64) *PutUserUpdateUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the put user update user ID params
func (o *PutUserUpdateUserIDParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PutUserUpdateUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param userID
	if err := r.SetPathParam("userID", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
