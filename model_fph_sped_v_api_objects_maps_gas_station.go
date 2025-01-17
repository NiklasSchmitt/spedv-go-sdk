/*
SpedV API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the FPHSpedVAPIObjectsMapsGasStation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsMapsGasStation{}

// FPHSpedVAPIObjectsMapsGasStation struct for FPHSpedVAPIObjectsMapsGasStation
type FPHSpedVAPIObjectsMapsGasStation struct {
	Id int32 `json:"id"`
	LiterPrize float64 `json:"literPrize"`
	LastReported time.Time `json:"lastReported"`
	Currency NullableString `json:"currency"`
	X float64 `json:"x"`
	Z float64 `json:"z"`
	ReportedBy NullableFPHSpedVAPIObjectsUsersUser `json:"reportedBy"`
}

type _FPHSpedVAPIObjectsMapsGasStation FPHSpedVAPIObjectsMapsGasStation

// NewFPHSpedVAPIObjectsMapsGasStation instantiates a new FPHSpedVAPIObjectsMapsGasStation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsMapsGasStation(id int32, literPrize float64, lastReported time.Time, currency NullableString, x float64, z float64, reportedBy NullableFPHSpedVAPIObjectsUsersUser) *FPHSpedVAPIObjectsMapsGasStation {
	this := FPHSpedVAPIObjectsMapsGasStation{}
	this.Id = id
	this.LiterPrize = literPrize
	this.LastReported = lastReported
	this.Currency = currency
	this.X = x
	this.Z = z
	this.ReportedBy = reportedBy
	return &this
}

// NewFPHSpedVAPIObjectsMapsGasStationWithDefaults instantiates a new FPHSpedVAPIObjectsMapsGasStation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsMapsGasStationWithDefaults() *FPHSpedVAPIObjectsMapsGasStation {
	this := FPHSpedVAPIObjectsMapsGasStation{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsMapsGasStation) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsGasStation) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsMapsGasStation) SetId(v int32) {
	o.Id = v
}

// GetLiterPrize returns the LiterPrize field value
func (o *FPHSpedVAPIObjectsMapsGasStation) GetLiterPrize() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.LiterPrize
}

// GetLiterPrizeOk returns a tuple with the LiterPrize field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsGasStation) GetLiterPrizeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LiterPrize, true
}

// SetLiterPrize sets field value
func (o *FPHSpedVAPIObjectsMapsGasStation) SetLiterPrize(v float64) {
	o.LiterPrize = v
}

// GetLastReported returns the LastReported field value
func (o *FPHSpedVAPIObjectsMapsGasStation) GetLastReported() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastReported
}

// GetLastReportedOk returns a tuple with the LastReported field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsGasStation) GetLastReportedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastReported, true
}

// SetLastReported sets field value
func (o *FPHSpedVAPIObjectsMapsGasStation) SetLastReported(v time.Time) {
	o.LastReported = v
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsMapsGasStation) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMapsGasStation) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *FPHSpedVAPIObjectsMapsGasStation) SetCurrency(v string) {
	o.Currency.Set(&v)
}

// GetX returns the X field value
func (o *FPHSpedVAPIObjectsMapsGasStation) GetX() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsGasStation) GetXOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *FPHSpedVAPIObjectsMapsGasStation) SetX(v float64) {
	o.X = v
}

// GetZ returns the Z field value
func (o *FPHSpedVAPIObjectsMapsGasStation) GetZ() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Z
}

// GetZOk returns a tuple with the Z field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMapsGasStation) GetZOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Z, true
}

// SetZ sets field value
func (o *FPHSpedVAPIObjectsMapsGasStation) SetZ(v float64) {
	o.Z = v
}

// GetReportedBy returns the ReportedBy field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsUsersUser will be returned
func (o *FPHSpedVAPIObjectsMapsGasStation) GetReportedBy() FPHSpedVAPIObjectsUsersUser {
	if o == nil || o.ReportedBy.Get() == nil {
		var ret FPHSpedVAPIObjectsUsersUser
		return ret
	}

	return *o.ReportedBy.Get()
}

// GetReportedByOk returns a tuple with the ReportedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMapsGasStation) GetReportedByOk() (*FPHSpedVAPIObjectsUsersUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportedBy.Get(), o.ReportedBy.IsSet()
}

// SetReportedBy sets field value
func (o *FPHSpedVAPIObjectsMapsGasStation) SetReportedBy(v FPHSpedVAPIObjectsUsersUser) {
	o.ReportedBy.Set(&v)
}

func (o FPHSpedVAPIObjectsMapsGasStation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsMapsGasStation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["literPrize"] = o.LiterPrize
	toSerialize["lastReported"] = o.LastReported
	toSerialize["currency"] = o.Currency.Get()
	toSerialize["x"] = o.X
	toSerialize["z"] = o.Z
	toSerialize["reportedBy"] = o.ReportedBy.Get()
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsMapsGasStation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"literPrize",
		"lastReported",
		"currency",
		"x",
		"z",
		"reportedBy",
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

	varFPHSpedVAPIObjectsMapsGasStation := _FPHSpedVAPIObjectsMapsGasStation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsMapsGasStation)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsMapsGasStation(varFPHSpedVAPIObjectsMapsGasStation)

	return err
}

type NullableFPHSpedVAPIObjectsMapsGasStation struct {
	value *FPHSpedVAPIObjectsMapsGasStation
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsMapsGasStation) Get() *FPHSpedVAPIObjectsMapsGasStation {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsMapsGasStation) Set(val *FPHSpedVAPIObjectsMapsGasStation) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsMapsGasStation) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsMapsGasStation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsMapsGasStation(val *FPHSpedVAPIObjectsMapsGasStation) *NullableFPHSpedVAPIObjectsMapsGasStation {
	return &NullableFPHSpedVAPIObjectsMapsGasStation{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsMapsGasStation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsMapsGasStation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


