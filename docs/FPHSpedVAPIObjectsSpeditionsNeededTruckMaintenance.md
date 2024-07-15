# FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**FPHSpedVAPIEnumsMaintenanceKind**](FPHSpedVAPIEnumsMaintenanceKind.md) |   0 &#x3D; Engine  1 &#x3D; OszilationDamper  2 &#x3D; Stabilizer  3 &#x3D; StoneChip  4 &#x3D; Transmission  5 &#x3D; Wishbone  6 &#x3D; BrakePads  7 &#x3D; BrakeDiscs  8 &#x3D; EngineMaintenance  9 &#x3D; TireChange  10 &#x3D; MainCheck  11 &#x3D; SafetyCheck  12 &#x3D; SaddlePlate  13 &#x3D; AirPressureUnit  14 &#x3D; Alternator  15 &#x3D; BrakeVentil  -1 &#x3D; NotSet | [readonly] 
**ExternalTime** | **string** |  | [readonly] 
**InternalTime** | **string** |  | [readonly] 
**ExternalCost** | **int32** |  | [readonly] 
**InternalCost** | **int32** |  | [readonly] 
**NeededImmediately** | **bool** |  | [readonly] 
**NeededTillKM** | **NullableInt32** |  | [readonly] 
**NeededTillDate** | **NullableTime** |  | [readonly] 
**RepeatTimeSpan** | **NullableString** |  | [readonly] 
**RepeatKM** | **NullableInt32** |  | [readonly] 
**CurrentKM** | **int32** |  | [readonly] 
**Game** | [**FPHSpedVAPIEnumsGameEnum**](FPHSpedVAPIEnumsGameEnum.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | [readonly] 
**KmSinceLast** | **NullableInt32** |  | [readonly] 
**TimeSinceLast** | **NullableString** |  | [readonly] 
**VisibleRemaining** | **NullableString** |  | [readonly] 
**RepeatSpan** | **float64** |  | [readonly] 
**SpanSinceLast** | **float64** |  | [readonly] 
**IsNeeded** | **bool** |  | [readonly] 
**IsOverdue** | **bool** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance

`func NewFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance(kind FPHSpedVAPIEnumsMaintenanceKind, externalTime string, internalTime string, externalCost int32, internalCost int32, neededImmediately bool, neededTillKM NullableInt32, neededTillDate NullableTime, repeatTimeSpan NullableString, repeatKM NullableInt32, currentKM int32, game FPHSpedVAPIEnumsGameEnum, kmSinceLast NullableInt32, timeSinceLast NullableString, visibleRemaining NullableString, repeatSpan float64, spanSinceLast float64, isNeeded bool, isOverdue bool, ) *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance`

NewFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance instantiates a new FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenanceWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenanceWithDefaults() *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance`

NewFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenanceWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetKind() FPHSpedVAPIEnumsMaintenanceKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetKindOk() (*FPHSpedVAPIEnumsMaintenanceKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetKind(v FPHSpedVAPIEnumsMaintenanceKind)`

SetKind sets Kind field to given value.


### GetExternalTime

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetExternalTime() string`

GetExternalTime returns the ExternalTime field if non-nil, zero value otherwise.

### GetExternalTimeOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetExternalTimeOk() (*string, bool)`

GetExternalTimeOk returns a tuple with the ExternalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTime

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetExternalTime(v string)`

SetExternalTime sets ExternalTime field to given value.


### GetInternalTime

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetInternalTime() string`

GetInternalTime returns the InternalTime field if non-nil, zero value otherwise.

### GetInternalTimeOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetInternalTimeOk() (*string, bool)`

GetInternalTimeOk returns a tuple with the InternalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTime

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetInternalTime(v string)`

SetInternalTime sets InternalTime field to given value.


### GetExternalCost

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetExternalCost() int32`

GetExternalCost returns the ExternalCost field if non-nil, zero value otherwise.

### GetExternalCostOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetExternalCostOk() (*int32, bool)`

GetExternalCostOk returns a tuple with the ExternalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCost

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetExternalCost(v int32)`

SetExternalCost sets ExternalCost field to given value.


### GetInternalCost

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetInternalCost() int32`

GetInternalCost returns the InternalCost field if non-nil, zero value otherwise.

### GetInternalCostOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetInternalCostOk() (*int32, bool)`

GetInternalCostOk returns a tuple with the InternalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalCost

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetInternalCost(v int32)`

SetInternalCost sets InternalCost field to given value.


### GetNeededImmediately

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetNeededImmediately() bool`

GetNeededImmediately returns the NeededImmediately field if non-nil, zero value otherwise.

### GetNeededImmediatelyOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetNeededImmediatelyOk() (*bool, bool)`

GetNeededImmediatelyOk returns a tuple with the NeededImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeededImmediately

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetNeededImmediately(v bool)`

SetNeededImmediately sets NeededImmediately field to given value.


### GetNeededTillKM

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetNeededTillKM() int32`

GetNeededTillKM returns the NeededTillKM field if non-nil, zero value otherwise.

### GetNeededTillKMOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetNeededTillKMOk() (*int32, bool)`

GetNeededTillKMOk returns a tuple with the NeededTillKM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeededTillKM

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetNeededTillKM(v int32)`

SetNeededTillKM sets NeededTillKM field to given value.


### SetNeededTillKMNil

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetNeededTillKMNil(b bool)`

 SetNeededTillKMNil sets the value for NeededTillKM to be an explicit nil

### UnsetNeededTillKM
`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) UnsetNeededTillKM()`

UnsetNeededTillKM ensures that no value is present for NeededTillKM, not even an explicit nil
### GetNeededTillDate

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetNeededTillDate() time.Time`

GetNeededTillDate returns the NeededTillDate field if non-nil, zero value otherwise.

### GetNeededTillDateOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetNeededTillDateOk() (*time.Time, bool)`

GetNeededTillDateOk returns a tuple with the NeededTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeededTillDate

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetNeededTillDate(v time.Time)`

SetNeededTillDate sets NeededTillDate field to given value.


### SetNeededTillDateNil

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetNeededTillDateNil(b bool)`

 SetNeededTillDateNil sets the value for NeededTillDate to be an explicit nil

### UnsetNeededTillDate
`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) UnsetNeededTillDate()`

UnsetNeededTillDate ensures that no value is present for NeededTillDate, not even an explicit nil
### GetRepeatTimeSpan

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetRepeatTimeSpan() string`

GetRepeatTimeSpan returns the RepeatTimeSpan field if non-nil, zero value otherwise.

### GetRepeatTimeSpanOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetRepeatTimeSpanOk() (*string, bool)`

GetRepeatTimeSpanOk returns a tuple with the RepeatTimeSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatTimeSpan

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetRepeatTimeSpan(v string)`

SetRepeatTimeSpan sets RepeatTimeSpan field to given value.


### SetRepeatTimeSpanNil

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetRepeatTimeSpanNil(b bool)`

 SetRepeatTimeSpanNil sets the value for RepeatTimeSpan to be an explicit nil

### UnsetRepeatTimeSpan
`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) UnsetRepeatTimeSpan()`

UnsetRepeatTimeSpan ensures that no value is present for RepeatTimeSpan, not even an explicit nil
### GetRepeatKM

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetRepeatKM() int32`

GetRepeatKM returns the RepeatKM field if non-nil, zero value otherwise.

### GetRepeatKMOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetRepeatKMOk() (*int32, bool)`

GetRepeatKMOk returns a tuple with the RepeatKM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatKM

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetRepeatKM(v int32)`

SetRepeatKM sets RepeatKM field to given value.


### SetRepeatKMNil

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetRepeatKMNil(b bool)`

 SetRepeatKMNil sets the value for RepeatKM to be an explicit nil

### UnsetRepeatKM
`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) UnsetRepeatKM()`

UnsetRepeatKM ensures that no value is present for RepeatKM, not even an explicit nil
### GetCurrentKM

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetCurrentKM() int32`

GetCurrentKM returns the CurrentKM field if non-nil, zero value otherwise.

### GetCurrentKMOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetCurrentKMOk() (*int32, bool)`

GetCurrentKMOk returns a tuple with the CurrentKM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentKM

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetCurrentKM(v int32)`

SetCurrentKM sets CurrentKM field to given value.


### GetGame

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetGame() FPHSpedVAPIEnumsGameEnum`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetGameOk() (*FPHSpedVAPIEnumsGameEnum, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetGame(v FPHSpedVAPIEnumsGameEnum)`

SetGame sets Game field to given value.


### GetKmSinceLast

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetKmSinceLast() int32`

GetKmSinceLast returns the KmSinceLast field if non-nil, zero value otherwise.

### GetKmSinceLastOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetKmSinceLastOk() (*int32, bool)`

GetKmSinceLastOk returns a tuple with the KmSinceLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmSinceLast

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetKmSinceLast(v int32)`

SetKmSinceLast sets KmSinceLast field to given value.


### SetKmSinceLastNil

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetKmSinceLastNil(b bool)`

 SetKmSinceLastNil sets the value for KmSinceLast to be an explicit nil

### UnsetKmSinceLast
`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) UnsetKmSinceLast()`

UnsetKmSinceLast ensures that no value is present for KmSinceLast, not even an explicit nil
### GetTimeSinceLast

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetTimeSinceLast() string`

GetTimeSinceLast returns the TimeSinceLast field if non-nil, zero value otherwise.

### GetTimeSinceLastOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetTimeSinceLastOk() (*string, bool)`

GetTimeSinceLastOk returns a tuple with the TimeSinceLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSinceLast

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetTimeSinceLast(v string)`

SetTimeSinceLast sets TimeSinceLast field to given value.


### SetTimeSinceLastNil

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetTimeSinceLastNil(b bool)`

 SetTimeSinceLastNil sets the value for TimeSinceLast to be an explicit nil

### UnsetTimeSinceLast
`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) UnsetTimeSinceLast()`

UnsetTimeSinceLast ensures that no value is present for TimeSinceLast, not even an explicit nil
### GetVisibleRemaining

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetVisibleRemaining() string`

GetVisibleRemaining returns the VisibleRemaining field if non-nil, zero value otherwise.

### GetVisibleRemainingOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetVisibleRemainingOk() (*string, bool)`

GetVisibleRemainingOk returns a tuple with the VisibleRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleRemaining

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetVisibleRemaining(v string)`

SetVisibleRemaining sets VisibleRemaining field to given value.


### SetVisibleRemainingNil

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetVisibleRemainingNil(b bool)`

 SetVisibleRemainingNil sets the value for VisibleRemaining to be an explicit nil

### UnsetVisibleRemaining
`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) UnsetVisibleRemaining()`

UnsetVisibleRemaining ensures that no value is present for VisibleRemaining, not even an explicit nil
### GetRepeatSpan

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetRepeatSpan() float64`

GetRepeatSpan returns the RepeatSpan field if non-nil, zero value otherwise.

### GetRepeatSpanOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetRepeatSpanOk() (*float64, bool)`

GetRepeatSpanOk returns a tuple with the RepeatSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatSpan

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetRepeatSpan(v float64)`

SetRepeatSpan sets RepeatSpan field to given value.


### GetSpanSinceLast

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetSpanSinceLast() float64`

GetSpanSinceLast returns the SpanSinceLast field if non-nil, zero value otherwise.

### GetSpanSinceLastOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetSpanSinceLastOk() (*float64, bool)`

GetSpanSinceLastOk returns a tuple with the SpanSinceLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanSinceLast

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetSpanSinceLast(v float64)`

SetSpanSinceLast sets SpanSinceLast field to given value.


### GetIsNeeded

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetIsNeeded() bool`

GetIsNeeded returns the IsNeeded field if non-nil, zero value otherwise.

### GetIsNeededOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetIsNeededOk() (*bool, bool)`

GetIsNeededOk returns a tuple with the IsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNeeded

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetIsNeeded(v bool)`

SetIsNeeded sets IsNeeded field to given value.


### GetIsOverdue

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetIsOverdue() bool`

GetIsOverdue returns the IsOverdue field if non-nil, zero value otherwise.

### GetIsOverdueOk

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetIsOverdueOk() (*bool, bool)`

GetIsOverdueOk returns a tuple with the IsOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverdue

`func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetIsOverdue(v bool)`

SetIsOverdue sets IsOverdue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


