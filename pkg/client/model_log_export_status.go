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

// LogExportStatus LogExportStatus encodes the possible states that a configuration can be in as it is created, deployed, and disabled.
type LogExportStatus string

// List of LogExportStatus.
const (
	LOGEXPORTSTATUS_DISABLED       LogExportStatus = "DISABLED"
	LOGEXPORTSTATUS_DISABLING      LogExportStatus = "DISABLING"
	LOGEXPORTSTATUS_DISABLE_FAILED LogExportStatus = "DISABLE_FAILED"
	LOGEXPORTSTATUS_ENABLED        LogExportStatus = "ENABLED"
	LOGEXPORTSTATUS_ENABLING       LogExportStatus = "ENABLING"
	LOGEXPORTSTATUS_ENABLE_FAILED  LogExportStatus = "ENABLE_FAILED"
)

// All allowed values of LogExportStatus enum.
var AllowedLogExportStatusEnumValues = []LogExportStatus{
	"DISABLED",
	"DISABLING",
	"DISABLE_FAILED",
	"ENABLED",
	"ENABLING",
	"ENABLE_FAILED",
}

func (v *LogExportStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogExportStatus(value)
	for _, existing := range AllowedLogExportStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogExportStatus", value)
}

// NewLogExportStatusFromValue returns a pointer to a valid LogExportStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogExportStatusFromValue(v string) (*LogExportStatus, error) {
	ev := LogExportStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogExportStatus: valid values are %v", v, AllowedLogExportStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogExportStatus) IsValid() bool {
	for _, existing := range AllowedLogExportStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogExportStatus value.
func (v LogExportStatus) Ptr() *LogExportStatus {
	return &v
}
