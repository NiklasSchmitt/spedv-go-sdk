# FPHSpedVAPIObjectsChatSystemChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Initiator** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | 
**StartTime** | **time.Time** |  | 
**Subject** | **NullableString** |  | 
**LastMessageTimestamp** | **NullableTime** |  | 
**LastMessageString** | **NullableString** |  | [readonly] 
**Participants** | [**[]FPHSpedVAPIObjectsChatSystemChatParticipant**](FPHSpedVAPIObjectsChatSystemChatParticipant.md) |  | 
**ParticipantString** | **NullableString** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsChatSystemChat

`func NewFPHSpedVAPIObjectsChatSystemChat(id int32, initiator NullableFPHSpedVAPIObjectsUsersUser, startTime time.Time, subject NullableString, lastMessageTimestamp NullableTime, lastMessageString NullableString, participants []FPHSpedVAPIObjectsChatSystemChatParticipant, participantString NullableString, ) *FPHSpedVAPIObjectsChatSystemChat`

NewFPHSpedVAPIObjectsChatSystemChat instantiates a new FPHSpedVAPIObjectsChatSystemChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsChatSystemChatWithDefaults

`func NewFPHSpedVAPIObjectsChatSystemChatWithDefaults() *FPHSpedVAPIObjectsChatSystemChat`

NewFPHSpedVAPIObjectsChatSystemChatWithDefaults instantiates a new FPHSpedVAPIObjectsChatSystemChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetId(v int32)`

SetId sets Id field to given value.


### GetInitiator

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetInitiator() FPHSpedVAPIObjectsUsersUser`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetInitiatorOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetInitiator(v FPHSpedVAPIObjectsUsersUser)`

SetInitiator sets Initiator field to given value.


### SetInitiatorNil

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetInitiatorNil(b bool)`

 SetInitiatorNil sets the value for Initiator to be an explicit nil

### UnsetInitiator
`func (o *FPHSpedVAPIObjectsChatSystemChat) UnsetInitiator()`

UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
### GetStartTime

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetSubject

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetSubject(v string)`

SetSubject sets Subject field to given value.


### SetSubjectNil

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *FPHSpedVAPIObjectsChatSystemChat) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetLastMessageTimestamp

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetLastMessageTimestamp() time.Time`

GetLastMessageTimestamp returns the LastMessageTimestamp field if non-nil, zero value otherwise.

### GetLastMessageTimestampOk

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetLastMessageTimestampOk() (*time.Time, bool)`

GetLastMessageTimestampOk returns a tuple with the LastMessageTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageTimestamp

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetLastMessageTimestamp(v time.Time)`

SetLastMessageTimestamp sets LastMessageTimestamp field to given value.


### SetLastMessageTimestampNil

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetLastMessageTimestampNil(b bool)`

 SetLastMessageTimestampNil sets the value for LastMessageTimestamp to be an explicit nil

### UnsetLastMessageTimestamp
`func (o *FPHSpedVAPIObjectsChatSystemChat) UnsetLastMessageTimestamp()`

UnsetLastMessageTimestamp ensures that no value is present for LastMessageTimestamp, not even an explicit nil
### GetLastMessageString

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetLastMessageString() string`

GetLastMessageString returns the LastMessageString field if non-nil, zero value otherwise.

### GetLastMessageStringOk

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetLastMessageStringOk() (*string, bool)`

GetLastMessageStringOk returns a tuple with the LastMessageString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageString

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetLastMessageString(v string)`

SetLastMessageString sets LastMessageString field to given value.


### SetLastMessageStringNil

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetLastMessageStringNil(b bool)`

 SetLastMessageStringNil sets the value for LastMessageString to be an explicit nil

### UnsetLastMessageString
`func (o *FPHSpedVAPIObjectsChatSystemChat) UnsetLastMessageString()`

UnsetLastMessageString ensures that no value is present for LastMessageString, not even an explicit nil
### GetParticipants

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetParticipants() []FPHSpedVAPIObjectsChatSystemChatParticipant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetParticipantsOk() (*[]FPHSpedVAPIObjectsChatSystemChatParticipant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetParticipants(v []FPHSpedVAPIObjectsChatSystemChatParticipant)`

SetParticipants sets Participants field to given value.


### SetParticipantsNil

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetParticipantsNil(b bool)`

 SetParticipantsNil sets the value for Participants to be an explicit nil

### UnsetParticipants
`func (o *FPHSpedVAPIObjectsChatSystemChat) UnsetParticipants()`

UnsetParticipants ensures that no value is present for Participants, not even an explicit nil
### GetParticipantString

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetParticipantString() string`

GetParticipantString returns the ParticipantString field if non-nil, zero value otherwise.

### GetParticipantStringOk

`func (o *FPHSpedVAPIObjectsChatSystemChat) GetParticipantStringOk() (*string, bool)`

GetParticipantStringOk returns a tuple with the ParticipantString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantString

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetParticipantString(v string)`

SetParticipantString sets ParticipantString field to given value.


### SetParticipantStringNil

`func (o *FPHSpedVAPIObjectsChatSystemChat) SetParticipantStringNil(b bool)`

 SetParticipantStringNil sets the value for ParticipantString to be an explicit nil

### UnsetParticipantString
`func (o *FPHSpedVAPIObjectsChatSystemChat) UnsetParticipantString()`

UnsetParticipantString ensures that no value is present for ParticipantString, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


