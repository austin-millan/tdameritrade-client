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

// ExpirationDate struct for ExpirationDate
type ExpirationDate struct {
	Date *string `json:"date,omitempty"`
}

// NewExpirationDate instantiates a new ExpirationDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpirationDate() *ExpirationDate {
	this := ExpirationDate{}
	return &this
}

// NewExpirationDateWithDefaults instantiates a new ExpirationDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpirationDateWithDefaults() *ExpirationDate {
	this := ExpirationDate{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ExpirationDate) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpirationDate) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ExpirationDate) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *ExpirationDate) SetDate(v string) {
	o.Date = &v
}

func (o ExpirationDate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

type NullableExpirationDate struct {
	value *ExpirationDate
	isSet bool
}

func (v NullableExpirationDate) Get() *ExpirationDate {
	return v.value
}

func (v *NullableExpirationDate) Set(val *ExpirationDate) {
	v.value = val
	v.isSet = true
}

func (v NullableExpirationDate) IsSet() bool {
	return v.isSet
}

func (v *NullableExpirationDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpirationDate(val *ExpirationDate) *NullableExpirationDate {
	return &NullableExpirationDate{value: val, isSet: true}
}

func (v NullableExpirationDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpirationDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

