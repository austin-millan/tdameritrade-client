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

// FixedIncome struct for FixedIncome
type FixedIncome struct {
	AssetType *AssetType `json:"assetType,omitempty"`
	Cusip *string `json:"cusip,omitempty"`
	Description *string `json:"description,omitempty"`
	Factor *float32 `json:"factor,omitempty"`
	MaturityDate *string `json:"maturityDate,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	VariableRate *float32 `json:"variableRate,omitempty"`
}

// NewFixedIncome instantiates a new FixedIncome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedIncome() *FixedIncome {
	this := FixedIncome{}
	return &this
}

// NewFixedIncomeWithDefaults instantiates a new FixedIncome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedIncomeWithDefaults() *FixedIncome {
	this := FixedIncome{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *FixedIncome) GetAssetType() AssetType {
	if o == nil || o.AssetType == nil {
		var ret AssetType
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedIncome) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil || o.AssetType == nil {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *FixedIncome) HasAssetType() bool {
	if o != nil && o.AssetType != nil {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given AssetType and assigns it to the AssetType field.
func (o *FixedIncome) SetAssetType(v AssetType) {
	o.AssetType = &v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *FixedIncome) GetCusip() string {
	if o == nil || o.Cusip == nil {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedIncome) GetCusipOk() (*string, bool) {
	if o == nil || o.Cusip == nil {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *FixedIncome) HasCusip() bool {
	if o != nil && o.Cusip != nil {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *FixedIncome) SetCusip(v string) {
	o.Cusip = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FixedIncome) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedIncome) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FixedIncome) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FixedIncome) SetDescription(v string) {
	o.Description = &v
}

// GetFactor returns the Factor field value if set, zero value otherwise.
func (o *FixedIncome) GetFactor() float32 {
	if o == nil || o.Factor == nil {
		var ret float32
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedIncome) GetFactorOk() (*float32, bool) {
	if o == nil || o.Factor == nil {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *FixedIncome) HasFactor() bool {
	if o != nil && o.Factor != nil {
		return true
	}

	return false
}

// SetFactor gets a reference to the given float32 and assigns it to the Factor field.
func (o *FixedIncome) SetFactor(v float32) {
	o.Factor = &v
}

// GetMaturityDate returns the MaturityDate field value if set, zero value otherwise.
func (o *FixedIncome) GetMaturityDate() string {
	if o == nil || o.MaturityDate == nil {
		var ret string
		return ret
	}
	return *o.MaturityDate
}

// GetMaturityDateOk returns a tuple with the MaturityDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedIncome) GetMaturityDateOk() (*string, bool) {
	if o == nil || o.MaturityDate == nil {
		return nil, false
	}
	return o.MaturityDate, true
}

// HasMaturityDate returns a boolean if a field has been set.
func (o *FixedIncome) HasMaturityDate() bool {
	if o != nil && o.MaturityDate != nil {
		return true
	}

	return false
}

// SetMaturityDate gets a reference to the given string and assigns it to the MaturityDate field.
func (o *FixedIncome) SetMaturityDate(v string) {
	o.MaturityDate = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *FixedIncome) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedIncome) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *FixedIncome) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *FixedIncome) SetSymbol(v string) {
	o.Symbol = &v
}

// GetVariableRate returns the VariableRate field value if set, zero value otherwise.
func (o *FixedIncome) GetVariableRate() float32 {
	if o == nil || o.VariableRate == nil {
		var ret float32
		return ret
	}
	return *o.VariableRate
}

// GetVariableRateOk returns a tuple with the VariableRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedIncome) GetVariableRateOk() (*float32, bool) {
	if o == nil || o.VariableRate == nil {
		return nil, false
	}
	return o.VariableRate, true
}

// HasVariableRate returns a boolean if a field has been set.
func (o *FixedIncome) HasVariableRate() bool {
	if o != nil && o.VariableRate != nil {
		return true
	}

	return false
}

// SetVariableRate gets a reference to the given float32 and assigns it to the VariableRate field.
func (o *FixedIncome) SetVariableRate(v float32) {
	o.VariableRate = &v
}

func (o FixedIncome) MarshalJSON() ([]byte, error) {
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
	if o.Factor != nil {
		toSerialize["factor"] = o.Factor
	}
	if o.MaturityDate != nil {
		toSerialize["maturityDate"] = o.MaturityDate
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.VariableRate != nil {
		toSerialize["variableRate"] = o.VariableRate
	}
	return json.Marshal(toSerialize)
}

type NullableFixedIncome struct {
	value *FixedIncome
	isSet bool
}

func (v NullableFixedIncome) Get() *FixedIncome {
	return v.value
}

func (v *NullableFixedIncome) Set(val *FixedIncome) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedIncome(val *FixedIncome) *NullableFixedIncome {
	return &NullableFixedIncome{value: val, isSet: true}
}

func (v NullableFixedIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


