// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatusMessageEnvelopeBody10 Package Build Skipped
//
// A message indicating that a requested package build has been skipped.
// swagger:model statusMessageEnvelopeBody10
type StatusMessageEnvelopeBody10 struct {

	// ingredient id
	// Required: true
	// Format: uuid
	IngredientID *strfmt.UUID `json:"ingredient_id"`

	// ingredient version
	// Required: true
	IngredientVersion *string `json:"ingredient_version"`

	// reason
	// Required: true
	Reason *string `json:"reason"`
}

// Validate validates this status message envelope body10
func (m *StatusMessageEnvelopeBody10) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIngredientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusMessageEnvelopeBody10) validateIngredientID(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_id", "body", m.IngredientID); err != nil {
		return err
	}

	if err := validate.FormatOf("ingredient_id", "body", "uuid", m.IngredientID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StatusMessageEnvelopeBody10) validateIngredientVersion(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_version", "body", m.IngredientVersion); err != nil {
		return err
	}

	return nil
}

func (m *StatusMessageEnvelopeBody10) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusMessageEnvelopeBody10) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusMessageEnvelopeBody10) UnmarshalBinary(b []byte) error {
	var res StatusMessageEnvelopeBody10
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
