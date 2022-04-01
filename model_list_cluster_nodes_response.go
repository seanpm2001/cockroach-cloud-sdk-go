/*
CockroachDB Cloud API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2021-12-28
Contact: support@cockroachlabs.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ListClusterNodesResponse struct for ListClusterNodesResponse
type ListClusterNodesResponse struct {
	Nodes []Node `json:"nodes"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewListClusterNodesResponse instantiates a new ListClusterNodesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListClusterNodesResponse(nodes []Node) *ListClusterNodesResponse {
	this := ListClusterNodesResponse{}
	this.Nodes = nodes
	return &this
}

// NewListClusterNodesResponseWithDefaults instantiates a new ListClusterNodesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListClusterNodesResponseWithDefaults() *ListClusterNodesResponse {
	this := ListClusterNodesResponse{}
	return &this
}

// GetNodes returns the Nodes field value
func (o *ListClusterNodesResponse) GetNodes() []Node {
	if o == nil {
		var ret []Node
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *ListClusterNodesResponse) GetNodesOk() ([]Node, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *ListClusterNodesResponse) SetNodes(v []Node) {
	o.Nodes = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListClusterNodesResponse) GetPagination() PaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterNodesResponse) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListClusterNodesResponse) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *ListClusterNodesResponse) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o ListClusterNodesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nodes"] = o.Nodes
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListClusterNodesResponse struct {
	value *ListClusterNodesResponse
	isSet bool
}

func (v NullableListClusterNodesResponse) Get() *ListClusterNodesResponse {
	return v.value
}

func (v *NullableListClusterNodesResponse) Set(val *ListClusterNodesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListClusterNodesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListClusterNodesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListClusterNodesResponse(val *ListClusterNodesResponse) *NullableListClusterNodesResponse {
	return &NullableListClusterNodesResponse{value: val, isSet: true}
}

func (v NullableListClusterNodesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListClusterNodesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


