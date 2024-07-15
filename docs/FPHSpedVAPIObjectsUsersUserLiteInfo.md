# FPHSpedVAPIObjectsUsersUserLiteInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**UserName** | **NullableString** |  | [readonly] 
**Rank** | [**NullableFPHSpedVAPIObjectsSpeditionsLiteRank**](FPHSpedVAPIObjectsSpeditionsLiteRank.md) |  | [readonly] 
**LastTaskTime** | **time.Time** |  | [readonly] 
**MainBankAccount** | [**NullableFPHSpedVAPIObjectsMoneyLiteBankAccount**](FPHSpedVAPIObjectsMoneyLiteBankAccount.md) |  | [readonly] 
**ProfilePicURL** | **NullableString** |  | [readonly] 
**Spedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**InSped** | **bool** |  | [readonly] 
**FormatName** | **NullableString** |  | [readonly] 
**SpeditionJoinDate** | **NullableTime** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsUsersUserLiteInfo

`func NewFPHSpedVAPIObjectsUsersUserLiteInfo(id int32, userName NullableString, rank NullableFPHSpedVAPIObjectsSpeditionsLiteRank, lastTaskTime time.Time, mainBankAccount NullableFPHSpedVAPIObjectsMoneyLiteBankAccount, profilePicURL NullableString, spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, inSped bool, formatName NullableString, speditionJoinDate NullableTime, ) *FPHSpedVAPIObjectsUsersUserLiteInfo`

NewFPHSpedVAPIObjectsUsersUserLiteInfo instantiates a new FPHSpedVAPIObjectsUsersUserLiteInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsUsersUserLiteInfoWithDefaults

`func NewFPHSpedVAPIObjectsUsersUserLiteInfoWithDefaults() *FPHSpedVAPIObjectsUsersUserLiteInfo`

NewFPHSpedVAPIObjectsUsersUserLiteInfoWithDefaults instantiates a new FPHSpedVAPIObjectsUsersUserLiteInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetId(v int32)`

SetId sets Id field to given value.


### GetUserName

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.


### SetUserNameNil

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetRank

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetRank() FPHSpedVAPIObjectsSpeditionsLiteRank`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetRankOk() (*FPHSpedVAPIObjectsSpeditionsLiteRank, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetRank(v FPHSpedVAPIObjectsSpeditionsLiteRank)`

SetRank sets Rank field to given value.


### SetRankNil

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetRankNil(b bool)`

 SetRankNil sets the value for Rank to be an explicit nil

### UnsetRank
`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) UnsetRank()`

UnsetRank ensures that no value is present for Rank, not even an explicit nil
### GetLastTaskTime

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetLastTaskTime() time.Time`

GetLastTaskTime returns the LastTaskTime field if non-nil, zero value otherwise.

### GetLastTaskTimeOk

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetLastTaskTimeOk() (*time.Time, bool)`

GetLastTaskTimeOk returns a tuple with the LastTaskTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTaskTime

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetLastTaskTime(v time.Time)`

SetLastTaskTime sets LastTaskTime field to given value.


### GetMainBankAccount

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetMainBankAccount() FPHSpedVAPIObjectsMoneyLiteBankAccount`

GetMainBankAccount returns the MainBankAccount field if non-nil, zero value otherwise.

### GetMainBankAccountOk

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetMainBankAccountOk() (*FPHSpedVAPIObjectsMoneyLiteBankAccount, bool)`

GetMainBankAccountOk returns a tuple with the MainBankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainBankAccount

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetMainBankAccount(v FPHSpedVAPIObjectsMoneyLiteBankAccount)`

SetMainBankAccount sets MainBankAccount field to given value.


### SetMainBankAccountNil

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetMainBankAccountNil(b bool)`

 SetMainBankAccountNil sets the value for MainBankAccount to be an explicit nil

### UnsetMainBankAccount
`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) UnsetMainBankAccount()`

UnsetMainBankAccount ensures that no value is present for MainBankAccount, not even an explicit nil
### GetProfilePicURL

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetProfilePicURL() string`

GetProfilePicURL returns the ProfilePicURL field if non-nil, zero value otherwise.

### GetProfilePicURLOk

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetProfilePicURLOk() (*string, bool)`

GetProfilePicURLOk returns a tuple with the ProfilePicURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicURL

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetProfilePicURL(v string)`

SetProfilePicURL sets ProfilePicURL field to given value.


### SetProfilePicURLNil

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetProfilePicURLNil(b bool)`

 SetProfilePicURLNil sets the value for ProfilePicURL to be an explicit nil

### UnsetProfilePicURL
`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) UnsetProfilePicURL()`

UnsetProfilePicURL ensures that no value is present for ProfilePicURL, not even an explicit nil
### GetSpedition

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetSpedition returns the Spedition field if non-nil, zero value otherwise.

### GetSpeditionOk

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetSpeditionOk returns a tuple with the Spedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpedition

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetSpedition sets Spedition field to given value.


### SetSpeditionNil

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetSpeditionNil(b bool)`

 SetSpeditionNil sets the value for Spedition to be an explicit nil

### UnsetSpedition
`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) UnsetSpedition()`

UnsetSpedition ensures that no value is present for Spedition, not even an explicit nil
### GetInSped

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetInSped() bool`

GetInSped returns the InSped field if non-nil, zero value otherwise.

### GetInSpedOk

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetInSpedOk() (*bool, bool)`

GetInSpedOk returns a tuple with the InSped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInSped

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetInSped(v bool)`

SetInSped sets InSped field to given value.


### GetFormatName

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetFormatName() string`

GetFormatName returns the FormatName field if non-nil, zero value otherwise.

### GetFormatNameOk

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetFormatNameOk() (*string, bool)`

GetFormatNameOk returns a tuple with the FormatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatName

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetFormatName(v string)`

SetFormatName sets FormatName field to given value.


### SetFormatNameNil

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetFormatNameNil(b bool)`

 SetFormatNameNil sets the value for FormatName to be an explicit nil

### UnsetFormatName
`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) UnsetFormatName()`

UnsetFormatName ensures that no value is present for FormatName, not even an explicit nil
### GetSpeditionJoinDate

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetSpeditionJoinDate() time.Time`

GetSpeditionJoinDate returns the SpeditionJoinDate field if non-nil, zero value otherwise.

### GetSpeditionJoinDateOk

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) GetSpeditionJoinDateOk() (*time.Time, bool)`

GetSpeditionJoinDateOk returns a tuple with the SpeditionJoinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeditionJoinDate

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetSpeditionJoinDate(v time.Time)`

SetSpeditionJoinDate sets SpeditionJoinDate field to given value.


### SetSpeditionJoinDateNil

`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) SetSpeditionJoinDateNil(b bool)`

 SetSpeditionJoinDateNil sets the value for SpeditionJoinDate to be an explicit nil

### UnsetSpeditionJoinDate
`func (o *FPHSpedVAPIObjectsUsersUserLiteInfo) UnsetSpeditionJoinDate()`

UnsetSpeditionJoinDate ensures that no value is present for SpeditionJoinDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


