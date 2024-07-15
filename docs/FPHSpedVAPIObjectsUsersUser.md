# FPHSpedVAPIObjectsUsersUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**UserName** | **NullableString** |  | [readonly] 
**ProfilePicURL** | **NullableString** |  | [readonly] 
**Spedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**InSped** | **bool** |  | [readonly] 
**FormatName** | **NullableString** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsUsersUser

`func NewFPHSpedVAPIObjectsUsersUser(id int32, userName NullableString, profilePicURL NullableString, spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, inSped bool, formatName NullableString, ) *FPHSpedVAPIObjectsUsersUser`

NewFPHSpedVAPIObjectsUsersUser instantiates a new FPHSpedVAPIObjectsUsersUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsUsersUserWithDefaults

`func NewFPHSpedVAPIObjectsUsersUserWithDefaults() *FPHSpedVAPIObjectsUsersUser`

NewFPHSpedVAPIObjectsUsersUserWithDefaults instantiates a new FPHSpedVAPIObjectsUsersUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsUsersUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsUsersUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsUsersUser) SetId(v int32)`

SetId sets Id field to given value.


### GetUserName

`func (o *FPHSpedVAPIObjectsUsersUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FPHSpedVAPIObjectsUsersUser) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FPHSpedVAPIObjectsUsersUser) SetUserName(v string)`

SetUserName sets UserName field to given value.


### SetUserNameNil

`func (o *FPHSpedVAPIObjectsUsersUser) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *FPHSpedVAPIObjectsUsersUser) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetProfilePicURL

`func (o *FPHSpedVAPIObjectsUsersUser) GetProfilePicURL() string`

GetProfilePicURL returns the ProfilePicURL field if non-nil, zero value otherwise.

### GetProfilePicURLOk

`func (o *FPHSpedVAPIObjectsUsersUser) GetProfilePicURLOk() (*string, bool)`

GetProfilePicURLOk returns a tuple with the ProfilePicURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicURL

`func (o *FPHSpedVAPIObjectsUsersUser) SetProfilePicURL(v string)`

SetProfilePicURL sets ProfilePicURL field to given value.


### SetProfilePicURLNil

`func (o *FPHSpedVAPIObjectsUsersUser) SetProfilePicURLNil(b bool)`

 SetProfilePicURLNil sets the value for ProfilePicURL to be an explicit nil

### UnsetProfilePicURL
`func (o *FPHSpedVAPIObjectsUsersUser) UnsetProfilePicURL()`

UnsetProfilePicURL ensures that no value is present for ProfilePicURL, not even an explicit nil
### GetSpedition

`func (o *FPHSpedVAPIObjectsUsersUser) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetSpedition returns the Spedition field if non-nil, zero value otherwise.

### GetSpeditionOk

`func (o *FPHSpedVAPIObjectsUsersUser) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetSpeditionOk returns a tuple with the Spedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpedition

`func (o *FPHSpedVAPIObjectsUsersUser) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetSpedition sets Spedition field to given value.


### SetSpeditionNil

`func (o *FPHSpedVAPIObjectsUsersUser) SetSpeditionNil(b bool)`

 SetSpeditionNil sets the value for Spedition to be an explicit nil

### UnsetSpedition
`func (o *FPHSpedVAPIObjectsUsersUser) UnsetSpedition()`

UnsetSpedition ensures that no value is present for Spedition, not even an explicit nil
### GetInSped

`func (o *FPHSpedVAPIObjectsUsersUser) GetInSped() bool`

GetInSped returns the InSped field if non-nil, zero value otherwise.

### GetInSpedOk

`func (o *FPHSpedVAPIObjectsUsersUser) GetInSpedOk() (*bool, bool)`

GetInSpedOk returns a tuple with the InSped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInSped

`func (o *FPHSpedVAPIObjectsUsersUser) SetInSped(v bool)`

SetInSped sets InSped field to given value.


### GetFormatName

`func (o *FPHSpedVAPIObjectsUsersUser) GetFormatName() string`

GetFormatName returns the FormatName field if non-nil, zero value otherwise.

### GetFormatNameOk

`func (o *FPHSpedVAPIObjectsUsersUser) GetFormatNameOk() (*string, bool)`

GetFormatNameOk returns a tuple with the FormatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatName

`func (o *FPHSpedVAPIObjectsUsersUser) SetFormatName(v string)`

SetFormatName sets FormatName field to given value.


### SetFormatNameNil

`func (o *FPHSpedVAPIObjectsUsersUser) SetFormatNameNil(b bool)`

 SetFormatNameNil sets the value for FormatName to be an explicit nil

### UnsetFormatName
`func (o *FPHSpedVAPIObjectsUsersUser) UnsetFormatName()`

UnsetFormatName ensures that no value is present for FormatName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


