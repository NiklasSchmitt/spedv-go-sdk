# FPHSpedVAPIObjectsOMSIStationLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Map** | [**NullableFPHSpedVAPIObjectsOMSIMap**](FPHSpedVAPIObjectsOMSIMap.md) |  | [readonly] 
**StartCubeID** | **int32** |  | [readonly] 
**DestCubeID** | **int32** |  | [readonly] 
**StartBusStop** | [**NullableFPHSpedVAPIObjectsOMSIBusStop**](FPHSpedVAPIObjectsOMSIBusStop.md) |  | [readonly] 
**DestBusStop** | [**NullableFPHSpedVAPIObjectsOMSIBusStop**](FPHSpedVAPIObjectsOMSIBusStop.md) |  | [readonly] 
**Length** | **int32** |  | [readonly] 
**Source** | **NullableString** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsOMSIStationLink

`func NewFPHSpedVAPIObjectsOMSIStationLink(id int32, map_ NullableFPHSpedVAPIObjectsOMSIMap, startCubeID int32, destCubeID int32, startBusStop NullableFPHSpedVAPIObjectsOMSIBusStop, destBusStop NullableFPHSpedVAPIObjectsOMSIBusStop, length int32, source NullableString, ) *FPHSpedVAPIObjectsOMSIStationLink`

NewFPHSpedVAPIObjectsOMSIStationLink instantiates a new FPHSpedVAPIObjectsOMSIStationLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsOMSIStationLinkWithDefaults

`func NewFPHSpedVAPIObjectsOMSIStationLinkWithDefaults() *FPHSpedVAPIObjectsOMSIStationLink`

NewFPHSpedVAPIObjectsOMSIStationLinkWithDefaults instantiates a new FPHSpedVAPIObjectsOMSIStationLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsOMSIStationLink) SetId(v int32)`

SetId sets Id field to given value.


### GetMap

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetMap() FPHSpedVAPIObjectsOMSIMap`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetMapOk() (*FPHSpedVAPIObjectsOMSIMap, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *FPHSpedVAPIObjectsOMSIStationLink) SetMap(v FPHSpedVAPIObjectsOMSIMap)`

SetMap sets Map field to given value.


### SetMapNil

`func (o *FPHSpedVAPIObjectsOMSIStationLink) SetMapNil(b bool)`

 SetMapNil sets the value for Map to be an explicit nil

### UnsetMap
`func (o *FPHSpedVAPIObjectsOMSIStationLink) UnsetMap()`

UnsetMap ensures that no value is present for Map, not even an explicit nil
### GetStartCubeID

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetStartCubeID() int32`

GetStartCubeID returns the StartCubeID field if non-nil, zero value otherwise.

### GetStartCubeIDOk

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetStartCubeIDOk() (*int32, bool)`

GetStartCubeIDOk returns a tuple with the StartCubeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCubeID

`func (o *FPHSpedVAPIObjectsOMSIStationLink) SetStartCubeID(v int32)`

SetStartCubeID sets StartCubeID field to given value.


### GetDestCubeID

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetDestCubeID() int32`

GetDestCubeID returns the DestCubeID field if non-nil, zero value otherwise.

### GetDestCubeIDOk

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetDestCubeIDOk() (*int32, bool)`

GetDestCubeIDOk returns a tuple with the DestCubeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCubeID

`func (o *FPHSpedVAPIObjectsOMSIStationLink) SetDestCubeID(v int32)`

SetDestCubeID sets DestCubeID field to given value.


### GetStartBusStop

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetStartBusStop() FPHSpedVAPIObjectsOMSIBusStop`

GetStartBusStop returns the StartBusStop field if non-nil, zero value otherwise.

### GetStartBusStopOk

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetStartBusStopOk() (*FPHSpedVAPIObjectsOMSIBusStop, bool)`

GetStartBusStopOk returns a tuple with the StartBusStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartBusStop

`func (o *FPHSpedVAPIObjectsOMSIStationLink) SetStartBusStop(v FPHSpedVAPIObjectsOMSIBusStop)`

SetStartBusStop sets StartBusStop field to given value.


### SetStartBusStopNil

`func (o *FPHSpedVAPIObjectsOMSIStationLink) SetStartBusStopNil(b bool)`

 SetStartBusStopNil sets the value for StartBusStop to be an explicit nil

### UnsetStartBusStop
`func (o *FPHSpedVAPIObjectsOMSIStationLink) UnsetStartBusStop()`

UnsetStartBusStop ensures that no value is present for StartBusStop, not even an explicit nil
### GetDestBusStop

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetDestBusStop() FPHSpedVAPIObjectsOMSIBusStop`

GetDestBusStop returns the DestBusStop field if non-nil, zero value otherwise.

### GetDestBusStopOk

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetDestBusStopOk() (*FPHSpedVAPIObjectsOMSIBusStop, bool)`

GetDestBusStopOk returns a tuple with the DestBusStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestBusStop

`func (o *FPHSpedVAPIObjectsOMSIStationLink) SetDestBusStop(v FPHSpedVAPIObjectsOMSIBusStop)`

SetDestBusStop sets DestBusStop field to given value.


### SetDestBusStopNil

`func (o *FPHSpedVAPIObjectsOMSIStationLink) SetDestBusStopNil(b bool)`

 SetDestBusStopNil sets the value for DestBusStop to be an explicit nil

### UnsetDestBusStop
`func (o *FPHSpedVAPIObjectsOMSIStationLink) UnsetDestBusStop()`

UnsetDestBusStop ensures that no value is present for DestBusStop, not even an explicit nil
### GetLength

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *FPHSpedVAPIObjectsOMSIStationLink) SetLength(v int32)`

SetLength sets Length field to given value.


### GetSource

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FPHSpedVAPIObjectsOMSIStationLink) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FPHSpedVAPIObjectsOMSIStationLink) SetSource(v string)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *FPHSpedVAPIObjectsOMSIStationLink) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *FPHSpedVAPIObjectsOMSIStationLink) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


