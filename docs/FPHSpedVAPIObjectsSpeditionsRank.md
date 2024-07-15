# FPHSpedVAPIObjectsSpeditionsRank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **NullableString** |  | [readonly] 
**Salary** | **int32** |  | [readonly] 
**KmSalary** | **float64** |  | [readonly] 
**Sorting** | **int32** |  | [readonly] 
**IsStartUpRank** | **bool** |  | [readonly] 
**IsLeader** | **bool** |  | [readonly] 
**IsInactiveRank** | **bool** |  | [readonly] 
**InactiveTime** | **int32** |  | [readonly] 
**Color** | **NullableString** |  | [readonly] 
**RankPermissions** | [**[]FPHSpedVAPIObjectsSpeditionsRankPermission**](FPHSpedVAPIObjectsSpeditionsRankPermission.md) |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsRank

`func NewFPHSpedVAPIObjectsSpeditionsRank(id int32, name NullableString, salary int32, kmSalary float64, sorting int32, isStartUpRank bool, isLeader bool, isInactiveRank bool, inactiveTime int32, color NullableString, rankPermissions []FPHSpedVAPIObjectsSpeditionsRankPermission, ) *FPHSpedVAPIObjectsSpeditionsRank`

NewFPHSpedVAPIObjectsSpeditionsRank instantiates a new FPHSpedVAPIObjectsSpeditionsRank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsRankWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsRankWithDefaults() *FPHSpedVAPIObjectsSpeditionsRank`

NewFPHSpedVAPIObjectsSpeditionsRankWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsRank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FPHSpedVAPIObjectsSpeditionsRank) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSalary

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetSalary() int32`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetSalaryOk() (*int32, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetSalary(v int32)`

SetSalary sets Salary field to given value.


### GetKmSalary

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetKmSalary() float64`

GetKmSalary returns the KmSalary field if non-nil, zero value otherwise.

### GetKmSalaryOk

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetKmSalaryOk() (*float64, bool)`

GetKmSalaryOk returns a tuple with the KmSalary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmSalary

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetKmSalary(v float64)`

SetKmSalary sets KmSalary field to given value.


### GetSorting

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetSorting() int32`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetSortingOk() (*int32, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetSorting(v int32)`

SetSorting sets Sorting field to given value.


### GetIsStartUpRank

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetIsStartUpRank() bool`

GetIsStartUpRank returns the IsStartUpRank field if non-nil, zero value otherwise.

### GetIsStartUpRankOk

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetIsStartUpRankOk() (*bool, bool)`

GetIsStartUpRankOk returns a tuple with the IsStartUpRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStartUpRank

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetIsStartUpRank(v bool)`

SetIsStartUpRank sets IsStartUpRank field to given value.


### GetIsLeader

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetIsLeader() bool`

GetIsLeader returns the IsLeader field if non-nil, zero value otherwise.

### GetIsLeaderOk

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetIsLeaderOk() (*bool, bool)`

GetIsLeaderOk returns a tuple with the IsLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeader

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetIsLeader(v bool)`

SetIsLeader sets IsLeader field to given value.


### GetIsInactiveRank

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetIsInactiveRank() bool`

GetIsInactiveRank returns the IsInactiveRank field if non-nil, zero value otherwise.

### GetIsInactiveRankOk

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetIsInactiveRankOk() (*bool, bool)`

GetIsInactiveRankOk returns a tuple with the IsInactiveRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInactiveRank

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetIsInactiveRank(v bool)`

SetIsInactiveRank sets IsInactiveRank field to given value.


### GetInactiveTime

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetInactiveTime() int32`

GetInactiveTime returns the InactiveTime field if non-nil, zero value otherwise.

### GetInactiveTimeOk

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetInactiveTimeOk() (*int32, bool)`

GetInactiveTimeOk returns a tuple with the InactiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveTime

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetInactiveTime(v int32)`

SetInactiveTime sets InactiveTime field to given value.


### GetColor

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetColor(v string)`

SetColor sets Color field to given value.


### SetColorNil

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *FPHSpedVAPIObjectsSpeditionsRank) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetRankPermissions

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetRankPermissions() []FPHSpedVAPIObjectsSpeditionsRankPermission`

GetRankPermissions returns the RankPermissions field if non-nil, zero value otherwise.

### GetRankPermissionsOk

`func (o *FPHSpedVAPIObjectsSpeditionsRank) GetRankPermissionsOk() (*[]FPHSpedVAPIObjectsSpeditionsRankPermission, bool)`

GetRankPermissionsOk returns a tuple with the RankPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankPermissions

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetRankPermissions(v []FPHSpedVAPIObjectsSpeditionsRankPermission)`

SetRankPermissions sets RankPermissions field to given value.


### SetRankPermissionsNil

`func (o *FPHSpedVAPIObjectsSpeditionsRank) SetRankPermissionsNil(b bool)`

 SetRankPermissionsNil sets the value for RankPermissions to be an explicit nil

### UnsetRankPermissions
`func (o *FPHSpedVAPIObjectsSpeditionsRank) UnsetRankPermissions()`

UnsetRankPermissions ensures that no value is present for RankPermissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


