// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-03-31

package client

import (
	"encoding/json"
)

// AllowlistEntry struct for AllowlistEntry.
type AllowlistEntry struct {
	CidrIp               string  `json:"cidr_ip"`
	CidrMask             *int32  `json:"cidr_mask,omitempty"`
	Ui                   bool    `json:"ui"`
	Sql                  bool    `json:"sql"`
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type allowlistEntry AllowlistEntry

// NewAllowlistEntry instantiates a new AllowlistEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowlistEntry(cidrIp string, ui bool, sql bool) *AllowlistEntry {
	p := AllowlistEntry{}
	p.CidrIp = cidrIp
	p.Ui = ui
	p.Sql = sql
	return &p
}

// NewAllowlistEntryWithDefaults instantiates a new AllowlistEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowlistEntryWithDefaults() *AllowlistEntry {
	p := AllowlistEntry{}
	return &p
}

// GetCidrIp returns the CidrIp field value.
func (o *AllowlistEntry) GetCidrIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CidrIp
}

// SetCidrIp sets field value.
func (o *AllowlistEntry) SetCidrIp(v string) {
	o.CidrIp = v
}

// GetCidrMask returns the CidrMask field value if set, zero value otherwise.
func (o *AllowlistEntry) GetCidrMask() int32 {
	if o == nil || o.CidrMask == nil {
		var ret int32
		return ret
	}
	return *o.CidrMask
}

// SetCidrMask gets a reference to the given int32 and assigns it to the CidrMask field.
func (o *AllowlistEntry) SetCidrMask(v int32) {
	o.CidrMask = &v
}

// GetUi returns the Ui field value.
func (o *AllowlistEntry) GetUi() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ui
}

// SetUi sets field value.
func (o *AllowlistEntry) SetUi(v bool) {
	o.Ui = v
}

// GetSql returns the Sql field value.
func (o *AllowlistEntry) GetSql() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sql
}

// SetSql sets field value.
func (o *AllowlistEntry) SetSql(v bool) {
	o.Sql = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AllowlistEntry) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AllowlistEntry) SetName(v string) {
	o.Name = &v
}

func (o AllowlistEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cidr_ip"] = o.CidrIp
	}
	if o.CidrMask != nil {
		toSerialize["cidr_mask"] = o.CidrMask
	}
	if true {
		toSerialize["ui"] = o.Ui
	}
	if true {
		toSerialize["sql"] = o.Sql
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AllowlistEntry) UnmarshalJSON(bytes []byte) (err error) {
	varAllowlistEntry := allowlistEntry{}

	if err = json.Unmarshal(bytes, &varAllowlistEntry); err == nil {
		*o = AllowlistEntry(varAllowlistEntry)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cidr_ip")
		delete(additionalProperties, "cidr_mask")
		delete(additionalProperties, "ui")
		delete(additionalProperties, "sql")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
