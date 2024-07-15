# FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**VisibleID** | **NullableString** |  | [readonly] 
**State** | [**FPHSpedVAPIEnumsETSTaskState**](FPHSpedVAPIEnumsETSTaskState.md) |   0 &#x3D; InDrive  1 &#x3D; Done  2 &#x3D; Settled  3 &#x3D; Fail  4 &#x3D; AdminCheck  5 &#x3D; Paused  6 &#x3D; Cancelled  7 &#x3D; Invalid  -1 &#x3D; NotAvaliable | [readonly] 
**Start** | **NullableString** |  | [readonly] 
**Dest** | **NullableString** |  | [readonly] 
**Freight** | **NullableString** |  | [readonly] 
**HasScreenshot** | **int32** |  | [readonly] 
**IsDeductable** | **bool** |  | [readonly] 
**DdCleaned** | **bool** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo

`func NewFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo(id int32, visibleID NullableString, state FPHSpedVAPIEnumsETSTaskState, start NullableString, dest NullableString, freight NullableString, hasScreenshot int32, isDeductable bool, ddCleaned bool, ) *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo`

NewFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo instantiates a new FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfoWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfoWithDefaults() *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo`

NewFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfoWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetId(v int32)`

SetId sets Id field to given value.


### GetVisibleID

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetVisibleID() string`

GetVisibleID returns the VisibleID field if non-nil, zero value otherwise.

### GetVisibleIDOk

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetVisibleIDOk() (*string, bool)`

GetVisibleIDOk returns a tuple with the VisibleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleID

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetVisibleID(v string)`

SetVisibleID sets VisibleID field to given value.


### SetVisibleIDNil

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetVisibleIDNil(b bool)`

 SetVisibleIDNil sets the value for VisibleID to be an explicit nil

### UnsetVisibleID
`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) UnsetVisibleID()`

UnsetVisibleID ensures that no value is present for VisibleID, not even an explicit nil
### GetState

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetState() FPHSpedVAPIEnumsETSTaskState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetStateOk() (*FPHSpedVAPIEnumsETSTaskState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetState(v FPHSpedVAPIEnumsETSTaskState)`

SetState sets State field to given value.


### GetStart

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetStart(v string)`

SetStart sets Start field to given value.


### SetStartNil

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetDest

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetDest() string`

GetDest returns the Dest field if non-nil, zero value otherwise.

### GetDestOk

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetDestOk() (*string, bool)`

GetDestOk returns a tuple with the Dest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDest

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetDest(v string)`

SetDest sets Dest field to given value.


### SetDestNil

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetDestNil(b bool)`

 SetDestNil sets the value for Dest to be an explicit nil

### UnsetDest
`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) UnsetDest()`

UnsetDest ensures that no value is present for Dest, not even an explicit nil
### GetFreight

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetFreight() string`

GetFreight returns the Freight field if non-nil, zero value otherwise.

### GetFreightOk

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetFreightOk() (*string, bool)`

GetFreightOk returns a tuple with the Freight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreight

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetFreight(v string)`

SetFreight sets Freight field to given value.


### SetFreightNil

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetFreightNil(b bool)`

 SetFreightNil sets the value for Freight to be an explicit nil

### UnsetFreight
`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) UnsetFreight()`

UnsetFreight ensures that no value is present for Freight, not even an explicit nil
### GetHasScreenshot

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetHasScreenshot() int32`

GetHasScreenshot returns the HasScreenshot field if non-nil, zero value otherwise.

### GetHasScreenshotOk

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetHasScreenshotOk() (*int32, bool)`

GetHasScreenshotOk returns a tuple with the HasScreenshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasScreenshot

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetHasScreenshot(v int32)`

SetHasScreenshot sets HasScreenshot field to given value.


### GetIsDeductable

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetIsDeductable() bool`

GetIsDeductable returns the IsDeductable field if non-nil, zero value otherwise.

### GetIsDeductableOk

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetIsDeductableOk() (*bool, bool)`

GetIsDeductableOk returns a tuple with the IsDeductable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeductable

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetIsDeductable(v bool)`

SetIsDeductable sets IsDeductable field to given value.


### GetDdCleaned

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetDdCleaned() bool`

GetDdCleaned returns the DdCleaned field if non-nil, zero value otherwise.

### GetDdCleanedOk

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetDdCleanedOk() (*bool, bool)`

GetDdCleanedOk returns a tuple with the DdCleaned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdCleaned

`func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetDdCleaned(v bool)`

SetDdCleaned sets DdCleaned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


