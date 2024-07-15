# FPHSpedVAPIObjectsMapsRealRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **NullableString** |  | [optional] 
**Destination** | Pointer to **NullableString** |  | [optional] 
**CachedDate** | Pointer to **time.Time** |  | [optional] 
**DistanceKM** | Pointer to **int32** |  | [optional] 
**Locations** | Pointer to **[][]float64** |  | [optional] 
**ViaFerryID** | Pointer to **int32** |  | [optional] 

## Methods

### NewFPHSpedVAPIObjectsMapsRealRoute

`func NewFPHSpedVAPIObjectsMapsRealRoute() *FPHSpedVAPIObjectsMapsRealRoute`

NewFPHSpedVAPIObjectsMapsRealRoute instantiates a new FPHSpedVAPIObjectsMapsRealRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsMapsRealRouteWithDefaults

`func NewFPHSpedVAPIObjectsMapsRealRouteWithDefaults() *FPHSpedVAPIObjectsMapsRealRoute`

NewFPHSpedVAPIObjectsMapsRealRouteWithDefaults instantiates a new FPHSpedVAPIObjectsMapsRealRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *FPHSpedVAPIObjectsMapsRealRoute) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *FPHSpedVAPIObjectsMapsRealRoute) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *FPHSpedVAPIObjectsMapsRealRoute) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *FPHSpedVAPIObjectsMapsRealRoute) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *FPHSpedVAPIObjectsMapsRealRoute) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *FPHSpedVAPIObjectsMapsRealRoute) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetDestination

`func (o *FPHSpedVAPIObjectsMapsRealRoute) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *FPHSpedVAPIObjectsMapsRealRoute) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *FPHSpedVAPIObjectsMapsRealRoute) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *FPHSpedVAPIObjectsMapsRealRoute) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *FPHSpedVAPIObjectsMapsRealRoute) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *FPHSpedVAPIObjectsMapsRealRoute) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetCachedDate

`func (o *FPHSpedVAPIObjectsMapsRealRoute) GetCachedDate() time.Time`

GetCachedDate returns the CachedDate field if non-nil, zero value otherwise.

### GetCachedDateOk

`func (o *FPHSpedVAPIObjectsMapsRealRoute) GetCachedDateOk() (*time.Time, bool)`

GetCachedDateOk returns a tuple with the CachedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedDate

`func (o *FPHSpedVAPIObjectsMapsRealRoute) SetCachedDate(v time.Time)`

SetCachedDate sets CachedDate field to given value.

### HasCachedDate

`func (o *FPHSpedVAPIObjectsMapsRealRoute) HasCachedDate() bool`

HasCachedDate returns a boolean if a field has been set.

### GetDistanceKM

`func (o *FPHSpedVAPIObjectsMapsRealRoute) GetDistanceKM() int32`

GetDistanceKM returns the DistanceKM field if non-nil, zero value otherwise.

### GetDistanceKMOk

`func (o *FPHSpedVAPIObjectsMapsRealRoute) GetDistanceKMOk() (*int32, bool)`

GetDistanceKMOk returns a tuple with the DistanceKM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceKM

`func (o *FPHSpedVAPIObjectsMapsRealRoute) SetDistanceKM(v int32)`

SetDistanceKM sets DistanceKM field to given value.

### HasDistanceKM

`func (o *FPHSpedVAPIObjectsMapsRealRoute) HasDistanceKM() bool`

HasDistanceKM returns a boolean if a field has been set.

### GetLocations

`func (o *FPHSpedVAPIObjectsMapsRealRoute) GetLocations() [][]float64`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *FPHSpedVAPIObjectsMapsRealRoute) GetLocationsOk() (*[][]float64, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *FPHSpedVAPIObjectsMapsRealRoute) SetLocations(v [][]float64)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *FPHSpedVAPIObjectsMapsRealRoute) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocationsNil

`func (o *FPHSpedVAPIObjectsMapsRealRoute) SetLocationsNil(b bool)`

 SetLocationsNil sets the value for Locations to be an explicit nil

### UnsetLocations
`func (o *FPHSpedVAPIObjectsMapsRealRoute) UnsetLocations()`

UnsetLocations ensures that no value is present for Locations, not even an explicit nil
### GetViaFerryID

`func (o *FPHSpedVAPIObjectsMapsRealRoute) GetViaFerryID() int32`

GetViaFerryID returns the ViaFerryID field if non-nil, zero value otherwise.

### GetViaFerryIDOk

`func (o *FPHSpedVAPIObjectsMapsRealRoute) GetViaFerryIDOk() (*int32, bool)`

GetViaFerryIDOk returns a tuple with the ViaFerryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaFerryID

`func (o *FPHSpedVAPIObjectsMapsRealRoute) SetViaFerryID(v int32)`

SetViaFerryID sets ViaFerryID field to given value.

### HasViaFerryID

`func (o *FPHSpedVAPIObjectsMapsRealRoute) HasViaFerryID() bool`

HasViaFerryID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


