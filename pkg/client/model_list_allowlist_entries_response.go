// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-07-12T00:00:00.000Z

package client

import (
	"encoding/json"
)

// ListAllowlistEntriesResponse struct for ListAllowlistEntriesResponse.
type ListAllowlistEntriesResponse struct {
	Allowlist            []AllowlistEntry          `json:"allowlist"`
	Propagating          bool                      `json:"propagating"`
	Pagination           *KeysetPaginationResponse `json:"pagination,omitempty"`
	AdditionalProperties map[string]interface{}
}

type listAllowlistEntriesResponse ListAllowlistEntriesResponse

// NewListAllowlistEntriesResponse instantiates a new ListAllowlistEntriesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllowlistEntriesResponse(allowlist []AllowlistEntry, propagating bool) *ListAllowlistEntriesResponse {
	p := ListAllowlistEntriesResponse{}
	p.Allowlist = allowlist
	p.Propagating = propagating
	return &p
}

// NewListAllowlistEntriesResponseWithDefaults instantiates a new ListAllowlistEntriesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllowlistEntriesResponseWithDefaults() *ListAllowlistEntriesResponse {
	p := ListAllowlistEntriesResponse{}
	return &p
}

// GetAllowlist returns the Allowlist field value.
func (o *ListAllowlistEntriesResponse) GetAllowlist() []AllowlistEntry {
	if o == nil {
		var ret []AllowlistEntry
		return ret
	}

	return o.Allowlist
}

// SetAllowlist sets field value.
func (o *ListAllowlistEntriesResponse) SetAllowlist(v []AllowlistEntry) {
	o.Allowlist = v
}

// GetPropagating returns the Propagating field value.
func (o *ListAllowlistEntriesResponse) GetPropagating() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Propagating
}

// SetPropagating sets field value.
func (o *ListAllowlistEntriesResponse) SetPropagating(v bool) {
	o.Propagating = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListAllowlistEntriesResponse) GetPagination() KeysetPaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret KeysetPaginationResponse
		return ret
	}
	return *o.Pagination
}

// SetPagination gets a reference to the given KeysetPaginationResponse and assigns it to the Pagination field.
func (o *ListAllowlistEntriesResponse) SetPagination(v KeysetPaginationResponse) {
	o.Pagination = &v
}

func (o ListAllowlistEntriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allowlist"] = o.Allowlist
	}
	if true {
		toSerialize["propagating"] = o.Propagating
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ListAllowlistEntriesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varListAllowlistEntriesResponse := listAllowlistEntriesResponse{}

	if err = json.Unmarshal(bytes, &varListAllowlistEntriesResponse); err == nil {
		*o = ListAllowlistEntriesResponse(varListAllowlistEntriesResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "allowlist")
		delete(additionalProperties, "propagating")
		delete(additionalProperties, "pagination")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
