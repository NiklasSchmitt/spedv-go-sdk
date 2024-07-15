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

// checks if the FPHSpedVAPIObjectsOMSIBusStop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsOMSIBusStop{}

// FPHSpedVAPIObjectsOMSIBusStop struct for FPHSpedVAPIObjectsOMSIBusStop
type FPHSpedVAPIObjectsOMSIBusStop struct {
	Id int32 `json:"id"`
	// Deprecated
	InternalName NullableString `json:"internalName"`
	InternalNames []string `json:"internalNames"`
	Anzeige1 NullableString `json:"anzeige1"`
	Anzeige2 NullableString `json:"anzeige2"`
	IbiS1 NullableString `json:"ibiS1"`
	IbiS2 NullableString `json:"ibiS2"`
}

type _FPHSpedVAPIObjectsOMSIBusStop FPHSpedVAPIObjectsOMSIBusStop

// NewFPHSpedVAPIObjectsOMSIBusStop instantiates a new FPHSpedVAPIObjectsOMSIBusStop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsOMSIBusStop(id int32, internalName NullableString, internalNames []string, anzeige1 NullableString, anzeige2 NullableString, ibiS1 NullableString, ibiS2 NullableString) *FPHSpedVAPIObjectsOMSIBusStop {
	this := FPHSpedVAPIObjectsOMSIBusStop{}
	this.Id = id
	this.InternalName = internalName
	this.InternalNames = internalNames
	this.Anzeige1 = anzeige1
	this.Anzeige2 = anzeige2
	this.IbiS1 = ibiS1
	this.IbiS2 = ibiS2
	return &this
}

// NewFPHSpedVAPIObjectsOMSIBusStopWithDefaults instantiates a new FPHSpedVAPIObjectsOMSIBusStop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsOMSIBusStopWithDefaults() *FPHSpedVAPIObjectsOMSIBusStop {
	this := FPHSpedVAPIObjectsOMSIBusStop{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsOMSIBusStop) SetId(v int32) {
	o.Id = v
}

// GetInternalName returns the InternalName field value
// If the value is explicit nil, the zero value for string will be returned
// Deprecated
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetInternalName() string {
	if o == nil || o.InternalName.Get() == nil {
		var ret string
		return ret
	}

	return *o.InternalName.Get()
}

// GetInternalNameOk returns a tuple with the InternalName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetInternalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternalName.Get(), o.InternalName.IsSet()
}

// SetInternalName sets field value
// Deprecated
func (o *FPHSpedVAPIObjectsOMSIBusStop) SetInternalName(v string) {
	o.InternalName.Set(&v)
}

// GetInternalNames returns the InternalNames field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetInternalNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InternalNames
}

// GetInternalNamesOk returns a tuple with the InternalNames field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetInternalNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.InternalNames) {
		return nil, false
	}
	return o.InternalNames, true
}

// SetInternalNames sets field value
func (o *FPHSpedVAPIObjectsOMSIBusStop) SetInternalNames(v []string) {
	o.InternalNames = v
}

// GetAnzeige1 returns the Anzeige1 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetAnzeige1() string {
	if o == nil || o.Anzeige1.Get() == nil {
		var ret string
		return ret
	}

	return *o.Anzeige1.Get()
}

// GetAnzeige1Ok returns a tuple with the Anzeige1 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetAnzeige1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Anzeige1.Get(), o.Anzeige1.IsSet()
}

// SetAnzeige1 sets field value
func (o *FPHSpedVAPIObjectsOMSIBusStop) SetAnzeige1(v string) {
	o.Anzeige1.Set(&v)
}

// GetAnzeige2 returns the Anzeige2 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetAnzeige2() string {
	if o == nil || o.Anzeige2.Get() == nil {
		var ret string
		return ret
	}

	return *o.Anzeige2.Get()
}

// GetAnzeige2Ok returns a tuple with the Anzeige2 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetAnzeige2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Anzeige2.Get(), o.Anzeige2.IsSet()
}

// SetAnzeige2 sets field value
func (o *FPHSpedVAPIObjectsOMSIBusStop) SetAnzeige2(v string) {
	o.Anzeige2.Set(&v)
}

// GetIbiS1 returns the IbiS1 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetIbiS1() string {
	if o == nil || o.IbiS1.Get() == nil {
		var ret string
		return ret
	}

	return *o.IbiS1.Get()
}

// GetIbiS1Ok returns a tuple with the IbiS1 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetIbiS1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IbiS1.Get(), o.IbiS1.IsSet()
}

// SetIbiS1 sets field value
func (o *FPHSpedVAPIObjectsOMSIBusStop) SetIbiS1(v string) {
	o.IbiS1.Set(&v)
}

// GetIbiS2 returns the IbiS2 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetIbiS2() string {
	if o == nil || o.IbiS2.Get() == nil {
		var ret string
		return ret
	}

	return *o.IbiS2.Get()
}

// GetIbiS2Ok returns a tuple with the IbiS2 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsOMSIBusStop) GetIbiS2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IbiS2.Get(), o.IbiS2.IsSet()
}

// SetIbiS2 sets field value
func (o *FPHSpedVAPIObjectsOMSIBusStop) SetIbiS2(v string) {
	o.IbiS2.Set(&v)
}

func (o FPHSpedVAPIObjectsOMSIBusStop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsOMSIBusStop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["internalName"] = o.InternalName.Get()
	if o.InternalNames != nil {
		toSerialize["internalNames"] = o.InternalNames
	}
	toSerialize["anzeige1"] = o.Anzeige1.Get()
	toSerialize["anzeige2"] = o.Anzeige2.Get()
	toSerialize["ibiS1"] = o.IbiS1.Get()
	toSerialize["ibiS2"] = o.IbiS2.Get()
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsOMSIBusStop) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"internalName",
		"internalNames",
		"anzeige1",
		"anzeige2",
		"ibiS1",
		"ibiS2",
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

	varFPHSpedVAPIObjectsOMSIBusStop := _FPHSpedVAPIObjectsOMSIBusStop{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsOMSIBusStop)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsOMSIBusStop(varFPHSpedVAPIObjectsOMSIBusStop)

	return err
}

type NullableFPHSpedVAPIObjectsOMSIBusStop struct {
	value *FPHSpedVAPIObjectsOMSIBusStop
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsOMSIBusStop) Get() *FPHSpedVAPIObjectsOMSIBusStop {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsOMSIBusStop) Set(val *FPHSpedVAPIObjectsOMSIBusStop) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsOMSIBusStop) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsOMSIBusStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsOMSIBusStop(val *FPHSpedVAPIObjectsOMSIBusStop) *NullableFPHSpedVAPIObjectsOMSIBusStop {
	return &NullableFPHSpedVAPIObjectsOMSIBusStop{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsOMSIBusStop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsOMSIBusStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

