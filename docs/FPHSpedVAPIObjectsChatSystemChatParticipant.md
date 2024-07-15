# FPHSpedVAPIObjectsChatSystemChatParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | 
**HasArchived** | **bool** |  | 
**LastView** | **time.Time** |  | 
**IsAdmin** | **bool** |  | 

## Methods

### NewFPHSpedVAPIObjectsChatSystemChatParticipant

`func NewFPHSpedVAPIObjectsChatSystemChatParticipant(user NullableFPHSpedVAPIObjectsUsersUser, hasArchived bool, lastView time.Time, isAdmin bool, ) *FPHSpedVAPIObjectsChatSystemChatParticipant`

NewFPHSpedVAPIObjectsChatSystemChatParticipant instantiates a new FPHSpedVAPIObjectsChatSystemChatParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsChatSystemChatParticipantWithDefaults

`func NewFPHSpedVAPIObjectsChatSystemChatParticipantWithDefaults() *FPHSpedVAPIObjectsChatSystemChatParticipant`

NewFPHSpedVAPIObjectsChatSystemChatParticipantWithDefaults instantiates a new FPHSpedVAPIObjectsChatSystemChatParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) GetUser() FPHSpedVAPIObjectsUsersUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) GetUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) SetUser(v FPHSpedVAPIObjectsUsersUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetHasArchived

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) GetHasArchived() bool`

GetHasArchived returns the HasArchived field if non-nil, zero value otherwise.

### GetHasArchivedOk

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) GetHasArchivedOk() (*bool, bool)`

GetHasArchivedOk returns a tuple with the HasArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasArchived

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) SetHasArchived(v bool)`

SetHasArchived sets HasArchived field to given value.


### GetLastView

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) GetLastView() time.Time`

GetLastView returns the LastView field if non-nil, zero value otherwise.

### GetLastViewOk

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) GetLastViewOk() (*time.Time, bool)`

GetLastViewOk returns a tuple with the LastView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastView

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) SetLastView(v time.Time)`

SetLastView sets LastView field to given value.


### GetIsAdmin

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *FPHSpedVAPIObjectsChatSystemChatParticipant) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


