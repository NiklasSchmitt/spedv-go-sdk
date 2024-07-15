# FPHSpedVAPIObjectsKontorLightTrailer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Spedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**Game** | [**FPHSpedVAPIEnumsGameEnum**](FPHSpedVAPIEnumsGameEnum.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | [readonly] 
**LicensePlate** | **NullableString** |  | [readonly] 
**Comment** | **NullableString** |  | [readonly] 
**Category** | **NullableString** |  | [readonly] 
**Skin** | **NullableString** |  | [readonly] 
**TrailerType** | [**FPHSpedVAPIEnumsKontorTrailerType**](FPHSpedVAPIEnumsKontorTrailerType.md) |   0 &#x3D; StandardTrailer  1 &#x3D; DoubleTrailer  2 &#x3D; BDoubleTrailer  3 &#x3D; TripleTrailer  4 &#x3D; ShortTrailer  -1 &#x3D; NotSet | [readonly] 
**DefaultDriver** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**AssignedScheduler** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**TrailerAvailable** | **time.Time** |  | [readonly] 
**TrailerHasNeededMaintenance** | **bool** |  | [readonly] 
**Km** | **int32** |  | [readonly] 
**MaxWeight** | **int32** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsKontorLightTrailer

`func NewFPHSpedVAPIObjectsKontorLightTrailer(id int32, spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, game FPHSpedVAPIEnumsGameEnum, licensePlate NullableString, comment NullableString, category NullableString, skin NullableString, trailerType FPHSpedVAPIEnumsKontorTrailerType, defaultDriver NullableFPHSpedVAPIObjectsUsersUser, assignedScheduler NullableFPHSpedVAPIObjectsUsersUser, trailerAvailable time.Time, trailerHasNeededMaintenance bool, km int32, maxWeight int32, ) *FPHSpedVAPIObjectsKontorLightTrailer`

NewFPHSpedVAPIObjectsKontorLightTrailer instantiates a new FPHSpedVAPIObjectsKontorLightTrailer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsKontorLightTrailerWithDefaults

`func NewFPHSpedVAPIObjectsKontorLightTrailerWithDefaults() *FPHSpedVAPIObjectsKontorLightTrailer`

NewFPHSpedVAPIObjectsKontorLightTrailerWithDefaults instantiates a new FPHSpedVAPIObjectsKontorLightTrailer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetId(v int32)`

SetId sets Id field to given value.


### GetSpedition

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetSpedition returns the Spedition field if non-nil, zero value otherwise.

### GetSpeditionOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetSpeditionOk returns a tuple with the Spedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpedition

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetSpedition sets Spedition field to given value.


### SetSpeditionNil

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetSpeditionNil(b bool)`

 SetSpeditionNil sets the value for Spedition to be an explicit nil

### UnsetSpedition
`func (o *FPHSpedVAPIObjectsKontorLightTrailer) UnsetSpedition()`

UnsetSpedition ensures that no value is present for Spedition, not even an explicit nil
### GetGame

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetGame() FPHSpedVAPIEnumsGameEnum`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetGameOk() (*FPHSpedVAPIEnumsGameEnum, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetGame(v FPHSpedVAPIEnumsGameEnum)`

SetGame sets Game field to given value.


### GetLicensePlate

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetLicensePlate() string`

GetLicensePlate returns the LicensePlate field if non-nil, zero value otherwise.

### GetLicensePlateOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetLicensePlateOk() (*string, bool)`

GetLicensePlateOk returns a tuple with the LicensePlate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlate

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetLicensePlate(v string)`

SetLicensePlate sets LicensePlate field to given value.


### SetLicensePlateNil

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetLicensePlateNil(b bool)`

 SetLicensePlateNil sets the value for LicensePlate to be an explicit nil

### UnsetLicensePlate
`func (o *FPHSpedVAPIObjectsKontorLightTrailer) UnsetLicensePlate()`

UnsetLicensePlate ensures that no value is present for LicensePlate, not even an explicit nil
### GetComment

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FPHSpedVAPIObjectsKontorLightTrailer) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCategory

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetCategory(v string)`

SetCategory sets Category field to given value.


### SetCategoryNil

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *FPHSpedVAPIObjectsKontorLightTrailer) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetSkin

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetSkin() string`

GetSkin returns the Skin field if non-nil, zero value otherwise.

### GetSkinOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetSkinOk() (*string, bool)`

GetSkinOk returns a tuple with the Skin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkin

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetSkin(v string)`

SetSkin sets Skin field to given value.


### SetSkinNil

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetSkinNil(b bool)`

 SetSkinNil sets the value for Skin to be an explicit nil

### UnsetSkin
`func (o *FPHSpedVAPIObjectsKontorLightTrailer) UnsetSkin()`

UnsetSkin ensures that no value is present for Skin, not even an explicit nil
### GetTrailerType

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetTrailerType() FPHSpedVAPIEnumsKontorTrailerType`

GetTrailerType returns the TrailerType field if non-nil, zero value otherwise.

### GetTrailerTypeOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetTrailerTypeOk() (*FPHSpedVAPIEnumsKontorTrailerType, bool)`

GetTrailerTypeOk returns a tuple with the TrailerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerType

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetTrailerType(v FPHSpedVAPIEnumsKontorTrailerType)`

SetTrailerType sets TrailerType field to given value.


### GetDefaultDriver

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetDefaultDriver() FPHSpedVAPIObjectsUsersUser`

GetDefaultDriver returns the DefaultDriver field if non-nil, zero value otherwise.

### GetDefaultDriverOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetDefaultDriverOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetDefaultDriverOk returns a tuple with the DefaultDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDriver

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetDefaultDriver(v FPHSpedVAPIObjectsUsersUser)`

SetDefaultDriver sets DefaultDriver field to given value.


### SetDefaultDriverNil

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetDefaultDriverNil(b bool)`

 SetDefaultDriverNil sets the value for DefaultDriver to be an explicit nil

### UnsetDefaultDriver
`func (o *FPHSpedVAPIObjectsKontorLightTrailer) UnsetDefaultDriver()`

UnsetDefaultDriver ensures that no value is present for DefaultDriver, not even an explicit nil
### GetAssignedScheduler

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetAssignedScheduler() FPHSpedVAPIObjectsUsersUser`

GetAssignedScheduler returns the AssignedScheduler field if non-nil, zero value otherwise.

### GetAssignedSchedulerOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetAssignedSchedulerOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetAssignedSchedulerOk returns a tuple with the AssignedScheduler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedScheduler

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetAssignedScheduler(v FPHSpedVAPIObjectsUsersUser)`

SetAssignedScheduler sets AssignedScheduler field to given value.


### SetAssignedSchedulerNil

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetAssignedSchedulerNil(b bool)`

 SetAssignedSchedulerNil sets the value for AssignedScheduler to be an explicit nil

### UnsetAssignedScheduler
`func (o *FPHSpedVAPIObjectsKontorLightTrailer) UnsetAssignedScheduler()`

UnsetAssignedScheduler ensures that no value is present for AssignedScheduler, not even an explicit nil
### GetTrailerAvailable

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetTrailerAvailable() time.Time`

GetTrailerAvailable returns the TrailerAvailable field if non-nil, zero value otherwise.

### GetTrailerAvailableOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetTrailerAvailableOk() (*time.Time, bool)`

GetTrailerAvailableOk returns a tuple with the TrailerAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerAvailable

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetTrailerAvailable(v time.Time)`

SetTrailerAvailable sets TrailerAvailable field to given value.


### GetTrailerHasNeededMaintenance

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetTrailerHasNeededMaintenance() bool`

GetTrailerHasNeededMaintenance returns the TrailerHasNeededMaintenance field if non-nil, zero value otherwise.

### GetTrailerHasNeededMaintenanceOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetTrailerHasNeededMaintenanceOk() (*bool, bool)`

GetTrailerHasNeededMaintenanceOk returns a tuple with the TrailerHasNeededMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerHasNeededMaintenance

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetTrailerHasNeededMaintenance(v bool)`

SetTrailerHasNeededMaintenance sets TrailerHasNeededMaintenance field to given value.


### GetKm

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetKm() int32`

GetKm returns the Km field if non-nil, zero value otherwise.

### GetKmOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetKmOk() (*int32, bool)`

GetKmOk returns a tuple with the Km field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKm

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetKm(v int32)`

SetKm sets Km field to given value.


### GetMaxWeight

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetMaxWeight() int32`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) GetMaxWeightOk() (*int32, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *FPHSpedVAPIObjectsKontorLightTrailer) SetMaxWeight(v int32)`

SetMaxWeight sets MaxWeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


