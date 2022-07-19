// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-07-12T00:00:00.000Z

package client

import (
	"encoding/json"
	"time"
)

// CMEKKeyInfo CMEKKeyInfo contains the status of a customer-provided key alongside the specification..
type CMEKKeyInfo struct {
	Status               *CMEKStatus           `json:"status,omitempty"`
	UserMessage          *string               `json:"user_message,omitempty"`
	Spec                 *CMEKKeySpecification `json:"spec,omitempty"`
	CreatedAt            *time.Time            `json:"created_at,omitempty"`
	UpdatedAt            *time.Time            `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type cMEKKeyInfo CMEKKeyInfo

// NewCMEKKeyInfo instantiates a new CMEKKeyInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCMEKKeyInfo() *CMEKKeyInfo {
	p := CMEKKeyInfo{}
	return &p
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CMEKKeyInfo) GetStatus() CMEKStatus {
	if o == nil || o.Status == nil {
		var ret CMEKStatus
		return ret
	}
	return *o.Status
}

// SetStatus gets a reference to the given CMEKStatus and assigns it to the Status field.
func (o *CMEKKeyInfo) SetStatus(v CMEKStatus) {
	o.Status = &v
}

// GetUserMessage returns the UserMessage field value if set, zero value otherwise.
func (o *CMEKKeyInfo) GetUserMessage() string {
	if o == nil || o.UserMessage == nil {
		var ret string
		return ret
	}
	return *o.UserMessage
}

// SetUserMessage gets a reference to the given string and assigns it to the UserMessage field.
func (o *CMEKKeyInfo) SetUserMessage(v string) {
	o.UserMessage = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *CMEKKeyInfo) GetSpec() CMEKKeySpecification {
	if o == nil || o.Spec == nil {
		var ret CMEKKeySpecification
		return ret
	}
	return *o.Spec
}

// SetSpec gets a reference to the given CMEKKeySpecification and assigns it to the Spec field.
func (o *CMEKKeyInfo) SetSpec(v CMEKKeySpecification) {
	o.Spec = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CMEKKeyInfo) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CMEKKeyInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CMEKKeyInfo) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CMEKKeyInfo) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o CMEKKeyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UserMessage != nil {
		toSerialize["user_message"] = o.UserMessage
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CMEKKeyInfo) UnmarshalJSON(bytes []byte) (err error) {
	varCMEKKeyInfo := cMEKKeyInfo{}

	if err = json.Unmarshal(bytes, &varCMEKKeyInfo); err == nil {
		*o = CMEKKeyInfo(varCMEKKeyInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "user_message")
		delete(additionalProperties, "spec")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
