# BingMapsRESTToolkitLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Point** | [**NullableBingMapsRESTToolkitPoint**](BingMapsRESTToolkitPoint.md) |  | 
**EntityType** | **NullableString** |  | 
**Address** | [**NullableBingMapsRESTToolkitAddress**](BingMapsRESTToolkitAddress.md) |  | 
**Confidence** | **NullableString** |  | 
**MatchCodes** | **[]string** |  | 
**GeocodePoints** | [**[]BingMapsRESTToolkitPoint**](BingMapsRESTToolkitPoint.md) |  | 
**QueryParseValues** | [**[]BingMapsRESTToolkitQueryParseValue**](BingMapsRESTToolkitQueryParseValue.md) |  | 
**Bbox** | Pointer to **[]float64** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBingMapsRESTToolkitLocation

`func NewBingMapsRESTToolkitLocation(name NullableString, point NullableBingMapsRESTToolkitPoint, entityType NullableString, address NullableBingMapsRESTToolkitAddress, confidence NullableString, matchCodes []string, geocodePoints []BingMapsRESTToolkitPoint, queryParseValues []BingMapsRESTToolkitQueryParseValue, ) *BingMapsRESTToolkitLocation`

NewBingMapsRESTToolkitLocation instantiates a new BingMapsRESTToolkitLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBingMapsRESTToolkitLocationWithDefaults

`func NewBingMapsRESTToolkitLocationWithDefaults() *BingMapsRESTToolkitLocation`

NewBingMapsRESTToolkitLocationWithDefaults instantiates a new BingMapsRESTToolkitLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BingMapsRESTToolkitLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BingMapsRESTToolkitLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BingMapsRESTToolkitLocation) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *BingMapsRESTToolkitLocation) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BingMapsRESTToolkitLocation) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPoint

`func (o *BingMapsRESTToolkitLocation) GetPoint() BingMapsRESTToolkitPoint`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *BingMapsRESTToolkitLocation) GetPointOk() (*BingMapsRESTToolkitPoint, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *BingMapsRESTToolkitLocation) SetPoint(v BingMapsRESTToolkitPoint)`

SetPoint sets Point field to given value.


### SetPointNil

`func (o *BingMapsRESTToolkitLocation) SetPointNil(b bool)`

 SetPointNil sets the value for Point to be an explicit nil

### UnsetPoint
`func (o *BingMapsRESTToolkitLocation) UnsetPoint()`

UnsetPoint ensures that no value is present for Point, not even an explicit nil
### GetEntityType

`func (o *BingMapsRESTToolkitLocation) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *BingMapsRESTToolkitLocation) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *BingMapsRESTToolkitLocation) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### SetEntityTypeNil

`func (o *BingMapsRESTToolkitLocation) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *BingMapsRESTToolkitLocation) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetAddress

`func (o *BingMapsRESTToolkitLocation) GetAddress() BingMapsRESTToolkitAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BingMapsRESTToolkitLocation) GetAddressOk() (*BingMapsRESTToolkitAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BingMapsRESTToolkitLocation) SetAddress(v BingMapsRESTToolkitAddress)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *BingMapsRESTToolkitLocation) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *BingMapsRESTToolkitLocation) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetConfidence

`func (o *BingMapsRESTToolkitLocation) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *BingMapsRESTToolkitLocation) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *BingMapsRESTToolkitLocation) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.


### SetConfidenceNil

`func (o *BingMapsRESTToolkitLocation) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *BingMapsRESTToolkitLocation) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetMatchCodes

`func (o *BingMapsRESTToolkitLocation) GetMatchCodes() []string`

GetMatchCodes returns the MatchCodes field if non-nil, zero value otherwise.

### GetMatchCodesOk

`func (o *BingMapsRESTToolkitLocation) GetMatchCodesOk() (*[]string, bool)`

GetMatchCodesOk returns a tuple with the MatchCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCodes

`func (o *BingMapsRESTToolkitLocation) SetMatchCodes(v []string)`

SetMatchCodes sets MatchCodes field to given value.


### SetMatchCodesNil

`func (o *BingMapsRESTToolkitLocation) SetMatchCodesNil(b bool)`

 SetMatchCodesNil sets the value for MatchCodes to be an explicit nil

### UnsetMatchCodes
`func (o *BingMapsRESTToolkitLocation) UnsetMatchCodes()`

UnsetMatchCodes ensures that no value is present for MatchCodes, not even an explicit nil
### GetGeocodePoints

`func (o *BingMapsRESTToolkitLocation) GetGeocodePoints() []BingMapsRESTToolkitPoint`

GetGeocodePoints returns the GeocodePoints field if non-nil, zero value otherwise.

### GetGeocodePointsOk

`func (o *BingMapsRESTToolkitLocation) GetGeocodePointsOk() (*[]BingMapsRESTToolkitPoint, bool)`

GetGeocodePointsOk returns a tuple with the GeocodePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocodePoints

`func (o *BingMapsRESTToolkitLocation) SetGeocodePoints(v []BingMapsRESTToolkitPoint)`

SetGeocodePoints sets GeocodePoints field to given value.


### SetGeocodePointsNil

`func (o *BingMapsRESTToolkitLocation) SetGeocodePointsNil(b bool)`

 SetGeocodePointsNil sets the value for GeocodePoints to be an explicit nil

### UnsetGeocodePoints
`func (o *BingMapsRESTToolkitLocation) UnsetGeocodePoints()`

UnsetGeocodePoints ensures that no value is present for GeocodePoints, not even an explicit nil
### GetQueryParseValues

`func (o *BingMapsRESTToolkitLocation) GetQueryParseValues() []BingMapsRESTToolkitQueryParseValue`

GetQueryParseValues returns the QueryParseValues field if non-nil, zero value otherwise.

### GetQueryParseValuesOk

`func (o *BingMapsRESTToolkitLocation) GetQueryParseValuesOk() (*[]BingMapsRESTToolkitQueryParseValue, bool)`

GetQueryParseValuesOk returns a tuple with the QueryParseValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParseValues

`func (o *BingMapsRESTToolkitLocation) SetQueryParseValues(v []BingMapsRESTToolkitQueryParseValue)`

SetQueryParseValues sets QueryParseValues field to given value.


### SetQueryParseValuesNil

`func (o *BingMapsRESTToolkitLocation) SetQueryParseValuesNil(b bool)`

 SetQueryParseValuesNil sets the value for QueryParseValues to be an explicit nil

### UnsetQueryParseValues
`func (o *BingMapsRESTToolkitLocation) UnsetQueryParseValues()`

UnsetQueryParseValues ensures that no value is present for QueryParseValues, not even an explicit nil
### GetBbox

`func (o *BingMapsRESTToolkitLocation) GetBbox() []float64`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *BingMapsRESTToolkitLocation) GetBboxOk() (*[]float64, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *BingMapsRESTToolkitLocation) SetBbox(v []float64)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *BingMapsRESTToolkitLocation) HasBbox() bool`

HasBbox returns a boolean if a field has been set.

### SetBboxNil

`func (o *BingMapsRESTToolkitLocation) SetBboxNil(b bool)`

 SetBboxNil sets the value for Bbox to be an explicit nil

### UnsetBbox
`func (o *BingMapsRESTToolkitLocation) UnsetBbox()`

UnsetBbox ensures that no value is present for Bbox, not even an explicit nil
### GetType

`func (o *BingMapsRESTToolkitLocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BingMapsRESTToolkitLocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BingMapsRESTToolkitLocation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BingMapsRESTToolkitLocation) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BingMapsRESTToolkitLocation) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BingMapsRESTToolkitLocation) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


