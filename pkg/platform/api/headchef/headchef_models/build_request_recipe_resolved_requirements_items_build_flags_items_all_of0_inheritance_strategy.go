// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategy The strategy describes how this flag is inherited from other build ingredients. If this is not set then the flag is entirely independent.
// swagger:model buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategy
type BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategy struct {

	// The name of the dependency from which to get the flag value when the strategy is 'dependency'.
	DependencyName string `json:"dependency_name,omitempty"`

	// If the strategy is 'language_core' then it is derived by looking for a flag of the same name in the language core build flags. If the strategy is 'dependency' then it looks at the value of a dependency matching a certain name in the build. If this is omitted then it is set unconditionally based on the recipe's settings for this ingredient version
	// Required: true
	// Enum: [dependency language_core]
	Strategy *string `json:"strategy"`
}

// Validate validates this build request recipe resolved requirements items build flags items all of0 inheritance strategy
func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategyTypeStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dependency","language_core"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategyTypeStrategyPropEnum = append(buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategyTypeStrategyPropEnum, v)
	}
}

const (

	// BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategyStrategyDependency captures enum value "dependency"
	BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategyStrategyDependency string = "dependency"

	// BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategyStrategyLanguageCore captures enum value "language_core"
	BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategyStrategyLanguageCore string = "language_core"
)

// prop value enum
func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategy) validateStrategyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategyTypeStrategyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategy) validateStrategy(formats strfmt.Registry) error {

	if err := validate.Required("strategy", "body", m.Strategy); err != nil {
		return err
	}

	// value enum
	if err := m.validateStrategyEnum("strategy", "body", *m.Strategy); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategy) UnmarshalBinary(b []byte) error {
	var res BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
