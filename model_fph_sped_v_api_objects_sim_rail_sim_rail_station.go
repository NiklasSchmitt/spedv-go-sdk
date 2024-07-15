/*
SpedV API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FPHSpedVAPIObjectsSimRailSimRailStation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsSimRailSimRailStation{}

// FPHSpedVAPIObjectsSimRailSimRailStation struct for FPHSpedVAPIObjectsSimRailSimRailStation
type FPHSpedVAPIObjectsSimRailSimRailStation struct {
	Id NullableString `json:"id"`
	Name NullableString `json:"name"`
	SupervisedBy NullableString `json:"supervisedBy"`
	StationCategory NullableString `json:"stationCategory"`
	Prefix NullableString `json:"prefix"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	DifficultyLevel int32 `json:"difficultyLevel"`
	MainImageURL NullableString `json:"mainImageURL"`
	AdditionalImage1URL NullableString `json:"additionalImage1URL"`
	AdditionalImage2URL NullableString `json:"additionalImage2URL"`
}

type _FPHSpedVAPIObjectsSimRailSimRailStation FPHSpedVAPIObjectsSimRailSimRailStation

// NewFPHSpedVAPIObjectsSimRailSimRailStation instantiates a new FPHSpedVAPIObjectsSimRailSimRailStation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsSimRailSimRailStation(id NullableString, name NullableString, supervisedBy NullableString, stationCategory NullableString, prefix NullableString, latitude float64, longitude float64, difficultyLevel int32, mainImageURL NullableString, additionalImage1URL NullableString, additionalImage2URL NullableString) *FPHSpedVAPIObjectsSimRailSimRailStation {
	this := FPHSpedVAPIObjectsSimRailSimRailStation{}
	this.Id = id
	this.Name = name
	this.SupervisedBy = supervisedBy
	this.StationCategory = stationCategory
	this.Prefix = prefix
	this.Latitude = latitude
	this.Longitude = longitude
	this.DifficultyLevel = difficultyLevel
	this.MainImageURL = mainImageURL
	this.AdditionalImage1URL = additionalImage1URL
	this.AdditionalImage2URL = additionalImage2URL
	return &this
}

// NewFPHSpedVAPIObjectsSimRailSimRailStationWithDefaults instantiates a new FPHSpedVAPIObjectsSimRailSimRailStation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsSimRailSimRailStationWithDefaults() *FPHSpedVAPIObjectsSimRailSimRailStation {
	this := FPHSpedVAPIObjectsSimRailSimRailStation{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) SetId(v string) {
	o.Id.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) SetName(v string) {
	o.Name.Set(&v)
}

// GetSupervisedBy returns the SupervisedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetSupervisedBy() string {
	if o == nil || o.SupervisedBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.SupervisedBy.Get()
}

// GetSupervisedByOk returns a tuple with the SupervisedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetSupervisedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupervisedBy.Get(), o.SupervisedBy.IsSet()
}

// SetSupervisedBy sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) SetSupervisedBy(v string) {
	o.SupervisedBy.Set(&v)
}

// GetStationCategory returns the StationCategory field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetStationCategory() string {
	if o == nil || o.StationCategory.Get() == nil {
		var ret string
		return ret
	}

	return *o.StationCategory.Get()
}

// GetStationCategoryOk returns a tuple with the StationCategory field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetStationCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StationCategory.Get(), o.StationCategory.IsSet()
}

// SetStationCategory sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) SetStationCategory(v string) {
	o.StationCategory.Set(&v)
}

// GetPrefix returns the Prefix field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetPrefix() string {
	if o == nil || o.Prefix.Get() == nil {
		var ret string
		return ret
	}

	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// SetPrefix sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) SetPrefix(v string) {
	o.Prefix.Set(&v)
}

// GetLatitude returns the Latitude field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) SetLatitude(v float64) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) SetLongitude(v float64) {
	o.Longitude = v
}

// GetDifficultyLevel returns the DifficultyLevel field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetDifficultyLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DifficultyLevel
}

// GetDifficultyLevelOk returns a tuple with the DifficultyLevel field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetDifficultyLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DifficultyLevel, true
}

// SetDifficultyLevel sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) SetDifficultyLevel(v int32) {
	o.DifficultyLevel = v
}

// GetMainImageURL returns the MainImageURL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetMainImageURL() string {
	if o == nil || o.MainImageURL.Get() == nil {
		var ret string
		return ret
	}

	return *o.MainImageURL.Get()
}

// GetMainImageURLOk returns a tuple with the MainImageURL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetMainImageURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MainImageURL.Get(), o.MainImageURL.IsSet()
}

// SetMainImageURL sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) SetMainImageURL(v string) {
	o.MainImageURL.Set(&v)
}

// GetAdditionalImage1URL returns the AdditionalImage1URL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetAdditionalImage1URL() string {
	if o == nil || o.AdditionalImage1URL.Get() == nil {
		var ret string
		return ret
	}

	return *o.AdditionalImage1URL.Get()
}

// GetAdditionalImage1URLOk returns a tuple with the AdditionalImage1URL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetAdditionalImage1URLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalImage1URL.Get(), o.AdditionalImage1URL.IsSet()
}

// SetAdditionalImage1URL sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) SetAdditionalImage1URL(v string) {
	o.AdditionalImage1URL.Set(&v)
}

// GetAdditionalImage2URL returns the AdditionalImage2URL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetAdditionalImage2URL() string {
	if o == nil || o.AdditionalImage2URL.Get() == nil {
		var ret string
		return ret
	}

	return *o.AdditionalImage2URL.Get()
}

// GetAdditionalImage2URLOk returns a tuple with the AdditionalImage2URL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) GetAdditionalImage2URLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalImage2URL.Get(), o.AdditionalImage2URL.IsSet()
}

// SetAdditionalImage2URL sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailStation) SetAdditionalImage2URL(v string) {
	o.AdditionalImage2URL.Set(&v)
}

func (o FPHSpedVAPIObjectsSimRailSimRailStation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsSimRailSimRailStation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["name"] = o.Name.Get()
	toSerialize["supervisedBy"] = o.SupervisedBy.Get()
	toSerialize["stationCategory"] = o.StationCategory.Get()
	toSerialize["prefix"] = o.Prefix.Get()
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	toSerialize["difficultyLevel"] = o.DifficultyLevel
	toSerialize["mainImageURL"] = o.MainImageURL.Get()
	toSerialize["additionalImage1URL"] = o.AdditionalImage1URL.Get()
	toSerialize["additionalImage2URL"] = o.AdditionalImage2URL.Get()
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsSimRailSimRailStation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"supervisedBy",
		"stationCategory",
		"prefix",
		"latitude",
		"longitude",
		"difficultyLevel",
		"mainImageURL",
		"additionalImage1URL",
		"additionalImage2URL",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFPHSpedVAPIObjectsSimRailSimRailStation := _FPHSpedVAPIObjectsSimRailSimRailStation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsSimRailSimRailStation)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsSimRailSimRailStation(varFPHSpedVAPIObjectsSimRailSimRailStation)

	return err
}

type NullableFPHSpedVAPIObjectsSimRailSimRailStation struct {
	value *FPHSpedVAPIObjectsSimRailSimRailStation
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailStation) Get() *FPHSpedVAPIObjectsSimRailSimRailStation {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailStation) Set(val *FPHSpedVAPIObjectsSimRailSimRailStation) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailStation) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailStation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSimRailSimRailStation(val *FPHSpedVAPIObjectsSimRailSimRailStation) *NullableFPHSpedVAPIObjectsSimRailSimRailStation {
	return &NullableFPHSpedVAPIObjectsSimRailSimRailStation{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailStation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailStation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


