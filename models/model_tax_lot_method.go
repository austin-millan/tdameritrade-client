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

// TaxLotMethod the model 'TaxLotMethod'
type TaxLotMethod string

// List of TaxLotMethod
const (
	TAXLOTMETHOD_FIFO TaxLotMethod = "FIFO"
	TAXLOTMETHOD_LIFO TaxLotMethod = "LIFO"
	TAXLOTMETHOD_HIGH_COST TaxLotMethod = "HIGH_COST"
	TAXLOTMETHOD_LOW_COST TaxLotMethod = "LOW_COST"
	TAXLOTMETHOD_AVERAGE_COST TaxLotMethod = "AVERAGE_COST"
	TAXLOTMETHOD_SPECIFIC_LOT TaxLotMethod = "SPECIFIC_LOT"
)

func (v *TaxLotMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaxLotMethod(value)
	for _, existing := range []TaxLotMethod{ "FIFO", "LIFO", "HIGH_COST", "LOW_COST", "AVERAGE_COST", "SPECIFIC_LOT",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaxLotMethod", value)
}

// Ptr returns reference to TaxLotMethod value
func (v TaxLotMethod) Ptr() *TaxLotMethod {
	return &v
}

type NullableTaxLotMethod struct {
	value *TaxLotMethod
	isSet bool
}

func (v NullableTaxLotMethod) Get() *TaxLotMethod {
	return v.value
}

func (v *NullableTaxLotMethod) Set(val *TaxLotMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxLotMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxLotMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxLotMethod(val *TaxLotMethod) *NullableTaxLotMethod {
	return &NullableTaxLotMethod{value: val, isSet: true}
}

func (v NullableTaxLotMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxLotMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

