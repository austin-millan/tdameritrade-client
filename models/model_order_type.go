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

// OrderType the model 'OrderType'
type OrderType string

// List of OrderType
const (
	ORDERTYPE_MARKET OrderType = "MARKET"
	ORDERTYPE_LIMIT OrderType = "LIMIT"
	ORDERTYPE_STOP OrderType = "STOP"
	ORDERTYPE_STOP_LIMIT OrderType = "STOP_LIMIT"
	ORDERTYPE_TRAILING_STOP OrderType = "TRAILING_STOP"
	ORDERTYPE_MARKET_ON_CLOSE OrderType = "MARKET_ON_CLOSE"
	ORDERTYPE_EXERCISE OrderType = "EXERCISE"
	ORDERTYPE_TRAILING_STOP_LIMIT OrderType = "TRAILING_STOP_LIMIT"
	ORDERTYPE_NET_DEBIT OrderType = "NET_DEBIT"
	ORDERTYPE_NET_CREDIT OrderType = "NET_CREDIT"
	ORDERTYPE_NET_ZERO OrderType = "NET_ZERO"
)

func (v *OrderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderType(value)
	for _, existing := range []OrderType{ "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP", "MARKET_ON_CLOSE", "EXERCISE", "TRAILING_STOP_LIMIT", "NET_DEBIT", "NET_CREDIT", "NET_ZERO",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderType", value)
}

// Ptr returns reference to OrderType value
func (v OrderType) Ptr() *OrderType {
	return &v
}

type NullableOrderType struct {
	value *OrderType
	isSet bool
}

func (v NullableOrderType) Get() *OrderType {
	return v.value
}

func (v *NullableOrderType) Set(val *OrderType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderType(val *OrderType) *NullableOrderType {
	return &NullableOrderType{value: val, isSet: true}
}

func (v NullableOrderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
