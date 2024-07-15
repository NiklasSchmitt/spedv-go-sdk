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

// FPHSpedVAPIEnumsBankAccountType   0 = System  1 = Spedition  2 = User  -1 = NotSet
type FPHSpedVAPIEnumsBankAccountType int32

// List of FPH.SpedV.API.Enums.BankAccountType
const (
	_0 FPHSpedVAPIEnumsBankAccountType = 0
	_1 FPHSpedVAPIEnumsBankAccountType = 1
	_2 FPHSpedVAPIEnumsBankAccountType = 2
	_MINUS_1 FPHSpedVAPIEnumsBankAccountType = -1
)

// All allowed values of FPHSpedVAPIEnumsBankAccountType enum
var AllowedFPHSpedVAPIEnumsBankAccountTypeEnumValues = []FPHSpedVAPIEnumsBankAccountType{
	0,
	1,
	2,
	-1,
}

func (v *FPHSpedVAPIEnumsBankAccountType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FPHSpedVAPIEnumsBankAccountType(value)
	for _, existing := range AllowedFPHSpedVAPIEnumsBankAccountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FPHSpedVAPIEnumsBankAccountType", value)
}

// NewFPHSpedVAPIEnumsBankAccountTypeFromValue returns a pointer to a valid FPHSpedVAPIEnumsBankAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFPHSpedVAPIEnumsBankAccountTypeFromValue(v int32) (*FPHSpedVAPIEnumsBankAccountType, error) {
	ev := FPHSpedVAPIEnumsBankAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FPHSpedVAPIEnumsBankAccountType: valid values are %v", v, AllowedFPHSpedVAPIEnumsBankAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FPHSpedVAPIEnumsBankAccountType) IsValid() bool {
	for _, existing := range AllowedFPHSpedVAPIEnumsBankAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FPH.SpedV.API.Enums.BankAccountType value
func (v FPHSpedVAPIEnumsBankAccountType) Ptr() *FPHSpedVAPIEnumsBankAccountType {
	return &v
}

type NullableFPHSpedVAPIEnumsBankAccountType struct {
	value *FPHSpedVAPIEnumsBankAccountType
	isSet bool
}

func (v NullableFPHSpedVAPIEnumsBankAccountType) Get() *FPHSpedVAPIEnumsBankAccountType {
	return v.value
}

func (v *NullableFPHSpedVAPIEnumsBankAccountType) Set(val *FPHSpedVAPIEnumsBankAccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIEnumsBankAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIEnumsBankAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIEnumsBankAccountType(val *FPHSpedVAPIEnumsBankAccountType) *NullableFPHSpedVAPIEnumsBankAccountType {
	return &NullableFPHSpedVAPIEnumsBankAccountType{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIEnumsBankAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIEnumsBankAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

