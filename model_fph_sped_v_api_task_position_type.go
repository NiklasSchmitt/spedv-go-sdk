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

// FPHSpedVAPITaskPositionType   0 = StateChange  1 = RefusedTruckInfo  2 = Refueled  3 = FerryUsed  -1 = NotSet
type FPHSpedVAPITaskPositionType int32

// List of FPH.SpedV.API.TaskPositionType
const (
	_0 FPHSpedVAPITaskPositionType = 0
	_1 FPHSpedVAPITaskPositionType = 1
	_2 FPHSpedVAPITaskPositionType = 2
	_3 FPHSpedVAPITaskPositionType = 3
	_MINUS_1 FPHSpedVAPITaskPositionType = -1
)

// All allowed values of FPHSpedVAPITaskPositionType enum
var AllowedFPHSpedVAPITaskPositionTypeEnumValues = []FPHSpedVAPITaskPositionType{
	0,
	1,
	2,
	3,
	-1,
}

func (v *FPHSpedVAPITaskPositionType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FPHSpedVAPITaskPositionType(value)
	for _, existing := range AllowedFPHSpedVAPITaskPositionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FPHSpedVAPITaskPositionType", value)
}

// NewFPHSpedVAPITaskPositionTypeFromValue returns a pointer to a valid FPHSpedVAPITaskPositionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFPHSpedVAPITaskPositionTypeFromValue(v int32) (*FPHSpedVAPITaskPositionType, error) {
	ev := FPHSpedVAPITaskPositionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FPHSpedVAPITaskPositionType: valid values are %v", v, AllowedFPHSpedVAPITaskPositionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FPHSpedVAPITaskPositionType) IsValid() bool {
	for _, existing := range AllowedFPHSpedVAPITaskPositionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FPH.SpedV.API.TaskPositionType value
func (v FPHSpedVAPITaskPositionType) Ptr() *FPHSpedVAPITaskPositionType {
	return &v
}

type NullableFPHSpedVAPITaskPositionType struct {
	value *FPHSpedVAPITaskPositionType
	isSet bool
}

func (v NullableFPHSpedVAPITaskPositionType) Get() *FPHSpedVAPITaskPositionType {
	return v.value
}

func (v *NullableFPHSpedVAPITaskPositionType) Set(val *FPHSpedVAPITaskPositionType) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPITaskPositionType) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPITaskPositionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPITaskPositionType(val *FPHSpedVAPITaskPositionType) *NullableFPHSpedVAPITaskPositionType {
	return &NullableFPHSpedVAPITaskPositionType{value: val, isSet: true}
}

func (v NullableFPHSpedVAPITaskPositionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPITaskPositionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
