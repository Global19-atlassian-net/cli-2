// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UserEditable user editable
// swagger:model UserEditable
type UserEditable struct {

	// e u l a accepted
	EULAAccepted *bool `json:"EULAAccepted,omitempty"`

	// datetime format
	DatetimeFormat string `json:"datetimeFormat,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// invite code
	InviteCode *string `json:"inviteCode,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// send marketing email
	SendMarketingEmail *bool `json:"sendMarketingEmail,omitempty"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this user editable
func (m *UserEditable) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserEditable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserEditable) UnmarshalBinary(b []byte) error {
	var res UserEditable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
