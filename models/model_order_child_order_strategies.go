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

// OrderChildOrderStrategies struct for OrderChildOrderStrategies
type OrderChildOrderStrategies struct {
	CancelTime *CashAccountCancelTime `json:"cancelTime,omitempty"`
	ComplexOrderStrategyType *ComplexOrderStrategyType `json:"complexOrderStrategyType,omitempty"`
	DestinationLinkName *string `json:"destinationLinkName,omitempty"`
	Duration *Duration `json:"duration,omitempty"`
	FilledQuantity *float32 `json:"filledQuantity,omitempty"`
	OrderType *OrderType `json:"orderType,omitempty"`
	Price *float32 `json:"price,omitempty"`
	PriceLinkBasis *PriceLinkBasis `json:"priceLinkBasis,omitempty"`
	PriceLinkType *PriceLinkType `json:"priceLinkType,omitempty"`
	Quantity *float32 `json:"quantity,omitempty"`
	ReleaseTime *string `json:"releaseTime,omitempty"`
	RemainingQuantity *float32 `json:"remainingQuantity,omitempty"`
	RequestedDestination *DestinationExchange `json:"requestedDestination,omitempty"`
	Session *Session `json:"session,omitempty"`
	StopPrice *float32 `json:"stopPrice,omitempty"`
	StopPriceLinkBasis *PriceLinkBasis `json:"stopPriceLinkBasis,omitempty"`
	StopPriceLinkType *PriceLinkType `json:"stopPriceLinkType,omitempty"`
	StopPriceOffset *float32 `json:"stopPriceOffset,omitempty"`
	StopType *StopType `json:"stopType,omitempty"`
	TaxLotMethod *TaxLotMethod `json:"taxLotMethod,omitempty"`
}

// NewOrderChildOrderStrategies instantiates a new OrderChildOrderStrategies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderChildOrderStrategies() *OrderChildOrderStrategies {
	this := OrderChildOrderStrategies{}
	return &this
}

// NewOrderChildOrderStrategiesWithDefaults instantiates a new OrderChildOrderStrategies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderChildOrderStrategiesWithDefaults() *OrderChildOrderStrategies {
	this := OrderChildOrderStrategies{}
	return &this
}

// GetCancelTime returns the CancelTime field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetCancelTime() CashAccountCancelTime {
	if o == nil || o.CancelTime == nil {
		var ret CashAccountCancelTime
		return ret
	}
	return *o.CancelTime
}

// GetCancelTimeOk returns a tuple with the CancelTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetCancelTimeOk() (*CashAccountCancelTime, bool) {
	if o == nil || o.CancelTime == nil {
		return nil, false
	}
	return o.CancelTime, true
}

// HasCancelTime returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasCancelTime() bool {
	if o != nil && o.CancelTime != nil {
		return true
	}

	return false
}

// SetCancelTime gets a reference to the given CashAccountCancelTime and assigns it to the CancelTime field.
func (o *OrderChildOrderStrategies) SetCancelTime(v CashAccountCancelTime) {
	o.CancelTime = &v
}

// GetComplexOrderStrategyType returns the ComplexOrderStrategyType field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetComplexOrderStrategyType() ComplexOrderStrategyType {
	if o == nil || o.ComplexOrderStrategyType == nil {
		var ret ComplexOrderStrategyType
		return ret
	}
	return *o.ComplexOrderStrategyType
}

// GetComplexOrderStrategyTypeOk returns a tuple with the ComplexOrderStrategyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetComplexOrderStrategyTypeOk() (*ComplexOrderStrategyType, bool) {
	if o == nil || o.ComplexOrderStrategyType == nil {
		return nil, false
	}
	return o.ComplexOrderStrategyType, true
}

// HasComplexOrderStrategyType returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasComplexOrderStrategyType() bool {
	if o != nil && o.ComplexOrderStrategyType != nil {
		return true
	}

	return false
}

// SetComplexOrderStrategyType gets a reference to the given ComplexOrderStrategyType and assigns it to the ComplexOrderStrategyType field.
func (o *OrderChildOrderStrategies) SetComplexOrderStrategyType(v ComplexOrderStrategyType) {
	o.ComplexOrderStrategyType = &v
}

// GetDestinationLinkName returns the DestinationLinkName field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetDestinationLinkName() string {
	if o == nil || o.DestinationLinkName == nil {
		var ret string
		return ret
	}
	return *o.DestinationLinkName
}

// GetDestinationLinkNameOk returns a tuple with the DestinationLinkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetDestinationLinkNameOk() (*string, bool) {
	if o == nil || o.DestinationLinkName == nil {
		return nil, false
	}
	return o.DestinationLinkName, true
}

// HasDestinationLinkName returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasDestinationLinkName() bool {
	if o != nil && o.DestinationLinkName != nil {
		return true
	}

	return false
}

// SetDestinationLinkName gets a reference to the given string and assigns it to the DestinationLinkName field.
func (o *OrderChildOrderStrategies) SetDestinationLinkName(v string) {
	o.DestinationLinkName = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetDuration() Duration {
	if o == nil || o.Duration == nil {
		var ret Duration
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetDurationOk() (*Duration, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given Duration and assigns it to the Duration field.
func (o *OrderChildOrderStrategies) SetDuration(v Duration) {
	o.Duration = &v
}

// GetFilledQuantity returns the FilledQuantity field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetFilledQuantity() float32 {
	if o == nil || o.FilledQuantity == nil {
		var ret float32
		return ret
	}
	return *o.FilledQuantity
}

// GetFilledQuantityOk returns a tuple with the FilledQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetFilledQuantityOk() (*float32, bool) {
	if o == nil || o.FilledQuantity == nil {
		return nil, false
	}
	return o.FilledQuantity, true
}

// HasFilledQuantity returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasFilledQuantity() bool {
	if o != nil && o.FilledQuantity != nil {
		return true
	}

	return false
}

// SetFilledQuantity gets a reference to the given float32 and assigns it to the FilledQuantity field.
func (o *OrderChildOrderStrategies) SetFilledQuantity(v float32) {
	o.FilledQuantity = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetOrderType() OrderType {
	if o == nil || o.OrderType == nil {
		var ret OrderType
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetOrderTypeOk() (*OrderType, bool) {
	if o == nil || o.OrderType == nil {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasOrderType() bool {
	if o != nil && o.OrderType != nil {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given OrderType and assigns it to the OrderType field.
func (o *OrderChildOrderStrategies) SetOrderType(v OrderType) {
	o.OrderType = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *OrderChildOrderStrategies) SetPrice(v float32) {
	o.Price = &v
}

// GetPriceLinkBasis returns the PriceLinkBasis field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetPriceLinkBasis() PriceLinkBasis {
	if o == nil || o.PriceLinkBasis == nil {
		var ret PriceLinkBasis
		return ret
	}
	return *o.PriceLinkBasis
}

// GetPriceLinkBasisOk returns a tuple with the PriceLinkBasis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetPriceLinkBasisOk() (*PriceLinkBasis, bool) {
	if o == nil || o.PriceLinkBasis == nil {
		return nil, false
	}
	return o.PriceLinkBasis, true
}

// HasPriceLinkBasis returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasPriceLinkBasis() bool {
	if o != nil && o.PriceLinkBasis != nil {
		return true
	}

	return false
}

// SetPriceLinkBasis gets a reference to the given PriceLinkBasis and assigns it to the PriceLinkBasis field.
func (o *OrderChildOrderStrategies) SetPriceLinkBasis(v PriceLinkBasis) {
	o.PriceLinkBasis = &v
}

// GetPriceLinkType returns the PriceLinkType field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetPriceLinkType() PriceLinkType {
	if o == nil || o.PriceLinkType == nil {
		var ret PriceLinkType
		return ret
	}
	return *o.PriceLinkType
}

// GetPriceLinkTypeOk returns a tuple with the PriceLinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetPriceLinkTypeOk() (*PriceLinkType, bool) {
	if o == nil || o.PriceLinkType == nil {
		return nil, false
	}
	return o.PriceLinkType, true
}

// HasPriceLinkType returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasPriceLinkType() bool {
	if o != nil && o.PriceLinkType != nil {
		return true
	}

	return false
}

// SetPriceLinkType gets a reference to the given PriceLinkType and assigns it to the PriceLinkType field.
func (o *OrderChildOrderStrategies) SetPriceLinkType(v PriceLinkType) {
	o.PriceLinkType = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetQuantity() float32 {
	if o == nil || o.Quantity == nil {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetQuantityOk() (*float32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *OrderChildOrderStrategies) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetReleaseTime returns the ReleaseTime field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetReleaseTime() string {
	if o == nil || o.ReleaseTime == nil {
		var ret string
		return ret
	}
	return *o.ReleaseTime
}

// GetReleaseTimeOk returns a tuple with the ReleaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetReleaseTimeOk() (*string, bool) {
	if o == nil || o.ReleaseTime == nil {
		return nil, false
	}
	return o.ReleaseTime, true
}

// HasReleaseTime returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasReleaseTime() bool {
	if o != nil && o.ReleaseTime != nil {
		return true
	}

	return false
}

// SetReleaseTime gets a reference to the given string and assigns it to the ReleaseTime field.
func (o *OrderChildOrderStrategies) SetReleaseTime(v string) {
	o.ReleaseTime = &v
}

// GetRemainingQuantity returns the RemainingQuantity field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetRemainingQuantity() float32 {
	if o == nil || o.RemainingQuantity == nil {
		var ret float32
		return ret
	}
	return *o.RemainingQuantity
}

// GetRemainingQuantityOk returns a tuple with the RemainingQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetRemainingQuantityOk() (*float32, bool) {
	if o == nil || o.RemainingQuantity == nil {
		return nil, false
	}
	return o.RemainingQuantity, true
}

// HasRemainingQuantity returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasRemainingQuantity() bool {
	if o != nil && o.RemainingQuantity != nil {
		return true
	}

	return false
}

// SetRemainingQuantity gets a reference to the given float32 and assigns it to the RemainingQuantity field.
func (o *OrderChildOrderStrategies) SetRemainingQuantity(v float32) {
	o.RemainingQuantity = &v
}

// GetRequestedDestination returns the RequestedDestination field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetRequestedDestination() DestinationExchange {
	if o == nil || o.RequestedDestination == nil {
		var ret DestinationExchange
		return ret
	}
	return *o.RequestedDestination
}

// GetRequestedDestinationOk returns a tuple with the RequestedDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetRequestedDestinationOk() (*DestinationExchange, bool) {
	if o == nil || o.RequestedDestination == nil {
		return nil, false
	}
	return o.RequestedDestination, true
}

// HasRequestedDestination returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasRequestedDestination() bool {
	if o != nil && o.RequestedDestination != nil {
		return true
	}

	return false
}

// SetRequestedDestination gets a reference to the given DestinationExchange and assigns it to the RequestedDestination field.
func (o *OrderChildOrderStrategies) SetRequestedDestination(v DestinationExchange) {
	o.RequestedDestination = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetSession() Session {
	if o == nil || o.Session == nil {
		var ret Session
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetSessionOk() (*Session, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given Session and assigns it to the Session field.
func (o *OrderChildOrderStrategies) SetSession(v Session) {
	o.Session = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetStopPrice() float32 {
	if o == nil || o.StopPrice == nil {
		var ret float32
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetStopPriceOk() (*float32, bool) {
	if o == nil || o.StopPrice == nil {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasStopPrice() bool {
	if o != nil && o.StopPrice != nil {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given float32 and assigns it to the StopPrice field.
func (o *OrderChildOrderStrategies) SetStopPrice(v float32) {
	o.StopPrice = &v
}

// GetStopPriceLinkBasis returns the StopPriceLinkBasis field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetStopPriceLinkBasis() PriceLinkBasis {
	if o == nil || o.StopPriceLinkBasis == nil {
		var ret PriceLinkBasis
		return ret
	}
	return *o.StopPriceLinkBasis
}

// GetStopPriceLinkBasisOk returns a tuple with the StopPriceLinkBasis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetStopPriceLinkBasisOk() (*PriceLinkBasis, bool) {
	if o == nil || o.StopPriceLinkBasis == nil {
		return nil, false
	}
	return o.StopPriceLinkBasis, true
}

// HasStopPriceLinkBasis returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasStopPriceLinkBasis() bool {
	if o != nil && o.StopPriceLinkBasis != nil {
		return true
	}

	return false
}

// SetStopPriceLinkBasis gets a reference to the given PriceLinkBasis and assigns it to the StopPriceLinkBasis field.
func (o *OrderChildOrderStrategies) SetStopPriceLinkBasis(v PriceLinkBasis) {
	o.StopPriceLinkBasis = &v
}

// GetStopPriceLinkType returns the StopPriceLinkType field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetStopPriceLinkType() PriceLinkType {
	if o == nil || o.StopPriceLinkType == nil {
		var ret PriceLinkType
		return ret
	}
	return *o.StopPriceLinkType
}

// GetStopPriceLinkTypeOk returns a tuple with the StopPriceLinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetStopPriceLinkTypeOk() (*PriceLinkType, bool) {
	if o == nil || o.StopPriceLinkType == nil {
		return nil, false
	}
	return o.StopPriceLinkType, true
}

// HasStopPriceLinkType returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasStopPriceLinkType() bool {
	if o != nil && o.StopPriceLinkType != nil {
		return true
	}

	return false
}

// SetStopPriceLinkType gets a reference to the given PriceLinkType and assigns it to the StopPriceLinkType field.
func (o *OrderChildOrderStrategies) SetStopPriceLinkType(v PriceLinkType) {
	o.StopPriceLinkType = &v
}

// GetStopPriceOffset returns the StopPriceOffset field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetStopPriceOffset() float32 {
	if o == nil || o.StopPriceOffset == nil {
		var ret float32
		return ret
	}
	return *o.StopPriceOffset
}

// GetStopPriceOffsetOk returns a tuple with the StopPriceOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetStopPriceOffsetOk() (*float32, bool) {
	if o == nil || o.StopPriceOffset == nil {
		return nil, false
	}
	return o.StopPriceOffset, true
}

// HasStopPriceOffset returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasStopPriceOffset() bool {
	if o != nil && o.StopPriceOffset != nil {
		return true
	}

	return false
}

// SetStopPriceOffset gets a reference to the given float32 and assigns it to the StopPriceOffset field.
func (o *OrderChildOrderStrategies) SetStopPriceOffset(v float32) {
	o.StopPriceOffset = &v
}

// GetStopType returns the StopType field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetStopType() StopType {
	if o == nil || o.StopType == nil {
		var ret StopType
		return ret
	}
	return *o.StopType
}

// GetStopTypeOk returns a tuple with the StopType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetStopTypeOk() (*StopType, bool) {
	if o == nil || o.StopType == nil {
		return nil, false
	}
	return o.StopType, true
}

// HasStopType returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasStopType() bool {
	if o != nil && o.StopType != nil {
		return true
	}

	return false
}

// SetStopType gets a reference to the given StopType and assigns it to the StopType field.
func (o *OrderChildOrderStrategies) SetStopType(v StopType) {
	o.StopType = &v
}

// GetTaxLotMethod returns the TaxLotMethod field value if set, zero value otherwise.
func (o *OrderChildOrderStrategies) GetTaxLotMethod() TaxLotMethod {
	if o == nil || o.TaxLotMethod == nil {
		var ret TaxLotMethod
		return ret
	}
	return *o.TaxLotMethod
}

// GetTaxLotMethodOk returns a tuple with the TaxLotMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChildOrderStrategies) GetTaxLotMethodOk() (*TaxLotMethod, bool) {
	if o == nil || o.TaxLotMethod == nil {
		return nil, false
	}
	return o.TaxLotMethod, true
}

// HasTaxLotMethod returns a boolean if a field has been set.
func (o *OrderChildOrderStrategies) HasTaxLotMethod() bool {
	if o != nil && o.TaxLotMethod != nil {
		return true
	}

	return false
}

// SetTaxLotMethod gets a reference to the given TaxLotMethod and assigns it to the TaxLotMethod field.
func (o *OrderChildOrderStrategies) SetTaxLotMethod(v TaxLotMethod) {
	o.TaxLotMethod = &v
}

func (o OrderChildOrderStrategies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CancelTime != nil {
		toSerialize["cancelTime"] = o.CancelTime
	}
	if o.ComplexOrderStrategyType != nil {
		toSerialize["complexOrderStrategyType"] = o.ComplexOrderStrategyType
	}
	if o.DestinationLinkName != nil {
		toSerialize["destinationLinkName"] = o.DestinationLinkName
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.FilledQuantity != nil {
		toSerialize["filledQuantity"] = o.FilledQuantity
	}
	if o.OrderType != nil {
		toSerialize["orderType"] = o.OrderType
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.PriceLinkBasis != nil {
		toSerialize["priceLinkBasis"] = o.PriceLinkBasis
	}
	if o.PriceLinkType != nil {
		toSerialize["priceLinkType"] = o.PriceLinkType
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.ReleaseTime != nil {
		toSerialize["releaseTime"] = o.ReleaseTime
	}
	if o.RemainingQuantity != nil {
		toSerialize["remainingQuantity"] = o.RemainingQuantity
	}
	if o.RequestedDestination != nil {
		toSerialize["requestedDestination"] = o.RequestedDestination
	}
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}
	if o.StopPrice != nil {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if o.StopPriceLinkBasis != nil {
		toSerialize["stopPriceLinkBasis"] = o.StopPriceLinkBasis
	}
	if o.StopPriceLinkType != nil {
		toSerialize["stopPriceLinkType"] = o.StopPriceLinkType
	}
	if o.StopPriceOffset != nil {
		toSerialize["stopPriceOffset"] = o.StopPriceOffset
	}
	if o.StopType != nil {
		toSerialize["stopType"] = o.StopType
	}
	if o.TaxLotMethod != nil {
		toSerialize["taxLotMethod"] = o.TaxLotMethod
	}
	return json.Marshal(toSerialize)
}

type NullableOrderChildOrderStrategies struct {
	value *OrderChildOrderStrategies
	isSet bool
}

func (v NullableOrderChildOrderStrategies) Get() *OrderChildOrderStrategies {
	return v.value
}

func (v *NullableOrderChildOrderStrategies) Set(val *OrderChildOrderStrategies) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderChildOrderStrategies) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderChildOrderStrategies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderChildOrderStrategies(val *OrderChildOrderStrategies) *NullableOrderChildOrderStrategies {
	return &NullableOrderChildOrderStrategies{value: val, isSet: true}
}

func (v NullableOrderChildOrderStrategies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderChildOrderStrategies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

