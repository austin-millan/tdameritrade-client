# UpdatePreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthTokenTimeout** | Pointer to **string** |  | [optional] 
**DefaultAdvancedToolLaunch** | Pointer to **string** |  | [optional] 
**DefaultEquityOrderDuration** | Pointer to [**Duration**](Duration.md) |  | [optional] 
**DefaultEquityOrderLegInstruction** | Pointer to **string** |  | [optional] 
**DefaultEquityOrderMarketSession** | Pointer to [**Session**](Session.md) |  | [optional] 
**DefaultEquityOrderPriceLinkType** | Pointer to [**EquityPriceLinkType**](EquityPriceLinkType.md) |  | [optional] 
**DefaultEquityOrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**DefaultEquityQuantity** | Pointer to **int32** |  | [optional] 
**DirectEquityRouting** | Pointer to **bool** |  | [optional] 
**DirectOptionsRouting** | Pointer to **bool** |  | [optional] 
**EquityTaxLotMethod** | Pointer to [**TaxLotMethod**](TaxLotMethod.md) |  | [optional] 
**ExpressTrading** | Pointer to **bool** |  | [optional] 
**MutualFundTaxLotMethod** | Pointer to [**TaxLotMethod**](TaxLotMethod.md) |  | [optional] 
**OptionTaxLotMethod** | Pointer to [**TaxLotMethod**](TaxLotMethod.md) |  | [optional] 

## Methods

### NewUpdatePreferences

`func NewUpdatePreferences() *UpdatePreferences`

NewUpdatePreferences instantiates a new UpdatePreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePreferencesWithDefaults

`func NewUpdatePreferencesWithDefaults() *UpdatePreferences`

NewUpdatePreferencesWithDefaults instantiates a new UpdatePreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTokenTimeout

`func (o *UpdatePreferences) GetAuthTokenTimeout() string`

GetAuthTokenTimeout returns the AuthTokenTimeout field if non-nil, zero value otherwise.

### GetAuthTokenTimeoutOk

`func (o *UpdatePreferences) GetAuthTokenTimeoutOk() (*string, bool)`

GetAuthTokenTimeoutOk returns a tuple with the AuthTokenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokenTimeout

`func (o *UpdatePreferences) SetAuthTokenTimeout(v string)`

SetAuthTokenTimeout sets AuthTokenTimeout field to given value.

### HasAuthTokenTimeout

`func (o *UpdatePreferences) HasAuthTokenTimeout() bool`

HasAuthTokenTimeout returns a boolean if a field has been set.

### GetDefaultAdvancedToolLaunch

`func (o *UpdatePreferences) GetDefaultAdvancedToolLaunch() string`

GetDefaultAdvancedToolLaunch returns the DefaultAdvancedToolLaunch field if non-nil, zero value otherwise.

### GetDefaultAdvancedToolLaunchOk

`func (o *UpdatePreferences) GetDefaultAdvancedToolLaunchOk() (*string, bool)`

GetDefaultAdvancedToolLaunchOk returns a tuple with the DefaultAdvancedToolLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAdvancedToolLaunch

`func (o *UpdatePreferences) SetDefaultAdvancedToolLaunch(v string)`

SetDefaultAdvancedToolLaunch sets DefaultAdvancedToolLaunch field to given value.

### HasDefaultAdvancedToolLaunch

`func (o *UpdatePreferences) HasDefaultAdvancedToolLaunch() bool`

HasDefaultAdvancedToolLaunch returns a boolean if a field has been set.

### GetDefaultEquityOrderDuration

`func (o *UpdatePreferences) GetDefaultEquityOrderDuration() Duration`

GetDefaultEquityOrderDuration returns the DefaultEquityOrderDuration field if non-nil, zero value otherwise.

### GetDefaultEquityOrderDurationOk

`func (o *UpdatePreferences) GetDefaultEquityOrderDurationOk() (*Duration, bool)`

GetDefaultEquityOrderDurationOk returns a tuple with the DefaultEquityOrderDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEquityOrderDuration

`func (o *UpdatePreferences) SetDefaultEquityOrderDuration(v Duration)`

SetDefaultEquityOrderDuration sets DefaultEquityOrderDuration field to given value.

### HasDefaultEquityOrderDuration

`func (o *UpdatePreferences) HasDefaultEquityOrderDuration() bool`

HasDefaultEquityOrderDuration returns a boolean if a field has been set.

### GetDefaultEquityOrderLegInstruction

`func (o *UpdatePreferences) GetDefaultEquityOrderLegInstruction() string`

GetDefaultEquityOrderLegInstruction returns the DefaultEquityOrderLegInstruction field if non-nil, zero value otherwise.

### GetDefaultEquityOrderLegInstructionOk

`func (o *UpdatePreferences) GetDefaultEquityOrderLegInstructionOk() (*string, bool)`

GetDefaultEquityOrderLegInstructionOk returns a tuple with the DefaultEquityOrderLegInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEquityOrderLegInstruction

`func (o *UpdatePreferences) SetDefaultEquityOrderLegInstruction(v string)`

SetDefaultEquityOrderLegInstruction sets DefaultEquityOrderLegInstruction field to given value.

### HasDefaultEquityOrderLegInstruction

`func (o *UpdatePreferences) HasDefaultEquityOrderLegInstruction() bool`

HasDefaultEquityOrderLegInstruction returns a boolean if a field has been set.

### GetDefaultEquityOrderMarketSession

`func (o *UpdatePreferences) GetDefaultEquityOrderMarketSession() Session`

GetDefaultEquityOrderMarketSession returns the DefaultEquityOrderMarketSession field if non-nil, zero value otherwise.

### GetDefaultEquityOrderMarketSessionOk

`func (o *UpdatePreferences) GetDefaultEquityOrderMarketSessionOk() (*Session, bool)`

GetDefaultEquityOrderMarketSessionOk returns a tuple with the DefaultEquityOrderMarketSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEquityOrderMarketSession

`func (o *UpdatePreferences) SetDefaultEquityOrderMarketSession(v Session)`

SetDefaultEquityOrderMarketSession sets DefaultEquityOrderMarketSession field to given value.

### HasDefaultEquityOrderMarketSession

`func (o *UpdatePreferences) HasDefaultEquityOrderMarketSession() bool`

HasDefaultEquityOrderMarketSession returns a boolean if a field has been set.

### GetDefaultEquityOrderPriceLinkType

`func (o *UpdatePreferences) GetDefaultEquityOrderPriceLinkType() EquityPriceLinkType`

GetDefaultEquityOrderPriceLinkType returns the DefaultEquityOrderPriceLinkType field if non-nil, zero value otherwise.

### GetDefaultEquityOrderPriceLinkTypeOk

`func (o *UpdatePreferences) GetDefaultEquityOrderPriceLinkTypeOk() (*EquityPriceLinkType, bool)`

GetDefaultEquityOrderPriceLinkTypeOk returns a tuple with the DefaultEquityOrderPriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEquityOrderPriceLinkType

`func (o *UpdatePreferences) SetDefaultEquityOrderPriceLinkType(v EquityPriceLinkType)`

SetDefaultEquityOrderPriceLinkType sets DefaultEquityOrderPriceLinkType field to given value.

### HasDefaultEquityOrderPriceLinkType

`func (o *UpdatePreferences) HasDefaultEquityOrderPriceLinkType() bool`

HasDefaultEquityOrderPriceLinkType returns a boolean if a field has been set.

### GetDefaultEquityOrderType

`func (o *UpdatePreferences) GetDefaultEquityOrderType() OrderType`

GetDefaultEquityOrderType returns the DefaultEquityOrderType field if non-nil, zero value otherwise.

### GetDefaultEquityOrderTypeOk

`func (o *UpdatePreferences) GetDefaultEquityOrderTypeOk() (*OrderType, bool)`

GetDefaultEquityOrderTypeOk returns a tuple with the DefaultEquityOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEquityOrderType

`func (o *UpdatePreferences) SetDefaultEquityOrderType(v OrderType)`

SetDefaultEquityOrderType sets DefaultEquityOrderType field to given value.

### HasDefaultEquityOrderType

`func (o *UpdatePreferences) HasDefaultEquityOrderType() bool`

HasDefaultEquityOrderType returns a boolean if a field has been set.

### GetDefaultEquityQuantity

`func (o *UpdatePreferences) GetDefaultEquityQuantity() int32`

GetDefaultEquityQuantity returns the DefaultEquityQuantity field if non-nil, zero value otherwise.

### GetDefaultEquityQuantityOk

`func (o *UpdatePreferences) GetDefaultEquityQuantityOk() (*int32, bool)`

GetDefaultEquityQuantityOk returns a tuple with the DefaultEquityQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEquityQuantity

`func (o *UpdatePreferences) SetDefaultEquityQuantity(v int32)`

SetDefaultEquityQuantity sets DefaultEquityQuantity field to given value.

### HasDefaultEquityQuantity

`func (o *UpdatePreferences) HasDefaultEquityQuantity() bool`

HasDefaultEquityQuantity returns a boolean if a field has been set.

### GetDirectEquityRouting

`func (o *UpdatePreferences) GetDirectEquityRouting() bool`

GetDirectEquityRouting returns the DirectEquityRouting field if non-nil, zero value otherwise.

### GetDirectEquityRoutingOk

`func (o *UpdatePreferences) GetDirectEquityRoutingOk() (*bool, bool)`

GetDirectEquityRoutingOk returns a tuple with the DirectEquityRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectEquityRouting

`func (o *UpdatePreferences) SetDirectEquityRouting(v bool)`

SetDirectEquityRouting sets DirectEquityRouting field to given value.

### HasDirectEquityRouting

`func (o *UpdatePreferences) HasDirectEquityRouting() bool`

HasDirectEquityRouting returns a boolean if a field has been set.

### GetDirectOptionsRouting

`func (o *UpdatePreferences) GetDirectOptionsRouting() bool`

GetDirectOptionsRouting returns the DirectOptionsRouting field if non-nil, zero value otherwise.

### GetDirectOptionsRoutingOk

`func (o *UpdatePreferences) GetDirectOptionsRoutingOk() (*bool, bool)`

GetDirectOptionsRoutingOk returns a tuple with the DirectOptionsRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectOptionsRouting

`func (o *UpdatePreferences) SetDirectOptionsRouting(v bool)`

SetDirectOptionsRouting sets DirectOptionsRouting field to given value.

### HasDirectOptionsRouting

`func (o *UpdatePreferences) HasDirectOptionsRouting() bool`

HasDirectOptionsRouting returns a boolean if a field has been set.

### GetEquityTaxLotMethod

`func (o *UpdatePreferences) GetEquityTaxLotMethod() TaxLotMethod`

GetEquityTaxLotMethod returns the EquityTaxLotMethod field if non-nil, zero value otherwise.

### GetEquityTaxLotMethodOk

`func (o *UpdatePreferences) GetEquityTaxLotMethodOk() (*TaxLotMethod, bool)`

GetEquityTaxLotMethodOk returns a tuple with the EquityTaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityTaxLotMethod

`func (o *UpdatePreferences) SetEquityTaxLotMethod(v TaxLotMethod)`

SetEquityTaxLotMethod sets EquityTaxLotMethod field to given value.

### HasEquityTaxLotMethod

`func (o *UpdatePreferences) HasEquityTaxLotMethod() bool`

HasEquityTaxLotMethod returns a boolean if a field has been set.

### GetExpressTrading

`func (o *UpdatePreferences) GetExpressTrading() bool`

GetExpressTrading returns the ExpressTrading field if non-nil, zero value otherwise.

### GetExpressTradingOk

`func (o *UpdatePreferences) GetExpressTradingOk() (*bool, bool)`

GetExpressTradingOk returns a tuple with the ExpressTrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressTrading

`func (o *UpdatePreferences) SetExpressTrading(v bool)`

SetExpressTrading sets ExpressTrading field to given value.

### HasExpressTrading

`func (o *UpdatePreferences) HasExpressTrading() bool`

HasExpressTrading returns a boolean if a field has been set.

### GetMutualFundTaxLotMethod

`func (o *UpdatePreferences) GetMutualFundTaxLotMethod() TaxLotMethod`

GetMutualFundTaxLotMethod returns the MutualFundTaxLotMethod field if non-nil, zero value otherwise.

### GetMutualFundTaxLotMethodOk

`func (o *UpdatePreferences) GetMutualFundTaxLotMethodOk() (*TaxLotMethod, bool)`

GetMutualFundTaxLotMethodOk returns a tuple with the MutualFundTaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualFundTaxLotMethod

`func (o *UpdatePreferences) SetMutualFundTaxLotMethod(v TaxLotMethod)`

SetMutualFundTaxLotMethod sets MutualFundTaxLotMethod field to given value.

### HasMutualFundTaxLotMethod

`func (o *UpdatePreferences) HasMutualFundTaxLotMethod() bool`

HasMutualFundTaxLotMethod returns a boolean if a field has been set.

### GetOptionTaxLotMethod

`func (o *UpdatePreferences) GetOptionTaxLotMethod() TaxLotMethod`

GetOptionTaxLotMethod returns the OptionTaxLotMethod field if non-nil, zero value otherwise.

### GetOptionTaxLotMethodOk

`func (o *UpdatePreferences) GetOptionTaxLotMethodOk() (*TaxLotMethod, bool)`

GetOptionTaxLotMethodOk returns a tuple with the OptionTaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTaxLotMethod

`func (o *UpdatePreferences) SetOptionTaxLotMethod(v TaxLotMethod)`

SetOptionTaxLotMethod sets OptionTaxLotMethod field to given value.

### HasOptionTaxLotMethod

`func (o *UpdatePreferences) HasOptionTaxLotMethod() bool`

HasOptionTaxLotMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


