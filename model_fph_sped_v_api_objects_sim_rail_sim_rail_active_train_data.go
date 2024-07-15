/*
SpedV API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FPHSpedVAPIObjectsSimRailSimRailActiveTrainData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsSimRailSimRailActiveTrainData{}

// FPHSpedVAPIObjectsSimRailSimRailActiveTrainData struct for FPHSpedVAPIObjectsSimRailSimRailActiveTrainData
type FPHSpedVAPIObjectsSimRailSimRailActiveTrainData struct {
	TrainName NullableString `json:"trainName"`
	TrainNumber NullableString `json:"trainNumber"`
	TrainNumberInternational NullableString `json:"trainNumberInternational"`
	TrainLength int32 `json:"trainLength"`
	TrainWeight int32 `json:"trainWeight"`
	ContinuesAs NullableString `json:"continuesAs"`
	ServerData NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo `json:"serverData"`
	Id NullableString `json:"id"`
	StartStation NullableString `json:"startStation"`
	EndStation NullableString `json:"endStation"`
	Vehicles []string `json:"vehicles"`
	//   0 = Ghost  1 = Bot  2 = Player
	ControlType FPHSpedVAPIObjectsSimRailSimRailTrainControlType `json:"controlType"`
	ControlledBySteamID NullableInt64 `json:"controlledBySteamID"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Velocity float64 `json:"velocity"`
	NextSignal NullableString `json:"nextSignal"`
	DistanceToNextSignal float64 `json:"distanceToNextSignal"`
	SignalInFrontSpeed int32 `json:"signalInFrontSpeed"`
	ApiTimetableIndex int32 `json:"apiTimetableIndex"`
	LastRealTimetableIndex int32 `json:"lastRealTimetableIndex"`
	//   0 = InStation  1 = BetweenStations  2 = ApproachingStation
	CurrentPositionType FPHSpedVAPIObjectsSimRailSimRailTrainPositionType `json:"currentPositionType"`
	Timetable []FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry `json:"timetable"`
	CurrentTimetableIndex int32 `json:"currentTimetableIndex"`
}

type _FPHSpedVAPIObjectsSimRailSimRailActiveTrainData FPHSpedVAPIObjectsSimRailSimRailActiveTrainData

// NewFPHSpedVAPIObjectsSimRailSimRailActiveTrainData instantiates a new FPHSpedVAPIObjectsSimRailSimRailActiveTrainData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsSimRailSimRailActiveTrainData(trainName NullableString, trainNumber NullableString, trainNumberInternational NullableString, trainLength int32, trainWeight int32, continuesAs NullableString, serverData NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo, id NullableString, startStation NullableString, endStation NullableString, vehicles []string, controlType FPHSpedVAPIObjectsSimRailSimRailTrainControlType, controlledBySteamID NullableInt64, latitude float64, longitude float64, velocity float64, nextSignal NullableString, distanceToNextSignal float64, signalInFrontSpeed int32, apiTimetableIndex int32, lastRealTimetableIndex int32, currentPositionType FPHSpedVAPIObjectsSimRailSimRailTrainPositionType, timetable []FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry, currentTimetableIndex int32) *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData {
	this := FPHSpedVAPIObjectsSimRailSimRailActiveTrainData{}
	this.TrainName = trainName
	this.TrainNumber = trainNumber
	this.TrainNumberInternational = trainNumberInternational
	this.TrainLength = trainLength
	this.TrainWeight = trainWeight
	this.ContinuesAs = continuesAs
	this.ServerData = serverData
	this.Id = id
	this.StartStation = startStation
	this.EndStation = endStation
	this.Vehicles = vehicles
	this.ControlType = controlType
	this.ControlledBySteamID = controlledBySteamID
	this.Latitude = latitude
	this.Longitude = longitude
	this.Velocity = velocity
	this.NextSignal = nextSignal
	this.DistanceToNextSignal = distanceToNextSignal
	this.SignalInFrontSpeed = signalInFrontSpeed
	this.ApiTimetableIndex = apiTimetableIndex
	this.LastRealTimetableIndex = lastRealTimetableIndex
	this.CurrentPositionType = currentPositionType
	this.Timetable = timetable
	this.CurrentTimetableIndex = currentTimetableIndex
	return &this
}

// NewFPHSpedVAPIObjectsSimRailSimRailActiveTrainDataWithDefaults instantiates a new FPHSpedVAPIObjectsSimRailSimRailActiveTrainData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsSimRailSimRailActiveTrainDataWithDefaults() *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData {
	this := FPHSpedVAPIObjectsSimRailSimRailActiveTrainData{}
	return &this
}

// GetTrainName returns the TrainName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainName() string {
	if o == nil || o.TrainName.Get() == nil {
		var ret string
		return ret
	}

	return *o.TrainName.Get()
}

// GetTrainNameOk returns a tuple with the TrainName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrainName.Get(), o.TrainName.IsSet()
}

// SetTrainName sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainName(v string) {
	o.TrainName.Set(&v)
}

// GetTrainNumber returns the TrainNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainNumber() string {
	if o == nil || o.TrainNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.TrainNumber.Get()
}

// GetTrainNumberOk returns a tuple with the TrainNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrainNumber.Get(), o.TrainNumber.IsSet()
}

// SetTrainNumber sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainNumber(v string) {
	o.TrainNumber.Set(&v)
}

// GetTrainNumberInternational returns the TrainNumberInternational field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainNumberInternational() string {
	if o == nil || o.TrainNumberInternational.Get() == nil {
		var ret string
		return ret
	}

	return *o.TrainNumberInternational.Get()
}

// GetTrainNumberInternationalOk returns a tuple with the TrainNumberInternational field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainNumberInternationalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrainNumberInternational.Get(), o.TrainNumberInternational.IsSet()
}

// SetTrainNumberInternational sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainNumberInternational(v string) {
	o.TrainNumberInternational.Set(&v)
}

// GetTrainLength returns the TrainLength field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrainLength
}

// GetTrainLengthOk returns a tuple with the TrainLength field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrainLength, true
}

// SetTrainLength sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainLength(v int32) {
	o.TrainLength = v
}

// GetTrainWeight returns the TrainWeight field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrainWeight
}

// GetTrainWeightOk returns a tuple with the TrainWeight field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTrainWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrainWeight, true
}

// SetTrainWeight sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTrainWeight(v int32) {
	o.TrainWeight = v
}

// GetContinuesAs returns the ContinuesAs field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetContinuesAs() string {
	if o == nil || o.ContinuesAs.Get() == nil {
		var ret string
		return ret
	}

	return *o.ContinuesAs.Get()
}

// GetContinuesAsOk returns a tuple with the ContinuesAs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetContinuesAsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinuesAs.Get(), o.ContinuesAs.IsSet()
}

// SetContinuesAs sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetContinuesAs(v string) {
	o.ContinuesAs.Set(&v)
}

// GetServerData returns the ServerData field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsSimRailSimRailServerInfo will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetServerData() FPHSpedVAPIObjectsSimRailSimRailServerInfo {
	if o == nil || o.ServerData.Get() == nil {
		var ret FPHSpedVAPIObjectsSimRailSimRailServerInfo
		return ret
	}

	return *o.ServerData.Get()
}

// GetServerDataOk returns a tuple with the ServerData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetServerDataOk() (*FPHSpedVAPIObjectsSimRailSimRailServerInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerData.Get(), o.ServerData.IsSet()
}

// SetServerData sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetServerData(v FPHSpedVAPIObjectsSimRailSimRailServerInfo) {
	o.ServerData.Set(&v)
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetId(v string) {
	o.Id.Set(&v)
}

// GetStartStation returns the StartStation field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetStartStation() string {
	if o == nil || o.StartStation.Get() == nil {
		var ret string
		return ret
	}

	return *o.StartStation.Get()
}

// GetStartStationOk returns a tuple with the StartStation field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetStartStationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartStation.Get(), o.StartStation.IsSet()
}

// SetStartStation sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetStartStation(v string) {
	o.StartStation.Set(&v)
}

// GetEndStation returns the EndStation field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetEndStation() string {
	if o == nil || o.EndStation.Get() == nil {
		var ret string
		return ret
	}

	return *o.EndStation.Get()
}

// GetEndStationOk returns a tuple with the EndStation field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetEndStationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndStation.Get(), o.EndStation.IsSet()
}

// SetEndStation sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetEndStation(v string) {
	o.EndStation.Set(&v)
}

// GetVehicles returns the Vehicles field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetVehicles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Vehicles
}

// GetVehiclesOk returns a tuple with the Vehicles field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetVehiclesOk() ([]string, bool) {
	if o == nil || IsNil(o.Vehicles) {
		return nil, false
	}
	return o.Vehicles, true
}

// SetVehicles sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetVehicles(v []string) {
	o.Vehicles = v
}

// GetControlType returns the ControlType field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetControlType() FPHSpedVAPIObjectsSimRailSimRailTrainControlType {
	if o == nil {
		var ret FPHSpedVAPIObjectsSimRailSimRailTrainControlType
		return ret
	}

	return o.ControlType
}

// GetControlTypeOk returns a tuple with the ControlType field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetControlTypeOk() (*FPHSpedVAPIObjectsSimRailSimRailTrainControlType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ControlType, true
}

// SetControlType sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetControlType(v FPHSpedVAPIObjectsSimRailSimRailTrainControlType) {
	o.ControlType = v
}

// GetControlledBySteamID returns the ControlledBySteamID field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetControlledBySteamID() int64 {
	if o == nil || o.ControlledBySteamID.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ControlledBySteamID.Get()
}

// GetControlledBySteamIDOk returns a tuple with the ControlledBySteamID field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetControlledBySteamIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ControlledBySteamID.Get(), o.ControlledBySteamID.IsSet()
}

// SetControlledBySteamID sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetControlledBySteamID(v int64) {
	o.ControlledBySteamID.Set(&v)
}

// GetLatitude returns the Latitude field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetLatitude(v float64) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetLongitude(v float64) {
	o.Longitude = v
}

// GetVelocity returns the Velocity field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetVelocity() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Velocity
}

// GetVelocityOk returns a tuple with the Velocity field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetVelocityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Velocity, true
}

// SetVelocity sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetVelocity(v float64) {
	o.Velocity = v
}

// GetNextSignal returns the NextSignal field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetNextSignal() string {
	if o == nil || o.NextSignal.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextSignal.Get()
}

// GetNextSignalOk returns a tuple with the NextSignal field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetNextSignalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextSignal.Get(), o.NextSignal.IsSet()
}

// SetNextSignal sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetNextSignal(v string) {
	o.NextSignal.Set(&v)
}

// GetDistanceToNextSignal returns the DistanceToNextSignal field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetDistanceToNextSignal() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DistanceToNextSignal
}

// GetDistanceToNextSignalOk returns a tuple with the DistanceToNextSignal field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetDistanceToNextSignalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistanceToNextSignal, true
}

// SetDistanceToNextSignal sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetDistanceToNextSignal(v float64) {
	o.DistanceToNextSignal = v
}

// GetSignalInFrontSpeed returns the SignalInFrontSpeed field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetSignalInFrontSpeed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SignalInFrontSpeed
}

// GetSignalInFrontSpeedOk returns a tuple with the SignalInFrontSpeed field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetSignalInFrontSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalInFrontSpeed, true
}

// SetSignalInFrontSpeed sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetSignalInFrontSpeed(v int32) {
	o.SignalInFrontSpeed = v
}

// GetApiTimetableIndex returns the ApiTimetableIndex field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetApiTimetableIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApiTimetableIndex
}

// GetApiTimetableIndexOk returns a tuple with the ApiTimetableIndex field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetApiTimetableIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiTimetableIndex, true
}

// SetApiTimetableIndex sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetApiTimetableIndex(v int32) {
	o.ApiTimetableIndex = v
}

// GetLastRealTimetableIndex returns the LastRealTimetableIndex field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetLastRealTimetableIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LastRealTimetableIndex
}

// GetLastRealTimetableIndexOk returns a tuple with the LastRealTimetableIndex field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetLastRealTimetableIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastRealTimetableIndex, true
}

// SetLastRealTimetableIndex sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetLastRealTimetableIndex(v int32) {
	o.LastRealTimetableIndex = v
}

// GetCurrentPositionType returns the CurrentPositionType field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetCurrentPositionType() FPHSpedVAPIObjectsSimRailSimRailTrainPositionType {
	if o == nil {
		var ret FPHSpedVAPIObjectsSimRailSimRailTrainPositionType
		return ret
	}

	return o.CurrentPositionType
}

// GetCurrentPositionTypeOk returns a tuple with the CurrentPositionType field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetCurrentPositionTypeOk() (*FPHSpedVAPIObjectsSimRailSimRailTrainPositionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPositionType, true
}

// SetCurrentPositionType sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetCurrentPositionType(v FPHSpedVAPIObjectsSimRailSimRailTrainPositionType) {
	o.CurrentPositionType = v
}

// GetTimetable returns the Timetable field value
// If the value is explicit nil, the zero value for []FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTimetable() []FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry {
	if o == nil {
		var ret []FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry
		return ret
	}

	return o.Timetable
}

// GetTimetableOk returns a tuple with the Timetable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetTimetableOk() ([]FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry, bool) {
	if o == nil || IsNil(o.Timetable) {
		return nil, false
	}
	return o.Timetable, true
}

// SetTimetable sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetTimetable(v []FPHSpedVAPIObjectsSimRailSimRailTrainLiveTimetableEntry) {
	o.Timetable = v
}

// GetCurrentTimetableIndex returns the CurrentTimetableIndex field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetCurrentTimetableIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentTimetableIndex
}

// GetCurrentTimetableIndexOk returns a tuple with the CurrentTimetableIndex field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) GetCurrentTimetableIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentTimetableIndex, true
}

// SetCurrentTimetableIndex sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) SetCurrentTimetableIndex(v int32) {
	o.CurrentTimetableIndex = v
}

func (o FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trainName"] = o.TrainName.Get()
	toSerialize["trainNumber"] = o.TrainNumber.Get()
	toSerialize["trainNumberInternational"] = o.TrainNumberInternational.Get()
	toSerialize["trainLength"] = o.TrainLength
	toSerialize["trainWeight"] = o.TrainWeight
	toSerialize["continuesAs"] = o.ContinuesAs.Get()
	toSerialize["serverData"] = o.ServerData.Get()
	toSerialize["id"] = o.Id.Get()
	toSerialize["startStation"] = o.StartStation.Get()
	toSerialize["endStation"] = o.EndStation.Get()
	if o.Vehicles != nil {
		toSerialize["vehicles"] = o.Vehicles
	}
	toSerialize["controlType"] = o.ControlType
	toSerialize["controlledBySteamID"] = o.ControlledBySteamID.Get()
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	toSerialize["velocity"] = o.Velocity
	toSerialize["nextSignal"] = o.NextSignal.Get()
	toSerialize["distanceToNextSignal"] = o.DistanceToNextSignal
	toSerialize["signalInFrontSpeed"] = o.SignalInFrontSpeed
	toSerialize["apiTimetableIndex"] = o.ApiTimetableIndex
	toSerialize["lastRealTimetableIndex"] = o.LastRealTimetableIndex
	toSerialize["currentPositionType"] = o.CurrentPositionType
	if o.Timetable != nil {
		toSerialize["timetable"] = o.Timetable
	}
	toSerialize["currentTimetableIndex"] = o.CurrentTimetableIndex
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"trainName",
		"trainNumber",
		"trainNumberInternational",
		"trainLength",
		"trainWeight",
		"continuesAs",
		"serverData",
		"id",
		"startStation",
		"endStation",
		"vehicles",
		"controlType",
		"controlledBySteamID",
		"latitude",
		"longitude",
		"velocity",
		"nextSignal",
		"distanceToNextSignal",
		"signalInFrontSpeed",
		"apiTimetableIndex",
		"lastRealTimetableIndex",
		"currentPositionType",
		"timetable",
		"currentTimetableIndex",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFPHSpedVAPIObjectsSimRailSimRailActiveTrainData := _FPHSpedVAPIObjectsSimRailSimRailActiveTrainData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsSimRailSimRailActiveTrainData)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsSimRailSimRailActiveTrainData(varFPHSpedVAPIObjectsSimRailSimRailActiveTrainData)

	return err
}

type NullableFPHSpedVAPIObjectsSimRailSimRailActiveTrainData struct {
	value *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailActiveTrainData) Get() *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailActiveTrainData) Set(val *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailActiveTrainData) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailActiveTrainData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSimRailSimRailActiveTrainData(val *FPHSpedVAPIObjectsSimRailSimRailActiveTrainData) *NullableFPHSpedVAPIObjectsSimRailSimRailActiveTrainData {
	return &NullableFPHSpedVAPIObjectsSimRailSimRailActiveTrainData{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailActiveTrainData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailActiveTrainData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

