# FPHSpedVAPIObjectsSpeditionsSpedition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **NullableString** |  | [readonly] 
**ShortName** | **NullableString** |  | [readonly] 
**SpeditionType** | [**FPHSpedVAPIEnumsSpeditionType**](FPHSpedVAPIEnumsSpeditionType.md) |   0 &#x3D; NonCompeting  1 &#x3D; LightRealism  2 &#x3D; RealEco  -1 &#x3D; NotSet | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsSpedition

`func NewFPHSpedVAPIObjectsSpeditionsSpedition(id int32, name NullableString, shortName NullableString, speditionType FPHSpedVAPIEnumsSpeditionType, ) *FPHSpedVAPIObjectsSpeditionsSpedition`

NewFPHSpedVAPIObjectsSpeditionsSpedition instantiates a new FPHSpedVAPIObjectsSpeditionsSpedition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsSpeditionWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsSpeditionWithDefaults() *FPHSpedVAPIObjectsSpeditionsSpedition`

NewFPHSpedVAPIObjectsSpeditionsSpeditionWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsSpedition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetShortName

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### SetShortNameNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) SetShortNameNil(b bool)`

 SetShortNameNil sets the value for ShortName to be an explicit nil

### UnsetShortName
`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) UnsetShortName()`

UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
### GetSpeditionType

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) GetSpeditionType() FPHSpedVAPIEnumsSpeditionType`

GetSpeditionType returns the SpeditionType field if non-nil, zero value otherwise.

### GetSpeditionTypeOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) GetSpeditionTypeOk() (*FPHSpedVAPIEnumsSpeditionType, bool)`

GetSpeditionTypeOk returns a tuple with the SpeditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeditionType

`func (o *FPHSpedVAPIObjectsSpeditionsSpedition) SetSpeditionType(v FPHSpedVAPIEnumsSpeditionType)`

SetSpeditionType sets SpeditionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


