# FPHSpedVAPIObjectsChatSystemChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Sender** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**SendDate** | **time.Time** |  | [readonly] 
**Message** | **NullableString** |  | [readonly] 
**AnswerTo** | [**NullableFPHSpedVAPIObjectsChatSystemChatMessage**](FPHSpedVAPIObjectsChatSystemChatMessage.md) |  | [readonly] 
**IsDeleted** | **bool** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsChatSystemChatMessage

`func NewFPHSpedVAPIObjectsChatSystemChatMessage(id int32, sender NullableFPHSpedVAPIObjectsUsersUser, sendDate time.Time, message NullableString, answerTo NullableFPHSpedVAPIObjectsChatSystemChatMessage, isDeleted bool, ) *FPHSpedVAPIObjectsChatSystemChatMessage`

NewFPHSpedVAPIObjectsChatSystemChatMessage instantiates a new FPHSpedVAPIObjectsChatSystemChatMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsChatSystemChatMessageWithDefaults

`func NewFPHSpedVAPIObjectsChatSystemChatMessageWithDefaults() *FPHSpedVAPIObjectsChatSystemChatMessage`

NewFPHSpedVAPIObjectsChatSystemChatMessageWithDefaults instantiates a new FPHSpedVAPIObjectsChatSystemChatMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) SetId(v int32)`

SetId sets Id field to given value.


### GetSender

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) GetSender() FPHSpedVAPIObjectsUsersUser`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) GetSenderOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) SetSender(v FPHSpedVAPIObjectsUsersUser)`

SetSender sets Sender field to given value.


### SetSenderNil

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetSendDate

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.


### GetMessage

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetAnswerTo

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) GetAnswerTo() FPHSpedVAPIObjectsChatSystemChatMessage`

GetAnswerTo returns the AnswerTo field if non-nil, zero value otherwise.

### GetAnswerToOk

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) GetAnswerToOk() (*FPHSpedVAPIObjectsChatSystemChatMessage, bool)`

GetAnswerToOk returns a tuple with the AnswerTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerTo

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) SetAnswerTo(v FPHSpedVAPIObjectsChatSystemChatMessage)`

SetAnswerTo sets AnswerTo field to given value.


### SetAnswerToNil

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) SetAnswerToNil(b bool)`

 SetAnswerToNil sets the value for AnswerTo to be an explicit nil

### UnsetAnswerTo
`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) UnsetAnswerTo()`

UnsetAnswerTo ensures that no value is present for AnswerTo, not even an explicit nil
### GetIsDeleted

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *FPHSpedVAPIObjectsChatSystemChatMessage) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


