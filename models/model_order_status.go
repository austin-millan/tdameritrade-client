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

// OrderStatus the model 'OrderStatus'
type OrderStatus string

// List of OrderStatus
const (
	ORDERSTATUS_AWAITING_PARENT_ORDER OrderStatus = "AWAITING_PARENT_ORDER"
	ORDERSTATUS_AWAITING_CONDITION OrderStatus = "AWAITING_CONDITION"
	ORDERSTATUS_AWAITING_MANUAL_REVIEW OrderStatus = "AWAITING_MANUAL_REVIEW"
	ORDERSTATUS_ACCEPTED OrderStatus = "ACCEPTED"
	ORDERSTATUS_AWAITING_UR_OUT OrderStatus = "AWAITING_UR_OUT"
	ORDERSTATUS_PENDING_ACTIVATION OrderStatus = "PENDING_ACTIVATION"
	ORDERSTATUS_WORKING OrderStatus = "WORKING"
	ORDERSTATUS_REJECTED OrderStatus = "REJECTED"
	ORDERSTATUS_QUEUED OrderStatus = "QUEUED"
	ORDERSTATUS_PENDING_CANCEL OrderStatus = "PENDING_CANCEL"
	ORDERSTATUS_CANCELED OrderStatus = "CANCELED"
	ORDERSTATUS_PENDING_REPLACE OrderStatus = "PENDING_REPLACE"
	ORDERSTATUS_REPLACED OrderStatus = "REPLACED"
	ORDERSTATUS_FILLED OrderStatus = "FILLED"
	ORDERSTATUS_EXPIRED OrderStatus = "EXPIRED"
)

func (v *OrderStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderStatus(value)
	for _, existing := range []OrderStatus{ "AWAITING_PARENT_ORDER", "AWAITING_CONDITION", "AWAITING_MANUAL_REVIEW", "ACCEPTED", "AWAITING_UR_OUT", "PENDING_ACTIVATION", "WORKING", "REJECTED", "QUEUED", "PENDING_CANCEL", "CANCELED", "PENDING_REPLACE", "REPLACED", "FILLED", "EXPIRED",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderStatus", value)
}

// Ptr returns reference to OrderStatus value
func (v OrderStatus) Ptr() *OrderStatus {
	return &v
}

type NullableOrderStatus struct {
	value *OrderStatus
	isSet bool
}

func (v NullableOrderStatus) Get() *OrderStatus {
	return v.value
}

func (v *NullableOrderStatus) Set(val *OrderStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatus(val *OrderStatus) *NullableOrderStatus {
	return &NullableOrderStatus{value: val, isSet: true}
}

func (v NullableOrderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
