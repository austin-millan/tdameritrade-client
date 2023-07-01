# WatchlistInstrument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetType** | Pointer to [**AssetType**](AssetType.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewWatchlistInstrument

`func NewWatchlistInstrument() *WatchlistInstrument`

NewWatchlistInstrument instantiates a new WatchlistInstrument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchlistInstrumentWithDefaults

`func NewWatchlistInstrumentWithDefaults() *WatchlistInstrument`

NewWatchlistInstrumentWithDefaults instantiates a new WatchlistInstrument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetType

`func (o *WatchlistInstrument) GetAssetType() AssetType`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *WatchlistInstrument) GetAssetTypeOk() (*AssetType, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *WatchlistInstrument) SetAssetType(v AssetType)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *WatchlistInstrument) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetDescription

`func (o *WatchlistInstrument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WatchlistInstrument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WatchlistInstrument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WatchlistInstrument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSymbol

`func (o *WatchlistInstrument) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *WatchlistInstrument) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *WatchlistInstrument) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *WatchlistInstrument) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


