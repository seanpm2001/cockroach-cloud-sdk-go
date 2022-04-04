// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-03-31

package client

import (
	"encoding/json"
)

// CreateClusterSpecification struct for CreateClusterSpecification.
type CreateClusterSpecification struct {
	Dedicated            *DedicatedClusterCreateSpecification `json:"dedicated,omitempty"`
	Serverless           *ServerlessClusterSpecification      `json:"serverless,omitempty"`
	AdditionalProperties map[string]interface{}
}

type createClusterSpecification CreateClusterSpecification

// NewCreateClusterSpecification instantiates a new CreateClusterSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterSpecification() *CreateClusterSpecification {
	p := CreateClusterSpecification{}
	return &p
}

// GetDedicated returns the Dedicated field value if set, zero value otherwise.
func (o *CreateClusterSpecification) GetDedicated() DedicatedClusterCreateSpecification {
	if o == nil || o.Dedicated == nil {
		var ret DedicatedClusterCreateSpecification
		return ret
	}
	return *o.Dedicated
}

// SetDedicated gets a reference to the given DedicatedClusterCreateSpecification and assigns it to the Dedicated field.
func (o *CreateClusterSpecification) SetDedicated(v DedicatedClusterCreateSpecification) {
	o.Dedicated = &v
}

// GetServerless returns the Serverless field value if set, zero value otherwise.
func (o *CreateClusterSpecification) GetServerless() ServerlessClusterSpecification {
	if o == nil || o.Serverless == nil {
		var ret ServerlessClusterSpecification
		return ret
	}
	return *o.Serverless
}

// SetServerless gets a reference to the given ServerlessClusterSpecification and assigns it to the Serverless field.
func (o *CreateClusterSpecification) SetServerless(v ServerlessClusterSpecification) {
	o.Serverless = &v
}

func (o CreateClusterSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dedicated != nil {
		toSerialize["dedicated"] = o.Dedicated
	}
	if o.Serverless != nil {
		toSerialize["serverless"] = o.Serverless
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateClusterSpecification) UnmarshalJSON(bytes []byte) (err error) {
	varCreateClusterSpecification := createClusterSpecification{}

	if err = json.Unmarshal(bytes, &varCreateClusterSpecification); err == nil {
		*o = CreateClusterSpecification(varCreateClusterSpecification)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dedicated")
		delete(additionalProperties, "serverless")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
