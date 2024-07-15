/*
SpedV API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// FPHSpedVAPIEnumsBoniTypeType   0 = Value  1 = KM  -1 = NotSet
type FPHSpedVAPIEnumsBoniTypeType int32

// List of FPH.SpedV.API.Enums.BoniTypeType
const (
	_0 FPHSpedVAPIEnumsBoniTypeType = 0
	_1 FPHSpedVAPIEnumsBoniTypeType = 1
	_MINUS_1 FPHSpedVAPIEnumsBoniTypeType = -1
)

// All allowed values of FPHSpedVAPIEnumsBoniTypeType enum
var AllowedFPHSpedVAPIEnumsBoniTypeTypeEnumValues = []FPHSpedVAPIEnumsBoniTypeType{
	0,
	1,
	-1,
}

func (v *FPHSpedVAPIEnumsBoniTypeType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FPHSpedVAPIEnumsBoniTypeType(value)
	for _, existing := range AllowedFPHSpedVAPIEnumsBoniTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FPHSpedVAPIEnumsBoniTypeType", value)
}

// NewFPHSpedVAPIEnumsBoniTypeTypeFromValue returns a pointer to a valid FPHSpedVAPIEnumsBoniTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFPHSpedVAPIEnumsBoniTypeTypeFromValue(v int32) (*FPHSpedVAPIEnumsBoniTypeType, error) {
	ev := FPHSpedVAPIEnumsBoniTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FPHSpedVAPIEnumsBoniTypeType: valid values are %v", v, AllowedFPHSpedVAPIEnumsBoniTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FPHSpedVAPIEnumsBoniTypeType) IsValid() bool {
	for _, existing := range AllowedFPHSpedVAPIEnumsBoniTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FPH.SpedV.API.Enums.BoniTypeType value
func (v FPHSpedVAPIEnumsBoniTypeType) Ptr() *FPHSpedVAPIEnumsBoniTypeType {
	return &v
}

type NullableFPHSpedVAPIEnumsBoniTypeType struct {
	value *FPHSpedVAPIEnumsBoniTypeType
	isSet bool
}

func (v NullableFPHSpedVAPIEnumsBoniTypeType) Get() *FPHSpedVAPIEnumsBoniTypeType {
	return v.value
}

func (v *NullableFPHSpedVAPIEnumsBoniTypeType) Set(val *FPHSpedVAPIEnumsBoniTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIEnumsBoniTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIEnumsBoniTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIEnumsBoniTypeType(val *FPHSpedVAPIEnumsBoniTypeType) *NullableFPHSpedVAPIEnumsBoniTypeType {
	return &NullableFPHSpedVAPIEnumsBoniTypeType{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIEnumsBoniTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIEnumsBoniTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

