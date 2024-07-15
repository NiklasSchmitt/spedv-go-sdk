# FPHSpedVAPIObjectsKontorTrailer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Spedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**LastLocation** | [**NullableFPHSpedVAPIObjectsMapsCompanyCity**](FPHSpedVAPIObjectsMapsCompanyCity.md) |  | [readonly] 
**ActLocation** | [**NullableFPHSpedVAPIObjectsMapsCompanyCity**](FPHSpedVAPIObjectsMapsCompanyCity.md) |  | [readonly] 
**Game** | [**FPHSpedVAPIEnumsGameEnum**](FPHSpedVAPIEnumsGameEnum.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | [readonly] 
**LicensePlate** | **NullableString** |  | [readonly] 
**Comment** | **NullableString** |  | [readonly] 
**Category** | **NullableString** |  | [readonly] 
**Skin** | **NullableString** |  | [readonly] 
**TrailerType** | [**FPHSpedVAPIEnumsKontorTrailerType**](FPHSpedVAPIEnumsKontorTrailerType.md) |   0 &#x3D; StandardTrailer  1 &#x3D; DoubleTrailer  2 &#x3D; BDoubleTrailer  3 &#x3D; TripleTrailer  4 &#x3D; ShortTrailer  -1 &#x3D; NotSet | [readonly] 
**DefaultDriver** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**AssignedScheduler** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**TrailerAvailable** | **time.Time** |  | [readonly] 
**NextMainInspection** | **time.Time** |  | [readonly] 
**NextSafetyCheck** | **time.Time** |  | [readonly] 
**NextTireChange** | **int32** |  | [readonly] 
**NextBrakePadsChange** | **int32** |  | [readonly] 
**NextBrakeDiscChange** | **int32** |  | [readonly] 
**HasOszilationDamperDefect** | **bool** |  | [readonly] 
**HasBrakeVentilDefect** | **bool** |  | [readonly] 
**NeededMaintenanceJobs** | [**[]FPHSpedVAPIObjectsKontorNeededTrailerMaintenance**](FPHSpedVAPIObjectsKontorNeededTrailerMaintenance.md) |  | [readonly] 
**Km** | **int32** |  | [readonly] 
**MaxWeight** | **int32** |  | [readonly] 
**IsPlanned** | **bool** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsKontorTrailer

`func NewFPHSpedVAPIObjectsKontorTrailer(id int32, spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, lastLocation NullableFPHSpedVAPIObjectsMapsCompanyCity, actLocation NullableFPHSpedVAPIObjectsMapsCompanyCity, game FPHSpedVAPIEnumsGameEnum, licensePlate NullableString, comment NullableString, category NullableString, skin NullableString, trailerType FPHSpedVAPIEnumsKontorTrailerType, defaultDriver NullableFPHSpedVAPIObjectsUsersUser, assignedScheduler NullableFPHSpedVAPIObjectsUsersUser, trailerAvailable time.Time, nextMainInspection time.Time, nextSafetyCheck time.Time, nextTireChange int32, nextBrakePadsChange int32, nextBrakeDiscChange int32, hasOszilationDamperDefect bool, hasBrakeVentilDefect bool, neededMaintenanceJobs []FPHSpedVAPIObjectsKontorNeededTrailerMaintenance, km int32, maxWeight int32, isPlanned bool, ) *FPHSpedVAPIObjectsKontorTrailer`

NewFPHSpedVAPIObjectsKontorTrailer instantiates a new FPHSpedVAPIObjectsKontorTrailer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsKontorTrailerWithDefaults

`func NewFPHSpedVAPIObjectsKontorTrailerWithDefaults() *FPHSpedVAPIObjectsKontorTrailer`

NewFPHSpedVAPIObjectsKontorTrailerWithDefaults instantiates a new FPHSpedVAPIObjectsKontorTrailer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetId(v int32)`

SetId sets Id field to given value.


### GetSpedition

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetSpedition returns the Spedition field if non-nil, zero value otherwise.

### GetSpeditionOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetSpeditionOk returns a tuple with the Spedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpedition

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetSpedition sets Spedition field to given value.


### SetSpeditionNil

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetSpeditionNil(b bool)`

 SetSpeditionNil sets the value for Spedition to be an explicit nil

### UnsetSpedition
`func (o *FPHSpedVAPIObjectsKontorTrailer) UnsetSpedition()`

UnsetSpedition ensures that no value is present for Spedition, not even an explicit nil
### GetLastLocation

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetLastLocation() FPHSpedVAPIObjectsMapsCompanyCity`

GetLastLocation returns the LastLocation field if non-nil, zero value otherwise.

### GetLastLocationOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetLastLocationOk() (*FPHSpedVAPIObjectsMapsCompanyCity, bool)`

GetLastLocationOk returns a tuple with the LastLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocation

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetLastLocation(v FPHSpedVAPIObjectsMapsCompanyCity)`

SetLastLocation sets LastLocation field to given value.


### SetLastLocationNil

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetLastLocationNil(b bool)`

 SetLastLocationNil sets the value for LastLocation to be an explicit nil

### UnsetLastLocation
`func (o *FPHSpedVAPIObjectsKontorTrailer) UnsetLastLocation()`

UnsetLastLocation ensures that no value is present for LastLocation, not even an explicit nil
### GetActLocation

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetActLocation() FPHSpedVAPIObjectsMapsCompanyCity`

GetActLocation returns the ActLocation field if non-nil, zero value otherwise.

### GetActLocationOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetActLocationOk() (*FPHSpedVAPIObjectsMapsCompanyCity, bool)`

GetActLocationOk returns a tuple with the ActLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActLocation

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetActLocation(v FPHSpedVAPIObjectsMapsCompanyCity)`

SetActLocation sets ActLocation field to given value.


### SetActLocationNil

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetActLocationNil(b bool)`

 SetActLocationNil sets the value for ActLocation to be an explicit nil

### UnsetActLocation
`func (o *FPHSpedVAPIObjectsKontorTrailer) UnsetActLocation()`

UnsetActLocation ensures that no value is present for ActLocation, not even an explicit nil
### GetGame

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetGame() FPHSpedVAPIEnumsGameEnum`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetGameOk() (*FPHSpedVAPIEnumsGameEnum, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetGame(v FPHSpedVAPIEnumsGameEnum)`

SetGame sets Game field to given value.


### GetLicensePlate

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetLicensePlate() string`

GetLicensePlate returns the LicensePlate field if non-nil, zero value otherwise.

### GetLicensePlateOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetLicensePlateOk() (*string, bool)`

GetLicensePlateOk returns a tuple with the LicensePlate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlate

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetLicensePlate(v string)`

SetLicensePlate sets LicensePlate field to given value.


### SetLicensePlateNil

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetLicensePlateNil(b bool)`

 SetLicensePlateNil sets the value for LicensePlate to be an explicit nil

### UnsetLicensePlate
`func (o *FPHSpedVAPIObjectsKontorTrailer) UnsetLicensePlate()`

UnsetLicensePlate ensures that no value is present for LicensePlate, not even an explicit nil
### GetComment

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FPHSpedVAPIObjectsKontorTrailer) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCategory

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetCategory(v string)`

SetCategory sets Category field to given value.


### SetCategoryNil

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *FPHSpedVAPIObjectsKontorTrailer) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetSkin

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetSkin() string`

GetSkin returns the Skin field if non-nil, zero value otherwise.

### GetSkinOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetSkinOk() (*string, bool)`

GetSkinOk returns a tuple with the Skin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkin

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetSkin(v string)`

SetSkin sets Skin field to given value.


### SetSkinNil

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetSkinNil(b bool)`

 SetSkinNil sets the value for Skin to be an explicit nil

### UnsetSkin
`func (o *FPHSpedVAPIObjectsKontorTrailer) UnsetSkin()`

UnsetSkin ensures that no value is present for Skin, not even an explicit nil
### GetTrailerType

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetTrailerType() FPHSpedVAPIEnumsKontorTrailerType`

GetTrailerType returns the TrailerType field if non-nil, zero value otherwise.

### GetTrailerTypeOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetTrailerTypeOk() (*FPHSpedVAPIEnumsKontorTrailerType, bool)`

GetTrailerTypeOk returns a tuple with the TrailerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerType

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetTrailerType(v FPHSpedVAPIEnumsKontorTrailerType)`

SetTrailerType sets TrailerType field to given value.


### GetDefaultDriver

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetDefaultDriver() FPHSpedVAPIObjectsUsersUser`

GetDefaultDriver returns the DefaultDriver field if non-nil, zero value otherwise.

### GetDefaultDriverOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetDefaultDriverOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetDefaultDriverOk returns a tuple with the DefaultDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDriver

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetDefaultDriver(v FPHSpedVAPIObjectsUsersUser)`

SetDefaultDriver sets DefaultDriver field to given value.


### SetDefaultDriverNil

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetDefaultDriverNil(b bool)`

 SetDefaultDriverNil sets the value for DefaultDriver to be an explicit nil

### UnsetDefaultDriver
`func (o *FPHSpedVAPIObjectsKontorTrailer) UnsetDefaultDriver()`

UnsetDefaultDriver ensures that no value is present for DefaultDriver, not even an explicit nil
### GetAssignedScheduler

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetAssignedScheduler() FPHSpedVAPIObjectsUsersUser`

GetAssignedScheduler returns the AssignedScheduler field if non-nil, zero value otherwise.

### GetAssignedSchedulerOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetAssignedSchedulerOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetAssignedSchedulerOk returns a tuple with the AssignedScheduler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedScheduler

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetAssignedScheduler(v FPHSpedVAPIObjectsUsersUser)`

SetAssignedScheduler sets AssignedScheduler field to given value.


### SetAssignedSchedulerNil

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetAssignedSchedulerNil(b bool)`

 SetAssignedSchedulerNil sets the value for AssignedScheduler to be an explicit nil

### UnsetAssignedScheduler
`func (o *FPHSpedVAPIObjectsKontorTrailer) UnsetAssignedScheduler()`

UnsetAssignedScheduler ensures that no value is present for AssignedScheduler, not even an explicit nil
### GetTrailerAvailable

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetTrailerAvailable() time.Time`

GetTrailerAvailable returns the TrailerAvailable field if non-nil, zero value otherwise.

### GetTrailerAvailableOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetTrailerAvailableOk() (*time.Time, bool)`

GetTrailerAvailableOk returns a tuple with the TrailerAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerAvailable

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetTrailerAvailable(v time.Time)`

SetTrailerAvailable sets TrailerAvailable field to given value.


### GetNextMainInspection

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetNextMainInspection() time.Time`

GetNextMainInspection returns the NextMainInspection field if non-nil, zero value otherwise.

### GetNextMainInspectionOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetNextMainInspectionOk() (*time.Time, bool)`

GetNextMainInspectionOk returns a tuple with the NextMainInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextMainInspection

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetNextMainInspection(v time.Time)`

SetNextMainInspection sets NextMainInspection field to given value.


### GetNextSafetyCheck

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetNextSafetyCheck() time.Time`

GetNextSafetyCheck returns the NextSafetyCheck field if non-nil, zero value otherwise.

### GetNextSafetyCheckOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetNextSafetyCheckOk() (*time.Time, bool)`

GetNextSafetyCheckOk returns a tuple with the NextSafetyCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSafetyCheck

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetNextSafetyCheck(v time.Time)`

SetNextSafetyCheck sets NextSafetyCheck field to given value.


### GetNextTireChange

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetNextTireChange() int32`

GetNextTireChange returns the NextTireChange field if non-nil, zero value otherwise.

### GetNextTireChangeOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetNextTireChangeOk() (*int32, bool)`

GetNextTireChangeOk returns a tuple with the NextTireChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTireChange

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetNextTireChange(v int32)`

SetNextTireChange sets NextTireChange field to given value.


### GetNextBrakePadsChange

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetNextBrakePadsChange() int32`

GetNextBrakePadsChange returns the NextBrakePadsChange field if non-nil, zero value otherwise.

### GetNextBrakePadsChangeOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetNextBrakePadsChangeOk() (*int32, bool)`

GetNextBrakePadsChangeOk returns a tuple with the NextBrakePadsChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBrakePadsChange

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetNextBrakePadsChange(v int32)`

SetNextBrakePadsChange sets NextBrakePadsChange field to given value.


### GetNextBrakeDiscChange

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetNextBrakeDiscChange() int32`

GetNextBrakeDiscChange returns the NextBrakeDiscChange field if non-nil, zero value otherwise.

### GetNextBrakeDiscChangeOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetNextBrakeDiscChangeOk() (*int32, bool)`

GetNextBrakeDiscChangeOk returns a tuple with the NextBrakeDiscChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBrakeDiscChange

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetNextBrakeDiscChange(v int32)`

SetNextBrakeDiscChange sets NextBrakeDiscChange field to given value.


### GetHasOszilationDamperDefect

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetHasOszilationDamperDefect() bool`

GetHasOszilationDamperDefect returns the HasOszilationDamperDefect field if non-nil, zero value otherwise.

### GetHasOszilationDamperDefectOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetHasOszilationDamperDefectOk() (*bool, bool)`

GetHasOszilationDamperDefectOk returns a tuple with the HasOszilationDamperDefect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOszilationDamperDefect

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetHasOszilationDamperDefect(v bool)`

SetHasOszilationDamperDefect sets HasOszilationDamperDefect field to given value.


### GetHasBrakeVentilDefect

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetHasBrakeVentilDefect() bool`

GetHasBrakeVentilDefect returns the HasBrakeVentilDefect field if non-nil, zero value otherwise.

### GetHasBrakeVentilDefectOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetHasBrakeVentilDefectOk() (*bool, bool)`

GetHasBrakeVentilDefectOk returns a tuple with the HasBrakeVentilDefect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBrakeVentilDefect

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetHasBrakeVentilDefect(v bool)`

SetHasBrakeVentilDefect sets HasBrakeVentilDefect field to given value.


### GetNeededMaintenanceJobs

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetNeededMaintenanceJobs() []FPHSpedVAPIObjectsKontorNeededTrailerMaintenance`

GetNeededMaintenanceJobs returns the NeededMaintenanceJobs field if non-nil, zero value otherwise.

### GetNeededMaintenanceJobsOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetNeededMaintenanceJobsOk() (*[]FPHSpedVAPIObjectsKontorNeededTrailerMaintenance, bool)`

GetNeededMaintenanceJobsOk returns a tuple with the NeededMaintenanceJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeededMaintenanceJobs

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetNeededMaintenanceJobs(v []FPHSpedVAPIObjectsKontorNeededTrailerMaintenance)`

SetNeededMaintenanceJobs sets NeededMaintenanceJobs field to given value.


### SetNeededMaintenanceJobsNil

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetNeededMaintenanceJobsNil(b bool)`

 SetNeededMaintenanceJobsNil sets the value for NeededMaintenanceJobs to be an explicit nil

### UnsetNeededMaintenanceJobs
`func (o *FPHSpedVAPIObjectsKontorTrailer) UnsetNeededMaintenanceJobs()`

UnsetNeededMaintenanceJobs ensures that no value is present for NeededMaintenanceJobs, not even an explicit nil
### GetKm

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetKm() int32`

GetKm returns the Km field if non-nil, zero value otherwise.

### GetKmOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetKmOk() (*int32, bool)`

GetKmOk returns a tuple with the Km field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKm

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetKm(v int32)`

SetKm sets Km field to given value.


### GetMaxWeight

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetMaxWeight() int32`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetMaxWeightOk() (*int32, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetMaxWeight(v int32)`

SetMaxWeight sets MaxWeight field to given value.


### GetIsPlanned

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetIsPlanned() bool`

GetIsPlanned returns the IsPlanned field if non-nil, zero value otherwise.

### GetIsPlannedOk

`func (o *FPHSpedVAPIObjectsKontorTrailer) GetIsPlannedOk() (*bool, bool)`

GetIsPlannedOk returns a tuple with the IsPlanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlanned

`func (o *FPHSpedVAPIObjectsKontorTrailer) SetIsPlanned(v bool)`

SetIsPlanned sets IsPlanned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


