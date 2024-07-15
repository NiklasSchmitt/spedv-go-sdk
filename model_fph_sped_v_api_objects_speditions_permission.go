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

// checks if the FPHSpedVAPIObjectsSpeditionsPermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsSpeditionsPermission{}

// FPHSpedVAPIObjectsSpeditionsPermission struct for FPHSpedVAPIObjectsSpeditionsPermission
type FPHSpedVAPIObjectsSpeditionsPermission struct {
	Id int32 `json:"id"`
	Key NullableString `json:"key"`
	DefaultValue bool `json:"defaultValue"`
	DefaultValueLeader bool `json:"defaultValueLeader"`
}

type _FPHSpedVAPIObjectsSpeditionsPermission FPHSpedVAPIObjectsSpeditionsPermission

// NewFPHSpedVAPIObjectsSpeditionsPermission instantiates a new FPHSpedVAPIObjectsSpeditionsPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsSpeditionsPermission(id int32, key NullableString, defaultValue bool, defaultValueLeader bool) *FPHSpedVAPIObjectsSpeditionsPermission {
	this := FPHSpedVAPIObjectsSpeditionsPermission{}
	this.Id = id
	this.Key = key
	this.DefaultValue = defaultValue
	this.DefaultValueLeader = defaultValueLeader
	return &this
}

// NewFPHSpedVAPIObjectsSpeditionsPermissionWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsSpeditionsPermissionWithDefaults() *FPHSpedVAPIObjectsSpeditionsPermission {
	this := FPHSpedVAPIObjectsSpeditionsPermission{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsSpeditionsPermission) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsPermission) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPermission) SetId(v int32) {
	o.Id = v
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSpeditionsPermission) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}

	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsPermission) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPermission) SetKey(v string) {
	o.Key.Set(&v)
}

// GetDefaultValue returns the DefaultValue field value
func (o *FPHSpedVAPIObjectsSpeditionsPermission) GetDefaultValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsPermission) GetDefaultValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// SetDefaultValue sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPermission) SetDefaultValue(v bool) {
	o.DefaultValue = v
}

// GetDefaultValueLeader returns the DefaultValueLeader field value
func (o *FPHSpedVAPIObjectsSpeditionsPermission) GetDefaultValueLeader() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultValueLeader
}

// GetDefaultValueLeaderOk returns a tuple with the DefaultValueLeader field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsPermission) GetDefaultValueLeaderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultValueLeader, true
}

// SetDefaultValueLeader sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPermission) SetDefaultValueLeader(v bool) {
	o.DefaultValueLeader = v
}

func (o FPHSpedVAPIObjectsSpeditionsPermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsSpeditionsPermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key.Get()
	toSerialize["defaultValue"] = o.DefaultValue
	toSerialize["defaultValueLeader"] = o.DefaultValueLeader
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsSpeditionsPermission) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"key",
		"defaultValue",
		"defaultValueLeader",
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

	varFPHSpedVAPIObjectsSpeditionsPermission := _FPHSpedVAPIObjectsSpeditionsPermission{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsSpeditionsPermission)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsSpeditionsPermission(varFPHSpedVAPIObjectsSpeditionsPermission)

	return err
}

type NullableFPHSpedVAPIObjectsSpeditionsPermission struct {
	value *FPHSpedVAPIObjectsSpeditionsPermission
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSpeditionsPermission) Get() *FPHSpedVAPIObjectsSpeditionsPermission {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsPermission) Set(val *FPHSpedVAPIObjectsSpeditionsPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSpeditionsPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSpeditionsPermission(val *FPHSpedVAPIObjectsSpeditionsPermission) *NullableFPHSpedVAPIObjectsSpeditionsPermission {
	return &NullableFPHSpedVAPIObjectsSpeditionsPermission{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSpeditionsPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


