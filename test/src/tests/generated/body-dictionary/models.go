package dictionarygroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// Error ...
type Error struct {
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// SetBase64URL ...
type SetBase64URL struct {
	autorest.Response `json:"-"`
	Value             map[string]*string `json:"value"`
}

// MarshalJSON is the custom marshaler for SetBase64URL.
func (sb6u SetBase64URL) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sb6u.Value != nil {
		objectMap["value"] = sb6u.Value
	}
	return json.Marshal(objectMap)
}

// SetBool ...
type SetBool struct {
	autorest.Response `json:"-"`
	Value             map[string]*bool `json:"value"`
}

// MarshalJSON is the custom marshaler for SetBool.
func (sb SetBool) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sb.Value != nil {
		objectMap["value"] = sb.Value
	}
	return json.Marshal(objectMap)
}

// SetByteArray ...
type SetByteArray struct {
	autorest.Response `json:"-"`
	Value             map[string][]byte `json:"value"`
}

// MarshalJSON is the custom marshaler for SetByteArray.
func (sba SetByteArray) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sba.Value != nil {
		objectMap["value"] = sba.Value
	}
	return json.Marshal(objectMap)
}

// SetDate ...
type SetDate struct {
	autorest.Response `json:"-"`
	Value             map[string]*date.Date `json:"value"`
}

// MarshalJSON is the custom marshaler for SetDate.
func (sd SetDate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sd.Value != nil {
		objectMap["value"] = sd.Value
	}
	return json.Marshal(objectMap)
}

// SetDateTime ...
type SetDateTime struct {
	autorest.Response `json:"-"`
	Value             map[string]*date.Time `json:"value"`
}

// MarshalJSON is the custom marshaler for SetDateTime.
func (sdt SetDateTime) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sdt.Value != nil {
		objectMap["value"] = sdt.Value
	}
	return json.Marshal(objectMap)
}

// SetDateTimeRfc1123 ...
type SetDateTimeRfc1123 struct {
	autorest.Response `json:"-"`
	Value             map[string]*date.TimeRFC1123 `json:"value"`
}

// MarshalJSON is the custom marshaler for SetDateTimeRfc1123.
func (sdtr1 SetDateTimeRfc1123) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sdtr1.Value != nil {
		objectMap["value"] = sdtr1.Value
	}
	return json.Marshal(objectMap)
}

// SetFloat64 ...
type SetFloat64 struct {
	autorest.Response `json:"-"`
	Value             map[string]*float64 `json:"value"`
}

// MarshalJSON is the custom marshaler for SetFloat64.
func (sf6 SetFloat64) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sf6.Value != nil {
		objectMap["value"] = sf6.Value
	}
	return json.Marshal(objectMap)
}

// SetInt32 ...
type SetInt32 struct {
	autorest.Response `json:"-"`
	Value             map[string]*int32 `json:"value"`
}

// MarshalJSON is the custom marshaler for SetInt32.
func (si3 SetInt32) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if si3.Value != nil {
		objectMap["value"] = si3.Value
	}
	return json.Marshal(objectMap)
}

// SetInt64 ...
type SetInt64 struct {
	autorest.Response `json:"-"`
	Value             map[string]*int64 `json:"value"`
}

// MarshalJSON is the custom marshaler for SetInt64.
func (si6 SetInt64) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if si6.Value != nil {
		objectMap["value"] = si6.Value
	}
	return json.Marshal(objectMap)
}

// SetListString ...
type SetListString struct {
	autorest.Response `json:"-"`
	Value             map[string][]string `json:"value"`
}

// MarshalJSON is the custom marshaler for SetListString.
func (sls SetListString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sls.Value != nil {
		objectMap["value"] = sls.Value
	}
	return json.Marshal(objectMap)
}

// SetSetString ...
type SetSetString struct {
	autorest.Response `json:"-"`
	Value             map[string]map[string]*string `json:"value"`
}

// MarshalJSON is the custom marshaler for SetSetString.
func (sss SetSetString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sss.Value != nil {
		objectMap["value"] = sss.Value
	}
	return json.Marshal(objectMap)
}

// SetString ...
type SetString struct {
	autorest.Response `json:"-"`
	Value             map[string]*string `json:"value"`
}

// MarshalJSON is the custom marshaler for SetString.
func (ss SetString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ss.Value != nil {
		objectMap["value"] = ss.Value
	}
	return json.Marshal(objectMap)
}

// SetTimeSpan ...
type SetTimeSpan struct {
	autorest.Response `json:"-"`
	Value             map[string]*string `json:"value"`
}

// MarshalJSON is the custom marshaler for SetTimeSpan.
func (sts SetTimeSpan) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sts.Value != nil {
		objectMap["value"] = sts.Value
	}
	return json.Marshal(objectMap)
}

// SetWidget ...
type SetWidget struct {
	autorest.Response `json:"-"`
	Value             map[string]*Widget `json:"value"`
}

// MarshalJSON is the custom marshaler for SetWidget.
func (sw SetWidget) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sw.Value != nil {
		objectMap["value"] = sw.Value
	}
	return json.Marshal(objectMap)
}

// Widget ...
type Widget struct {
	Integer *int32  `json:"integer,omitempty"`
	String  *string `json:"string,omitempty"`
}
