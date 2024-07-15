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

// FPHSpedVAPIObjectsSimRailSimRailTrainControlType   0 = Ghost  1 = Bot  2 = Player
type FPHSpedVAPIObjectsSimRailSimRailTrainControlType int32

// List of FPH.SpedV.API.Objects.SimRail.SimRailTrainControlType
const (
	_0 FPHSpedVAPIObjectsSimRailSimRailTrainControlType = 0
	_1 FPHSpedVAPIObjectsSimRailSimRailTrainControlType = 1
	_2 FPHSpedVAPIObjectsSimRailSimRailTrainControlType = 2
)

// All allowed values of FPHSpedVAPIObjectsSimRailSimRailTrainControlType enum
var AllowedFPHSpedVAPIObjectsSimRailSimRailTrainControlTypeEnumValues = []FPHSpedVAPIObjectsSimRailSimRailTrainControlType{
	0,
	1,
	2,
}

func (v *FPHSpedVAPIObjectsSimRailSimRailTrainControlType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FPHSpedVAPIObjectsSimRailSimRailTrainControlType(value)
	for _, existing := range AllowedFPHSpedVAPIObjectsSimRailSimRailTrainControlTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FPHSpedVAPIObjectsSimRailSimRailTrainControlType", value)
}

// NewFPHSpedVAPIObjectsSimRailSimRailTrainControlTypeFromValue returns a pointer to a valid FPHSpedVAPIObjectsSimRailSimRailTrainControlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFPHSpedVAPIObjectsSimRailSimRailTrainControlTypeFromValue(v int32) (*FPHSpedVAPIObjectsSimRailSimRailTrainControlType, error) {
	ev := FPHSpedVAPIObjectsSimRailSimRailTrainControlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FPHSpedVAPIObjectsSimRailSimRailTrainControlType: valid values are %v", v, AllowedFPHSpedVAPIObjectsSimRailSimRailTrainControlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FPHSpedVAPIObjectsSimRailSimRailTrainControlType) IsValid() bool {
	for _, existing := range AllowedFPHSpedVAPIObjectsSimRailSimRailTrainControlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FPH.SpedV.API.Objects.SimRail.SimRailTrainControlType value
func (v FPHSpedVAPIObjectsSimRailSimRailTrainControlType) Ptr() *FPHSpedVAPIObjectsSimRailSimRailTrainControlType {
	return &v
}

type NullableFPHSpedVAPIObjectsSimRailSimRailTrainControlType struct {
	value *FPHSpedVAPIObjectsSimRailSimRailTrainControlType
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailTrainControlType) Get() *FPHSpedVAPIObjectsSimRailSimRailTrainControlType {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailTrainControlType) Set(val *FPHSpedVAPIObjectsSimRailSimRailTrainControlType) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailTrainControlType) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailTrainControlType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSimRailSimRailTrainControlType(val *FPHSpedVAPIObjectsSimRailSimRailTrainControlType) *NullableFPHSpedVAPIObjectsSimRailSimRailTrainControlType {
	return &NullableFPHSpedVAPIObjectsSimRailSimRailTrainControlType{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailTrainControlType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailTrainControlType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
