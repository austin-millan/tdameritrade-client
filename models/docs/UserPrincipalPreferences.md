# UserPrincipalPreferences

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
**DefaultEquityQuantity** | Pointer to **float32** |  | [optional] 
**DirectEquityRouting** | Pointer to **bool** |  | [optional] 
**DirectOptionsRouting** | Pointer to **bool** |  | [optional] 
**EquityTaxLotMethod** | Pointer to [**TaxLotMethod**](TaxLotMethod.md) |  | [optional] 
**ExpressTrading** | Pointer to **bool** |  | [optional] 
**MutualFundTaxLotMethod** | Pointer to [**TaxLotMethod**](TaxLotMethod.md) |  | [optional] 
**OptionTaxLotMethod** | Pointer to [**TaxLotMethod**](TaxLotMethod.md) |  | [optional] 

## Methods

### NewUserPrincipalPreferences

`func NewUserPrincipalPreferences() *UserPrincipalPreferences`

NewUserPrincipalPreferences instantiates a new UserPrincipalPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPrincipalPreferencesWithDefaults

`func NewUserPrincipalPreferencesWithDefaults() *UserPrincipalPreferences`

NewUserPrincipalPreferencesWithDefaults instantiates a new UserPrincipalPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTokenTimeout

`func (o *UserPrincipalPreferences) GetAuthTokenTimeout() string`

GetAuthTokenTimeout returns the AuthTokenTimeout field if non-nil, zero value otherwise.

### GetAuthTokenTimeoutOk

`func (o *UserPrincipalPreferences) GetAuthTokenTimeoutOk() (*string, bool)`

GetAuthTokenTimeoutOk returns a tuple with the AuthTokenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokenTimeout

`func (o *UserPrincipalPreferences) SetAuthTokenTimeout(v string)`

SetAuthTokenTimeout sets AuthTokenTimeout field to given value.

### HasAuthTokenTimeout

`func (o *UserPrincipalPreferences) HasAuthTokenTimeout() bool`

HasAuthTokenTimeout returns a boolean if a field has been set.

### GetDefaultAdvancedToolLaunch

`func (o *UserPrincipalPreferences) GetDefaultAdvancedToolLaunch() string`

GetDefaultAdvancedToolLaunch returns the DefaultAdvancedToolLaunch field if non-nil, zero value otherwise.

### GetDefaultAdvancedToolLaunchOk

`func (o *UserPrincipalPreferences) GetDefaultAdvancedToolLaunchOk() (*string, bool)`

GetDefaultAdvancedToolLaunchOk returns a tuple with the DefaultAdvancedToolLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAdvancedToolLaunch

`func (o *UserPrincipalPreferences) SetDefaultAdvancedToolLaunch(v string)`

SetDefaultAdvancedToolLaunch sets DefaultAdvancedToolLaunch field to given value.

### HasDefaultAdvancedToolLaunch

`func (o *UserPrincipalPreferences) HasDefaultAdvancedToolLaunch() bool`

HasDefaultAdvancedToolLaunch returns a boolean if a field has been set.

### GetDefaultEquityOrderDuration

`func (o *UserPrincipalPreferences) GetDefaultEquityOrderDuration() Duration`

GetDefaultEquityOrderDuration returns the DefaultEquityOrderDuration field if non-nil, zero value otherwise.

### GetDefaultEquityOrderDurationOk

`func (o *UserPrincipalPreferences) GetDefaultEquityOrderDurationOk() (*Duration, bool)`

GetDefaultEquityOrderDurationOk returns a tuple with the DefaultEquityOrderDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEquityOrderDuration

`func (o *UserPrincipalPreferences) SetDefaultEquityOrderDuration(v Duration)`

SetDefaultEquityOrderDuration sets DefaultEquityOrderDuration field to given value.

### HasDefaultEquityOrderDuration

`func (o *UserPrincipalPreferences) HasDefaultEquityOrderDuration() bool`

HasDefaultEquityOrderDuration returns a boolean if a field has been set.

### GetDefaultEquityOrderLegInstruction

`func (o *UserPrincipalPreferences) GetDefaultEquityOrderLegInstruction() string`

GetDefaultEquityOrderLegInstruction returns the DefaultEquityOrderLegInstruction field if non-nil, zero value otherwise.

### GetDefaultEquityOrderLegInstructionOk

`func (o *UserPrincipalPreferences) GetDefaultEquityOrderLegInstructionOk() (*string, bool)`

GetDefaultEquityOrderLegInstructionOk returns a tuple with the DefaultEquityOrderLegInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEquityOrderLegInstruction

`func (o *UserPrincipalPreferences) SetDefaultEquityOrderLegInstruction(v string)`

SetDefaultEquityOrderLegInstruction sets DefaultEquityOrderLegInstruction field to given value.

### HasDefaultEquityOrderLegInstruction

`func (o *UserPrincipalPreferences) HasDefaultEquityOrderLegInstruction() bool`

HasDefaultEquityOrderLegInstruction returns a boolean if a field has been set.

### GetDefaultEquityOrderMarketSession

`func (o *UserPrincipalPreferences) GetDefaultEquityOrderMarketSession() Session`

GetDefaultEquityOrderMarketSession returns the DefaultEquityOrderMarketSession field if non-nil, zero value otherwise.

### GetDefaultEquityOrderMarketSessionOk

`func (o *UserPrincipalPreferences) GetDefaultEquityOrderMarketSessionOk() (*Session, bool)`

GetDefaultEquityOrderMarketSessionOk returns a tuple with the DefaultEquityOrderMarketSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEquityOrderMarketSession

`func (o *UserPrincipalPreferences) SetDefaultEquityOrderMarketSession(v Session)`

SetDefaultEquityOrderMarketSession sets DefaultEquityOrderMarketSession field to given value.

### HasDefaultEquityOrderMarketSession

`func (o *UserPrincipalPreferences) HasDefaultEquityOrderMarketSession() bool`

HasDefaultEquityOrderMarketSession returns a boolean if a field has been set.

### GetDefaultEquityOrderPriceLinkType

`func (o *UserPrincipalPreferences) GetDefaultEquityOrderPriceLinkType() EquityPriceLinkType`

GetDefaultEquityOrderPriceLinkType returns the DefaultEquityOrderPriceLinkType field if non-nil, zero value otherwise.

### GetDefaultEquityOrderPriceLinkTypeOk

`func (o *UserPrincipalPreferences) GetDefaultEquityOrderPriceLinkTypeOk() (*EquityPriceLinkType, bool)`

GetDefaultEquityOrderPriceLinkTypeOk returns a tuple with the DefaultEquityOrderPriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEquityOrderPriceLinkType

`func (o *UserPrincipalPreferences) SetDefaultEquityOrderPriceLinkType(v EquityPriceLinkType)`

SetDefaultEquityOrderPriceLinkType sets DefaultEquityOrderPriceLinkType field to given value.

### HasDefaultEquityOrderPriceLinkType

`func (o *UserPrincipalPreferences) HasDefaultEquityOrderPriceLinkType() bool`

HasDefaultEquityOrderPriceLinkType returns a boolean if a field has been set.

### GetDefaultEquityOrderType

`func (o *UserPrincipalPreferences) GetDefaultEquityOrderType() OrderType`

GetDefaultEquityOrderType returns the DefaultEquityOrderType field if non-nil, zero value otherwise.

### GetDefaultEquityOrderTypeOk

`func (o *UserPrincipalPreferences) GetDefaultEquityOrderTypeOk() (*OrderType, bool)`

GetDefaultEquityOrderTypeOk returns a tuple with the DefaultEquityOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEquityOrderType

`func (o *UserPrincipalPreferences) SetDefaultEquityOrderType(v OrderType)`

SetDefaultEquityOrderType sets DefaultEquityOrderType field to given value.

### HasDefaultEquityOrderType

`func (o *UserPrincipalPreferences) HasDefaultEquityOrderType() bool`

HasDefaultEquityOrderType returns a boolean if a field has been set.

### GetDefaultEquityQuantity

`func (o *UserPrincipalPreferences) GetDefaultEquityQuantity() float32`

GetDefaultEquityQuantity returns the DefaultEquityQuantity field if non-nil, zero value otherwise.

### GetDefaultEquityQuantityOk

`func (o *UserPrincipalPreferences) GetDefaultEquityQuantityOk() (*float32, bool)`

GetDefaultEquityQuantityOk returns a tuple with the DefaultEquityQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEquityQuantity

`func (o *UserPrincipalPreferences) SetDefaultEquityQuantity(v float32)`

SetDefaultEquityQuantity sets DefaultEquityQuantity field to given value.

### HasDefaultEquityQuantity

`func (o *UserPrincipalPreferences) HasDefaultEquityQuantity() bool`

HasDefaultEquityQuantity returns a boolean if a field has been set.

### GetDirectEquityRouting

`func (o *UserPrincipalPreferences) GetDirectEquityRouting() bool`

GetDirectEquityRouting returns the DirectEquityRouting field if non-nil, zero value otherwise.

### GetDirectEquityRoutingOk

`func (o *UserPrincipalPreferences) GetDirectEquityRoutingOk() (*bool, bool)`

GetDirectEquityRoutingOk returns a tuple with the DirectEquityRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectEquityRouting

`func (o *UserPrincipalPreferences) SetDirectEquityRouting(v bool)`

SetDirectEquityRouting sets DirectEquityRouting field to given value.

### HasDirectEquityRouting

`func (o *UserPrincipalPreferences) HasDirectEquityRouting() bool`

HasDirectEquityRouting returns a boolean if a field has been set.

### GetDirectOptionsRouting

`func (o *UserPrincipalPreferences) GetDirectOptionsRouting() bool`

GetDirectOptionsRouting returns the DirectOptionsRouting field if non-nil, zero value otherwise.

### GetDirectOptionsRoutingOk

`func (o *UserPrincipalPreferences) GetDirectOptionsRoutingOk() (*bool, bool)`

GetDirectOptionsRoutingOk returns a tuple with the DirectOptionsRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectOptionsRouting

`func (o *UserPrincipalPreferences) SetDirectOptionsRouting(v bool)`

SetDirectOptionsRouting sets DirectOptionsRouting field to given value.

### HasDirectOptionsRouting

`func (o *UserPrincipalPreferences) HasDirectOptionsRouting() bool`

HasDirectOptionsRouting returns a boolean if a field has been set.

### GetEquityTaxLotMethod

`func (o *UserPrincipalPreferences) GetEquityTaxLotMethod() TaxLotMethod`

GetEquityTaxLotMethod returns the EquityTaxLotMethod field if non-nil, zero value otherwise.

### GetEquityTaxLotMethodOk

`func (o *UserPrincipalPreferences) GetEquityTaxLotMethodOk() (*TaxLotMethod, bool)`

GetEquityTaxLotMethodOk returns a tuple with the EquityTaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityTaxLotMethod

`func (o *UserPrincipalPreferences) SetEquityTaxLotMethod(v TaxLotMethod)`

SetEquityTaxLotMethod sets EquityTaxLotMethod field to given value.

### HasEquityTaxLotMethod

`func (o *UserPrincipalPreferences) HasEquityTaxLotMethod() bool`

HasEquityTaxLotMethod returns a boolean if a field has been set.

### GetExpressTrading

`func (o *UserPrincipalPreferences) GetExpressTrading() bool`

GetExpressTrading returns the ExpressTrading field if non-nil, zero value otherwise.

### GetExpressTradingOk

`func (o *UserPrincipalPreferences) GetExpressTradingOk() (*bool, bool)`

GetExpressTradingOk returns a tuple with the ExpressTrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressTrading

`func (o *UserPrincipalPreferences) SetExpressTrading(v bool)`

SetExpressTrading sets ExpressTrading field to given value.

### HasExpressTrading

`func (o *UserPrincipalPreferences) HasExpressTrading() bool`

HasExpressTrading returns a boolean if a field has been set.

### GetMutualFundTaxLotMethod

`func (o *UserPrincipalPreferences) GetMutualFundTaxLotMethod() TaxLotMethod`

GetMutualFundTaxLotMethod returns the MutualFundTaxLotMethod field if non-nil, zero value otherwise.

### GetMutualFundTaxLotMethodOk

`func (o *UserPrincipalPreferences) GetMutualFundTaxLotMethodOk() (*TaxLotMethod, bool)`

GetMutualFundTaxLotMethodOk returns a tuple with the MutualFundTaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualFundTaxLotMethod

`func (o *UserPrincipalPreferences) SetMutualFundTaxLotMethod(v TaxLotMethod)`

SetMutualFundTaxLotMethod sets MutualFundTaxLotMethod field to given value.

### HasMutualFundTaxLotMethod

`func (o *UserPrincipalPreferences) HasMutualFundTaxLotMethod() bool`

HasMutualFundTaxLotMethod returns a boolean if a field has been set.

### GetOptionTaxLotMethod

`func (o *UserPrincipalPreferences) GetOptionTaxLotMethod() TaxLotMethod`

GetOptionTaxLotMethod returns the OptionTaxLotMethod field if non-nil, zero value otherwise.

### GetOptionTaxLotMethodOk

`func (o *UserPrincipalPreferences) GetOptionTaxLotMethodOk() (*TaxLotMethod, bool)`

GetOptionTaxLotMethodOk returns a tuple with the OptionTaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTaxLotMethod

`func (o *UserPrincipalPreferences) SetOptionTaxLotMethod(v TaxLotMethod)`

SetOptionTaxLotMethod sets OptionTaxLotMethod field to given value.

### HasOptionTaxLotMethod

`func (o *UserPrincipalPreferences) HasOptionTaxLotMethod() bool`

HasOptionTaxLotMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


