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

// checks if the FPHSpedVAPIObjectsMapsHarbor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsMapsHarbor{}

// FPHSpedVAPIObjectsMapsHarbor struct for FPHSpedVAPIObjectsMapsHarbor
type FPHSpedVAPIObjectsMapsHarbor struct {
	Id int32 `json:"id"`
	X float64 `json:"x"`
	Z float64 `json:"z"`
	Name NullableString `json:"name"`
}

type _FPHSpedVAPIObjectsMapsHarbor FPHSpedVAPIObjectsMapsHarbor

// NewFPHSpedVAPIObjectsMapsHarbor instantiates a new FPHSpedVAPIObjectsMapsHarbor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsMapsHarbor(id int32, x float64, z float64, name NullableString) *FPHSpedVAPIObjectsMapsHarbor {
	this := FPHSpedVAPIObjectsMapsHarbor{}
	this.Id = id
	this.X = x
	this.Z = z
	this.Name = name
	return &this
}

// NewFPHSpedVAPIObjectsMapsHarborWithDefaults instantiates a new FPHSpedVAPIObjectsMapsHarbor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsMapsHarborWithDefaults() *FPHSpedVAPIObjectsMapsHarbor {
	this := FPHSpedVAPIObjectsMapsHarbor{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsMapsHarbor) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsHarbor) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsMapsHarbor) SetId(v int32) {
	o.Id = v
}

// GetX returns the X field value
func (o *FPHSpedVAPIObjectsMapsHarbor) GetX() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsHarbor) GetXOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *FPHSpedVAPIObjectsMapsHarbor) SetX(v float64) {
	o.X = v
}

// GetZ returns the Z field value
func (o *FPHSpedVAPIObjectsMapsHarbor) GetZ() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Z
}

// GetZOk returns a tuple with the Z field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsHarbor) GetZOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Z, true
}

// SetZ sets field value
func (o *FPHSpedVAPIObjectsMapsHarbor) SetZ(v float64) {
	o.Z = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsMapsHarbor) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMapsHarbor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *FPHSpedVAPIObjectsMapsHarbor) SetName(v string) {
	o.Name.Set(&v)
}

func (o FPHSpedVAPIObjectsMapsHarbor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsMapsHarbor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["x"] = o.X
	toSerialize["z"] = o.Z
	toSerialize["name"] = o.Name.Get()
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsMapsHarbor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"x",
		"z",
		"name",
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

	varFPHSpedVAPIObjectsMapsHarbor := _FPHSpedVAPIObjectsMapsHarbor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsMapsHarbor)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsMapsHarbor(varFPHSpedVAPIObjectsMapsHarbor)

	return err
}

type NullableFPHSpedVAPIObjectsMapsHarbor struct {
	value *FPHSpedVAPIObjectsMapsHarbor
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsMapsHarbor) Get() *FPHSpedVAPIObjectsMapsHarbor {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsMapsHarbor) Set(val *FPHSpedVAPIObjectsMapsHarbor) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsMapsHarbor) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsMapsHarbor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsMapsHarbor(val *FPHSpedVAPIObjectsMapsHarbor) *NullableFPHSpedVAPIObjectsMapsHarbor {
	return &NullableFPHSpedVAPIObjectsMapsHarbor{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsMapsHarbor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsMapsHarbor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

