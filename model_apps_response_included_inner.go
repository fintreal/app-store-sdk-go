/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// AppsResponseIncludedInner - struct for AppsResponseIncludedInner
type AppsResponseIncludedInner struct {
	AppInfo *AppInfo
}

// AppInfoAsAppsResponseIncludedInner is a convenience function that returns AppInfo wrapped in AppsResponseIncludedInner
func AppInfoAsAppsResponseIncludedInner(v *AppInfo) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		AppInfo: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppInfo
	err = newStrictDecoder(data).Decode(&dst.AppInfo)
	if err == nil {
		jsonAppInfo, _ := json.Marshal(dst.AppInfo)
		if string(jsonAppInfo) == "{}" { // empty struct
			dst.AppInfo = nil
		} else {
			if err = validator.Validate(dst.AppInfo); err != nil {
				dst.AppInfo = nil
			} else {
				match++
			}
		}
	} else {
		dst.AppInfo = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppInfo = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppInfo != nil {
		return json.Marshal(&src.AppInfo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppInfo != nil {
		return obj.AppInfo
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AppsResponseIncludedInner) GetActualInstanceValue() (interface{}) {
	if obj.AppInfo != nil {
		return *obj.AppInfo
	}

	// all schemas are nil
	return nil
}

type NullableAppsResponseIncludedInner struct {
	value *AppsResponseIncludedInner
	isSet bool
}

func (v NullableAppsResponseIncludedInner) Get() *AppsResponseIncludedInner {
	return v.value
}

func (v *NullableAppsResponseIncludedInner) Set(val *AppsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsResponseIncludedInner(val *AppsResponseIncludedInner) *NullableAppsResponseIncludedInner {
	return &NullableAppsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


