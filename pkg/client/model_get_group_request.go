// Copyright 2023 The Cockroach Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2023-04-10

package client

// GetGroupRequest struct for GetGroupRequest.
type GetGroupRequest struct {
	Attributes         *string `json:"attributes,omitempty"`
	ExcludedAttributes *string `json:"excludedAttributes,omitempty"`
}

// NewGetGroupRequest instantiates a new GetGroupRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGroupRequest() *GetGroupRequest {
	p := GetGroupRequest{}
	return &p
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GetGroupRequest) GetAttributes() string {
	if o == nil || o.Attributes == nil {
		var ret string
		return ret
	}
	return *o.Attributes
}

// SetAttributes gets a reference to the given string and assigns it to the Attributes field.
func (o *GetGroupRequest) SetAttributes(v string) {
	o.Attributes = &v
}

// GetExcludedAttributes returns the ExcludedAttributes field value if set, zero value otherwise.
func (o *GetGroupRequest) GetExcludedAttributes() string {
	if o == nil || o.ExcludedAttributes == nil {
		var ret string
		return ret
	}
	return *o.ExcludedAttributes
}

// SetExcludedAttributes gets a reference to the given string and assigns it to the ExcludedAttributes field.
func (o *GetGroupRequest) SetExcludedAttributes(v string) {
	o.ExcludedAttributes = &v
}
