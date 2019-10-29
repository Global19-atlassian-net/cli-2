// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	inventory_models "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// AddIngredientVersionAuthorReader is a Reader for the AddIngredientVersionAuthor structure.
type AddIngredientVersionAuthorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddIngredientVersionAuthorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddIngredientVersionAuthorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddIngredientVersionAuthorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddIngredientVersionAuthorDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddIngredientVersionAuthorOK creates a AddIngredientVersionAuthorOK with default headers values
func NewAddIngredientVersionAuthorOK() *AddIngredientVersionAuthorOK {
	return &AddIngredientVersionAuthorOK{}
}

/*AddIngredientVersionAuthorOK handles this case with default header values.

The author added to the ingredient version
*/
type AddIngredientVersionAuthorOK struct {
	Payload *inventory_models.V1Author
}

func (o *AddIngredientVersionAuthorOK) Error() string {
	return fmt.Sprintf("[POST /v1/ingredients/{ingredient_id}/versions/{ingredient_version_id}/authors][%d] addIngredientVersionAuthorOK  %+v", 200, o.Payload)
}

func (o *AddIngredientVersionAuthorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1Author)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddIngredientVersionAuthorBadRequest creates a AddIngredientVersionAuthorBadRequest with default headers values
func NewAddIngredientVersionAuthorBadRequest() *AddIngredientVersionAuthorBadRequest {
	return &AddIngredientVersionAuthorBadRequest{}
}

/*AddIngredientVersionAuthorBadRequest handles this case with default header values.

If the author ID doesn't exist
*/
type AddIngredientVersionAuthorBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddIngredientVersionAuthorBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/ingredients/{ingredient_id}/versions/{ingredient_version_id}/authors][%d] addIngredientVersionAuthorBadRequest  %+v", 400, o.Payload)
}

func (o *AddIngredientVersionAuthorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddIngredientVersionAuthorDefault creates a AddIngredientVersionAuthorDefault with default headers values
func NewAddIngredientVersionAuthorDefault(code int) *AddIngredientVersionAuthorDefault {
	return &AddIngredientVersionAuthorDefault{
		_statusCode: code,
	}
}

/*AddIngredientVersionAuthorDefault handles this case with default header values.

generic error response
*/
type AddIngredientVersionAuthorDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add ingredient version author default response
func (o *AddIngredientVersionAuthorDefault) Code() int {
	return o._statusCode
}

func (o *AddIngredientVersionAuthorDefault) Error() string {
	return fmt.Sprintf("[POST /v1/ingredients/{ingredient_id}/versions/{ingredient_version_id}/authors][%d] addIngredientVersionAuthor default  %+v", o._statusCode, o.Payload)
}

func (o *AddIngredientVersionAuthorDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}