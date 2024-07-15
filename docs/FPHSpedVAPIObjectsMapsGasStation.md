# FPHSpedVAPIObjectsMapsGasStation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**LiterPrize** | **float64** |  | [readonly] 
**LastReported** | **time.Time** |  | [readonly] 
**Currency** | **NullableString** |  | [readonly] 
**X** | **float64** |  | [readonly] 
**Z** | **float64** |  | [readonly] 
**ReportedBy** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsMapsGasStation

`func NewFPHSpedVAPIObjectsMapsGasStation(id int32, literPrize float64, lastReported time.Time, currency NullableString, x float64, z float64, reportedBy NullableFPHSpedVAPIObjectsUsersUser, ) *FPHSpedVAPIObjectsMapsGasStation`

NewFPHSpedVAPIObjectsMapsGasStation instantiates a new FPHSpedVAPIObjectsMapsGasStation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsMapsGasStationWithDefaults

`func NewFPHSpedVAPIObjectsMapsGasStationWithDefaults() *FPHSpedVAPIObjectsMapsGasStation`

NewFPHSpedVAPIObjectsMapsGasStationWithDefaults instantiates a new FPHSpedVAPIObjectsMapsGasStation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsMapsGasStation) SetId(v int32)`

SetId sets Id field to given value.


### GetLiterPrize

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetLiterPrize() float64`

GetLiterPrize returns the LiterPrize field if non-nil, zero value otherwise.

### GetLiterPrizeOk

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetLiterPrizeOk() (*float64, bool)`

GetLiterPrizeOk returns a tuple with the LiterPrize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiterPrize

`func (o *FPHSpedVAPIObjectsMapsGasStation) SetLiterPrize(v float64)`

SetLiterPrize sets LiterPrize field to given value.


### GetLastReported

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetLastReported() time.Time`

GetLastReported returns the LastReported field if non-nil, zero value otherwise.

### GetLastReportedOk

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetLastReportedOk() (*time.Time, bool)`

GetLastReportedOk returns a tuple with the LastReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReported

`func (o *FPHSpedVAPIObjectsMapsGasStation) SetLastReported(v time.Time)`

SetLastReported sets LastReported field to given value.


### GetCurrency

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FPHSpedVAPIObjectsMapsGasStation) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *FPHSpedVAPIObjectsMapsGasStation) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *FPHSpedVAPIObjectsMapsGasStation) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetX

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *FPHSpedVAPIObjectsMapsGasStation) SetX(v float64)`

SetX sets X field to given value.


### GetZ

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetZ() float64`

GetZ returns the Z field if non-nil, zero value otherwise.

### GetZOk

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetZOk() (*float64, bool)`

GetZOk returns a tuple with the Z field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZ

`func (o *FPHSpedVAPIObjectsMapsGasStation) SetZ(v float64)`

SetZ sets Z field to given value.


### GetReportedBy

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetReportedBy() FPHSpedVAPIObjectsUsersUser`

GetReportedBy returns the ReportedBy field if non-nil, zero value otherwise.

### GetReportedByOk

`func (o *FPHSpedVAPIObjectsMapsGasStation) GetReportedByOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetReportedByOk returns a tuple with the ReportedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedBy

`func (o *FPHSpedVAPIObjectsMapsGasStation) SetReportedBy(v FPHSpedVAPIObjectsUsersUser)`

SetReportedBy sets ReportedBy field to given value.


### SetReportedByNil

`func (o *FPHSpedVAPIObjectsMapsGasStation) SetReportedByNil(b bool)`

 SetReportedByNil sets the value for ReportedBy to be an explicit nil

### UnsetReportedBy
`func (o *FPHSpedVAPIObjectsMapsGasStation) UnsetReportedBy()`

UnsetReportedBy ensures that no value is present for ReportedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


