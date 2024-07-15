# FPHSpedVAPIObjectsMapsFerry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Start** | [**NullableFPHSpedVAPIObjectsMapsHarbor**](FPHSpedVAPIObjectsMapsHarbor.md) |  | [readonly] 
**Destination** | [**NullableFPHSpedVAPIObjectsMapsHarbor**](FPHSpedVAPIObjectsMapsHarbor.md) |  | [readonly] 
**Amount** | **int32** |  | [readonly] 
**Company** | **NullableString** |  | [readonly] 
**Currency** | **NullableString** |  | [readonly] 
**Distance** | **int32** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsMapsFerry

`func NewFPHSpedVAPIObjectsMapsFerry(id int32, start NullableFPHSpedVAPIObjectsMapsHarbor, destination NullableFPHSpedVAPIObjectsMapsHarbor, amount int32, company NullableString, currency NullableString, distance int32, ) *FPHSpedVAPIObjectsMapsFerry`

NewFPHSpedVAPIObjectsMapsFerry instantiates a new FPHSpedVAPIObjectsMapsFerry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsMapsFerryWithDefaults

`func NewFPHSpedVAPIObjectsMapsFerryWithDefaults() *FPHSpedVAPIObjectsMapsFerry`

NewFPHSpedVAPIObjectsMapsFerryWithDefaults instantiates a new FPHSpedVAPIObjectsMapsFerry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsMapsFerry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsMapsFerry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsMapsFerry) SetId(v int32)`

SetId sets Id field to given value.


### GetStart

`func (o *FPHSpedVAPIObjectsMapsFerry) GetStart() FPHSpedVAPIObjectsMapsHarbor`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *FPHSpedVAPIObjectsMapsFerry) GetStartOk() (*FPHSpedVAPIObjectsMapsHarbor, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *FPHSpedVAPIObjectsMapsFerry) SetStart(v FPHSpedVAPIObjectsMapsHarbor)`

SetStart sets Start field to given value.


### SetStartNil

`func (o *FPHSpedVAPIObjectsMapsFerry) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *FPHSpedVAPIObjectsMapsFerry) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetDestination

`func (o *FPHSpedVAPIObjectsMapsFerry) GetDestination() FPHSpedVAPIObjectsMapsHarbor`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *FPHSpedVAPIObjectsMapsFerry) GetDestinationOk() (*FPHSpedVAPIObjectsMapsHarbor, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *FPHSpedVAPIObjectsMapsFerry) SetDestination(v FPHSpedVAPIObjectsMapsHarbor)`

SetDestination sets Destination field to given value.


### SetDestinationNil

`func (o *FPHSpedVAPIObjectsMapsFerry) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *FPHSpedVAPIObjectsMapsFerry) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetAmount

`func (o *FPHSpedVAPIObjectsMapsFerry) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FPHSpedVAPIObjectsMapsFerry) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FPHSpedVAPIObjectsMapsFerry) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCompany

`func (o *FPHSpedVAPIObjectsMapsFerry) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *FPHSpedVAPIObjectsMapsFerry) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *FPHSpedVAPIObjectsMapsFerry) SetCompany(v string)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *FPHSpedVAPIObjectsMapsFerry) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *FPHSpedVAPIObjectsMapsFerry) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCurrency

`func (o *FPHSpedVAPIObjectsMapsFerry) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FPHSpedVAPIObjectsMapsFerry) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FPHSpedVAPIObjectsMapsFerry) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *FPHSpedVAPIObjectsMapsFerry) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *FPHSpedVAPIObjectsMapsFerry) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetDistance

`func (o *FPHSpedVAPIObjectsMapsFerry) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *FPHSpedVAPIObjectsMapsFerry) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *FPHSpedVAPIObjectsMapsFerry) SetDistance(v int32)`

SetDistance sets Distance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


