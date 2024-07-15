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

// checks if the FPHSpedVAPIObjectsSaveGameSaveGameColor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsSaveGameSaveGameColor{}

// FPHSpedVAPIObjectsSaveGameSaveGameColor struct for FPHSpedVAPIObjectsSaveGameSaveGameColor
type FPHSpedVAPIObjectsSaveGameSaveGameColor struct {
	Index int32 `json:"index"`
	Color1 NullableString `json:"color1"`
	Color2 NullableString `json:"color2"`
	Color3 NullableString `json:"color3"`
	Color4 NullableString `json:"color4"`
}

type _FPHSpedVAPIObjectsSaveGameSaveGameColor FPHSpedVAPIObjectsSaveGameSaveGameColor

// NewFPHSpedVAPIObjectsSaveGameSaveGameColor instantiates a new FPHSpedVAPIObjectsSaveGameSaveGameColor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsSaveGameSaveGameColor(index int32, color1 NullableString, color2 NullableString, color3 NullableString, color4 NullableString) *FPHSpedVAPIObjectsSaveGameSaveGameColor {
	this := FPHSpedVAPIObjectsSaveGameSaveGameColor{}
	this.Index = index
	this.Color1 = color1
	this.Color2 = color2
	this.Color3 = color3
	this.Color4 = color4
	return &this
}

// NewFPHSpedVAPIObjectsSaveGameSaveGameColorWithDefaults instantiates a new FPHSpedVAPIObjectsSaveGameSaveGameColor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsSaveGameSaveGameColorWithDefaults() *FPHSpedVAPIObjectsSaveGameSaveGameColor {
	this := FPHSpedVAPIObjectsSaveGameSaveGameColor{}
	return &this
}

// GetIndex returns the Index field value
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) SetIndex(v int32) {
	o.Index = v
}

// GetColor1 returns the Color1 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) GetColor1() string {
	if o == nil || o.Color1.Get() == nil {
		var ret string
		return ret
	}

	return *o.Color1.Get()
}

// GetColor1Ok returns a tuple with the Color1 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) GetColor1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color1.Get(), o.Color1.IsSet()
}

// SetColor1 sets field value
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) SetColor1(v string) {
	o.Color1.Set(&v)
}

// GetColor2 returns the Color2 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) GetColor2() string {
	if o == nil || o.Color2.Get() == nil {
		var ret string
		return ret
	}

	return *o.Color2.Get()
}

// GetColor2Ok returns a tuple with the Color2 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) GetColor2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color2.Get(), o.Color2.IsSet()
}

// SetColor2 sets field value
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) SetColor2(v string) {
	o.Color2.Set(&v)
}

// GetColor3 returns the Color3 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) GetColor3() string {
	if o == nil || o.Color3.Get() == nil {
		var ret string
		return ret
	}

	return *o.Color3.Get()
}

// GetColor3Ok returns a tuple with the Color3 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) GetColor3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color3.Get(), o.Color3.IsSet()
}

// SetColor3 sets field value
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) SetColor3(v string) {
	o.Color3.Set(&v)
}

// GetColor4 returns the Color4 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) GetColor4() string {
	if o == nil || o.Color4.Get() == nil {
		var ret string
		return ret
	}

	return *o.Color4.Get()
}

// GetColor4Ok returns a tuple with the Color4 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) GetColor4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color4.Get(), o.Color4.IsSet()
}

// SetColor4 sets field value
func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) SetColor4(v string) {
	o.Color4.Set(&v)
}

func (o FPHSpedVAPIObjectsSaveGameSaveGameColor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsSaveGameSaveGameColor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["color1"] = o.Color1.Get()
	toSerialize["color2"] = o.Color2.Get()
	toSerialize["color3"] = o.Color3.Get()
	toSerialize["color4"] = o.Color4.Get()
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsSaveGameSaveGameColor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"color1",
		"color2",
		"color3",
		"color4",
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

	varFPHSpedVAPIObjectsSaveGameSaveGameColor := _FPHSpedVAPIObjectsSaveGameSaveGameColor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsSaveGameSaveGameColor)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsSaveGameSaveGameColor(varFPHSpedVAPIObjectsSaveGameSaveGameColor)

	return err
}

type NullableFPHSpedVAPIObjectsSaveGameSaveGameColor struct {
	value *FPHSpedVAPIObjectsSaveGameSaveGameColor
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSaveGameSaveGameColor) Get() *FPHSpedVAPIObjectsSaveGameSaveGameColor {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSaveGameSaveGameColor) Set(val *FPHSpedVAPIObjectsSaveGameSaveGameColor) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSaveGameSaveGameColor) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSaveGameSaveGameColor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSaveGameSaveGameColor(val *FPHSpedVAPIObjectsSaveGameSaveGameColor) *NullableFPHSpedVAPIObjectsSaveGameSaveGameColor {
	return &NullableFPHSpedVAPIObjectsSaveGameSaveGameColor{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSaveGameSaveGameColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSaveGameSaveGameColor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

