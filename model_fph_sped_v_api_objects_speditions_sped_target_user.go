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

// checks if the FPHSpedVAPIObjectsSpeditionsSpedTargetUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsSpeditionsSpedTargetUser{}

// FPHSpedVAPIObjectsSpeditionsSpedTargetUser struct for FPHSpedVAPIObjectsSpeditionsSpedTargetUser
type FPHSpedVAPIObjectsSpeditionsSpedTargetUser struct {
	UserObject NullableFPHSpedVAPIObjectsUsersUser `json:"userObject"`
	User NullableString `json:"user"`
	// Deprecated
	Amount int32 `json:"amount"`
	AmountReached float64 `json:"amountReached"`
	Unit NullableString `json:"unit"`
	// Deprecated
	Value NullableString `json:"value"`
	VisAmountReached NullableString `json:"visAmountReached"`
	Max int32 `json:"max"`
}

type _FPHSpedVAPIObjectsSpeditionsSpedTargetUser FPHSpedVAPIObjectsSpeditionsSpedTargetUser

// NewFPHSpedVAPIObjectsSpeditionsSpedTargetUser instantiates a new FPHSpedVAPIObjectsSpeditionsSpedTargetUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsSpeditionsSpedTargetUser(userObject NullableFPHSpedVAPIObjectsUsersUser, user NullableString, amount int32, amountReached float64, unit NullableString, value NullableString, visAmountReached NullableString, max int32) *FPHSpedVAPIObjectsSpeditionsSpedTargetUser {
	this := FPHSpedVAPIObjectsSpeditionsSpedTargetUser{}
	this.UserObject = userObject
	this.User = user
	this.Amount = amount
	this.AmountReached = amountReached
	this.Unit = unit
	this.Value = value
	this.VisAmountReached = visAmountReached
	this.Max = max
	return &this
}

// NewFPHSpedVAPIObjectsSpeditionsSpedTargetUserWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsSpedTargetUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsSpeditionsSpedTargetUserWithDefaults() *FPHSpedVAPIObjectsSpeditionsSpedTargetUser {
	this := FPHSpedVAPIObjectsSpeditionsSpedTargetUser{}
	return &this
}

// GetUserObject returns the UserObject field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsUsersUser will be returned
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetUserObject() FPHSpedVAPIObjectsUsersUser {
	if o == nil || o.UserObject.Get() == nil {
		var ret FPHSpedVAPIObjectsUsersUser
		return ret
	}

	return *o.UserObject.Get()
}

// GetUserObjectOk returns a tuple with the UserObject field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetUserObjectOk() (*FPHSpedVAPIObjectsUsersUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserObject.Get(), o.UserObject.IsSet()
}

// SetUserObject sets field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetUserObject(v FPHSpedVAPIObjectsUsersUser) {
	o.UserObject.Set(&v)
}

// GetUser returns the User field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetUser() string {
	if o == nil || o.User.Get() == nil {
		var ret string
		return ret
	}

	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// SetUser sets field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetUser(v string) {
	o.User.Set(&v)
}

// GetAmount returns the Amount field value
// Deprecated
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
// Deprecated
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetAmount(v int32) {
	o.Amount = v
}

// GetAmountReached returns the AmountReached field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetAmountReached() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AmountReached
}

// GetAmountReachedOk returns a tuple with the AmountReached field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetAmountReachedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountReached, true
}

// SetAmountReached sets field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetAmountReached(v float64) {
	o.AmountReached = v
}

// GetUnit returns the Unit field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetUnit() string {
	if o == nil || o.Unit.Get() == nil {
		var ret string
		return ret
	}

	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// SetUnit sets field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetUnit(v string) {
	o.Unit.Set(&v)
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
// Deprecated
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
// Deprecated
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetValue(v string) {
	o.Value.Set(&v)
}

// GetVisAmountReached returns the VisAmountReached field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetVisAmountReached() string {
	if o == nil || o.VisAmountReached.Get() == nil {
		var ret string
		return ret
	}

	return *o.VisAmountReached.Get()
}

// GetVisAmountReachedOk returns a tuple with the VisAmountReached field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetVisAmountReachedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VisAmountReached.Get(), o.VisAmountReached.IsSet()
}

// SetVisAmountReached sets field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetVisAmountReached(v string) {
	o.VisAmountReached.Set(&v)
}

// GetMax returns the Max field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) GetMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) SetMax(v int32) {
	o.Max = v
}

func (o FPHSpedVAPIObjectsSpeditionsSpedTargetUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsSpeditionsSpedTargetUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userObject"] = o.UserObject.Get()
	toSerialize["user"] = o.User.Get()
	toSerialize["amount"] = o.Amount
	toSerialize["amountReached"] = o.AmountReached
	toSerialize["unit"] = o.Unit.Get()
	toSerialize["value"] = o.Value.Get()
	toSerialize["visAmountReached"] = o.VisAmountReached.Get()
	toSerialize["max"] = o.Max
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userObject",
		"user",
		"amount",
		"amountReached",
		"unit",
		"value",
		"visAmountReached",
		"max",
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

	varFPHSpedVAPIObjectsSpeditionsSpedTargetUser := _FPHSpedVAPIObjectsSpeditionsSpedTargetUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsSpeditionsSpedTargetUser)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsSpeditionsSpedTargetUser(varFPHSpedVAPIObjectsSpeditionsSpedTargetUser)

	return err
}

type NullableFPHSpedVAPIObjectsSpeditionsSpedTargetUser struct {
	value *FPHSpedVAPIObjectsSpeditionsSpedTargetUser
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSpeditionsSpedTargetUser) Get() *FPHSpedVAPIObjectsSpeditionsSpedTargetUser {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsSpedTargetUser) Set(val *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSpeditionsSpedTargetUser) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsSpedTargetUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSpeditionsSpedTargetUser(val *FPHSpedVAPIObjectsSpeditionsSpedTargetUser) *NullableFPHSpedVAPIObjectsSpeditionsSpedTargetUser {
	return &NullableFPHSpedVAPIObjectsSpeditionsSpedTargetUser{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSpeditionsSpedTargetUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsSpedTargetUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


