// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1PlatformPagedListPlatformsItemsImagesItems Image
//
// The full image data model
// swagger:model v1PlatformPagedListPlatformsItemsImagesItems
type V1PlatformPagedListPlatformsItemsImagesItems struct {
	V1PlatformPagedListPlatformsItemsImagesItemsAllOf0

	V1PlatformPagedListPlatformsItemsImagesItemsAllOf1

	V1PlatformPagedListPlatformsItemsImagesItemsAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1PlatformPagedListPlatformsItemsImagesItems) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1PlatformPagedListPlatformsItemsImagesItemsAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1PlatformPagedListPlatformsItemsImagesItemsAllOf0 = aO0

	// AO1
	var aO1 V1PlatformPagedListPlatformsItemsImagesItemsAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1PlatformPagedListPlatformsItemsImagesItemsAllOf1 = aO1

	// AO2
	var aO2 V1PlatformPagedListPlatformsItemsImagesItemsAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.V1PlatformPagedListPlatformsItemsImagesItemsAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1PlatformPagedListPlatformsItemsImagesItems) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.V1PlatformPagedListPlatformsItemsImagesItemsAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1PlatformPagedListPlatformsItemsImagesItemsAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.V1PlatformPagedListPlatformsItemsImagesItemsAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 platform paged list platforms items images items
func (m *V1PlatformPagedListPlatformsItemsImagesItems) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1PlatformPagedListPlatformsItemsImagesItemsAllOf0
	if err := m.V1PlatformPagedListPlatformsItemsImagesItemsAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1PlatformPagedListPlatformsItemsImagesItemsAllOf1
	if err := m.V1PlatformPagedListPlatformsItemsImagesItemsAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1PlatformPagedListPlatformsItemsImagesItemsAllOf2
	if err := m.V1PlatformPagedListPlatformsItemsImagesItemsAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1PlatformPagedListPlatformsItemsImagesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PlatformPagedListPlatformsItemsImagesItems) UnmarshalBinary(b []byte) error {
	var res V1PlatformPagedListPlatformsItemsImagesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}