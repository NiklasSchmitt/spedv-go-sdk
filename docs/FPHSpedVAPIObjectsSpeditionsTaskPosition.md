# FPHSpedVAPIObjectsSpeditionsTaskPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Timestamp** | **time.Time** |  | 
**ChangeType** | [**FPHSpedVAPITaskPositionType**](FPHSpedVAPITaskPositionType.md) |   0 &#x3D; StateChange  1 &#x3D; RefusedTruckInfo  2 &#x3D; Refueled  3 &#x3D; FerryUsed  -1 &#x3D; NotSet | [readonly] 
**ChangedBy** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | 
**RefusedTruck** | [**NullableFPHSpedVAPIObjectsSpeditionsTruck**](FPHSpedVAPIObjectsSpeditionsTruck.md) |  | 
**StateChangedFrom** | [**NullableFPHSpedVAPIEnumsETSTaskState**](FPHSpedVAPIEnumsETSTaskState.md) |   0 &#x3D; InDrive  1 &#x3D; Done  2 &#x3D; Settled  3 &#x3D; Fail  4 &#x3D; AdminCheck  5 &#x3D; Paused  6 &#x3D; Cancelled  7 &#x3D; Invalid  -1 &#x3D; NotAvaliable | 
**StateChangeTo** | [**NullableFPHSpedVAPIEnumsETSTaskState**](FPHSpedVAPIEnumsETSTaskState.md) |   0 &#x3D; InDrive  1 &#x3D; Done  2 &#x3D; Settled  3 &#x3D; Fail  4 &#x3D; AdminCheck  5 &#x3D; Paused  6 &#x3D; Cancelled  7 &#x3D; Invalid  -1 &#x3D; NotAvaliable | 
**RefueledLiter** | **NullableFloat64** |  | 
**RefueledSum** | **NullableFloat64** |  | 
**UsedGasStation** | [**NullableFPHSpedVAPIObjectsMapsGasStation**](FPHSpedVAPIObjectsMapsGasStation.md) |  | [readonly] 
**UsedFerry** | [**NullableFPHSpedVAPIObjectsMapsFerry**](FPHSpedVAPIObjectsMapsFerry.md) |  | [readonly] 
**Coordinates** | [**NullableFPHSpedVAPIObjectsTelemetryETSVektor**](FPHSpedVAPIObjectsTelemetryETSVektor.md) |  | 
**Comment** | **NullableString** |  | 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsTaskPosition

`func NewFPHSpedVAPIObjectsSpeditionsTaskPosition(id int32, timestamp time.Time, changeType FPHSpedVAPITaskPositionType, changedBy NullableFPHSpedVAPIObjectsUsersUser, refusedTruck NullableFPHSpedVAPIObjectsSpeditionsTruck, stateChangedFrom NullableFPHSpedVAPIEnumsETSTaskState, stateChangeTo NullableFPHSpedVAPIEnumsETSTaskState, refueledLiter NullableFloat64, refueledSum NullableFloat64, usedGasStation NullableFPHSpedVAPIObjectsMapsGasStation, usedFerry NullableFPHSpedVAPIObjectsMapsFerry, coordinates NullableFPHSpedVAPIObjectsTelemetryETSVektor, comment NullableString, ) *FPHSpedVAPIObjectsSpeditionsTaskPosition`

NewFPHSpedVAPIObjectsSpeditionsTaskPosition instantiates a new FPHSpedVAPIObjectsSpeditionsTaskPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsTaskPositionWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsTaskPositionWithDefaults() *FPHSpedVAPIObjectsSpeditionsTaskPosition`

NewFPHSpedVAPIObjectsSpeditionsTaskPositionWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsTaskPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetId(v int32)`

SetId sets Id field to given value.


### GetTimestamp

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetChangeType

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetChangeType() FPHSpedVAPITaskPositionType`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetChangeTypeOk() (*FPHSpedVAPITaskPositionType, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetChangeType(v FPHSpedVAPITaskPositionType)`

SetChangeType sets ChangeType field to given value.


### GetChangedBy

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetChangedBy() FPHSpedVAPIObjectsUsersUser`

GetChangedBy returns the ChangedBy field if non-nil, zero value otherwise.

### GetChangedByOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetChangedByOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetChangedByOk returns a tuple with the ChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedBy

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetChangedBy(v FPHSpedVAPIObjectsUsersUser)`

SetChangedBy sets ChangedBy field to given value.


### SetChangedByNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetChangedByNil(b bool)`

 SetChangedByNil sets the value for ChangedBy to be an explicit nil

### UnsetChangedBy
`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) UnsetChangedBy()`

UnsetChangedBy ensures that no value is present for ChangedBy, not even an explicit nil
### GetRefusedTruck

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetRefusedTruck() FPHSpedVAPIObjectsSpeditionsTruck`

GetRefusedTruck returns the RefusedTruck field if non-nil, zero value otherwise.

### GetRefusedTruckOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetRefusedTruckOk() (*FPHSpedVAPIObjectsSpeditionsTruck, bool)`

GetRefusedTruckOk returns a tuple with the RefusedTruck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusedTruck

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetRefusedTruck(v FPHSpedVAPIObjectsSpeditionsTruck)`

SetRefusedTruck sets RefusedTruck field to given value.


### SetRefusedTruckNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetRefusedTruckNil(b bool)`

 SetRefusedTruckNil sets the value for RefusedTruck to be an explicit nil

### UnsetRefusedTruck
`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) UnsetRefusedTruck()`

UnsetRefusedTruck ensures that no value is present for RefusedTruck, not even an explicit nil
### GetStateChangedFrom

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetStateChangedFrom() FPHSpedVAPIEnumsETSTaskState`

GetStateChangedFrom returns the StateChangedFrom field if non-nil, zero value otherwise.

### GetStateChangedFromOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetStateChangedFromOk() (*FPHSpedVAPIEnumsETSTaskState, bool)`

GetStateChangedFromOk returns a tuple with the StateChangedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangedFrom

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetStateChangedFrom(v FPHSpedVAPIEnumsETSTaskState)`

SetStateChangedFrom sets StateChangedFrom field to given value.


### SetStateChangedFromNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetStateChangedFromNil(b bool)`

 SetStateChangedFromNil sets the value for StateChangedFrom to be an explicit nil

### UnsetStateChangedFrom
`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) UnsetStateChangedFrom()`

UnsetStateChangedFrom ensures that no value is present for StateChangedFrom, not even an explicit nil
### GetStateChangeTo

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetStateChangeTo() FPHSpedVAPIEnumsETSTaskState`

GetStateChangeTo returns the StateChangeTo field if non-nil, zero value otherwise.

### GetStateChangeToOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetStateChangeToOk() (*FPHSpedVAPIEnumsETSTaskState, bool)`

GetStateChangeToOk returns a tuple with the StateChangeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangeTo

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetStateChangeTo(v FPHSpedVAPIEnumsETSTaskState)`

SetStateChangeTo sets StateChangeTo field to given value.


### SetStateChangeToNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetStateChangeToNil(b bool)`

 SetStateChangeToNil sets the value for StateChangeTo to be an explicit nil

### UnsetStateChangeTo
`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) UnsetStateChangeTo()`

UnsetStateChangeTo ensures that no value is present for StateChangeTo, not even an explicit nil
### GetRefueledLiter

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetRefueledLiter() float64`

GetRefueledLiter returns the RefueledLiter field if non-nil, zero value otherwise.

### GetRefueledLiterOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetRefueledLiterOk() (*float64, bool)`

GetRefueledLiterOk returns a tuple with the RefueledLiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefueledLiter

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetRefueledLiter(v float64)`

SetRefueledLiter sets RefueledLiter field to given value.


### SetRefueledLiterNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetRefueledLiterNil(b bool)`

 SetRefueledLiterNil sets the value for RefueledLiter to be an explicit nil

### UnsetRefueledLiter
`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) UnsetRefueledLiter()`

UnsetRefueledLiter ensures that no value is present for RefueledLiter, not even an explicit nil
### GetRefueledSum

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetRefueledSum() float64`

GetRefueledSum returns the RefueledSum field if non-nil, zero value otherwise.

### GetRefueledSumOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetRefueledSumOk() (*float64, bool)`

GetRefueledSumOk returns a tuple with the RefueledSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefueledSum

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetRefueledSum(v float64)`

SetRefueledSum sets RefueledSum field to given value.


### SetRefueledSumNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetRefueledSumNil(b bool)`

 SetRefueledSumNil sets the value for RefueledSum to be an explicit nil

### UnsetRefueledSum
`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) UnsetRefueledSum()`

UnsetRefueledSum ensures that no value is present for RefueledSum, not even an explicit nil
### GetUsedGasStation

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetUsedGasStation() FPHSpedVAPIObjectsMapsGasStation`

GetUsedGasStation returns the UsedGasStation field if non-nil, zero value otherwise.

### GetUsedGasStationOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetUsedGasStationOk() (*FPHSpedVAPIObjectsMapsGasStation, bool)`

GetUsedGasStationOk returns a tuple with the UsedGasStation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedGasStation

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetUsedGasStation(v FPHSpedVAPIObjectsMapsGasStation)`

SetUsedGasStation sets UsedGasStation field to given value.


### SetUsedGasStationNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetUsedGasStationNil(b bool)`

 SetUsedGasStationNil sets the value for UsedGasStation to be an explicit nil

### UnsetUsedGasStation
`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) UnsetUsedGasStation()`

UnsetUsedGasStation ensures that no value is present for UsedGasStation, not even an explicit nil
### GetUsedFerry

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetUsedFerry() FPHSpedVAPIObjectsMapsFerry`

GetUsedFerry returns the UsedFerry field if non-nil, zero value otherwise.

### GetUsedFerryOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetUsedFerryOk() (*FPHSpedVAPIObjectsMapsFerry, bool)`

GetUsedFerryOk returns a tuple with the UsedFerry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedFerry

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetUsedFerry(v FPHSpedVAPIObjectsMapsFerry)`

SetUsedFerry sets UsedFerry field to given value.


### SetUsedFerryNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetUsedFerryNil(b bool)`

 SetUsedFerryNil sets the value for UsedFerry to be an explicit nil

### UnsetUsedFerry
`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) UnsetUsedFerry()`

UnsetUsedFerry ensures that no value is present for UsedFerry, not even an explicit nil
### GetCoordinates

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetCoordinates() FPHSpedVAPIObjectsTelemetryETSVektor`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetCoordinatesOk() (*FPHSpedVAPIObjectsTelemetryETSVektor, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetCoordinates(v FPHSpedVAPIObjectsTelemetryETSVektor)`

SetCoordinates sets Coordinates field to given value.


### SetCoordinatesNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetCoordinatesNil(b bool)`

 SetCoordinatesNil sets the value for Coordinates to be an explicit nil

### UnsetCoordinates
`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) UnsetCoordinates()`

UnsetCoordinates ensures that no value is present for Coordinates, not even an explicit nil
### GetComment

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FPHSpedVAPIObjectsSpeditionsTaskPosition) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


