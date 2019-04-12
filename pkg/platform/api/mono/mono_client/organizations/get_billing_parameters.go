// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBillingParams creates a new GetBillingParams object
// with the default values initialized.
func NewGetBillingParams() *GetBillingParams {
	var ()
	return &GetBillingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBillingParamsWithTimeout creates a new GetBillingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBillingParamsWithTimeout(timeout time.Duration) *GetBillingParams {
	var ()
	return &GetBillingParams{

		timeout: timeout,
	}
}

// NewGetBillingParamsWithContext creates a new GetBillingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBillingParamsWithContext(ctx context.Context) *GetBillingParams {
	var ()
	return &GetBillingParams{

		Context: ctx,
	}
}

// NewGetBillingParamsWithHTTPClient creates a new GetBillingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBillingParamsWithHTTPClient(client *http.Client) *GetBillingParams {
	var ()
	return &GetBillingParams{
		HTTPClient: client,
	}
}

/*GetBillingParams contains all the parameters to send to the API endpoint
for the get billing operation typically these are written to a http.Request
*/
type GetBillingParams struct {

	/*OrganizationName
	  desired organization

	*/
	OrganizationName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get billing params
func (o *GetBillingParams) WithTimeout(timeout time.Duration) *GetBillingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get billing params
func (o *GetBillingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get billing params
func (o *GetBillingParams) WithContext(ctx context.Context) *GetBillingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get billing params
func (o *GetBillingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get billing params
func (o *GetBillingParams) WithHTTPClient(client *http.Client) *GetBillingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get billing params
func (o *GetBillingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationName adds the organizationName to the get billing params
func (o *GetBillingParams) WithOrganizationName(organizationName string) *GetBillingParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the get billing params
func (o *GetBillingParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WriteToRequest writes these params to a swagger request
func (o *GetBillingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}