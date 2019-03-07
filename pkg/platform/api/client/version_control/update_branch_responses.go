// Code generated by go-swagger; DO NOT EDIT.

package version_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/pkg/platform/api/models"
)

// UpdateBranchReader is a Reader for the UpdateBranch structure.
type UpdateBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateBranchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateBranchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateBranchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateBranchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateBranchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateBranchOK creates a UpdateBranchOK with default headers values
func NewUpdateBranchOK() *UpdateBranchOK {
	return &UpdateBranchOK{}
}

/*UpdateBranchOK handles this case with default header values.

Branch was updated, returns resulting branch
*/
type UpdateBranchOK struct {
	Payload *models.Branch
}

func (o *UpdateBranchOK) Error() string {
	return fmt.Sprintf("[PUT /vcs/branch/{branchID}][%d] updateBranchOK  %+v", 200, o.Payload)
}

func (o *UpdateBranchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Branch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBranchBadRequest creates a UpdateBranchBadRequest with default headers values
func NewUpdateBranchBadRequest() *UpdateBranchBadRequest {
	return &UpdateBranchBadRequest{}
}

/*UpdateBranchBadRequest handles this case with default header values.

Bad Request
*/
type UpdateBranchBadRequest struct {
	Payload *models.Message
}

func (o *UpdateBranchBadRequest) Error() string {
	return fmt.Sprintf("[PUT /vcs/branch/{branchID}][%d] updateBranchBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateBranchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBranchForbidden creates a UpdateBranchForbidden with default headers values
func NewUpdateBranchForbidden() *UpdateBranchForbidden {
	return &UpdateBranchForbidden{}
}

/*UpdateBranchForbidden handles this case with default header values.

Forbidden
*/
type UpdateBranchForbidden struct {
	Payload *models.Message
}

func (o *UpdateBranchForbidden) Error() string {
	return fmt.Sprintf("[PUT /vcs/branch/{branchID}][%d] updateBranchForbidden  %+v", 403, o.Payload)
}

func (o *UpdateBranchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBranchNotFound creates a UpdateBranchNotFound with default headers values
func NewUpdateBranchNotFound() *UpdateBranchNotFound {
	return &UpdateBranchNotFound{}
}

/*UpdateBranchNotFound handles this case with default header values.

branch was not found
*/
type UpdateBranchNotFound struct {
	Payload *models.Message
}

func (o *UpdateBranchNotFound) Error() string {
	return fmt.Sprintf("[PUT /vcs/branch/{branchID}][%d] updateBranchNotFound  %+v", 404, o.Payload)
}

func (o *UpdateBranchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBranchInternalServerError creates a UpdateBranchInternalServerError with default headers values
func NewUpdateBranchInternalServerError() *UpdateBranchInternalServerError {
	return &UpdateBranchInternalServerError{}
}

/*UpdateBranchInternalServerError handles this case with default header values.

Server Error
*/
type UpdateBranchInternalServerError struct {
	Payload *models.Message
}

func (o *UpdateBranchInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /vcs/branch/{branchID}][%d] updateBranchInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateBranchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}