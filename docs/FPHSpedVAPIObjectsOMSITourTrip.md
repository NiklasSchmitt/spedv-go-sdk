# FPHSpedVAPIObjectsOMSITourTrip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tour** | [**NullableFPHSpedVAPIObjectsOMSITour**](FPHSpedVAPIObjectsOMSITour.md) |  | [readonly] 
**Trip** | [**NullableFPHSpedVAPIObjectsOMSITrip**](FPHSpedVAPIObjectsOMSITrip.md) |  | [readonly] 
**BusStops** | [**[]FPHSpedVAPIObjectsOMSITripBusStop**](FPHSpedVAPIObjectsOMSITripBusStop.md) |  | [readonly] 
**Profile** | **int32** |  | [readonly] 
**StartTime** | **time.Time** |  | [readonly] 
**BaseDate** | **time.Time** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsOMSITourTrip

`func NewFPHSpedVAPIObjectsOMSITourTrip(id int32, tour NullableFPHSpedVAPIObjectsOMSITour, trip NullableFPHSpedVAPIObjectsOMSITrip, busStops []FPHSpedVAPIObjectsOMSITripBusStop, profile int32, startTime time.Time, baseDate time.Time, ) *FPHSpedVAPIObjectsOMSITourTrip`

NewFPHSpedVAPIObjectsOMSITourTrip instantiates a new FPHSpedVAPIObjectsOMSITourTrip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsOMSITourTripWithDefaults

`func NewFPHSpedVAPIObjectsOMSITourTripWithDefaults() *FPHSpedVAPIObjectsOMSITourTrip`

NewFPHSpedVAPIObjectsOMSITourTripWithDefaults instantiates a new FPHSpedVAPIObjectsOMSITourTrip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsOMSITourTrip) SetId(v int32)`

SetId sets Id field to given value.


### GetTour

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetTour() FPHSpedVAPIObjectsOMSITour`

GetTour returns the Tour field if non-nil, zero value otherwise.

### GetTourOk

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetTourOk() (*FPHSpedVAPIObjectsOMSITour, bool)`

GetTourOk returns a tuple with the Tour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTour

`func (o *FPHSpedVAPIObjectsOMSITourTrip) SetTour(v FPHSpedVAPIObjectsOMSITour)`

SetTour sets Tour field to given value.


### SetTourNil

`func (o *FPHSpedVAPIObjectsOMSITourTrip) SetTourNil(b bool)`

 SetTourNil sets the value for Tour to be an explicit nil

### UnsetTour
`func (o *FPHSpedVAPIObjectsOMSITourTrip) UnsetTour()`

UnsetTour ensures that no value is present for Tour, not even an explicit nil
### GetTrip

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetTrip() FPHSpedVAPIObjectsOMSITrip`

GetTrip returns the Trip field if non-nil, zero value otherwise.

### GetTripOk

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetTripOk() (*FPHSpedVAPIObjectsOMSITrip, bool)`

GetTripOk returns a tuple with the Trip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrip

`func (o *FPHSpedVAPIObjectsOMSITourTrip) SetTrip(v FPHSpedVAPIObjectsOMSITrip)`

SetTrip sets Trip field to given value.


### SetTripNil

`func (o *FPHSpedVAPIObjectsOMSITourTrip) SetTripNil(b bool)`

 SetTripNil sets the value for Trip to be an explicit nil

### UnsetTrip
`func (o *FPHSpedVAPIObjectsOMSITourTrip) UnsetTrip()`

UnsetTrip ensures that no value is present for Trip, not even an explicit nil
### GetBusStops

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetBusStops() []FPHSpedVAPIObjectsOMSITripBusStop`

GetBusStops returns the BusStops field if non-nil, zero value otherwise.

### GetBusStopsOk

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetBusStopsOk() (*[]FPHSpedVAPIObjectsOMSITripBusStop, bool)`

GetBusStopsOk returns a tuple with the BusStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusStops

`func (o *FPHSpedVAPIObjectsOMSITourTrip) SetBusStops(v []FPHSpedVAPIObjectsOMSITripBusStop)`

SetBusStops sets BusStops field to given value.


### SetBusStopsNil

`func (o *FPHSpedVAPIObjectsOMSITourTrip) SetBusStopsNil(b bool)`

 SetBusStopsNil sets the value for BusStops to be an explicit nil

### UnsetBusStops
`func (o *FPHSpedVAPIObjectsOMSITourTrip) UnsetBusStops()`

UnsetBusStops ensures that no value is present for BusStops, not even an explicit nil
### GetProfile

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetProfile() int32`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetProfileOk() (*int32, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FPHSpedVAPIObjectsOMSITourTrip) SetProfile(v int32)`

SetProfile sets Profile field to given value.


### GetStartTime

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *FPHSpedVAPIObjectsOMSITourTrip) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetBaseDate

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetBaseDate() time.Time`

GetBaseDate returns the BaseDate field if non-nil, zero value otherwise.

### GetBaseDateOk

`func (o *FPHSpedVAPIObjectsOMSITourTrip) GetBaseDateOk() (*time.Time, bool)`

GetBaseDateOk returns a tuple with the BaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDate

`func (o *FPHSpedVAPIObjectsOMSITourTrip) SetBaseDate(v time.Time)`

SetBaseDate sets BaseDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


