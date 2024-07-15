# FPHSpedVAPIObjectsKontorTrailerMaintenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Trailer** | [**NullableFPHSpedVAPIObjectsKontorTrailer**](FPHSpedVAPIObjectsKontorTrailer.md) |  | [readonly] 
**OwnedBranch** | [**NullableFPHSpedVAPIObjectsSpeditionsOwnedBranch**](FPHSpedVAPIObjectsSpeditionsOwnedBranch.md) |  | [readonly] 
**InitiatedByUser** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**MaintenanceKind** | [**FPHSpedVAPIEnumsMaintenanceKind**](FPHSpedVAPIEnumsMaintenanceKind.md) |   0 &#x3D; Engine  1 &#x3D; OszilationDamper  2 &#x3D; Stabilizer  3 &#x3D; StoneChip  4 &#x3D; Transmission  5 &#x3D; Wishbone  6 &#x3D; BrakePads  7 &#x3D; BrakeDiscs  8 &#x3D; EngineMaintenance  9 &#x3D; TireChange  10 &#x3D; MainCheck  11 &#x3D; SafetyCheck  12 &#x3D; SaddlePlate  13 &#x3D; AirPressureUnit  14 &#x3D; Alternator  15 &#x3D; BrakeVentil  -1 &#x3D; NotSet | [readonly] 
**SparePartWarningSent** | **bool** |  | [readonly] 
**QueueDate** | **time.Time** |  | [readonly] 
**StartDate** | **NullableTime** |  | [readonly] 
**EndDate** | **NullableTime** |  | [readonly] 
**State** | [**FPHSpedVAPIEnumsMaintenanceState**](FPHSpedVAPIEnumsMaintenanceState.md) |   0 &#x3D; Enqueued  11 &#x3D; InDrive  12 &#x3D; NotInOwnedBranch  13 &#x3D; OtherMaintenanceActive  21 &#x3D; MissingSpareParts  22 &#x3D; NotEnoughMaintenancePlaces  91 &#x3D; Processing  92 &#x3D; ProcessingButDelayed  93 &#x3D; Finished  -1 &#x3D; NotSet | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsKontorTrailerMaintenance

`func NewFPHSpedVAPIObjectsKontorTrailerMaintenance(id int32, trailer NullableFPHSpedVAPIObjectsKontorTrailer, ownedBranch NullableFPHSpedVAPIObjectsSpeditionsOwnedBranch, initiatedByUser NullableFPHSpedVAPIObjectsUsersUser, maintenanceKind FPHSpedVAPIEnumsMaintenanceKind, sparePartWarningSent bool, queueDate time.Time, startDate NullableTime, endDate NullableTime, state FPHSpedVAPIEnumsMaintenanceState, ) *FPHSpedVAPIObjectsKontorTrailerMaintenance`

NewFPHSpedVAPIObjectsKontorTrailerMaintenance instantiates a new FPHSpedVAPIObjectsKontorTrailerMaintenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsKontorTrailerMaintenanceWithDefaults

`func NewFPHSpedVAPIObjectsKontorTrailerMaintenanceWithDefaults() *FPHSpedVAPIObjectsKontorTrailerMaintenance`

NewFPHSpedVAPIObjectsKontorTrailerMaintenanceWithDefaults instantiates a new FPHSpedVAPIObjectsKontorTrailerMaintenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetId(v int32)`

SetId sets Id field to given value.


### GetTrailer

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetTrailer() FPHSpedVAPIObjectsKontorTrailer`

GetTrailer returns the Trailer field if non-nil, zero value otherwise.

### GetTrailerOk

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetTrailerOk() (*FPHSpedVAPIObjectsKontorTrailer, bool)`

GetTrailerOk returns a tuple with the Trailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailer

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetTrailer(v FPHSpedVAPIObjectsKontorTrailer)`

SetTrailer sets Trailer field to given value.


### SetTrailerNil

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetTrailerNil(b bool)`

 SetTrailerNil sets the value for Trailer to be an explicit nil

### UnsetTrailer
`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) UnsetTrailer()`

UnsetTrailer ensures that no value is present for Trailer, not even an explicit nil
### GetOwnedBranch

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetOwnedBranch() FPHSpedVAPIObjectsSpeditionsOwnedBranch`

GetOwnedBranch returns the OwnedBranch field if non-nil, zero value otherwise.

### GetOwnedBranchOk

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetOwnedBranchOk() (*FPHSpedVAPIObjectsSpeditionsOwnedBranch, bool)`

GetOwnedBranchOk returns a tuple with the OwnedBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBranch

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetOwnedBranch(v FPHSpedVAPIObjectsSpeditionsOwnedBranch)`

SetOwnedBranch sets OwnedBranch field to given value.


### SetOwnedBranchNil

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetOwnedBranchNil(b bool)`

 SetOwnedBranchNil sets the value for OwnedBranch to be an explicit nil

### UnsetOwnedBranch
`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) UnsetOwnedBranch()`

UnsetOwnedBranch ensures that no value is present for OwnedBranch, not even an explicit nil
### GetInitiatedByUser

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetInitiatedByUser() FPHSpedVAPIObjectsUsersUser`

GetInitiatedByUser returns the InitiatedByUser field if non-nil, zero value otherwise.

### GetInitiatedByUserOk

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetInitiatedByUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetInitiatedByUserOk returns a tuple with the InitiatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedByUser

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetInitiatedByUser(v FPHSpedVAPIObjectsUsersUser)`

SetInitiatedByUser sets InitiatedByUser field to given value.


### SetInitiatedByUserNil

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetInitiatedByUserNil(b bool)`

 SetInitiatedByUserNil sets the value for InitiatedByUser to be an explicit nil

### UnsetInitiatedByUser
`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) UnsetInitiatedByUser()`

UnsetInitiatedByUser ensures that no value is present for InitiatedByUser, not even an explicit nil
### GetMaintenanceKind

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetMaintenanceKind() FPHSpedVAPIEnumsMaintenanceKind`

GetMaintenanceKind returns the MaintenanceKind field if non-nil, zero value otherwise.

### GetMaintenanceKindOk

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetMaintenanceKindOk() (*FPHSpedVAPIEnumsMaintenanceKind, bool)`

GetMaintenanceKindOk returns a tuple with the MaintenanceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceKind

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetMaintenanceKind(v FPHSpedVAPIEnumsMaintenanceKind)`

SetMaintenanceKind sets MaintenanceKind field to given value.


### GetSparePartWarningSent

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetSparePartWarningSent() bool`

GetSparePartWarningSent returns the SparePartWarningSent field if non-nil, zero value otherwise.

### GetSparePartWarningSentOk

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetSparePartWarningSentOk() (*bool, bool)`

GetSparePartWarningSentOk returns a tuple with the SparePartWarningSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparePartWarningSent

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetSparePartWarningSent(v bool)`

SetSparePartWarningSent sets SparePartWarningSent field to given value.


### GetQueueDate

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetQueueDate() time.Time`

GetQueueDate returns the QueueDate field if non-nil, zero value otherwise.

### GetQueueDateOk

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetQueueDateOk() (*time.Time, bool)`

GetQueueDateOk returns a tuple with the QueueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDate

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetQueueDate(v time.Time)`

SetQueueDate sets QueueDate field to given value.


### GetStartDate

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### SetStartDateNil

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetState

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetState() FPHSpedVAPIEnumsMaintenanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) GetStateOk() (*FPHSpedVAPIEnumsMaintenanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FPHSpedVAPIObjectsKontorTrailerMaintenance) SetState(v FPHSpedVAPIEnumsMaintenanceState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


