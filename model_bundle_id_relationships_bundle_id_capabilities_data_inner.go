/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BundleIdRelationshipsBundleIdCapabilitiesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdRelationshipsBundleIdCapabilitiesDataInner{}

// BundleIdRelationshipsBundleIdCapabilitiesDataInner struct for BundleIdRelationshipsBundleIdCapabilitiesDataInner
type BundleIdRelationshipsBundleIdCapabilitiesDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

type _BundleIdRelationshipsBundleIdCapabilitiesDataInner BundleIdRelationshipsBundleIdCapabilitiesDataInner

// NewBundleIdRelationshipsBundleIdCapabilitiesDataInner instantiates a new BundleIdRelationshipsBundleIdCapabilitiesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdRelationshipsBundleIdCapabilitiesDataInner(type_ string, id string) *BundleIdRelationshipsBundleIdCapabilitiesDataInner {
	this := BundleIdRelationshipsBundleIdCapabilitiesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBundleIdRelationshipsBundleIdCapabilitiesDataInnerWithDefaults instantiates a new BundleIdRelationshipsBundleIdCapabilitiesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdRelationshipsBundleIdCapabilitiesDataInnerWithDefaults() *BundleIdRelationshipsBundleIdCapabilitiesDataInner {
	this := BundleIdRelationshipsBundleIdCapabilitiesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *BundleIdRelationshipsBundleIdCapabilitiesDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BundleIdRelationshipsBundleIdCapabilitiesDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BundleIdRelationshipsBundleIdCapabilitiesDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BundleIdRelationshipsBundleIdCapabilitiesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BundleIdRelationshipsBundleIdCapabilitiesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BundleIdRelationshipsBundleIdCapabilitiesDataInner) SetId(v string) {
	o.Id = v
}

func (o BundleIdRelationshipsBundleIdCapabilitiesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdRelationshipsBundleIdCapabilitiesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *BundleIdRelationshipsBundleIdCapabilitiesDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBundleIdRelationshipsBundleIdCapabilitiesDataInner := _BundleIdRelationshipsBundleIdCapabilitiesDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varBundleIdRelationshipsBundleIdCapabilitiesDataInner)

	if err != nil {
		return err
	}

	*o = BundleIdRelationshipsBundleIdCapabilitiesDataInner(varBundleIdRelationshipsBundleIdCapabilitiesDataInner)

	return err
}

type NullableBundleIdRelationshipsBundleIdCapabilitiesDataInner struct {
	value *BundleIdRelationshipsBundleIdCapabilitiesDataInner
	isSet bool
}

func (v NullableBundleIdRelationshipsBundleIdCapabilitiesDataInner) Get() *BundleIdRelationshipsBundleIdCapabilitiesDataInner {
	return v.value
}

func (v *NullableBundleIdRelationshipsBundleIdCapabilitiesDataInner) Set(val *BundleIdRelationshipsBundleIdCapabilitiesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdRelationshipsBundleIdCapabilitiesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdRelationshipsBundleIdCapabilitiesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdRelationshipsBundleIdCapabilitiesDataInner(val *BundleIdRelationshipsBundleIdCapabilitiesDataInner) *NullableBundleIdRelationshipsBundleIdCapabilitiesDataInner {
	return &NullableBundleIdRelationshipsBundleIdCapabilitiesDataInner{value: val, isSet: true}
}

func (v NullableBundleIdRelationshipsBundleIdCapabilitiesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdRelationshipsBundleIdCapabilitiesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


