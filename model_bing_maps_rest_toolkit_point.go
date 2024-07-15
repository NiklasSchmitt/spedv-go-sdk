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

// checks if the BingMapsRESTToolkitPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BingMapsRESTToolkitPoint{}

// BingMapsRESTToolkitPoint struct for BingMapsRESTToolkitPoint
type BingMapsRESTToolkitPoint struct {
	Type NullableString `json:"type"`
	Coordinates []float64 `json:"coordinates"`
	CalculationMethod NullableString `json:"calculationMethod"`
	UsageTypes []string `json:"usageTypes"`
	BoundingBox []float64 `json:"boundingBox"`
}

type _BingMapsRESTToolkitPoint BingMapsRESTToolkitPoint

// NewBingMapsRESTToolkitPoint instantiates a new BingMapsRESTToolkitPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBingMapsRESTToolkitPoint(type_ NullableString, coordinates []float64, calculationMethod NullableString, usageTypes []string, boundingBox []float64) *BingMapsRESTToolkitPoint {
	this := BingMapsRESTToolkitPoint{}
	this.Type = type_
	this.Coordinates = coordinates
	this.CalculationMethod = calculationMethod
	this.UsageTypes = usageTypes
	this.BoundingBox = boundingBox
	return &this
}

// NewBingMapsRESTToolkitPointWithDefaults instantiates a new BingMapsRESTToolkitPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBingMapsRESTToolkitPointWithDefaults() *BingMapsRESTToolkitPoint {
	this := BingMapsRESTToolkitPoint{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BingMapsRESTToolkitPoint) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BingMapsRESTToolkitPoint) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *BingMapsRESTToolkitPoint) SetType(v string) {
	o.Type.Set(&v)
}

// GetCoordinates returns the Coordinates field value
// If the value is explicit nil, the zero value for []float64 will be returned
func (o *BingMapsRESTToolkitPoint) GetCoordinates() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}

	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BingMapsRESTToolkitPoint) GetCoordinatesOk() ([]float64, bool) {
	if o == nil || IsNil(o.Coordinates) {
		return nil, false
	}
	return o.Coordinates, true
}

// SetCoordinates sets field value
func (o *BingMapsRESTToolkitPoint) SetCoordinates(v []float64) {
	o.Coordinates = v
}

// GetCalculationMethod returns the CalculationMethod field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BingMapsRESTToolkitPoint) GetCalculationMethod() string {
	if o == nil || o.CalculationMethod.Get() == nil {
		var ret string
		return ret
	}

	return *o.CalculationMethod.Get()
}

// GetCalculationMethodOk returns a tuple with the CalculationMethod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BingMapsRESTToolkitPoint) GetCalculationMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CalculationMethod.Get(), o.CalculationMethod.IsSet()
}

// SetCalculationMethod sets field value
func (o *BingMapsRESTToolkitPoint) SetCalculationMethod(v string) {
	o.CalculationMethod.Set(&v)
}

// GetUsageTypes returns the UsageTypes field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *BingMapsRESTToolkitPoint) GetUsageTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UsageTypes
}

// GetUsageTypesOk returns a tuple with the UsageTypes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BingMapsRESTToolkitPoint) GetUsageTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.UsageTypes) {
		return nil, false
	}
	return o.UsageTypes, true
}

// SetUsageTypes sets field value
func (o *BingMapsRESTToolkitPoint) SetUsageTypes(v []string) {
	o.UsageTypes = v
}

// GetBoundingBox returns the BoundingBox field value
// If the value is explicit nil, the zero value for []float64 will be returned
func (o *BingMapsRESTToolkitPoint) GetBoundingBox() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}

	return o.BoundingBox
}

// GetBoundingBoxOk returns a tuple with the BoundingBox field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BingMapsRESTToolkitPoint) GetBoundingBoxOk() ([]float64, bool) {
	if o == nil || IsNil(o.BoundingBox) {
		return nil, false
	}
	return o.BoundingBox, true
}

// SetBoundingBox sets field value
func (o *BingMapsRESTToolkitPoint) SetBoundingBox(v []float64) {
	o.BoundingBox = v
}

func (o BingMapsRESTToolkitPoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BingMapsRESTToolkitPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type.Get()
	if o.Coordinates != nil {
		toSerialize["coordinates"] = o.Coordinates
	}
	toSerialize["calculationMethod"] = o.CalculationMethod.Get()
	if o.UsageTypes != nil {
		toSerialize["usageTypes"] = o.UsageTypes
	}
	if o.BoundingBox != nil {
		toSerialize["boundingBox"] = o.BoundingBox
	}
	return toSerialize, nil
}

func (o *BingMapsRESTToolkitPoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"coordinates",
		"calculationMethod",
		"usageTypes",
		"boundingBox",
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

	varBingMapsRESTToolkitPoint := _BingMapsRESTToolkitPoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBingMapsRESTToolkitPoint)

	if err != nil {
		return err
	}

	*o = BingMapsRESTToolkitPoint(varBingMapsRESTToolkitPoint)

	return err
}

type NullableBingMapsRESTToolkitPoint struct {
	value *BingMapsRESTToolkitPoint
	isSet bool
}

func (v NullableBingMapsRESTToolkitPoint) Get() *BingMapsRESTToolkitPoint {
	return v.value
}

func (v *NullableBingMapsRESTToolkitPoint) Set(val *BingMapsRESTToolkitPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableBingMapsRESTToolkitPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableBingMapsRESTToolkitPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBingMapsRESTToolkitPoint(val *BingMapsRESTToolkitPoint) *NullableBingMapsRESTToolkitPoint {
	return &NullableBingMapsRESTToolkitPoint{value: val, isSet: true}
}

func (v NullableBingMapsRESTToolkitPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBingMapsRESTToolkitPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

