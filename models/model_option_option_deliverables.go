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

// OptionOptionDeliverables struct for OptionOptionDeliverables
type OptionOptionDeliverables struct {
	AssetType *AssetType `json:"assetType,omitempty"`
	CurrencyType *CurrencyType `json:"currencyType,omitempty"`
	DeliverableUnits *float32 `json:"deliverableUnits,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewOptionOptionDeliverables instantiates a new OptionOptionDeliverables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionOptionDeliverables() *OptionOptionDeliverables {
	this := OptionOptionDeliverables{}
	return &this
}

// NewOptionOptionDeliverablesWithDefaults instantiates a new OptionOptionDeliverables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionOptionDeliverablesWithDefaults() *OptionOptionDeliverables {
	this := OptionOptionDeliverables{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *OptionOptionDeliverables) GetAssetType() AssetType {
	if o == nil || o.AssetType == nil {
		var ret AssetType
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOptionDeliverables) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil || o.AssetType == nil {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *OptionOptionDeliverables) HasAssetType() bool {
	if o != nil && o.AssetType != nil {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given AssetType and assigns it to the AssetType field.
func (o *OptionOptionDeliverables) SetAssetType(v AssetType) {
	o.AssetType = &v
}

// GetCurrencyType returns the CurrencyType field value if set, zero value otherwise.
func (o *OptionOptionDeliverables) GetCurrencyType() CurrencyType {
	if o == nil || o.CurrencyType == nil {
		var ret CurrencyType
		return ret
	}
	return *o.CurrencyType
}

// GetCurrencyTypeOk returns a tuple with the CurrencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOptionDeliverables) GetCurrencyTypeOk() (*CurrencyType, bool) {
	if o == nil || o.CurrencyType == nil {
		return nil, false
	}
	return o.CurrencyType, true
}

// HasCurrencyType returns a boolean if a field has been set.
func (o *OptionOptionDeliverables) HasCurrencyType() bool {
	if o != nil && o.CurrencyType != nil {
		return true
	}

	return false
}

// SetCurrencyType gets a reference to the given CurrencyType and assigns it to the CurrencyType field.
func (o *OptionOptionDeliverables) SetCurrencyType(v CurrencyType) {
	o.CurrencyType = &v
}

// GetDeliverableUnits returns the DeliverableUnits field value if set, zero value otherwise.
func (o *OptionOptionDeliverables) GetDeliverableUnits() float32 {
	if o == nil || o.DeliverableUnits == nil {
		var ret float32
		return ret
	}
	return *o.DeliverableUnits
}

// GetDeliverableUnitsOk returns a tuple with the DeliverableUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOptionDeliverables) GetDeliverableUnitsOk() (*float32, bool) {
	if o == nil || o.DeliverableUnits == nil {
		return nil, false
	}
	return o.DeliverableUnits, true
}

// HasDeliverableUnits returns a boolean if a field has been set.
func (o *OptionOptionDeliverables) HasDeliverableUnits() bool {
	if o != nil && o.DeliverableUnits != nil {
		return true
	}

	return false
}

// SetDeliverableUnits gets a reference to the given float32 and assigns it to the DeliverableUnits field.
func (o *OptionOptionDeliverables) SetDeliverableUnits(v float32) {
	o.DeliverableUnits = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OptionOptionDeliverables) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOptionDeliverables) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OptionOptionDeliverables) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OptionOptionDeliverables) SetSymbol(v string) {
	o.Symbol = &v
}

func (o OptionOptionDeliverables) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetType != nil {
		toSerialize["assetType"] = o.AssetType
	}
	if o.CurrencyType != nil {
		toSerialize["currencyType"] = o.CurrencyType
	}
	if o.DeliverableUnits != nil {
		toSerialize["deliverableUnits"] = o.DeliverableUnits
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableOptionOptionDeliverables struct {
	value *OptionOptionDeliverables
	isSet bool
}

func (v NullableOptionOptionDeliverables) Get() *OptionOptionDeliverables {
	return v.value
}

func (v *NullableOptionOptionDeliverables) Set(val *OptionOptionDeliverables) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionOptionDeliverables) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionOptionDeliverables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionOptionDeliverables(val *OptionOptionDeliverables) *NullableOptionOptionDeliverables {
	return &NullableOptionOptionDeliverables{value: val, isSet: true}
}

func (v NullableOptionOptionDeliverables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionOptionDeliverables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

