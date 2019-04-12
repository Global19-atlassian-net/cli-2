// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	secrets_models "github.com/ActiveState/cli/pkg/platform/api/secrets/secrets_models"
)

// GetAllUserSecretsReader is a Reader for the GetAllUserSecrets structure.
type GetAllUserSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllUserSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllUserSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAllUserSecretsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAllUserSecretsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllUserSecretsOK creates a GetAllUserSecretsOK with default headers values
func NewGetAllUserSecretsOK() *GetAllUserSecretsOK {
	return &GetAllUserSecretsOK{}
}

/*GetAllUserSecretsOK handles this case with default header values.

Success
*/
type GetAllUserSecretsOK struct {
	Payload []*secrets_models.UserSecret
}

func (o *GetAllUserSecretsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationID}/user_secrets][%d] getAllUserSecretsOK  %+v", 200, o.Payload)
}

func (o *GetAllUserSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllUserSecretsUnauthorized creates a GetAllUserSecretsUnauthorized with default headers values
func NewGetAllUserSecretsUnauthorized() *GetAllUserSecretsUnauthorized {
	return &GetAllUserSecretsUnauthorized{}
}

/*GetAllUserSecretsUnauthorized handles this case with default header values.

Invalid credentials
*/
type GetAllUserSecretsUnauthorized struct {
	Payload *secrets_models.Message
}

func (o *GetAllUserSecretsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationID}/user_secrets][%d] getAllUserSecretsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllUserSecretsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(secrets_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllUserSecretsInternalServerError creates a GetAllUserSecretsInternalServerError with default headers values
func NewGetAllUserSecretsInternalServerError() *GetAllUserSecretsInternalServerError {
	return &GetAllUserSecretsInternalServerError{}
}

/*GetAllUserSecretsInternalServerError handles this case with default header values.

Server Error
*/
type GetAllUserSecretsInternalServerError struct {
	Payload *secrets_models.Message
}

func (o *GetAllUserSecretsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationID}/user_secrets][%d] getAllUserSecretsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllUserSecretsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(secrets_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}