# FPHSpedVAPIObjectsSpeditionsTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | 
**Id** | **int32** |  | 
**VisibleID** | **NullableString** |  | [readonly] 
**IsDeductable** | **bool** |  | [readonly] 
**User** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | 
**Maps** | [**[]FPHSpedVAPIObjectsMapsMap**](FPHSpedVAPIObjectsMapsMap.md) |  | 
**VisMaps** | **NullableString** |  | [readonly] 
**StartCity** | [**NullableFPHSpedVAPIObjectsMapsCity**](FPHSpedVAPIObjectsMapsCity.md) |  | [readonly] 
**DestCity** | [**NullableFPHSpedVAPIObjectsMapsCity**](FPHSpedVAPIObjectsMapsCity.md) |  | [readonly] 
**StartCompany** | [**NullableFPHSpedVAPIObjectsMapsCompany**](FPHSpedVAPIObjectsMapsCompany.md) |  | [readonly] 
**DestCompany** | [**NullableFPHSpedVAPIObjectsMapsCompany**](FPHSpedVAPIObjectsMapsCompany.md) |  | [readonly] 
**Freight** | [**NullableFPHSpedVAPIObjectsMapsFreight**](FPHSpedVAPIObjectsMapsFreight.md) |  | 
**Freightweight** | **int32** |  | 
**VisFreightWeight** | **NullableString** |  | [readonly] 
**DmgReport** | **NullableString** |  | 
**Value** | **float64** |  | 
**Damage** | **float64** |  | 
**DamagePercent** | **float64** |  | [readonly] 
**VisDamagePercent** | **NullableString** |  | [readonly] 
**Taxes** | **float64** |  | 
**Maintenance** | **float64** |  | 
**Toll** | **float64** |  | 
**Income** | **float64** |  | [readonly] 
**State** | [**FPHSpedVAPIEnumsETSTaskState**](FPHSpedVAPIEnumsETSTaskState.md) |   0 &#x3D; InDrive  1 &#x3D; Done  2 &#x3D; Settled  3 &#x3D; Fail  4 &#x3D; AdminCheck  5 &#x3D; Paused  6 &#x3D; Cancelled  7 &#x3D; Invalid  -1 &#x3D; NotAvaliable | 
**Starttime** | **time.Time** |  | 
**Endtime** | **time.Time** |  | 
**NeededTime** | **string** |  | [readonly] 
**Ferry** | **int32** |  | 
**FerryKM** | **int32** |  | 
**Refuel** | **int32** |  | 
**Boni** | **float64** |  | 
**Currency** | **NullableString** |  | 
**CurrencySymbol** | **NullableString** |  | [readonly] 
**Truck** | [**NullableFPHSpedVAPIObjectsSpeditionsTruck**](FPHSpedVAPIObjectsSpeditionsTruck.md) |  | 
**TruckType** | [**NullableFPHSpedVAPIObjectsSpeditionsTruckType**](FPHSpedVAPIObjectsSpeditionsTruckType.md) |  | 
**RentedTruck** | **bool** |  | [readonly] 
**Startkm** | **int32** |  | 
**Endkm** | **int32** |  | 
**DistKM** | **float64** |  | [readonly] 
**DistMI** | **float64** |  | [readonly] 
**VisDistance** | **NullableString** |  | [readonly] 
**Eurokm** | **float64** |  | [readonly] 
**Dollarmi** | **float64** |  | [readonly] 
**VisEuroKM** | **NullableString** |  | [readonly] 
**FuelUsedLi** | **int32** |  | [readonly] 
**FuelUsedGal** | **float64** |  | [readonly] 
**FuelRefuledLi** | **int32** |  | [readonly] 
**FuelRefuledGal** | **float64** |  | [readonly] 
**FuelAvg100KM** | **float64** |  | [readonly] 
**FuelAvgMiGal** | **float64** |  | [readonly] 
**VisFuelEfficiency** | **NullableString** |  | [readonly] 
**TruckDamage** | **float64** |  | [readonly] 
**MaxVelocityKMH** | **int32** |  | [readonly] 
**MaxVelocityMPH** | **int32** |  | [readonly] 
**AvgVelocityKMH** | **int32** |  | [readonly] 
**AvgVelocityMPH** | **int32** |  | [readonly] 
**DdCleaned** | **bool** |  | 
**CargoMarket** | [**FPHSpedVAPIEnumsCargoMarketType**](FPHSpedVAPIEnumsCargoMarketType.md) |   0 &#x3D; CargoMarket  1 &#x3D; QuickJob  2 &#x3D; FreightMarket  3 &#x3D; ExternalContract  4 &#x3D; ExternalMarket  -1 &#x3D; NotSet | 
**ScreenshotState** | [**FPHSpedVAPIEnumsScreenshotState**](FPHSpedVAPIEnumsScreenshotState.md) |   0 &#x3D; None  1 &#x3D; Uploaded  2 &#x3D; Processed  3 &#x3D; Rejected  4 &#x3D; TaskSettled | 
**ScreenshotCheckedBy** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | 
**ScreenshotCheckDate** | **time.Time** |  | 
**KontorPart** | [**NullableFPHSpedVAPIObjectsKontorJobPart**](FPHSpedVAPIObjectsKontorJobPart.md) |  | [readonly] 
**UserSetRealEco** | **bool** |  | 
**IsRealScale** | **bool** |  | 
**TaskPositions** | [**[]FPHSpedVAPIObjectsSpeditionsTaskPosition**](FPHSpedVAPIObjectsSpeditionsTaskPosition.md) |  | 
**CurrentConvoyInfo** | [**NullableFPHSpedVAPIObjectsLiveConvoyInfo**](FPHSpedVAPIObjectsLiveConvoyInfo.md) |  | [readonly] 
**MeetsSpeditionVelocityRequirements** | **bool** |  | [readonly] 
**IsOK** | **bool** |  | [readonly] 
**IsQuestionable** | **bool** |  | [readonly] 
**IsNegative** | **bool** |  | [readonly] 
**Fail** | **float64** |  | [readonly] 
**ScreenshotURL** | **NullableString** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsTask

`func NewFPHSpedVAPIObjectsSpeditionsTask(spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, id int32, visibleID NullableString, isDeductable bool, user NullableFPHSpedVAPIObjectsUsersUser, maps []FPHSpedVAPIObjectsMapsMap, visMaps NullableString, startCity NullableFPHSpedVAPIObjectsMapsCity, destCity NullableFPHSpedVAPIObjectsMapsCity, startCompany NullableFPHSpedVAPIObjectsMapsCompany, destCompany NullableFPHSpedVAPIObjectsMapsCompany, freight NullableFPHSpedVAPIObjectsMapsFreight, freightweight int32, visFreightWeight NullableString, dmgReport NullableString, value float64, damage float64, damagePercent float64, visDamagePercent NullableString, taxes float64, maintenance float64, toll float64, income float64, state FPHSpedVAPIEnumsETSTaskState, starttime time.Time, endtime time.Time, neededTime string, ferry int32, ferryKM int32, refuel int32, boni float64, currency NullableString, currencySymbol NullableString, truck NullableFPHSpedVAPIObjectsSpeditionsTruck, truckType NullableFPHSpedVAPIObjectsSpeditionsTruckType, rentedTruck bool, startkm int32, endkm int32, distKM float64, distMI float64, visDistance NullableString, eurokm float64, dollarmi float64, visEuroKM NullableString, fuelUsedLi int32, fuelUsedGal float64, fuelRefuledLi int32, fuelRefuledGal float64, fuelAvg100KM float64, fuelAvgMiGal float64, visFuelEfficiency NullableString, truckDamage float64, maxVelocityKMH int32, maxVelocityMPH int32, avgVelocityKMH int32, avgVelocityMPH int32, ddCleaned bool, cargoMarket FPHSpedVAPIEnumsCargoMarketType, screenshotState FPHSpedVAPIEnumsScreenshotState, screenshotCheckedBy NullableFPHSpedVAPIObjectsUsersUser, screenshotCheckDate time.Time, kontorPart NullableFPHSpedVAPIObjectsKontorJobPart, userSetRealEco bool, isRealScale bool, taskPositions []FPHSpedVAPIObjectsSpeditionsTaskPosition, currentConvoyInfo NullableFPHSpedVAPIObjectsLiveConvoyInfo, meetsSpeditionVelocityRequirements bool, isOK bool, isQuestionable bool, isNegative bool, fail float64, screenshotURL NullableString, ) *FPHSpedVAPIObjectsSpeditionsTask`

NewFPHSpedVAPIObjectsSpeditionsTask instantiates a new FPHSpedVAPIObjectsSpeditionsTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsTaskWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsTaskWithDefaults() *FPHSpedVAPIObjectsSpeditionsTask`

NewFPHSpedVAPIObjectsSpeditionsTaskWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpedition

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetSpedition returns the Spedition field if non-nil, zero value otherwise.

### GetSpeditionOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetSpeditionOk returns a tuple with the Spedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpedition

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetSpedition sets Spedition field to given value.


### SetSpeditionNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetSpeditionNil(b bool)`

 SetSpeditionNil sets the value for Spedition to be an explicit nil

### UnsetSpedition
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetSpedition()`

UnsetSpedition ensures that no value is present for Spedition, not even an explicit nil
### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetId(v int32)`

SetId sets Id field to given value.


### GetVisibleID

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisibleID() string`

GetVisibleID returns the VisibleID field if non-nil, zero value otherwise.

### GetVisibleIDOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisibleIDOk() (*string, bool)`

GetVisibleIDOk returns a tuple with the VisibleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleID

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisibleID(v string)`

SetVisibleID sets VisibleID field to given value.


### SetVisibleIDNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisibleIDNil(b bool)`

 SetVisibleIDNil sets the value for VisibleID to be an explicit nil

### UnsetVisibleID
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetVisibleID()`

UnsetVisibleID ensures that no value is present for VisibleID, not even an explicit nil
### GetIsDeductable

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIsDeductable() bool`

GetIsDeductable returns the IsDeductable field if non-nil, zero value otherwise.

### GetIsDeductableOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIsDeductableOk() (*bool, bool)`

GetIsDeductableOk returns a tuple with the IsDeductable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeductable

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetIsDeductable(v bool)`

SetIsDeductable sets IsDeductable field to given value.


### GetUser

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetUser() FPHSpedVAPIObjectsUsersUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetUser(v FPHSpedVAPIObjectsUsersUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetMaps

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetMaps() []FPHSpedVAPIObjectsMapsMap`

GetMaps returns the Maps field if non-nil, zero value otherwise.

### GetMapsOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetMapsOk() (*[]FPHSpedVAPIObjectsMapsMap, bool)`

GetMapsOk returns a tuple with the Maps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaps

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetMaps(v []FPHSpedVAPIObjectsMapsMap)`

SetMaps sets Maps field to given value.


### SetMapsNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetMapsNil(b bool)`

 SetMapsNil sets the value for Maps to be an explicit nil

### UnsetMaps
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetMaps()`

UnsetMaps ensures that no value is present for Maps, not even an explicit nil
### GetVisMaps

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisMaps() string`

GetVisMaps returns the VisMaps field if non-nil, zero value otherwise.

### GetVisMapsOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisMapsOk() (*string, bool)`

GetVisMapsOk returns a tuple with the VisMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisMaps

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisMaps(v string)`

SetVisMaps sets VisMaps field to given value.


### SetVisMapsNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisMapsNil(b bool)`

 SetVisMapsNil sets the value for VisMaps to be an explicit nil

### UnsetVisMaps
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetVisMaps()`

UnsetVisMaps ensures that no value is present for VisMaps, not even an explicit nil
### GetStartCity

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetStartCity() FPHSpedVAPIObjectsMapsCity`

GetStartCity returns the StartCity field if non-nil, zero value otherwise.

### GetStartCityOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetStartCityOk() (*FPHSpedVAPIObjectsMapsCity, bool)`

GetStartCityOk returns a tuple with the StartCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCity

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetStartCity(v FPHSpedVAPIObjectsMapsCity)`

SetStartCity sets StartCity field to given value.


### SetStartCityNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetStartCityNil(b bool)`

 SetStartCityNil sets the value for StartCity to be an explicit nil

### UnsetStartCity
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetStartCity()`

UnsetStartCity ensures that no value is present for StartCity, not even an explicit nil
### GetDestCity

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDestCity() FPHSpedVAPIObjectsMapsCity`

GetDestCity returns the DestCity field if non-nil, zero value otherwise.

### GetDestCityOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDestCityOk() (*FPHSpedVAPIObjectsMapsCity, bool)`

GetDestCityOk returns a tuple with the DestCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCity

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetDestCity(v FPHSpedVAPIObjectsMapsCity)`

SetDestCity sets DestCity field to given value.


### SetDestCityNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetDestCityNil(b bool)`

 SetDestCityNil sets the value for DestCity to be an explicit nil

### UnsetDestCity
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetDestCity()`

UnsetDestCity ensures that no value is present for DestCity, not even an explicit nil
### GetStartCompany

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetStartCompany() FPHSpedVAPIObjectsMapsCompany`

GetStartCompany returns the StartCompany field if non-nil, zero value otherwise.

### GetStartCompanyOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetStartCompanyOk() (*FPHSpedVAPIObjectsMapsCompany, bool)`

GetStartCompanyOk returns a tuple with the StartCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCompany

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetStartCompany(v FPHSpedVAPIObjectsMapsCompany)`

SetStartCompany sets StartCompany field to given value.


### SetStartCompanyNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetStartCompanyNil(b bool)`

 SetStartCompanyNil sets the value for StartCompany to be an explicit nil

### UnsetStartCompany
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetStartCompany()`

UnsetStartCompany ensures that no value is present for StartCompany, not even an explicit nil
### GetDestCompany

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDestCompany() FPHSpedVAPIObjectsMapsCompany`

GetDestCompany returns the DestCompany field if non-nil, zero value otherwise.

### GetDestCompanyOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDestCompanyOk() (*FPHSpedVAPIObjectsMapsCompany, bool)`

GetDestCompanyOk returns a tuple with the DestCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCompany

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetDestCompany(v FPHSpedVAPIObjectsMapsCompany)`

SetDestCompany sets DestCompany field to given value.


### SetDestCompanyNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetDestCompanyNil(b bool)`

 SetDestCompanyNil sets the value for DestCompany to be an explicit nil

### UnsetDestCompany
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetDestCompany()`

UnsetDestCompany ensures that no value is present for DestCompany, not even an explicit nil
### GetFreight

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFreight() FPHSpedVAPIObjectsMapsFreight`

GetFreight returns the Freight field if non-nil, zero value otherwise.

### GetFreightOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFreightOk() (*FPHSpedVAPIObjectsMapsFreight, bool)`

GetFreightOk returns a tuple with the Freight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreight

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetFreight(v FPHSpedVAPIObjectsMapsFreight)`

SetFreight sets Freight field to given value.


### SetFreightNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetFreightNil(b bool)`

 SetFreightNil sets the value for Freight to be an explicit nil

### UnsetFreight
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetFreight()`

UnsetFreight ensures that no value is present for Freight, not even an explicit nil
### GetFreightweight

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFreightweight() int32`

GetFreightweight returns the Freightweight field if non-nil, zero value otherwise.

### GetFreightweightOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFreightweightOk() (*int32, bool)`

GetFreightweightOk returns a tuple with the Freightweight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightweight

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetFreightweight(v int32)`

SetFreightweight sets Freightweight field to given value.


### GetVisFreightWeight

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisFreightWeight() string`

GetVisFreightWeight returns the VisFreightWeight field if non-nil, zero value otherwise.

### GetVisFreightWeightOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisFreightWeightOk() (*string, bool)`

GetVisFreightWeightOk returns a tuple with the VisFreightWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisFreightWeight

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisFreightWeight(v string)`

SetVisFreightWeight sets VisFreightWeight field to given value.


### SetVisFreightWeightNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisFreightWeightNil(b bool)`

 SetVisFreightWeightNil sets the value for VisFreightWeight to be an explicit nil

### UnsetVisFreightWeight
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetVisFreightWeight()`

UnsetVisFreightWeight ensures that no value is present for VisFreightWeight, not even an explicit nil
### GetDmgReport

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDmgReport() string`

GetDmgReport returns the DmgReport field if non-nil, zero value otherwise.

### GetDmgReportOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDmgReportOk() (*string, bool)`

GetDmgReportOk returns a tuple with the DmgReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmgReport

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetDmgReport(v string)`

SetDmgReport sets DmgReport field to given value.


### SetDmgReportNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetDmgReportNil(b bool)`

 SetDmgReportNil sets the value for DmgReport to be an explicit nil

### UnsetDmgReport
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetDmgReport()`

UnsetDmgReport ensures that no value is present for DmgReport, not even an explicit nil
### GetValue

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetValue(v float64)`

SetValue sets Value field to given value.


### GetDamage

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDamage() float64`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDamageOk() (*float64, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetDamage(v float64)`

SetDamage sets Damage field to given value.


### GetDamagePercent

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDamagePercent() float64`

GetDamagePercent returns the DamagePercent field if non-nil, zero value otherwise.

### GetDamagePercentOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDamagePercentOk() (*float64, bool)`

GetDamagePercentOk returns a tuple with the DamagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamagePercent

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetDamagePercent(v float64)`

SetDamagePercent sets DamagePercent field to given value.


### GetVisDamagePercent

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisDamagePercent() string`

GetVisDamagePercent returns the VisDamagePercent field if non-nil, zero value otherwise.

### GetVisDamagePercentOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisDamagePercentOk() (*string, bool)`

GetVisDamagePercentOk returns a tuple with the VisDamagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisDamagePercent

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisDamagePercent(v string)`

SetVisDamagePercent sets VisDamagePercent field to given value.


### SetVisDamagePercentNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisDamagePercentNil(b bool)`

 SetVisDamagePercentNil sets the value for VisDamagePercent to be an explicit nil

### UnsetVisDamagePercent
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetVisDamagePercent()`

UnsetVisDamagePercent ensures that no value is present for VisDamagePercent, not even an explicit nil
### GetTaxes

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetTaxes() float64`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetTaxesOk() (*float64, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetTaxes(v float64)`

SetTaxes sets Taxes field to given value.


### GetMaintenance

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetMaintenance() float64`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetMaintenanceOk() (*float64, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetMaintenance(v float64)`

SetMaintenance sets Maintenance field to given value.


### GetToll

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetToll() float64`

GetToll returns the Toll field if non-nil, zero value otherwise.

### GetTollOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetTollOk() (*float64, bool)`

GetTollOk returns a tuple with the Toll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToll

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetToll(v float64)`

SetToll sets Toll field to given value.


### GetIncome

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIncome() float64`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIncomeOk() (*float64, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetIncome(v float64)`

SetIncome sets Income field to given value.


### GetState

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetState() FPHSpedVAPIEnumsETSTaskState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetStateOk() (*FPHSpedVAPIEnumsETSTaskState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetState(v FPHSpedVAPIEnumsETSTaskState)`

SetState sets State field to given value.


### GetStarttime

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetStarttime() time.Time`

GetStarttime returns the Starttime field if non-nil, zero value otherwise.

### GetStarttimeOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetStarttimeOk() (*time.Time, bool)`

GetStarttimeOk returns a tuple with the Starttime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarttime

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetStarttime(v time.Time)`

SetStarttime sets Starttime field to given value.


### GetEndtime

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetEndtime() time.Time`

GetEndtime returns the Endtime field if non-nil, zero value otherwise.

### GetEndtimeOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetEndtimeOk() (*time.Time, bool)`

GetEndtimeOk returns a tuple with the Endtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndtime

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetEndtime(v time.Time)`

SetEndtime sets Endtime field to given value.


### GetNeededTime

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetNeededTime() string`

GetNeededTime returns the NeededTime field if non-nil, zero value otherwise.

### GetNeededTimeOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetNeededTimeOk() (*string, bool)`

GetNeededTimeOk returns a tuple with the NeededTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeededTime

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetNeededTime(v string)`

SetNeededTime sets NeededTime field to given value.


### GetFerry

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFerry() int32`

GetFerry returns the Ferry field if non-nil, zero value otherwise.

### GetFerryOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFerryOk() (*int32, bool)`

GetFerryOk returns a tuple with the Ferry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFerry

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetFerry(v int32)`

SetFerry sets Ferry field to given value.


### GetFerryKM

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFerryKM() int32`

GetFerryKM returns the FerryKM field if non-nil, zero value otherwise.

### GetFerryKMOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFerryKMOk() (*int32, bool)`

GetFerryKMOk returns a tuple with the FerryKM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFerryKM

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetFerryKM(v int32)`

SetFerryKM sets FerryKM field to given value.


### GetRefuel

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetRefuel() int32`

GetRefuel returns the Refuel field if non-nil, zero value otherwise.

### GetRefuelOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetRefuelOk() (*int32, bool)`

GetRefuelOk returns a tuple with the Refuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefuel

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetRefuel(v int32)`

SetRefuel sets Refuel field to given value.


### GetBoni

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetBoni() float64`

GetBoni returns the Boni field if non-nil, zero value otherwise.

### GetBoniOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetBoniOk() (*float64, bool)`

GetBoniOk returns a tuple with the Boni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoni

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetBoni(v float64)`

SetBoni sets Boni field to given value.


### GetCurrency

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCurrencySymbol

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.


### SetCurrencySymbolNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetCurrencySymbolNil(b bool)`

 SetCurrencySymbolNil sets the value for CurrencySymbol to be an explicit nil

### UnsetCurrencySymbol
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetCurrencySymbol()`

UnsetCurrencySymbol ensures that no value is present for CurrencySymbol, not even an explicit nil
### GetTruck

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetTruck() FPHSpedVAPIObjectsSpeditionsTruck`

GetTruck returns the Truck field if non-nil, zero value otherwise.

### GetTruckOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetTruckOk() (*FPHSpedVAPIObjectsSpeditionsTruck, bool)`

GetTruckOk returns a tuple with the Truck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruck

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetTruck(v FPHSpedVAPIObjectsSpeditionsTruck)`

SetTruck sets Truck field to given value.


### SetTruckNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetTruckNil(b bool)`

 SetTruckNil sets the value for Truck to be an explicit nil

### UnsetTruck
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetTruck()`

UnsetTruck ensures that no value is present for Truck, not even an explicit nil
### GetTruckType

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetTruckType() FPHSpedVAPIObjectsSpeditionsTruckType`

GetTruckType returns the TruckType field if non-nil, zero value otherwise.

### GetTruckTypeOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetTruckTypeOk() (*FPHSpedVAPIObjectsSpeditionsTruckType, bool)`

GetTruckTypeOk returns a tuple with the TruckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckType

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetTruckType(v FPHSpedVAPIObjectsSpeditionsTruckType)`

SetTruckType sets TruckType field to given value.


### SetTruckTypeNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetTruckTypeNil(b bool)`

 SetTruckTypeNil sets the value for TruckType to be an explicit nil

### UnsetTruckType
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetTruckType()`

UnsetTruckType ensures that no value is present for TruckType, not even an explicit nil
### GetRentedTruck

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetRentedTruck() bool`

GetRentedTruck returns the RentedTruck field if non-nil, zero value otherwise.

### GetRentedTruckOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetRentedTruckOk() (*bool, bool)`

GetRentedTruckOk returns a tuple with the RentedTruck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRentedTruck

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetRentedTruck(v bool)`

SetRentedTruck sets RentedTruck field to given value.


### GetStartkm

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetStartkm() int32`

GetStartkm returns the Startkm field if non-nil, zero value otherwise.

### GetStartkmOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetStartkmOk() (*int32, bool)`

GetStartkmOk returns a tuple with the Startkm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartkm

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetStartkm(v int32)`

SetStartkm sets Startkm field to given value.


### GetEndkm

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetEndkm() int32`

GetEndkm returns the Endkm field if non-nil, zero value otherwise.

### GetEndkmOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetEndkmOk() (*int32, bool)`

GetEndkmOk returns a tuple with the Endkm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndkm

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetEndkm(v int32)`

SetEndkm sets Endkm field to given value.


### GetDistKM

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDistKM() float64`

GetDistKM returns the DistKM field if non-nil, zero value otherwise.

### GetDistKMOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDistKMOk() (*float64, bool)`

GetDistKMOk returns a tuple with the DistKM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistKM

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetDistKM(v float64)`

SetDistKM sets DistKM field to given value.


### GetDistMI

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDistMI() float64`

GetDistMI returns the DistMI field if non-nil, zero value otherwise.

### GetDistMIOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDistMIOk() (*float64, bool)`

GetDistMIOk returns a tuple with the DistMI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistMI

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetDistMI(v float64)`

SetDistMI sets DistMI field to given value.


### GetVisDistance

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisDistance() string`

GetVisDistance returns the VisDistance field if non-nil, zero value otherwise.

### GetVisDistanceOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisDistanceOk() (*string, bool)`

GetVisDistanceOk returns a tuple with the VisDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisDistance

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisDistance(v string)`

SetVisDistance sets VisDistance field to given value.


### SetVisDistanceNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisDistanceNil(b bool)`

 SetVisDistanceNil sets the value for VisDistance to be an explicit nil

### UnsetVisDistance
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetVisDistance()`

UnsetVisDistance ensures that no value is present for VisDistance, not even an explicit nil
### GetEurokm

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetEurokm() float64`

GetEurokm returns the Eurokm field if non-nil, zero value otherwise.

### GetEurokmOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetEurokmOk() (*float64, bool)`

GetEurokmOk returns a tuple with the Eurokm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEurokm

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetEurokm(v float64)`

SetEurokm sets Eurokm field to given value.


### GetDollarmi

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDollarmi() float64`

GetDollarmi returns the Dollarmi field if non-nil, zero value otherwise.

### GetDollarmiOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDollarmiOk() (*float64, bool)`

GetDollarmiOk returns a tuple with the Dollarmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDollarmi

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetDollarmi(v float64)`

SetDollarmi sets Dollarmi field to given value.


### GetVisEuroKM

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisEuroKM() string`

GetVisEuroKM returns the VisEuroKM field if non-nil, zero value otherwise.

### GetVisEuroKMOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisEuroKMOk() (*string, bool)`

GetVisEuroKMOk returns a tuple with the VisEuroKM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisEuroKM

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisEuroKM(v string)`

SetVisEuroKM sets VisEuroKM field to given value.


### SetVisEuroKMNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisEuroKMNil(b bool)`

 SetVisEuroKMNil sets the value for VisEuroKM to be an explicit nil

### UnsetVisEuroKM
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetVisEuroKM()`

UnsetVisEuroKM ensures that no value is present for VisEuroKM, not even an explicit nil
### GetFuelUsedLi

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFuelUsedLi() int32`

GetFuelUsedLi returns the FuelUsedLi field if non-nil, zero value otherwise.

### GetFuelUsedLiOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFuelUsedLiOk() (*int32, bool)`

GetFuelUsedLiOk returns a tuple with the FuelUsedLi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelUsedLi

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetFuelUsedLi(v int32)`

SetFuelUsedLi sets FuelUsedLi field to given value.


### GetFuelUsedGal

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFuelUsedGal() float64`

GetFuelUsedGal returns the FuelUsedGal field if non-nil, zero value otherwise.

### GetFuelUsedGalOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFuelUsedGalOk() (*float64, bool)`

GetFuelUsedGalOk returns a tuple with the FuelUsedGal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelUsedGal

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetFuelUsedGal(v float64)`

SetFuelUsedGal sets FuelUsedGal field to given value.


### GetFuelRefuledLi

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFuelRefuledLi() int32`

GetFuelRefuledLi returns the FuelRefuledLi field if non-nil, zero value otherwise.

### GetFuelRefuledLiOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFuelRefuledLiOk() (*int32, bool)`

GetFuelRefuledLiOk returns a tuple with the FuelRefuledLi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelRefuledLi

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetFuelRefuledLi(v int32)`

SetFuelRefuledLi sets FuelRefuledLi field to given value.


### GetFuelRefuledGal

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFuelRefuledGal() float64`

GetFuelRefuledGal returns the FuelRefuledGal field if non-nil, zero value otherwise.

### GetFuelRefuledGalOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFuelRefuledGalOk() (*float64, bool)`

GetFuelRefuledGalOk returns a tuple with the FuelRefuledGal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelRefuledGal

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetFuelRefuledGal(v float64)`

SetFuelRefuledGal sets FuelRefuledGal field to given value.


### GetFuelAvg100KM

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFuelAvg100KM() float64`

GetFuelAvg100KM returns the FuelAvg100KM field if non-nil, zero value otherwise.

### GetFuelAvg100KMOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFuelAvg100KMOk() (*float64, bool)`

GetFuelAvg100KMOk returns a tuple with the FuelAvg100KM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelAvg100KM

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetFuelAvg100KM(v float64)`

SetFuelAvg100KM sets FuelAvg100KM field to given value.


### GetFuelAvgMiGal

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFuelAvgMiGal() float64`

GetFuelAvgMiGal returns the FuelAvgMiGal field if non-nil, zero value otherwise.

### GetFuelAvgMiGalOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFuelAvgMiGalOk() (*float64, bool)`

GetFuelAvgMiGalOk returns a tuple with the FuelAvgMiGal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelAvgMiGal

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetFuelAvgMiGal(v float64)`

SetFuelAvgMiGal sets FuelAvgMiGal field to given value.


### GetVisFuelEfficiency

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisFuelEfficiency() string`

GetVisFuelEfficiency returns the VisFuelEfficiency field if non-nil, zero value otherwise.

### GetVisFuelEfficiencyOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetVisFuelEfficiencyOk() (*string, bool)`

GetVisFuelEfficiencyOk returns a tuple with the VisFuelEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisFuelEfficiency

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisFuelEfficiency(v string)`

SetVisFuelEfficiency sets VisFuelEfficiency field to given value.


### SetVisFuelEfficiencyNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetVisFuelEfficiencyNil(b bool)`

 SetVisFuelEfficiencyNil sets the value for VisFuelEfficiency to be an explicit nil

### UnsetVisFuelEfficiency
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetVisFuelEfficiency()`

UnsetVisFuelEfficiency ensures that no value is present for VisFuelEfficiency, not even an explicit nil
### GetTruckDamage

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetTruckDamage() float64`

GetTruckDamage returns the TruckDamage field if non-nil, zero value otherwise.

### GetTruckDamageOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetTruckDamageOk() (*float64, bool)`

GetTruckDamageOk returns a tuple with the TruckDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckDamage

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetTruckDamage(v float64)`

SetTruckDamage sets TruckDamage field to given value.


### GetMaxVelocityKMH

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetMaxVelocityKMH() int32`

GetMaxVelocityKMH returns the MaxVelocityKMH field if non-nil, zero value otherwise.

### GetMaxVelocityKMHOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetMaxVelocityKMHOk() (*int32, bool)`

GetMaxVelocityKMHOk returns a tuple with the MaxVelocityKMH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVelocityKMH

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetMaxVelocityKMH(v int32)`

SetMaxVelocityKMH sets MaxVelocityKMH field to given value.


### GetMaxVelocityMPH

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetMaxVelocityMPH() int32`

GetMaxVelocityMPH returns the MaxVelocityMPH field if non-nil, zero value otherwise.

### GetMaxVelocityMPHOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetMaxVelocityMPHOk() (*int32, bool)`

GetMaxVelocityMPHOk returns a tuple with the MaxVelocityMPH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVelocityMPH

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetMaxVelocityMPH(v int32)`

SetMaxVelocityMPH sets MaxVelocityMPH field to given value.


### GetAvgVelocityKMH

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetAvgVelocityKMH() int32`

GetAvgVelocityKMH returns the AvgVelocityKMH field if non-nil, zero value otherwise.

### GetAvgVelocityKMHOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetAvgVelocityKMHOk() (*int32, bool)`

GetAvgVelocityKMHOk returns a tuple with the AvgVelocityKMH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgVelocityKMH

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetAvgVelocityKMH(v int32)`

SetAvgVelocityKMH sets AvgVelocityKMH field to given value.


### GetAvgVelocityMPH

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetAvgVelocityMPH() int32`

GetAvgVelocityMPH returns the AvgVelocityMPH field if non-nil, zero value otherwise.

### GetAvgVelocityMPHOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetAvgVelocityMPHOk() (*int32, bool)`

GetAvgVelocityMPHOk returns a tuple with the AvgVelocityMPH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgVelocityMPH

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetAvgVelocityMPH(v int32)`

SetAvgVelocityMPH sets AvgVelocityMPH field to given value.


### GetDdCleaned

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDdCleaned() bool`

GetDdCleaned returns the DdCleaned field if non-nil, zero value otherwise.

### GetDdCleanedOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetDdCleanedOk() (*bool, bool)`

GetDdCleanedOk returns a tuple with the DdCleaned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdCleaned

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetDdCleaned(v bool)`

SetDdCleaned sets DdCleaned field to given value.


### GetCargoMarket

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetCargoMarket() FPHSpedVAPIEnumsCargoMarketType`

GetCargoMarket returns the CargoMarket field if non-nil, zero value otherwise.

### GetCargoMarketOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetCargoMarketOk() (*FPHSpedVAPIEnumsCargoMarketType, bool)`

GetCargoMarketOk returns a tuple with the CargoMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoMarket

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetCargoMarket(v FPHSpedVAPIEnumsCargoMarketType)`

SetCargoMarket sets CargoMarket field to given value.


### GetScreenshotState

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetScreenshotState() FPHSpedVAPIEnumsScreenshotState`

GetScreenshotState returns the ScreenshotState field if non-nil, zero value otherwise.

### GetScreenshotStateOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetScreenshotStateOk() (*FPHSpedVAPIEnumsScreenshotState, bool)`

GetScreenshotStateOk returns a tuple with the ScreenshotState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotState

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetScreenshotState(v FPHSpedVAPIEnumsScreenshotState)`

SetScreenshotState sets ScreenshotState field to given value.


### GetScreenshotCheckedBy

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetScreenshotCheckedBy() FPHSpedVAPIObjectsUsersUser`

GetScreenshotCheckedBy returns the ScreenshotCheckedBy field if non-nil, zero value otherwise.

### GetScreenshotCheckedByOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetScreenshotCheckedByOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetScreenshotCheckedByOk returns a tuple with the ScreenshotCheckedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotCheckedBy

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetScreenshotCheckedBy(v FPHSpedVAPIObjectsUsersUser)`

SetScreenshotCheckedBy sets ScreenshotCheckedBy field to given value.


### SetScreenshotCheckedByNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetScreenshotCheckedByNil(b bool)`

 SetScreenshotCheckedByNil sets the value for ScreenshotCheckedBy to be an explicit nil

### UnsetScreenshotCheckedBy
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetScreenshotCheckedBy()`

UnsetScreenshotCheckedBy ensures that no value is present for ScreenshotCheckedBy, not even an explicit nil
### GetScreenshotCheckDate

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetScreenshotCheckDate() time.Time`

GetScreenshotCheckDate returns the ScreenshotCheckDate field if non-nil, zero value otherwise.

### GetScreenshotCheckDateOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetScreenshotCheckDateOk() (*time.Time, bool)`

GetScreenshotCheckDateOk returns a tuple with the ScreenshotCheckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotCheckDate

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetScreenshotCheckDate(v time.Time)`

SetScreenshotCheckDate sets ScreenshotCheckDate field to given value.


### GetKontorPart

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetKontorPart() FPHSpedVAPIObjectsKontorJobPart`

GetKontorPart returns the KontorPart field if non-nil, zero value otherwise.

### GetKontorPartOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetKontorPartOk() (*FPHSpedVAPIObjectsKontorJobPart, bool)`

GetKontorPartOk returns a tuple with the KontorPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKontorPart

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetKontorPart(v FPHSpedVAPIObjectsKontorJobPart)`

SetKontorPart sets KontorPart field to given value.


### SetKontorPartNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetKontorPartNil(b bool)`

 SetKontorPartNil sets the value for KontorPart to be an explicit nil

### UnsetKontorPart
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetKontorPart()`

UnsetKontorPart ensures that no value is present for KontorPart, not even an explicit nil
### GetUserSetRealEco

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetUserSetRealEco() bool`

GetUserSetRealEco returns the UserSetRealEco field if non-nil, zero value otherwise.

### GetUserSetRealEcoOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetUserSetRealEcoOk() (*bool, bool)`

GetUserSetRealEcoOk returns a tuple with the UserSetRealEco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSetRealEco

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetUserSetRealEco(v bool)`

SetUserSetRealEco sets UserSetRealEco field to given value.


### GetIsRealScale

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIsRealScale() bool`

GetIsRealScale returns the IsRealScale field if non-nil, zero value otherwise.

### GetIsRealScaleOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIsRealScaleOk() (*bool, bool)`

GetIsRealScaleOk returns a tuple with the IsRealScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealScale

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetIsRealScale(v bool)`

SetIsRealScale sets IsRealScale field to given value.


### GetTaskPositions

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetTaskPositions() []FPHSpedVAPIObjectsSpeditionsTaskPosition`

GetTaskPositions returns the TaskPositions field if non-nil, zero value otherwise.

### GetTaskPositionsOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetTaskPositionsOk() (*[]FPHSpedVAPIObjectsSpeditionsTaskPosition, bool)`

GetTaskPositionsOk returns a tuple with the TaskPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPositions

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetTaskPositions(v []FPHSpedVAPIObjectsSpeditionsTaskPosition)`

SetTaskPositions sets TaskPositions field to given value.


### SetTaskPositionsNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetTaskPositionsNil(b bool)`

 SetTaskPositionsNil sets the value for TaskPositions to be an explicit nil

### UnsetTaskPositions
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetTaskPositions()`

UnsetTaskPositions ensures that no value is present for TaskPositions, not even an explicit nil
### GetCurrentConvoyInfo

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetCurrentConvoyInfo() FPHSpedVAPIObjectsLiveConvoyInfo`

GetCurrentConvoyInfo returns the CurrentConvoyInfo field if non-nil, zero value otherwise.

### GetCurrentConvoyInfoOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetCurrentConvoyInfoOk() (*FPHSpedVAPIObjectsLiveConvoyInfo, bool)`

GetCurrentConvoyInfoOk returns a tuple with the CurrentConvoyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConvoyInfo

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetCurrentConvoyInfo(v FPHSpedVAPIObjectsLiveConvoyInfo)`

SetCurrentConvoyInfo sets CurrentConvoyInfo field to given value.


### SetCurrentConvoyInfoNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetCurrentConvoyInfoNil(b bool)`

 SetCurrentConvoyInfoNil sets the value for CurrentConvoyInfo to be an explicit nil

### UnsetCurrentConvoyInfo
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetCurrentConvoyInfo()`

UnsetCurrentConvoyInfo ensures that no value is present for CurrentConvoyInfo, not even an explicit nil
### GetMeetsSpeditionVelocityRequirements

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetMeetsSpeditionVelocityRequirements() bool`

GetMeetsSpeditionVelocityRequirements returns the MeetsSpeditionVelocityRequirements field if non-nil, zero value otherwise.

### GetMeetsSpeditionVelocityRequirementsOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetMeetsSpeditionVelocityRequirementsOk() (*bool, bool)`

GetMeetsSpeditionVelocityRequirementsOk returns a tuple with the MeetsSpeditionVelocityRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetsSpeditionVelocityRequirements

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetMeetsSpeditionVelocityRequirements(v bool)`

SetMeetsSpeditionVelocityRequirements sets MeetsSpeditionVelocityRequirements field to given value.


### GetIsOK

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIsOK() bool`

GetIsOK returns the IsOK field if non-nil, zero value otherwise.

### GetIsOKOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIsOKOk() (*bool, bool)`

GetIsOKOk returns a tuple with the IsOK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOK

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetIsOK(v bool)`

SetIsOK sets IsOK field to given value.


### GetIsQuestionable

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIsQuestionable() bool`

GetIsQuestionable returns the IsQuestionable field if non-nil, zero value otherwise.

### GetIsQuestionableOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIsQuestionableOk() (*bool, bool)`

GetIsQuestionableOk returns a tuple with the IsQuestionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuestionable

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetIsQuestionable(v bool)`

SetIsQuestionable sets IsQuestionable field to given value.


### GetIsNegative

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIsNegative() bool`

GetIsNegative returns the IsNegative field if non-nil, zero value otherwise.

### GetIsNegativeOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetIsNegativeOk() (*bool, bool)`

GetIsNegativeOk returns a tuple with the IsNegative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNegative

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetIsNegative(v bool)`

SetIsNegative sets IsNegative field to given value.


### GetFail

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFail() float64`

GetFail returns the Fail field if non-nil, zero value otherwise.

### GetFailOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetFailOk() (*float64, bool)`

GetFailOk returns a tuple with the Fail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFail

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetFail(v float64)`

SetFail sets Fail field to given value.


### GetScreenshotURL

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetScreenshotURL() string`

GetScreenshotURL returns the ScreenshotURL field if non-nil, zero value otherwise.

### GetScreenshotURLOk

`func (o *FPHSpedVAPIObjectsSpeditionsTask) GetScreenshotURLOk() (*string, bool)`

GetScreenshotURLOk returns a tuple with the ScreenshotURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotURL

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetScreenshotURL(v string)`

SetScreenshotURL sets ScreenshotURL field to given value.


### SetScreenshotURLNil

`func (o *FPHSpedVAPIObjectsSpeditionsTask) SetScreenshotURLNil(b bool)`

 SetScreenshotURLNil sets the value for ScreenshotURL to be an explicit nil

### UnsetScreenshotURL
`func (o *FPHSpedVAPIObjectsSpeditionsTask) UnsetScreenshotURL()`

UnsetScreenshotURL ensures that no value is present for ScreenshotURL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


