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

// checks if the FPHSpedVAPIObjectsKontorJobOfferFreight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsKontorJobOfferFreight{}

// FPHSpedVAPIObjectsKontorJobOfferFreight struct for FPHSpedVAPIObjectsKontorJobOfferFreight
type FPHSpedVAPIObjectsKontorJobOfferFreight struct {
	Id int32 `json:"id"`
	Name NullableString `json:"name"`
	PrizeCoeffizient float64 `json:"prizeCoeffizient"`
	IsSpecialFreight bool `json:"isSpecialFreight"`
}

type _FPHSpedVAPIObjectsKontorJobOfferFreight FPHSpedVAPIObjectsKontorJobOfferFreight

// NewFPHSpedVAPIObjectsKontorJobOfferFreight instantiates a new FPHSpedVAPIObjectsKontorJobOfferFreight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsKontorJobOfferFreight(id int32, name NullableString, prizeCoeffizient float64, isSpecialFreight bool) *FPHSpedVAPIObjectsKontorJobOfferFreight {
	this := FPHSpedVAPIObjectsKontorJobOfferFreight{}
	this.Id = id
	this.Name = name
	this.PrizeCoeffizient = prizeCoeffizient
	this.IsSpecialFreight = isSpecialFreight
	return &this
}

// NewFPHSpedVAPIObjectsKontorJobOfferFreightWithDefaults instantiates a new FPHSpedVAPIObjectsKontorJobOfferFreight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsKontorJobOfferFreightWithDefaults() *FPHSpedVAPIObjectsKontorJobOfferFreight {
	this := FPHSpedVAPIObjectsKontorJobOfferFreight{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) SetName(v string) {
	o.Name.Set(&v)
}

// GetPrizeCoeffizient returns the PrizeCoeffizient field value
func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) GetPrizeCoeffizient() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.PrizeCoeffizient
}

// GetPrizeCoeffizientOk returns a tuple with the PrizeCoeffizient field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) GetPrizeCoeffizientOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrizeCoeffizient, true
}

// SetPrizeCoeffizient sets field value
func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) SetPrizeCoeffizient(v float64) {
	o.PrizeCoeffizient = v
}

// GetIsSpecialFreight returns the IsSpecialFreight field value
func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) GetIsSpecialFreight() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpecialFreight
}

// GetIsSpecialFreightOk returns a tuple with the IsSpecialFreight field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) GetIsSpecialFreightOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSpecialFreight, true
}

// SetIsSpecialFreight sets field value
func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) SetIsSpecialFreight(v bool) {
	o.IsSpecialFreight = v
}

func (o FPHSpedVAPIObjectsKontorJobOfferFreight) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsKontorJobOfferFreight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["prizeCoeffizient"] = o.PrizeCoeffizient
	toSerialize["isSpecialFreight"] = o.IsSpecialFreight
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsKontorJobOfferFreight) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"prizeCoeffizient",
		"isSpecialFreight",
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

	varFPHSpedVAPIObjectsKontorJobOfferFreight := _FPHSpedVAPIObjectsKontorJobOfferFreight{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsKontorJobOfferFreight)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsKontorJobOfferFreight(varFPHSpedVAPIObjectsKontorJobOfferFreight)

	return err
}

type NullableFPHSpedVAPIObjectsKontorJobOfferFreight struct {
	value *FPHSpedVAPIObjectsKontorJobOfferFreight
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsKontorJobOfferFreight) Get() *FPHSpedVAPIObjectsKontorJobOfferFreight {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsKontorJobOfferFreight) Set(val *FPHSpedVAPIObjectsKontorJobOfferFreight) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsKontorJobOfferFreight) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsKontorJobOfferFreight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsKontorJobOfferFreight(val *FPHSpedVAPIObjectsKontorJobOfferFreight) *NullableFPHSpedVAPIObjectsKontorJobOfferFreight {
	return &NullableFPHSpedVAPIObjectsKontorJobOfferFreight{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsKontorJobOfferFreight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsKontorJobOfferFreight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


