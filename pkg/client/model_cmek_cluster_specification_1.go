// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-07-12T00:00:00.000Z

package client

import (
	"encoding/json"
)

// CMEKClusterSpecification1 struct for CMEKClusterSpecification1.
type CMEKClusterSpecification1 struct {
	RegionSpecs          []CMEKRegionSpecification `json:"region_specs"`
	AdditionalProperties map[string]interface{}
}

type cMEKClusterSpecification1 CMEKClusterSpecification1

// NewCMEKClusterSpecification1 instantiates a new CMEKClusterSpecification1 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCMEKClusterSpecification1(regionSpecs []CMEKRegionSpecification) *CMEKClusterSpecification1 {
	p := CMEKClusterSpecification1{}
	p.RegionSpecs = regionSpecs
	return &p
}

// NewCMEKClusterSpecification1WithDefaults instantiates a new CMEKClusterSpecification1 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCMEKClusterSpecification1WithDefaults() *CMEKClusterSpecification1 {
	p := CMEKClusterSpecification1{}
	return &p
}

// GetRegionSpecs returns the RegionSpecs field value.
func (o *CMEKClusterSpecification1) GetRegionSpecs() []CMEKRegionSpecification {
	if o == nil {
		var ret []CMEKRegionSpecification
		return ret
	}

	return o.RegionSpecs
}

// SetRegionSpecs sets field value.
func (o *CMEKClusterSpecification1) SetRegionSpecs(v []CMEKRegionSpecification) {
	o.RegionSpecs = v
}

func (o CMEKClusterSpecification1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["region_specs"] = o.RegionSpecs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CMEKClusterSpecification1) UnmarshalJSON(bytes []byte) (err error) {
	varCMEKClusterSpecification1 := cMEKClusterSpecification1{}

	if err = json.Unmarshal(bytes, &varCMEKClusterSpecification1); err == nil {
		*o = CMEKClusterSpecification1(varCMEKClusterSpecification1)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "region_specs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
