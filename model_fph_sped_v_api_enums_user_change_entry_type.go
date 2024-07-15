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

// FPHSpedVAPIEnumsUserChangeEntryType   0 = ChangedName  1 = JoinedSpedition  2 = DismissedBySpedition  3 = Quitted  4 = ChangedProfileData  5 = FoundedSpedition  -1 = NotSet
type FPHSpedVAPIEnumsUserChangeEntryType int32

// List of FPH.SpedV.API.Enums.UserChangeEntryType
const (
	_0 FPHSpedVAPIEnumsUserChangeEntryType = 0
	_1 FPHSpedVAPIEnumsUserChangeEntryType = 1
	_2 FPHSpedVAPIEnumsUserChangeEntryType = 2
	_3 FPHSpedVAPIEnumsUserChangeEntryType = 3
	_4 FPHSpedVAPIEnumsUserChangeEntryType = 4
	_5 FPHSpedVAPIEnumsUserChangeEntryType = 5
	_MINUS_1 FPHSpedVAPIEnumsUserChangeEntryType = -1
)

// All allowed values of FPHSpedVAPIEnumsUserChangeEntryType enum
var AllowedFPHSpedVAPIEnumsUserChangeEntryTypeEnumValues = []FPHSpedVAPIEnumsUserChangeEntryType{
	0,
	1,
	2,
	3,
	4,
	5,
	-1,
}

func (v *FPHSpedVAPIEnumsUserChangeEntryType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FPHSpedVAPIEnumsUserChangeEntryType(value)
	for _, existing := range AllowedFPHSpedVAPIEnumsUserChangeEntryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FPHSpedVAPIEnumsUserChangeEntryType", value)
}

// NewFPHSpedVAPIEnumsUserChangeEntryTypeFromValue returns a pointer to a valid FPHSpedVAPIEnumsUserChangeEntryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFPHSpedVAPIEnumsUserChangeEntryTypeFromValue(v int32) (*FPHSpedVAPIEnumsUserChangeEntryType, error) {
	ev := FPHSpedVAPIEnumsUserChangeEntryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FPHSpedVAPIEnumsUserChangeEntryType: valid values are %v", v, AllowedFPHSpedVAPIEnumsUserChangeEntryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FPHSpedVAPIEnumsUserChangeEntryType) IsValid() bool {
	for _, existing := range AllowedFPHSpedVAPIEnumsUserChangeEntryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FPH.SpedV.API.Enums.UserChangeEntryType value
func (v FPHSpedVAPIEnumsUserChangeEntryType) Ptr() *FPHSpedVAPIEnumsUserChangeEntryType {
	return &v
}

type NullableFPHSpedVAPIEnumsUserChangeEntryType struct {
	value *FPHSpedVAPIEnumsUserChangeEntryType
	isSet bool
}

func (v NullableFPHSpedVAPIEnumsUserChangeEntryType) Get() *FPHSpedVAPIEnumsUserChangeEntryType {
	return v.value
}

func (v *NullableFPHSpedVAPIEnumsUserChangeEntryType) Set(val *FPHSpedVAPIEnumsUserChangeEntryType) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIEnumsUserChangeEntryType) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIEnumsUserChangeEntryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIEnumsUserChangeEntryType(val *FPHSpedVAPIEnumsUserChangeEntryType) *NullableFPHSpedVAPIEnumsUserChangeEntryType {
	return &NullableFPHSpedVAPIEnumsUserChangeEntryType{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIEnumsUserChangeEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIEnumsUserChangeEntryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

