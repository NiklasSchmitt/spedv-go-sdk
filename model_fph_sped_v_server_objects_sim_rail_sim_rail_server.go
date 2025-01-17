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

// checks if the FPHSpedVServerObjectsSimRailSimRailServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVServerObjectsSimRailSimRailServer{}

// FPHSpedVServerObjectsSimRailSimRailServer struct for FPHSpedVServerObjectsSimRailSimRailServer
type FPHSpedVServerObjectsSimRailSimRailServer struct {
	ServerData NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo `json:"serverData"`
	LastOnlineState bool `json:"lastOnlineState"`
	CurrentTrainList []FPHSpedVServerObjectsSimRailSimRailActiveTrain `json:"currentTrainList"`
	DispatchStationList []FPHSpedVServerObjectsSimRailSimRailDispatchStation `json:"dispatchStationList"`
}

type _FPHSpedVServerObjectsSimRailSimRailServer FPHSpedVServerObjectsSimRailSimRailServer

// NewFPHSpedVServerObjectsSimRailSimRailServer instantiates a new FPHSpedVServerObjectsSimRailSimRailServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVServerObjectsSimRailSimRailServer(serverData NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo, lastOnlineState bool, currentTrainList []FPHSpedVServerObjectsSimRailSimRailActiveTrain, dispatchStationList []FPHSpedVServerObjectsSimRailSimRailDispatchStation) *FPHSpedVServerObjectsSimRailSimRailServer {
	this := FPHSpedVServerObjectsSimRailSimRailServer{}
	this.ServerData = serverData
	this.LastOnlineState = lastOnlineState
	this.CurrentTrainList = currentTrainList
	this.DispatchStationList = dispatchStationList
	return &this
}

// NewFPHSpedVServerObjectsSimRailSimRailServerWithDefaults instantiates a new FPHSpedVServerObjectsSimRailSimRailServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVServerObjectsSimRailSimRailServerWithDefaults() *FPHSpedVServerObjectsSimRailSimRailServer {
	this := FPHSpedVServerObjectsSimRailSimRailServer{}
	return &this
}

// GetServerData returns the ServerData field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsSimRailSimRailServerInfo will be returned
func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetServerData() FPHSpedVAPIObjectsSimRailSimRailServerInfo {
	if o == nil || o.ServerData.Get() == nil {
		var ret FPHSpedVAPIObjectsSimRailSimRailServerInfo
		return ret
	}

	return *o.ServerData.Get()
}

// GetServerDataOk returns a tuple with the ServerData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetServerDataOk() (*FPHSpedVAPIObjectsSimRailSimRailServerInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerData.Get(), o.ServerData.IsSet()
}

// SetServerData sets field value
func (o *FPHSpedVServerObjectsSimRailSimRailServer) SetServerData(v FPHSpedVAPIObjectsSimRailSimRailServerInfo) {
	o.ServerData.Set(&v)
}

// GetLastOnlineState returns the LastOnlineState field value
func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetLastOnlineState() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LastOnlineState
}

// GetLastOnlineStateOk returns a tuple with the LastOnlineState field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetLastOnlineStateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastOnlineState, true
}

// SetLastOnlineState sets field value
func (o *FPHSpedVServerObjectsSimRailSimRailServer) SetLastOnlineState(v bool) {
	o.LastOnlineState = v
}

// GetCurrentTrainList returns the CurrentTrainList field value
// If the value is explicit nil, the zero value for []FPHSpedVServerObjectsSimRailSimRailActiveTrain will be returned
func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetCurrentTrainList() []FPHSpedVServerObjectsSimRailSimRailActiveTrain {
	if o == nil {
		var ret []FPHSpedVServerObjectsSimRailSimRailActiveTrain
		return ret
	}

	return o.CurrentTrainList
}

// GetCurrentTrainListOk returns a tuple with the CurrentTrainList field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetCurrentTrainListOk() ([]FPHSpedVServerObjectsSimRailSimRailActiveTrain, bool) {
	if o == nil || IsNil(o.CurrentTrainList) {
		return nil, false
	}
	return o.CurrentTrainList, true
}

// SetCurrentTrainList sets field value
func (o *FPHSpedVServerObjectsSimRailSimRailServer) SetCurrentTrainList(v []FPHSpedVServerObjectsSimRailSimRailActiveTrain) {
	o.CurrentTrainList = v
}

// GetDispatchStationList returns the DispatchStationList field value
// If the value is explicit nil, the zero value for []FPHSpedVServerObjectsSimRailSimRailDispatchStation will be returned
func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetDispatchStationList() []FPHSpedVServerObjectsSimRailSimRailDispatchStation {
	if o == nil {
		var ret []FPHSpedVServerObjectsSimRailSimRailDispatchStation
		return ret
	}

	return o.DispatchStationList
}

// GetDispatchStationListOk returns a tuple with the DispatchStationList field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetDispatchStationListOk() ([]FPHSpedVServerObjectsSimRailSimRailDispatchStation, bool) {
	if o == nil || IsNil(o.DispatchStationList) {
		return nil, false
	}
	return o.DispatchStationList, true
}

// SetDispatchStationList sets field value
func (o *FPHSpedVServerObjectsSimRailSimRailServer) SetDispatchStationList(v []FPHSpedVServerObjectsSimRailSimRailDispatchStation) {
	o.DispatchStationList = v
}

func (o FPHSpedVServerObjectsSimRailSimRailServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVServerObjectsSimRailSimRailServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverData"] = o.ServerData.Get()
	toSerialize["lastOnlineState"] = o.LastOnlineState
	if o.CurrentTrainList != nil {
		toSerialize["currentTrainList"] = o.CurrentTrainList
	}
	if o.DispatchStationList != nil {
		toSerialize["dispatchStationList"] = o.DispatchStationList
	}
	return toSerialize, nil
}

func (o *FPHSpedVServerObjectsSimRailSimRailServer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serverData",
		"lastOnlineState",
		"currentTrainList",
		"dispatchStationList",
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

	varFPHSpedVServerObjectsSimRailSimRailServer := _FPHSpedVServerObjectsSimRailSimRailServer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVServerObjectsSimRailSimRailServer)

	if err != nil {
		return err
	}

	*o = FPHSpedVServerObjectsSimRailSimRailServer(varFPHSpedVServerObjectsSimRailSimRailServer)

	return err
}

type NullableFPHSpedVServerObjectsSimRailSimRailServer struct {
	value *FPHSpedVServerObjectsSimRailSimRailServer
	isSet bool
}

func (v NullableFPHSpedVServerObjectsSimRailSimRailServer) Get() *FPHSpedVServerObjectsSimRailSimRailServer {
	return v.value
}

func (v *NullableFPHSpedVServerObjectsSimRailSimRailServer) Set(val *FPHSpedVServerObjectsSimRailSimRailServer) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVServerObjectsSimRailSimRailServer) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVServerObjectsSimRailSimRailServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVServerObjectsSimRailSimRailServer(val *FPHSpedVServerObjectsSimRailSimRailServer) *NullableFPHSpedVServerObjectsSimRailSimRailServer {
	return &NullableFPHSpedVServerObjectsSimRailSimRailServer{value: val, isSet: true}
}

func (v NullableFPHSpedVServerObjectsSimRailSimRailServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVServerObjectsSimRailSimRailServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


