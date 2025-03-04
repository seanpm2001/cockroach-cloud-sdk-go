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

import (
	"encoding/json"
	"fmt"
)

// ClusterVersionDeferralPolicyType the model 'ClusterVersionDeferralPolicyType'.
type ClusterVersionDeferralPolicyType string

// List of ClusterVersionDeferralPolicy.Type.
const (
	CLUSTERVERSIONDEFERRALPOLICYTYPE_NOT_DEFERRED   ClusterVersionDeferralPolicyType = "NOT_DEFERRED"
	CLUSTERVERSIONDEFERRALPOLICYTYPE_FIXED_DEFERRAL ClusterVersionDeferralPolicyType = "FIXED_DEFERRAL"
)

// All allowed values of ClusterVersionDeferralPolicyType enum.
var AllowedClusterVersionDeferralPolicyTypeEnumValues = []ClusterVersionDeferralPolicyType{
	"NOT_DEFERRED",
	"FIXED_DEFERRAL",
}

func (v *ClusterVersionDeferralPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterVersionDeferralPolicyType(value)
	for _, existing := range AllowedClusterVersionDeferralPolicyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterVersionDeferralPolicyType", value)
}

// NewClusterVersionDeferralPolicyTypeFromValue returns a pointer to a valid ClusterVersionDeferralPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClusterVersionDeferralPolicyTypeFromValue(v string) (*ClusterVersionDeferralPolicyType, error) {
	ev := ClusterVersionDeferralPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterVersionDeferralPolicyType: valid values are %v", v, AllowedClusterVersionDeferralPolicyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClusterVersionDeferralPolicyType) IsValid() bool {
	for _, existing := range AllowedClusterVersionDeferralPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterVersionDeferralPolicy.Type value.
func (v ClusterVersionDeferralPolicyType) Ptr() *ClusterVersionDeferralPolicyType {
	return &v
}
