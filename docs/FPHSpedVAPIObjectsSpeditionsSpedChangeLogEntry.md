# FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**User** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**Timestamp** | **time.Time** |  | [readonly] 
**Type** | [**FPHSpedVAPIEnumsSpeditionChangeEntryType**](FPHSpedVAPIEnumsSpeditionChangeEntryType.md) |   0 &#x3D; TaskDeducted  1 &#x3D; TaskRejected  2 &#x3D; RankChanged  3 &#x3D; RankDeleted  4 &#x3D; UserRankChanged  5 &#x3D; UserDismissed  6 &#x3D; SpeditionSettingChanged  7 &#x3D; ApplicationTextChanged  8 &#x3D; ApplicationAcknowledged  9 &#x3D; ApplicationRejected  10 &#x3D; BranchBought  11 &#x3D; BranchExpandedTruckParkplace  12 &#x3D; BranchShrinkedTruckParkplace  13 &#x3D; BranchDeleted  14 &#x3D; TruckBought  15 &#x3D; TruckSold  16 &#x3D; TruckLicensePlateChanged  17 &#x3D; BoniAdded  18 &#x3D; BoniDeleted  19 &#x3D; RankAddedPermission  20 &#x3D; RankDeletedPermission  21 &#x3D; BranchExpandedTrailerParkplace  22 &#x3D; BranchShrinkedTrailerParkplace  23 &#x3D; BranchExpandedMaintenancePlace  24 &#x3D; BranchShrinkedMaintenancePlace  25 &#x3D; TruckTransfered  26 &#x3D; TruckMaintenancePlanned  27 &#x3D; TruckMaintenanceAborted  -1 &#x3D; NotSet | [readonly] 
**Comment** | **NullableString** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry

`func NewFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry(id int32, user NullableFPHSpedVAPIObjectsUsersUser, timestamp time.Time, type_ FPHSpedVAPIEnumsSpeditionChangeEntryType, comment NullableString, ) *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry`

NewFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry instantiates a new FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntryWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntryWithDefaults() *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry`

NewFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntryWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) SetId(v int32)`

SetId sets Id field to given value.


### GetUser

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetUser() FPHSpedVAPIObjectsUsersUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) SetUser(v FPHSpedVAPIObjectsUsersUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetTimestamp

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetType() FPHSpedVAPIEnumsSpeditionChangeEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetTypeOk() (*FPHSpedVAPIEnumsSpeditionChangeEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) SetType(v FPHSpedVAPIEnumsSpeditionChangeEntryType)`

SetType sets Type field to given value.


### GetComment

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


