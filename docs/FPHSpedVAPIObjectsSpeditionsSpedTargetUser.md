# FPHSpedVAPIObjectsSpeditionsSpedTargetUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserObject** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**User** | **NullableString** |  | [readonly] 
**Amount** | **int32** |  | [readonly] 
**AmountReached** | **float64** |  | [readonly] 
**Unit** | **NullableString** |  | [readonly] 
**Value** | **NullableString** |  | [readonly] 
**VisAmountReached** | **NullableString** |  | [readonly] 
**Max** | **int32** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsSpedTargetUser

`func NewFPHSpedVAPIObjectsSpeditionsSpedTargetUser(userObject NullableFPHSpedVAPIObjectsUsersUser, user NullableString, amount int32, amountReached float64, unit NullableString, value NullableString, visAmountReached NullableString, max int32, ) *FPHSpedVAPIObjectsSpeditionsSpedTargetUser`

NewFPHSpedVAPIObjectsSpeditionsSpedTargetUser instantiates a new FPHSpedVAPIObjectsSpeditionsSpedTargetUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsSpedTargetUserWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsSpedTargetUserWithDefaults() *FPHSpedVAPIObjectsSpeditionsSpedTargetUser`

NewFPHSpedVAPIObjectsSpeditionsSpedTargetUserWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsSpedTargetUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserObject

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetUserObject() FPHSpedVAPIObjectsUsersUser`

GetUserObject returns the UserObject field if non-nil, zero value otherwise.

### GetUserObjectOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetUserObjectOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetUserObjectOk returns a tuple with the UserObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObject

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetUserObject(v FPHSpedVAPIObjectsUsersUser)`

SetUserObject sets UserObject field to given value.


### SetUserObjectNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetUserObjectNil(b bool)`

 SetUserObjectNil sets the value for UserObject to be an explicit nil

### UnsetUserObject
`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) UnsetUserObject()`

UnsetUserObject ensures that no value is present for UserObject, not even an explicit nil
### GetUser

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetUser(v string)`

SetUser sets User field to given value.


### SetUserNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetAmount

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAmountReached

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetAmountReached() float64`

GetAmountReached returns the AmountReached field if non-nil, zero value otherwise.

### GetAmountReachedOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetAmountReachedOk() (*float64, bool)`

GetAmountReachedOk returns a tuple with the AmountReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountReached

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetAmountReached(v float64)`

SetAmountReached sets AmountReached field to given value.


### GetUnit

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetValue

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetVisAmountReached

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetVisAmountReached() string`

GetVisAmountReached returns the VisAmountReached field if non-nil, zero value otherwise.

### GetVisAmountReachedOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetVisAmountReachedOk() (*string, bool)`

GetVisAmountReachedOk returns a tuple with the VisAmountReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisAmountReached

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetVisAmountReached(v string)`

SetVisAmountReached sets VisAmountReached field to given value.


### SetVisAmountReachedNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetVisAmountReachedNil(b bool)`

 SetVisAmountReachedNil sets the value for VisAmountReached to be an explicit nil

### UnsetVisAmountReached
`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) UnsetVisAmountReached()`

UnsetVisAmountReached ensures that no value is present for VisAmountReached, not even an explicit nil
### GetMax

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetMax(v int32)`

SetMax sets Max field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


