# FPHSpedVAPIObjectsSpeditionsSparePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**LocaleName** | **NullableString** |  | [readonly] 
**NeededForMaintenanceKind** | [**FPHSpedVAPIEnumsMaintenanceKind**](FPHSpedVAPIEnumsMaintenanceKind.md) |   0 &#x3D; Engine  1 &#x3D; OszilationDamper  2 &#x3D; Stabilizer  3 &#x3D; StoneChip  4 &#x3D; Transmission  5 &#x3D; Wishbone  6 &#x3D; BrakePads  7 &#x3D; BrakeDiscs  8 &#x3D; EngineMaintenance  9 &#x3D; TireChange  10 &#x3D; MainCheck  11 &#x3D; SafetyCheck  12 &#x3D; SaddlePlate  13 &#x3D; AirPressureUnit  14 &#x3D; Alternator  15 &#x3D; BrakeVentil  -1 &#x3D; NotSet | [readonly] 
**PartsPerTon** | **int32** |  | [readonly] 
**PricePerPartNormal** | **int32** |  | [readonly] 
**PricePerPartRealEco** | **int32** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsSparePart

`func NewFPHSpedVAPIObjectsSpeditionsSparePart(id int32, localeName NullableString, neededForMaintenanceKind FPHSpedVAPIEnumsMaintenanceKind, partsPerTon int32, pricePerPartNormal int32, pricePerPartRealEco int32, ) *FPHSpedVAPIObjectsSpeditionsSparePart`

NewFPHSpedVAPIObjectsSpeditionsSparePart instantiates a new FPHSpedVAPIObjectsSpeditionsSparePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsSparePartWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsSparePartWithDefaults() *FPHSpedVAPIObjectsSpeditionsSparePart`

NewFPHSpedVAPIObjectsSpeditionsSparePartWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsSparePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) SetId(v int32)`

SetId sets Id field to given value.


### GetLocaleName

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) GetLocaleName() string`

GetLocaleName returns the LocaleName field if non-nil, zero value otherwise.

### GetLocaleNameOk

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) GetLocaleNameOk() (*string, bool)`

GetLocaleNameOk returns a tuple with the LocaleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaleName

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) SetLocaleName(v string)`

SetLocaleName sets LocaleName field to given value.


### SetLocaleNameNil

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) SetLocaleNameNil(b bool)`

 SetLocaleNameNil sets the value for LocaleName to be an explicit nil

### UnsetLocaleName
`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) UnsetLocaleName()`

UnsetLocaleName ensures that no value is present for LocaleName, not even an explicit nil
### GetNeededForMaintenanceKind

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) GetNeededForMaintenanceKind() FPHSpedVAPIEnumsMaintenanceKind`

GetNeededForMaintenanceKind returns the NeededForMaintenanceKind field if non-nil, zero value otherwise.

### GetNeededForMaintenanceKindOk

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) GetNeededForMaintenanceKindOk() (*FPHSpedVAPIEnumsMaintenanceKind, bool)`

GetNeededForMaintenanceKindOk returns a tuple with the NeededForMaintenanceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeededForMaintenanceKind

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) SetNeededForMaintenanceKind(v FPHSpedVAPIEnumsMaintenanceKind)`

SetNeededForMaintenanceKind sets NeededForMaintenanceKind field to given value.


### GetPartsPerTon

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) GetPartsPerTon() int32`

GetPartsPerTon returns the PartsPerTon field if non-nil, zero value otherwise.

### GetPartsPerTonOk

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) GetPartsPerTonOk() (*int32, bool)`

GetPartsPerTonOk returns a tuple with the PartsPerTon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsPerTon

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) SetPartsPerTon(v int32)`

SetPartsPerTon sets PartsPerTon field to given value.


### GetPricePerPartNormal

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) GetPricePerPartNormal() int32`

GetPricePerPartNormal returns the PricePerPartNormal field if non-nil, zero value otherwise.

### GetPricePerPartNormalOk

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) GetPricePerPartNormalOk() (*int32, bool)`

GetPricePerPartNormalOk returns a tuple with the PricePerPartNormal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerPartNormal

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) SetPricePerPartNormal(v int32)`

SetPricePerPartNormal sets PricePerPartNormal field to given value.


### GetPricePerPartRealEco

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) GetPricePerPartRealEco() int32`

GetPricePerPartRealEco returns the PricePerPartRealEco field if non-nil, zero value otherwise.

### GetPricePerPartRealEcoOk

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) GetPricePerPartRealEcoOk() (*int32, bool)`

GetPricePerPartRealEcoOk returns a tuple with the PricePerPartRealEco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerPartRealEco

`func (o *FPHSpedVAPIObjectsSpeditionsSparePart) SetPricePerPartRealEco(v int32)`

SetPricePerPartRealEco sets PricePerPartRealEco field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


