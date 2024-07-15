# FPHSpedVAPIObjectsUsersFullUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**SteamID** | **int64** |  | [readonly] 
**AuthCode** | **NullableString** |  | [readonly] 
**PlusUntil** | **time.Time** |  | [readonly] 
**PlusActive** | **bool** |  | [readonly] 
**BannedTill** | **time.Time** |  | [readonly] 
**IsBanned** | **bool** |  | [readonly] 
**BannedReason** | **NullableString** |  | [readonly] 
**EMail** | **NullableString** |  | [readonly] 
**RealName** | **NullableString** |  | [readonly] 
**AboutMe** | **NullableString** |  | [readonly] 
**Birthday** | **time.Time** |  | [readonly] 
**SteamImgURL** | **NullableString** |  | [readonly] 
**Residence** | **NullableString** |  | [readonly] 
**UserName** | **NullableString** |  | [readonly] 
**KontorActivated** | **bool** |  | [readonly] 
**ChangelogVisible** | **bool** |  | [readonly] 
**DeletionFrist** | **int32** |  | [readonly] 
**TfaActivated** | **bool** |  | [readonly] 
**IsAlphaTester** | **bool** |  | [readonly] 
**IsTrustedUser** | **bool** |  | [readonly] 
**SystemAdminPermissions** | [**FPHSpedVAPIAdminPermissionEnum**](FPHSpedVAPIAdminPermissionEnum.md) |   0 &#x3D; None  1 &#x3D; SeeLiveData  2 &#x3D; HandleScreenshots  4 &#x3D; HandleTaskCheck  8 &#x3D; TicketSystem  16 &#x3D; ViewData  32 &#x3D; ChangeData  64 &#x3D; BanUser  128 &#x3D; MapImport  256 &#x3D; KontorManagement  512 &#x3D; IsManagement | [readonly] 
**SystemAdminRankName** | **NullableString** |  | [readonly] 
**TwitchName** | **NullableString** |  | [readonly] 
**TwitchID** | **NullableInt32** |  | [readonly] 
**DiscordUID** | **NullableInt64** |  | [readonly] 
**Spedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**Rank** | [**NullableFPHSpedVAPIObjectsSpeditionsRank**](FPHSpedVAPIObjectsSpeditionsRank.md) |  | [readonly] 
**Achievements** | **map[string]interface{}** |  | [readonly] 
**InSped** | **bool** |  | [readonly] 
**FormatName** | **NullableString** |  | [readonly] 
**LastTaskTime** | **time.Time** |  | [readonly] 
**IsInCooldown** | **bool** |  | [readonly] 
**MainBankAccount** | [**NullableFPHSpedVAPIObjectsMoneyLiteBankAccount**](FPHSpedVAPIObjectsMoneyLiteBankAccount.md) |  | [readonly] 
**SpeditionJoinDate** | **NullableTime** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsUsersFullUserInfo

`func NewFPHSpedVAPIObjectsUsersFullUserInfo(id int32, steamID int64, authCode NullableString, plusUntil time.Time, plusActive bool, bannedTill time.Time, isBanned bool, bannedReason NullableString, eMail NullableString, realName NullableString, aboutMe NullableString, birthday time.Time, steamImgURL NullableString, residence NullableString, userName NullableString, kontorActivated bool, changelogVisible bool, deletionFrist int32, tfaActivated bool, isAlphaTester bool, isTrustedUser bool, systemAdminPermissions FPHSpedVAPIAdminPermissionEnum, systemAdminRankName NullableString, twitchName NullableString, twitchID NullableInt32, discordUID NullableInt64, spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, rank NullableFPHSpedVAPIObjectsSpeditionsRank, achievements map[string]interface{}, inSped bool, formatName NullableString, lastTaskTime time.Time, isInCooldown bool, mainBankAccount NullableFPHSpedVAPIObjectsMoneyLiteBankAccount, speditionJoinDate NullableTime, ) *FPHSpedVAPIObjectsUsersFullUserInfo`

NewFPHSpedVAPIObjectsUsersFullUserInfo instantiates a new FPHSpedVAPIObjectsUsersFullUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsUsersFullUserInfoWithDefaults

`func NewFPHSpedVAPIObjectsUsersFullUserInfoWithDefaults() *FPHSpedVAPIObjectsUsersFullUserInfo`

NewFPHSpedVAPIObjectsUsersFullUserInfoWithDefaults instantiates a new FPHSpedVAPIObjectsUsersFullUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetId(v int32)`

SetId sets Id field to given value.


### GetSteamID

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetSteamID() int64`

GetSteamID returns the SteamID field if non-nil, zero value otherwise.

### GetSteamIDOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetSteamIDOk() (*int64, bool)`

GetSteamIDOk returns a tuple with the SteamID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamID

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetSteamID(v int64)`

SetSteamID sets SteamID field to given value.


### GetAuthCode

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### SetAuthCodeNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetAuthCodeNil(b bool)`

 SetAuthCodeNil sets the value for AuthCode to be an explicit nil

### UnsetAuthCode
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetAuthCode()`

UnsetAuthCode ensures that no value is present for AuthCode, not even an explicit nil
### GetPlusUntil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetPlusUntil() time.Time`

GetPlusUntil returns the PlusUntil field if non-nil, zero value otherwise.

### GetPlusUntilOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetPlusUntilOk() (*time.Time, bool)`

GetPlusUntilOk returns a tuple with the PlusUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlusUntil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetPlusUntil(v time.Time)`

SetPlusUntil sets PlusUntil field to given value.


### GetPlusActive

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetPlusActive() bool`

GetPlusActive returns the PlusActive field if non-nil, zero value otherwise.

### GetPlusActiveOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetPlusActiveOk() (*bool, bool)`

GetPlusActiveOk returns a tuple with the PlusActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlusActive

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetPlusActive(v bool)`

SetPlusActive sets PlusActive field to given value.


### GetBannedTill

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetBannedTill() time.Time`

GetBannedTill returns the BannedTill field if non-nil, zero value otherwise.

### GetBannedTillOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetBannedTillOk() (*time.Time, bool)`

GetBannedTillOk returns a tuple with the BannedTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedTill

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetBannedTill(v time.Time)`

SetBannedTill sets BannedTill field to given value.


### GetIsBanned

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetIsBanned() bool`

GetIsBanned returns the IsBanned field if non-nil, zero value otherwise.

### GetIsBannedOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetIsBannedOk() (*bool, bool)`

GetIsBannedOk returns a tuple with the IsBanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBanned

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetIsBanned(v bool)`

SetIsBanned sets IsBanned field to given value.


### GetBannedReason

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetBannedReason() string`

GetBannedReason returns the BannedReason field if non-nil, zero value otherwise.

### GetBannedReasonOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetBannedReasonOk() (*string, bool)`

GetBannedReasonOk returns a tuple with the BannedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedReason

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetBannedReason(v string)`

SetBannedReason sets BannedReason field to given value.


### SetBannedReasonNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetBannedReasonNil(b bool)`

 SetBannedReasonNil sets the value for BannedReason to be an explicit nil

### UnsetBannedReason
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetBannedReason()`

UnsetBannedReason ensures that no value is present for BannedReason, not even an explicit nil
### GetEMail

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetEMail() string`

GetEMail returns the EMail field if non-nil, zero value otherwise.

### GetEMailOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetEMailOk() (*string, bool)`

GetEMailOk returns a tuple with the EMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEMail

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetEMail(v string)`

SetEMail sets EMail field to given value.


### SetEMailNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetEMailNil(b bool)`

 SetEMailNil sets the value for EMail to be an explicit nil

### UnsetEMail
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetEMail()`

UnsetEMail ensures that no value is present for EMail, not even an explicit nil
### GetRealName

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetRealName(v string)`

SetRealName sets RealName field to given value.


### SetRealNameNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetRealNameNil(b bool)`

 SetRealNameNil sets the value for RealName to be an explicit nil

### UnsetRealName
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetRealName()`

UnsetRealName ensures that no value is present for RealName, not even an explicit nil
### GetAboutMe

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetAboutMe() string`

GetAboutMe returns the AboutMe field if non-nil, zero value otherwise.

### GetAboutMeOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetAboutMeOk() (*string, bool)`

GetAboutMeOk returns a tuple with the AboutMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboutMe

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetAboutMe(v string)`

SetAboutMe sets AboutMe field to given value.


### SetAboutMeNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetAboutMeNil(b bool)`

 SetAboutMeNil sets the value for AboutMe to be an explicit nil

### UnsetAboutMe
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetAboutMe()`

UnsetAboutMe ensures that no value is present for AboutMe, not even an explicit nil
### GetBirthday

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.


### GetSteamImgURL

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetSteamImgURL() string`

GetSteamImgURL returns the SteamImgURL field if non-nil, zero value otherwise.

### GetSteamImgURLOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetSteamImgURLOk() (*string, bool)`

GetSteamImgURLOk returns a tuple with the SteamImgURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamImgURL

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetSteamImgURL(v string)`

SetSteamImgURL sets SteamImgURL field to given value.


### SetSteamImgURLNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetSteamImgURLNil(b bool)`

 SetSteamImgURLNil sets the value for SteamImgURL to be an explicit nil

### UnsetSteamImgURL
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetSteamImgURL()`

UnsetSteamImgURL ensures that no value is present for SteamImgURL, not even an explicit nil
### GetResidence

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetResidence() string`

GetResidence returns the Residence field if non-nil, zero value otherwise.

### GetResidenceOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetResidenceOk() (*string, bool)`

GetResidenceOk returns a tuple with the Residence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidence

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetResidence(v string)`

SetResidence sets Residence field to given value.


### SetResidenceNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetResidenceNil(b bool)`

 SetResidenceNil sets the value for Residence to be an explicit nil

### UnsetResidence
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetResidence()`

UnsetResidence ensures that no value is present for Residence, not even an explicit nil
### GetUserName

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.


### SetUserNameNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetKontorActivated

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetKontorActivated() bool`

GetKontorActivated returns the KontorActivated field if non-nil, zero value otherwise.

### GetKontorActivatedOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetKontorActivatedOk() (*bool, bool)`

GetKontorActivatedOk returns a tuple with the KontorActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKontorActivated

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetKontorActivated(v bool)`

SetKontorActivated sets KontorActivated field to given value.


### GetChangelogVisible

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetChangelogVisible() bool`

GetChangelogVisible returns the ChangelogVisible field if non-nil, zero value otherwise.

### GetChangelogVisibleOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetChangelogVisibleOk() (*bool, bool)`

GetChangelogVisibleOk returns a tuple with the ChangelogVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogVisible

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetChangelogVisible(v bool)`

SetChangelogVisible sets ChangelogVisible field to given value.


### GetDeletionFrist

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetDeletionFrist() int32`

GetDeletionFrist returns the DeletionFrist field if non-nil, zero value otherwise.

### GetDeletionFristOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetDeletionFristOk() (*int32, bool)`

GetDeletionFristOk returns a tuple with the DeletionFrist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionFrist

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetDeletionFrist(v int32)`

SetDeletionFrist sets DeletionFrist field to given value.


### GetTfaActivated

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetTfaActivated() bool`

GetTfaActivated returns the TfaActivated field if non-nil, zero value otherwise.

### GetTfaActivatedOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetTfaActivatedOk() (*bool, bool)`

GetTfaActivatedOk returns a tuple with the TfaActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfaActivated

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetTfaActivated(v bool)`

SetTfaActivated sets TfaActivated field to given value.


### GetIsAlphaTester

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetIsAlphaTester() bool`

GetIsAlphaTester returns the IsAlphaTester field if non-nil, zero value otherwise.

### GetIsAlphaTesterOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetIsAlphaTesterOk() (*bool, bool)`

GetIsAlphaTesterOk returns a tuple with the IsAlphaTester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlphaTester

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetIsAlphaTester(v bool)`

SetIsAlphaTester sets IsAlphaTester field to given value.


### GetIsTrustedUser

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetIsTrustedUser() bool`

GetIsTrustedUser returns the IsTrustedUser field if non-nil, zero value otherwise.

### GetIsTrustedUserOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetIsTrustedUserOk() (*bool, bool)`

GetIsTrustedUserOk returns a tuple with the IsTrustedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrustedUser

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetIsTrustedUser(v bool)`

SetIsTrustedUser sets IsTrustedUser field to given value.


### GetSystemAdminPermissions

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetSystemAdminPermissions() FPHSpedVAPIAdminPermissionEnum`

GetSystemAdminPermissions returns the SystemAdminPermissions field if non-nil, zero value otherwise.

### GetSystemAdminPermissionsOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetSystemAdminPermissionsOk() (*FPHSpedVAPIAdminPermissionEnum, bool)`

GetSystemAdminPermissionsOk returns a tuple with the SystemAdminPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAdminPermissions

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetSystemAdminPermissions(v FPHSpedVAPIAdminPermissionEnum)`

SetSystemAdminPermissions sets SystemAdminPermissions field to given value.


### GetSystemAdminRankName

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetSystemAdminRankName() string`

GetSystemAdminRankName returns the SystemAdminRankName field if non-nil, zero value otherwise.

### GetSystemAdminRankNameOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetSystemAdminRankNameOk() (*string, bool)`

GetSystemAdminRankNameOk returns a tuple with the SystemAdminRankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAdminRankName

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetSystemAdminRankName(v string)`

SetSystemAdminRankName sets SystemAdminRankName field to given value.


### SetSystemAdminRankNameNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetSystemAdminRankNameNil(b bool)`

 SetSystemAdminRankNameNil sets the value for SystemAdminRankName to be an explicit nil

### UnsetSystemAdminRankName
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetSystemAdminRankName()`

UnsetSystemAdminRankName ensures that no value is present for SystemAdminRankName, not even an explicit nil
### GetTwitchName

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetTwitchName() string`

GetTwitchName returns the TwitchName field if non-nil, zero value otherwise.

### GetTwitchNameOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetTwitchNameOk() (*string, bool)`

GetTwitchNameOk returns a tuple with the TwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitchName

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetTwitchName(v string)`

SetTwitchName sets TwitchName field to given value.


### SetTwitchNameNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetTwitchNameNil(b bool)`

 SetTwitchNameNil sets the value for TwitchName to be an explicit nil

### UnsetTwitchName
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetTwitchName()`

UnsetTwitchName ensures that no value is present for TwitchName, not even an explicit nil
### GetTwitchID

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetTwitchID() int32`

GetTwitchID returns the TwitchID field if non-nil, zero value otherwise.

### GetTwitchIDOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetTwitchIDOk() (*int32, bool)`

GetTwitchIDOk returns a tuple with the TwitchID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitchID

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetTwitchID(v int32)`

SetTwitchID sets TwitchID field to given value.


### SetTwitchIDNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetTwitchIDNil(b bool)`

 SetTwitchIDNil sets the value for TwitchID to be an explicit nil

### UnsetTwitchID
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetTwitchID()`

UnsetTwitchID ensures that no value is present for TwitchID, not even an explicit nil
### GetDiscordUID

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetDiscordUID() int64`

GetDiscordUID returns the DiscordUID field if non-nil, zero value otherwise.

### GetDiscordUIDOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetDiscordUIDOk() (*int64, bool)`

GetDiscordUIDOk returns a tuple with the DiscordUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordUID

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetDiscordUID(v int64)`

SetDiscordUID sets DiscordUID field to given value.


### SetDiscordUIDNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetDiscordUIDNil(b bool)`

 SetDiscordUIDNil sets the value for DiscordUID to be an explicit nil

### UnsetDiscordUID
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetDiscordUID()`

UnsetDiscordUID ensures that no value is present for DiscordUID, not even an explicit nil
### GetSpedition

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetSpedition returns the Spedition field if non-nil, zero value otherwise.

### GetSpeditionOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetSpeditionOk returns a tuple with the Spedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpedition

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetSpedition sets Spedition field to given value.


### SetSpeditionNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetSpeditionNil(b bool)`

 SetSpeditionNil sets the value for Spedition to be an explicit nil

### UnsetSpedition
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetSpedition()`

UnsetSpedition ensures that no value is present for Spedition, not even an explicit nil
### GetRank

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetRank() FPHSpedVAPIObjectsSpeditionsRank`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetRankOk() (*FPHSpedVAPIObjectsSpeditionsRank, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetRank(v FPHSpedVAPIObjectsSpeditionsRank)`

SetRank sets Rank field to given value.


### SetRankNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetRankNil(b bool)`

 SetRankNil sets the value for Rank to be an explicit nil

### UnsetRank
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetRank()`

UnsetRank ensures that no value is present for Rank, not even an explicit nil
### GetAchievements

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetAchievements() map[string]interface{}`

GetAchievements returns the Achievements field if non-nil, zero value otherwise.

### GetAchievementsOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetAchievementsOk() (*map[string]interface{}, bool)`

GetAchievementsOk returns a tuple with the Achievements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievements

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetAchievements(v map[string]interface{})`

SetAchievements sets Achievements field to given value.


### SetAchievementsNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetAchievementsNil(b bool)`

 SetAchievementsNil sets the value for Achievements to be an explicit nil

### UnsetAchievements
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetAchievements()`

UnsetAchievements ensures that no value is present for Achievements, not even an explicit nil
### GetInSped

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetInSped() bool`

GetInSped returns the InSped field if non-nil, zero value otherwise.

### GetInSpedOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetInSpedOk() (*bool, bool)`

GetInSpedOk returns a tuple with the InSped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInSped

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetInSped(v bool)`

SetInSped sets InSped field to given value.


### GetFormatName

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetFormatName() string`

GetFormatName returns the FormatName field if non-nil, zero value otherwise.

### GetFormatNameOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetFormatNameOk() (*string, bool)`

GetFormatNameOk returns a tuple with the FormatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatName

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetFormatName(v string)`

SetFormatName sets FormatName field to given value.


### SetFormatNameNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetFormatNameNil(b bool)`

 SetFormatNameNil sets the value for FormatName to be an explicit nil

### UnsetFormatName
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetFormatName()`

UnsetFormatName ensures that no value is present for FormatName, not even an explicit nil
### GetLastTaskTime

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetLastTaskTime() time.Time`

GetLastTaskTime returns the LastTaskTime field if non-nil, zero value otherwise.

### GetLastTaskTimeOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetLastTaskTimeOk() (*time.Time, bool)`

GetLastTaskTimeOk returns a tuple with the LastTaskTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTaskTime

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetLastTaskTime(v time.Time)`

SetLastTaskTime sets LastTaskTime field to given value.


### GetIsInCooldown

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetIsInCooldown() bool`

GetIsInCooldown returns the IsInCooldown field if non-nil, zero value otherwise.

### GetIsInCooldownOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetIsInCooldownOk() (*bool, bool)`

GetIsInCooldownOk returns a tuple with the IsInCooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInCooldown

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetIsInCooldown(v bool)`

SetIsInCooldown sets IsInCooldown field to given value.


### GetMainBankAccount

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetMainBankAccount() FPHSpedVAPIObjectsMoneyLiteBankAccount`

GetMainBankAccount returns the MainBankAccount field if non-nil, zero value otherwise.

### GetMainBankAccountOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetMainBankAccountOk() (*FPHSpedVAPIObjectsMoneyLiteBankAccount, bool)`

GetMainBankAccountOk returns a tuple with the MainBankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainBankAccount

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetMainBankAccount(v FPHSpedVAPIObjectsMoneyLiteBankAccount)`

SetMainBankAccount sets MainBankAccount field to given value.


### SetMainBankAccountNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetMainBankAccountNil(b bool)`

 SetMainBankAccountNil sets the value for MainBankAccount to be an explicit nil

### UnsetMainBankAccount
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetMainBankAccount()`

UnsetMainBankAccount ensures that no value is present for MainBankAccount, not even an explicit nil
### GetSpeditionJoinDate

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetSpeditionJoinDate() time.Time`

GetSpeditionJoinDate returns the SpeditionJoinDate field if non-nil, zero value otherwise.

### GetSpeditionJoinDateOk

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) GetSpeditionJoinDateOk() (*time.Time, bool)`

GetSpeditionJoinDateOk returns a tuple with the SpeditionJoinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeditionJoinDate

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetSpeditionJoinDate(v time.Time)`

SetSpeditionJoinDate sets SpeditionJoinDate field to given value.


### SetSpeditionJoinDateNil

`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) SetSpeditionJoinDateNil(b bool)`

 SetSpeditionJoinDateNil sets the value for SpeditionJoinDate to be an explicit nil

### UnsetSpeditionJoinDate
`func (o *FPHSpedVAPIObjectsUsersFullUserInfo) UnsetSpeditionJoinDate()`

UnsetSpeditionJoinDate ensures that no value is present for SpeditionJoinDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


