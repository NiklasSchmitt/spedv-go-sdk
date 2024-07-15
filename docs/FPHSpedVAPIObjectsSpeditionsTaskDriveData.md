# FPHSpedVAPIObjectsSpeditionsTaskDriveData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Timestamp** | **time.Time** |  | 
**Velocity** | **int32** |  | 
**SpeedLimit** | **int32** |  | 
**CruiseControl** | **int32** |  | 
**XCoord** | **float64** |  | 
**ZCoord** | **float64** |  | 
**MpServer** | **NullableString** |  | 
**Odometer** | **NullableFloat64** |  | 
**FreightDamagePercent** | **NullableFloat64** |  | 
**TruckDamagePercent** | **NullableFloat64** |  | 
**TruckFuel** | **NullableFloat64** |  | 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsTaskDriveData

`func NewFPHSpedVAPIObjectsSpeditionsTaskDriveData(id int64, timestamp time.Time, velocity int32, speedLimit int32, cruiseControl int32, xCoord float64, zCoord float64, mpServer NullableString, odometer NullableFloat64, freightDamagePercent NullableFloat64, truckDamagePercent NullableFloat64, truckFuel NullableFloat64, ) *FPHSpedVAPIObjectsSpeditionsTaskDriveData`

NewFPHSpedVAPIObjectsSpeditionsTaskDriveData instantiates a new FPHSpedVAPIObjectsSpeditionsTaskDriveData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsTaskDriveDataWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsTaskDriveDataWithDefaults() *FPHSpedVAPIObjectsSpeditionsTaskDriveData`

NewFPHSpedVAPIObjectsSpeditionsTaskDriveDataWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsTaskDriveData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetId(v int64)`

SetId sets Id field to given value.


### GetTimestamp

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetVelocity

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetVelocity() int32`

GetVelocity returns the Velocity field if non-nil, zero value otherwise.

### GetVelocityOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetVelocityOk() (*int32, bool)`

GetVelocityOk returns a tuple with the Velocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocity

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetVelocity(v int32)`

SetVelocity sets Velocity field to given value.


### GetSpeedLimit

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetSpeedLimit() int32`

GetSpeedLimit returns the SpeedLimit field if non-nil, zero value otherwise.

### GetSpeedLimitOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetSpeedLimitOk() (*int32, bool)`

GetSpeedLimitOk returns a tuple with the SpeedLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedLimit

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetSpeedLimit(v int32)`

SetSpeedLimit sets SpeedLimit field to given value.


### GetCruiseControl

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetCruiseControl() int32`

GetCruiseControl returns the CruiseControl field if non-nil, zero value otherwise.

### GetCruiseControlOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetCruiseControlOk() (*int32, bool)`

GetCruiseControlOk returns a tuple with the CruiseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCruiseControl

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetCruiseControl(v int32)`

SetCruiseControl sets CruiseControl field to given value.


### GetXCoord

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetXCoord() float64`

GetXCoord returns the XCoord field if non-nil, zero value otherwise.

### GetXCoordOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetXCoordOk() (*float64, bool)`

GetXCoordOk returns a tuple with the XCoord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXCoord

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetXCoord(v float64)`

SetXCoord sets XCoord field to given value.


### GetZCoord

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetZCoord() float64`

GetZCoord returns the ZCoord field if non-nil, zero value otherwise.

### GetZCoordOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetZCoordOk() (*float64, bool)`

GetZCoordOk returns a tuple with the ZCoord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZCoord

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetZCoord(v float64)`

SetZCoord sets ZCoord field to given value.


### GetMpServer

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetMpServer() string`

GetMpServer returns the MpServer field if non-nil, zero value otherwise.

### GetMpServerOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetMpServerOk() (*string, bool)`

GetMpServerOk returns a tuple with the MpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpServer

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetMpServer(v string)`

SetMpServer sets MpServer field to given value.


### SetMpServerNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetMpServerNil(b bool)`

 SetMpServerNil sets the value for MpServer to be an explicit nil

### UnsetMpServer
`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) UnsetMpServer()`

UnsetMpServer ensures that no value is present for MpServer, not even an explicit nil
### GetOdometer

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetOdometer() float64`

GetOdometer returns the Odometer field if non-nil, zero value otherwise.

### GetOdometerOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetOdometerOk() (*float64, bool)`

GetOdometerOk returns a tuple with the Odometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometer

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetOdometer(v float64)`

SetOdometer sets Odometer field to given value.


### SetOdometerNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetOdometerNil(b bool)`

 SetOdometerNil sets the value for Odometer to be an explicit nil

### UnsetOdometer
`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) UnsetOdometer()`

UnsetOdometer ensures that no value is present for Odometer, not even an explicit nil
### GetFreightDamagePercent

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetFreightDamagePercent() float64`

GetFreightDamagePercent returns the FreightDamagePercent field if non-nil, zero value otherwise.

### GetFreightDamagePercentOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetFreightDamagePercentOk() (*float64, bool)`

GetFreightDamagePercentOk returns a tuple with the FreightDamagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightDamagePercent

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetFreightDamagePercent(v float64)`

SetFreightDamagePercent sets FreightDamagePercent field to given value.


### SetFreightDamagePercentNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetFreightDamagePercentNil(b bool)`

 SetFreightDamagePercentNil sets the value for FreightDamagePercent to be an explicit nil

### UnsetFreightDamagePercent
`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) UnsetFreightDamagePercent()`

UnsetFreightDamagePercent ensures that no value is present for FreightDamagePercent, not even an explicit nil
### GetTruckDamagePercent

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetTruckDamagePercent() float64`

GetTruckDamagePercent returns the TruckDamagePercent field if non-nil, zero value otherwise.

### GetTruckDamagePercentOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetTruckDamagePercentOk() (*float64, bool)`

GetTruckDamagePercentOk returns a tuple with the TruckDamagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckDamagePercent

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetTruckDamagePercent(v float64)`

SetTruckDamagePercent sets TruckDamagePercent field to given value.


### SetTruckDamagePercentNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetTruckDamagePercentNil(b bool)`

 SetTruckDamagePercentNil sets the value for TruckDamagePercent to be an explicit nil

### UnsetTruckDamagePercent
`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) UnsetTruckDamagePercent()`

UnsetTruckDamagePercent ensures that no value is present for TruckDamagePercent, not even an explicit nil
### GetTruckFuel

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetTruckFuel() float64`

GetTruckFuel returns the TruckFuel field if non-nil, zero value otherwise.

### GetTruckFuelOk

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) GetTruckFuelOk() (*float64, bool)`

GetTruckFuelOk returns a tuple with the TruckFuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckFuel

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetTruckFuel(v float64)`

SetTruckFuel sets TruckFuel field to given value.


### SetTruckFuelNil

`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) SetTruckFuelNil(b bool)`

 SetTruckFuelNil sets the value for TruckFuel to be an explicit nil

### UnsetTruckFuel
`func (o *FPHSpedVAPIObjectsSpeditionsTaskDriveData) UnsetTruckFuel()`

UnsetTruckFuel ensures that no value is present for TruckFuel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


