# FPHSpedVAPIObjectsKontorJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Source** | [**NullableFPHSpedVAPIObjectsMapsCompanyCity**](FPHSpedVAPIObjectsMapsCompanyCity.md) |  | [readonly] 
**Destination** | [**NullableFPHSpedVAPIObjectsMapsCompanyCity**](FPHSpedVAPIObjectsMapsCompanyCity.md) |  | [readonly] 
**Weight** | **int32** |  | [readonly] 
**VisWeight** | **NullableString** |  | [readonly] 
**Spedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**Freight** | [**NullableFPHSpedVAPIObjectsKontorFreight**](FPHSpedVAPIObjectsKontorFreight.md) |  | [readonly] 
**Value** | **int32** |  | [readonly] 
**AcceptedBy** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**Comment** | **NullableString** |  | [readonly] 
**VisValue** | **NullableString** |  | [readonly] 
**JobParts** | [**[]FPHSpedVAPIObjectsKontorJobPartLoaded**](FPHSpedVAPIObjectsKontorJobPartLoaded.md) |  | [readonly] 
**State** | [**FPHSpedVAPIEnumsJobOfferState**](FPHSpedVAPIEnumsJobOfferState.md) |   0 &#x3D; WaitingSpedition  1 &#x3D; WaitingDrive  2 &#x3D; InDrive  3 &#x3D; Finished  -1 &#x3D; NotAvaliable | [readonly] 
**ExpirationDate** | **time.Time** |  | [readonly] 
**Distance** | **int32** |  | [readonly] 
**VisDistance** | **NullableString** |  | [readonly] 
**IsRealEco** | **bool** |  | [readonly] 
**CompleteScheduled** | **bool** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsKontorJob

`func NewFPHSpedVAPIObjectsKontorJob(id int32, source NullableFPHSpedVAPIObjectsMapsCompanyCity, destination NullableFPHSpedVAPIObjectsMapsCompanyCity, weight int32, visWeight NullableString, spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, freight NullableFPHSpedVAPIObjectsKontorFreight, value int32, acceptedBy NullableFPHSpedVAPIObjectsUsersUser, comment NullableString, visValue NullableString, jobParts []FPHSpedVAPIObjectsKontorJobPartLoaded, state FPHSpedVAPIEnumsJobOfferState, expirationDate time.Time, distance int32, visDistance NullableString, isRealEco bool, completeScheduled bool, ) *FPHSpedVAPIObjectsKontorJob`

NewFPHSpedVAPIObjectsKontorJob instantiates a new FPHSpedVAPIObjectsKontorJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsKontorJobWithDefaults

`func NewFPHSpedVAPIObjectsKontorJobWithDefaults() *FPHSpedVAPIObjectsKontorJob`

NewFPHSpedVAPIObjectsKontorJobWithDefaults instantiates a new FPHSpedVAPIObjectsKontorJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsKontorJob) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsKontorJob) SetId(v int32)`

SetId sets Id field to given value.


### GetSource

`func (o *FPHSpedVAPIObjectsKontorJob) GetSource() FPHSpedVAPIObjectsMapsCompanyCity`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetSourceOk() (*FPHSpedVAPIObjectsMapsCompanyCity, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FPHSpedVAPIObjectsKontorJob) SetSource(v FPHSpedVAPIObjectsMapsCompanyCity)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *FPHSpedVAPIObjectsKontorJob) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *FPHSpedVAPIObjectsKontorJob) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetDestination

`func (o *FPHSpedVAPIObjectsKontorJob) GetDestination() FPHSpedVAPIObjectsMapsCompanyCity`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetDestinationOk() (*FPHSpedVAPIObjectsMapsCompanyCity, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *FPHSpedVAPIObjectsKontorJob) SetDestination(v FPHSpedVAPIObjectsMapsCompanyCity)`

SetDestination sets Destination field to given value.


### SetDestinationNil

`func (o *FPHSpedVAPIObjectsKontorJob) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *FPHSpedVAPIObjectsKontorJob) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetWeight

`func (o *FPHSpedVAPIObjectsKontorJob) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *FPHSpedVAPIObjectsKontorJob) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetVisWeight

`func (o *FPHSpedVAPIObjectsKontorJob) GetVisWeight() string`

GetVisWeight returns the VisWeight field if non-nil, zero value otherwise.

### GetVisWeightOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetVisWeightOk() (*string, bool)`

GetVisWeightOk returns a tuple with the VisWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisWeight

`func (o *FPHSpedVAPIObjectsKontorJob) SetVisWeight(v string)`

SetVisWeight sets VisWeight field to given value.


### SetVisWeightNil

`func (o *FPHSpedVAPIObjectsKontorJob) SetVisWeightNil(b bool)`

 SetVisWeightNil sets the value for VisWeight to be an explicit nil

### UnsetVisWeight
`func (o *FPHSpedVAPIObjectsKontorJob) UnsetVisWeight()`

UnsetVisWeight ensures that no value is present for VisWeight, not even an explicit nil
### GetSpedition

`func (o *FPHSpedVAPIObjectsKontorJob) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetSpedition returns the Spedition field if non-nil, zero value otherwise.

### GetSpeditionOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetSpeditionOk returns a tuple with the Spedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpedition

`func (o *FPHSpedVAPIObjectsKontorJob) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetSpedition sets Spedition field to given value.


### SetSpeditionNil

`func (o *FPHSpedVAPIObjectsKontorJob) SetSpeditionNil(b bool)`

 SetSpeditionNil sets the value for Spedition to be an explicit nil

### UnsetSpedition
`func (o *FPHSpedVAPIObjectsKontorJob) UnsetSpedition()`

UnsetSpedition ensures that no value is present for Spedition, not even an explicit nil
### GetFreight

`func (o *FPHSpedVAPIObjectsKontorJob) GetFreight() FPHSpedVAPIObjectsKontorFreight`

GetFreight returns the Freight field if non-nil, zero value otherwise.

### GetFreightOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetFreightOk() (*FPHSpedVAPIObjectsKontorFreight, bool)`

GetFreightOk returns a tuple with the Freight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreight

`func (o *FPHSpedVAPIObjectsKontorJob) SetFreight(v FPHSpedVAPIObjectsKontorFreight)`

SetFreight sets Freight field to given value.


### SetFreightNil

`func (o *FPHSpedVAPIObjectsKontorJob) SetFreightNil(b bool)`

 SetFreightNil sets the value for Freight to be an explicit nil

### UnsetFreight
`func (o *FPHSpedVAPIObjectsKontorJob) UnsetFreight()`

UnsetFreight ensures that no value is present for Freight, not even an explicit nil
### GetValue

`func (o *FPHSpedVAPIObjectsKontorJob) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FPHSpedVAPIObjectsKontorJob) SetValue(v int32)`

SetValue sets Value field to given value.


### GetAcceptedBy

`func (o *FPHSpedVAPIObjectsKontorJob) GetAcceptedBy() FPHSpedVAPIObjectsUsersUser`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetAcceptedByOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *FPHSpedVAPIObjectsKontorJob) SetAcceptedBy(v FPHSpedVAPIObjectsUsersUser)`

SetAcceptedBy sets AcceptedBy field to given value.


### SetAcceptedByNil

`func (o *FPHSpedVAPIObjectsKontorJob) SetAcceptedByNil(b bool)`

 SetAcceptedByNil sets the value for AcceptedBy to be an explicit nil

### UnsetAcceptedBy
`func (o *FPHSpedVAPIObjectsKontorJob) UnsetAcceptedBy()`

UnsetAcceptedBy ensures that no value is present for AcceptedBy, not even an explicit nil
### GetComment

`func (o *FPHSpedVAPIObjectsKontorJob) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FPHSpedVAPIObjectsKontorJob) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *FPHSpedVAPIObjectsKontorJob) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FPHSpedVAPIObjectsKontorJob) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetVisValue

`func (o *FPHSpedVAPIObjectsKontorJob) GetVisValue() string`

GetVisValue returns the VisValue field if non-nil, zero value otherwise.

### GetVisValueOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetVisValueOk() (*string, bool)`

GetVisValueOk returns a tuple with the VisValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisValue

`func (o *FPHSpedVAPIObjectsKontorJob) SetVisValue(v string)`

SetVisValue sets VisValue field to given value.


### SetVisValueNil

`func (o *FPHSpedVAPIObjectsKontorJob) SetVisValueNil(b bool)`

 SetVisValueNil sets the value for VisValue to be an explicit nil

### UnsetVisValue
`func (o *FPHSpedVAPIObjectsKontorJob) UnsetVisValue()`

UnsetVisValue ensures that no value is present for VisValue, not even an explicit nil
### GetJobParts

`func (o *FPHSpedVAPIObjectsKontorJob) GetJobParts() []FPHSpedVAPIObjectsKontorJobPartLoaded`

GetJobParts returns the JobParts field if non-nil, zero value otherwise.

### GetJobPartsOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetJobPartsOk() (*[]FPHSpedVAPIObjectsKontorJobPartLoaded, bool)`

GetJobPartsOk returns a tuple with the JobParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobParts

`func (o *FPHSpedVAPIObjectsKontorJob) SetJobParts(v []FPHSpedVAPIObjectsKontorJobPartLoaded)`

SetJobParts sets JobParts field to given value.


### SetJobPartsNil

`func (o *FPHSpedVAPIObjectsKontorJob) SetJobPartsNil(b bool)`

 SetJobPartsNil sets the value for JobParts to be an explicit nil

### UnsetJobParts
`func (o *FPHSpedVAPIObjectsKontorJob) UnsetJobParts()`

UnsetJobParts ensures that no value is present for JobParts, not even an explicit nil
### GetState

`func (o *FPHSpedVAPIObjectsKontorJob) GetState() FPHSpedVAPIEnumsJobOfferState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetStateOk() (*FPHSpedVAPIEnumsJobOfferState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FPHSpedVAPIObjectsKontorJob) SetState(v FPHSpedVAPIEnumsJobOfferState)`

SetState sets State field to given value.


### GetExpirationDate

`func (o *FPHSpedVAPIObjectsKontorJob) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FPHSpedVAPIObjectsKontorJob) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.


### GetDistance

`func (o *FPHSpedVAPIObjectsKontorJob) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *FPHSpedVAPIObjectsKontorJob) SetDistance(v int32)`

SetDistance sets Distance field to given value.


### GetVisDistance

`func (o *FPHSpedVAPIObjectsKontorJob) GetVisDistance() string`

GetVisDistance returns the VisDistance field if non-nil, zero value otherwise.

### GetVisDistanceOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetVisDistanceOk() (*string, bool)`

GetVisDistanceOk returns a tuple with the VisDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisDistance

`func (o *FPHSpedVAPIObjectsKontorJob) SetVisDistance(v string)`

SetVisDistance sets VisDistance field to given value.


### SetVisDistanceNil

`func (o *FPHSpedVAPIObjectsKontorJob) SetVisDistanceNil(b bool)`

 SetVisDistanceNil sets the value for VisDistance to be an explicit nil

### UnsetVisDistance
`func (o *FPHSpedVAPIObjectsKontorJob) UnsetVisDistance()`

UnsetVisDistance ensures that no value is present for VisDistance, not even an explicit nil
### GetIsRealEco

`func (o *FPHSpedVAPIObjectsKontorJob) GetIsRealEco() bool`

GetIsRealEco returns the IsRealEco field if non-nil, zero value otherwise.

### GetIsRealEcoOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetIsRealEcoOk() (*bool, bool)`

GetIsRealEcoOk returns a tuple with the IsRealEco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealEco

`func (o *FPHSpedVAPIObjectsKontorJob) SetIsRealEco(v bool)`

SetIsRealEco sets IsRealEco field to given value.


### GetCompleteScheduled

`func (o *FPHSpedVAPIObjectsKontorJob) GetCompleteScheduled() bool`

GetCompleteScheduled returns the CompleteScheduled field if non-nil, zero value otherwise.

### GetCompleteScheduledOk

`func (o *FPHSpedVAPIObjectsKontorJob) GetCompleteScheduledOk() (*bool, bool)`

GetCompleteScheduledOk returns a tuple with the CompleteScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteScheduled

`func (o *FPHSpedVAPIObjectsKontorJob) SetCompleteScheduled(v bool)`

SetCompleteScheduled sets CompleteScheduled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


