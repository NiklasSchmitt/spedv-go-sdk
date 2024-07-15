# FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Station** | [**NullableFPHSpedVAPIObjectsSimRailSimRailStation**](FPHSpedVAPIObjectsSimRailSimRailStation.md) |  | [readonly] 
**ScheduledArrival** | **string** |  | [readonly] 
**ScheduledDeparture** | **string** |  | [readonly] 
**Platform** | **NullableString** |  | [readonly] 
**Track** | **NullableInt32** |  | [readonly] 
**StopType** | **NullableString** |  | [readonly] 
**Line** | **int32** |  | [readonly] 
**Km** | **float64** |  | [readonly] 
**MaxSpeed** | **int32** |  | [readonly] 
**RealArrival** | **NullableString** |  | [readonly] 
**RealDeparture** | **NullableString** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry

`func NewFPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry(station NullableFPHSpedVAPIObjectsSimRailSimRailStation, scheduledArrival string, scheduledDeparture string, platform NullableString, track NullableInt32, stopType NullableString, line int32, km float64, maxSpeed int32, realArrival NullableString, realDeparture NullableString, ) *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry`

NewFPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry instantiates a new FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntryWithDefaults

`func NewFPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntryWithDefaults() *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry`

NewFPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntryWithDefaults instantiates a new FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStation

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetStation() FPHSpedVAPIObjectsSimRailSimRailStation`

GetStation returns the Station field if non-nil, zero value otherwise.

### GetStationOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetStationOk() (*FPHSpedVAPIObjectsSimRailSimRailStation, bool)`

GetStationOk returns a tuple with the Station field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStation

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetStation(v FPHSpedVAPIObjectsSimRailSimRailStation)`

SetStation sets Station field to given value.


### SetStationNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetStationNil(b bool)`

 SetStationNil sets the value for Station to be an explicit nil

### UnsetStation
`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) UnsetStation()`

UnsetStation ensures that no value is present for Station, not even an explicit nil
### GetScheduledArrival

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetScheduledArrival() string`

GetScheduledArrival returns the ScheduledArrival field if non-nil, zero value otherwise.

### GetScheduledArrivalOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetScheduledArrivalOk() (*string, bool)`

GetScheduledArrivalOk returns a tuple with the ScheduledArrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledArrival

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetScheduledArrival(v string)`

SetScheduledArrival sets ScheduledArrival field to given value.


### GetScheduledDeparture

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetScheduledDeparture() string`

GetScheduledDeparture returns the ScheduledDeparture field if non-nil, zero value otherwise.

### GetScheduledDepartureOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetScheduledDepartureOk() (*string, bool)`

GetScheduledDepartureOk returns a tuple with the ScheduledDeparture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeparture

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetScheduledDeparture(v string)`

SetScheduledDeparture sets ScheduledDeparture field to given value.


### GetPlatform

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### SetPlatformNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetTrack

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetTrack() int32`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetTrackOk() (*int32, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetTrack(v int32)`

SetTrack sets Track field to given value.


### SetTrackNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetTrackNil(b bool)`

 SetTrackNil sets the value for Track to be an explicit nil

### UnsetTrack
`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) UnsetTrack()`

UnsetTrack ensures that no value is present for Track, not even an explicit nil
### GetStopType

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetStopType() string`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetStopTypeOk() (*string, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetStopType(v string)`

SetStopType sets StopType field to given value.


### SetStopTypeNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetStopTypeNil(b bool)`

 SetStopTypeNil sets the value for StopType to be an explicit nil

### UnsetStopType
`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) UnsetStopType()`

UnsetStopType ensures that no value is present for StopType, not even an explicit nil
### GetLine

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetLine(v int32)`

SetLine sets Line field to given value.


### GetKm

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetKm() float64`

GetKm returns the Km field if non-nil, zero value otherwise.

### GetKmOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetKmOk() (*float64, bool)`

GetKmOk returns a tuple with the Km field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKm

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetKm(v float64)`

SetKm sets Km field to given value.


### GetMaxSpeed

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetMaxSpeed() int32`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetMaxSpeedOk() (*int32, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetMaxSpeed(v int32)`

SetMaxSpeed sets MaxSpeed field to given value.


### GetRealArrival

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetRealArrival() string`

GetRealArrival returns the RealArrival field if non-nil, zero value otherwise.

### GetRealArrivalOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetRealArrivalOk() (*string, bool)`

GetRealArrivalOk returns a tuple with the RealArrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealArrival

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetRealArrival(v string)`

SetRealArrival sets RealArrival field to given value.


### SetRealArrivalNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetRealArrivalNil(b bool)`

 SetRealArrivalNil sets the value for RealArrival to be an explicit nil

### UnsetRealArrival
`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) UnsetRealArrival()`

UnsetRealArrival ensures that no value is present for RealArrival, not even an explicit nil
### GetRealDeparture

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetRealDeparture() string`

GetRealDeparture returns the RealDeparture field if non-nil, zero value otherwise.

### GetRealDepartureOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) GetRealDepartureOk() (*string, bool)`

GetRealDepartureOk returns a tuple with the RealDeparture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealDeparture

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetRealDeparture(v string)`

SetRealDeparture sets RealDeparture field to given value.


### SetRealDepartureNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) SetRealDepartureNil(b bool)`

 SetRealDepartureNil sets the value for RealDeparture to be an explicit nil

### UnsetRealDeparture
`func (o *FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) UnsetRealDeparture()`

UnsetRealDeparture ensures that no value is present for RealDeparture, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


