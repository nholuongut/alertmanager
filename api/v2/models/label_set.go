// Code generated by go-swagger; DO NOT EDIT.

// Copyright Nho luong DevOps
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// LabelSet label set
//
// swagger:model labelSet
type LabelSet map[string]string

// Validate validates this label set
func (m LabelSet) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this label set based on context it is used
func (m LabelSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
