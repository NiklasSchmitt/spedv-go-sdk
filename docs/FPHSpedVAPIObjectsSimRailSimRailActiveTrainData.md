# FPHSpedVAPIObjectsSimRailSimRailActiveTrainData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrainName** | **NullableString** |  | [readonly] 
**TrainNumber** | **NullableString** |  | [readonly] 
**TrainNumberInternational** | **NullableString** |  | [readonly] 
**TrainLength** | **int32** |  | [readonly] 
**TrainWeight** | **int32** |  | [readonly] 
**ContinuesAs** | **NullableString** |  | [readonly] 
**ServerData** | [**NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo**](FPHSpedVAPIObjectsSimRailSimRailServerInfo.md) |  | [readonly] 
**Id** | **NullableString** |  | [readonly] 
**StartStation** | **NullableString** |  | [readonly] 
**EndStation** | **NullableString** |  | [readonly] 
**Vehicles** | **[]string** |  | [readonly] 
**ControlType** | [**FPHSpedVAPIObjectsSimRailSimRailTrainControlType**](FPHSpedVAPIObjectsSimRailSimRailTrainControlType.md) |   0 &#x3D; Ghost  1 &#x3D; Bot  2 &#x3D; Player | [readonly] 
**ControlledBySteamID** | **NullableInt64** |  | [readonly] 
**Latitude** | **float64** |  | [readonly] 
**Longitude** | **float64** |  | [readonly] 
**Velocity** | **float64** |  | [readonly] 
**NextSignal** | **NullableString** |  | [readonly] 
**DistanceToNextSignal** | **float64** |  | [readonly] 
**SignalInFrontSpeed** | **int32** |  | [readonly] 
**ApiTimetableIndex** | **int32** |  | [readonly] 
**LastRealTimetableIndex** | **int32** |  | [readonly] 
**CurrentPositionType** | [**FPHSpedVAPIObjectsSimRailSimRailTrainPositionType**](FPHSpedVAPIObjectsSimRailSimRailTrainPositionType.md) |   0 &#x3D; InStation  1 &#x3D; BetweenStations  2 &#x3D; ApproachingStation | [readonly] 
**Timetable** | [**[]FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry**](FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry.md) |  | [readonly] 
**CurrentTimetableIndex** | **int32** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSimRailSimRailActiveTrainData

`func NewFPHSpedVAPIObjectsSimRailSimRailActiveTrainData(trainName NullableString, trainNumber NullableString, trainNumberInternational NullableString, trainLength int32, trainWeight int32, continuesAs NullableString, serverData NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo, id NullableString, startStation NullableString, endStation NullableString, vehicles []string, controlType FPHSpedVAPIObjectsSimRailSimRailTrainControlType, controlledBySteamID NullableInt64, latitude float64, longitude float64, velocity float64, nextSignal NullableString, distanceToNextSignal float64, signalInFrontSpeed int32, apiTimetableIndex int32, lastRealTimetableIndex int32, currentPositionType FPHSpedVAPIObjectsSimRailSimRailTrainPositionType, timetable []FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry, currentTimetableIndex int32, ) *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData`

NewFPHSpedVAPIObjectsSimRailSimRailActiveTrainData instantiates a new FPHSpedVAPIObjectsSimRailSimRailActiveTrainData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSimRailSimRailActiveTrainDataWithDefaults

`func NewFPHSpedVAPIObjectsSimRailSimRailActiveTrainDataWithDefaults() *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData`

NewFPHSpedVAPIObjectsSimRailSimRailActiveTrainDataWithDefaults instantiates a new FPHSpedVAPIObjectsSimRailSimRailActiveTrainData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrainName

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainName() string`

GetTrainName returns the TrainName field if non-nil, zero value otherwise.

### GetTrainNameOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainNameOk() (*string, bool)`

GetTrainNameOk returns a tuple with the TrainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainName

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainName(v string)`

SetTrainName sets TrainName field to given value.


### SetTrainNameNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainNameNil(b bool)`

 SetTrainNameNil sets the value for TrainName to be an explicit nil

### UnsetTrainName
`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnsetTrainName()`

UnsetTrainName ensures that no value is present for TrainName, not even an explicit nil
### GetTrainNumber

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainNumber() string`

GetTrainNumber returns the TrainNumber field if non-nil, zero value otherwise.

### GetTrainNumberOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainNumberOk() (*string, bool)`

GetTrainNumberOk returns a tuple with the TrainNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainNumber

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainNumber(v string)`

SetTrainNumber sets TrainNumber field to given value.


### SetTrainNumberNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainNumberNil(b bool)`

 SetTrainNumberNil sets the value for TrainNumber to be an explicit nil

### UnsetTrainNumber
`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnsetTrainNumber()`

UnsetTrainNumber ensures that no value is present for TrainNumber, not even an explicit nil
### GetTrainNumberInternational

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainNumberInternational() string`

GetTrainNumberInternational returns the TrainNumberInternational field if non-nil, zero value otherwise.

### GetTrainNumberInternationalOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainNumberInternationalOk() (*string, bool)`

GetTrainNumberInternationalOk returns a tuple with the TrainNumberInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainNumberInternational

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainNumberInternational(v string)`

SetTrainNumberInternational sets TrainNumberInternational field to given value.


### SetTrainNumberInternationalNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainNumberInternationalNil(b bool)`

 SetTrainNumberInternationalNil sets the value for TrainNumberInternational to be an explicit nil

### UnsetTrainNumberInternational
`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnsetTrainNumberInternational()`

UnsetTrainNumberInternational ensures that no value is present for TrainNumberInternational, not even an explicit nil
### GetTrainLength

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainLength() int32`

GetTrainLength returns the TrainLength field if non-nil, zero value otherwise.

### GetTrainLengthOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainLengthOk() (*int32, bool)`

GetTrainLengthOk returns a tuple with the TrainLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainLength

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainLength(v int32)`

SetTrainLength sets TrainLength field to given value.


### GetTrainWeight

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainWeight() int32`

GetTrainWeight returns the TrainWeight field if non-nil, zero value otherwise.

### GetTrainWeightOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainWeightOk() (*int32, bool)`

GetTrainWeightOk returns a tuple with the TrainWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainWeight

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainWeight(v int32)`

SetTrainWeight sets TrainWeight field to given value.


### GetContinuesAs

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetContinuesAs() string`

GetContinuesAs returns the ContinuesAs field if non-nil, zero value otherwise.

### GetContinuesAsOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetContinuesAsOk() (*string, bool)`

GetContinuesAsOk returns a tuple with the ContinuesAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuesAs

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetContinuesAs(v string)`

SetContinuesAs sets ContinuesAs field to given value.


### SetContinuesAsNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetContinuesAsNil(b bool)`

 SetContinuesAsNil sets the value for ContinuesAs to be an explicit nil

### UnsetContinuesAs
`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnsetContinuesAs()`

UnsetContinuesAs ensures that no value is present for ContinuesAs, not even an explicit nil
### GetServerData

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetServerData() FPHSpedVAPIObjectsSimRailSimRailServerInfo`

GetServerData returns the ServerData field if non-nil, zero value otherwise.

### GetServerDataOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetServerDataOk() (*FPHSpedVAPIObjectsSimRailSimRailServerInfo, bool)`

GetServerDataOk returns a tuple with the ServerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerData

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetServerData(v FPHSpedVAPIObjectsSimRailSimRailServerInfo)`

SetServerData sets ServerData field to given value.


### SetServerDataNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetServerDataNil(b bool)`

 SetServerDataNil sets the value for ServerData to be an explicit nil

### UnsetServerData
`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnsetServerData()`

UnsetServerData ensures that no value is present for ServerData, not even an explicit nil
### GetId

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetStartStation

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetStartStation() string`

GetStartStation returns the StartStation field if non-nil, zero value otherwise.

### GetStartStationOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetStartStationOk() (*string, bool)`

GetStartStationOk returns a tuple with the StartStation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartStation

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetStartStation(v string)`

SetStartStation sets StartStation field to given value.


### SetStartStationNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetStartStationNil(b bool)`

 SetStartStationNil sets the value for StartStation to be an explicit nil

### UnsetStartStation
`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnsetStartStation()`

UnsetStartStation ensures that no value is present for StartStation, not even an explicit nil
### GetEndStation

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetEndStation() string`

GetEndStation returns the EndStation field if non-nil, zero value otherwise.

### GetEndStationOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetEndStationOk() (*string, bool)`

GetEndStationOk returns a tuple with the EndStation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndStation

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetEndStation(v string)`

SetEndStation sets EndStation field to given value.


### SetEndStationNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetEndStationNil(b bool)`

 SetEndStationNil sets the value for EndStation to be an explicit nil

### UnsetEndStation
`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnsetEndStation()`

UnsetEndStation ensures that no value is present for EndStation, not even an explicit nil
### GetVehicles

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetVehicles() []string`

GetVehicles returns the Vehicles field if non-nil, zero value otherwise.

### GetVehiclesOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetVehiclesOk() (*[]string, bool)`

GetVehiclesOk returns a tuple with the Vehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicles

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetVehicles(v []string)`

SetVehicles sets Vehicles field to given value.


### SetVehiclesNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetVehiclesNil(b bool)`

 SetVehiclesNil sets the value for Vehicles to be an explicit nil

### UnsetVehicles
`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnsetVehicles()`

UnsetVehicles ensures that no value is present for Vehicles, not even an explicit nil
### GetControlType

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetControlType() FPHSpedVAPIObjectsSimRailSimRailTrainControlType`

GetControlType returns the ControlType field if non-nil, zero value otherwise.

### GetControlTypeOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetControlTypeOk() (*FPHSpedVAPIObjectsSimRailSimRailTrainControlType, bool)`

GetControlTypeOk returns a tuple with the ControlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlType

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetControlType(v FPHSpedVAPIObjectsSimRailSimRailTrainControlType)`

SetControlType sets ControlType field to given value.


### GetControlledBySteamID

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetControlledBySteamID() int64`

GetControlledBySteamID returns the ControlledBySteamID field if non-nil, zero value otherwise.

### GetControlledBySteamIDOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetControlledBySteamIDOk() (*int64, bool)`

GetControlledBySteamIDOk returns a tuple with the ControlledBySteamID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlledBySteamID

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetControlledBySteamID(v int64)`

SetControlledBySteamID sets ControlledBySteamID field to given value.


### SetControlledBySteamIDNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetControlledBySteamIDNil(b bool)`

 SetControlledBySteamIDNil sets the value for ControlledBySteamID to be an explicit nil

### UnsetControlledBySteamID
`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnsetControlledBySteamID()`

UnsetControlledBySteamID ensures that no value is present for ControlledBySteamID, not even an explicit nil
### GetLatitude

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetVelocity

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetVelocity() float64`

GetVelocity returns the Velocity field if non-nil, zero value otherwise.

### GetVelocityOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetVelocityOk() (*float64, bool)`

GetVelocityOk returns a tuple with the Velocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocity

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetVelocity(v float64)`

SetVelocity sets Velocity field to given value.


### GetNextSignal

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetNextSignal() string`

GetNextSignal returns the NextSignal field if non-nil, zero value otherwise.

### GetNextSignalOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetNextSignalOk() (*string, bool)`

GetNextSignalOk returns a tuple with the NextSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSignal

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetNextSignal(v string)`

SetNextSignal sets NextSignal field to given value.


### SetNextSignalNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetNextSignalNil(b bool)`

 SetNextSignalNil sets the value for NextSignal to be an explicit nil

### UnsetNextSignal
`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnsetNextSignal()`

UnsetNextSignal ensures that no value is present for NextSignal, not even an explicit nil
### GetDistanceToNextSignal

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetDistanceToNextSignal() float64`

GetDistanceToNextSignal returns the DistanceToNextSignal field if non-nil, zero value otherwise.

### GetDistanceToNextSignalOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetDistanceToNextSignalOk() (*float64, bool)`

GetDistanceToNextSignalOk returns a tuple with the DistanceToNextSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceToNextSignal

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetDistanceToNextSignal(v float64)`

SetDistanceToNextSignal sets DistanceToNextSignal field to given value.


### GetSignalInFrontSpeed

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetSignalInFrontSpeed() int32`

GetSignalInFrontSpeed returns the SignalInFrontSpeed field if non-nil, zero value otherwise.

### GetSignalInFrontSpeedOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetSignalInFrontSpeedOk() (*int32, bool)`

GetSignalInFrontSpeedOk returns a tuple with the SignalInFrontSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalInFrontSpeed

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetSignalInFrontSpeed(v int32)`

SetSignalInFrontSpeed sets SignalInFrontSpeed field to given value.


### GetApiTimetableIndex

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetApiTimetableIndex() int32`

GetApiTimetableIndex returns the ApiTimetableIndex field if non-nil, zero value otherwise.

### GetApiTimetableIndexOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetApiTimetableIndexOk() (*int32, bool)`

GetApiTimetableIndexOk returns a tuple with the ApiTimetableIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTimetableIndex

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetApiTimetableIndex(v int32)`

SetApiTimetableIndex sets ApiTimetableIndex field to given value.


### GetLastRealTimetableIndex

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetLastRealTimetableIndex() int32`

GetLastRealTimetableIndex returns the LastRealTimetableIndex field if non-nil, zero value otherwise.

### GetLastRealTimetableIndexOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetLastRealTimetableIndexOk() (*int32, bool)`

GetLastRealTimetableIndexOk returns a tuple with the LastRealTimetableIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRealTimetableIndex

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetLastRealTimetableIndex(v int32)`

SetLastRealTimetableIndex sets LastRealTimetableIndex field to given value.


### GetCurrentPositionType

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetCurrentPositionType() FPHSpedVAPIObjectsSimRailSimRailTrainPositionType`

GetCurrentPositionType returns the CurrentPositionType field if non-nil, zero value otherwise.

### GetCurrentPositionTypeOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetCurrentPositionTypeOk() (*FPHSpedVAPIObjectsSimRailSimRailTrainPositionType, bool)`

GetCurrentPositionTypeOk returns a tuple with the CurrentPositionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPositionType

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetCurrentPositionType(v FPHSpedVAPIObjectsSimRailSimRailTrainPositionType)`

SetCurrentPositionType sets CurrentPositionType field to given value.


### GetTimetable

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTimetable() []FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry`

GetTimetable returns the Timetable field if non-nil, zero value otherwise.

### GetTimetableOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTimetableOk() (*[]FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry, bool)`

GetTimetableOk returns a tuple with the Timetable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimetable

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTimetable(v []FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry)`

SetTimetable sets Timetable field to given value.


### SetTimetableNil

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTimetableNil(b bool)`

 SetTimetableNil sets the value for Timetable to be an explicit nil

### UnsetTimetable
`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnsetTimetable()`

UnsetTimetable ensures that no value is present for Timetable, not even an explicit nil
### GetCurrentTimetableIndex

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetCurrentTimetableIndex() int32`

GetCurrentTimetableIndex returns the CurrentTimetableIndex field if non-nil, zero value otherwise.

### GetCurrentTimetableIndexOk

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetCurrentTimetableIndexOk() (*int32, bool)`

GetCurrentTimetableIndexOk returns a tuple with the CurrentTimetableIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTimetableIndex

`func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetCurrentTimetableIndex(v int32)`

SetCurrentTimetableIndex sets CurrentTimetableIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


