# FPHSpedVAPIObjectsUsersUserStatUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserObject** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**User** | **NullableString** |  | [readonly] 
**LastTaskTime** | **time.Time** |  | [readonly] 
**DrivenKM** | **float64** |  | [readonly] 
**Income** | **float64** |  | [readonly] 
**Damage** | **float64** |  | [readonly] 
**Tasks** | **float64** |  | [readonly] 
**Fuel** | **float64** |  | [readonly] 
**FuelPer100KM** | **float64** |  | [readonly] 
**DamagePercent** | **float64** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsUsersUserStatUser

`func NewFPHSpedVAPIObjectsUsersUserStatUser(userObject NullableFPHSpedVAPIObjectsUsersUser, user NullableString, lastTaskTime time.Time, drivenKM float64, income float64, damage float64, tasks float64, fuel float64, fuelPer100KM float64, damagePercent float64, ) *FPHSpedVAPIObjectsUsersUserStatUser`

NewFPHSpedVAPIObjectsUsersUserStatUser instantiates a new FPHSpedVAPIObjectsUsersUserStatUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsUsersUserStatUserWithDefaults

`func NewFPHSpedVAPIObjectsUsersUserStatUserWithDefaults() *FPHSpedVAPIObjectsUsersUserStatUser`

NewFPHSpedVAPIObjectsUsersUserStatUserWithDefaults instantiates a new FPHSpedVAPIObjectsUsersUserStatUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserObject

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetUserObject() FPHSpedVAPIObjectsUsersUser`

GetUserObject returns the UserObject field if non-nil, zero value otherwise.

### GetUserObjectOk

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetUserObjectOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetUserObjectOk returns a tuple with the UserObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObject

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) SetUserObject(v FPHSpedVAPIObjectsUsersUser)`

SetUserObject sets UserObject field to given value.


### SetUserObjectNil

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) SetUserObjectNil(b bool)`

 SetUserObjectNil sets the value for UserObject to be an explicit nil

### UnsetUserObject
`func (o *FPHSpedVAPIObjectsUsersUserStatUser) UnsetUserObject()`

UnsetUserObject ensures that no value is present for UserObject, not even an explicit nil
### GetUser

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) SetUser(v string)`

SetUser sets User field to given value.


### SetUserNil

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *FPHSpedVAPIObjectsUsersUserStatUser) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetLastTaskTime

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetLastTaskTime() time.Time`

GetLastTaskTime returns the LastTaskTime field if non-nil, zero value otherwise.

### GetLastTaskTimeOk

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetLastTaskTimeOk() (*time.Time, bool)`

GetLastTaskTimeOk returns a tuple with the LastTaskTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTaskTime

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) SetLastTaskTime(v time.Time)`

SetLastTaskTime sets LastTaskTime field to given value.


### GetDrivenKM

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetDrivenKM() float64`

GetDrivenKM returns the DrivenKM field if non-nil, zero value otherwise.

### GetDrivenKMOk

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetDrivenKMOk() (*float64, bool)`

GetDrivenKMOk returns a tuple with the DrivenKM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivenKM

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) SetDrivenKM(v float64)`

SetDrivenKM sets DrivenKM field to given value.


### GetIncome

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetIncome() float64`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetIncomeOk() (*float64, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) SetIncome(v float64)`

SetIncome sets Income field to given value.


### GetDamage

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetDamage() float64`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetDamageOk() (*float64, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) SetDamage(v float64)`

SetDamage sets Damage field to given value.


### GetTasks

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetTasks() float64`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetTasksOk() (*float64, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) SetTasks(v float64)`

SetTasks sets Tasks field to given value.


### GetFuel

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetFuel() float64`

GetFuel returns the Fuel field if non-nil, zero value otherwise.

### GetFuelOk

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetFuelOk() (*float64, bool)`

GetFuelOk returns a tuple with the Fuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuel

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) SetFuel(v float64)`

SetFuel sets Fuel field to given value.


### GetFuelPer100KM

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetFuelPer100KM() float64`

GetFuelPer100KM returns the FuelPer100KM field if non-nil, zero value otherwise.

### GetFuelPer100KMOk

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetFuelPer100KMOk() (*float64, bool)`

GetFuelPer100KMOk returns a tuple with the FuelPer100KM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelPer100KM

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) SetFuelPer100KM(v float64)`

SetFuelPer100KM sets FuelPer100KM field to given value.


### GetDamagePercent

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetDamagePercent() float64`

GetDamagePercent returns the DamagePercent field if non-nil, zero value otherwise.

### GetDamagePercentOk

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) GetDamagePercentOk() (*float64, bool)`

GetDamagePercentOk returns a tuple with the DamagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamagePercent

`func (o *FPHSpedVAPIObjectsUsersUserStatUser) SetDamagePercent(v float64)`

SetDamagePercent sets DamagePercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


