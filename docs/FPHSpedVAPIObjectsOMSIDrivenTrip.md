# FPHSpedVAPIObjectsOMSIDrivenTrip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**VisibleID** | **NullableString** |  | [readonly] 
**StartDate** | **time.Time** |  | [readonly] 
**EndDate** | **NullableTime** |  | [readonly] 
**NeededTime** | **NullableString** |  | [readonly] 
**BaseDate** | **time.Time** |  | [readonly] 
**TourTrip** | [**NullableFPHSpedVAPIObjectsOMSITourTrip**](FPHSpedVAPIObjectsOMSITourTrip.md) |  | [readonly] 
**Map** | [**NullableFPHSpedVAPIObjectsOMSIMap**](FPHSpedVAPIObjectsOMSIMap.md) |  | [readonly] 
**StartKM** | **float64** |  | [readonly] 
**EndKM** | **NullableFloat64** |  | [readonly] 
**VisDistance** | **NullableString** |  | [readonly] 
**Value** | **NullableInt32** |  | [readonly] 
**RealtimeBoni** | **NullableFloat64** |  | [readonly] 
**FollowupBoni** | **NullableFloat64** |  | [readonly] 
**Tax** | **NullableFloat64** |  | [readonly] 
**ResultValue** | **NullableFloat64** |  | [readonly] 
**Currency** | **NullableString** |  | [readonly] 
**State** | [**FPHSpedVAPIEnumsOMSITaskState**](FPHSpedVAPIEnumsOMSITaskState.md) |   0 &#x3D; InDrive  1 &#x3D; Finished  2 &#x3D; Cancelled  -1 &#x3D; NotAvaliable | [readonly] 
**User** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**Spedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**RegisteredBusStops** | [**[]FPHSpedVAPIObjectsOMSIDrivenTripBusStop**](FPHSpedVAPIObjectsOMSIDrivenTripBusStop.md) |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsOMSIDrivenTrip

`func NewFPHSpedVAPIObjectsOMSIDrivenTrip(id int32, visibleID NullableString, startDate time.Time, endDate NullableTime, neededTime NullableString, baseDate time.Time, tourTrip NullableFPHSpedVAPIObjectsOMSITourTrip, map_ NullableFPHSpedVAPIObjectsOMSIMap, startKM float64, endKM NullableFloat64, visDistance NullableString, value NullableInt32, realtimeBoni NullableFloat64, followupBoni NullableFloat64, tax NullableFloat64, resultValue NullableFloat64, currency NullableString, state FPHSpedVAPIEnumsOMSITaskState, user NullableFPHSpedVAPIObjectsUsersUser, spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, registeredBusStops []FPHSpedVAPIObjectsOMSIDrivenTripBusStop, ) *FPHSpedVAPIObjectsOMSIDrivenTrip`

NewFPHSpedVAPIObjectsOMSIDrivenTrip instantiates a new FPHSpedVAPIObjectsOMSIDrivenTrip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsOMSIDrivenTripWithDefaults

`func NewFPHSpedVAPIObjectsOMSIDrivenTripWithDefaults() *FPHSpedVAPIObjectsOMSIDrivenTrip`

NewFPHSpedVAPIObjectsOMSIDrivenTripWithDefaults instantiates a new FPHSpedVAPIObjectsOMSIDrivenTrip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetId(v int32)`

SetId sets Id field to given value.


### GetVisibleID

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetVisibleID() string`

GetVisibleID returns the VisibleID field if non-nil, zero value otherwise.

### GetVisibleIDOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetVisibleIDOk() (*string, bool)`

GetVisibleIDOk returns a tuple with the VisibleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleID

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetVisibleID(v string)`

SetVisibleID sets VisibleID field to given value.


### SetVisibleIDNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetVisibleIDNil(b bool)`

 SetVisibleIDNil sets the value for VisibleID to be an explicit nil

### UnsetVisibleID
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetVisibleID()`

UnsetVisibleID ensures that no value is present for VisibleID, not even an explicit nil
### GetStartDate

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetNeededTime

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetNeededTime() string`

GetNeededTime returns the NeededTime field if non-nil, zero value otherwise.

### GetNeededTimeOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetNeededTimeOk() (*string, bool)`

GetNeededTimeOk returns a tuple with the NeededTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeededTime

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetNeededTime(v string)`

SetNeededTime sets NeededTime field to given value.


### SetNeededTimeNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetNeededTimeNil(b bool)`

 SetNeededTimeNil sets the value for NeededTime to be an explicit nil

### UnsetNeededTime
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetNeededTime()`

UnsetNeededTime ensures that no value is present for NeededTime, not even an explicit nil
### GetBaseDate

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetBaseDate() time.Time`

GetBaseDate returns the BaseDate field if non-nil, zero value otherwise.

### GetBaseDateOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetBaseDateOk() (*time.Time, bool)`

GetBaseDateOk returns a tuple with the BaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDate

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetBaseDate(v time.Time)`

SetBaseDate sets BaseDate field to given value.


### GetTourTrip

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetTourTrip() FPHSpedVAPIObjectsOMSITourTrip`

GetTourTrip returns the TourTrip field if non-nil, zero value otherwise.

### GetTourTripOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetTourTripOk() (*FPHSpedVAPIObjectsOMSITourTrip, bool)`

GetTourTripOk returns a tuple with the TourTrip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTourTrip

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetTourTrip(v FPHSpedVAPIObjectsOMSITourTrip)`

SetTourTrip sets TourTrip field to given value.


### SetTourTripNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetTourTripNil(b bool)`

 SetTourTripNil sets the value for TourTrip to be an explicit nil

### UnsetTourTrip
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetTourTrip()`

UnsetTourTrip ensures that no value is present for TourTrip, not even an explicit nil
### GetMap

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetMap() FPHSpedVAPIObjectsOMSIMap`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetMapOk() (*FPHSpedVAPIObjectsOMSIMap, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetMap(v FPHSpedVAPIObjectsOMSIMap)`

SetMap sets Map field to given value.


### SetMapNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetMapNil(b bool)`

 SetMapNil sets the value for Map to be an explicit nil

### UnsetMap
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetMap()`

UnsetMap ensures that no value is present for Map, not even an explicit nil
### GetStartKM

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetStartKM() float64`

GetStartKM returns the StartKM field if non-nil, zero value otherwise.

### GetStartKMOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetStartKMOk() (*float64, bool)`

GetStartKMOk returns a tuple with the StartKM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartKM

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetStartKM(v float64)`

SetStartKM sets StartKM field to given value.


### GetEndKM

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetEndKM() float64`

GetEndKM returns the EndKM field if non-nil, zero value otherwise.

### GetEndKMOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetEndKMOk() (*float64, bool)`

GetEndKMOk returns a tuple with the EndKM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndKM

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetEndKM(v float64)`

SetEndKM sets EndKM field to given value.


### SetEndKMNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetEndKMNil(b bool)`

 SetEndKMNil sets the value for EndKM to be an explicit nil

### UnsetEndKM
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetEndKM()`

UnsetEndKM ensures that no value is present for EndKM, not even an explicit nil
### GetVisDistance

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetVisDistance() string`

GetVisDistance returns the VisDistance field if non-nil, zero value otherwise.

### GetVisDistanceOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetVisDistanceOk() (*string, bool)`

GetVisDistanceOk returns a tuple with the VisDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisDistance

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetVisDistance(v string)`

SetVisDistance sets VisDistance field to given value.


### SetVisDistanceNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetVisDistanceNil(b bool)`

 SetVisDistanceNil sets the value for VisDistance to be an explicit nil

### UnsetVisDistance
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetVisDistance()`

UnsetVisDistance ensures that no value is present for VisDistance, not even an explicit nil
### GetValue

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetValue(v int32)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetRealtimeBoni

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetRealtimeBoni() float64`

GetRealtimeBoni returns the RealtimeBoni field if non-nil, zero value otherwise.

### GetRealtimeBoniOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetRealtimeBoniOk() (*float64, bool)`

GetRealtimeBoniOk returns a tuple with the RealtimeBoni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtimeBoni

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetRealtimeBoni(v float64)`

SetRealtimeBoni sets RealtimeBoni field to given value.


### SetRealtimeBoniNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetRealtimeBoniNil(b bool)`

 SetRealtimeBoniNil sets the value for RealtimeBoni to be an explicit nil

### UnsetRealtimeBoni
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetRealtimeBoni()`

UnsetRealtimeBoni ensures that no value is present for RealtimeBoni, not even an explicit nil
### GetFollowupBoni

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetFollowupBoni() float64`

GetFollowupBoni returns the FollowupBoni field if non-nil, zero value otherwise.

### GetFollowupBoniOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetFollowupBoniOk() (*float64, bool)`

GetFollowupBoniOk returns a tuple with the FollowupBoni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowupBoni

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetFollowupBoni(v float64)`

SetFollowupBoni sets FollowupBoni field to given value.


### SetFollowupBoniNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetFollowupBoniNil(b bool)`

 SetFollowupBoniNil sets the value for FollowupBoni to be an explicit nil

### UnsetFollowupBoni
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetFollowupBoni()`

UnsetFollowupBoni ensures that no value is present for FollowupBoni, not even an explicit nil
### GetTax

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetTax() float64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetTaxOk() (*float64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetTax(v float64)`

SetTax sets Tax field to given value.


### SetTaxNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetTaxNil(b bool)`

 SetTaxNil sets the value for Tax to be an explicit nil

### UnsetTax
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetTax()`

UnsetTax ensures that no value is present for Tax, not even an explicit nil
### GetResultValue

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetResultValue() float64`

GetResultValue returns the ResultValue field if non-nil, zero value otherwise.

### GetResultValueOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetResultValueOk() (*float64, bool)`

GetResultValueOk returns a tuple with the ResultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultValue

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetResultValue(v float64)`

SetResultValue sets ResultValue field to given value.


### SetResultValueNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetResultValueNil(b bool)`

 SetResultValueNil sets the value for ResultValue to be an explicit nil

### UnsetResultValue
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetResultValue()`

UnsetResultValue ensures that no value is present for ResultValue, not even an explicit nil
### GetCurrency

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetState

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetState() FPHSpedVAPIEnumsOMSITaskState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetStateOk() (*FPHSpedVAPIEnumsOMSITaskState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetState(v FPHSpedVAPIEnumsOMSITaskState)`

SetState sets State field to given value.


### GetUser

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetUser() FPHSpedVAPIObjectsUsersUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetUser(v FPHSpedVAPIObjectsUsersUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetSpedition

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetSpedition returns the Spedition field if non-nil, zero value otherwise.

### GetSpeditionOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetSpeditionOk returns a tuple with the Spedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpedition

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetSpedition sets Spedition field to given value.


### SetSpeditionNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetSpeditionNil(b bool)`

 SetSpeditionNil sets the value for Spedition to be an explicit nil

### UnsetSpedition
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetSpedition()`

UnsetSpedition ensures that no value is present for Spedition, not even an explicit nil
### GetRegisteredBusStops

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetRegisteredBusStops() []FPHSpedVAPIObjectsOMSIDrivenTripBusStop`

GetRegisteredBusStops returns the RegisteredBusStops field if non-nil, zero value otherwise.

### GetRegisteredBusStopsOk

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) GetRegisteredBusStopsOk() (*[]FPHSpedVAPIObjectsOMSIDrivenTripBusStop, bool)`

GetRegisteredBusStopsOk returns a tuple with the RegisteredBusStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredBusStops

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetRegisteredBusStops(v []FPHSpedVAPIObjectsOMSIDrivenTripBusStop)`

SetRegisteredBusStops sets RegisteredBusStops field to given value.


### SetRegisteredBusStopsNil

`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) SetRegisteredBusStopsNil(b bool)`

 SetRegisteredBusStopsNil sets the value for RegisteredBusStops to be an explicit nil

### UnsetRegisteredBusStops
`func (o *FPHSpedVAPIObjectsOMSIDrivenTrip) UnsetRegisteredBusStops()`

UnsetRegisteredBusStops ensures that no value is present for RegisteredBusStops, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


