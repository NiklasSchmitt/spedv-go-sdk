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

// checks if the FPHSpedVAPIObjectsSimRailSimRailServerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsSimRailSimRailServerInfo{}

// FPHSpedVAPIObjectsSimRailSimRailServerInfo struct for FPHSpedVAPIObjectsSimRailSimRailServerInfo
type FPHSpedVAPIObjectsSimRailSimRailServerInfo struct {
	ServerCode NullableString `json:"serverCode"`
	ServerName NullableString `json:"serverName"`
	ServerRegion NullableString `json:"serverRegion"`
	UtcOffset int32 `json:"utcOffset"`
}

type _FPHSpedVAPIObjectsSimRailSimRailServerInfo FPHSpedVAPIObjectsSimRailSimRailServerInfo

// NewFPHSpedVAPIObjectsSimRailSimRailServerInfo instantiates a new FPHSpedVAPIObjectsSimRailSimRailServerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsSimRailSimRailServerInfo(serverCode NullableString, serverName NullableString, serverRegion NullableString, utcOffset int32) *FPHSpedVAPIObjectsSimRailSimRailServerInfo {
	this := FPHSpedVAPIObjectsSimRailSimRailServerInfo{}
	this.ServerCode = serverCode
	this.ServerName = serverName
	this.ServerRegion = serverRegion
	this.UtcOffset = utcOffset
	return &this
}

// NewFPHSpedVAPIObjectsSimRailSimRailServerInfoWithDefaults instantiates a new FPHSpedVAPIObjectsSimRailSimRailServerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsSimRailSimRailServerInfoWithDefaults() *FPHSpedVAPIObjectsSimRailSimRailServerInfo {
	this := FPHSpedVAPIObjectsSimRailSimRailServerInfo{}
	return &this
}

// GetServerCode returns the ServerCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) GetServerCode() string {
	if o == nil || o.ServerCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.ServerCode.Get()
}

// GetServerCodeOk returns a tuple with the ServerCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) GetServerCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerCode.Get(), o.ServerCode.IsSet()
}

// SetServerCode sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) SetServerCode(v string) {
	o.ServerCode.Set(&v)
}

// GetServerName returns the ServerName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) GetServerName() string {
	if o == nil || o.ServerName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ServerName.Get()
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerName.Get(), o.ServerName.IsSet()
}

// SetServerName sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) SetServerName(v string) {
	o.ServerName.Set(&v)
}

// GetServerRegion returns the ServerRegion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) GetServerRegion() string {
	if o == nil || o.ServerRegion.Get() == nil {
		var ret string
		return ret
	}

	return *o.ServerRegion.Get()
}

// GetServerRegionOk returns a tuple with the ServerRegion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) GetServerRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerRegion.Get(), o.ServerRegion.IsSet()
}

// SetServerRegion sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) SetServerRegion(v string) {
	o.ServerRegion.Set(&v)
}

// GetUtcOffset returns the UtcOffset field value
func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) GetUtcOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UtcOffset
}

// GetUtcOffsetOk returns a tuple with the UtcOffset field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) GetUtcOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UtcOffset, true
}

// SetUtcOffset sets field value
func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) SetUtcOffset(v int32) {
	o.UtcOffset = v
}

func (o FPHSpedVAPIObjectsSimRailSimRailServerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsSimRailSimRailServerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverCode"] = o.ServerCode.Get()
	toSerialize["serverName"] = o.ServerName.Get()
	toSerialize["serverRegion"] = o.ServerRegion.Get()
	toSerialize["utcOffset"] = o.UtcOffset
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsSimRailSimRailServerInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serverCode",
		"serverName",
		"serverRegion",
		"utcOffset",
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

	varFPHSpedVAPIObjectsSimRailSimRailServerInfo := _FPHSpedVAPIObjectsSimRailSimRailServerInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsSimRailSimRailServerInfo)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsSimRailSimRailServerInfo(varFPHSpedVAPIObjectsSimRailSimRailServerInfo)

	return err
}

type NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo struct {
	value *FPHSpedVAPIObjectsSimRailSimRailServerInfo
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo) Get() *FPHSpedVAPIObjectsSimRailSimRailServerInfo {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo) Set(val *FPHSpedVAPIObjectsSimRailSimRailServerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSimRailSimRailServerInfo(val *FPHSpedVAPIObjectsSimRailSimRailServerInfo) *NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo {
	return &NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

