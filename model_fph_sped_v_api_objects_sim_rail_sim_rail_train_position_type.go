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

// FPHSpedVAPIObjectsSimRailSimRailTrainPositionType   0 = InStation  1 = BetweenStations  2 = ApproachingStation
type FPHSpedVAPIObjectsSimRailSimRailTrainPositionType int32

// List of FPH.SpedV.API.Objects.SimRail.SimRailTrainPositionType
const (
	_0 FPHSpedVAPIObjectsSimRailSimRailTrainPositionType = 0
	_1 FPHSpedVAPIObjectsSimRailSimRailTrainPositionType = 1
	_2 FPHSpedVAPIObjectsSimRailSimRailTrainPositionType = 2
)

// All allowed values of FPHSpedVAPIObjectsSimRailSimRailTrainPositionType enum
var AllowedFPHSpedVAPIObjectsSimRailSimRailTrainPositionTypeEnumValues = []FPHSpedVAPIObjectsSimRailSimRailTrainPositionType{
	0,
	1,
	2,
}

func (v *FPHSpedVAPIObjectsSimRailSimRailTrainPositionType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FPHSpedVAPIObjectsSimRailSimRailTrainPositionType(value)
	for _, existing := range AllowedFPHSpedVAPIObjectsSimRailSimRailTrainPositionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FPHSpedVAPIObjectsSimRailSimRailTrainPositionType", value)
}

// NewFPHSpedVAPIObjectsSimRailSimRailTrainPositionTypeFromValue returns a pointer to a valid FPHSpedVAPIObjectsSimRailSimRailTrainPositionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFPHSpedVAPIObjectsSimRailSimRailTrainPositionTypeFromValue(v int32) (*FPHSpedVAPIObjectsSimRailSimRailTrainPositionType, error) {
	ev := FPHSpedVAPIObjectsSimRailSimRailTrainPositionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FPHSpedVAPIObjectsSimRailSimRailTrainPositionType: valid values are %v", v, AllowedFPHSpedVAPIObjectsSimRailSimRailTrainPositionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FPHSpedVAPIObjectsSimRailSimRailTrainPositionType) IsValid() bool {
	for _, existing := range AllowedFPHSpedVAPIObjectsSimRailSimRailTrainPositionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FPH.SpedV.API.Objects.SimRail.SimRailTrainPositionType value
func (v FPHSpedVAPIObjectsSimRailSimRailTrainPositionType) Ptr() *FPHSpedVAPIObjectsSimRailSimRailTrainPositionType {
	return &v
}

type NullableFPHSpedVAPIObjectsSimRailSimRailTrainPositionType struct {
	value *FPHSpedVAPIObjectsSimRailSimRailTrainPositionType
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailTrainPositionType) Get() *FPHSpedVAPIObjectsSimRailSimRailTrainPositionType {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailTrainPositionType) Set(val *FPHSpedVAPIObjectsSimRailSimRailTrainPositionType) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailTrainPositionType) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailTrainPositionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSimRailSimRailTrainPositionType(val *FPHSpedVAPIObjectsSimRailSimRailTrainPositionType) *NullableFPHSpedVAPIObjectsSimRailSimRailTrainPositionType {
	return &NullableFPHSpedVAPIObjectsSimRailSimRailTrainPositionType{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSimRailSimRailTrainPositionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSimRailSimRailTrainPositionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

