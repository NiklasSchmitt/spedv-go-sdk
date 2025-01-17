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

// checks if the FPHSpedVAPIObjectsSimRailSimRailDispatchStationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsSimRailSimRailDispatchStationData{}

// FPHSpedVAPIObjectsSimRailSimRailDispatchStationData struct for FPHSpedVAPIObjectsSimRailSimRailDispatchStationData
type FPHSpedVAPIObjectsSimRailSimRailDispatchStationData struct {
	Id NullableString `json:"id"`
	Name NullableString `json:"name"`
	Prefix NullableString `json:"prefix"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	DifficultyLevel int32 `json:"difficultyLevel"`
	MainImageURL NullableString `json:"mainImageURL"`
	AdditionalImage1URL NullableString `json:"additionalImage1URL"`
	AdditionalImage2URL NullableString `json:"additionalImage2URL"`
	Server NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo `json:"server"`
	SteamId NullableInt64 `json:"steamId"`
}

type _FPHSpedVAPIObjectsSimRailSimRailDispatchStationData FPHSpedVAPIObjectsSimRailSimRailDispatchStationData

// NewFPHSpedVAPIObjectsSimRailSimRailDispatchStationData instantiates a new FPHSpedVAPIObjectsSimRailSimRailDispatchStationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsSimRailSimRailDispatchStationData(id NullableString, name NullableString, prefix NullableString, latitude float64, longitude float64, difficultyLevel int32, mainImageURL NullableString, additionalImage1URL NullableString, additionalImage2URL NullableString, server NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo, steamId NullableInt64) *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData {
	this := FPHSpedVAPIObjectsSimRailSimRailDispatchStationData{}
	this.Id = id
	this.Name = name
	this.Prefix = prefix
	this.Latitude = latitude
	this.Longitude = longitude
	this.DifficultyLevel = difficultyLevel
	this.MainImageURL = mainImageURL
	this.AdditionalImage1URL = additionalImage1URL
	this.AdditionalImage2URL = additionalImage2URL
	this.Server = server
	this.SteamId = steamId
	return &this
}

// NewFPHSpedVAPIObjectsSimRailSimRailDispatchStationDataWithDefaults instantiates a new FPHSpedVAPIObjectsSimRailSimRailDispatchStationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsSimRailSimRailDispatchStationDataWithDefaults() *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData {
	this := FPHSpedVAPIObjectsSimRailSimRailDispatchStationData{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetId(v string) {
	o.Id.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetName(v string) {
	o.Name.Set(&v)
}

// GetPrefix returns the Prefix field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetPrefix() string {
	if o == nil || o.Prefix.Get() == nil {
		var ret string
		return ret
	}

	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// SetPrefix sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetPrefix(v string) {
	o.Prefix.Set(&v)
}

// GetLatitude returns the Latitude field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetLatitude(v float64) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetLongitude(v float64) {
	o.Longitude = v
}

// GetDifficultyLevel returns the DifficultyLevel field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetDifficultyLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DifficultyLevel
}

// GetDifficultyLevelOk returns a tuple with the DifficultyLevel field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetDifficultyLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DifficultyLevel, true
}

// SetDifficultyLevel sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetDifficultyLevel(v int32) {
	o.DifficultyLevel = v
}

// GetMainImageURL returns the MainImageURL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetMainImageURL() string {
	if o == nil || o.MainImageURL.Get() == nil {
		var ret string
		return ret
	}

	return *o.MainImageURL.Get()
}

// GetMainImageURLOk returns a tuple with the MainImageURL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetMainImageURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MainImageURL.Get(), o.MainImageURL.IsSet()
}

// SetMainImageURL sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetMainImageURL(v string) {
	o.MainImageURL.Set(&v)
}

// GetAdditionalImage1URL returns the AdditionalImage1URL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetAdditionalImage1URL() string {
	if o == nil || o.AdditionalImage1URL.Get() == nil {
		var ret string
		return ret
	}

	return *o.AdditionalImage1URL.Get()
}

// GetAdditionalImage1URLOk returns a tuple with the AdditionalImage1URL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetAdditionalImage1URLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalImage1URL.Get(), o.AdditionalImage1URL.IsSet()
}

// SetAdditionalImage1URL sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetAdditionalImage1URL(v string) {
	o.AdditionalImage1URL.Set(&v)
}

// GetAdditionalImage2URL returns the AdditionalImage2URL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetAdditionalImage2URL() string {
	if o == nil || o.AdditionalImage2URL.Get() == nil {
		var ret string
		return ret
	}

	return *o.AdditionalImage2URL.Get()
}

// GetAdditionalImage2URLOk returns a tuple with the AdditionalImage2URL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetAdditionalImage2URLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalImage2URL.Get(), o.AdditionalImage2URL.IsSet()
}

// SetAdditionalImage2URL sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetAdditionalImage2URL(v string) {
	o.AdditionalImage2URL.Set(&v)
}

// GetServer returns the Server field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsSimRailSimRailServerInfo will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetServer() FPHSpedVAPIObjectsSimRailSimRailServerInfo {
	if o == nil || o.Server.Get() == nil {
		var ret FPHSpedVAPIObjectsSimRailSimRailServerInfo
		return ret
	}

	return *o.Server.Get()
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetServerOk() (*FPHSpedVAPIObjectsSimRailSimRailServerInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Server.Get(), o.Server.IsSet()
}

// SetServer sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetServer(v FPHSpedVAPIObjectsSimRailSimRailServerInfo) {
	o.Server.Set(&v)
}

// GetSteamId returns the SteamId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetSteamId() int64 {
	if o == nil || o.SteamId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.SteamId.Get()
}

// GetSteamIdOk returns a tuple with the SteamId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) GetSteamIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SteamId.Get(), o.SteamId.IsSet()
}

// SetSteamId sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) SetSteamId(v int64) {
	o.SteamId.Set(&v)
}

func (o FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["name"] = o.Name.Get()
	toSerialize["prefix"] = o.Prefix.Get()
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	toSerialize["difficultyLevel"] = o.DifficultyLevel
	toSerialize["mainImageURL"] = o.MainImageURL.Get()
	toSerialize["additionalImage1URL"] = o.AdditionalImage1URL.Get()
	toSerialize["additionalImage2URL"] = o.AdditionalImage2URL.Get()
	toSerialize["server"] = o.Server.Get()
	toSerialize["steamId"] = o.SteamId.Get()
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"prefix",
		"latitude",
		"longitude",
		"difficultyLevel",
		"mainImageURL",
		"additionalImage1URL",
		"additionalImage2URL",
		"server",
		"steamId",
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

	varFPHSpedVAPIObjectsSimRailSimRailDispatchStationData := _FPHSpedVAPIObjectsSimRailSimRailDispatchStationData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsSimRailSimRailDispatchStationData)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsSimRailSimRailDispatchStationData(varFPHSpedVAPIObjectsSimRailSimRailDispatchStationData)

	return err
}

type NullableFPHSpedVAPIObjectsSimRailSimRailDispatchStationData struct {
	value *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailDispatchStationData) Get() *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailDispatchStationData) Set(val *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailDispatchStationData) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailDispatchStationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSimRailSimRailDispatchStationData(val *FPHSpedVAPIObjectsSimRailSimRailDispatchStationData) *NullableFPHSpedVAPIObjectsSimRailSimRailDispatchStationData {
	return &NullableFPHSpedVAPIObjectsSimRailSimRailDispatchStationData{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailDispatchStationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailDispatchStationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


