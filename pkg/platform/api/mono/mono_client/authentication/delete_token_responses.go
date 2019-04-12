// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// DeleteTokenReader is a Reader for the DeleteToken structure.
type DeleteTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteTokenOK creates a DeleteTokenOK with default headers values
func NewDeleteTokenOK() *DeleteTokenOK {
	return &DeleteTokenOK{}
}

/*DeleteTokenOK handles this case with default header values.

Token deleted
*/
type DeleteTokenOK struct {
	Payload *mono_models.Message
}

func (o *DeleteTokenOK) Error() string {
	return fmt.Sprintf("[DELETE /apikeys/{tokenID}][%d] deleteTokenOK  %+v", 200, o.Payload)
}

func (o *DeleteTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokenBadRequest creates a DeleteTokenBadRequest with default headers values
func NewDeleteTokenBadRequest() *DeleteTokenBadRequest {
	return &DeleteTokenBadRequest{}
}

/*DeleteTokenBadRequest handles this case with default header values.

Bad Request
*/
type DeleteTokenBadRequest struct {
	Payload *mono_models.Message
}

func (o *DeleteTokenBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apikeys/{tokenID}][%d] deleteTokenBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokenForbidden creates a DeleteTokenForbidden with default headers values
func NewDeleteTokenForbidden() *DeleteTokenForbidden {
	return &DeleteTokenForbidden{}
}

/*DeleteTokenForbidden handles this case with default header values.

Forbidden
*/
type DeleteTokenForbidden struct {
}

func (o *DeleteTokenForbidden) Error() string {
	return fmt.Sprintf("[DELETE /apikeys/{tokenID}][%d] deleteTokenForbidden ", 403)
}

func (o *DeleteTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}