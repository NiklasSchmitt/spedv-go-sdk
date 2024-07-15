# FPHSpedVAPIObjectsUsersUserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**SteamID** | **int64** |  | [readonly] 
**IncomeSum** | **float64** |  | [readonly] 
**DamageSum** | **float64** |  | [readonly] 
**KmSum** | **float64** |  | [readonly] 
**RealName** | **NullableString** |  | [readonly] 
**AboutMe** | **NullableString** |  | [readonly] 
**Birthday** | **time.Time** |  | [readonly] 
**SteamImgURL** | **NullableString** |  | [readonly] 
**Residence** | **NullableString** |  | [readonly] 
**UserName** | **NullableString** |  | [readonly] 
**ChangeLog** | [**[]FPHSpedVAPIObjectsUsersUserChangeLogEntry**](FPHSpedVAPIObjectsUsersUserChangeLogEntry.md) |  | [readonly] 
**Spedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**InSped** | **bool** |  | [readonly] 
**FormatName** | **NullableString** |  | [readonly] 
**LastTaskTime** | **time.Time** |  | [readonly] 
**MainBankAccount** | [**NullableFPHSpedVAPIObjectsMoneyLiteBankAccount**](FPHSpedVAPIObjectsMoneyLiteBankAccount.md) |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsUsersUserProfile

`func NewFPHSpedVAPIObjectsUsersUserProfile(id int32, steamID int64, incomeSum float64, damageSum float64, kmSum float64, realName NullableString, aboutMe NullableString, birthday time.Time, steamImgURL NullableString, residence NullableString, userName NullableString, changeLog []FPHSpedVAPIObjectsUsersUserChangeLogEntry, spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, inSped bool, formatName NullableString, lastTaskTime time.Time, mainBankAccount NullableFPHSpedVAPIObjectsMoneyLiteBankAccount, ) *FPHSpedVAPIObjectsUsersUserProfile`

NewFPHSpedVAPIObjectsUsersUserProfile instantiates a new FPHSpedVAPIObjectsUsersUserProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsUsersUserProfileWithDefaults

`func NewFPHSpedVAPIObjectsUsersUserProfileWithDefaults() *FPHSpedVAPIObjectsUsersUserProfile`

NewFPHSpedVAPIObjectsUsersUserProfileWithDefaults instantiates a new FPHSpedVAPIObjectsUsersUserProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetId(v int32)`

SetId sets Id field to given value.


### GetSteamID

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetSteamID() int64`

GetSteamID returns the SteamID field if non-nil, zero value otherwise.

### GetSteamIDOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetSteamIDOk() (*int64, bool)`

GetSteamIDOk returns a tuple with the SteamID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamID

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetSteamID(v int64)`

SetSteamID sets SteamID field to given value.


### GetIncomeSum

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetIncomeSum() float64`

GetIncomeSum returns the IncomeSum field if non-nil, zero value otherwise.

### GetIncomeSumOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetIncomeSumOk() (*float64, bool)`

GetIncomeSumOk returns a tuple with the IncomeSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeSum

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetIncomeSum(v float64)`

SetIncomeSum sets IncomeSum field to given value.


### GetDamageSum

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetDamageSum() float64`

GetDamageSum returns the DamageSum field if non-nil, zero value otherwise.

### GetDamageSumOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetDamageSumOk() (*float64, bool)`

GetDamageSumOk returns a tuple with the DamageSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageSum

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetDamageSum(v float64)`

SetDamageSum sets DamageSum field to given value.


### GetKmSum

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetKmSum() float64`

GetKmSum returns the KmSum field if non-nil, zero value otherwise.

### GetKmSumOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetKmSumOk() (*float64, bool)`

GetKmSumOk returns a tuple with the KmSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmSum

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetKmSum(v float64)`

SetKmSum sets KmSum field to given value.


### GetRealName

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetRealName(v string)`

SetRealName sets RealName field to given value.


### SetRealNameNil

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetRealNameNil(b bool)`

 SetRealNameNil sets the value for RealName to be an explicit nil

### UnsetRealName
`func (o *FPHSpedVAPIObjectsUsersUserProfile) UnsetRealName()`

UnsetRealName ensures that no value is present for RealName, not even an explicit nil
### GetAboutMe

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetAboutMe() string`

GetAboutMe returns the AboutMe field if non-nil, zero value otherwise.

### GetAboutMeOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetAboutMeOk() (*string, bool)`

GetAboutMeOk returns a tuple with the AboutMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboutMe

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetAboutMe(v string)`

SetAboutMe sets AboutMe field to given value.


### SetAboutMeNil

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetAboutMeNil(b bool)`

 SetAboutMeNil sets the value for AboutMe to be an explicit nil

### UnsetAboutMe
`func (o *FPHSpedVAPIObjectsUsersUserProfile) UnsetAboutMe()`

UnsetAboutMe ensures that no value is present for AboutMe, not even an explicit nil
### GetBirthday

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.


### GetSteamImgURL

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetSteamImgURL() string`

GetSteamImgURL returns the SteamImgURL field if non-nil, zero value otherwise.

### GetSteamImgURLOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetSteamImgURLOk() (*string, bool)`

GetSteamImgURLOk returns a tuple with the SteamImgURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamImgURL

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetSteamImgURL(v string)`

SetSteamImgURL sets SteamImgURL field to given value.


### SetSteamImgURLNil

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetSteamImgURLNil(b bool)`

 SetSteamImgURLNil sets the value for SteamImgURL to be an explicit nil

### UnsetSteamImgURL
`func (o *FPHSpedVAPIObjectsUsersUserProfile) UnsetSteamImgURL()`

UnsetSteamImgURL ensures that no value is present for SteamImgURL, not even an explicit nil
### GetResidence

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetResidence() string`

GetResidence returns the Residence field if non-nil, zero value otherwise.

### GetResidenceOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetResidenceOk() (*string, bool)`

GetResidenceOk returns a tuple with the Residence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidence

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetResidence(v string)`

SetResidence sets Residence field to given value.


### SetResidenceNil

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetResidenceNil(b bool)`

 SetResidenceNil sets the value for Residence to be an explicit nil

### UnsetResidence
`func (o *FPHSpedVAPIObjectsUsersUserProfile) UnsetResidence()`

UnsetResidence ensures that no value is present for Residence, not even an explicit nil
### GetUserName

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetUserName(v string)`

SetUserName sets UserName field to given value.


### SetUserNameNil

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *FPHSpedVAPIObjectsUsersUserProfile) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetChangeLog

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetChangeLog() []FPHSpedVAPIObjectsUsersUserChangeLogEntry`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetChangeLogOk() (*[]FPHSpedVAPIObjectsUsersUserChangeLogEntry, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetChangeLog(v []FPHSpedVAPIObjectsUsersUserChangeLogEntry)`

SetChangeLog sets ChangeLog field to given value.


### SetChangeLogNil

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetChangeLogNil(b bool)`

 SetChangeLogNil sets the value for ChangeLog to be an explicit nil

### UnsetChangeLog
`func (o *FPHSpedVAPIObjectsUsersUserProfile) UnsetChangeLog()`

UnsetChangeLog ensures that no value is present for ChangeLog, not even an explicit nil
### GetSpedition

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetSpedition returns the Spedition field if non-nil, zero value otherwise.

### GetSpeditionOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetSpeditionOk returns a tuple with the Spedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpedition

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetSpedition sets Spedition field to given value.


### SetSpeditionNil

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetSpeditionNil(b bool)`

 SetSpeditionNil sets the value for Spedition to be an explicit nil

### UnsetSpedition
`func (o *FPHSpedVAPIObjectsUsersUserProfile) UnsetSpedition()`

UnsetSpedition ensures that no value is present for Spedition, not even an explicit nil
### GetInSped

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetInSped() bool`

GetInSped returns the InSped field if non-nil, zero value otherwise.

### GetInSpedOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetInSpedOk() (*bool, bool)`

GetInSpedOk returns a tuple with the InSped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInSped

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetInSped(v bool)`

SetInSped sets InSped field to given value.


### GetFormatName

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetFormatName() string`

GetFormatName returns the FormatName field if non-nil, zero value otherwise.

### GetFormatNameOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetFormatNameOk() (*string, bool)`

GetFormatNameOk returns a tuple with the FormatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatName

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetFormatName(v string)`

SetFormatName sets FormatName field to given value.


### SetFormatNameNil

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetFormatNameNil(b bool)`

 SetFormatNameNil sets the value for FormatName to be an explicit nil

### UnsetFormatName
`func (o *FPHSpedVAPIObjectsUsersUserProfile) UnsetFormatName()`

UnsetFormatName ensures that no value is present for FormatName, not even an explicit nil
### GetLastTaskTime

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetLastTaskTime() time.Time`

GetLastTaskTime returns the LastTaskTime field if non-nil, zero value otherwise.

### GetLastTaskTimeOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetLastTaskTimeOk() (*time.Time, bool)`

GetLastTaskTimeOk returns a tuple with the LastTaskTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTaskTime

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetLastTaskTime(v time.Time)`

SetLastTaskTime sets LastTaskTime field to given value.


### GetMainBankAccount

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetMainBankAccount() FPHSpedVAPIObjectsMoneyLiteBankAccount`

GetMainBankAccount returns the MainBankAccount field if non-nil, zero value otherwise.

### GetMainBankAccountOk

`func (o *FPHSpedVAPIObjectsUsersUserProfile) GetMainBankAccountOk() (*FPHSpedVAPIObjectsMoneyLiteBankAccount, bool)`

GetMainBankAccountOk returns a tuple with the MainBankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainBankAccount

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetMainBankAccount(v FPHSpedVAPIObjectsMoneyLiteBankAccount)`

SetMainBankAccount sets MainBankAccount field to given value.


### SetMainBankAccountNil

`func (o *FPHSpedVAPIObjectsUsersUserProfile) SetMainBankAccountNil(b bool)`

 SetMainBankAccountNil sets the value for MainBankAccount to be an explicit nil

### UnsetMainBankAccount
`func (o *FPHSpedVAPIObjectsUsersUserProfile) UnsetMainBankAccount()`

UnsetMainBankAccount ensures that no value is present for MainBankAccount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


