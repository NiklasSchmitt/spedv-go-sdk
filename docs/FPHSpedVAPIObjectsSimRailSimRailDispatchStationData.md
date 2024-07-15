# FPHSpedVAPIObjectsSimRailSimRailDispatchStationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | [readonly] 
**Name** | **NullableString** |  | [readonly] 
**Prefix** | **NullableString** |  | [readonly] 
**Latitude** | **float64** |  | [readonly] 
**Longitude** | **float64** |  | [readonly] 
**DifficultyLevel** | **int32** |  | [readonly] 
**MainImageURL** | **NullableString** |  | [readonly] 
**AdditionalImage1URL** | **NullableString** |  | [readonly] 
**AdditionalImage2URL** | **NullableString** |  | [readonly] 
**Server** | [**NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo**](FPHSpedVAPIObjectsSimRailSimRailServerInfo.md) |  | [readonly] 
**SteamId** | **NullableInt64** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSimRailSimRailDispatchStationData

`func NewFPHSpedVAPIObjectsSimRailSimRailDispatchStationData(id NullableString, name NullableString, prefix NullableString, latitude float64, longitude float64, difficultyLevel int32, mainImageURL NullableString, additionalImage1URL NullableString, additionalImage2URL NullableString, server NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo, steamId NullableInt64, ) *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData`

NewFPHSpedVAPIObjectsSimRailSimRailDispatchStationData instantiates a new FPHSpedVAPIObjectsSimRailSimRailDispatchStationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSimRailSimRailDispatchStationDataWithDefaults

`func NewFPHSpedVAPIObjectsSimRailSimRailDispatchStationDataWithDefaults() *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData`

NewFPHSpedVAPIObjectsSimRailSimRailDispatchStationDataWithDefaults instantiates a new FPHSpedVAPIObjectsSimRailSimRailDispatchStationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPrefix

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### SetPrefixNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetLatitude

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetDifficultyLevel

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetDifficultyLevel() int32`

GetDifficultyLevel returns the DifficultyLevel field if non-nil, zero value otherwise.

### GetDifficultyLevelOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetDifficultyLevelOk() (*int32, bool)`

GetDifficultyLevelOk returns a tuple with the DifficultyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficultyLevel

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetDifficultyLevel(v int32)`

SetDifficultyLevel sets DifficultyLevel field to given value.


### GetMainImageURL

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetMainImageURL() string`

GetMainImageURL returns the MainImageURL field if non-nil, zero value otherwise.

### GetMainImageURLOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetMainImageURLOk() (*string, bool)`

GetMainImageURLOk returns a tuple with the MainImageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainImageURL

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetMainImageURL(v string)`

SetMainImageURL sets MainImageURL field to given value.


### SetMainImageURLNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetMainImageURLNil(b bool)`

 SetMainImageURLNil sets the value for MainImageURL to be an explicit nil

### UnsetMainImageURL
`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) UnsetMainImageURL()`

UnsetMainImageURL ensures that no value is present for MainImageURL, not even an explicit nil
### GetAdditionalImage1URL

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetAdditionalImage1URL() string`

GetAdditionalImage1URL returns the AdditionalImage1URL field if non-nil, zero value otherwise.

### GetAdditionalImage1URLOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetAdditionalImage1URLOk() (*string, bool)`

GetAdditionalImage1URLOk returns a tuple with the AdditionalImage1URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalImage1URL

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetAdditionalImage1URL(v string)`

SetAdditionalImage1URL sets AdditionalImage1URL field to given value.


### SetAdditionalImage1URLNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetAdditionalImage1URLNil(b bool)`

 SetAdditionalImage1URLNil sets the value for AdditionalImage1URL to be an explicit nil

### UnsetAdditionalImage1URL
`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) UnsetAdditionalImage1URL()`

UnsetAdditionalImage1URL ensures that no value is present for AdditionalImage1URL, not even an explicit nil
### GetAdditionalImage2URL

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetAdditionalImage2URL() string`

GetAdditionalImage2URL returns the AdditionalImage2URL field if non-nil, zero value otherwise.

### GetAdditionalImage2URLOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetAdditionalImage2URLOk() (*string, bool)`

GetAdditionalImage2URLOk returns a tuple with the AdditionalImage2URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalImage2URL

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetAdditionalImage2URL(v string)`

SetAdditionalImage2URL sets AdditionalImage2URL field to given value.


### SetAdditionalImage2URLNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetAdditionalImage2URLNil(b bool)`

 SetAdditionalImage2URLNil sets the value for AdditionalImage2URL to be an explicit nil

### UnsetAdditionalImage2URL
`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) UnsetAdditionalImage2URL()`

UnsetAdditionalImage2URL ensures that no value is present for AdditionalImage2URL, not even an explicit nil
### GetServer

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetServer() FPHSpedVAPIObjectsSimRailSimRailServerInfo`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetServerOk() (*FPHSpedVAPIObjectsSimRailSimRailServerInfo, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetServer(v FPHSpedVAPIObjectsSimRailSimRailServerInfo)`

SetServer sets Server field to given value.


### SetServerNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetSteamId

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetSteamId() int64`

GetSteamId returns the SteamId field if non-nil, zero value otherwise.

### GetSteamIdOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetSteamIdOk() (*int64, bool)`

GetSteamIdOk returns a tuple with the SteamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamId

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetSteamId(v int64)`

SetSteamId sets SteamId field to given value.


### SetSteamIdNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetSteamIdNil(b bool)`

 SetSteamIdNil sets the value for SteamId to be an explicit nil

### UnsetSteamId
`func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) UnsetSteamId()`

UnsetSteamId ensures that no value is present for SteamId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


