// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1 Revision
//
// The properties of any revisioned resource which can be modified by a new revision
// swagger:model v1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1
type V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1 struct {

	// Whether this revision should be considered 'stable'. When a new stable revision is created, it supercedes any existing stable revision and becomes the default revision of the revisioned resource going forward.
	IsStableRevision *bool `json:"is_stable_revision,omitempty"`

	// provided features
	// Required: true
	ProvidedFeatures []*V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1ProvidedFeaturesItems `json:"provided_features"`
}

// Validate validates this v1 build request recipe platform operating system version all of1 all of1
func (m *V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvidedFeatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1) validateProvidedFeatures(formats strfmt.Registry) error {

	if err := validate.Required("provided_features", "body", m.ProvidedFeatures); err != nil {
		return err
	}

	for i := 0; i < len(m.ProvidedFeatures); i++ {
		if swag.IsZero(m.ProvidedFeatures[i]) { // not required
			continue
		}

		if m.ProvidedFeatures[i] != nil {
			if err := m.ProvidedFeatures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("provided_features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1) UnmarshalBinary(b []byte) error {
	var res V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
