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

// Preferences struct for Preferences
type Preferences struct {
	AuthTokenTimeout *string `json:"authTokenTimeout,omitempty"`
	DefaultAdvancedToolLaunch *string `json:"defaultAdvancedToolLaunch,omitempty"`
	DefaultEquityOrderDuration *Duration `json:"defaultEquityOrderDuration,omitempty"`
	DefaultEquityOrderLegInstruction *string `json:"defaultEquityOrderLegInstruction,omitempty"`
	DefaultEquityOrderMarketSession *Session `json:"defaultEquityOrderMarketSession,omitempty"`
	DefaultEquityOrderPriceLinkType *EquityPriceLinkType `json:"defaultEquityOrderPriceLinkType,omitempty"`
	DefaultEquityOrderType *OrderType `json:"defaultEquityOrderType,omitempty"`
	DefaultEquityQuantity *float32 `json:"defaultEquityQuantity,omitempty"`
	DirectEquityRouting *bool `json:"directEquityRouting,omitempty"`
	DirectOptionsRouting *bool `json:"directOptionsRouting,omitempty"`
	EquityTaxLotMethod *TaxLotMethod `json:"equityTaxLotMethod,omitempty"`
	ExpressTrading *bool `json:"expressTrading,omitempty"`
	MutualFundTaxLotMethod *TaxLotMethod `json:"mutualFundTaxLotMethod,omitempty"`
	OptionTaxLotMethod *TaxLotMethod `json:"optionTaxLotMethod,omitempty"`
}

// NewPreferences instantiates a new Preferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferences() *Preferences {
	this := Preferences{}
	return &this
}

// NewPreferencesWithDefaults instantiates a new Preferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferencesWithDefaults() *Preferences {
	this := Preferences{}
	return &this
}

// GetAuthTokenTimeout returns the AuthTokenTimeout field value if set, zero value otherwise.
func (o *Preferences) GetAuthTokenTimeout() string {
	if o == nil || o.AuthTokenTimeout == nil {
		var ret string
		return ret
	}
	return *o.AuthTokenTimeout
}

// GetAuthTokenTimeoutOk returns a tuple with the AuthTokenTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetAuthTokenTimeoutOk() (*string, bool) {
	if o == nil || o.AuthTokenTimeout == nil {
		return nil, false
	}
	return o.AuthTokenTimeout, true
}

// HasAuthTokenTimeout returns a boolean if a field has been set.
func (o *Preferences) HasAuthTokenTimeout() bool {
	if o != nil && o.AuthTokenTimeout != nil {
		return true
	}

	return false
}

// SetAuthTokenTimeout gets a reference to the given string and assigns it to the AuthTokenTimeout field.
func (o *Preferences) SetAuthTokenTimeout(v string) {
	o.AuthTokenTimeout = &v
}

// GetDefaultAdvancedToolLaunch returns the DefaultAdvancedToolLaunch field value if set, zero value otherwise.
func (o *Preferences) GetDefaultAdvancedToolLaunch() string {
	if o == nil || o.DefaultAdvancedToolLaunch == nil {
		var ret string
		return ret
	}
	return *o.DefaultAdvancedToolLaunch
}

// GetDefaultAdvancedToolLaunchOk returns a tuple with the DefaultAdvancedToolLaunch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetDefaultAdvancedToolLaunchOk() (*string, bool) {
	if o == nil || o.DefaultAdvancedToolLaunch == nil {
		return nil, false
	}
	return o.DefaultAdvancedToolLaunch, true
}

// HasDefaultAdvancedToolLaunch returns a boolean if a field has been set.
func (o *Preferences) HasDefaultAdvancedToolLaunch() bool {
	if o != nil && o.DefaultAdvancedToolLaunch != nil {
		return true
	}

	return false
}

// SetDefaultAdvancedToolLaunch gets a reference to the given string and assigns it to the DefaultAdvancedToolLaunch field.
func (o *Preferences) SetDefaultAdvancedToolLaunch(v string) {
	o.DefaultAdvancedToolLaunch = &v
}

// GetDefaultEquityOrderDuration returns the DefaultEquityOrderDuration field value if set, zero value otherwise.
func (o *Preferences) GetDefaultEquityOrderDuration() Duration {
	if o == nil || o.DefaultEquityOrderDuration == nil {
		var ret Duration
		return ret
	}
	return *o.DefaultEquityOrderDuration
}

// GetDefaultEquityOrderDurationOk returns a tuple with the DefaultEquityOrderDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetDefaultEquityOrderDurationOk() (*Duration, bool) {
	if o == nil || o.DefaultEquityOrderDuration == nil {
		return nil, false
	}
	return o.DefaultEquityOrderDuration, true
}

// HasDefaultEquityOrderDuration returns a boolean if a field has been set.
func (o *Preferences) HasDefaultEquityOrderDuration() bool {
	if o != nil && o.DefaultEquityOrderDuration != nil {
		return true
	}

	return false
}

// SetDefaultEquityOrderDuration gets a reference to the given Duration and assigns it to the DefaultEquityOrderDuration field.
func (o *Preferences) SetDefaultEquityOrderDuration(v Duration) {
	o.DefaultEquityOrderDuration = &v
}

// GetDefaultEquityOrderLegInstruction returns the DefaultEquityOrderLegInstruction field value if set, zero value otherwise.
func (o *Preferences) GetDefaultEquityOrderLegInstruction() string {
	if o == nil || o.DefaultEquityOrderLegInstruction == nil {
		var ret string
		return ret
	}
	return *o.DefaultEquityOrderLegInstruction
}

// GetDefaultEquityOrderLegInstructionOk returns a tuple with the DefaultEquityOrderLegInstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetDefaultEquityOrderLegInstructionOk() (*string, bool) {
	if o == nil || o.DefaultEquityOrderLegInstruction == nil {
		return nil, false
	}
	return o.DefaultEquityOrderLegInstruction, true
}

// HasDefaultEquityOrderLegInstruction returns a boolean if a field has been set.
func (o *Preferences) HasDefaultEquityOrderLegInstruction() bool {
	if o != nil && o.DefaultEquityOrderLegInstruction != nil {
		return true
	}

	return false
}

// SetDefaultEquityOrderLegInstruction gets a reference to the given string and assigns it to the DefaultEquityOrderLegInstruction field.
func (o *Preferences) SetDefaultEquityOrderLegInstruction(v string) {
	o.DefaultEquityOrderLegInstruction = &v
}

// GetDefaultEquityOrderMarketSession returns the DefaultEquityOrderMarketSession field value if set, zero value otherwise.
func (o *Preferences) GetDefaultEquityOrderMarketSession() Session {
	if o == nil || o.DefaultEquityOrderMarketSession == nil {
		var ret Session
		return ret
	}
	return *o.DefaultEquityOrderMarketSession
}

// GetDefaultEquityOrderMarketSessionOk returns a tuple with the DefaultEquityOrderMarketSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetDefaultEquityOrderMarketSessionOk() (*Session, bool) {
	if o == nil || o.DefaultEquityOrderMarketSession == nil {
		return nil, false
	}
	return o.DefaultEquityOrderMarketSession, true
}

// HasDefaultEquityOrderMarketSession returns a boolean if a field has been set.
func (o *Preferences) HasDefaultEquityOrderMarketSession() bool {
	if o != nil && o.DefaultEquityOrderMarketSession != nil {
		return true
	}

	return false
}

// SetDefaultEquityOrderMarketSession gets a reference to the given Session and assigns it to the DefaultEquityOrderMarketSession field.
func (o *Preferences) SetDefaultEquityOrderMarketSession(v Session) {
	o.DefaultEquityOrderMarketSession = &v
}

// GetDefaultEquityOrderPriceLinkType returns the DefaultEquityOrderPriceLinkType field value if set, zero value otherwise.
func (o *Preferences) GetDefaultEquityOrderPriceLinkType() EquityPriceLinkType {
	if o == nil || o.DefaultEquityOrderPriceLinkType == nil {
		var ret EquityPriceLinkType
		return ret
	}
	return *o.DefaultEquityOrderPriceLinkType
}

// GetDefaultEquityOrderPriceLinkTypeOk returns a tuple with the DefaultEquityOrderPriceLinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetDefaultEquityOrderPriceLinkTypeOk() (*EquityPriceLinkType, bool) {
	if o == nil || o.DefaultEquityOrderPriceLinkType == nil {
		return nil, false
	}
	return o.DefaultEquityOrderPriceLinkType, true
}

// HasDefaultEquityOrderPriceLinkType returns a boolean if a field has been set.
func (o *Preferences) HasDefaultEquityOrderPriceLinkType() bool {
	if o != nil && o.DefaultEquityOrderPriceLinkType != nil {
		return true
	}

	return false
}

// SetDefaultEquityOrderPriceLinkType gets a reference to the given EquityPriceLinkType and assigns it to the DefaultEquityOrderPriceLinkType field.
func (o *Preferences) SetDefaultEquityOrderPriceLinkType(v EquityPriceLinkType) {
	o.DefaultEquityOrderPriceLinkType = &v
}

// GetDefaultEquityOrderType returns the DefaultEquityOrderType field value if set, zero value otherwise.
func (o *Preferences) GetDefaultEquityOrderType() OrderType {
	if o == nil || o.DefaultEquityOrderType == nil {
		var ret OrderType
		return ret
	}
	return *o.DefaultEquityOrderType
}

// GetDefaultEquityOrderTypeOk returns a tuple with the DefaultEquityOrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetDefaultEquityOrderTypeOk() (*OrderType, bool) {
	if o == nil || o.DefaultEquityOrderType == nil {
		return nil, false
	}
	return o.DefaultEquityOrderType, true
}

// HasDefaultEquityOrderType returns a boolean if a field has been set.
func (o *Preferences) HasDefaultEquityOrderType() bool {
	if o != nil && o.DefaultEquityOrderType != nil {
		return true
	}

	return false
}

// SetDefaultEquityOrderType gets a reference to the given OrderType and assigns it to the DefaultEquityOrderType field.
func (o *Preferences) SetDefaultEquityOrderType(v OrderType) {
	o.DefaultEquityOrderType = &v
}

// GetDefaultEquityQuantity returns the DefaultEquityQuantity field value if set, zero value otherwise.
func (o *Preferences) GetDefaultEquityQuantity() float32 {
	if o == nil || o.DefaultEquityQuantity == nil {
		var ret float32
		return ret
	}
	return *o.DefaultEquityQuantity
}

// GetDefaultEquityQuantityOk returns a tuple with the DefaultEquityQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetDefaultEquityQuantityOk() (*float32, bool) {
	if o == nil || o.DefaultEquityQuantity == nil {
		return nil, false
	}
	return o.DefaultEquityQuantity, true
}

// HasDefaultEquityQuantity returns a boolean if a field has been set.
func (o *Preferences) HasDefaultEquityQuantity() bool {
	if o != nil && o.DefaultEquityQuantity != nil {
		return true
	}

	return false
}

// SetDefaultEquityQuantity gets a reference to the given float32 and assigns it to the DefaultEquityQuantity field.
func (o *Preferences) SetDefaultEquityQuantity(v float32) {
	o.DefaultEquityQuantity = &v
}

// GetDirectEquityRouting returns the DirectEquityRouting field value if set, zero value otherwise.
func (o *Preferences) GetDirectEquityRouting() bool {
	if o == nil || o.DirectEquityRouting == nil {
		var ret bool
		return ret
	}
	return *o.DirectEquityRouting
}

// GetDirectEquityRoutingOk returns a tuple with the DirectEquityRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetDirectEquityRoutingOk() (*bool, bool) {
	if o == nil || o.DirectEquityRouting == nil {
		return nil, false
	}
	return o.DirectEquityRouting, true
}

// HasDirectEquityRouting returns a boolean if a field has been set.
func (o *Preferences) HasDirectEquityRouting() bool {
	if o != nil && o.DirectEquityRouting != nil {
		return true
	}

	return false
}

// SetDirectEquityRouting gets a reference to the given bool and assigns it to the DirectEquityRouting field.
func (o *Preferences) SetDirectEquityRouting(v bool) {
	o.DirectEquityRouting = &v
}

// GetDirectOptionsRouting returns the DirectOptionsRouting field value if set, zero value otherwise.
func (o *Preferences) GetDirectOptionsRouting() bool {
	if o == nil || o.DirectOptionsRouting == nil {
		var ret bool
		return ret
	}
	return *o.DirectOptionsRouting
}

// GetDirectOptionsRoutingOk returns a tuple with the DirectOptionsRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetDirectOptionsRoutingOk() (*bool, bool) {
	if o == nil || o.DirectOptionsRouting == nil {
		return nil, false
	}
	return o.DirectOptionsRouting, true
}

// HasDirectOptionsRouting returns a boolean if a field has been set.
func (o *Preferences) HasDirectOptionsRouting() bool {
	if o != nil && o.DirectOptionsRouting != nil {
		return true
	}

	return false
}

// SetDirectOptionsRouting gets a reference to the given bool and assigns it to the DirectOptionsRouting field.
func (o *Preferences) SetDirectOptionsRouting(v bool) {
	o.DirectOptionsRouting = &v
}

// GetEquityTaxLotMethod returns the EquityTaxLotMethod field value if set, zero value otherwise.
func (o *Preferences) GetEquityTaxLotMethod() TaxLotMethod {
	if o == nil || o.EquityTaxLotMethod == nil {
		var ret TaxLotMethod
		return ret
	}
	return *o.EquityTaxLotMethod
}

// GetEquityTaxLotMethodOk returns a tuple with the EquityTaxLotMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetEquityTaxLotMethodOk() (*TaxLotMethod, bool) {
	if o == nil || o.EquityTaxLotMethod == nil {
		return nil, false
	}
	return o.EquityTaxLotMethod, true
}

// HasEquityTaxLotMethod returns a boolean if a field has been set.
func (o *Preferences) HasEquityTaxLotMethod() bool {
	if o != nil && o.EquityTaxLotMethod != nil {
		return true
	}

	return false
}

// SetEquityTaxLotMethod gets a reference to the given TaxLotMethod and assigns it to the EquityTaxLotMethod field.
func (o *Preferences) SetEquityTaxLotMethod(v TaxLotMethod) {
	o.EquityTaxLotMethod = &v
}

// GetExpressTrading returns the ExpressTrading field value if set, zero value otherwise.
func (o *Preferences) GetExpressTrading() bool {
	if o == nil || o.ExpressTrading == nil {
		var ret bool
		return ret
	}
	return *o.ExpressTrading
}

// GetExpressTradingOk returns a tuple with the ExpressTrading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetExpressTradingOk() (*bool, bool) {
	if o == nil || o.ExpressTrading == nil {
		return nil, false
	}
	return o.ExpressTrading, true
}

// HasExpressTrading returns a boolean if a field has been set.
func (o *Preferences) HasExpressTrading() bool {
	if o != nil && o.ExpressTrading != nil {
		return true
	}

	return false
}

// SetExpressTrading gets a reference to the given bool and assigns it to the ExpressTrading field.
func (o *Preferences) SetExpressTrading(v bool) {
	o.ExpressTrading = &v
}

// GetMutualFundTaxLotMethod returns the MutualFundTaxLotMethod field value if set, zero value otherwise.
func (o *Preferences) GetMutualFundTaxLotMethod() TaxLotMethod {
	if o == nil || o.MutualFundTaxLotMethod == nil {
		var ret TaxLotMethod
		return ret
	}
	return *o.MutualFundTaxLotMethod
}

// GetMutualFundTaxLotMethodOk returns a tuple with the MutualFundTaxLotMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetMutualFundTaxLotMethodOk() (*TaxLotMethod, bool) {
	if o == nil || o.MutualFundTaxLotMethod == nil {
		return nil, false
	}
	return o.MutualFundTaxLotMethod, true
}

// HasMutualFundTaxLotMethod returns a boolean if a field has been set.
func (o *Preferences) HasMutualFundTaxLotMethod() bool {
	if o != nil && o.MutualFundTaxLotMethod != nil {
		return true
	}

	return false
}

// SetMutualFundTaxLotMethod gets a reference to the given TaxLotMethod and assigns it to the MutualFundTaxLotMethod field.
func (o *Preferences) SetMutualFundTaxLotMethod(v TaxLotMethod) {
	o.MutualFundTaxLotMethod = &v
}

// GetOptionTaxLotMethod returns the OptionTaxLotMethod field value if set, zero value otherwise.
func (o *Preferences) GetOptionTaxLotMethod() TaxLotMethod {
	if o == nil || o.OptionTaxLotMethod == nil {
		var ret TaxLotMethod
		return ret
	}
	return *o.OptionTaxLotMethod
}

// GetOptionTaxLotMethodOk returns a tuple with the OptionTaxLotMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preferences) GetOptionTaxLotMethodOk() (*TaxLotMethod, bool) {
	if o == nil || o.OptionTaxLotMethod == nil {
		return nil, false
	}
	return o.OptionTaxLotMethod, true
}

// HasOptionTaxLotMethod returns a boolean if a field has been set.
func (o *Preferences) HasOptionTaxLotMethod() bool {
	if o != nil && o.OptionTaxLotMethod != nil {
		return true
	}

	return false
}

// SetOptionTaxLotMethod gets a reference to the given TaxLotMethod and assigns it to the OptionTaxLotMethod field.
func (o *Preferences) SetOptionTaxLotMethod(v TaxLotMethod) {
	o.OptionTaxLotMethod = &v
}

func (o Preferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthTokenTimeout != nil {
		toSerialize["authTokenTimeout"] = o.AuthTokenTimeout
	}
	if o.DefaultAdvancedToolLaunch != nil {
		toSerialize["defaultAdvancedToolLaunch"] = o.DefaultAdvancedToolLaunch
	}
	if o.DefaultEquityOrderDuration != nil {
		toSerialize["defaultEquityOrderDuration"] = o.DefaultEquityOrderDuration
	}
	if o.DefaultEquityOrderLegInstruction != nil {
		toSerialize["defaultEquityOrderLegInstruction"] = o.DefaultEquityOrderLegInstruction
	}
	if o.DefaultEquityOrderMarketSession != nil {
		toSerialize["defaultEquityOrderMarketSession"] = o.DefaultEquityOrderMarketSession
	}
	if o.DefaultEquityOrderPriceLinkType != nil {
		toSerialize["defaultEquityOrderPriceLinkType"] = o.DefaultEquityOrderPriceLinkType
	}
	if o.DefaultEquityOrderType != nil {
		toSerialize["defaultEquityOrderType"] = o.DefaultEquityOrderType
	}
	if o.DefaultEquityQuantity != nil {
		toSerialize["defaultEquityQuantity"] = o.DefaultEquityQuantity
	}
	if o.DirectEquityRouting != nil {
		toSerialize["directEquityRouting"] = o.DirectEquityRouting
	}
	if o.DirectOptionsRouting != nil {
		toSerialize["directOptionsRouting"] = o.DirectOptionsRouting
	}
	if o.EquityTaxLotMethod != nil {
		toSerialize["equityTaxLotMethod"] = o.EquityTaxLotMethod
	}
	if o.ExpressTrading != nil {
		toSerialize["expressTrading"] = o.ExpressTrading
	}
	if o.MutualFundTaxLotMethod != nil {
		toSerialize["mutualFundTaxLotMethod"] = o.MutualFundTaxLotMethod
	}
	if o.OptionTaxLotMethod != nil {
		toSerialize["optionTaxLotMethod"] = o.OptionTaxLotMethod
	}
	return json.Marshal(toSerialize)
}

type NullablePreferences struct {
	value *Preferences
	isSet bool
}

func (v NullablePreferences) Get() *Preferences {
	return v.value
}

func (v *NullablePreferences) Set(val *Preferences) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferences) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferences(val *Preferences) *NullablePreferences {
	return &NullablePreferences{value: val, isSet: true}
}

func (v NullablePreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


