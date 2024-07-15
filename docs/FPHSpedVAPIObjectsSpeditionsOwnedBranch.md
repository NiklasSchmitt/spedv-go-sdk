# FPHSpedVAPIObjectsSpeditionsOwnedBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Spedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**CompanyCity** | [**NullableFPHSpedVAPIObjectsMapsCompanyCity**](FPHSpedVAPIObjectsMapsCompanyCity.md) |  | [readonly] 
**Comment** | **NullableString** |  | [readonly] 
**MaintenancePlaces** | **int32** |  | [readonly] 
**TruckParkplaces** | **int32** |  | [readonly] 
**TrailerParkplaces** | **int32** |  | [readonly] 
**BreakbulkCargoPlace** | **int32** |  | [readonly] 
**TruckParkplacesUsed** | **int32** |  | [readonly] 
**MaintenancePlacesUsed** | **int32** |  | [readonly] 
**SparePartsInStorage** | [**[]FPHSpedVAPIObjectsSpeditionsSparePartStock**](FPHSpedVAPIObjectsSpeditionsSparePartStock.md) |  | [readonly] 
**TruckParkplacesNotFull** | **bool** |  | [readonly] 
**MaintenancePlacesNotFull** | **bool** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsOwnedBranch

`func NewFPHSpedVAPIObjectsSpeditionsOwnedBranch(id int32, spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, companyCity NullableFPHSpedVAPIObjectsMapsCompanyCity, comment NullableString, maintenancePlaces int32, truckParkplaces int32, trailerParkplaces int32, breakbulkCargoPlace int32, truckParkplacesUsed int32, maintenancePlacesUsed int32, sparePartsInStorage []FPHSpedVAPIObjectsSpeditionsSparePartStock, truckParkplacesNotFull bool, maintenancePlacesNotFull bool, ) *FPHSpedVAPIObjectsSpeditionsOwnedBranch`

NewFPHSpedVAPIObjectsSpeditionsOwnedBranch instantiates a new FPHSpedVAPIObjectsSpeditionsOwnedBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsOwnedBranchWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsOwnedBranchWithDefaults() *FPHSpedVAPIObjectsSpeditionsOwnedBranch`

NewFPHSpedVAPIObjectsSpeditionsOwnedBranchWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsOwnedBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetId(v int32)`

SetId sets Id field to given value.


### GetSpedition

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetSpedition returns the Spedition field if non-nil, zero value otherwise.

### GetSpeditionOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetSpeditionOk returns a tuple with the Spedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpedition

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetSpedition sets Spedition field to given value.


### SetSpeditionNil

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetSpeditionNil(b bool)`

 SetSpeditionNil sets the value for Spedition to be an explicit nil

### UnsetSpedition
`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) UnsetSpedition()`

UnsetSpedition ensures that no value is present for Spedition, not even an explicit nil
### GetCompanyCity

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetCompanyCity() FPHSpedVAPIObjectsMapsCompanyCity`

GetCompanyCity returns the CompanyCity field if non-nil, zero value otherwise.

### GetCompanyCityOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetCompanyCityOk() (*FPHSpedVAPIObjectsMapsCompanyCity, bool)`

GetCompanyCityOk returns a tuple with the CompanyCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCity

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetCompanyCity(v FPHSpedVAPIObjectsMapsCompanyCity)`

SetCompanyCity sets CompanyCity field to given value.


### SetCompanyCityNil

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetCompanyCityNil(b bool)`

 SetCompanyCityNil sets the value for CompanyCity to be an explicit nil

### UnsetCompanyCity
`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) UnsetCompanyCity()`

UnsetCompanyCity ensures that no value is present for CompanyCity, not even an explicit nil
### GetComment

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetMaintenancePlaces

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetMaintenancePlaces() int32`

GetMaintenancePlaces returns the MaintenancePlaces field if non-nil, zero value otherwise.

### GetMaintenancePlacesOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetMaintenancePlacesOk() (*int32, bool)`

GetMaintenancePlacesOk returns a tuple with the MaintenancePlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenancePlaces

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetMaintenancePlaces(v int32)`

SetMaintenancePlaces sets MaintenancePlaces field to given value.


### GetTruckParkplaces

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetTruckParkplaces() int32`

GetTruckParkplaces returns the TruckParkplaces field if non-nil, zero value otherwise.

### GetTruckParkplacesOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetTruckParkplacesOk() (*int32, bool)`

GetTruckParkplacesOk returns a tuple with the TruckParkplaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckParkplaces

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetTruckParkplaces(v int32)`

SetTruckParkplaces sets TruckParkplaces field to given value.


### GetTrailerParkplaces

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetTrailerParkplaces() int32`

GetTrailerParkplaces returns the TrailerParkplaces field if non-nil, zero value otherwise.

### GetTrailerParkplacesOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetTrailerParkplacesOk() (*int32, bool)`

GetTrailerParkplacesOk returns a tuple with the TrailerParkplaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerParkplaces

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetTrailerParkplaces(v int32)`

SetTrailerParkplaces sets TrailerParkplaces field to given value.


### GetBreakbulkCargoPlace

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetBreakbulkCargoPlace() int32`

GetBreakbulkCargoPlace returns the BreakbulkCargoPlace field if non-nil, zero value otherwise.

### GetBreakbulkCargoPlaceOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetBreakbulkCargoPlaceOk() (*int32, bool)`

GetBreakbulkCargoPlaceOk returns a tuple with the BreakbulkCargoPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakbulkCargoPlace

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetBreakbulkCargoPlace(v int32)`

SetBreakbulkCargoPlace sets BreakbulkCargoPlace field to given value.


### GetTruckParkplacesUsed

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetTruckParkplacesUsed() int32`

GetTruckParkplacesUsed returns the TruckParkplacesUsed field if non-nil, zero value otherwise.

### GetTruckParkplacesUsedOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetTruckParkplacesUsedOk() (*int32, bool)`

GetTruckParkplacesUsedOk returns a tuple with the TruckParkplacesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckParkplacesUsed

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetTruckParkplacesUsed(v int32)`

SetTruckParkplacesUsed sets TruckParkplacesUsed field to given value.


### GetMaintenancePlacesUsed

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetMaintenancePlacesUsed() int32`

GetMaintenancePlacesUsed returns the MaintenancePlacesUsed field if non-nil, zero value otherwise.

### GetMaintenancePlacesUsedOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetMaintenancePlacesUsedOk() (*int32, bool)`

GetMaintenancePlacesUsedOk returns a tuple with the MaintenancePlacesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenancePlacesUsed

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetMaintenancePlacesUsed(v int32)`

SetMaintenancePlacesUsed sets MaintenancePlacesUsed field to given value.


### GetSparePartsInStorage

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetSparePartsInStorage() []FPHSpedVAPIObjectsSpeditionsSparePartStock`

GetSparePartsInStorage returns the SparePartsInStorage field if non-nil, zero value otherwise.

### GetSparePartsInStorageOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetSparePartsInStorageOk() (*[]FPHSpedVAPIObjectsSpeditionsSparePartStock, bool)`

GetSparePartsInStorageOk returns a tuple with the SparePartsInStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparePartsInStorage

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetSparePartsInStorage(v []FPHSpedVAPIObjectsSpeditionsSparePartStock)`

SetSparePartsInStorage sets SparePartsInStorage field to given value.


### SetSparePartsInStorageNil

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetSparePartsInStorageNil(b bool)`

 SetSparePartsInStorageNil sets the value for SparePartsInStorage to be an explicit nil

### UnsetSparePartsInStorage
`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) UnsetSparePartsInStorage()`

UnsetSparePartsInStorage ensures that no value is present for SparePartsInStorage, not even an explicit nil
### GetTruckParkplacesNotFull

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetTruckParkplacesNotFull() bool`

GetTruckParkplacesNotFull returns the TruckParkplacesNotFull field if non-nil, zero value otherwise.

### GetTruckParkplacesNotFullOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetTruckParkplacesNotFullOk() (*bool, bool)`

GetTruckParkplacesNotFullOk returns a tuple with the TruckParkplacesNotFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckParkplacesNotFull

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetTruckParkplacesNotFull(v bool)`

SetTruckParkplacesNotFull sets TruckParkplacesNotFull field to given value.


### GetMaintenancePlacesNotFull

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetMaintenancePlacesNotFull() bool`

GetMaintenancePlacesNotFull returns the MaintenancePlacesNotFull field if non-nil, zero value otherwise.

### GetMaintenancePlacesNotFullOk

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) GetMaintenancePlacesNotFullOk() (*bool, bool)`

GetMaintenancePlacesNotFullOk returns a tuple with the MaintenancePlacesNotFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenancePlacesNotFull

`func (o *FPHSpedVAPIObjectsSpeditionsOwnedBranch) SetMaintenancePlacesNotFull(v bool)`

SetMaintenancePlacesNotFull sets MaintenancePlacesNotFull field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


