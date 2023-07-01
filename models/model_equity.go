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

// Equity struct for Equity
type Equity struct {
	AssetType *AssetType `json:"assetType,omitempty"`
	Cusip *string `json:"cusip,omitempty"`
	Description *string `json:"description,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewEquity instantiates a new Equity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquity() *Equity {
	this := Equity{}
	return &this
}

// NewEquityWithDefaults instantiates a new Equity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityWithDefaults() *Equity {
	this := Equity{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *Equity) GetAssetType() AssetType {
	if o == nil || o.AssetType == nil {
		var ret AssetType
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Equity) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil || o.AssetType == nil {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *Equity) HasAssetType() bool {
	if o != nil && o.AssetType != nil {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given AssetType and assigns it to the AssetType field.
func (o *Equity) SetAssetType(v AssetType) {
	o.AssetType = &v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *Equity) GetCusip() string {
	if o == nil || o.Cusip == nil {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Equity) GetCusipOk() (*string, bool) {
	if o == nil || o.Cusip == nil {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *Equity) HasCusip() bool {
	if o != nil && o.Cusip != nil {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *Equity) SetCusip(v string) {
	o.Cusip = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Equity) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Equity) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Equity) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Equity) SetDescription(v string) {
	o.Description = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Equity) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Equity) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Equity) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Equity) SetSymbol(v string) {
	o.Symbol = &v
}

func (o Equity) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableEquity struct {
	value *Equity
	isSet bool
}

func (v NullableEquity) Get() *Equity {
	return v.value
}

func (v *NullableEquity) Set(val *Equity) {
	v.value = val
	v.isSet = true
}

func (v NullableEquity) IsSet() bool {
	return v.isSet
}

func (v *NullableEquity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquity(val *Equity) *NullableEquity {
	return &NullableEquity{value: val, isSet: true}
}

func (v NullableEquity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

