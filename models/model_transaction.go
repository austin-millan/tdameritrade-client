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

// Transaction struct for Transaction
type Transaction struct {
	AccruedInterest *float32 `json:"accruedInterest,omitempty"`
	AchStatus *string `json:"achStatus,omitempty"`
	CashBalanceEffectFlag *bool `json:"cashBalanceEffectFlag,omitempty"`
	ClearingReferenceNumber *string `json:"clearingReferenceNumber,omitempty"`
	DayTradeBuyingPowerEffect *float32 `json:"dayTradeBuyingPowerEffect,omitempty"`
	Description *string `json:"description,omitempty"`
	Fees *string `json:"fees,omitempty"`
	NetAmount *float32 `json:"netAmount,omitempty"`
	OrderDate *string `json:"orderDate,omitempty"`
	OrderId *string `json:"orderId,omitempty"`
	RequirementReallocationAmount *float32 `json:"requirementReallocationAmount,omitempty"`
	SettlementDate *string `json:"settlementDate,omitempty"`
	Sma *float32 `json:"sma,omitempty"`
	SubAccount *string `json:"subAccount,omitempty"`
	TransactionDate *string `json:"transactionDate,omitempty"`
	TransactionId *float32 `json:"transactionId,omitempty"`
	TransactionItem *TransactionTransactionItem `json:"transactionItem,omitempty"`
	TransactionSubType *string `json:"transactionSubType,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction() *Transaction {
	this := Transaction{}
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetAccruedInterest returns the AccruedInterest field value if set, zero value otherwise.
func (o *Transaction) GetAccruedInterest() float32 {
	if o == nil || o.AccruedInterest == nil {
		var ret float32
		return ret
	}
	return *o.AccruedInterest
}

// GetAccruedInterestOk returns a tuple with the AccruedInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetAccruedInterestOk() (*float32, bool) {
	if o == nil || o.AccruedInterest == nil {
		return nil, false
	}
	return o.AccruedInterest, true
}

// HasAccruedInterest returns a boolean if a field has been set.
func (o *Transaction) HasAccruedInterest() bool {
	if o != nil && o.AccruedInterest != nil {
		return true
	}

	return false
}

// SetAccruedInterest gets a reference to the given float32 and assigns it to the AccruedInterest field.
func (o *Transaction) SetAccruedInterest(v float32) {
	o.AccruedInterest = &v
}

// GetAchStatus returns the AchStatus field value if set, zero value otherwise.
func (o *Transaction) GetAchStatus() string {
	if o == nil || o.AchStatus == nil {
		var ret string
		return ret
	}
	return *o.AchStatus
}

// GetAchStatusOk returns a tuple with the AchStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetAchStatusOk() (*string, bool) {
	if o == nil || o.AchStatus == nil {
		return nil, false
	}
	return o.AchStatus, true
}

// HasAchStatus returns a boolean if a field has been set.
func (o *Transaction) HasAchStatus() bool {
	if o != nil && o.AchStatus != nil {
		return true
	}

	return false
}

// SetAchStatus gets a reference to the given string and assigns it to the AchStatus field.
func (o *Transaction) SetAchStatus(v string) {
	o.AchStatus = &v
}

// GetCashBalanceEffectFlag returns the CashBalanceEffectFlag field value if set, zero value otherwise.
func (o *Transaction) GetCashBalanceEffectFlag() bool {
	if o == nil || o.CashBalanceEffectFlag == nil {
		var ret bool
		return ret
	}
	return *o.CashBalanceEffectFlag
}

// GetCashBalanceEffectFlagOk returns a tuple with the CashBalanceEffectFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCashBalanceEffectFlagOk() (*bool, bool) {
	if o == nil || o.CashBalanceEffectFlag == nil {
		return nil, false
	}
	return o.CashBalanceEffectFlag, true
}

// HasCashBalanceEffectFlag returns a boolean if a field has been set.
func (o *Transaction) HasCashBalanceEffectFlag() bool {
	if o != nil && o.CashBalanceEffectFlag != nil {
		return true
	}

	return false
}

// SetCashBalanceEffectFlag gets a reference to the given bool and assigns it to the CashBalanceEffectFlag field.
func (o *Transaction) SetCashBalanceEffectFlag(v bool) {
	o.CashBalanceEffectFlag = &v
}

// GetClearingReferenceNumber returns the ClearingReferenceNumber field value if set, zero value otherwise.
func (o *Transaction) GetClearingReferenceNumber() string {
	if o == nil || o.ClearingReferenceNumber == nil {
		var ret string
		return ret
	}
	return *o.ClearingReferenceNumber
}

// GetClearingReferenceNumberOk returns a tuple with the ClearingReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetClearingReferenceNumberOk() (*string, bool) {
	if o == nil || o.ClearingReferenceNumber == nil {
		return nil, false
	}
	return o.ClearingReferenceNumber, true
}

// HasClearingReferenceNumber returns a boolean if a field has been set.
func (o *Transaction) HasClearingReferenceNumber() bool {
	if o != nil && o.ClearingReferenceNumber != nil {
		return true
	}

	return false
}

// SetClearingReferenceNumber gets a reference to the given string and assigns it to the ClearingReferenceNumber field.
func (o *Transaction) SetClearingReferenceNumber(v string) {
	o.ClearingReferenceNumber = &v
}

// GetDayTradeBuyingPowerEffect returns the DayTradeBuyingPowerEffect field value if set, zero value otherwise.
func (o *Transaction) GetDayTradeBuyingPowerEffect() float32 {
	if o == nil || o.DayTradeBuyingPowerEffect == nil {
		var ret float32
		return ret
	}
	return *o.DayTradeBuyingPowerEffect
}

// GetDayTradeBuyingPowerEffectOk returns a tuple with the DayTradeBuyingPowerEffect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetDayTradeBuyingPowerEffectOk() (*float32, bool) {
	if o == nil || o.DayTradeBuyingPowerEffect == nil {
		return nil, false
	}
	return o.DayTradeBuyingPowerEffect, true
}

// HasDayTradeBuyingPowerEffect returns a boolean if a field has been set.
func (o *Transaction) HasDayTradeBuyingPowerEffect() bool {
	if o != nil && o.DayTradeBuyingPowerEffect != nil {
		return true
	}

	return false
}

// SetDayTradeBuyingPowerEffect gets a reference to the given float32 and assigns it to the DayTradeBuyingPowerEffect field.
func (o *Transaction) SetDayTradeBuyingPowerEffect(v float32) {
	o.DayTradeBuyingPowerEffect = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Transaction) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Transaction) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Transaction) SetDescription(v string) {
	o.Description = &v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *Transaction) GetFees() string {
	if o == nil || o.Fees == nil {
		var ret string
		return ret
	}
	return *o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetFeesOk() (*string, bool) {
	if o == nil || o.Fees == nil {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *Transaction) HasFees() bool {
	if o != nil && o.Fees != nil {
		return true
	}

	return false
}

// SetFees gets a reference to the given string and assigns it to the Fees field.
func (o *Transaction) SetFees(v string) {
	o.Fees = &v
}

// GetNetAmount returns the NetAmount field value if set, zero value otherwise.
func (o *Transaction) GetNetAmount() float32 {
	if o == nil || o.NetAmount == nil {
		var ret float32
		return ret
	}
	return *o.NetAmount
}

// GetNetAmountOk returns a tuple with the NetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetNetAmountOk() (*float32, bool) {
	if o == nil || o.NetAmount == nil {
		return nil, false
	}
	return o.NetAmount, true
}

// HasNetAmount returns a boolean if a field has been set.
func (o *Transaction) HasNetAmount() bool {
	if o != nil && o.NetAmount != nil {
		return true
	}

	return false
}

// SetNetAmount gets a reference to the given float32 and assigns it to the NetAmount field.
func (o *Transaction) SetNetAmount(v float32) {
	o.NetAmount = &v
}

// GetOrderDate returns the OrderDate field value if set, zero value otherwise.
func (o *Transaction) GetOrderDate() string {
	if o == nil || o.OrderDate == nil {
		var ret string
		return ret
	}
	return *o.OrderDate
}

// GetOrderDateOk returns a tuple with the OrderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetOrderDateOk() (*string, bool) {
	if o == nil || o.OrderDate == nil {
		return nil, false
	}
	return o.OrderDate, true
}

// HasOrderDate returns a boolean if a field has been set.
func (o *Transaction) HasOrderDate() bool {
	if o != nil && o.OrderDate != nil {
		return true
	}

	return false
}

// SetOrderDate gets a reference to the given string and assigns it to the OrderDate field.
func (o *Transaction) SetOrderDate(v string) {
	o.OrderDate = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *Transaction) GetOrderId() string {
	if o == nil || o.OrderId == nil {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetOrderIdOk() (*string, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *Transaction) HasOrderId() bool {
	if o != nil && o.OrderId != nil {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *Transaction) SetOrderId(v string) {
	o.OrderId = &v
}

// GetRequirementReallocationAmount returns the RequirementReallocationAmount field value if set, zero value otherwise.
func (o *Transaction) GetRequirementReallocationAmount() float32 {
	if o == nil || o.RequirementReallocationAmount == nil {
		var ret float32
		return ret
	}
	return *o.RequirementReallocationAmount
}

// GetRequirementReallocationAmountOk returns a tuple with the RequirementReallocationAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetRequirementReallocationAmountOk() (*float32, bool) {
	if o == nil || o.RequirementReallocationAmount == nil {
		return nil, false
	}
	return o.RequirementReallocationAmount, true
}

// HasRequirementReallocationAmount returns a boolean if a field has been set.
func (o *Transaction) HasRequirementReallocationAmount() bool {
	if o != nil && o.RequirementReallocationAmount != nil {
		return true
	}

	return false
}

// SetRequirementReallocationAmount gets a reference to the given float32 and assigns it to the RequirementReallocationAmount field.
func (o *Transaction) SetRequirementReallocationAmount(v float32) {
	o.RequirementReallocationAmount = &v
}

// GetSettlementDate returns the SettlementDate field value if set, zero value otherwise.
func (o *Transaction) GetSettlementDate() string {
	if o == nil || o.SettlementDate == nil {
		var ret string
		return ret
	}
	return *o.SettlementDate
}

// GetSettlementDateOk returns a tuple with the SettlementDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetSettlementDateOk() (*string, bool) {
	if o == nil || o.SettlementDate == nil {
		return nil, false
	}
	return o.SettlementDate, true
}

// HasSettlementDate returns a boolean if a field has been set.
func (o *Transaction) HasSettlementDate() bool {
	if o != nil && o.SettlementDate != nil {
		return true
	}

	return false
}

// SetSettlementDate gets a reference to the given string and assigns it to the SettlementDate field.
func (o *Transaction) SetSettlementDate(v string) {
	o.SettlementDate = &v
}

// GetSma returns the Sma field value if set, zero value otherwise.
func (o *Transaction) GetSma() float32 {
	if o == nil || o.Sma == nil {
		var ret float32
		return ret
	}
	return *o.Sma
}

// GetSmaOk returns a tuple with the Sma field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetSmaOk() (*float32, bool) {
	if o == nil || o.Sma == nil {
		return nil, false
	}
	return o.Sma, true
}

// HasSma returns a boolean if a field has been set.
func (o *Transaction) HasSma() bool {
	if o != nil && o.Sma != nil {
		return true
	}

	return false
}

// SetSma gets a reference to the given float32 and assigns it to the Sma field.
func (o *Transaction) SetSma(v float32) {
	o.Sma = &v
}

// GetSubAccount returns the SubAccount field value if set, zero value otherwise.
func (o *Transaction) GetSubAccount() string {
	if o == nil || o.SubAccount == nil {
		var ret string
		return ret
	}
	return *o.SubAccount
}

// GetSubAccountOk returns a tuple with the SubAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetSubAccountOk() (*string, bool) {
	if o == nil || o.SubAccount == nil {
		return nil, false
	}
	return o.SubAccount, true
}

// HasSubAccount returns a boolean if a field has been set.
func (o *Transaction) HasSubAccount() bool {
	if o != nil && o.SubAccount != nil {
		return true
	}

	return false
}

// SetSubAccount gets a reference to the given string and assigns it to the SubAccount field.
func (o *Transaction) SetSubAccount(v string) {
	o.SubAccount = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *Transaction) GetTransactionDate() string {
	if o == nil || o.TransactionDate == nil {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionDateOk() (*string, bool) {
	if o == nil || o.TransactionDate == nil {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *Transaction) HasTransactionDate() bool {
	if o != nil && o.TransactionDate != nil {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *Transaction) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *Transaction) GetTransactionId() float32 {
	if o == nil || o.TransactionId == nil {
		var ret float32
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionIdOk() (*float32, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *Transaction) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given float32 and assigns it to the TransactionId field.
func (o *Transaction) SetTransactionId(v float32) {
	o.TransactionId = &v
}

// GetTransactionItem returns the TransactionItem field value if set, zero value otherwise.
func (o *Transaction) GetTransactionItem() TransactionTransactionItem {
	if o == nil || o.TransactionItem == nil {
		var ret TransactionTransactionItem
		return ret
	}
	return *o.TransactionItem
}

// GetTransactionItemOk returns a tuple with the TransactionItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionItemOk() (*TransactionTransactionItem, bool) {
	if o == nil || o.TransactionItem == nil {
		return nil, false
	}
	return o.TransactionItem, true
}

// HasTransactionItem returns a boolean if a field has been set.
func (o *Transaction) HasTransactionItem() bool {
	if o != nil && o.TransactionItem != nil {
		return true
	}

	return false
}

// SetTransactionItem gets a reference to the given TransactionTransactionItem and assigns it to the TransactionItem field.
func (o *Transaction) SetTransactionItem(v TransactionTransactionItem) {
	o.TransactionItem = &v
}

// GetTransactionSubType returns the TransactionSubType field value if set, zero value otherwise.
func (o *Transaction) GetTransactionSubType() string {
	if o == nil || o.TransactionSubType == nil {
		var ret string
		return ret
	}
	return *o.TransactionSubType
}

// GetTransactionSubTypeOk returns a tuple with the TransactionSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionSubTypeOk() (*string, bool) {
	if o == nil || o.TransactionSubType == nil {
		return nil, false
	}
	return o.TransactionSubType, true
}

// HasTransactionSubType returns a boolean if a field has been set.
func (o *Transaction) HasTransactionSubType() bool {
	if o != nil && o.TransactionSubType != nil {
		return true
	}

	return false
}

// SetTransactionSubType gets a reference to the given string and assigns it to the TransactionSubType field.
func (o *Transaction) SetTransactionSubType(v string) {
	o.TransactionSubType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Transaction) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Transaction) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Transaction) SetType(v string) {
	o.Type = &v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccruedInterest != nil {
		toSerialize["accruedInterest"] = o.AccruedInterest
	}
	if o.AchStatus != nil {
		toSerialize["achStatus"] = o.AchStatus
	}
	if o.CashBalanceEffectFlag != nil {
		toSerialize["cashBalanceEffectFlag"] = o.CashBalanceEffectFlag
	}
	if o.ClearingReferenceNumber != nil {
		toSerialize["clearingReferenceNumber"] = o.ClearingReferenceNumber
	}
	if o.DayTradeBuyingPowerEffect != nil {
		toSerialize["dayTradeBuyingPowerEffect"] = o.DayTradeBuyingPowerEffect
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Fees != nil {
		toSerialize["fees"] = o.Fees
	}
	if o.NetAmount != nil {
		toSerialize["netAmount"] = o.NetAmount
	}
	if o.OrderDate != nil {
		toSerialize["orderDate"] = o.OrderDate
	}
	if o.OrderId != nil {
		toSerialize["orderId"] = o.OrderId
	}
	if o.RequirementReallocationAmount != nil {
		toSerialize["requirementReallocationAmount"] = o.RequirementReallocationAmount
	}
	if o.SettlementDate != nil {
		toSerialize["settlementDate"] = o.SettlementDate
	}
	if o.Sma != nil {
		toSerialize["sma"] = o.Sma
	}
	if o.SubAccount != nil {
		toSerialize["subAccount"] = o.SubAccount
	}
	if o.TransactionDate != nil {
		toSerialize["transactionDate"] = o.TransactionDate
	}
	if o.TransactionId != nil {
		toSerialize["transactionId"] = o.TransactionId
	}
	if o.TransactionItem != nil {
		toSerialize["transactionItem"] = o.TransactionItem
	}
	if o.TransactionSubType != nil {
		toSerialize["transactionSubType"] = o.TransactionSubType
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

