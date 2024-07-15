# FPHSpedVAPIObjectsUsersUserChangeLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**User** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**Timestamp** | **time.Time** |  | [readonly] 
**Type** | [**FPHSpedVAPIEnumsUserChangeEntryType**](FPHSpedVAPIEnumsUserChangeEntryType.md) |   0 &#x3D; ChangedName  1 &#x3D; JoinedSpedition  2 &#x3D; DismissedBySpedition  3 &#x3D; Quitted  4 &#x3D; ChangedProfileData  5 &#x3D; FoundedSpedition  -1 &#x3D; NotSet | [readonly] 
**Comment** | **NullableString** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsUsersUserChangeLogEntry

`func NewFPHSpedVAPIObjectsUsersUserChangeLogEntry(id int32, user NullableFPHSpedVAPIObjectsUsersUser, timestamp time.Time, type_ FPHSpedVAPIEnumsUserChangeEntryType, comment NullableString, ) *FPHSpedVAPIObjectsUsersUserChangeLogEntry`

NewFPHSpedVAPIObjectsUsersUserChangeLogEntry instantiates a new FPHSpedVAPIObjectsUsersUserChangeLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsUsersUserChangeLogEntryWithDefaults

`func NewFPHSpedVAPIObjectsUsersUserChangeLogEntryWithDefaults() *FPHSpedVAPIObjectsUsersUserChangeLogEntry`

NewFPHSpedVAPIObjectsUsersUserChangeLogEntryWithDefaults instantiates a new FPHSpedVAPIObjectsUsersUserChangeLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) SetId(v int32)`

SetId sets Id field to given value.


### GetUser

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) GetUser() FPHSpedVAPIObjectsUsersUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) GetUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) SetUser(v FPHSpedVAPIObjectsUsersUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetTimestamp

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) GetType() FPHSpedVAPIEnumsUserChangeEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) GetTypeOk() (*FPHSpedVAPIEnumsUserChangeEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) SetType(v FPHSpedVAPIEnumsUserChangeEntryType)`

SetType sets Type field to given value.


### GetComment

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FPHSpedVAPIObjectsUsersUserChangeLogEntry) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


