# FPHSpedVAPIObjectsKontorJobPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Spedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**Source** | [**NullableFPHSpedVAPIObjectsMapsCompanyCity**](FPHSpedVAPIObjectsMapsCompanyCity.md) |  | [readonly] 
**Destination** | [**NullableFPHSpedVAPIObjectsMapsCompanyCity**](FPHSpedVAPIObjectsMapsCompanyCity.md) |  | [readonly] 
**Driver** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**OfferUser** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**OfferDisplay** | **NullableString** |  | [readonly] 
**Value** | **int32** |  | [readonly] 
**VisValue** | **NullableString** |  | [readonly] 
**State** | [**FPHSpedVAPIEnumsJobPartState**](FPHSpedVAPIEnumsJobPartState.md) |   0 &#x3D; WaitingDriver  1 &#x3D; WaitingPrecondition  2 &#x3D; WaitingDrive  3 &#x3D; InDrive  4 &#x3D; Finished  -1 &#x3D; NotAvaliable | [readonly] 
**IsReady** | **bool** |  | [readonly] 
**Trailer** | [**NullableFPHSpedVAPIObjectsKontorLightTrailer**](FPHSpedVAPIObjectsKontorLightTrailer.md) |  | [readonly] 
**Weight** | **int32** |  | [readonly] 
**VisWeight** | **NullableString** |  | [readonly] 
**PublicState** | [**FPHSpedVAPIEnumsJobPublicState**](FPHSpedVAPIEnumsJobPublicState.md) |   0 &#x3D; Assigned  1 &#x3D; Intern  2 &#x3D; Extern  -1 &#x3D; NotAvailable | [readonly] 
**Distance** | **int32** |  | [readonly] 
**VisDistance** | **NullableString** |  | [readonly] 
**ExpirationDate** | **NullableTime** |  | [readonly] 
**CurrentConvoyInfo** | [**NullableFPHSpedVAPIObjectsLiveConvoyInfo**](FPHSpedVAPIObjectsLiveConvoyInfo.md) |  | [readonly] 
**IsRealEco** | **bool** |  | 

## Methods

### NewFPHSpedVAPIObjectsKontorJobPart

`func NewFPHSpedVAPIObjectsKontorJobPart(id int32, spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, source NullableFPHSpedVAPIObjectsMapsCompanyCity, destination NullableFPHSpedVAPIObjectsMapsCompanyCity, driver NullableFPHSpedVAPIObjectsUsersUser, offerUser NullableFPHSpedVAPIObjectsUsersUser, offerDisplay NullableString, value int32, visValue NullableString, state FPHSpedVAPIEnumsJobPartState, isReady bool, trailer NullableFPHSpedVAPIObjectsKontorLightTrailer, weight int32, visWeight NullableString, publicState FPHSpedVAPIEnumsJobPublicState, distance int32, visDistance NullableString, expirationDate NullableTime, currentConvoyInfo NullableFPHSpedVAPIObjectsLiveConvoyInfo, isRealEco bool, ) *FPHSpedVAPIObjectsKontorJobPart`

NewFPHSpedVAPIObjectsKontorJobPart instantiates a new FPHSpedVAPIObjectsKontorJobPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsKontorJobPartWithDefaults

`func NewFPHSpedVAPIObjectsKontorJobPartWithDefaults() *FPHSpedVAPIObjectsKontorJobPart`

NewFPHSpedVAPIObjectsKontorJobPartWithDefaults instantiates a new FPHSpedVAPIObjectsKontorJobPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetId(v int32)`

SetId sets Id field to given value.


### GetSpedition

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetSpedition returns the Spedition field if non-nil, zero value otherwise.

### GetSpeditionOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetSpeditionOk returns a tuple with the Spedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpedition

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetSpedition sets Spedition field to given value.


### SetSpeditionNil

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetSpeditionNil(b bool)`

 SetSpeditionNil sets the value for Spedition to be an explicit nil

### UnsetSpedition
`func (o *FPHSpedVAPIObjectsKontorJobPart) UnsetSpedition()`

UnsetSpedition ensures that no value is present for Spedition, not even an explicit nil
### GetSource

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetSource() FPHSpedVAPIObjectsMapsCompanyCity`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetSourceOk() (*FPHSpedVAPIObjectsMapsCompanyCity, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetSource(v FPHSpedVAPIObjectsMapsCompanyCity)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *FPHSpedVAPIObjectsKontorJobPart) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetDestination

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetDestination() FPHSpedVAPIObjectsMapsCompanyCity`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetDestinationOk() (*FPHSpedVAPIObjectsMapsCompanyCity, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetDestination(v FPHSpedVAPIObjectsMapsCompanyCity)`

SetDestination sets Destination field to given value.


### SetDestinationNil

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *FPHSpedVAPIObjectsKontorJobPart) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetDriver

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetDriver() FPHSpedVAPIObjectsUsersUser`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetDriverOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetDriver(v FPHSpedVAPIObjectsUsersUser)`

SetDriver sets Driver field to given value.


### SetDriverNil

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetDriverNil(b bool)`

 SetDriverNil sets the value for Driver to be an explicit nil

### UnsetDriver
`func (o *FPHSpedVAPIObjectsKontorJobPart) UnsetDriver()`

UnsetDriver ensures that no value is present for Driver, not even an explicit nil
### GetOfferUser

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetOfferUser() FPHSpedVAPIObjectsUsersUser`

GetOfferUser returns the OfferUser field if non-nil, zero value otherwise.

### GetOfferUserOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetOfferUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetOfferUserOk returns a tuple with the OfferUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferUser

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetOfferUser(v FPHSpedVAPIObjectsUsersUser)`

SetOfferUser sets OfferUser field to given value.


### SetOfferUserNil

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetOfferUserNil(b bool)`

 SetOfferUserNil sets the value for OfferUser to be an explicit nil

### UnsetOfferUser
`func (o *FPHSpedVAPIObjectsKontorJobPart) UnsetOfferUser()`

UnsetOfferUser ensures that no value is present for OfferUser, not even an explicit nil
### GetOfferDisplay

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetOfferDisplay() string`

GetOfferDisplay returns the OfferDisplay field if non-nil, zero value otherwise.

### GetOfferDisplayOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetOfferDisplayOk() (*string, bool)`

GetOfferDisplayOk returns a tuple with the OfferDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferDisplay

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetOfferDisplay(v string)`

SetOfferDisplay sets OfferDisplay field to given value.


### SetOfferDisplayNil

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetOfferDisplayNil(b bool)`

 SetOfferDisplayNil sets the value for OfferDisplay to be an explicit nil

### UnsetOfferDisplay
`func (o *FPHSpedVAPIObjectsKontorJobPart) UnsetOfferDisplay()`

UnsetOfferDisplay ensures that no value is present for OfferDisplay, not even an explicit nil
### GetValue

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetValue(v int32)`

SetValue sets Value field to given value.


### GetVisValue

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetVisValue() string`

GetVisValue returns the VisValue field if non-nil, zero value otherwise.

### GetVisValueOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetVisValueOk() (*string, bool)`

GetVisValueOk returns a tuple with the VisValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisValue

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetVisValue(v string)`

SetVisValue sets VisValue field to given value.


### SetVisValueNil

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetVisValueNil(b bool)`

 SetVisValueNil sets the value for VisValue to be an explicit nil

### UnsetVisValue
`func (o *FPHSpedVAPIObjectsKontorJobPart) UnsetVisValue()`

UnsetVisValue ensures that no value is present for VisValue, not even an explicit nil
### GetState

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetState() FPHSpedVAPIEnumsJobPartState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetStateOk() (*FPHSpedVAPIEnumsJobPartState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetState(v FPHSpedVAPIEnumsJobPartState)`

SetState sets State field to given value.


### GetIsReady

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.


### GetTrailer

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetTrailer() FPHSpedVAPIObjectsKontorLightTrailer`

GetTrailer returns the Trailer field if non-nil, zero value otherwise.

### GetTrailerOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetTrailerOk() (*FPHSpedVAPIObjectsKontorLightTrailer, bool)`

GetTrailerOk returns a tuple with the Trailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailer

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetTrailer(v FPHSpedVAPIObjectsKontorLightTrailer)`

SetTrailer sets Trailer field to given value.


### SetTrailerNil

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetTrailerNil(b bool)`

 SetTrailerNil sets the value for Trailer to be an explicit nil

### UnsetTrailer
`func (o *FPHSpedVAPIObjectsKontorJobPart) UnsetTrailer()`

UnsetTrailer ensures that no value is present for Trailer, not even an explicit nil
### GetWeight

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetVisWeight

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetVisWeight() string`

GetVisWeight returns the VisWeight field if non-nil, zero value otherwise.

### GetVisWeightOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetVisWeightOk() (*string, bool)`

GetVisWeightOk returns a tuple with the VisWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisWeight

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetVisWeight(v string)`

SetVisWeight sets VisWeight field to given value.


### SetVisWeightNil

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetVisWeightNil(b bool)`

 SetVisWeightNil sets the value for VisWeight to be an explicit nil

### UnsetVisWeight
`func (o *FPHSpedVAPIObjectsKontorJobPart) UnsetVisWeight()`

UnsetVisWeight ensures that no value is present for VisWeight, not even an explicit nil
### GetPublicState

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetPublicState() FPHSpedVAPIEnumsJobPublicState`

GetPublicState returns the PublicState field if non-nil, zero value otherwise.

### GetPublicStateOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetPublicStateOk() (*FPHSpedVAPIEnumsJobPublicState, bool)`

GetPublicStateOk returns a tuple with the PublicState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicState

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetPublicState(v FPHSpedVAPIEnumsJobPublicState)`

SetPublicState sets PublicState field to given value.


### GetDistance

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetDistance(v int32)`

SetDistance sets Distance field to given value.


### GetVisDistance

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetVisDistance() string`

GetVisDistance returns the VisDistance field if non-nil, zero value otherwise.

### GetVisDistanceOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetVisDistanceOk() (*string, bool)`

GetVisDistanceOk returns a tuple with the VisDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisDistance

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetVisDistance(v string)`

SetVisDistance sets VisDistance field to given value.


### SetVisDistanceNil

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetVisDistanceNil(b bool)`

 SetVisDistanceNil sets the value for VisDistance to be an explicit nil

### UnsetVisDistance
`func (o *FPHSpedVAPIObjectsKontorJobPart) UnsetVisDistance()`

UnsetVisDistance ensures that no value is present for VisDistance, not even an explicit nil
### GetExpirationDate

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.


### SetExpirationDateNil

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *FPHSpedVAPIObjectsKontorJobPart) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetCurrentConvoyInfo

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetCurrentConvoyInfo() FPHSpedVAPIObjectsLiveConvoyInfo`

GetCurrentConvoyInfo returns the CurrentConvoyInfo field if non-nil, zero value otherwise.

### GetCurrentConvoyInfoOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetCurrentConvoyInfoOk() (*FPHSpedVAPIObjectsLiveConvoyInfo, bool)`

GetCurrentConvoyInfoOk returns a tuple with the CurrentConvoyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConvoyInfo

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetCurrentConvoyInfo(v FPHSpedVAPIObjectsLiveConvoyInfo)`

SetCurrentConvoyInfo sets CurrentConvoyInfo field to given value.


### SetCurrentConvoyInfoNil

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetCurrentConvoyInfoNil(b bool)`

 SetCurrentConvoyInfoNil sets the value for CurrentConvoyInfo to be an explicit nil

### UnsetCurrentConvoyInfo
`func (o *FPHSpedVAPIObjectsKontorJobPart) UnsetCurrentConvoyInfo()`

UnsetCurrentConvoyInfo ensures that no value is present for CurrentConvoyInfo, not even an explicit nil
### GetIsRealEco

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetIsRealEco() bool`

GetIsRealEco returns the IsRealEco field if non-nil, zero value otherwise.

### GetIsRealEcoOk

`func (o *FPHSpedVAPIObjectsKontorJobPart) GetIsRealEcoOk() (*bool, bool)`

GetIsRealEcoOk returns a tuple with the IsRealEco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealEco

`func (o *FPHSpedVAPIObjectsKontorJobPart) SetIsRealEco(v bool)`

SetIsRealEco sets IsRealEco field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


