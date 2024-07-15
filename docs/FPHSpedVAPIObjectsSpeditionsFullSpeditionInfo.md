# FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **NullableString** |  | [readonly] 
**ShortName** | **NullableString** |  | [readonly] 
**ApplImg** | **NullableString** |  | [readonly] 
**ApplText** | **NullableString** |  | [readonly] 
**AdminAnnounce** | **NullableString** |  | [readonly] 
**DiscordWebHook** | **NullableString** |  | [readonly] 
**Expenses** | **int32** |  | [readonly] 
**MaxTruckPerUser** | **int32** |  | [readonly] 
**MaxAllowedVelocity** | **int32** |  | [readonly] 
**MaxAllowedVelocityATS** | **int32** |  | [readonly] 
**DmgReportFromPercent** | **int32** |  | [readonly] 
**MinCoolDown** | **float64** |  | [readonly] 
**CoolDown** | **float64** |  | [readonly] 
**SpeditionType** | [**FPHSpedVAPIEnumsSpeditionType**](FPHSpedVAPIEnumsSpeditionType.md) |   0 &#x3D; NonCompeting  1 &#x3D; LightRealism  2 &#x3D; RealEco  -1 &#x3D; NotSet | [readonly] 
**IsVisible** | **bool** |  | [readonly] 
**AcceptsApplicants** | **bool** |  | [readonly] 
**OneTimeJoinPasswords** | **[]string** |  | [readonly] 
**SpeditionFounder** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**ColorCollections** | [**[]FPHSpedVAPIObjectsSaveGameSaveGameColor**](FPHSpedVAPIObjectsSaveGameSaveGameColor.md) |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsFullSpeditionInfo

`func NewFPHSpedVAPIObjectsSpeditionsFullSpeditionInfo(id int32, name NullableString, shortName NullableString, applImg NullableString, applText NullableString, adminAnnounce NullableString, discordWebHook NullableString, expenses int32, maxTruckPerUser int32, maxAllowedVelocity int32, maxAllowedVelocityATS int32, dmgReportFromPercent int32, minCoolDown float64, coolDown float64, speditionType FPHSpedVAPIEnumsSpeditionType, isVisible bool, acceptsApplicants bool, oneTimeJoinPasswords []string, speditionFounder NullableFPHSpedVAPIObjectsUsersUser, colorCollections []FPHSpedVAPIObjectsSaveGameSaveGameColor, ) *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo`

NewFPHSpedVAPIObjectsSpeditionsFullSpeditionInfo instantiates a new FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsFullSpeditionInfoWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsFullSpeditionInfoWithDefaults() *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo`

NewFPHSpedVAPIObjectsSpeditionsFullSpeditionInfoWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetShortName

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### SetShortNameNil

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetShortNameNil(b bool)`

 SetShortNameNil sets the value for ShortName to be an explicit nil

### UnsetShortName
`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) UnsetShortName()`

UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
### GetApplImg

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetApplImg() string`

GetApplImg returns the ApplImg field if non-nil, zero value otherwise.

### GetApplImgOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetApplImgOk() (*string, bool)`

GetApplImgOk returns a tuple with the ApplImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplImg

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetApplImg(v string)`

SetApplImg sets ApplImg field to given value.


### SetApplImgNil

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetApplImgNil(b bool)`

 SetApplImgNil sets the value for ApplImg to be an explicit nil

### UnsetApplImg
`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) UnsetApplImg()`

UnsetApplImg ensures that no value is present for ApplImg, not even an explicit nil
### GetApplText

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetApplText() string`

GetApplText returns the ApplText field if non-nil, zero value otherwise.

### GetApplTextOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetApplTextOk() (*string, bool)`

GetApplTextOk returns a tuple with the ApplText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplText

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetApplText(v string)`

SetApplText sets ApplText field to given value.


### SetApplTextNil

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetApplTextNil(b bool)`

 SetApplTextNil sets the value for ApplText to be an explicit nil

### UnsetApplText
`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) UnsetApplText()`

UnsetApplText ensures that no value is present for ApplText, not even an explicit nil
### GetAdminAnnounce

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetAdminAnnounce() string`

GetAdminAnnounce returns the AdminAnnounce field if non-nil, zero value otherwise.

### GetAdminAnnounceOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetAdminAnnounceOk() (*string, bool)`

GetAdminAnnounceOk returns a tuple with the AdminAnnounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAnnounce

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetAdminAnnounce(v string)`

SetAdminAnnounce sets AdminAnnounce field to given value.


### SetAdminAnnounceNil

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetAdminAnnounceNil(b bool)`

 SetAdminAnnounceNil sets the value for AdminAnnounce to be an explicit nil

### UnsetAdminAnnounce
`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) UnsetAdminAnnounce()`

UnsetAdminAnnounce ensures that no value is present for AdminAnnounce, not even an explicit nil
### GetDiscordWebHook

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetDiscordWebHook() string`

GetDiscordWebHook returns the DiscordWebHook field if non-nil, zero value otherwise.

### GetDiscordWebHookOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetDiscordWebHookOk() (*string, bool)`

GetDiscordWebHookOk returns a tuple with the DiscordWebHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordWebHook

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetDiscordWebHook(v string)`

SetDiscordWebHook sets DiscordWebHook field to given value.


### SetDiscordWebHookNil

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetDiscordWebHookNil(b bool)`

 SetDiscordWebHookNil sets the value for DiscordWebHook to be an explicit nil

### UnsetDiscordWebHook
`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) UnsetDiscordWebHook()`

UnsetDiscordWebHook ensures that no value is present for DiscordWebHook, not even an explicit nil
### GetExpenses

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetExpenses() int32`

GetExpenses returns the Expenses field if non-nil, zero value otherwise.

### GetExpensesOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetExpensesOk() (*int32, bool)`

GetExpensesOk returns a tuple with the Expenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenses

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetExpenses(v int32)`

SetExpenses sets Expenses field to given value.


### GetMaxTruckPerUser

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetMaxTruckPerUser() int32`

GetMaxTruckPerUser returns the MaxTruckPerUser field if non-nil, zero value otherwise.

### GetMaxTruckPerUserOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetMaxTruckPerUserOk() (*int32, bool)`

GetMaxTruckPerUserOk returns a tuple with the MaxTruckPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTruckPerUser

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetMaxTruckPerUser(v int32)`

SetMaxTruckPerUser sets MaxTruckPerUser field to given value.


### GetMaxAllowedVelocity

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetMaxAllowedVelocity() int32`

GetMaxAllowedVelocity returns the MaxAllowedVelocity field if non-nil, zero value otherwise.

### GetMaxAllowedVelocityOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetMaxAllowedVelocityOk() (*int32, bool)`

GetMaxAllowedVelocityOk returns a tuple with the MaxAllowedVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedVelocity

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetMaxAllowedVelocity(v int32)`

SetMaxAllowedVelocity sets MaxAllowedVelocity field to given value.


### GetMaxAllowedVelocityATS

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetMaxAllowedVelocityATS() int32`

GetMaxAllowedVelocityATS returns the MaxAllowedVelocityATS field if non-nil, zero value otherwise.

### GetMaxAllowedVelocityATSOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetMaxAllowedVelocityATSOk() (*int32, bool)`

GetMaxAllowedVelocityATSOk returns a tuple with the MaxAllowedVelocityATS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedVelocityATS

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetMaxAllowedVelocityATS(v int32)`

SetMaxAllowedVelocityATS sets MaxAllowedVelocityATS field to given value.


### GetDmgReportFromPercent

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetDmgReportFromPercent() int32`

GetDmgReportFromPercent returns the DmgReportFromPercent field if non-nil, zero value otherwise.

### GetDmgReportFromPercentOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetDmgReportFromPercentOk() (*int32, bool)`

GetDmgReportFromPercentOk returns a tuple with the DmgReportFromPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmgReportFromPercent

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetDmgReportFromPercent(v int32)`

SetDmgReportFromPercent sets DmgReportFromPercent field to given value.


### GetMinCoolDown

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetMinCoolDown() float64`

GetMinCoolDown returns the MinCoolDown field if non-nil, zero value otherwise.

### GetMinCoolDownOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetMinCoolDownOk() (*float64, bool)`

GetMinCoolDownOk returns a tuple with the MinCoolDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCoolDown

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetMinCoolDown(v float64)`

SetMinCoolDown sets MinCoolDown field to given value.


### GetCoolDown

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetCoolDown() float64`

GetCoolDown returns the CoolDown field if non-nil, zero value otherwise.

### GetCoolDownOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetCoolDownOk() (*float64, bool)`

GetCoolDownOk returns a tuple with the CoolDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolDown

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetCoolDown(v float64)`

SetCoolDown sets CoolDown field to given value.


### GetSpeditionType

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetSpeditionType() FPHSpedVAPIEnumsSpeditionType`

GetSpeditionType returns the SpeditionType field if non-nil, zero value otherwise.

### GetSpeditionTypeOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetSpeditionTypeOk() (*FPHSpedVAPIEnumsSpeditionType, bool)`

GetSpeditionTypeOk returns a tuple with the SpeditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeditionType

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetSpeditionType(v FPHSpedVAPIEnumsSpeditionType)`

SetSpeditionType sets SpeditionType field to given value.


### GetIsVisible

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.


### GetAcceptsApplicants

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetAcceptsApplicants() bool`

GetAcceptsApplicants returns the AcceptsApplicants field if non-nil, zero value otherwise.

### GetAcceptsApplicantsOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetAcceptsApplicantsOk() (*bool, bool)`

GetAcceptsApplicantsOk returns a tuple with the AcceptsApplicants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsApplicants

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetAcceptsApplicants(v bool)`

SetAcceptsApplicants sets AcceptsApplicants field to given value.


### GetOneTimeJoinPasswords

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetOneTimeJoinPasswords() []string`

GetOneTimeJoinPasswords returns the OneTimeJoinPasswords field if non-nil, zero value otherwise.

### GetOneTimeJoinPasswordsOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetOneTimeJoinPasswordsOk() (*[]string, bool)`

GetOneTimeJoinPasswordsOk returns a tuple with the OneTimeJoinPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimeJoinPasswords

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetOneTimeJoinPasswords(v []string)`

SetOneTimeJoinPasswords sets OneTimeJoinPasswords field to given value.


### SetOneTimeJoinPasswordsNil

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetOneTimeJoinPasswordsNil(b bool)`

 SetOneTimeJoinPasswordsNil sets the value for OneTimeJoinPasswords to be an explicit nil

### UnsetOneTimeJoinPasswords
`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) UnsetOneTimeJoinPasswords()`

UnsetOneTimeJoinPasswords ensures that no value is present for OneTimeJoinPasswords, not even an explicit nil
### GetSpeditionFounder

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetSpeditionFounder() FPHSpedVAPIObjectsUsersUser`

GetSpeditionFounder returns the SpeditionFounder field if non-nil, zero value otherwise.

### GetSpeditionFounderOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetSpeditionFounderOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetSpeditionFounderOk returns a tuple with the SpeditionFounder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeditionFounder

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetSpeditionFounder(v FPHSpedVAPIObjectsUsersUser)`

SetSpeditionFounder sets SpeditionFounder field to given value.


### SetSpeditionFounderNil

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetSpeditionFounderNil(b bool)`

 SetSpeditionFounderNil sets the value for SpeditionFounder to be an explicit nil

### UnsetSpeditionFounder
`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) UnsetSpeditionFounder()`

UnsetSpeditionFounder ensures that no value is present for SpeditionFounder, not even an explicit nil
### GetColorCollections

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetColorCollections() []FPHSpedVAPIObjectsSaveGameSaveGameColor`

GetColorCollections returns the ColorCollections field if non-nil, zero value otherwise.

### GetColorCollectionsOk

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) GetColorCollectionsOk() (*[]FPHSpedVAPIObjectsSaveGameSaveGameColor, bool)`

GetColorCollectionsOk returns a tuple with the ColorCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorCollections

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetColorCollections(v []FPHSpedVAPIObjectsSaveGameSaveGameColor)`

SetColorCollections sets ColorCollections field to given value.


### SetColorCollectionsNil

`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) SetColorCollectionsNil(b bool)`

 SetColorCollectionsNil sets the value for ColorCollections to be an explicit nil

### UnsetColorCollections
`func (o *FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo) UnsetColorCollections()`

UnsetColorCollections ensures that no value is present for ColorCollections, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


