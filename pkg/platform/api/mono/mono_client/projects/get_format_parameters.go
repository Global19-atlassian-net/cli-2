// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewGetFormatParams creates a new GetFormatParams object
// with the default values initialized.
func NewGetFormatParams() *GetFormatParams {
	var ()
	return &GetFormatParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFormatParamsWithTimeout creates a new GetFormatParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFormatParamsWithTimeout(timeout time.Duration) *GetFormatParams {
	var ()
	return &GetFormatParams{

		timeout: timeout,
	}
}

// NewGetFormatParamsWithContext creates a new GetFormatParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFormatParamsWithContext(ctx context.Context) *GetFormatParams {
	var ()
	return &GetFormatParams{

		Context: ctx,
	}
}

// NewGetFormatParamsWithHTTPClient creates a new GetFormatParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFormatParamsWithHTTPClient(client *http.Client) *GetFormatParams {
	var ()
	return &GetFormatParams{
		HTTPClient: client,
	}
}

/*GetFormatParams contains all the parameters to send to the API endpoint
for the get format operation typically these are written to a http.Request
*/
type GetFormatParams struct {

	/*DistroID
	  desired distro

	*/
	DistroID strfmt.UUID
	/*FormatID
	  desired format

	*/
	FormatID strfmt.UUID
	/*OrganizationName
	  desired organization

	*/
	OrganizationName string
	/*ProjectName
	  desired project

	*/
	ProjectName string
	/*ReleaseID
	  desired release

	*/
	ReleaseID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get format params
func (o *GetFormatParams) WithTimeout(timeout time.Duration) *GetFormatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get format params
func (o *GetFormatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get format params
func (o *GetFormatParams) WithContext(ctx context.Context) *GetFormatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get format params
func (o *GetFormatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get format params
func (o *GetFormatParams) WithHTTPClient(client *http.Client) *GetFormatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get format params
func (o *GetFormatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDistroID adds the distroID to the get format params
func (o *GetFormatParams) WithDistroID(distroID strfmt.UUID) *GetFormatParams {
	o.SetDistroID(distroID)
	return o
}

// SetDistroID adds the distroId to the get format params
func (o *GetFormatParams) SetDistroID(distroID strfmt.UUID) {
	o.DistroID = distroID
}

// WithFormatID adds the formatID to the get format params
func (o *GetFormatParams) WithFormatID(formatID strfmt.UUID) *GetFormatParams {
	o.SetFormatID(formatID)
	return o
}

// SetFormatID adds the formatId to the get format params
func (o *GetFormatParams) SetFormatID(formatID strfmt.UUID) {
	o.FormatID = formatID
}

// WithOrganizationName adds the organizationName to the get format params
func (o *GetFormatParams) WithOrganizationName(organizationName string) *GetFormatParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the get format params
func (o *GetFormatParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WithProjectName adds the projectName to the get format params
func (o *GetFormatParams) WithProjectName(projectName string) *GetFormatParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the get format params
func (o *GetFormatParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithReleaseID adds the releaseID to the get format params
func (o *GetFormatParams) WithReleaseID(releaseID strfmt.UUID) *GetFormatParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the get format params
func (o *GetFormatParams) SetReleaseID(releaseID strfmt.UUID) {
	o.ReleaseID = releaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFormatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param distroID
	if err := r.SetPathParam("distroID", o.DistroID.String()); err != nil {
		return err
	}

	// path param formatID
	if err := r.SetPathParam("formatID", o.FormatID.String()); err != nil {
		return err
	}

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	// path param projectName
	if err := r.SetPathParam("projectName", o.ProjectName); err != nil {
		return err
	}

	// path param releaseID
	if err := r.SetPathParam("releaseID", o.ReleaseID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}