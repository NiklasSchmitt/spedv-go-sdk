# FPHSpedVAPIObjectsOMSITrip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**LineStr** | **NullableString** |  | [readonly] 
**Dest** | [**NullableFPHSpedVAPIObjectsOMSIDestination**](FPHSpedVAPIObjectsOMSIDestination.md) |  | [readonly] 
**Line** | **int32** |  | [readonly] 
**Circ** | **int32** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsOMSITrip

`func NewFPHSpedVAPIObjectsOMSITrip(id int32, lineStr NullableString, dest NullableFPHSpedVAPIObjectsOMSIDestination, line int32, circ int32, ) *FPHSpedVAPIObjectsOMSITrip`

NewFPHSpedVAPIObjectsOMSITrip instantiates a new FPHSpedVAPIObjectsOMSITrip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsOMSITripWithDefaults

`func NewFPHSpedVAPIObjectsOMSITripWithDefaults() *FPHSpedVAPIObjectsOMSITrip`

NewFPHSpedVAPIObjectsOMSITripWithDefaults instantiates a new FPHSpedVAPIObjectsOMSITrip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsOMSITrip) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsOMSITrip) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsOMSITrip) SetId(v int32)`

SetId sets Id field to given value.


### GetLineStr

`func (o *FPHSpedVAPIObjectsOMSITrip) GetLineStr() string`

GetLineStr returns the LineStr field if non-nil, zero value otherwise.

### GetLineStrOk

`func (o *FPHSpedVAPIObjectsOMSITrip) GetLineStrOk() (*string, bool)`

GetLineStrOk returns a tuple with the LineStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStr

`func (o *FPHSpedVAPIObjectsOMSITrip) SetLineStr(v string)`

SetLineStr sets LineStr field to given value.


### SetLineStrNil

`func (o *FPHSpedVAPIObjectsOMSITrip) SetLineStrNil(b bool)`

 SetLineStrNil sets the value for LineStr to be an explicit nil

### UnsetLineStr
`func (o *FPHSpedVAPIObjectsOMSITrip) UnsetLineStr()`

UnsetLineStr ensures that no value is present for LineStr, not even an explicit nil
### GetDest

`func (o *FPHSpedVAPIObjectsOMSITrip) GetDest() FPHSpedVAPIObjectsOMSIDestination`

GetDest returns the Dest field if non-nil, zero value otherwise.

### GetDestOk

`func (o *FPHSpedVAPIObjectsOMSITrip) GetDestOk() (*FPHSpedVAPIObjectsOMSIDestination, bool)`

GetDestOk returns a tuple with the Dest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDest

`func (o *FPHSpedVAPIObjectsOMSITrip) SetDest(v FPHSpedVAPIObjectsOMSIDestination)`

SetDest sets Dest field to given value.


### SetDestNil

`func (o *FPHSpedVAPIObjectsOMSITrip) SetDestNil(b bool)`

 SetDestNil sets the value for Dest to be an explicit nil

### UnsetDest
`func (o *FPHSpedVAPIObjectsOMSITrip) UnsetDest()`

UnsetDest ensures that no value is present for Dest, not even an explicit nil
### GetLine

`func (o *FPHSpedVAPIObjectsOMSITrip) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *FPHSpedVAPIObjectsOMSITrip) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *FPHSpedVAPIObjectsOMSITrip) SetLine(v int32)`

SetLine sets Line field to given value.


### GetCirc

`func (o *FPHSpedVAPIObjectsOMSITrip) GetCirc() int32`

GetCirc returns the Circ field if non-nil, zero value otherwise.

### GetCircOk

`func (o *FPHSpedVAPIObjectsOMSITrip) GetCircOk() (*int32, bool)`

GetCircOk returns a tuple with the Circ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCirc

`func (o *FPHSpedVAPIObjectsOMSITrip) SetCirc(v int32)`

SetCirc sets Circ field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


