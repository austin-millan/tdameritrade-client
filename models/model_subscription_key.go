/*
 * TD Ameritrade API
 *
 * TD Ameritrade API
 *
 * API version: 3.0.1
 * Contact: austin.millan@protonmail.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SubscriptionKey struct for SubscriptionKey
type SubscriptionKey struct {
	Keys *[]SubscriptionKeyKeys `json:"keys,omitempty"`
}

// NewSubscriptionKey instantiates a new SubscriptionKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionKey() *SubscriptionKey {
	this := SubscriptionKey{}
	return &this
}

// NewSubscriptionKeyWithDefaults instantiates a new SubscriptionKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionKeyWithDefaults() *SubscriptionKey {
	this := SubscriptionKey{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *SubscriptionKey) GetKeys() []SubscriptionKeyKeys {
	if o == nil || o.Keys == nil {
		var ret []SubscriptionKeyKeys
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionKey) GetKeysOk() (*[]SubscriptionKeyKeys, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *SubscriptionKey) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []SubscriptionKeyKeys and assigns it to the Keys field.
func (o *SubscriptionKey) SetKeys(v []SubscriptionKeyKeys) {
	o.Keys = &v
}

func (o SubscriptionKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionKey struct {
	value *SubscriptionKey
	isSet bool
}

func (v NullableSubscriptionKey) Get() *SubscriptionKey {
	return v.value
}

func (v *NullableSubscriptionKey) Set(val *SubscriptionKey) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionKey) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionKey(val *SubscriptionKey) *NullableSubscriptionKey {
	return &NullableSubscriptionKey{value: val, isSet: true}
}

func (v NullableSubscriptionKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


