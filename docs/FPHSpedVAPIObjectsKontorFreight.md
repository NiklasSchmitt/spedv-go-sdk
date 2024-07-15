# FPHSpedVAPIObjectsKontorFreight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **NullableString** |  | [readonly] 
**PrizeCoeffizient** | **float64** |  | [readonly] 
**Enabled** | **bool** |  | [readonly] 
**IsSpecialFreight** | **bool** |  | [readonly] 
**SourceAllowedCompanies** | [**[]FPHSpedVAPIObjectsMapsCompanyCategory**](FPHSpedVAPIObjectsMapsCompanyCategory.md) |  | [readonly] 
**SinkAllowedCompanies** | [**[]FPHSpedVAPIObjectsMapsCompanyCategory**](FPHSpedVAPIObjectsMapsCompanyCategory.md) |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsKontorFreight

`func NewFPHSpedVAPIObjectsKontorFreight(id int32, name NullableString, prizeCoeffizient float64, enabled bool, isSpecialFreight bool, sourceAllowedCompanies []FPHSpedVAPIObjectsMapsCompanyCategory, sinkAllowedCompanies []FPHSpedVAPIObjectsMapsCompanyCategory, ) *FPHSpedVAPIObjectsKontorFreight`

NewFPHSpedVAPIObjectsKontorFreight instantiates a new FPHSpedVAPIObjectsKontorFreight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsKontorFreightWithDefaults

`func NewFPHSpedVAPIObjectsKontorFreightWithDefaults() *FPHSpedVAPIObjectsKontorFreight`

NewFPHSpedVAPIObjectsKontorFreightWithDefaults instantiates a new FPHSpedVAPIObjectsKontorFreight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsKontorFreight) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsKontorFreight) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsKontorFreight) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FPHSpedVAPIObjectsKontorFreight) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPHSpedVAPIObjectsKontorFreight) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPHSpedVAPIObjectsKontorFreight) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *FPHSpedVAPIObjectsKontorFreight) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FPHSpedVAPIObjectsKontorFreight) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPrizeCoeffizient

`func (o *FPHSpedVAPIObjectsKontorFreight) GetPrizeCoeffizient() float64`

GetPrizeCoeffizient returns the PrizeCoeffizient field if non-nil, zero value otherwise.

### GetPrizeCoeffizientOk

`func (o *FPHSpedVAPIObjectsKontorFreight) GetPrizeCoeffizientOk() (*float64, bool)`

GetPrizeCoeffizientOk returns a tuple with the PrizeCoeffizient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrizeCoeffizient

`func (o *FPHSpedVAPIObjectsKontorFreight) SetPrizeCoeffizient(v float64)`

SetPrizeCoeffizient sets PrizeCoeffizient field to given value.


### GetEnabled

`func (o *FPHSpedVAPIObjectsKontorFreight) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FPHSpedVAPIObjectsKontorFreight) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FPHSpedVAPIObjectsKontorFreight) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIsSpecialFreight

`func (o *FPHSpedVAPIObjectsKontorFreight) GetIsSpecialFreight() bool`

GetIsSpecialFreight returns the IsSpecialFreight field if non-nil, zero value otherwise.

### GetIsSpecialFreightOk

`func (o *FPHSpedVAPIObjectsKontorFreight) GetIsSpecialFreightOk() (*bool, bool)`

GetIsSpecialFreightOk returns a tuple with the IsSpecialFreight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialFreight

`func (o *FPHSpedVAPIObjectsKontorFreight) SetIsSpecialFreight(v bool)`

SetIsSpecialFreight sets IsSpecialFreight field to given value.


### GetSourceAllowedCompanies

`func (o *FPHSpedVAPIObjectsKontorFreight) GetSourceAllowedCompanies() []FPHSpedVAPIObjectsMapsCompanyCategory`

GetSourceAllowedCompanies returns the SourceAllowedCompanies field if non-nil, zero value otherwise.

### GetSourceAllowedCompaniesOk

`func (o *FPHSpedVAPIObjectsKontorFreight) GetSourceAllowedCompaniesOk() (*[]FPHSpedVAPIObjectsMapsCompanyCategory, bool)`

GetSourceAllowedCompaniesOk returns a tuple with the SourceAllowedCompanies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAllowedCompanies

`func (o *FPHSpedVAPIObjectsKontorFreight) SetSourceAllowedCompanies(v []FPHSpedVAPIObjectsMapsCompanyCategory)`

SetSourceAllowedCompanies sets SourceAllowedCompanies field to given value.


### SetSourceAllowedCompaniesNil

`func (o *FPHSpedVAPIObjectsKontorFreight) SetSourceAllowedCompaniesNil(b bool)`

 SetSourceAllowedCompaniesNil sets the value for SourceAllowedCompanies to be an explicit nil

### UnsetSourceAllowedCompanies
`func (o *FPHSpedVAPIObjectsKontorFreight) UnsetSourceAllowedCompanies()`

UnsetSourceAllowedCompanies ensures that no value is present for SourceAllowedCompanies, not even an explicit nil
### GetSinkAllowedCompanies

`func (o *FPHSpedVAPIObjectsKontorFreight) GetSinkAllowedCompanies() []FPHSpedVAPIObjectsMapsCompanyCategory`

GetSinkAllowedCompanies returns the SinkAllowedCompanies field if non-nil, zero value otherwise.

### GetSinkAllowedCompaniesOk

`func (o *FPHSpedVAPIObjectsKontorFreight) GetSinkAllowedCompaniesOk() (*[]FPHSpedVAPIObjectsMapsCompanyCategory, bool)`

GetSinkAllowedCompaniesOk returns a tuple with the SinkAllowedCompanies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinkAllowedCompanies

`func (o *FPHSpedVAPIObjectsKontorFreight) SetSinkAllowedCompanies(v []FPHSpedVAPIObjectsMapsCompanyCategory)`

SetSinkAllowedCompanies sets SinkAllowedCompanies field to given value.


### SetSinkAllowedCompaniesNil

`func (o *FPHSpedVAPIObjectsKontorFreight) SetSinkAllowedCompaniesNil(b bool)`

 SetSinkAllowedCompaniesNil sets the value for SinkAllowedCompanies to be an explicit nil

### UnsetSinkAllowedCompanies
`func (o *FPHSpedVAPIObjectsKontorFreight) UnsetSinkAllowedCompanies()`

UnsetSinkAllowedCompanies ensures that no value is present for SinkAllowedCompanies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


