# FPHSpedVAPIObjectsLiveConvoyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**SpedVTeamRoleName** | **NullableString** |  | [readonly] 
**TmpStaffRoleName** | **NullableString** |  | [readonly] 
**ShowAsOffline** | **bool** |  | 
**State** | [**FPHSpedVAPIEnumsConvoyInfoDriveState**](FPHSpedVAPIEnumsConvoyInfoDriveState.md) |   0 &#x3D; NoGame  1 &#x3D; Paused  2 &#x3D; InDrive  3 &#x3D; OMSI  4 &#x3D; StellwerkSim  5 &#x3D; Bot  6 &#x3D; SimRail_Train  7 &#x3D; SimRail_Dispatch | 
**SteamID** | **int64** |  | [readonly] 
**SteamPic** | **NullableString** |  | [readonly] 
**TwitchURL** | **NullableString** |  | 
**DiscordUID** | **int64** |  | 
**BotMessage** | **NullableString** |  | 
**BotTask** | **NullableString** |  | 
**BotProgressMax** | **int32** |  | 
**BotProgressCur** | **int32** |  | 
**ScSGame** | [**FPHSpedVAPIEnumsGameEnum**](FPHSpedVAPIEnumsGameEnum.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | 
**ScSTotalDistance** | **int32** |  | 
**ScSCurrentDistance** | **int32** |  | 
**ScSOdometer** | **float64** |  | 
**ScSMaxFuel** | **int32** |  | 
**ScSActFuel** | **int32** |  | 
**ScSVelocity** | **float64** |  | 
**ScSSpeedLimit** | **int32** |  | 
**ScSCruiseControl** | **int32** |  | 
**ScSMPServer** | **NullableString** |  | 
**ScSEstArrival** | **time.Time** |  | 
**ScSDamageTruck** | **float64** |  | 
**ScSDamageTrailer** | **float64** |  | 
**ScSDamageFreight** | **float64** |  | 
**ScSTruckManufactor** | **NullableString** |  | 
**ScSTruckModel** | **NullableString** |  | 
**ScSTruckPlate** | **NullableString** |  | 
**ScSXCoord** | **float64** |  | 
**ScSZCoord** | **float64** |  | 
**ScSNearBy** | **NullableString** |  | 
**ScSSCCFollowUser** | **NullableString** |  | 
**TaskWeight** | **NullableString** |  | 
**TaskFreight** | **NullableString** |  | 
**TaskStart** | **NullableString** |  | 
**TaskDest** | **NullableString** |  | 
**TaskStartCompany** | **NullableString** |  | 
**TaskDestCompany** | **NullableString** |  | 
**TaskMap** | **NullableString** |  | 
**OmsILine** | **NullableString** |  | 
**OmsICirc** | **NullableString** |  | 
**OmsIDest** | **NullableString** |  | 
**OmsINext** | **NullableString** |  | 
**OmsIDelay** | **NullableString** |  | 
**OmsIMap** | **NullableString** |  | 
**StWAnlage** | **NullableString** |  | 
**StWCurTrains** | **int32** |  | 
**StWTotalTrains** | **int32** |  | 
**StWCurDelay** | **int32** |  | 
**SpotifyTitle** | **NullableString** |  | 
**SpotifyArtist** | **NullableString** |  | 
**SpotifyAlbum** | **NullableString** |  | 
**SpotifyCoverURL** | **NullableString** |  | 
**SpotifyPlayPosition** | **int32** |  | 
**SpotifyTitleURI** | **NullableString** |  | 
**SimRailServer** | **NullableString** |  | 
**SimRailLatitude** | **float64** |  | 
**SimRailLongitude** | **float64** |  | 
**SimRailTrainNumber** | **NullableString** |  | 
**SimRailTrainVelocity** | **float64** |  | 
**SimRailTrainNextStop** | **NullableString** |  | 
**SimRailTrainDestination** | **NullableString** |  | 
**SimRailDispatchStation** | **NullableString** |  | 

## Methods

### NewFPHSpedVAPIObjectsLiveConvoyInfo

`func NewFPHSpedVAPIObjectsLiveConvoyInfo(user NullableFPHSpedVAPIObjectsUsersUser, spedVTeamRoleName NullableString, tmpStaffRoleName NullableString, showAsOffline bool, state FPHSpedVAPIEnumsConvoyInfoDriveState, steamID int64, steamPic NullableString, twitchURL NullableString, discordUID int64, botMessage NullableString, botTask NullableString, botProgressMax int32, botProgressCur int32, scSGame FPHSpedVAPIEnumsGameEnum, scSTotalDistance int32, scSCurrentDistance int32, scSOdometer float64, scSMaxFuel int32, scSActFuel int32, scSVelocity float64, scSSpeedLimit int32, scSCruiseControl int32, scSMPServer NullableString, scSEstArrival time.Time, scSDamageTruck float64, scSDamageTrailer float64, scSDamageFreight float64, scSTruckManufactor NullableString, scSTruckModel NullableString, scSTruckPlate NullableString, scSXCoord float64, scSZCoord float64, scSNearBy NullableString, scSSCCFollowUser NullableString, taskWeight NullableString, taskFreight NullableString, taskStart NullableString, taskDest NullableString, taskStartCompany NullableString, taskDestCompany NullableString, taskMap NullableString, omsILine NullableString, omsICirc NullableString, omsIDest NullableString, omsINext NullableString, omsIDelay NullableString, omsIMap NullableString, stWAnlage NullableString, stWCurTrains int32, stWTotalTrains int32, stWCurDelay int32, spotifyTitle NullableString, spotifyArtist NullableString, spotifyAlbum NullableString, spotifyCoverURL NullableString, spotifyPlayPosition int32, spotifyTitleURI NullableString, simRailServer NullableString, simRailLatitude float64, simRailLongitude float64, simRailTrainNumber NullableString, simRailTrainVelocity float64, simRailTrainNextStop NullableString, simRailTrainDestination NullableString, simRailDispatchStation NullableString, ) *FPHSpedVAPIObjectsLiveConvoyInfo`

NewFPHSpedVAPIObjectsLiveConvoyInfo instantiates a new FPHSpedVAPIObjectsLiveConvoyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsLiveConvoyInfoWithDefaults

`func NewFPHSpedVAPIObjectsLiveConvoyInfoWithDefaults() *FPHSpedVAPIObjectsLiveConvoyInfo`

NewFPHSpedVAPIObjectsLiveConvoyInfoWithDefaults instantiates a new FPHSpedVAPIObjectsLiveConvoyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetUser() FPHSpedVAPIObjectsUsersUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetUser(v FPHSpedVAPIObjectsUsersUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetSpedVTeamRoleName

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpedVTeamRoleName() string`

GetSpedVTeamRoleName returns the SpedVTeamRoleName field if non-nil, zero value otherwise.

### GetSpedVTeamRoleNameOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpedVTeamRoleNameOk() (*string, bool)`

GetSpedVTeamRoleNameOk returns a tuple with the SpedVTeamRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpedVTeamRoleName

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpedVTeamRoleName(v string)`

SetSpedVTeamRoleName sets SpedVTeamRoleName field to given value.


### SetSpedVTeamRoleNameNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpedVTeamRoleNameNil(b bool)`

 SetSpedVTeamRoleNameNil sets the value for SpedVTeamRoleName to be an explicit nil

### UnsetSpedVTeamRoleName
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetSpedVTeamRoleName()`

UnsetSpedVTeamRoleName ensures that no value is present for SpedVTeamRoleName, not even an explicit nil
### GetTmpStaffRoleName

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTmpStaffRoleName() string`

GetTmpStaffRoleName returns the TmpStaffRoleName field if non-nil, zero value otherwise.

### GetTmpStaffRoleNameOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTmpStaffRoleNameOk() (*string, bool)`

GetTmpStaffRoleNameOk returns a tuple with the TmpStaffRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpStaffRoleName

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTmpStaffRoleName(v string)`

SetTmpStaffRoleName sets TmpStaffRoleName field to given value.


### SetTmpStaffRoleNameNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTmpStaffRoleNameNil(b bool)`

 SetTmpStaffRoleNameNil sets the value for TmpStaffRoleName to be an explicit nil

### UnsetTmpStaffRoleName
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetTmpStaffRoleName()`

UnsetTmpStaffRoleName ensures that no value is present for TmpStaffRoleName, not even an explicit nil
### GetShowAsOffline

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetShowAsOffline() bool`

GetShowAsOffline returns the ShowAsOffline field if non-nil, zero value otherwise.

### GetShowAsOfflineOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetShowAsOfflineOk() (*bool, bool)`

GetShowAsOfflineOk returns a tuple with the ShowAsOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAsOffline

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetShowAsOffline(v bool)`

SetShowAsOffline sets ShowAsOffline field to given value.


### GetState

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetState() FPHSpedVAPIEnumsConvoyInfoDriveState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetStateOk() (*FPHSpedVAPIEnumsConvoyInfoDriveState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetState(v FPHSpedVAPIEnumsConvoyInfoDriveState)`

SetState sets State field to given value.


### GetSteamID

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSteamID() int64`

GetSteamID returns the SteamID field if non-nil, zero value otherwise.

### GetSteamIDOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSteamIDOk() (*int64, bool)`

GetSteamIDOk returns a tuple with the SteamID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamID

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSteamID(v int64)`

SetSteamID sets SteamID field to given value.


### GetSteamPic

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSteamPic() string`

GetSteamPic returns the SteamPic field if non-nil, zero value otherwise.

### GetSteamPicOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSteamPicOk() (*string, bool)`

GetSteamPicOk returns a tuple with the SteamPic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamPic

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSteamPic(v string)`

SetSteamPic sets SteamPic field to given value.


### SetSteamPicNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSteamPicNil(b bool)`

 SetSteamPicNil sets the value for SteamPic to be an explicit nil

### UnsetSteamPic
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetSteamPic()`

UnsetSteamPic ensures that no value is present for SteamPic, not even an explicit nil
### GetTwitchURL

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTwitchURL() string`

GetTwitchURL returns the TwitchURL field if non-nil, zero value otherwise.

### GetTwitchURLOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTwitchURLOk() (*string, bool)`

GetTwitchURLOk returns a tuple with the TwitchURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitchURL

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTwitchURL(v string)`

SetTwitchURL sets TwitchURL field to given value.


### SetTwitchURLNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTwitchURLNil(b bool)`

 SetTwitchURLNil sets the value for TwitchURL to be an explicit nil

### UnsetTwitchURL
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetTwitchURL()`

UnsetTwitchURL ensures that no value is present for TwitchURL, not even an explicit nil
### GetDiscordUID

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetDiscordUID() int64`

GetDiscordUID returns the DiscordUID field if non-nil, zero value otherwise.

### GetDiscordUIDOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetDiscordUIDOk() (*int64, bool)`

GetDiscordUIDOk returns a tuple with the DiscordUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordUID

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetDiscordUID(v int64)`

SetDiscordUID sets DiscordUID field to given value.


### GetBotMessage

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetBotMessage() string`

GetBotMessage returns the BotMessage field if non-nil, zero value otherwise.

### GetBotMessageOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetBotMessageOk() (*string, bool)`

GetBotMessageOk returns a tuple with the BotMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotMessage

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetBotMessage(v string)`

SetBotMessage sets BotMessage field to given value.


### SetBotMessageNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetBotMessageNil(b bool)`

 SetBotMessageNil sets the value for BotMessage to be an explicit nil

### UnsetBotMessage
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetBotMessage()`

UnsetBotMessage ensures that no value is present for BotMessage, not even an explicit nil
### GetBotTask

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetBotTask() string`

GetBotTask returns the BotTask field if non-nil, zero value otherwise.

### GetBotTaskOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetBotTaskOk() (*string, bool)`

GetBotTaskOk returns a tuple with the BotTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotTask

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetBotTask(v string)`

SetBotTask sets BotTask field to given value.


### SetBotTaskNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetBotTaskNil(b bool)`

 SetBotTaskNil sets the value for BotTask to be an explicit nil

### UnsetBotTask
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetBotTask()`

UnsetBotTask ensures that no value is present for BotTask, not even an explicit nil
### GetBotProgressMax

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetBotProgressMax() int32`

GetBotProgressMax returns the BotProgressMax field if non-nil, zero value otherwise.

### GetBotProgressMaxOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetBotProgressMaxOk() (*int32, bool)`

GetBotProgressMaxOk returns a tuple with the BotProgressMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotProgressMax

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetBotProgressMax(v int32)`

SetBotProgressMax sets BotProgressMax field to given value.


### GetBotProgressCur

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetBotProgressCur() int32`

GetBotProgressCur returns the BotProgressCur field if non-nil, zero value otherwise.

### GetBotProgressCurOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetBotProgressCurOk() (*int32, bool)`

GetBotProgressCurOk returns a tuple with the BotProgressCur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotProgressCur

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetBotProgressCur(v int32)`

SetBotProgressCur sets BotProgressCur field to given value.


### GetScSGame

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSGame() FPHSpedVAPIEnumsGameEnum`

GetScSGame returns the ScSGame field if non-nil, zero value otherwise.

### GetScSGameOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSGameOk() (*FPHSpedVAPIEnumsGameEnum, bool)`

GetScSGameOk returns a tuple with the ScSGame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSGame

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSGame(v FPHSpedVAPIEnumsGameEnum)`

SetScSGame sets ScSGame field to given value.


### GetScSTotalDistance

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSTotalDistance() int32`

GetScSTotalDistance returns the ScSTotalDistance field if non-nil, zero value otherwise.

### GetScSTotalDistanceOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSTotalDistanceOk() (*int32, bool)`

GetScSTotalDistanceOk returns a tuple with the ScSTotalDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSTotalDistance

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSTotalDistance(v int32)`

SetScSTotalDistance sets ScSTotalDistance field to given value.


### GetScSCurrentDistance

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSCurrentDistance() int32`

GetScSCurrentDistance returns the ScSCurrentDistance field if non-nil, zero value otherwise.

### GetScSCurrentDistanceOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSCurrentDistanceOk() (*int32, bool)`

GetScSCurrentDistanceOk returns a tuple with the ScSCurrentDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSCurrentDistance

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSCurrentDistance(v int32)`

SetScSCurrentDistance sets ScSCurrentDistance field to given value.


### GetScSOdometer

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSOdometer() float64`

GetScSOdometer returns the ScSOdometer field if non-nil, zero value otherwise.

### GetScSOdometerOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSOdometerOk() (*float64, bool)`

GetScSOdometerOk returns a tuple with the ScSOdometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSOdometer

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSOdometer(v float64)`

SetScSOdometer sets ScSOdometer field to given value.


### GetScSMaxFuel

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSMaxFuel() int32`

GetScSMaxFuel returns the ScSMaxFuel field if non-nil, zero value otherwise.

### GetScSMaxFuelOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSMaxFuelOk() (*int32, bool)`

GetScSMaxFuelOk returns a tuple with the ScSMaxFuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSMaxFuel

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSMaxFuel(v int32)`

SetScSMaxFuel sets ScSMaxFuel field to given value.


### GetScSActFuel

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSActFuel() int32`

GetScSActFuel returns the ScSActFuel field if non-nil, zero value otherwise.

### GetScSActFuelOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSActFuelOk() (*int32, bool)`

GetScSActFuelOk returns a tuple with the ScSActFuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSActFuel

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSActFuel(v int32)`

SetScSActFuel sets ScSActFuel field to given value.


### GetScSVelocity

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSVelocity() float64`

GetScSVelocity returns the ScSVelocity field if non-nil, zero value otherwise.

### GetScSVelocityOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSVelocityOk() (*float64, bool)`

GetScSVelocityOk returns a tuple with the ScSVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSVelocity

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSVelocity(v float64)`

SetScSVelocity sets ScSVelocity field to given value.


### GetScSSpeedLimit

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSSpeedLimit() int32`

GetScSSpeedLimit returns the ScSSpeedLimit field if non-nil, zero value otherwise.

### GetScSSpeedLimitOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSSpeedLimitOk() (*int32, bool)`

GetScSSpeedLimitOk returns a tuple with the ScSSpeedLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSSpeedLimit

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSSpeedLimit(v int32)`

SetScSSpeedLimit sets ScSSpeedLimit field to given value.


### GetScSCruiseControl

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSCruiseControl() int32`

GetScSCruiseControl returns the ScSCruiseControl field if non-nil, zero value otherwise.

### GetScSCruiseControlOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSCruiseControlOk() (*int32, bool)`

GetScSCruiseControlOk returns a tuple with the ScSCruiseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSCruiseControl

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSCruiseControl(v int32)`

SetScSCruiseControl sets ScSCruiseControl field to given value.


### GetScSMPServer

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSMPServer() string`

GetScSMPServer returns the ScSMPServer field if non-nil, zero value otherwise.

### GetScSMPServerOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSMPServerOk() (*string, bool)`

GetScSMPServerOk returns a tuple with the ScSMPServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSMPServer

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSMPServer(v string)`

SetScSMPServer sets ScSMPServer field to given value.


### SetScSMPServerNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSMPServerNil(b bool)`

 SetScSMPServerNil sets the value for ScSMPServer to be an explicit nil

### UnsetScSMPServer
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetScSMPServer()`

UnsetScSMPServer ensures that no value is present for ScSMPServer, not even an explicit nil
### GetScSEstArrival

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSEstArrival() time.Time`

GetScSEstArrival returns the ScSEstArrival field if non-nil, zero value otherwise.

### GetScSEstArrivalOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSEstArrivalOk() (*time.Time, bool)`

GetScSEstArrivalOk returns a tuple with the ScSEstArrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSEstArrival

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSEstArrival(v time.Time)`

SetScSEstArrival sets ScSEstArrival field to given value.


### GetScSDamageTruck

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSDamageTruck() float64`

GetScSDamageTruck returns the ScSDamageTruck field if non-nil, zero value otherwise.

### GetScSDamageTruckOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSDamageTruckOk() (*float64, bool)`

GetScSDamageTruckOk returns a tuple with the ScSDamageTruck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSDamageTruck

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSDamageTruck(v float64)`

SetScSDamageTruck sets ScSDamageTruck field to given value.


### GetScSDamageTrailer

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSDamageTrailer() float64`

GetScSDamageTrailer returns the ScSDamageTrailer field if non-nil, zero value otherwise.

### GetScSDamageTrailerOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSDamageTrailerOk() (*float64, bool)`

GetScSDamageTrailerOk returns a tuple with the ScSDamageTrailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSDamageTrailer

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSDamageTrailer(v float64)`

SetScSDamageTrailer sets ScSDamageTrailer field to given value.


### GetScSDamageFreight

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSDamageFreight() float64`

GetScSDamageFreight returns the ScSDamageFreight field if non-nil, zero value otherwise.

### GetScSDamageFreightOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSDamageFreightOk() (*float64, bool)`

GetScSDamageFreightOk returns a tuple with the ScSDamageFreight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSDamageFreight

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSDamageFreight(v float64)`

SetScSDamageFreight sets ScSDamageFreight field to given value.


### GetScSTruckManufactor

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSTruckManufactor() string`

GetScSTruckManufactor returns the ScSTruckManufactor field if non-nil, zero value otherwise.

### GetScSTruckManufactorOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSTruckManufactorOk() (*string, bool)`

GetScSTruckManufactorOk returns a tuple with the ScSTruckManufactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSTruckManufactor

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSTruckManufactor(v string)`

SetScSTruckManufactor sets ScSTruckManufactor field to given value.


### SetScSTruckManufactorNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSTruckManufactorNil(b bool)`

 SetScSTruckManufactorNil sets the value for ScSTruckManufactor to be an explicit nil

### UnsetScSTruckManufactor
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetScSTruckManufactor()`

UnsetScSTruckManufactor ensures that no value is present for ScSTruckManufactor, not even an explicit nil
### GetScSTruckModel

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSTruckModel() string`

GetScSTruckModel returns the ScSTruckModel field if non-nil, zero value otherwise.

### GetScSTruckModelOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSTruckModelOk() (*string, bool)`

GetScSTruckModelOk returns a tuple with the ScSTruckModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSTruckModel

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSTruckModel(v string)`

SetScSTruckModel sets ScSTruckModel field to given value.


### SetScSTruckModelNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSTruckModelNil(b bool)`

 SetScSTruckModelNil sets the value for ScSTruckModel to be an explicit nil

### UnsetScSTruckModel
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetScSTruckModel()`

UnsetScSTruckModel ensures that no value is present for ScSTruckModel, not even an explicit nil
### GetScSTruckPlate

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSTruckPlate() string`

GetScSTruckPlate returns the ScSTruckPlate field if non-nil, zero value otherwise.

### GetScSTruckPlateOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSTruckPlateOk() (*string, bool)`

GetScSTruckPlateOk returns a tuple with the ScSTruckPlate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSTruckPlate

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSTruckPlate(v string)`

SetScSTruckPlate sets ScSTruckPlate field to given value.


### SetScSTruckPlateNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSTruckPlateNil(b bool)`

 SetScSTruckPlateNil sets the value for ScSTruckPlate to be an explicit nil

### UnsetScSTruckPlate
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetScSTruckPlate()`

UnsetScSTruckPlate ensures that no value is present for ScSTruckPlate, not even an explicit nil
### GetScSXCoord

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSXCoord() float64`

GetScSXCoord returns the ScSXCoord field if non-nil, zero value otherwise.

### GetScSXCoordOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSXCoordOk() (*float64, bool)`

GetScSXCoordOk returns a tuple with the ScSXCoord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSXCoord

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSXCoord(v float64)`

SetScSXCoord sets ScSXCoord field to given value.


### GetScSZCoord

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSZCoord() float64`

GetScSZCoord returns the ScSZCoord field if non-nil, zero value otherwise.

### GetScSZCoordOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSZCoordOk() (*float64, bool)`

GetScSZCoordOk returns a tuple with the ScSZCoord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSZCoord

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSZCoord(v float64)`

SetScSZCoord sets ScSZCoord field to given value.


### GetScSNearBy

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSNearBy() string`

GetScSNearBy returns the ScSNearBy field if non-nil, zero value otherwise.

### GetScSNearByOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSNearByOk() (*string, bool)`

GetScSNearByOk returns a tuple with the ScSNearBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSNearBy

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSNearBy(v string)`

SetScSNearBy sets ScSNearBy field to given value.


### SetScSNearByNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSNearByNil(b bool)`

 SetScSNearByNil sets the value for ScSNearBy to be an explicit nil

### UnsetScSNearBy
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetScSNearBy()`

UnsetScSNearBy ensures that no value is present for ScSNearBy, not even an explicit nil
### GetScSSCCFollowUser

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSSCCFollowUser() string`

GetScSSCCFollowUser returns the ScSSCCFollowUser field if non-nil, zero value otherwise.

### GetScSSCCFollowUserOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetScSSCCFollowUserOk() (*string, bool)`

GetScSSCCFollowUserOk returns a tuple with the ScSSCCFollowUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScSSCCFollowUser

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSSCCFollowUser(v string)`

SetScSSCCFollowUser sets ScSSCCFollowUser field to given value.


### SetScSSCCFollowUserNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetScSSCCFollowUserNil(b bool)`

 SetScSSCCFollowUserNil sets the value for ScSSCCFollowUser to be an explicit nil

### UnsetScSSCCFollowUser
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetScSSCCFollowUser()`

UnsetScSSCCFollowUser ensures that no value is present for ScSSCCFollowUser, not even an explicit nil
### GetTaskWeight

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskWeight() string`

GetTaskWeight returns the TaskWeight field if non-nil, zero value otherwise.

### GetTaskWeightOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskWeightOk() (*string, bool)`

GetTaskWeightOk returns a tuple with the TaskWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWeight

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskWeight(v string)`

SetTaskWeight sets TaskWeight field to given value.


### SetTaskWeightNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskWeightNil(b bool)`

 SetTaskWeightNil sets the value for TaskWeight to be an explicit nil

### UnsetTaskWeight
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetTaskWeight()`

UnsetTaskWeight ensures that no value is present for TaskWeight, not even an explicit nil
### GetTaskFreight

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskFreight() string`

GetTaskFreight returns the TaskFreight field if non-nil, zero value otherwise.

### GetTaskFreightOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskFreightOk() (*string, bool)`

GetTaskFreightOk returns a tuple with the TaskFreight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskFreight

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskFreight(v string)`

SetTaskFreight sets TaskFreight field to given value.


### SetTaskFreightNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskFreightNil(b bool)`

 SetTaskFreightNil sets the value for TaskFreight to be an explicit nil

### UnsetTaskFreight
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetTaskFreight()`

UnsetTaskFreight ensures that no value is present for TaskFreight, not even an explicit nil
### GetTaskStart

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskStart() string`

GetTaskStart returns the TaskStart field if non-nil, zero value otherwise.

### GetTaskStartOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskStartOk() (*string, bool)`

GetTaskStartOk returns a tuple with the TaskStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStart

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskStart(v string)`

SetTaskStart sets TaskStart field to given value.


### SetTaskStartNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskStartNil(b bool)`

 SetTaskStartNil sets the value for TaskStart to be an explicit nil

### UnsetTaskStart
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetTaskStart()`

UnsetTaskStart ensures that no value is present for TaskStart, not even an explicit nil
### GetTaskDest

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskDest() string`

GetTaskDest returns the TaskDest field if non-nil, zero value otherwise.

### GetTaskDestOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskDestOk() (*string, bool)`

GetTaskDestOk returns a tuple with the TaskDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDest

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskDest(v string)`

SetTaskDest sets TaskDest field to given value.


### SetTaskDestNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskDestNil(b bool)`

 SetTaskDestNil sets the value for TaskDest to be an explicit nil

### UnsetTaskDest
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetTaskDest()`

UnsetTaskDest ensures that no value is present for TaskDest, not even an explicit nil
### GetTaskStartCompany

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskStartCompany() string`

GetTaskStartCompany returns the TaskStartCompany field if non-nil, zero value otherwise.

### GetTaskStartCompanyOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskStartCompanyOk() (*string, bool)`

GetTaskStartCompanyOk returns a tuple with the TaskStartCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStartCompany

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskStartCompany(v string)`

SetTaskStartCompany sets TaskStartCompany field to given value.


### SetTaskStartCompanyNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskStartCompanyNil(b bool)`

 SetTaskStartCompanyNil sets the value for TaskStartCompany to be an explicit nil

### UnsetTaskStartCompany
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetTaskStartCompany()`

UnsetTaskStartCompany ensures that no value is present for TaskStartCompany, not even an explicit nil
### GetTaskDestCompany

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskDestCompany() string`

GetTaskDestCompany returns the TaskDestCompany field if non-nil, zero value otherwise.

### GetTaskDestCompanyOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskDestCompanyOk() (*string, bool)`

GetTaskDestCompanyOk returns a tuple with the TaskDestCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDestCompany

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskDestCompany(v string)`

SetTaskDestCompany sets TaskDestCompany field to given value.


### SetTaskDestCompanyNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskDestCompanyNil(b bool)`

 SetTaskDestCompanyNil sets the value for TaskDestCompany to be an explicit nil

### UnsetTaskDestCompany
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetTaskDestCompany()`

UnsetTaskDestCompany ensures that no value is present for TaskDestCompany, not even an explicit nil
### GetTaskMap

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskMap() string`

GetTaskMap returns the TaskMap field if non-nil, zero value otherwise.

### GetTaskMapOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetTaskMapOk() (*string, bool)`

GetTaskMapOk returns a tuple with the TaskMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskMap

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskMap(v string)`

SetTaskMap sets TaskMap field to given value.


### SetTaskMapNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetTaskMapNil(b bool)`

 SetTaskMapNil sets the value for TaskMap to be an explicit nil

### UnsetTaskMap
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetTaskMap()`

UnsetTaskMap ensures that no value is present for TaskMap, not even an explicit nil
### GetOmsILine

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetOmsILine() string`

GetOmsILine returns the OmsILine field if non-nil, zero value otherwise.

### GetOmsILineOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetOmsILineOk() (*string, bool)`

GetOmsILineOk returns a tuple with the OmsILine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmsILine

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetOmsILine(v string)`

SetOmsILine sets OmsILine field to given value.


### SetOmsILineNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetOmsILineNil(b bool)`

 SetOmsILineNil sets the value for OmsILine to be an explicit nil

### UnsetOmsILine
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetOmsILine()`

UnsetOmsILine ensures that no value is present for OmsILine, not even an explicit nil
### GetOmsICirc

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetOmsICirc() string`

GetOmsICirc returns the OmsICirc field if non-nil, zero value otherwise.

### GetOmsICircOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetOmsICircOk() (*string, bool)`

GetOmsICircOk returns a tuple with the OmsICirc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmsICirc

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetOmsICirc(v string)`

SetOmsICirc sets OmsICirc field to given value.


### SetOmsICircNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetOmsICircNil(b bool)`

 SetOmsICircNil sets the value for OmsICirc to be an explicit nil

### UnsetOmsICirc
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetOmsICirc()`

UnsetOmsICirc ensures that no value is present for OmsICirc, not even an explicit nil
### GetOmsIDest

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetOmsIDest() string`

GetOmsIDest returns the OmsIDest field if non-nil, zero value otherwise.

### GetOmsIDestOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetOmsIDestOk() (*string, bool)`

GetOmsIDestOk returns a tuple with the OmsIDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmsIDest

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetOmsIDest(v string)`

SetOmsIDest sets OmsIDest field to given value.


### SetOmsIDestNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetOmsIDestNil(b bool)`

 SetOmsIDestNil sets the value for OmsIDest to be an explicit nil

### UnsetOmsIDest
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetOmsIDest()`

UnsetOmsIDest ensures that no value is present for OmsIDest, not even an explicit nil
### GetOmsINext

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetOmsINext() string`

GetOmsINext returns the OmsINext field if non-nil, zero value otherwise.

### GetOmsINextOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetOmsINextOk() (*string, bool)`

GetOmsINextOk returns a tuple with the OmsINext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmsINext

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetOmsINext(v string)`

SetOmsINext sets OmsINext field to given value.


### SetOmsINextNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetOmsINextNil(b bool)`

 SetOmsINextNil sets the value for OmsINext to be an explicit nil

### UnsetOmsINext
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetOmsINext()`

UnsetOmsINext ensures that no value is present for OmsINext, not even an explicit nil
### GetOmsIDelay

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetOmsIDelay() string`

GetOmsIDelay returns the OmsIDelay field if non-nil, zero value otherwise.

### GetOmsIDelayOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetOmsIDelayOk() (*string, bool)`

GetOmsIDelayOk returns a tuple with the OmsIDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmsIDelay

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetOmsIDelay(v string)`

SetOmsIDelay sets OmsIDelay field to given value.


### SetOmsIDelayNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetOmsIDelayNil(b bool)`

 SetOmsIDelayNil sets the value for OmsIDelay to be an explicit nil

### UnsetOmsIDelay
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetOmsIDelay()`

UnsetOmsIDelay ensures that no value is present for OmsIDelay, not even an explicit nil
### GetOmsIMap

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetOmsIMap() string`

GetOmsIMap returns the OmsIMap field if non-nil, zero value otherwise.

### GetOmsIMapOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetOmsIMapOk() (*string, bool)`

GetOmsIMapOk returns a tuple with the OmsIMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmsIMap

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetOmsIMap(v string)`

SetOmsIMap sets OmsIMap field to given value.


### SetOmsIMapNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetOmsIMapNil(b bool)`

 SetOmsIMapNil sets the value for OmsIMap to be an explicit nil

### UnsetOmsIMap
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetOmsIMap()`

UnsetOmsIMap ensures that no value is present for OmsIMap, not even an explicit nil
### GetStWAnlage

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetStWAnlage() string`

GetStWAnlage returns the StWAnlage field if non-nil, zero value otherwise.

### GetStWAnlageOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetStWAnlageOk() (*string, bool)`

GetStWAnlageOk returns a tuple with the StWAnlage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStWAnlage

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetStWAnlage(v string)`

SetStWAnlage sets StWAnlage field to given value.


### SetStWAnlageNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetStWAnlageNil(b bool)`

 SetStWAnlageNil sets the value for StWAnlage to be an explicit nil

### UnsetStWAnlage
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetStWAnlage()`

UnsetStWAnlage ensures that no value is present for StWAnlage, not even an explicit nil
### GetStWCurTrains

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetStWCurTrains() int32`

GetStWCurTrains returns the StWCurTrains field if non-nil, zero value otherwise.

### GetStWCurTrainsOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetStWCurTrainsOk() (*int32, bool)`

GetStWCurTrainsOk returns a tuple with the StWCurTrains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStWCurTrains

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetStWCurTrains(v int32)`

SetStWCurTrains sets StWCurTrains field to given value.


### GetStWTotalTrains

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetStWTotalTrains() int32`

GetStWTotalTrains returns the StWTotalTrains field if non-nil, zero value otherwise.

### GetStWTotalTrainsOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetStWTotalTrainsOk() (*int32, bool)`

GetStWTotalTrainsOk returns a tuple with the StWTotalTrains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStWTotalTrains

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetStWTotalTrains(v int32)`

SetStWTotalTrains sets StWTotalTrains field to given value.


### GetStWCurDelay

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetStWCurDelay() int32`

GetStWCurDelay returns the StWCurDelay field if non-nil, zero value otherwise.

### GetStWCurDelayOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetStWCurDelayOk() (*int32, bool)`

GetStWCurDelayOk returns a tuple with the StWCurDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStWCurDelay

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetStWCurDelay(v int32)`

SetStWCurDelay sets StWCurDelay field to given value.


### GetSpotifyTitle

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpotifyTitle() string`

GetSpotifyTitle returns the SpotifyTitle field if non-nil, zero value otherwise.

### GetSpotifyTitleOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpotifyTitleOk() (*string, bool)`

GetSpotifyTitleOk returns a tuple with the SpotifyTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotifyTitle

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpotifyTitle(v string)`

SetSpotifyTitle sets SpotifyTitle field to given value.


### SetSpotifyTitleNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpotifyTitleNil(b bool)`

 SetSpotifyTitleNil sets the value for SpotifyTitle to be an explicit nil

### UnsetSpotifyTitle
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetSpotifyTitle()`

UnsetSpotifyTitle ensures that no value is present for SpotifyTitle, not even an explicit nil
### GetSpotifyArtist

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpotifyArtist() string`

GetSpotifyArtist returns the SpotifyArtist field if non-nil, zero value otherwise.

### GetSpotifyArtistOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpotifyArtistOk() (*string, bool)`

GetSpotifyArtistOk returns a tuple with the SpotifyArtist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotifyArtist

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpotifyArtist(v string)`

SetSpotifyArtist sets SpotifyArtist field to given value.


### SetSpotifyArtistNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpotifyArtistNil(b bool)`

 SetSpotifyArtistNil sets the value for SpotifyArtist to be an explicit nil

### UnsetSpotifyArtist
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetSpotifyArtist()`

UnsetSpotifyArtist ensures that no value is present for SpotifyArtist, not even an explicit nil
### GetSpotifyAlbum

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpotifyAlbum() string`

GetSpotifyAlbum returns the SpotifyAlbum field if non-nil, zero value otherwise.

### GetSpotifyAlbumOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpotifyAlbumOk() (*string, bool)`

GetSpotifyAlbumOk returns a tuple with the SpotifyAlbum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotifyAlbum

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpotifyAlbum(v string)`

SetSpotifyAlbum sets SpotifyAlbum field to given value.


### SetSpotifyAlbumNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpotifyAlbumNil(b bool)`

 SetSpotifyAlbumNil sets the value for SpotifyAlbum to be an explicit nil

### UnsetSpotifyAlbum
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetSpotifyAlbum()`

UnsetSpotifyAlbum ensures that no value is present for SpotifyAlbum, not even an explicit nil
### GetSpotifyCoverURL

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpotifyCoverURL() string`

GetSpotifyCoverURL returns the SpotifyCoverURL field if non-nil, zero value otherwise.

### GetSpotifyCoverURLOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpotifyCoverURLOk() (*string, bool)`

GetSpotifyCoverURLOk returns a tuple with the SpotifyCoverURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotifyCoverURL

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpotifyCoverURL(v string)`

SetSpotifyCoverURL sets SpotifyCoverURL field to given value.


### SetSpotifyCoverURLNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpotifyCoverURLNil(b bool)`

 SetSpotifyCoverURLNil sets the value for SpotifyCoverURL to be an explicit nil

### UnsetSpotifyCoverURL
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetSpotifyCoverURL()`

UnsetSpotifyCoverURL ensures that no value is present for SpotifyCoverURL, not even an explicit nil
### GetSpotifyPlayPosition

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpotifyPlayPosition() int32`

GetSpotifyPlayPosition returns the SpotifyPlayPosition field if non-nil, zero value otherwise.

### GetSpotifyPlayPositionOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpotifyPlayPositionOk() (*int32, bool)`

GetSpotifyPlayPositionOk returns a tuple with the SpotifyPlayPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotifyPlayPosition

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpotifyPlayPosition(v int32)`

SetSpotifyPlayPosition sets SpotifyPlayPosition field to given value.


### GetSpotifyTitleURI

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpotifyTitleURI() string`

GetSpotifyTitleURI returns the SpotifyTitleURI field if non-nil, zero value otherwise.

### GetSpotifyTitleURIOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSpotifyTitleURIOk() (*string, bool)`

GetSpotifyTitleURIOk returns a tuple with the SpotifyTitleURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotifyTitleURI

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpotifyTitleURI(v string)`

SetSpotifyTitleURI sets SpotifyTitleURI field to given value.


### SetSpotifyTitleURINil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSpotifyTitleURINil(b bool)`

 SetSpotifyTitleURINil sets the value for SpotifyTitleURI to be an explicit nil

### UnsetSpotifyTitleURI
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetSpotifyTitleURI()`

UnsetSpotifyTitleURI ensures that no value is present for SpotifyTitleURI, not even an explicit nil
### GetSimRailServer

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailServer() string`

GetSimRailServer returns the SimRailServer field if non-nil, zero value otherwise.

### GetSimRailServerOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailServerOk() (*string, bool)`

GetSimRailServerOk returns a tuple with the SimRailServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimRailServer

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailServer(v string)`

SetSimRailServer sets SimRailServer field to given value.


### SetSimRailServerNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailServerNil(b bool)`

 SetSimRailServerNil sets the value for SimRailServer to be an explicit nil

### UnsetSimRailServer
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetSimRailServer()`

UnsetSimRailServer ensures that no value is present for SimRailServer, not even an explicit nil
### GetSimRailLatitude

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailLatitude() float64`

GetSimRailLatitude returns the SimRailLatitude field if non-nil, zero value otherwise.

### GetSimRailLatitudeOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailLatitudeOk() (*float64, bool)`

GetSimRailLatitudeOk returns a tuple with the SimRailLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimRailLatitude

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailLatitude(v float64)`

SetSimRailLatitude sets SimRailLatitude field to given value.


### GetSimRailLongitude

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailLongitude() float64`

GetSimRailLongitude returns the SimRailLongitude field if non-nil, zero value otherwise.

### GetSimRailLongitudeOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailLongitudeOk() (*float64, bool)`

GetSimRailLongitudeOk returns a tuple with the SimRailLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimRailLongitude

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailLongitude(v float64)`

SetSimRailLongitude sets SimRailLongitude field to given value.


### GetSimRailTrainNumber

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailTrainNumber() string`

GetSimRailTrainNumber returns the SimRailTrainNumber field if non-nil, zero value otherwise.

### GetSimRailTrainNumberOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailTrainNumberOk() (*string, bool)`

GetSimRailTrainNumberOk returns a tuple with the SimRailTrainNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimRailTrainNumber

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailTrainNumber(v string)`

SetSimRailTrainNumber sets SimRailTrainNumber field to given value.


### SetSimRailTrainNumberNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailTrainNumberNil(b bool)`

 SetSimRailTrainNumberNil sets the value for SimRailTrainNumber to be an explicit nil

### UnsetSimRailTrainNumber
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetSimRailTrainNumber()`

UnsetSimRailTrainNumber ensures that no value is present for SimRailTrainNumber, not even an explicit nil
### GetSimRailTrainVelocity

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailTrainVelocity() float64`

GetSimRailTrainVelocity returns the SimRailTrainVelocity field if non-nil, zero value otherwise.

### GetSimRailTrainVelocityOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailTrainVelocityOk() (*float64, bool)`

GetSimRailTrainVelocityOk returns a tuple with the SimRailTrainVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimRailTrainVelocity

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailTrainVelocity(v float64)`

SetSimRailTrainVelocity sets SimRailTrainVelocity field to given value.


### GetSimRailTrainNextStop

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailTrainNextStop() string`

GetSimRailTrainNextStop returns the SimRailTrainNextStop field if non-nil, zero value otherwise.

### GetSimRailTrainNextStopOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailTrainNextStopOk() (*string, bool)`

GetSimRailTrainNextStopOk returns a tuple with the SimRailTrainNextStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimRailTrainNextStop

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailTrainNextStop(v string)`

SetSimRailTrainNextStop sets SimRailTrainNextStop field to given value.


### SetSimRailTrainNextStopNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailTrainNextStopNil(b bool)`

 SetSimRailTrainNextStopNil sets the value for SimRailTrainNextStop to be an explicit nil

### UnsetSimRailTrainNextStop
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetSimRailTrainNextStop()`

UnsetSimRailTrainNextStop ensures that no value is present for SimRailTrainNextStop, not even an explicit nil
### GetSimRailTrainDestination

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailTrainDestination() string`

GetSimRailTrainDestination returns the SimRailTrainDestination field if non-nil, zero value otherwise.

### GetSimRailTrainDestinationOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailTrainDestinationOk() (*string, bool)`

GetSimRailTrainDestinationOk returns a tuple with the SimRailTrainDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimRailTrainDestination

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailTrainDestination(v string)`

SetSimRailTrainDestination sets SimRailTrainDestination field to given value.


### SetSimRailTrainDestinationNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailTrainDestinationNil(b bool)`

 SetSimRailTrainDestinationNil sets the value for SimRailTrainDestination to be an explicit nil

### UnsetSimRailTrainDestination
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetSimRailTrainDestination()`

UnsetSimRailTrainDestination ensures that no value is present for SimRailTrainDestination, not even an explicit nil
### GetSimRailDispatchStation

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailDispatchStation() string`

GetSimRailDispatchStation returns the SimRailDispatchStation field if non-nil, zero value otherwise.

### GetSimRailDispatchStationOk

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) GetSimRailDispatchStationOk() (*string, bool)`

GetSimRailDispatchStationOk returns a tuple with the SimRailDispatchStation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimRailDispatchStation

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailDispatchStation(v string)`

SetSimRailDispatchStation sets SimRailDispatchStation field to given value.


### SetSimRailDispatchStationNil

`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) SetSimRailDispatchStationNil(b bool)`

 SetSimRailDispatchStationNil sets the value for SimRailDispatchStation to be an explicit nil

### UnsetSimRailDispatchStation
`func (o *FPHSpedVAPIObjectsLiveConvoyInfo) UnsetSimRailDispatchStation()`

UnsetSimRailDispatchStation ensures that no value is present for SimRailDispatchStation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


