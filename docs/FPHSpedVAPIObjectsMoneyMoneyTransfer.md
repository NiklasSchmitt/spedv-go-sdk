# FPHSpedVAPIObjectsMoneyMoneyTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**SenderAccount** | [**NullableFPHSpedVAPIObjectsMoneyLiteBankAccount**](FPHSpedVAPIObjectsMoneyLiteBankAccount.md) |  | [readonly] 
**Sender** | **NullableString** |  | [readonly] 
**ReceiverAccount** | [**NullableFPHSpedVAPIObjectsMoneyLiteBankAccount**](FPHSpedVAPIObjectsMoneyLiteBankAccount.md) |  | [readonly] 
**Receiver** | **NullableString** |  | [readonly] 
**Cause** | **NullableString** |  | [readonly] 
**CauseGroup** | **NullableString** |  | [readonly] 
**Value** | **float64** |  | [readonly] 
**Money** | **NullableString** |  | [readonly] 
**ResponsibleUser** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**Responsible** | **NullableString** |  | [readonly] 
**DateTime** | **time.Time** |  | [readonly] 
**Time** | **NullableString** |  | [readonly] 
**IsNegative** | **bool** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsMoneyMoneyTransfer

`func NewFPHSpedVAPIObjectsMoneyMoneyTransfer(id int32, senderAccount NullableFPHSpedVAPIObjectsMoneyLiteBankAccount, sender NullableString, receiverAccount NullableFPHSpedVAPIObjectsMoneyLiteBankAccount, receiver NullableString, cause NullableString, causeGroup NullableString, value float64, money NullableString, responsibleUser NullableFPHSpedVAPIObjectsUsersUser, responsible NullableString, dateTime time.Time, time NullableString, isNegative bool, ) *FPHSpedVAPIObjectsMoneyMoneyTransfer`

NewFPHSpedVAPIObjectsMoneyMoneyTransfer instantiates a new FPHSpedVAPIObjectsMoneyMoneyTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsMoneyMoneyTransferWithDefaults

`func NewFPHSpedVAPIObjectsMoneyMoneyTransferWithDefaults() *FPHSpedVAPIObjectsMoneyMoneyTransfer`

NewFPHSpedVAPIObjectsMoneyMoneyTransferWithDefaults instantiates a new FPHSpedVAPIObjectsMoneyMoneyTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetId(v int32)`

SetId sets Id field to given value.


### GetSenderAccount

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetSenderAccount() FPHSpedVAPIObjectsMoneyLiteBankAccount`

GetSenderAccount returns the SenderAccount field if non-nil, zero value otherwise.

### GetSenderAccountOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetSenderAccountOk() (*FPHSpedVAPIObjectsMoneyLiteBankAccount, bool)`

GetSenderAccountOk returns a tuple with the SenderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccount

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetSenderAccount(v FPHSpedVAPIObjectsMoneyLiteBankAccount)`

SetSenderAccount sets SenderAccount field to given value.


### SetSenderAccountNil

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetSenderAccountNil(b bool)`

 SetSenderAccountNil sets the value for SenderAccount to be an explicit nil

### UnsetSenderAccount
`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) UnsetSenderAccount()`

UnsetSenderAccount ensures that no value is present for SenderAccount, not even an explicit nil
### GetSender

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetSender(v string)`

SetSender sets Sender field to given value.


### SetSenderNil

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetReceiverAccount

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetReceiverAccount() FPHSpedVAPIObjectsMoneyLiteBankAccount`

GetReceiverAccount returns the ReceiverAccount field if non-nil, zero value otherwise.

### GetReceiverAccountOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetReceiverAccountOk() (*FPHSpedVAPIObjectsMoneyLiteBankAccount, bool)`

GetReceiverAccountOk returns a tuple with the ReceiverAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverAccount

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetReceiverAccount(v FPHSpedVAPIObjectsMoneyLiteBankAccount)`

SetReceiverAccount sets ReceiverAccount field to given value.


### SetReceiverAccountNil

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetReceiverAccountNil(b bool)`

 SetReceiverAccountNil sets the value for ReceiverAccount to be an explicit nil

### UnsetReceiverAccount
`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) UnsetReceiverAccount()`

UnsetReceiverAccount ensures that no value is present for ReceiverAccount, not even an explicit nil
### GetReceiver

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### SetReceiverNil

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetReceiverNil(b bool)`

 SetReceiverNil sets the value for Receiver to be an explicit nil

### UnsetReceiver
`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) UnsetReceiver()`

UnsetReceiver ensures that no value is present for Receiver, not even an explicit nil
### GetCause

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetCause(v string)`

SetCause sets Cause field to given value.


### SetCauseNil

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetCauseNil(b bool)`

 SetCauseNil sets the value for Cause to be an explicit nil

### UnsetCause
`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) UnsetCause()`

UnsetCause ensures that no value is present for Cause, not even an explicit nil
### GetCauseGroup

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetCauseGroup() string`

GetCauseGroup returns the CauseGroup field if non-nil, zero value otherwise.

### GetCauseGroupOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetCauseGroupOk() (*string, bool)`

GetCauseGroupOk returns a tuple with the CauseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseGroup

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetCauseGroup(v string)`

SetCauseGroup sets CauseGroup field to given value.


### SetCauseGroupNil

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetCauseGroupNil(b bool)`

 SetCauseGroupNil sets the value for CauseGroup to be an explicit nil

### UnsetCauseGroup
`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) UnsetCauseGroup()`

UnsetCauseGroup ensures that no value is present for CauseGroup, not even an explicit nil
### GetValue

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetValue(v float64)`

SetValue sets Value field to given value.


### GetMoney

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetMoney() string`

GetMoney returns the Money field if non-nil, zero value otherwise.

### GetMoneyOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetMoneyOk() (*string, bool)`

GetMoneyOk returns a tuple with the Money field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoney

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetMoney(v string)`

SetMoney sets Money field to given value.


### SetMoneyNil

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetMoneyNil(b bool)`

 SetMoneyNil sets the value for Money to be an explicit nil

### UnsetMoney
`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) UnsetMoney()`

UnsetMoney ensures that no value is present for Money, not even an explicit nil
### GetResponsibleUser

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetResponsibleUser() FPHSpedVAPIObjectsUsersUser`

GetResponsibleUser returns the ResponsibleUser field if non-nil, zero value otherwise.

### GetResponsibleUserOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetResponsibleUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetResponsibleUserOk returns a tuple with the ResponsibleUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleUser

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetResponsibleUser(v FPHSpedVAPIObjectsUsersUser)`

SetResponsibleUser sets ResponsibleUser field to given value.


### SetResponsibleUserNil

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetResponsibleUserNil(b bool)`

 SetResponsibleUserNil sets the value for ResponsibleUser to be an explicit nil

### UnsetResponsibleUser
`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) UnsetResponsibleUser()`

UnsetResponsibleUser ensures that no value is present for ResponsibleUser, not even an explicit nil
### GetResponsible

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetResponsible() string`

GetResponsible returns the Responsible field if non-nil, zero value otherwise.

### GetResponsibleOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetResponsibleOk() (*string, bool)`

GetResponsibleOk returns a tuple with the Responsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsible

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetResponsible(v string)`

SetResponsible sets Responsible field to given value.


### SetResponsibleNil

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetResponsibleNil(b bool)`

 SetResponsibleNil sets the value for Responsible to be an explicit nil

### UnsetResponsible
`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) UnsetResponsible()`

UnsetResponsible ensures that no value is present for Responsible, not even an explicit nil
### GetDateTime

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.


### GetTime

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetTime(v string)`

SetTime sets Time field to given value.


### SetTimeNil

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil
### GetIsNegative

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetIsNegative() bool`

GetIsNegative returns the IsNegative field if non-nil, zero value otherwise.

### GetIsNegativeOk

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) GetIsNegativeOk() (*bool, bool)`

GetIsNegativeOk returns a tuple with the IsNegative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNegative

`func (o *FPHSpedVAPIObjectsMoneyMoneyTransfer) SetIsNegative(v bool)`

SetIsNegative sets IsNegative field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


