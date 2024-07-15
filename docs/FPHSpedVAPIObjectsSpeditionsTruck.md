# FPHSpedVAPIObjectsSpeditionsTruck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Type** | [**NullableFPHSpedVAPIObjectsSpeditionsTruckType**](FPHSpedVAPIObjectsSpeditionsTruckType.md) |  | [readonly] 
**Branch** | [**NullableFPHSpedVAPIObjectsSpeditionsOwnedBranch**](FPHSpedVAPIObjectsSpeditionsOwnedBranch.md) |  | [readonly] 
**LicensePlate** | **NullableString** |  | [readonly] 
**User** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**DrivenKM** | **int32** |  | [readonly] 
**LastCity** | [**NullableFPHSpedVAPIObjectsMapsCity**](FPHSpedVAPIObjectsMapsCity.md) |  | [readonly] 
**LastTransfer** | **time.Time** |  | [readonly] 
**FreeUsageAllowed** | **bool** |  | [readonly] 
**IsParked** | **bool** |  | [readonly] 
**TruckAvailable** | **time.Time** |  | [readonly] 
**NextMainInspection** | **time.Time** |  | [readonly] 
**NextSafetyCheck** | **time.Time** |  | [readonly] 
**NextEngineMaintenance** | **int32** |  | [readonly] 
**NextTireChange** | **int32** |  | [readonly] 
**NextBrakePadsChange** | **int32** |  | [readonly] 
**NextBrakeDiscChange** | **int32** |  | [readonly] 
**HasWishboneDefect** | **bool** |  | [readonly] 
**HasStabilizerDefect** | **bool** |  | [readonly] 
**HasOszilationDamperDefect** | **bool** |  | [readonly] 
**HasEngineMalfunction** | **bool** |  | [readonly] 
**HasTransmissionDamage** | **bool** |  | [readonly] 
**HasSaddlePlateDamage** | **bool** |  | [readonly] 
**HasAPUDamage** | **bool** |  | [readonly] 
**HasAlternatorDamage** | **bool** |  | [readonly] 
**HasStoneChip** | **bool** |  | [readonly] 
**CurrentTaskID** | **int32** |  | [readonly] 
**NeededMaintenanceJobs** | [**[]FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance**](FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance.md) |  | [readonly] 
**ImageURL** | **NullableString** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsTruck

`func NewFPHSpedVAPIObjectsSpeditionsTruck(id int32, type_ NullableFPHSpedVAPIObjectsSpeditionsTruckType, branch NullableFPHSpedVAPIObjectsSpeditionsOwnedBranch, licensePlate NullableString, user NullableFPHSpedVAPIObjectsUsersUser, drivenKM int32, lastCity NullableFPHSpedVAPIObjectsMapsCity, lastTransfer time.Time, freeUsageAllowed bool, isParked bool, truckAvailable time.Time, nextMainInspection time.Time, nextSafetyCheck time.Time, nextEngineMaintenance int32, nextTireChange int32, nextBrakePadsChange int32, nextBrakeDiscChange int32, hasWishboneDefect bool, hasStabilizerDefect bool, hasOszilationDamperDefect bool, hasEngineMalfunction bool, hasTransmissionDamage bool, hasSaddlePlateDamage bool, hasAPUDamage bool, hasAlternatorDamage bool, hasStoneChip bool, currentTaskID int32, neededMaintenanceJobs []FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance, imageURL NullableString, ) *FPHSpedVAPIObjectsSpeditionsTruck`

NewFPHSpedVAPIObjectsSpeditionsTruck instantiates a new FPHSpedVAPIObjectsSpeditionsTruck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsTruckWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsTruckWithDefaults() *FPHSpedVAPIObjectsSpeditionsTruck`

NewFPHSpedVAPIObjectsSpeditionsTruckWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsTruck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetType() FPHSpedVAPIObjectsSpeditionsTruckType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetTypeOk() (*FPHSpedVAPIObjectsSpeditionsTruckType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetType(v FPHSpedVAPIObjectsSpeditionsTruckType)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FPHSpedVAPIObjectsSpeditionsTruck) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetBranch

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetBranch() FPHSpedVAPIObjectsSpeditionsOwnedBranch`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetBranchOk() (*FPHSpedVAPIObjectsSpeditionsOwnedBranch, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetBranch(v FPHSpedVAPIObjectsSpeditionsOwnedBranch)`

SetBranch sets Branch field to given value.


### SetBranchNil

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *FPHSpedVAPIObjectsSpeditionsTruck) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil
### GetLicensePlate

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetLicensePlate() string`

GetLicensePlate returns the LicensePlate field if non-nil, zero value otherwise.

### GetLicensePlateOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetLicensePlateOk() (*string, bool)`

GetLicensePlateOk returns a tuple with the LicensePlate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlate

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetLicensePlate(v string)`

SetLicensePlate sets LicensePlate field to given value.


### SetLicensePlateNil

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetLicensePlateNil(b bool)`

 SetLicensePlateNil sets the value for LicensePlate to be an explicit nil

### UnsetLicensePlate
`func (o *FPHSpedVAPIObjectsSpeditionsTruck) UnsetLicensePlate()`

UnsetLicensePlate ensures that no value is present for LicensePlate, not even an explicit nil
### GetUser

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetUser() FPHSpedVAPIObjectsUsersUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetUser(v FPHSpedVAPIObjectsUsersUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *FPHSpedVAPIObjectsSpeditionsTruck) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetDrivenKM

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetDrivenKM() int32`

GetDrivenKM returns the DrivenKM field if non-nil, zero value otherwise.

### GetDrivenKMOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetDrivenKMOk() (*int32, bool)`

GetDrivenKMOk returns a tuple with the DrivenKM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivenKM

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetDrivenKM(v int32)`

SetDrivenKM sets DrivenKM field to given value.


### GetLastCity

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetLastCity() FPHSpedVAPIObjectsMapsCity`

GetLastCity returns the LastCity field if non-nil, zero value otherwise.

### GetLastCityOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetLastCityOk() (*FPHSpedVAPIObjectsMapsCity, bool)`

GetLastCityOk returns a tuple with the LastCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCity

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetLastCity(v FPHSpedVAPIObjectsMapsCity)`

SetLastCity sets LastCity field to given value.


### SetLastCityNil

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetLastCityNil(b bool)`

 SetLastCityNil sets the value for LastCity to be an explicit nil

### UnsetLastCity
`func (o *FPHSpedVAPIObjectsSpeditionsTruck) UnsetLastCity()`

UnsetLastCity ensures that no value is present for LastCity, not even an explicit nil
### GetLastTransfer

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetLastTransfer() time.Time`

GetLastTransfer returns the LastTransfer field if non-nil, zero value otherwise.

### GetLastTransferOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetLastTransferOk() (*time.Time, bool)`

GetLastTransferOk returns a tuple with the LastTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransfer

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetLastTransfer(v time.Time)`

SetLastTransfer sets LastTransfer field to given value.


### GetFreeUsageAllowed

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetFreeUsageAllowed() bool`

GetFreeUsageAllowed returns the FreeUsageAllowed field if non-nil, zero value otherwise.

### GetFreeUsageAllowedOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetFreeUsageAllowedOk() (*bool, bool)`

GetFreeUsageAllowedOk returns a tuple with the FreeUsageAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeUsageAllowed

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetFreeUsageAllowed(v bool)`

SetFreeUsageAllowed sets FreeUsageAllowed field to given value.


### GetIsParked

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetIsParked() bool`

GetIsParked returns the IsParked field if non-nil, zero value otherwise.

### GetIsParkedOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetIsParkedOk() (*bool, bool)`

GetIsParkedOk returns a tuple with the IsParked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsParked

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetIsParked(v bool)`

SetIsParked sets IsParked field to given value.


### GetTruckAvailable

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetTruckAvailable() time.Time`

GetTruckAvailable returns the TruckAvailable field if non-nil, zero value otherwise.

### GetTruckAvailableOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetTruckAvailableOk() (*time.Time, bool)`

GetTruckAvailableOk returns a tuple with the TruckAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckAvailable

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetTruckAvailable(v time.Time)`

SetTruckAvailable sets TruckAvailable field to given value.


### GetNextMainInspection

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNextMainInspection() time.Time`

GetNextMainInspection returns the NextMainInspection field if non-nil, zero value otherwise.

### GetNextMainInspectionOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNextMainInspectionOk() (*time.Time, bool)`

GetNextMainInspectionOk returns a tuple with the NextMainInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextMainInspection

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetNextMainInspection(v time.Time)`

SetNextMainInspection sets NextMainInspection field to given value.


### GetNextSafetyCheck

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNextSafetyCheck() time.Time`

GetNextSafetyCheck returns the NextSafetyCheck field if non-nil, zero value otherwise.

### GetNextSafetyCheckOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNextSafetyCheckOk() (*time.Time, bool)`

GetNextSafetyCheckOk returns a tuple with the NextSafetyCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSafetyCheck

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetNextSafetyCheck(v time.Time)`

SetNextSafetyCheck sets NextSafetyCheck field to given value.


### GetNextEngineMaintenance

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNextEngineMaintenance() int32`

GetNextEngineMaintenance returns the NextEngineMaintenance field if non-nil, zero value otherwise.

### GetNextEngineMaintenanceOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNextEngineMaintenanceOk() (*int32, bool)`

GetNextEngineMaintenanceOk returns a tuple with the NextEngineMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEngineMaintenance

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetNextEngineMaintenance(v int32)`

SetNextEngineMaintenance sets NextEngineMaintenance field to given value.


### GetNextTireChange

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNextTireChange() int32`

GetNextTireChange returns the NextTireChange field if non-nil, zero value otherwise.

### GetNextTireChangeOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNextTireChangeOk() (*int32, bool)`

GetNextTireChangeOk returns a tuple with the NextTireChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTireChange

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetNextTireChange(v int32)`

SetNextTireChange sets NextTireChange field to given value.


### GetNextBrakePadsChange

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNextBrakePadsChange() int32`

GetNextBrakePadsChange returns the NextBrakePadsChange field if non-nil, zero value otherwise.

### GetNextBrakePadsChangeOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNextBrakePadsChangeOk() (*int32, bool)`

GetNextBrakePadsChangeOk returns a tuple with the NextBrakePadsChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBrakePadsChange

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetNextBrakePadsChange(v int32)`

SetNextBrakePadsChange sets NextBrakePadsChange field to given value.


### GetNextBrakeDiscChange

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNextBrakeDiscChange() int32`

GetNextBrakeDiscChange returns the NextBrakeDiscChange field if non-nil, zero value otherwise.

### GetNextBrakeDiscChangeOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNextBrakeDiscChangeOk() (*int32, bool)`

GetNextBrakeDiscChangeOk returns a tuple with the NextBrakeDiscChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBrakeDiscChange

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetNextBrakeDiscChange(v int32)`

SetNextBrakeDiscChange sets NextBrakeDiscChange field to given value.


### GetHasWishboneDefect

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasWishboneDefect() bool`

GetHasWishboneDefect returns the HasWishboneDefect field if non-nil, zero value otherwise.

### GetHasWishboneDefectOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasWishboneDefectOk() (*bool, bool)`

GetHasWishboneDefectOk returns a tuple with the HasWishboneDefect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWishboneDefect

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetHasWishboneDefect(v bool)`

SetHasWishboneDefect sets HasWishboneDefect field to given value.


### GetHasStabilizerDefect

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasStabilizerDefect() bool`

GetHasStabilizerDefect returns the HasStabilizerDefect field if non-nil, zero value otherwise.

### GetHasStabilizerDefectOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasStabilizerDefectOk() (*bool, bool)`

GetHasStabilizerDefectOk returns a tuple with the HasStabilizerDefect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStabilizerDefect

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetHasStabilizerDefect(v bool)`

SetHasStabilizerDefect sets HasStabilizerDefect field to given value.


### GetHasOszilationDamperDefect

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasOszilationDamperDefect() bool`

GetHasOszilationDamperDefect returns the HasOszilationDamperDefect field if non-nil, zero value otherwise.

### GetHasOszilationDamperDefectOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasOszilationDamperDefectOk() (*bool, bool)`

GetHasOszilationDamperDefectOk returns a tuple with the HasOszilationDamperDefect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOszilationDamperDefect

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetHasOszilationDamperDefect(v bool)`

SetHasOszilationDamperDefect sets HasOszilationDamperDefect field to given value.


### GetHasEngineMalfunction

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasEngineMalfunction() bool`

GetHasEngineMalfunction returns the HasEngineMalfunction field if non-nil, zero value otherwise.

### GetHasEngineMalfunctionOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasEngineMalfunctionOk() (*bool, bool)`

GetHasEngineMalfunctionOk returns a tuple with the HasEngineMalfunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEngineMalfunction

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetHasEngineMalfunction(v bool)`

SetHasEngineMalfunction sets HasEngineMalfunction field to given value.


### GetHasTransmissionDamage

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasTransmissionDamage() bool`

GetHasTransmissionDamage returns the HasTransmissionDamage field if non-nil, zero value otherwise.

### GetHasTransmissionDamageOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasTransmissionDamageOk() (*bool, bool)`

GetHasTransmissionDamageOk returns a tuple with the HasTransmissionDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTransmissionDamage

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetHasTransmissionDamage(v bool)`

SetHasTransmissionDamage sets HasTransmissionDamage field to given value.


### GetHasSaddlePlateDamage

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasSaddlePlateDamage() bool`

GetHasSaddlePlateDamage returns the HasSaddlePlateDamage field if non-nil, zero value otherwise.

### GetHasSaddlePlateDamageOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasSaddlePlateDamageOk() (*bool, bool)`

GetHasSaddlePlateDamageOk returns a tuple with the HasSaddlePlateDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSaddlePlateDamage

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetHasSaddlePlateDamage(v bool)`

SetHasSaddlePlateDamage sets HasSaddlePlateDamage field to given value.


### GetHasAPUDamage

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasAPUDamage() bool`

GetHasAPUDamage returns the HasAPUDamage field if non-nil, zero value otherwise.

### GetHasAPUDamageOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasAPUDamageOk() (*bool, bool)`

GetHasAPUDamageOk returns a tuple with the HasAPUDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAPUDamage

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetHasAPUDamage(v bool)`

SetHasAPUDamage sets HasAPUDamage field to given value.


### GetHasAlternatorDamage

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasAlternatorDamage() bool`

GetHasAlternatorDamage returns the HasAlternatorDamage field if non-nil, zero value otherwise.

### GetHasAlternatorDamageOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasAlternatorDamageOk() (*bool, bool)`

GetHasAlternatorDamageOk returns a tuple with the HasAlternatorDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAlternatorDamage

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetHasAlternatorDamage(v bool)`

SetHasAlternatorDamage sets HasAlternatorDamage field to given value.


### GetHasStoneChip

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasStoneChip() bool`

GetHasStoneChip returns the HasStoneChip field if non-nil, zero value otherwise.

### GetHasStoneChipOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetHasStoneChipOk() (*bool, bool)`

GetHasStoneChipOk returns a tuple with the HasStoneChip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStoneChip

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetHasStoneChip(v bool)`

SetHasStoneChip sets HasStoneChip field to given value.


### GetCurrentTaskID

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetCurrentTaskID() int32`

GetCurrentTaskID returns the CurrentTaskID field if non-nil, zero value otherwise.

### GetCurrentTaskIDOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetCurrentTaskIDOk() (*int32, bool)`

GetCurrentTaskIDOk returns a tuple with the CurrentTaskID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTaskID

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetCurrentTaskID(v int32)`

SetCurrentTaskID sets CurrentTaskID field to given value.


### GetNeededMaintenanceJobs

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNeededMaintenanceJobs() []FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance`

GetNeededMaintenanceJobs returns the NeededMaintenanceJobs field if non-nil, zero value otherwise.

### GetNeededMaintenanceJobsOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetNeededMaintenanceJobsOk() (*[]FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance, bool)`

GetNeededMaintenanceJobsOk returns a tuple with the NeededMaintenanceJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeededMaintenanceJobs

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetNeededMaintenanceJobs(v []FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance)`

SetNeededMaintenanceJobs sets NeededMaintenanceJobs field to given value.


### SetNeededMaintenanceJobsNil

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetNeededMaintenanceJobsNil(b bool)`

 SetNeededMaintenanceJobsNil sets the value for NeededMaintenanceJobs to be an explicit nil

### UnsetNeededMaintenanceJobs
`func (o *FPHSpedVAPIObjectsSpeditionsTruck) UnsetNeededMaintenanceJobs()`

UnsetNeededMaintenanceJobs ensures that no value is present for NeededMaintenanceJobs, not even an explicit nil
### GetImageURL

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetImageURL() string`

GetImageURL returns the ImageURL field if non-nil, zero value otherwise.

### GetImageURLOk

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) GetImageURLOk() (*string, bool)`

GetImageURLOk returns a tuple with the ImageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageURL

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetImageURL(v string)`

SetImageURL sets ImageURL field to given value.


### SetImageURLNil

`func (o *FPHSpedVAPIObjectsSpeditionsTruck) SetImageURLNil(b bool)`

 SetImageURLNil sets the value for ImageURL to be an explicit nil

### UnsetImageURL
`func (o *FPHSpedVAPIObjectsSpeditionsTruck) UnsetImageURL()`

UnsetImageURL ensures that no value is present for ImageURL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


