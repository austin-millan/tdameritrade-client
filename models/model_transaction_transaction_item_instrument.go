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

// TransactionTransactionItemInstrument struct for TransactionTransactionItemInstrument
type TransactionTransactionItemInstrument struct {
	AssetType *AssetType `json:"assetType,omitempty"`
	BondInterestRate *float32 `json:"bondInterestRate,omitempty"`
	BondMaturityDate *string `json:"bondMaturityDate,omitempty"`
	Cusip *string `json:"cusip,omitempty"`
	Description *string `json:"description,omitempty"`
	OptionExpirationDate *string `json:"optionExpirationDate,omitempty"`
	OptionStrikePrice *float32 `json:"optionStrikePrice,omitempty"`
	PutCall *PutOrCall `json:"putCall,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	UnderlyingSymbol *string `json:"underlyingSymbol,omitempty"`
}

// NewTransactionTransactionItemInstrument instantiates a new TransactionTransactionItemInstrument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionTransactionItemInstrument() *TransactionTransactionItemInstrument {
	this := TransactionTransactionItemInstrument{}
	return &this
}

// NewTransactionTransactionItemInstrumentWithDefaults instantiates a new TransactionTransactionItemInstrument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionTransactionItemInstrumentWithDefaults() *TransactionTransactionItemInstrument {
	this := TransactionTransactionItemInstrument{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *TransactionTransactionItemInstrument) GetAssetType() AssetType {
	if o == nil || o.AssetType == nil {
		var ret AssetType
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransactionItemInstrument) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil || o.AssetType == nil {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *TransactionTransactionItemInstrument) HasAssetType() bool {
	if o != nil && o.AssetType != nil {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given AssetType and assigns it to the AssetType field.
func (o *TransactionTransactionItemInstrument) SetAssetType(v AssetType) {
	o.AssetType = &v
}

// GetBondInterestRate returns the BondInterestRate field value if set, zero value otherwise.
func (o *TransactionTransactionItemInstrument) GetBondInterestRate() float32 {
	if o == nil || o.BondInterestRate == nil {
		var ret float32
		return ret
	}
	return *o.BondInterestRate
}

// GetBondInterestRateOk returns a tuple with the BondInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransactionItemInstrument) GetBondInterestRateOk() (*float32, bool) {
	if o == nil || o.BondInterestRate == nil {
		return nil, false
	}
	return o.BondInterestRate, true
}

// HasBondInterestRate returns a boolean if a field has been set.
func (o *TransactionTransactionItemInstrument) HasBondInterestRate() bool {
	if o != nil && o.BondInterestRate != nil {
		return true
	}

	return false
}

// SetBondInterestRate gets a reference to the given float32 and assigns it to the BondInterestRate field.
func (o *TransactionTransactionItemInstrument) SetBondInterestRate(v float32) {
	o.BondInterestRate = &v
}

// GetBondMaturityDate returns the BondMaturityDate field value if set, zero value otherwise.
func (o *TransactionTransactionItemInstrument) GetBondMaturityDate() string {
	if o == nil || o.BondMaturityDate == nil {
		var ret string
		return ret
	}
	return *o.BondMaturityDate
}

// GetBondMaturityDateOk returns a tuple with the BondMaturityDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransactionItemInstrument) GetBondMaturityDateOk() (*string, bool) {
	if o == nil || o.BondMaturityDate == nil {
		return nil, false
	}
	return o.BondMaturityDate, true
}

// HasBondMaturityDate returns a boolean if a field has been set.
func (o *TransactionTransactionItemInstrument) HasBondMaturityDate() bool {
	if o != nil && o.BondMaturityDate != nil {
		return true
	}

	return false
}

// SetBondMaturityDate gets a reference to the given string and assigns it to the BondMaturityDate field.
func (o *TransactionTransactionItemInstrument) SetBondMaturityDate(v string) {
	o.BondMaturityDate = &v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *TransactionTransactionItemInstrument) GetCusip() string {
	if o == nil || o.Cusip == nil {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransactionItemInstrument) GetCusipOk() (*string, bool) {
	if o == nil || o.Cusip == nil {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *TransactionTransactionItemInstrument) HasCusip() bool {
	if o != nil && o.Cusip != nil {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *TransactionTransactionItemInstrument) SetCusip(v string) {
	o.Cusip = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransactionTransactionItemInstrument) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransactionItemInstrument) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransactionTransactionItemInstrument) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransactionTransactionItemInstrument) SetDescription(v string) {
	o.Description = &v
}

// GetOptionExpirationDate returns the OptionExpirationDate field value if set, zero value otherwise.
func (o *TransactionTransactionItemInstrument) GetOptionExpirationDate() string {
	if o == nil || o.OptionExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.OptionExpirationDate
}

// GetOptionExpirationDateOk returns a tuple with the OptionExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransactionItemInstrument) GetOptionExpirationDateOk() (*string, bool) {
	if o == nil || o.OptionExpirationDate == nil {
		return nil, false
	}
	return o.OptionExpirationDate, true
}

// HasOptionExpirationDate returns a boolean if a field has been set.
func (o *TransactionTransactionItemInstrument) HasOptionExpirationDate() bool {
	if o != nil && o.OptionExpirationDate != nil {
		return true
	}

	return false
}

// SetOptionExpirationDate gets a reference to the given string and assigns it to the OptionExpirationDate field.
func (o *TransactionTransactionItemInstrument) SetOptionExpirationDate(v string) {
	o.OptionExpirationDate = &v
}

// GetOptionStrikePrice returns the OptionStrikePrice field value if set, zero value otherwise.
func (o *TransactionTransactionItemInstrument) GetOptionStrikePrice() float32 {
	if o == nil || o.OptionStrikePrice == nil {
		var ret float32
		return ret
	}
	return *o.OptionStrikePrice
}

// GetOptionStrikePriceOk returns a tuple with the OptionStrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransactionItemInstrument) GetOptionStrikePriceOk() (*float32, bool) {
	if o == nil || o.OptionStrikePrice == nil {
		return nil, false
	}
	return o.OptionStrikePrice, true
}

// HasOptionStrikePrice returns a boolean if a field has been set.
func (o *TransactionTransactionItemInstrument) HasOptionStrikePrice() bool {
	if o != nil && o.OptionStrikePrice != nil {
		return true
	}

	return false
}

// SetOptionStrikePrice gets a reference to the given float32 and assigns it to the OptionStrikePrice field.
func (o *TransactionTransactionItemInstrument) SetOptionStrikePrice(v float32) {
	o.OptionStrikePrice = &v
}

// GetPutCall returns the PutCall field value if set, zero value otherwise.
func (o *TransactionTransactionItemInstrument) GetPutCall() PutOrCall {
	if o == nil || o.PutCall == nil {
		var ret PutOrCall
		return ret
	}
	return *o.PutCall
}

// GetPutCallOk returns a tuple with the PutCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransactionItemInstrument) GetPutCallOk() (*PutOrCall, bool) {
	if o == nil || o.PutCall == nil {
		return nil, false
	}
	return o.PutCall, true
}

// HasPutCall returns a boolean if a field has been set.
func (o *TransactionTransactionItemInstrument) HasPutCall() bool {
	if o != nil && o.PutCall != nil {
		return true
	}

	return false
}

// SetPutCall gets a reference to the given PutOrCall and assigns it to the PutCall field.
func (o *TransactionTransactionItemInstrument) SetPutCall(v PutOrCall) {
	o.PutCall = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TransactionTransactionItemInstrument) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransactionItemInstrument) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TransactionTransactionItemInstrument) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TransactionTransactionItemInstrument) SetSymbol(v string) {
	o.Symbol = &v
}

// GetUnderlyingSymbol returns the UnderlyingSymbol field value if set, zero value otherwise.
func (o *TransactionTransactionItemInstrument) GetUnderlyingSymbol() string {
	if o == nil || o.UnderlyingSymbol == nil {
		var ret string
		return ret
	}
	return *o.UnderlyingSymbol
}

// GetUnderlyingSymbolOk returns a tuple with the UnderlyingSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransactionItemInstrument) GetUnderlyingSymbolOk() (*string, bool) {
	if o == nil || o.UnderlyingSymbol == nil {
		return nil, false
	}
	return o.UnderlyingSymbol, true
}

// HasUnderlyingSymbol returns a boolean if a field has been set.
func (o *TransactionTransactionItemInstrument) HasUnderlyingSymbol() bool {
	if o != nil && o.UnderlyingSymbol != nil {
		return true
	}

	return false
}

// SetUnderlyingSymbol gets a reference to the given string and assigns it to the UnderlyingSymbol field.
func (o *TransactionTransactionItemInstrument) SetUnderlyingSymbol(v string) {
	o.UnderlyingSymbol = &v
}

func (o TransactionTransactionItemInstrument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetType != nil {
		toSerialize["assetType"] = o.AssetType
	}
	if o.BondInterestRate != nil {
		toSerialize["bondInterestRate"] = o.BondInterestRate
	}
	if o.BondMaturityDate != nil {
		toSerialize["bondMaturityDate"] = o.BondMaturityDate
	}
	if o.Cusip != nil {
		toSerialize["cusip"] = o.Cusip
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.OptionExpirationDate != nil {
		toSerialize["optionExpirationDate"] = o.OptionExpirationDate
	}
	if o.OptionStrikePrice != nil {
		toSerialize["optionStrikePrice"] = o.OptionStrikePrice
	}
	if o.PutCall != nil {
		toSerialize["putCall"] = o.PutCall
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.UnderlyingSymbol != nil {
		toSerialize["underlyingSymbol"] = o.UnderlyingSymbol
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionTransactionItemInstrument struct {
	value *TransactionTransactionItemInstrument
	isSet bool
}

func (v NullableTransactionTransactionItemInstrument) Get() *TransactionTransactionItemInstrument {
	return v.value
}

func (v *NullableTransactionTransactionItemInstrument) Set(val *TransactionTransactionItemInstrument) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionTransactionItemInstrument) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionTransactionItemInstrument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionTransactionItemInstrument(val *TransactionTransactionItemInstrument) *NullableTransactionTransactionItemInstrument {
	return &NullableTransactionTransactionItemInstrument{value: val, isSet: true}
}

func (v NullableTransactionTransactionItemInstrument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionTransactionItemInstrument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

