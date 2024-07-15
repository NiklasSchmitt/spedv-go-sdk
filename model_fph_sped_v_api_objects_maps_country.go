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

// checks if the FPHSpedVAPIObjectsMapsCountry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsMapsCountry{}

// FPHSpedVAPIObjectsMapsCountry struct for FPHSpedVAPIObjectsMapsCountry
type FPHSpedVAPIObjectsMapsCountry struct {
	Id int32 `json:"id"`
	Name NullableString `json:"name"`
	ShortName NullableString `json:"shortName"`
	Mwst float64 `json:"mwst"`
	Currency NullableString `json:"currency"`
}

type _FPHSpedVAPIObjectsMapsCountry FPHSpedVAPIObjectsMapsCountry

// NewFPHSpedVAPIObjectsMapsCountry instantiates a new FPHSpedVAPIObjectsMapsCountry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsMapsCountry(id int32, name NullableString, shortName NullableString, mwst float64, currency NullableString) *FPHSpedVAPIObjectsMapsCountry {
	this := FPHSpedVAPIObjectsMapsCountry{}
	this.Id = id
	this.Name = name
	this.ShortName = shortName
	this.Mwst = mwst
	this.Currency = currency
	return &this
}

// NewFPHSpedVAPIObjectsMapsCountryWithDefaults instantiates a new FPHSpedVAPIObjectsMapsCountry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsMapsCountryWithDefaults() *FPHSpedVAPIObjectsMapsCountry {
	this := FPHSpedVAPIObjectsMapsCountry{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsMapsCountry) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsCountry) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsMapsCountry) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsMapsCountry) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMapsCountry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *FPHSpedVAPIObjectsMapsCountry) SetName(v string) {
	o.Name.Set(&v)
}

// GetShortName returns the ShortName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsMapsCountry) GetShortName() string {
	if o == nil || o.ShortName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShortName.Get()
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMapsCountry) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortName.Get(), o.ShortName.IsSet()
}

// SetShortName sets field value
func (o *FPHSpedVAPIObjectsMapsCountry) SetShortName(v string) {
	o.ShortName.Set(&v)
}

// GetMwst returns the Mwst field value
func (o *FPHSpedVAPIObjectsMapsCountry) GetMwst() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Mwst
}

// GetMwstOk returns a tuple with the Mwst field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsCountry) GetMwstOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mwst, true
}

// SetMwst sets field value
func (o *FPHSpedVAPIObjectsMapsCountry) SetMwst(v float64) {
	o.Mwst = v
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsMapsCountry) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMapsCountry) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *FPHSpedVAPIObjectsMapsCountry) SetCurrency(v string) {
	o.Currency.Set(&v)
}

func (o FPHSpedVAPIObjectsMapsCountry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsMapsCountry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["shortName"] = o.ShortName.Get()
	toSerialize["mwst"] = o.Mwst
	toSerialize["currency"] = o.Currency.Get()
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsMapsCountry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"shortName",
		"mwst",
		"currency",
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

	varFPHSpedVAPIObjectsMapsCountry := _FPHSpedVAPIObjectsMapsCountry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsMapsCountry)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsMapsCountry(varFPHSpedVAPIObjectsMapsCountry)

	return err
}

type NullableFPHSpedVAPIObjectsMapsCountry struct {
	value *FPHSpedVAPIObjectsMapsCountry
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsMapsCountry) Get() *FPHSpedVAPIObjectsMapsCountry {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsMapsCountry) Set(val *FPHSpedVAPIObjectsMapsCountry) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsMapsCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsMapsCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsMapsCountry(val *FPHSpedVAPIObjectsMapsCountry) *NullableFPHSpedVAPIObjectsMapsCountry {
	return &NullableFPHSpedVAPIObjectsMapsCountry{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsMapsCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsMapsCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

