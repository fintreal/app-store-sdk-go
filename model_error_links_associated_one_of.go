/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ErrorLinksAssociatedOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorLinksAssociatedOneOf{}

// ErrorLinksAssociatedOneOf struct for ErrorLinksAssociatedOneOf
type ErrorLinksAssociatedOneOf struct {
	Href *string `json:"href,omitempty"`
	Meta *ErrorLinksAssociatedOneOfMeta `json:"meta,omitempty"`
}

// NewErrorLinksAssociatedOneOf instantiates a new ErrorLinksAssociatedOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorLinksAssociatedOneOf() *ErrorLinksAssociatedOneOf {
	this := ErrorLinksAssociatedOneOf{}
	return &this
}

// NewErrorLinksAssociatedOneOfWithDefaults instantiates a new ErrorLinksAssociatedOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorLinksAssociatedOneOfWithDefaults() *ErrorLinksAssociatedOneOf {
	this := ErrorLinksAssociatedOneOf{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ErrorLinksAssociatedOneOf) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLinksAssociatedOneOf) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ErrorLinksAssociatedOneOf) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ErrorLinksAssociatedOneOf) SetHref(v string) {
	o.Href = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ErrorLinksAssociatedOneOf) GetMeta() ErrorLinksAssociatedOneOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ErrorLinksAssociatedOneOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLinksAssociatedOneOf) GetMetaOk() (*ErrorLinksAssociatedOneOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ErrorLinksAssociatedOneOf) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ErrorLinksAssociatedOneOfMeta and assigns it to the Meta field.
func (o *ErrorLinksAssociatedOneOf) SetMeta(v ErrorLinksAssociatedOneOfMeta) {
	o.Meta = &v
}

func (o ErrorLinksAssociatedOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorLinksAssociatedOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableErrorLinksAssociatedOneOf struct {
	value *ErrorLinksAssociatedOneOf
	isSet bool
}

func (v NullableErrorLinksAssociatedOneOf) Get() *ErrorLinksAssociatedOneOf {
	return v.value
}

func (v *NullableErrorLinksAssociatedOneOf) Set(val *ErrorLinksAssociatedOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorLinksAssociatedOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorLinksAssociatedOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorLinksAssociatedOneOf(val *ErrorLinksAssociatedOneOf) *NullableErrorLinksAssociatedOneOf {
	return &NullableErrorLinksAssociatedOneOf{value: val, isSet: true}
}

func (v NullableErrorLinksAssociatedOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorLinksAssociatedOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


