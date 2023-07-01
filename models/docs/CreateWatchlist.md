# CreateWatchlist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**WatchlistItems** | Pointer to [**[]CreateWatchlistWatchlistItems**](CreateWatchlistWatchlistItems.md) |  | [optional] 

## Methods

### NewCreateWatchlist

`func NewCreateWatchlist() *CreateWatchlist`

NewCreateWatchlist instantiates a new CreateWatchlist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWatchlistWithDefaults

`func NewCreateWatchlistWithDefaults() *CreateWatchlist`

NewCreateWatchlistWithDefaults instantiates a new CreateWatchlist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWatchlist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWatchlist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWatchlist) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateWatchlist) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWatchlistItems

`func (o *CreateWatchlist) GetWatchlistItems() []CreateWatchlistWatchlistItems`

GetWatchlistItems returns the WatchlistItems field if non-nil, zero value otherwise.

### GetWatchlistItemsOk

`func (o *CreateWatchlist) GetWatchlistItemsOk() (*[]CreateWatchlistWatchlistItems, bool)`

GetWatchlistItemsOk returns a tuple with the WatchlistItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistItems

`func (o *CreateWatchlist) SetWatchlistItems(v []CreateWatchlistWatchlistItems)`

SetWatchlistItems sets WatchlistItems field to given value.

### HasWatchlistItems

`func (o *CreateWatchlist) HasWatchlistItems() bool`

HasWatchlistItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


