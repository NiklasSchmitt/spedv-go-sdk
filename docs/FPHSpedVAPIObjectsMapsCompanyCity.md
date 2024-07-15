# FPHSpedVAPIObjectsMapsCompanyCity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Company** | [**NullableFPHSpedVAPIObjectsMapsCompany**](FPHSpedVAPIObjectsMapsCompany.md) |  | [readonly] 
**City** | [**NullableFPHSpedVAPIObjectsMapsCity**](FPHSpedVAPIObjectsMapsCity.md) |  | [readonly] 
**Map** | [**NullableFPHSpedVAPIObjectsMapsMap**](FPHSpedVAPIObjectsMapsMap.md) |  | [readonly] 
**SupportsDouble** | **bool** |  | [readonly] 
**SupportsTriple** | **bool** |  | [readonly] 
**IsDeleted** | **bool** |  | [readonly] 
**XCoord** | **NullableFloat64** |  | [readonly] 
**ZCoord** | **NullableFloat64** |  | [readonly] 
**DisabledOnMaps** | [**[]FPHSpedVAPIObjectsMapsMap**](FPHSpedVAPIObjectsMapsMap.md) |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsMapsCompanyCity

`func NewFPHSpedVAPIObjectsMapsCompanyCity(id int32, company NullableFPHSpedVAPIObjectsMapsCompany, city NullableFPHSpedVAPIObjectsMapsCity, map_ NullableFPHSpedVAPIObjectsMapsMap, supportsDouble bool, supportsTriple bool, isDeleted bool, xCoord NullableFloat64, zCoord NullableFloat64, disabledOnMaps []FPHSpedVAPIObjectsMapsMap, ) *FPHSpedVAPIObjectsMapsCompanyCity`

NewFPHSpedVAPIObjectsMapsCompanyCity instantiates a new FPHSpedVAPIObjectsMapsCompanyCity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsMapsCompanyCityWithDefaults

`func NewFPHSpedVAPIObjectsMapsCompanyCityWithDefaults() *FPHSpedVAPIObjectsMapsCompanyCity`

NewFPHSpedVAPIObjectsMapsCompanyCityWithDefaults instantiates a new FPHSpedVAPIObjectsMapsCompanyCity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetId(v int32)`

SetId sets Id field to given value.


### GetCompany

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetCompany() FPHSpedVAPIObjectsMapsCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetCompanyOk() (*FPHSpedVAPIObjectsMapsCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetCompany(v FPHSpedVAPIObjectsMapsCompany)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *FPHSpedVAPIObjectsMapsCompanyCity) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCity

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetCity() FPHSpedVAPIObjectsMapsCity`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetCityOk() (*FPHSpedVAPIObjectsMapsCity, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetCity(v FPHSpedVAPIObjectsMapsCity)`

SetCity sets City field to given value.


### SetCityNil

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *FPHSpedVAPIObjectsMapsCompanyCity) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetMap

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetMap() FPHSpedVAPIObjectsMapsMap`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetMapOk() (*FPHSpedVAPIObjectsMapsMap, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetMap(v FPHSpedVAPIObjectsMapsMap)`

SetMap sets Map field to given value.


### SetMapNil

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetMapNil(b bool)`

 SetMapNil sets the value for Map to be an explicit nil

### UnsetMap
`func (o *FPHSpedVAPIObjectsMapsCompanyCity) UnsetMap()`

UnsetMap ensures that no value is present for Map, not even an explicit nil
### GetSupportsDouble

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetSupportsDouble() bool`

GetSupportsDouble returns the SupportsDouble field if non-nil, zero value otherwise.

### GetSupportsDoubleOk

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetSupportsDoubleOk() (*bool, bool)`

GetSupportsDoubleOk returns a tuple with the SupportsDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDouble

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetSupportsDouble(v bool)`

SetSupportsDouble sets SupportsDouble field to given value.


### GetSupportsTriple

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetSupportsTriple() bool`

GetSupportsTriple returns the SupportsTriple field if non-nil, zero value otherwise.

### GetSupportsTripleOk

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetSupportsTripleOk() (*bool, bool)`

GetSupportsTripleOk returns a tuple with the SupportsTriple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsTriple

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetSupportsTriple(v bool)`

SetSupportsTriple sets SupportsTriple field to given value.


### GetIsDeleted

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetXCoord

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetXCoord() float64`

GetXCoord returns the XCoord field if non-nil, zero value otherwise.

### GetXCoordOk

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetXCoordOk() (*float64, bool)`

GetXCoordOk returns a tuple with the XCoord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXCoord

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetXCoord(v float64)`

SetXCoord sets XCoord field to given value.


### SetXCoordNil

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetXCoordNil(b bool)`

 SetXCoordNil sets the value for XCoord to be an explicit nil

### UnsetXCoord
`func (o *FPHSpedVAPIObjectsMapsCompanyCity) UnsetXCoord()`

UnsetXCoord ensures that no value is present for XCoord, not even an explicit nil
### GetZCoord

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetZCoord() float64`

GetZCoord returns the ZCoord field if non-nil, zero value otherwise.

### GetZCoordOk

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetZCoordOk() (*float64, bool)`

GetZCoordOk returns a tuple with the ZCoord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZCoord

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetZCoord(v float64)`

SetZCoord sets ZCoord field to given value.


### SetZCoordNil

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetZCoordNil(b bool)`

 SetZCoordNil sets the value for ZCoord to be an explicit nil

### UnsetZCoord
`func (o *FPHSpedVAPIObjectsMapsCompanyCity) UnsetZCoord()`

UnsetZCoord ensures that no value is present for ZCoord, not even an explicit nil
### GetDisabledOnMaps

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetDisabledOnMaps() []FPHSpedVAPIObjectsMapsMap`

GetDisabledOnMaps returns the DisabledOnMaps field if non-nil, zero value otherwise.

### GetDisabledOnMapsOk

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) GetDisabledOnMapsOk() (*[]FPHSpedVAPIObjectsMapsMap, bool)`

GetDisabledOnMapsOk returns a tuple with the DisabledOnMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledOnMaps

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetDisabledOnMaps(v []FPHSpedVAPIObjectsMapsMap)`

SetDisabledOnMaps sets DisabledOnMaps field to given value.


### SetDisabledOnMapsNil

`func (o *FPHSpedVAPIObjectsMapsCompanyCity) SetDisabledOnMapsNil(b bool)`

 SetDisabledOnMapsNil sets the value for DisabledOnMaps to be an explicit nil

### UnsetDisabledOnMaps
`func (o *FPHSpedVAPIObjectsMapsCompanyCity) UnsetDisabledOnMaps()`

UnsetDisabledOnMaps ensures that no value is present for DisabledOnMaps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


