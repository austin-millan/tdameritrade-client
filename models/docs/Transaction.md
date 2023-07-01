# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedInterest** | Pointer to **float32** |  | [optional] 
**AchStatus** | Pointer to **string** |  | [optional] 
**CashBalanceEffectFlag** | Pointer to **bool** |  | [optional] 
**ClearingReferenceNumber** | Pointer to **string** |  | [optional] 
**DayTradeBuyingPowerEffect** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Fees** | Pointer to **string** |  | [optional] 
**NetAmount** | Pointer to **float32** |  | [optional] 
**OrderDate** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**RequirementReallocationAmount** | Pointer to **float32** |  | [optional] 
**SettlementDate** | Pointer to **string** |  | [optional] 
**Sma** | Pointer to **float32** |  | [optional] 
**SubAccount** | Pointer to **string** |  | [optional] 
**TransactionDate** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **float32** |  | [optional] 
**TransactionItem** | Pointer to [**TransactionTransactionItem**](Transaction_transactionItem.md) |  | [optional] 
**TransactionSubType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedInterest

`func (o *Transaction) GetAccruedInterest() float32`

GetAccruedInterest returns the AccruedInterest field if non-nil, zero value otherwise.

### GetAccruedInterestOk

`func (o *Transaction) GetAccruedInterestOk() (*float32, bool)`

GetAccruedInterestOk returns a tuple with the AccruedInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedInterest

`func (o *Transaction) SetAccruedInterest(v float32)`

SetAccruedInterest sets AccruedInterest field to given value.

### HasAccruedInterest

`func (o *Transaction) HasAccruedInterest() bool`

HasAccruedInterest returns a boolean if a field has been set.

### GetAchStatus

`func (o *Transaction) GetAchStatus() string`

GetAchStatus returns the AchStatus field if non-nil, zero value otherwise.

### GetAchStatusOk

`func (o *Transaction) GetAchStatusOk() (*string, bool)`

GetAchStatusOk returns a tuple with the AchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchStatus

`func (o *Transaction) SetAchStatus(v string)`

SetAchStatus sets AchStatus field to given value.

### HasAchStatus

`func (o *Transaction) HasAchStatus() bool`

HasAchStatus returns a boolean if a field has been set.

### GetCashBalanceEffectFlag

`func (o *Transaction) GetCashBalanceEffectFlag() bool`

GetCashBalanceEffectFlag returns the CashBalanceEffectFlag field if non-nil, zero value otherwise.

### GetCashBalanceEffectFlagOk

`func (o *Transaction) GetCashBalanceEffectFlagOk() (*bool, bool)`

GetCashBalanceEffectFlagOk returns a tuple with the CashBalanceEffectFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBalanceEffectFlag

`func (o *Transaction) SetCashBalanceEffectFlag(v bool)`

SetCashBalanceEffectFlag sets CashBalanceEffectFlag field to given value.

### HasCashBalanceEffectFlag

`func (o *Transaction) HasCashBalanceEffectFlag() bool`

HasCashBalanceEffectFlag returns a boolean if a field has been set.

### GetClearingReferenceNumber

`func (o *Transaction) GetClearingReferenceNumber() string`

GetClearingReferenceNumber returns the ClearingReferenceNumber field if non-nil, zero value otherwise.

### GetClearingReferenceNumberOk

`func (o *Transaction) GetClearingReferenceNumberOk() (*string, bool)`

GetClearingReferenceNumberOk returns a tuple with the ClearingReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingReferenceNumber

`func (o *Transaction) SetClearingReferenceNumber(v string)`

SetClearingReferenceNumber sets ClearingReferenceNumber field to given value.

### HasClearingReferenceNumber

`func (o *Transaction) HasClearingReferenceNumber() bool`

HasClearingReferenceNumber returns a boolean if a field has been set.

### GetDayTradeBuyingPowerEffect

`func (o *Transaction) GetDayTradeBuyingPowerEffect() float32`

GetDayTradeBuyingPowerEffect returns the DayTradeBuyingPowerEffect field if non-nil, zero value otherwise.

### GetDayTradeBuyingPowerEffectOk

`func (o *Transaction) GetDayTradeBuyingPowerEffectOk() (*float32, bool)`

GetDayTradeBuyingPowerEffectOk returns a tuple with the DayTradeBuyingPowerEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradeBuyingPowerEffect

`func (o *Transaction) SetDayTradeBuyingPowerEffect(v float32)`

SetDayTradeBuyingPowerEffect sets DayTradeBuyingPowerEffect field to given value.

### HasDayTradeBuyingPowerEffect

`func (o *Transaction) HasDayTradeBuyingPowerEffect() bool`

HasDayTradeBuyingPowerEffect returns a boolean if a field has been set.

### GetDescription

`func (o *Transaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFees

`func (o *Transaction) GetFees() string`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *Transaction) GetFeesOk() (*string, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *Transaction) SetFees(v string)`

SetFees sets Fees field to given value.

### HasFees

`func (o *Transaction) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetNetAmount

`func (o *Transaction) GetNetAmount() float32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *Transaction) GetNetAmountOk() (*float32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *Transaction) SetNetAmount(v float32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *Transaction) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetOrderDate

`func (o *Transaction) GetOrderDate() string`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *Transaction) GetOrderDateOk() (*string, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *Transaction) SetOrderDate(v string)`

SetOrderDate sets OrderDate field to given value.

### HasOrderDate

`func (o *Transaction) HasOrderDate() bool`

HasOrderDate returns a boolean if a field has been set.

### GetOrderId

`func (o *Transaction) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Transaction) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Transaction) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Transaction) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetRequirementReallocationAmount

`func (o *Transaction) GetRequirementReallocationAmount() float32`

GetRequirementReallocationAmount returns the RequirementReallocationAmount field if non-nil, zero value otherwise.

### GetRequirementReallocationAmountOk

`func (o *Transaction) GetRequirementReallocationAmountOk() (*float32, bool)`

GetRequirementReallocationAmountOk returns a tuple with the RequirementReallocationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementReallocationAmount

`func (o *Transaction) SetRequirementReallocationAmount(v float32)`

SetRequirementReallocationAmount sets RequirementReallocationAmount field to given value.

### HasRequirementReallocationAmount

`func (o *Transaction) HasRequirementReallocationAmount() bool`

HasRequirementReallocationAmount returns a boolean if a field has been set.

### GetSettlementDate

`func (o *Transaction) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *Transaction) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *Transaction) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *Transaction) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetSma

`func (o *Transaction) GetSma() float32`

GetSma returns the Sma field if non-nil, zero value otherwise.

### GetSmaOk

`func (o *Transaction) GetSmaOk() (*float32, bool)`

GetSmaOk returns a tuple with the Sma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSma

`func (o *Transaction) SetSma(v float32)`

SetSma sets Sma field to given value.

### HasSma

`func (o *Transaction) HasSma() bool`

HasSma returns a boolean if a field has been set.

### GetSubAccount

`func (o *Transaction) GetSubAccount() string`

GetSubAccount returns the SubAccount field if non-nil, zero value otherwise.

### GetSubAccountOk

`func (o *Transaction) GetSubAccountOk() (*string, bool)`

GetSubAccountOk returns a tuple with the SubAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccount

`func (o *Transaction) SetSubAccount(v string)`

SetSubAccount sets SubAccount field to given value.

### HasSubAccount

`func (o *Transaction) HasSubAccount() bool`

HasSubAccount returns a boolean if a field has been set.

### GetTransactionDate

`func (o *Transaction) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *Transaction) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *Transaction) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *Transaction) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *Transaction) GetTransactionId() float32`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Transaction) GetTransactionIdOk() (*float32, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Transaction) SetTransactionId(v float32)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Transaction) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionItem

`func (o *Transaction) GetTransactionItem() TransactionTransactionItem`

GetTransactionItem returns the TransactionItem field if non-nil, zero value otherwise.

### GetTransactionItemOk

`func (o *Transaction) GetTransactionItemOk() (*TransactionTransactionItem, bool)`

GetTransactionItemOk returns a tuple with the TransactionItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionItem

`func (o *Transaction) SetTransactionItem(v TransactionTransactionItem)`

SetTransactionItem sets TransactionItem field to given value.

### HasTransactionItem

`func (o *Transaction) HasTransactionItem() bool`

HasTransactionItem returns a boolean if a field has been set.

### GetTransactionSubType

`func (o *Transaction) GetTransactionSubType() string`

GetTransactionSubType returns the TransactionSubType field if non-nil, zero value otherwise.

### GetTransactionSubTypeOk

`func (o *Transaction) GetTransactionSubTypeOk() (*string, bool)`

GetTransactionSubTypeOk returns a tuple with the TransactionSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSubType

`func (o *Transaction) SetTransactionSubType(v string)`

SetTransactionSubType sets TransactionSubType field to given value.

### HasTransactionSubType

`func (o *Transaction) HasTransactionSubType() bool`

HasTransactionSubType returns a boolean if a field has been set.

### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Transaction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


