// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-07-12T00:00:00.000Z

package client

import (
	"encoding/json"
)

// CMEKClusterSpecification struct for CMEKClusterSpecification.
type CMEKClusterSpecification struct {
	RegionSpecs          []CMEKRegionSpecification `json:"region_specs"`
	AdditionalProperties map[string]interface{}
}

type cMEKClusterSpecification CMEKClusterSpecification

// NewCMEKClusterSpecification instantiates a new CMEKClusterSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCMEKClusterSpecification(regionSpecs []CMEKRegionSpecification) *CMEKClusterSpecification {
	p := CMEKClusterSpecification{}
	p.RegionSpecs = regionSpecs
	return &p
}

// NewCMEKClusterSpecificationWithDefaults instantiates a new CMEKClusterSpecification object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCMEKClusterSpecificationWithDefaults() *CMEKClusterSpecification {
	p := CMEKClusterSpecification{}
	return &p
}

// GetRegionSpecs returns the RegionSpecs field value.
func (o *CMEKClusterSpecification) GetRegionSpecs() []CMEKRegionSpecification {
	if o == nil {
		var ret []CMEKRegionSpecification
		return ret
	}

	return o.RegionSpecs
}

// SetRegionSpecs sets field value.
func (o *CMEKClusterSpecification) SetRegionSpecs(v []CMEKRegionSpecification) {
	o.RegionSpecs = v
}

func (o CMEKClusterSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["region_specs"] = o.RegionSpecs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CMEKClusterSpecification) UnmarshalJSON(bytes []byte) (err error) {
	varCMEKClusterSpecification := cMEKClusterSpecification{}

	if err = json.Unmarshal(bytes, &varCMEKClusterSpecification); err == nil {
		*o = CMEKClusterSpecification(varCMEKClusterSpecification)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "region_specs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
