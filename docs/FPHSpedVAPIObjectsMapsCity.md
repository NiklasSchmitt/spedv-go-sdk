# FPHSpedVAPIObjectsMapsCity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **NullableString** |  | [readonly] 
**VisibleName** | **NullableString** |  | [readonly] 
**InGameNameDictionary** | **map[string]string** |  | [readonly] 
**Country** | [**NullableFPHSpedVAPIObjectsMapsCountry**](FPHSpedVAPIObjectsMapsCountry.md) |  | [readonly] 
**RealWorldLati** | **float64** |  | [readonly] 
**RealWorldLong** | **float64** |  | [readonly] 
**InGameNames** | **[]string** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsMapsCity

`func NewFPHSpedVAPIObjectsMapsCity(id int32, name NullableString, visibleName NullableString, inGameNameDictionary map[string]string, country NullableFPHSpedVAPIObjectsMapsCountry, realWorldLati float64, realWorldLong float64, inGameNames []string, ) *FPHSpedVAPIObjectsMapsCity`

NewFPHSpedVAPIObjectsMapsCity instantiates a new FPHSpedVAPIObjectsMapsCity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsMapsCityWithDefaults

`func NewFPHSpedVAPIObjectsMapsCityWithDefaults() *FPHSpedVAPIObjectsMapsCity`

NewFPHSpedVAPIObjectsMapsCityWithDefaults instantiates a new FPHSpedVAPIObjectsMapsCity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsMapsCity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsMapsCity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsMapsCity) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FPHSpedVAPIObjectsMapsCity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPHSpedVAPIObjectsMapsCity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPHSpedVAPIObjectsMapsCity) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *FPHSpedVAPIObjectsMapsCity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FPHSpedVAPIObjectsMapsCity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVisibleName

`func (o *FPHSpedVAPIObjectsMapsCity) GetVisibleName() string`

GetVisibleName returns the VisibleName field if non-nil, zero value otherwise.

### GetVisibleNameOk

`func (o *FPHSpedVAPIObjectsMapsCity) GetVisibleNameOk() (*string, bool)`

GetVisibleNameOk returns a tuple with the VisibleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleName

`func (o *FPHSpedVAPIObjectsMapsCity) SetVisibleName(v string)`

SetVisibleName sets VisibleName field to given value.


### SetVisibleNameNil

`func (o *FPHSpedVAPIObjectsMapsCity) SetVisibleNameNil(b bool)`

 SetVisibleNameNil sets the value for VisibleName to be an explicit nil

### UnsetVisibleName
`func (o *FPHSpedVAPIObjectsMapsCity) UnsetVisibleName()`

UnsetVisibleName ensures that no value is present for VisibleName, not even an explicit nil
### GetInGameNameDictionary

`func (o *FPHSpedVAPIObjectsMapsCity) GetInGameNameDictionary() map[string]string`

GetInGameNameDictionary returns the InGameNameDictionary field if non-nil, zero value otherwise.

### GetInGameNameDictionaryOk

`func (o *FPHSpedVAPIObjectsMapsCity) GetInGameNameDictionaryOk() (*map[string]string, bool)`

GetInGameNameDictionaryOk returns a tuple with the InGameNameDictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInGameNameDictionary

`func (o *FPHSpedVAPIObjectsMapsCity) SetInGameNameDictionary(v map[string]string)`

SetInGameNameDictionary sets InGameNameDictionary field to given value.


### SetInGameNameDictionaryNil

`func (o *FPHSpedVAPIObjectsMapsCity) SetInGameNameDictionaryNil(b bool)`

 SetInGameNameDictionaryNil sets the value for InGameNameDictionary to be an explicit nil

### UnsetInGameNameDictionary
`func (o *FPHSpedVAPIObjectsMapsCity) UnsetInGameNameDictionary()`

UnsetInGameNameDictionary ensures that no value is present for InGameNameDictionary, not even an explicit nil
### GetCountry

`func (o *FPHSpedVAPIObjectsMapsCity) GetCountry() FPHSpedVAPIObjectsMapsCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FPHSpedVAPIObjectsMapsCity) GetCountryOk() (*FPHSpedVAPIObjectsMapsCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FPHSpedVAPIObjectsMapsCity) SetCountry(v FPHSpedVAPIObjectsMapsCountry)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *FPHSpedVAPIObjectsMapsCity) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *FPHSpedVAPIObjectsMapsCity) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetRealWorldLati

`func (o *FPHSpedVAPIObjectsMapsCity) GetRealWorldLati() float64`

GetRealWorldLati returns the RealWorldLati field if non-nil, zero value otherwise.

### GetRealWorldLatiOk

`func (o *FPHSpedVAPIObjectsMapsCity) GetRealWorldLatiOk() (*float64, bool)`

GetRealWorldLatiOk returns a tuple with the RealWorldLati field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealWorldLati

`func (o *FPHSpedVAPIObjectsMapsCity) SetRealWorldLati(v float64)`

SetRealWorldLati sets RealWorldLati field to given value.


### GetRealWorldLong

`func (o *FPHSpedVAPIObjectsMapsCity) GetRealWorldLong() float64`

GetRealWorldLong returns the RealWorldLong field if non-nil, zero value otherwise.

### GetRealWorldLongOk

`func (o *FPHSpedVAPIObjectsMapsCity) GetRealWorldLongOk() (*float64, bool)`

GetRealWorldLongOk returns a tuple with the RealWorldLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealWorldLong

`func (o *FPHSpedVAPIObjectsMapsCity) SetRealWorldLong(v float64)`

SetRealWorldLong sets RealWorldLong field to given value.


### GetInGameNames

`func (o *FPHSpedVAPIObjectsMapsCity) GetInGameNames() []string`

GetInGameNames returns the InGameNames field if non-nil, zero value otherwise.

### GetInGameNamesOk

`func (o *FPHSpedVAPIObjectsMapsCity) GetInGameNamesOk() (*[]string, bool)`

GetInGameNamesOk returns a tuple with the InGameNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInGameNames

`func (o *FPHSpedVAPIObjectsMapsCity) SetInGameNames(v []string)`

SetInGameNames sets InGameNames field to given value.


### SetInGameNamesNil

`func (o *FPHSpedVAPIObjectsMapsCity) SetInGameNamesNil(b bool)`

 SetInGameNamesNil sets the value for InGameNames to be an explicit nil

### UnsetInGameNames
`func (o *FPHSpedVAPIObjectsMapsCity) UnsetInGameNames()`

UnsetInGameNames ensures that no value is present for InGameNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


