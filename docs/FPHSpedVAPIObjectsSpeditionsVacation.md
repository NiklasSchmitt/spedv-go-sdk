# FPHSpedVAPIObjectsSpeditionsVacation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**User** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**UserName** | **NullableString** |  | [readonly] 
**Description** | **NullableString** |  | [readonly] 
**Start** | **time.Time** |  | [readonly] 
**Ende** | **time.Time** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsVacation

`func NewFPHSpedVAPIObjectsSpeditionsVacation(id int32, user NullableFPHSpedVAPIObjectsUsersUser, userName NullableString, description NullableString, start time.Time, ende time.Time, ) *FPHSpedVAPIObjectsSpeditionsVacation`

NewFPHSpedVAPIObjectsSpeditionsVacation instantiates a new FPHSpedVAPIObjectsSpeditionsVacation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsVacationWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsVacationWithDefaults() *FPHSpedVAPIObjectsSpeditionsVacation`

NewFPHSpedVAPIObjectsSpeditionsVacationWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsVacation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) SetId(v int32)`

SetId sets Id field to given value.


### GetUser

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) GetUser() FPHSpedVAPIObjectsUsersUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) GetUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) SetUser(v FPHSpedVAPIObjectsUsersUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *FPHSpedVAPIObjectsSpeditionsVacation) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUserName

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) SetUserName(v string)`

SetUserName sets UserName field to given value.


### SetUserNameNil

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *FPHSpedVAPIObjectsSpeditionsVacation) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetDescription

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FPHSpedVAPIObjectsSpeditionsVacation) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStart

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnde

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) GetEnde() time.Time`

GetEnde returns the Ende field if non-nil, zero value otherwise.

### GetEndeOk

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) GetEndeOk() (*time.Time, bool)`

GetEndeOk returns a tuple with the Ende field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnde

`func (o *FPHSpedVAPIObjectsSpeditionsVacation) SetEnde(v time.Time)`

SetEnde sets Ende field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


