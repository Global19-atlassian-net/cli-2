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

// IngredientVersionsReader is a Reader for the IngredientVersions structure.
type IngredientVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IngredientVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIngredientVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewIngredientVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIngredientVersionsOK creates a IngredientVersionsOK with default headers values
func NewIngredientVersionsOK() *IngredientVersionsOK {
	return &IngredientVersionsOK{}
}

/*IngredientVersionsOK handles this case with default header values.

Returns information about all available version for a given ingredient.
*/
type IngredientVersionsOK struct {
	Payload []*inventory_models.IngredientVersion
}

func (o *IngredientVersionsOK) Error() string {
	return fmt.Sprintf("[GET /ingredients/{ingredient_id}/versions][%d] ingredientVersionsOK  %+v", 200, o.Payload)
}

func (o *IngredientVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIngredientVersionsDefault creates a IngredientVersionsDefault with default headers values
func NewIngredientVersionsDefault(code int) *IngredientVersionsDefault {
	return &IngredientVersionsDefault{
		_statusCode: code,
	}
}

/*IngredientVersionsDefault handles this case with default header values.

generic error response
*/
type IngredientVersionsDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the ingredient versions default response
func (o *IngredientVersionsDefault) Code() int {
	return o._statusCode
}

func (o *IngredientVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /ingredients/{ingredient_id}/versions][%d] ingredientVersions default  %+v", o._statusCode, o.Payload)
}

func (o *IngredientVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}