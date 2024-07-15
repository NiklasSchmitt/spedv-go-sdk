# FPHSpedVAPIObjectsSpeditionsSpedTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Comment** | **NullableString** |  | [readonly] 
**Moneyamount** | **int32** |  | [readonly] 
**Amount** | **float64** |  | [readonly] 
**StartDate** | **time.Time** |  | [readonly] 
**DoneDate** | **time.Time** |  | [readonly] 
**Bonus** | **int32** |  | [readonly] 
**BonusSum** | **float64** |  | [readonly] 
**Users** | [**[]FPHSpedVAPIObjectsSpeditionsSpedTargetUser**](FPHSpedVAPIObjectsSpeditionsSpedTargetUser.md) |  | [readonly] 
**IgnoredUsers** | [**[]FPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**Type** | [**FPHSpedVAPIEnumsSpedTargetType**](FPHSpedVAPIEnumsSpedTargetType.md) |   0 &#x3D; Value  1 &#x3D; Distance  2 &#x3D; WeightDistance | [readonly] 
**KontorHandling** | [**FPHSpedVAPIEnumsSpedTargetKontorHandling**](FPHSpedVAPIEnumsSpedTargetKontorHandling.md) |   0 &#x3D; AllTasks  1 &#x3D; OnlyKontorTasks  2 &#x3D; NoKontorTasks | [readonly] 
**OmsiHandling** | [**FPHSpedVAPIEnumsSpedTargetOMSIHandling**](FPHSpedVAPIEnumsSpedTargetOMSIHandling.md) |   0 &#x3D; AllTasks  1 &#x3D; OnlyOMSITasks  2 &#x3D; NoOMSITasks | [readonly] 
**Alrreached** | **int32** |  | [readonly] 
**AmountReachedSum** | **float64** |  | [readonly] 
**Reached** | **bool** |  | [readonly] 
**IsReached** | **bool** |  | [readonly] 
**IsActive** | **bool** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsSpedTarget

`func NewFPHSpedVAPIObjectsSpeditionsSpedTarget(id int32, comment NullableString, moneyamount int32, amount float64, startDate time.Time, doneDate time.Time, bonus int32, bonusSum float64, users []FPHSpedVAPIObjectsSpeditionsSpedTargetUser, ignoredUsers []FPHSpedVAPIObjectsUsersUser, type_ FPHSpedVAPIEnumsSpedTargetType, kontorHandling FPHSpedVAPIEnumsSpedTargetKontorHandling, omsiHandling FPHSpedVAPIEnumsSpedTargetOMSIHandling, alrreached int32, amountReachedSum float64, reached bool, isReached bool, isActive bool, ) *FPHSpedVAPIObjectsSpeditionsSpedTarget`

NewFPHSpedVAPIObjectsSpeditionsSpedTarget instantiates a new FPHSpedVAPIObjectsSpeditionsSpedTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsSpedTargetWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsSpedTargetWithDefaults() *FPHSpedVAPIObjectsSpeditionsSpedTarget`

NewFPHSpedVAPIObjectsSpeditionsSpedTargetWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsSpedTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetId(v int32)`

SetId sets Id field to given value.


### GetComment

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetMoneyamount

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetMoneyamount() int32`

GetMoneyamount returns the Moneyamount field if non-nil, zero value otherwise.

### GetMoneyamountOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetMoneyamountOk() (*int32, bool)`

GetMoneyamountOk returns a tuple with the Moneyamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyamount

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetMoneyamount(v int32)`

SetMoneyamount sets Moneyamount field to given value.


### GetAmount

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetStartDate

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetDoneDate

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetDoneDate() time.Time`

GetDoneDate returns the DoneDate field if non-nil, zero value otherwise.

### GetDoneDateOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetDoneDateOk() (*time.Time, bool)`

GetDoneDateOk returns a tuple with the DoneDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneDate

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetDoneDate(v time.Time)`

SetDoneDate sets DoneDate field to given value.


### GetBonus

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetBonus() int32`

GetBonus returns the Bonus field if non-nil, zero value otherwise.

### GetBonusOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetBonusOk() (*int32, bool)`

GetBonusOk returns a tuple with the Bonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonus

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetBonus(v int32)`

SetBonus sets Bonus field to given value.


### GetBonusSum

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetBonusSum() float64`

GetBonusSum returns the BonusSum field if non-nil, zero value otherwise.

### GetBonusSumOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetBonusSumOk() (*float64, bool)`

GetBonusSumOk returns a tuple with the BonusSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonusSum

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetBonusSum(v float64)`

SetBonusSum sets BonusSum field to given value.


### GetUsers

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetUsers() []FPHSpedVAPIObjectsSpeditionsSpedTargetUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetUsersOk() (*[]FPHSpedVAPIObjectsSpeditionsSpedTargetUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetUsers(v []FPHSpedVAPIObjectsSpeditionsSpedTargetUser)`

SetUsers sets Users field to given value.


### SetUsersNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetIgnoredUsers

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetIgnoredUsers() []FPHSpedVAPIObjectsUsersUser`

GetIgnoredUsers returns the IgnoredUsers field if non-nil, zero value otherwise.

### GetIgnoredUsersOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetIgnoredUsersOk() (*[]FPHSpedVAPIObjectsUsersUser, bool)`

GetIgnoredUsersOk returns a tuple with the IgnoredUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredUsers

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetIgnoredUsers(v []FPHSpedVAPIObjectsUsersUser)`

SetIgnoredUsers sets IgnoredUsers field to given value.


### SetIgnoredUsersNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetIgnoredUsersNil(b bool)`

 SetIgnoredUsersNil sets the value for IgnoredUsers to be an explicit nil

### UnsetIgnoredUsers
`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) UnsetIgnoredUsers()`

UnsetIgnoredUsers ensures that no value is present for IgnoredUsers, not even an explicit nil
### GetType

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetType() FPHSpedVAPIEnumsSpedTargetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetTypeOk() (*FPHSpedVAPIEnumsSpedTargetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetType(v FPHSpedVAPIEnumsSpedTargetType)`

SetType sets Type field to given value.


### GetKontorHandling

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetKontorHandling() FPHSpedVAPIEnumsSpedTargetKontorHandling`

GetKontorHandling returns the KontorHandling field if non-nil, zero value otherwise.

### GetKontorHandlingOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetKontorHandlingOk() (*FPHSpedVAPIEnumsSpedTargetKontorHandling, bool)`

GetKontorHandlingOk returns a tuple with the KontorHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKontorHandling

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetKontorHandling(v FPHSpedVAPIEnumsSpedTargetKontorHandling)`

SetKontorHandling sets KontorHandling field to given value.


### GetOmsiHandling

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetOmsiHandling() FPHSpedVAPIEnumsSpedTargetOMSIHandling`

GetOmsiHandling returns the OmsiHandling field if non-nil, zero value otherwise.

### GetOmsiHandlingOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetOmsiHandlingOk() (*FPHSpedVAPIEnumsSpedTargetOMSIHandling, bool)`

GetOmsiHandlingOk returns a tuple with the OmsiHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmsiHandling

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetOmsiHandling(v FPHSpedVAPIEnumsSpedTargetOMSIHandling)`

SetOmsiHandling sets OmsiHandling field to given value.


### GetAlrreached

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetAlrreached() int32`

GetAlrreached returns the Alrreached field if non-nil, zero value otherwise.

### GetAlrreachedOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetAlrreachedOk() (*int32, bool)`

GetAlrreachedOk returns a tuple with the Alrreached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlrreached

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetAlrreached(v int32)`

SetAlrreached sets Alrreached field to given value.


### GetAmountReachedSum

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetAmountReachedSum() float64`

GetAmountReachedSum returns the AmountReachedSum field if non-nil, zero value otherwise.

### GetAmountReachedSumOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetAmountReachedSumOk() (*float64, bool)`

GetAmountReachedSumOk returns a tuple with the AmountReachedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountReachedSum

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetAmountReachedSum(v float64)`

SetAmountReachedSum sets AmountReachedSum field to given value.


### GetReached

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetReached() bool`

GetReached returns the Reached field if non-nil, zero value otherwise.

### GetReachedOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetReachedOk() (*bool, bool)`

GetReachedOk returns a tuple with the Reached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReached

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetReached(v bool)`

SetReached sets Reached field to given value.


### GetIsReached

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetIsReached() bool`

GetIsReached returns the IsReached field if non-nil, zero value otherwise.

### GetIsReachedOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetIsReachedOk() (*bool, bool)`

GetIsReachedOk returns a tuple with the IsReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReached

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetIsReached(v bool)`

SetIsReached sets IsReached field to given value.


### GetIsActive

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *FPHSpedVAPIObjectsSpeditionsSpedTarget) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


