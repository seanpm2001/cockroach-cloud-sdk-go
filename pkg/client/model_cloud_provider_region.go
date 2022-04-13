// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-03-31

package client

import (
	"encoding/json"
)

// CloudProviderRegion struct for CloudProviderRegion.
type CloudProviderRegion struct {
	Name                 string           `json:"name"`
	Location             string           `json:"location"`
	Provider             ApiCloudProvider `json:"provider"`
	Serverless           bool             `json:"serverless"`
	Distance             float32          `json:"distance"`
	AdditionalProperties map[string]interface{}
}

type cloudProviderRegion CloudProviderRegion

// NewCloudProviderRegion instantiates a new CloudProviderRegion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderRegion(name string, location string, provider ApiCloudProvider, serverless bool, distance float32) *CloudProviderRegion {
	p := CloudProviderRegion{}
	p.Name = name
	p.Location = location
	p.Provider = provider
	p.Serverless = serverless
	p.Distance = distance
	return &p
}

// NewCloudProviderRegionWithDefaults instantiates a new CloudProviderRegion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderRegionWithDefaults() *CloudProviderRegion {
	p := CloudProviderRegion{}
	var provider ApiCloudProvider = APICLOUDPROVIDER_UNSPECIFIED
	p.Provider = provider
	return &p
}

// GetName returns the Name field value.
func (o *CloudProviderRegion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value.
func (o *CloudProviderRegion) SetName(v string) {
	o.Name = v
}

// GetLocation returns the Location field value.
func (o *CloudProviderRegion) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// SetLocation sets field value.
func (o *CloudProviderRegion) SetLocation(v string) {
	o.Location = v
}

// GetProvider returns the Provider field value.
func (o *CloudProviderRegion) GetProvider() ApiCloudProvider {
	if o == nil {
		var ret ApiCloudProvider
		return ret
	}

	return o.Provider
}

// SetProvider sets field value.
func (o *CloudProviderRegion) SetProvider(v ApiCloudProvider) {
	o.Provider = v
}

// GetServerless returns the Serverless field value.
func (o *CloudProviderRegion) GetServerless() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Serverless
}

// SetServerless sets field value.
func (o *CloudProviderRegion) SetServerless(v bool) {
	o.Serverless = v
}

// GetDistance returns the Distance field value.
func (o *CloudProviderRegion) GetDistance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Distance
}

// SetDistance sets field value.
func (o *CloudProviderRegion) SetDistance(v float32) {
	o.Distance = v
}

func (o CloudProviderRegion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["serverless"] = o.Serverless
	}
	if true {
		toSerialize["distance"] = o.Distance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudProviderRegion) UnmarshalJSON(bytes []byte) (err error) {
	varCloudProviderRegion := cloudProviderRegion{}

	if err = json.Unmarshal(bytes, &varCloudProviderRegion); err == nil {
		*o = CloudProviderRegion(varCloudProviderRegion)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "location")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "serverless")
		delete(additionalProperties, "distance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
