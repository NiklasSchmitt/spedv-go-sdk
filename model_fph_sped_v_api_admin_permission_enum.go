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

// FPHSpedVAPIAdminPermissionEnum   0 = None  1 = SeeLiveData  2 = HandleScreenshots  4 = HandleTaskCheck  8 = TicketSystem  16 = ViewData  32 = ChangeData  64 = BanUser  128 = MapImport  256 = KontorManagement  512 = IsManagement
type FPHSpedVAPIAdminPermissionEnum int32

// List of FPH.SpedV.API.AdminPermissionEnum
const (
	_0 FPHSpedVAPIAdminPermissionEnum = 0
	_1 FPHSpedVAPIAdminPermissionEnum = 1
	_2 FPHSpedVAPIAdminPermissionEnum = 2
	_4 FPHSpedVAPIAdminPermissionEnum = 4
	_8 FPHSpedVAPIAdminPermissionEnum = 8
	_16 FPHSpedVAPIAdminPermissionEnum = 16
	_32 FPHSpedVAPIAdminPermissionEnum = 32
	_64 FPHSpedVAPIAdminPermissionEnum = 64
	_128 FPHSpedVAPIAdminPermissionEnum = 128
	_256 FPHSpedVAPIAdminPermissionEnum = 256
	_512 FPHSpedVAPIAdminPermissionEnum = 512
)

// All allowed values of FPHSpedVAPIAdminPermissionEnum enum
var AllowedFPHSpedVAPIAdminPermissionEnumEnumValues = []FPHSpedVAPIAdminPermissionEnum{
	0,
	1,
	2,
	4,
	8,
	16,
	32,
	64,
	128,
	256,
	512,
}

func (v *FPHSpedVAPIAdminPermissionEnum) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FPHSpedVAPIAdminPermissionEnum(value)
	for _, existing := range AllowedFPHSpedVAPIAdminPermissionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FPHSpedVAPIAdminPermissionEnum", value)
}

// NewFPHSpedVAPIAdminPermissionEnumFromValue returns a pointer to a valid FPHSpedVAPIAdminPermissionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFPHSpedVAPIAdminPermissionEnumFromValue(v int32) (*FPHSpedVAPIAdminPermissionEnum, error) {
	ev := FPHSpedVAPIAdminPermissionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FPHSpedVAPIAdminPermissionEnum: valid values are %v", v, AllowedFPHSpedVAPIAdminPermissionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FPHSpedVAPIAdminPermissionEnum) IsValid() bool {
	for _, existing := range AllowedFPHSpedVAPIAdminPermissionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FPH.SpedV.API.AdminPermissionEnum value
func (v FPHSpedVAPIAdminPermissionEnum) Ptr() *FPHSpedVAPIAdminPermissionEnum {
	return &v
}

type NullableFPHSpedVAPIAdminPermissionEnum struct {
	value *FPHSpedVAPIAdminPermissionEnum
	isSet bool
}

func (v NullableFPHSpedVAPIAdminPermissionEnum) Get() *FPHSpedVAPIAdminPermissionEnum {
	return v.value
}

func (v *NullableFPHSpedVAPIAdminPermissionEnum) Set(val *FPHSpedVAPIAdminPermissionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIAdminPermissionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIAdminPermissionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIAdminPermissionEnum(val *FPHSpedVAPIAdminPermissionEnum) *NullableFPHSpedVAPIAdminPermissionEnum {
	return &NullableFPHSpedVAPIAdminPermissionEnum{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIAdminPermissionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIAdminPermissionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
