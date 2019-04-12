// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TokenEditable token editable
// swagger:model TokenEditable
type TokenEditable struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this token editable
func (m *TokenEditable) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TokenEditable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenEditable) UnmarshalBinary(b []byte) error {
	var res TokenEditable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}