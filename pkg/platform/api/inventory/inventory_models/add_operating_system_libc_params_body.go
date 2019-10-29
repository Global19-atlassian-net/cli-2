// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddOperatingSystemLibcParamsBody add operating system libc params body
// swagger:model addOperatingSystemLibcParamsBody
type AddOperatingSystemLibcParamsBody struct {

	// The ID of the libc that can be used with this operating system
	// Required: true
	// Format: uuid
	LibcID *strfmt.UUID `json:"libc_id"`
}

// Validate validates this add operating system libc params body
func (m *AddOperatingSystemLibcParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLibcID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddOperatingSystemLibcParamsBody) validateLibcID(formats strfmt.Registry) error {

	if err := validate.Required("libc_id", "body", m.LibcID); err != nil {
		return err
	}

	if err := validate.FormatOf("libc_id", "body", "uuid", m.LibcID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddOperatingSystemLibcParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddOperatingSystemLibcParamsBody) UnmarshalBinary(b []byte) error {
	var res AddOperatingSystemLibcParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}