// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Diff diff
// swagger:model diff
type Diff struct {

	// direction
	Direction string `json:"direction,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// path type
	PathType string `json:"path_type,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this diff
func (m *Diff) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Diff) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Diff) UnmarshalBinary(b []byte) error {
	var res Diff
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
