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

// checks if the FPHSpedVAPIObjectsMapsMap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsMapsMap{}

// FPHSpedVAPIObjectsMapsMap struct for FPHSpedVAPIObjectsMapsMap
type FPHSpedVAPIObjectsMapsMap struct {
	Id int32 `json:"id"`
	Name NullableString `json:"name"`
	ShortName NullableString `json:"shortName"`
	Game NullableString `json:"game"`
	IsImported bool `json:"isImported"`
	IsOriginal bool `json:"isOriginal"`
	Needs NullableFPHSpedVAPIObjectsMapsMap `json:"needs"`
}

type _FPHSpedVAPIObjectsMapsMap FPHSpedVAPIObjectsMapsMap

// NewFPHSpedVAPIObjectsMapsMap instantiates a new FPHSpedVAPIObjectsMapsMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsMapsMap(id int32, name NullableString, shortName NullableString, game NullableString, isImported bool, isOriginal bool, needs NullableFPHSpedVAPIObjectsMapsMap) *FPHSpedVAPIObjectsMapsMap {
	this := FPHSpedVAPIObjectsMapsMap{}
	this.Id = id
	this.Name = name
	this.ShortName = shortName
	this.Game = game
	this.IsImported = isImported
	this.IsOriginal = isOriginal
	this.Needs = needs
	return &this
}

// NewFPHSpedVAPIObjectsMapsMapWithDefaults instantiates a new FPHSpedVAPIObjectsMapsMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsMapsMapWithDefaults() *FPHSpedVAPIObjectsMapsMap {
	this := FPHSpedVAPIObjectsMapsMap{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsMapsMap) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsMap) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsMapsMap) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsMapsMap) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMapsMap) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *FPHSpedVAPIObjectsMapsMap) SetName(v string) {
	o.Name.Set(&v)
}

// GetShortName returns the ShortName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsMapsMap) GetShortName() string {
	if o == nil || o.ShortName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShortName.Get()
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMapsMap) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortName.Get(), o.ShortName.IsSet()
}

// SetShortName sets field value
func (o *FPHSpedVAPIObjectsMapsMap) SetShortName(v string) {
	o.ShortName.Set(&v)
}

// GetGame returns the Game field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsMapsMap) GetGame() string {
	if o == nil || o.Game.Get() == nil {
		var ret string
		return ret
	}

	return *o.Game.Get()
}

// GetGameOk returns a tuple with the Game field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMapsMap) GetGameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Game.Get(), o.Game.IsSet()
}

// SetGame sets field value
func (o *FPHSpedVAPIObjectsMapsMap) SetGame(v string) {
	o.Game.Set(&v)
}

// GetIsImported returns the IsImported field value
func (o *FPHSpedVAPIObjectsMapsMap) GetIsImported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsImported
}

// GetIsImportedOk returns a tuple with the IsImported field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsMap) GetIsImportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsImported, true
}

// SetIsImported sets field value
func (o *FPHSpedVAPIObjectsMapsMap) SetIsImported(v bool) {
	o.IsImported = v
}

// GetIsOriginal returns the IsOriginal field value
func (o *FPHSpedVAPIObjectsMapsMap) GetIsOriginal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsOriginal
}

// GetIsOriginalOk returns a tuple with the IsOriginal field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsMap) GetIsOriginalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsOriginal, true
}

// SetIsOriginal sets field value
func (o *FPHSpedVAPIObjectsMapsMap) SetIsOriginal(v bool) {
	o.IsOriginal = v
}

// GetNeeds returns the Needs field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsMapsMap will be returned
func (o *FPHSpedVAPIObjectsMapsMap) GetNeeds() FPHSpedVAPIObjectsMapsMap {
	if o == nil || o.Needs.Get() == nil {
		var ret FPHSpedVAPIObjectsMapsMap
		return ret
	}

	return *o.Needs.Get()
}

// GetNeedsOk returns a tuple with the Needs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMapsMap) GetNeedsOk() (*FPHSpedVAPIObjectsMapsMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Needs.Get(), o.Needs.IsSet()
}

// SetNeeds sets field value
func (o *FPHSpedVAPIObjectsMapsMap) SetNeeds(v FPHSpedVAPIObjectsMapsMap) {
	o.Needs.Set(&v)
}

func (o FPHSpedVAPIObjectsMapsMap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsMapsMap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["shortName"] = o.ShortName.Get()
	toSerialize["game"] = o.Game.Get()
	toSerialize["isImported"] = o.IsImported
	toSerialize["isOriginal"] = o.IsOriginal
	toSerialize["needs"] = o.Needs.Get()
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsMapsMap) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"shortName",
		"game",
		"isImported",
		"isOriginal",
		"needs",
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

	varFPHSpedVAPIObjectsMapsMap := _FPHSpedVAPIObjectsMapsMap{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsMapsMap)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsMapsMap(varFPHSpedVAPIObjectsMapsMap)

	return err
}

type NullableFPHSpedVAPIObjectsMapsMap struct {
	value *FPHSpedVAPIObjectsMapsMap
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsMapsMap) Get() *FPHSpedVAPIObjectsMapsMap {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsMapsMap) Set(val *FPHSpedVAPIObjectsMapsMap) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsMapsMap) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsMapsMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsMapsMap(val *FPHSpedVAPIObjectsMapsMap) *NullableFPHSpedVAPIObjectsMapsMap {
	return &NullableFPHSpedVAPIObjectsMapsMap{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsMapsMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsMapsMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


