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
	"fmt"
)

// Duration the model 'Duration'
type Duration string

// List of Duration
const (
	DURATION_DAY Duration = "DAY"
	DURATION_GOOD_TILL_CANCEL Duration = "GOOD_TILL_CANCEL"
	DURATION_FILL_OR_KILL Duration = "FILL_OR_KILL"
)

func (v *Duration) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Duration(value)
	for _, existing := range []Duration{ "DAY", "GOOD_TILL_CANCEL", "FILL_OR_KILL",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Duration", value)
}

// Ptr returns reference to Duration value
func (v Duration) Ptr() *Duration {
	return &v
}

type NullableDuration struct {
	value *Duration
	isSet bool
}

func (v NullableDuration) Get() *Duration {
	return v.value
}

func (v *NullableDuration) Set(val *Duration) {
	v.value = val
	v.isSet = true
}

func (v NullableDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuration(val *Duration) *NullableDuration {
	return &NullableDuration{value: val, isSet: true}
}

func (v NullableDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
