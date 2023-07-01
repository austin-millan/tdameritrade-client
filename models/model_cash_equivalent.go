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

// CashEquivalent struct for CashEquivalent
type CashEquivalent struct {
	AssetType *AssetType `json:"assetType,omitempty"`
	Cusip *string `json:"cusip,omitempty"`
	Description *string `json:"description,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewCashEquivalent instantiates a new CashEquivalent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashEquivalent() *CashEquivalent {
	this := CashEquivalent{}
	return &this
}

// NewCashEquivalentWithDefaults instantiates a new CashEquivalent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashEquivalentWithDefaults() *CashEquivalent {
	this := CashEquivalent{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *CashEquivalent) GetAssetType() AssetType {
	if o == nil || o.AssetType == nil {
		var ret AssetType
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashEquivalent) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil || o.AssetType == nil {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *CashEquivalent) HasAssetType() bool {
	if o != nil && o.AssetType != nil {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given AssetType and assigns it to the AssetType field.
func (o *CashEquivalent) SetAssetType(v AssetType) {
	o.AssetType = &v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *CashEquivalent) GetCusip() string {
	if o == nil || o.Cusip == nil {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashEquivalent) GetCusipOk() (*string, bool) {
	if o == nil || o.Cusip == nil {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *CashEquivalent) HasCusip() bool {
	if o != nil && o.Cusip != nil {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *CashEquivalent) SetCusip(v string) {
	o.Cusip = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CashEquivalent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashEquivalent) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CashEquivalent) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CashEquivalent) SetDescription(v string) {
	o.Description = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CashEquivalent) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashEquivalent) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CashEquivalent) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CashEquivalent) SetSymbol(v string) {
	o.Symbol = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CashEquivalent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashEquivalent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CashEquivalent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CashEquivalent) SetType(v string) {
	o.Type = &v
}

func (o CashEquivalent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetType != nil {
		toSerialize["assetType"] = o.AssetType
	}
	if o.Cusip != nil {
		toSerialize["cusip"] = o.Cusip
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCashEquivalent struct {
	value *CashEquivalent
	isSet bool
}

func (v NullableCashEquivalent) Get() *CashEquivalent {
	return v.value
}

func (v *NullableCashEquivalent) Set(val *CashEquivalent) {
	v.value = val
	v.isSet = true
}

func (v NullableCashEquivalent) IsSet() bool {
	return v.isSet
}

func (v *NullableCashEquivalent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashEquivalent(val *CashEquivalent) *NullableCashEquivalent {
	return &NullableCashEquivalent{value: val, isSet: true}
}

func (v NullableCashEquivalent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashEquivalent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


