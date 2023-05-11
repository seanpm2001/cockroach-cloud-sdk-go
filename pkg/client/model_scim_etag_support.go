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

// ScimEtagSupport struct for ScimEtagSupport.
type ScimEtagSupport struct {
	MaxResults *int32 `json:"maxResults,omitempty"`
	Supported  *bool  `json:"supported,omitempty"`
}

// NewScimEtagSupport instantiates a new ScimEtagSupport object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimEtagSupport() *ScimEtagSupport {
	p := ScimEtagSupport{}
	return &p
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *ScimEtagSupport) GetMaxResults() int32 {
	if o == nil || o.MaxResults == nil {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *ScimEtagSupport) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetSupported returns the Supported field value if set, zero value otherwise.
func (o *ScimEtagSupport) GetSupported() bool {
	if o == nil || o.Supported == nil {
		var ret bool
		return ret
	}
	return *o.Supported
}

// SetSupported gets a reference to the given bool and assigns it to the Supported field.
func (o *ScimEtagSupport) SetSupported(v bool) {
	o.Supported = &v
}
