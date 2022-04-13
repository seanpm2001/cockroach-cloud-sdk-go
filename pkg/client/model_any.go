// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-03-31

package client

import (
	"encoding/json"
)

// Any struct for Any.
type Any struct {
	Type *string `json:"@type,omitempty"`
}

// NewAny instantiates a new Any object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAny() *Any {
	p := Any{}
	return &p
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Any) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Any) SetType(v string) {
	o.Type = &v
}

func (o Any) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["@type"] = o.Type
	}
	return json.Marshal(toSerialize)
}
