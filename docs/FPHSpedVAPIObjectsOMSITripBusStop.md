# FPHSpedVAPIObjectsOMSITripBusStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**BusStop** | [**NullableFPHSpedVAPIObjectsOMSIBusStop**](FPHSpedVAPIObjectsOMSIBusStop.md) |  | [readonly] 
**Time** | **string** |  | [readonly] 
**TimeSinceLast** | **string** |  | [readonly] 
**Departure** | **time.Time** |  | [readonly] 
**IsGoThrough** | **bool** |  | [readonly] 
**DistanceSinceLast** | **int32** |  | [readonly] 
**NeededVelocity** | **NullableString** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsOMSITripBusStop

`func NewFPHSpedVAPIObjectsOMSITripBusStop(id int32, busStop NullableFPHSpedVAPIObjectsOMSIBusStop, time string, timeSinceLast string, departure time.Time, isGoThrough bool, distanceSinceLast int32, neededVelocity NullableString, ) *FPHSpedVAPIObjectsOMSITripBusStop`

NewFPHSpedVAPIObjectsOMSITripBusStop instantiates a new FPHSpedVAPIObjectsOMSITripBusStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsOMSITripBusStopWithDefaults

`func NewFPHSpedVAPIObjectsOMSITripBusStopWithDefaults() *FPHSpedVAPIObjectsOMSITripBusStop`

NewFPHSpedVAPIObjectsOMSITripBusStopWithDefaults instantiates a new FPHSpedVAPIObjectsOMSITripBusStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) SetId(v int32)`

SetId sets Id field to given value.


### GetBusStop

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetBusStop() FPHSpedVAPIObjectsOMSIBusStop`

GetBusStop returns the BusStop field if non-nil, zero value otherwise.

### GetBusStopOk

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetBusStopOk() (*FPHSpedVAPIObjectsOMSIBusStop, bool)`

GetBusStopOk returns a tuple with the BusStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusStop

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) SetBusStop(v FPHSpedVAPIObjectsOMSIBusStop)`

SetBusStop sets BusStop field to given value.


### SetBusStopNil

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) SetBusStopNil(b bool)`

 SetBusStopNil sets the value for BusStop to be an explicit nil

### UnsetBusStop
`func (o *FPHSpedVAPIObjectsOMSITripBusStop) UnsetBusStop()`

UnsetBusStop ensures that no value is present for BusStop, not even an explicit nil
### GetTime

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) SetTime(v string)`

SetTime sets Time field to given value.


### GetTimeSinceLast

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetTimeSinceLast() string`

GetTimeSinceLast returns the TimeSinceLast field if non-nil, zero value otherwise.

### GetTimeSinceLastOk

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetTimeSinceLastOk() (*string, bool)`

GetTimeSinceLastOk returns a tuple with the TimeSinceLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSinceLast

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) SetTimeSinceLast(v string)`

SetTimeSinceLast sets TimeSinceLast field to given value.


### GetDeparture

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetDeparture() time.Time`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetDepartureOk() (*time.Time, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) SetDeparture(v time.Time)`

SetDeparture sets Departure field to given value.


### GetIsGoThrough

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetIsGoThrough() bool`

GetIsGoThrough returns the IsGoThrough field if non-nil, zero value otherwise.

### GetIsGoThroughOk

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetIsGoThroughOk() (*bool, bool)`

GetIsGoThroughOk returns a tuple with the IsGoThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGoThrough

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) SetIsGoThrough(v bool)`

SetIsGoThrough sets IsGoThrough field to given value.


### GetDistanceSinceLast

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetDistanceSinceLast() int32`

GetDistanceSinceLast returns the DistanceSinceLast field if non-nil, zero value otherwise.

### GetDistanceSinceLastOk

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetDistanceSinceLastOk() (*int32, bool)`

GetDistanceSinceLastOk returns a tuple with the DistanceSinceLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceSinceLast

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) SetDistanceSinceLast(v int32)`

SetDistanceSinceLast sets DistanceSinceLast field to given value.


### GetNeededVelocity

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetNeededVelocity() string`

GetNeededVelocity returns the NeededVelocity field if non-nil, zero value otherwise.

### GetNeededVelocityOk

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) GetNeededVelocityOk() (*string, bool)`

GetNeededVelocityOk returns a tuple with the NeededVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeededVelocity

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) SetNeededVelocity(v string)`

SetNeededVelocity sets NeededVelocity field to given value.


### SetNeededVelocityNil

`func (o *FPHSpedVAPIObjectsOMSITripBusStop) SetNeededVelocityNil(b bool)`

 SetNeededVelocityNil sets the value for NeededVelocity to be an explicit nil

### UnsetNeededVelocity
`func (o *FPHSpedVAPIObjectsOMSITripBusStop) UnsetNeededVelocity()`

UnsetNeededVelocity ensures that no value is present for NeededVelocity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


