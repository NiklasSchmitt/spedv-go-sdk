# FPHSpedVAPIObjectsKontorJobOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | [readonly] 
**Source** | [**NullableFPHSpedVAPIObjectsMapsCompanyCity**](FPHSpedVAPIObjectsMapsCompanyCity.md) |  | [readonly] 
**Destination** | [**NullableFPHSpedVAPIObjectsMapsCompanyCity**](FPHSpedVAPIObjectsMapsCompanyCity.md) |  | [readonly] 
**Weight** | **int32** |  | [readonly] 
**Freight** | [**NullableFPHSpedVAPIObjectsKontorJobOfferFreight**](FPHSpedVAPIObjectsKontorJobOfferFreight.md) |  | [readonly] 
**PayingSpedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**SparePartsForBranchID** | **int32** |  | [readonly] 
**Value** | **int32** |  | [readonly] 
**ExpirationDate** | **time.Time** |  | [readonly] 
**Distance** | **int32** |  | [readonly] 
**VisDistance** | **NullableString** |  | [readonly] 
**VisWeight** | **NullableString** |  | [readonly] 
**VisValue** | **NullableString** |  | [readonly] 
**IsRealEco** | **bool** |  | 

## Methods

### NewFPHSpedVAPIObjectsKontorJobOffer

`func NewFPHSpedVAPIObjectsKontorJobOffer(id NullableString, source NullableFPHSpedVAPIObjectsMapsCompanyCity, destination NullableFPHSpedVAPIObjectsMapsCompanyCity, weight int32, freight NullableFPHSpedVAPIObjectsKontorJobOfferFreight, payingSpedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, sparePartsForBranchID int32, value int32, expirationDate time.Time, distance int32, visDistance NullableString, visWeight NullableString, visValue NullableString, isRealEco bool, ) *FPHSpedVAPIObjectsKontorJobOffer`

NewFPHSpedVAPIObjectsKontorJobOffer instantiates a new FPHSpedVAPIObjectsKontorJobOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsKontorJobOfferWithDefaults

`func NewFPHSpedVAPIObjectsKontorJobOfferWithDefaults() *FPHSpedVAPIObjectsKontorJobOffer`

NewFPHSpedVAPIObjectsKontorJobOfferWithDefaults instantiates a new FPHSpedVAPIObjectsKontorJobOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FPHSpedVAPIObjectsKontorJobOffer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSource

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetSource() FPHSpedVAPIObjectsMapsCompanyCity`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetSourceOk() (*FPHSpedVAPIObjectsMapsCompanyCity, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetSource(v FPHSpedVAPIObjectsMapsCompanyCity)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *FPHSpedVAPIObjectsKontorJobOffer) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetDestination

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetDestination() FPHSpedVAPIObjectsMapsCompanyCity`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetDestinationOk() (*FPHSpedVAPIObjectsMapsCompanyCity, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetDestination(v FPHSpedVAPIObjectsMapsCompanyCity)`

SetDestination sets Destination field to given value.


### SetDestinationNil

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *FPHSpedVAPIObjectsKontorJobOffer) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetWeight

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetFreight

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetFreight() FPHSpedVAPIObjectsKontorJobOfferFreight`

GetFreight returns the Freight field if non-nil, zero value otherwise.

### GetFreightOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetFreightOk() (*FPHSpedVAPIObjectsKontorJobOfferFreight, bool)`

GetFreightOk returns a tuple with the Freight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreight

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetFreight(v FPHSpedVAPIObjectsKontorJobOfferFreight)`

SetFreight sets Freight field to given value.


### SetFreightNil

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetFreightNil(b bool)`

 SetFreightNil sets the value for Freight to be an explicit nil

### UnsetFreight
`func (o *FPHSpedVAPIObjectsKontorJobOffer) UnsetFreight()`

UnsetFreight ensures that no value is present for Freight, not even an explicit nil
### GetPayingSpedition

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetPayingSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetPayingSpedition returns the PayingSpedition field if non-nil, zero value otherwise.

### GetPayingSpeditionOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetPayingSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetPayingSpeditionOk returns a tuple with the PayingSpedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayingSpedition

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetPayingSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetPayingSpedition sets PayingSpedition field to given value.


### SetPayingSpeditionNil

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetPayingSpeditionNil(b bool)`

 SetPayingSpeditionNil sets the value for PayingSpedition to be an explicit nil

### UnsetPayingSpedition
`func (o *FPHSpedVAPIObjectsKontorJobOffer) UnsetPayingSpedition()`

UnsetPayingSpedition ensures that no value is present for PayingSpedition, not even an explicit nil
### GetSparePartsForBranchID

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetSparePartsForBranchID() int32`

GetSparePartsForBranchID returns the SparePartsForBranchID field if non-nil, zero value otherwise.

### GetSparePartsForBranchIDOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetSparePartsForBranchIDOk() (*int32, bool)`

GetSparePartsForBranchIDOk returns a tuple with the SparePartsForBranchID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparePartsForBranchID

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetSparePartsForBranchID(v int32)`

SetSparePartsForBranchID sets SparePartsForBranchID field to given value.


### GetValue

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetValue(v int32)`

SetValue sets Value field to given value.


### GetExpirationDate

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.


### GetDistance

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetDistance(v int32)`

SetDistance sets Distance field to given value.


### GetVisDistance

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetVisDistance() string`

GetVisDistance returns the VisDistance field if non-nil, zero value otherwise.

### GetVisDistanceOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetVisDistanceOk() (*string, bool)`

GetVisDistanceOk returns a tuple with the VisDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisDistance

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetVisDistance(v string)`

SetVisDistance sets VisDistance field to given value.


### SetVisDistanceNil

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetVisDistanceNil(b bool)`

 SetVisDistanceNil sets the value for VisDistance to be an explicit nil

### UnsetVisDistance
`func (o *FPHSpedVAPIObjectsKontorJobOffer) UnsetVisDistance()`

UnsetVisDistance ensures that no value is present for VisDistance, not even an explicit nil
### GetVisWeight

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetVisWeight() string`

GetVisWeight returns the VisWeight field if non-nil, zero value otherwise.

### GetVisWeightOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetVisWeightOk() (*string, bool)`

GetVisWeightOk returns a tuple with the VisWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisWeight

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetVisWeight(v string)`

SetVisWeight sets VisWeight field to given value.


### SetVisWeightNil

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetVisWeightNil(b bool)`

 SetVisWeightNil sets the value for VisWeight to be an explicit nil

### UnsetVisWeight
`func (o *FPHSpedVAPIObjectsKontorJobOffer) UnsetVisWeight()`

UnsetVisWeight ensures that no value is present for VisWeight, not even an explicit nil
### GetVisValue

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetVisValue() string`

GetVisValue returns the VisValue field if non-nil, zero value otherwise.

### GetVisValueOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetVisValueOk() (*string, bool)`

GetVisValueOk returns a tuple with the VisValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisValue

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetVisValue(v string)`

SetVisValue sets VisValue field to given value.


### SetVisValueNil

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetVisValueNil(b bool)`

 SetVisValueNil sets the value for VisValue to be an explicit nil

### UnsetVisValue
`func (o *FPHSpedVAPIObjectsKontorJobOffer) UnsetVisValue()`

UnsetVisValue ensures that no value is present for VisValue, not even an explicit nil
### GetIsRealEco

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetIsRealEco() bool`

GetIsRealEco returns the IsRealEco field if non-nil, zero value otherwise.

### GetIsRealEcoOk

`func (o *FPHSpedVAPIObjectsKontorJobOffer) GetIsRealEcoOk() (*bool, bool)`

GetIsRealEcoOk returns a tuple with the IsRealEco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealEco

`func (o *FPHSpedVAPIObjectsKontorJobOffer) SetIsRealEco(v bool)`

SetIsRealEco sets IsRealEco field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


