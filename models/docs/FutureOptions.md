# FutureOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskPriceInDouble** | Pointer to **float32** |  | [optional] 
**BidPriceInDouble** | Pointer to **float32** |  | [optional] 
**ClosePriceInDouble** | Pointer to **float32** |  | [optional] 
**ContractType** | Pointer to [**PutOrCall**](PutOrCall.md) |  | [optional] 
**DeltaInDouble** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Digits** | Pointer to **float32** |  | [optional] 
**ExchangeName** | Pointer to **string** |  | [optional] 
**ExerciseType** | Pointer to **string** |  | [optional] 
**ExpirationType** | Pointer to **string** |  | [optional] 
**FutureExpirationDate** | Pointer to **float32** |  | [optional] 
**FutureIsActive** | Pointer to **bool** |  | [optional] 
**FutureIsTradable** | Pointer to **bool** |  | [optional] 
**FuturePercentChange** | Pointer to **float32** |  | [optional] 
**FutureTradingHours** | Pointer to **string** |  | [optional] 
**GammaInDouble** | Pointer to **float32** |  | [optional] 
**HighPriceInDouble** | Pointer to **float32** |  | [optional] 
**InTheMoney** | Pointer to **bool** |  | [optional] 
**LastPriceInDouble** | Pointer to **float32** |  | [optional] 
**LowPriceInDouble** | Pointer to **float32** |  | [optional] 
**Mark** | Pointer to **float32** |  | [optional] 
**MoneyIntrinsicValueInDouble** | Pointer to **float32** |  | [optional] 
**MultiplierInDouble** | Pointer to **float32** |  | [optional] 
**NetChangeInDouble** | Pointer to **float32** |  | [optional] 
**OpenInterest** | Pointer to **float32** |  | [optional] 
**OpenPriceInDouble** | Pointer to **float32** |  | [optional] 
**RhoInDouble** | Pointer to **float32** |  | [optional] 
**SecurityStatus** | Pointer to **string** |  | [optional] 
**StrikePriceInDouble** | Pointer to **float32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**ThetaInDouble** | Pointer to **float32** |  | [optional] 
**Tick** | Pointer to **float32** |  | [optional] 
**TickAmount** | Pointer to **float32** |  | [optional] 
**TimeValueInDouble** | Pointer to **float32** |  | [optional] 
**Underlying** | Pointer to **string** |  | [optional] 
**VegaInDouble** | Pointer to **float32** |  | [optional] 
**Volatility** | Pointer to **float32** |  | [optional] 

## Methods

### NewFutureOptions

`func NewFutureOptions() *FutureOptions`

NewFutureOptions instantiates a new FutureOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFutureOptionsWithDefaults

`func NewFutureOptionsWithDefaults() *FutureOptions`

NewFutureOptionsWithDefaults instantiates a new FutureOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPriceInDouble

`func (o *FutureOptions) GetAskPriceInDouble() float32`

GetAskPriceInDouble returns the AskPriceInDouble field if non-nil, zero value otherwise.

### GetAskPriceInDoubleOk

`func (o *FutureOptions) GetAskPriceInDoubleOk() (*float32, bool)`

GetAskPriceInDoubleOk returns a tuple with the AskPriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPriceInDouble

`func (o *FutureOptions) SetAskPriceInDouble(v float32)`

SetAskPriceInDouble sets AskPriceInDouble field to given value.

### HasAskPriceInDouble

`func (o *FutureOptions) HasAskPriceInDouble() bool`

HasAskPriceInDouble returns a boolean if a field has been set.

### GetBidPriceInDouble

`func (o *FutureOptions) GetBidPriceInDouble() float32`

GetBidPriceInDouble returns the BidPriceInDouble field if non-nil, zero value otherwise.

### GetBidPriceInDoubleOk

`func (o *FutureOptions) GetBidPriceInDoubleOk() (*float32, bool)`

GetBidPriceInDoubleOk returns a tuple with the BidPriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPriceInDouble

`func (o *FutureOptions) SetBidPriceInDouble(v float32)`

SetBidPriceInDouble sets BidPriceInDouble field to given value.

### HasBidPriceInDouble

`func (o *FutureOptions) HasBidPriceInDouble() bool`

HasBidPriceInDouble returns a boolean if a field has been set.

### GetClosePriceInDouble

`func (o *FutureOptions) GetClosePriceInDouble() float32`

GetClosePriceInDouble returns the ClosePriceInDouble field if non-nil, zero value otherwise.

### GetClosePriceInDoubleOk

`func (o *FutureOptions) GetClosePriceInDoubleOk() (*float32, bool)`

GetClosePriceInDoubleOk returns a tuple with the ClosePriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePriceInDouble

`func (o *FutureOptions) SetClosePriceInDouble(v float32)`

SetClosePriceInDouble sets ClosePriceInDouble field to given value.

### HasClosePriceInDouble

`func (o *FutureOptions) HasClosePriceInDouble() bool`

HasClosePriceInDouble returns a boolean if a field has been set.

### GetContractType

`func (o *FutureOptions) GetContractType() PutOrCall`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *FutureOptions) GetContractTypeOk() (*PutOrCall, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *FutureOptions) SetContractType(v PutOrCall)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *FutureOptions) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetDeltaInDouble

`func (o *FutureOptions) GetDeltaInDouble() float32`

GetDeltaInDouble returns the DeltaInDouble field if non-nil, zero value otherwise.

### GetDeltaInDoubleOk

`func (o *FutureOptions) GetDeltaInDoubleOk() (*float32, bool)`

GetDeltaInDoubleOk returns a tuple with the DeltaInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaInDouble

`func (o *FutureOptions) SetDeltaInDouble(v float32)`

SetDeltaInDouble sets DeltaInDouble field to given value.

### HasDeltaInDouble

`func (o *FutureOptions) HasDeltaInDouble() bool`

HasDeltaInDouble returns a boolean if a field has been set.

### GetDescription

`func (o *FutureOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FutureOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FutureOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FutureOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDigits

`func (o *FutureOptions) GetDigits() float32`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *FutureOptions) GetDigitsOk() (*float32, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *FutureOptions) SetDigits(v float32)`

SetDigits sets Digits field to given value.

### HasDigits

`func (o *FutureOptions) HasDigits() bool`

HasDigits returns a boolean if a field has been set.

### GetExchangeName

`func (o *FutureOptions) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *FutureOptions) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *FutureOptions) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *FutureOptions) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetExerciseType

`func (o *FutureOptions) GetExerciseType() string`

GetExerciseType returns the ExerciseType field if non-nil, zero value otherwise.

### GetExerciseTypeOk

`func (o *FutureOptions) GetExerciseTypeOk() (*string, bool)`

GetExerciseTypeOk returns a tuple with the ExerciseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExerciseType

`func (o *FutureOptions) SetExerciseType(v string)`

SetExerciseType sets ExerciseType field to given value.

### HasExerciseType

`func (o *FutureOptions) HasExerciseType() bool`

HasExerciseType returns a boolean if a field has been set.

### GetExpirationType

`func (o *FutureOptions) GetExpirationType() string`

GetExpirationType returns the ExpirationType field if non-nil, zero value otherwise.

### GetExpirationTypeOk

`func (o *FutureOptions) GetExpirationTypeOk() (*string, bool)`

GetExpirationTypeOk returns a tuple with the ExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationType

`func (o *FutureOptions) SetExpirationType(v string)`

SetExpirationType sets ExpirationType field to given value.

### HasExpirationType

`func (o *FutureOptions) HasExpirationType() bool`

HasExpirationType returns a boolean if a field has been set.

### GetFutureExpirationDate

`func (o *FutureOptions) GetFutureExpirationDate() float32`

GetFutureExpirationDate returns the FutureExpirationDate field if non-nil, zero value otherwise.

### GetFutureExpirationDateOk

`func (o *FutureOptions) GetFutureExpirationDateOk() (*float32, bool)`

GetFutureExpirationDateOk returns a tuple with the FutureExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureExpirationDate

`func (o *FutureOptions) SetFutureExpirationDate(v float32)`

SetFutureExpirationDate sets FutureExpirationDate field to given value.

### HasFutureExpirationDate

`func (o *FutureOptions) HasFutureExpirationDate() bool`

HasFutureExpirationDate returns a boolean if a field has been set.

### GetFutureIsActive

`func (o *FutureOptions) GetFutureIsActive() bool`

GetFutureIsActive returns the FutureIsActive field if non-nil, zero value otherwise.

### GetFutureIsActiveOk

`func (o *FutureOptions) GetFutureIsActiveOk() (*bool, bool)`

GetFutureIsActiveOk returns a tuple with the FutureIsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureIsActive

`func (o *FutureOptions) SetFutureIsActive(v bool)`

SetFutureIsActive sets FutureIsActive field to given value.

### HasFutureIsActive

`func (o *FutureOptions) HasFutureIsActive() bool`

HasFutureIsActive returns a boolean if a field has been set.

### GetFutureIsTradable

`func (o *FutureOptions) GetFutureIsTradable() bool`

GetFutureIsTradable returns the FutureIsTradable field if non-nil, zero value otherwise.

### GetFutureIsTradableOk

`func (o *FutureOptions) GetFutureIsTradableOk() (*bool, bool)`

GetFutureIsTradableOk returns a tuple with the FutureIsTradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureIsTradable

`func (o *FutureOptions) SetFutureIsTradable(v bool)`

SetFutureIsTradable sets FutureIsTradable field to given value.

### HasFutureIsTradable

`func (o *FutureOptions) HasFutureIsTradable() bool`

HasFutureIsTradable returns a boolean if a field has been set.

### GetFuturePercentChange

`func (o *FutureOptions) GetFuturePercentChange() float32`

GetFuturePercentChange returns the FuturePercentChange field if non-nil, zero value otherwise.

### GetFuturePercentChangeOk

`func (o *FutureOptions) GetFuturePercentChangeOk() (*float32, bool)`

GetFuturePercentChangeOk returns a tuple with the FuturePercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuturePercentChange

`func (o *FutureOptions) SetFuturePercentChange(v float32)`

SetFuturePercentChange sets FuturePercentChange field to given value.

### HasFuturePercentChange

`func (o *FutureOptions) HasFuturePercentChange() bool`

HasFuturePercentChange returns a boolean if a field has been set.

### GetFutureTradingHours

`func (o *FutureOptions) GetFutureTradingHours() string`

GetFutureTradingHours returns the FutureTradingHours field if non-nil, zero value otherwise.

### GetFutureTradingHoursOk

`func (o *FutureOptions) GetFutureTradingHoursOk() (*string, bool)`

GetFutureTradingHoursOk returns a tuple with the FutureTradingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureTradingHours

`func (o *FutureOptions) SetFutureTradingHours(v string)`

SetFutureTradingHours sets FutureTradingHours field to given value.

### HasFutureTradingHours

`func (o *FutureOptions) HasFutureTradingHours() bool`

HasFutureTradingHours returns a boolean if a field has been set.

### GetGammaInDouble

`func (o *FutureOptions) GetGammaInDouble() float32`

GetGammaInDouble returns the GammaInDouble field if non-nil, zero value otherwise.

### GetGammaInDoubleOk

`func (o *FutureOptions) GetGammaInDoubleOk() (*float32, bool)`

GetGammaInDoubleOk returns a tuple with the GammaInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGammaInDouble

`func (o *FutureOptions) SetGammaInDouble(v float32)`

SetGammaInDouble sets GammaInDouble field to given value.

### HasGammaInDouble

`func (o *FutureOptions) HasGammaInDouble() bool`

HasGammaInDouble returns a boolean if a field has been set.

### GetHighPriceInDouble

`func (o *FutureOptions) GetHighPriceInDouble() float32`

GetHighPriceInDouble returns the HighPriceInDouble field if non-nil, zero value otherwise.

### GetHighPriceInDoubleOk

`func (o *FutureOptions) GetHighPriceInDoubleOk() (*float32, bool)`

GetHighPriceInDoubleOk returns a tuple with the HighPriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPriceInDouble

`func (o *FutureOptions) SetHighPriceInDouble(v float32)`

SetHighPriceInDouble sets HighPriceInDouble field to given value.

### HasHighPriceInDouble

`func (o *FutureOptions) HasHighPriceInDouble() bool`

HasHighPriceInDouble returns a boolean if a field has been set.

### GetInTheMoney

`func (o *FutureOptions) GetInTheMoney() bool`

GetInTheMoney returns the InTheMoney field if non-nil, zero value otherwise.

### GetInTheMoneyOk

`func (o *FutureOptions) GetInTheMoneyOk() (*bool, bool)`

GetInTheMoneyOk returns a tuple with the InTheMoney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTheMoney

`func (o *FutureOptions) SetInTheMoney(v bool)`

SetInTheMoney sets InTheMoney field to given value.

### HasInTheMoney

`func (o *FutureOptions) HasInTheMoney() bool`

HasInTheMoney returns a boolean if a field has been set.

### GetLastPriceInDouble

`func (o *FutureOptions) GetLastPriceInDouble() float32`

GetLastPriceInDouble returns the LastPriceInDouble field if non-nil, zero value otherwise.

### GetLastPriceInDoubleOk

`func (o *FutureOptions) GetLastPriceInDoubleOk() (*float32, bool)`

GetLastPriceInDoubleOk returns a tuple with the LastPriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPriceInDouble

`func (o *FutureOptions) SetLastPriceInDouble(v float32)`

SetLastPriceInDouble sets LastPriceInDouble field to given value.

### HasLastPriceInDouble

`func (o *FutureOptions) HasLastPriceInDouble() bool`

HasLastPriceInDouble returns a boolean if a field has been set.

### GetLowPriceInDouble

`func (o *FutureOptions) GetLowPriceInDouble() float32`

GetLowPriceInDouble returns the LowPriceInDouble field if non-nil, zero value otherwise.

### GetLowPriceInDoubleOk

`func (o *FutureOptions) GetLowPriceInDoubleOk() (*float32, bool)`

GetLowPriceInDoubleOk returns a tuple with the LowPriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPriceInDouble

`func (o *FutureOptions) SetLowPriceInDouble(v float32)`

SetLowPriceInDouble sets LowPriceInDouble field to given value.

### HasLowPriceInDouble

`func (o *FutureOptions) HasLowPriceInDouble() bool`

HasLowPriceInDouble returns a boolean if a field has been set.

### GetMark

`func (o *FutureOptions) GetMark() float32`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *FutureOptions) GetMarkOk() (*float32, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *FutureOptions) SetMark(v float32)`

SetMark sets Mark field to given value.

### HasMark

`func (o *FutureOptions) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetMoneyIntrinsicValueInDouble

`func (o *FutureOptions) GetMoneyIntrinsicValueInDouble() float32`

GetMoneyIntrinsicValueInDouble returns the MoneyIntrinsicValueInDouble field if non-nil, zero value otherwise.

### GetMoneyIntrinsicValueInDoubleOk

`func (o *FutureOptions) GetMoneyIntrinsicValueInDoubleOk() (*float32, bool)`

GetMoneyIntrinsicValueInDoubleOk returns a tuple with the MoneyIntrinsicValueInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyIntrinsicValueInDouble

`func (o *FutureOptions) SetMoneyIntrinsicValueInDouble(v float32)`

SetMoneyIntrinsicValueInDouble sets MoneyIntrinsicValueInDouble field to given value.

### HasMoneyIntrinsicValueInDouble

`func (o *FutureOptions) HasMoneyIntrinsicValueInDouble() bool`

HasMoneyIntrinsicValueInDouble returns a boolean if a field has been set.

### GetMultiplierInDouble

`func (o *FutureOptions) GetMultiplierInDouble() float32`

GetMultiplierInDouble returns the MultiplierInDouble field if non-nil, zero value otherwise.

### GetMultiplierInDoubleOk

`func (o *FutureOptions) GetMultiplierInDoubleOk() (*float32, bool)`

GetMultiplierInDoubleOk returns a tuple with the MultiplierInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierInDouble

`func (o *FutureOptions) SetMultiplierInDouble(v float32)`

SetMultiplierInDouble sets MultiplierInDouble field to given value.

### HasMultiplierInDouble

`func (o *FutureOptions) HasMultiplierInDouble() bool`

HasMultiplierInDouble returns a boolean if a field has been set.

### GetNetChangeInDouble

`func (o *FutureOptions) GetNetChangeInDouble() float32`

GetNetChangeInDouble returns the NetChangeInDouble field if non-nil, zero value otherwise.

### GetNetChangeInDoubleOk

`func (o *FutureOptions) GetNetChangeInDoubleOk() (*float32, bool)`

GetNetChangeInDoubleOk returns a tuple with the NetChangeInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChangeInDouble

`func (o *FutureOptions) SetNetChangeInDouble(v float32)`

SetNetChangeInDouble sets NetChangeInDouble field to given value.

### HasNetChangeInDouble

`func (o *FutureOptions) HasNetChangeInDouble() bool`

HasNetChangeInDouble returns a boolean if a field has been set.

### GetOpenInterest

`func (o *FutureOptions) GetOpenInterest() float32`

GetOpenInterest returns the OpenInterest field if non-nil, zero value otherwise.

### GetOpenInterestOk

`func (o *FutureOptions) GetOpenInterestOk() (*float32, bool)`

GetOpenInterestOk returns a tuple with the OpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterest

`func (o *FutureOptions) SetOpenInterest(v float32)`

SetOpenInterest sets OpenInterest field to given value.

### HasOpenInterest

`func (o *FutureOptions) HasOpenInterest() bool`

HasOpenInterest returns a boolean if a field has been set.

### GetOpenPriceInDouble

`func (o *FutureOptions) GetOpenPriceInDouble() float32`

GetOpenPriceInDouble returns the OpenPriceInDouble field if non-nil, zero value otherwise.

### GetOpenPriceInDoubleOk

`func (o *FutureOptions) GetOpenPriceInDoubleOk() (*float32, bool)`

GetOpenPriceInDoubleOk returns a tuple with the OpenPriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPriceInDouble

`func (o *FutureOptions) SetOpenPriceInDouble(v float32)`

SetOpenPriceInDouble sets OpenPriceInDouble field to given value.

### HasOpenPriceInDouble

`func (o *FutureOptions) HasOpenPriceInDouble() bool`

HasOpenPriceInDouble returns a boolean if a field has been set.

### GetRhoInDouble

`func (o *FutureOptions) GetRhoInDouble() float32`

GetRhoInDouble returns the RhoInDouble field if non-nil, zero value otherwise.

### GetRhoInDoubleOk

`func (o *FutureOptions) GetRhoInDoubleOk() (*float32, bool)`

GetRhoInDoubleOk returns a tuple with the RhoInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhoInDouble

`func (o *FutureOptions) SetRhoInDouble(v float32)`

SetRhoInDouble sets RhoInDouble field to given value.

### HasRhoInDouble

`func (o *FutureOptions) HasRhoInDouble() bool`

HasRhoInDouble returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *FutureOptions) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *FutureOptions) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *FutureOptions) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *FutureOptions) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetStrikePriceInDouble

`func (o *FutureOptions) GetStrikePriceInDouble() float32`

GetStrikePriceInDouble returns the StrikePriceInDouble field if non-nil, zero value otherwise.

### GetStrikePriceInDoubleOk

`func (o *FutureOptions) GetStrikePriceInDoubleOk() (*float32, bool)`

GetStrikePriceInDoubleOk returns a tuple with the StrikePriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePriceInDouble

`func (o *FutureOptions) SetStrikePriceInDouble(v float32)`

SetStrikePriceInDouble sets StrikePriceInDouble field to given value.

### HasStrikePriceInDouble

`func (o *FutureOptions) HasStrikePriceInDouble() bool`

HasStrikePriceInDouble returns a boolean if a field has been set.

### GetSymbol

`func (o *FutureOptions) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FutureOptions) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FutureOptions) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FutureOptions) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetThetaInDouble

`func (o *FutureOptions) GetThetaInDouble() float32`

GetThetaInDouble returns the ThetaInDouble field if non-nil, zero value otherwise.

### GetThetaInDoubleOk

`func (o *FutureOptions) GetThetaInDoubleOk() (*float32, bool)`

GetThetaInDoubleOk returns a tuple with the ThetaInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThetaInDouble

`func (o *FutureOptions) SetThetaInDouble(v float32)`

SetThetaInDouble sets ThetaInDouble field to given value.

### HasThetaInDouble

`func (o *FutureOptions) HasThetaInDouble() bool`

HasThetaInDouble returns a boolean if a field has been set.

### GetTick

`func (o *FutureOptions) GetTick() float32`

GetTick returns the Tick field if non-nil, zero value otherwise.

### GetTickOk

`func (o *FutureOptions) GetTickOk() (*float32, bool)`

GetTickOk returns a tuple with the Tick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTick

`func (o *FutureOptions) SetTick(v float32)`

SetTick sets Tick field to given value.

### HasTick

`func (o *FutureOptions) HasTick() bool`

HasTick returns a boolean if a field has been set.

### GetTickAmount

`func (o *FutureOptions) GetTickAmount() float32`

GetTickAmount returns the TickAmount field if non-nil, zero value otherwise.

### GetTickAmountOk

`func (o *FutureOptions) GetTickAmountOk() (*float32, bool)`

GetTickAmountOk returns a tuple with the TickAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickAmount

`func (o *FutureOptions) SetTickAmount(v float32)`

SetTickAmount sets TickAmount field to given value.

### HasTickAmount

`func (o *FutureOptions) HasTickAmount() bool`

HasTickAmount returns a boolean if a field has been set.

### GetTimeValueInDouble

`func (o *FutureOptions) GetTimeValueInDouble() float32`

GetTimeValueInDouble returns the TimeValueInDouble field if non-nil, zero value otherwise.

### GetTimeValueInDoubleOk

`func (o *FutureOptions) GetTimeValueInDoubleOk() (*float32, bool)`

GetTimeValueInDoubleOk returns a tuple with the TimeValueInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeValueInDouble

`func (o *FutureOptions) SetTimeValueInDouble(v float32)`

SetTimeValueInDouble sets TimeValueInDouble field to given value.

### HasTimeValueInDouble

`func (o *FutureOptions) HasTimeValueInDouble() bool`

HasTimeValueInDouble returns a boolean if a field has been set.

### GetUnderlying

`func (o *FutureOptions) GetUnderlying() string`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *FutureOptions) GetUnderlyingOk() (*string, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *FutureOptions) SetUnderlying(v string)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *FutureOptions) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.

### GetVegaInDouble

`func (o *FutureOptions) GetVegaInDouble() float32`

GetVegaInDouble returns the VegaInDouble field if non-nil, zero value otherwise.

### GetVegaInDoubleOk

`func (o *FutureOptions) GetVegaInDoubleOk() (*float32, bool)`

GetVegaInDoubleOk returns a tuple with the VegaInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVegaInDouble

`func (o *FutureOptions) SetVegaInDouble(v float32)`

SetVegaInDouble sets VegaInDouble field to given value.

### HasVegaInDouble

`func (o *FutureOptions) HasVegaInDouble() bool`

HasVegaInDouble returns a boolean if a field has been set.

### GetVolatility

`func (o *FutureOptions) GetVolatility() float32`

GetVolatility returns the Volatility field if non-nil, zero value otherwise.

### GetVolatilityOk

`func (o *FutureOptions) GetVolatilityOk() (*float32, bool)`

GetVolatilityOk returns a tuple with the Volatility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolatility

`func (o *FutureOptions) SetVolatility(v float32)`

SetVolatility sets Volatility field to given value.

### HasVolatility

`func (o *FutureOptions) HasVolatility() bool`

HasVolatility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


