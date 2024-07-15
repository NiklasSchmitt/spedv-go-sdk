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

// checks if the FPHSpedVAPIObjectsUsersUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsUsersUser{}

// FPHSpedVAPIObjectsUsersUser struct for FPHSpedVAPIObjectsUsersUser
type FPHSpedVAPIObjectsUsersUser struct {
	Id int32 `json:"id"`
	UserName NullableString `json:"userName"`
	ProfilePicURL NullableString `json:"profilePicURL"`
	Spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition `json:"spedition"`
	InSped bool `json:"inSped"`
	FormatName NullableString `json:"formatName"`
}

type _FPHSpedVAPIObjectsUsersUser FPHSpedVAPIObjectsUsersUser

// NewFPHSpedVAPIObjectsUsersUser instantiates a new FPHSpedVAPIObjectsUsersUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsUsersUser(id int32, userName NullableString, profilePicURL NullableString, spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, inSped bool, formatName NullableString) *FPHSpedVAPIObjectsUsersUser {
	this := FPHSpedVAPIObjectsUsersUser{}
	this.Id = id
	this.UserName = userName
	this.ProfilePicURL = profilePicURL
	this.Spedition = spedition
	this.InSped = inSped
	this.FormatName = formatName
	return &this
}

// NewFPHSpedVAPIObjectsUsersUserWithDefaults instantiates a new FPHSpedVAPIObjectsUsersUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsUsersUserWithDefaults() *FPHSpedVAPIObjectsUsersUser {
	this := FPHSpedVAPIObjectsUsersUser{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsUsersUser) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsUsersUser) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsUsersUser) SetId(v int32) {
	o.Id = v
}

// GetUserName returns the UserName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsUsersUser) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}

	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUser) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// SetUserName sets field value
func (o *FPHSpedVAPIObjectsUsersUser) SetUserName(v string) {
	o.UserName.Set(&v)
}

// GetProfilePicURL returns the ProfilePicURL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsUsersUser) GetProfilePicURL() string {
	if o == nil || o.ProfilePicURL.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProfilePicURL.Get()
}

// GetProfilePicURLOk returns a tuple with the ProfilePicURL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUser) GetProfilePicURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilePicURL.Get(), o.ProfilePicURL.IsSet()
}

// SetProfilePicURL sets field value
func (o *FPHSpedVAPIObjectsUsersUser) SetProfilePicURL(v string) {
	o.ProfilePicURL.Set(&v)
}

// GetSpedition returns the Spedition field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsSpeditionsSpedition will be returned
func (o *FPHSpedVAPIObjectsUsersUser) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition {
	if o == nil || o.Spedition.Get() == nil {
		var ret FPHSpedVAPIObjectsSpeditionsSpedition
		return ret
	}

	return *o.Spedition.Get()
}

// GetSpeditionOk returns a tuple with the Spedition field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUser) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Spedition.Get(), o.Spedition.IsSet()
}

// SetSpedition sets field value
func (o *FPHSpedVAPIObjectsUsersUser) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition) {
	o.Spedition.Set(&v)
}

// GetInSped returns the InSped field value
func (o *FPHSpedVAPIObjectsUsersUser) GetInSped() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InSped
}

// GetInSpedOk returns a tuple with the InSped field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsUsersUser) GetInSpedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InSped, true
}

// SetInSped sets field value
func (o *FPHSpedVAPIObjectsUsersUser) SetInSped(v bool) {
	o.InSped = v
}

// GetFormatName returns the FormatName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsUsersUser) GetFormatName() string {
	if o == nil || o.FormatName.Get() == nil {
		var ret string
		return ret
	}

	return *o.FormatName.Get()
}

// GetFormatNameOk returns a tuple with the FormatName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUser) GetFormatNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormatName.Get(), o.FormatName.IsSet()
}

// SetFormatName sets field value
func (o *FPHSpedVAPIObjectsUsersUser) SetFormatName(v string) {
	o.FormatName.Set(&v)
}

func (o FPHSpedVAPIObjectsUsersUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsUsersUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["userName"] = o.UserName.Get()
	toSerialize["profilePicURL"] = o.ProfilePicURL.Get()
	toSerialize["spedition"] = o.Spedition.Get()
	toSerialize["inSped"] = o.InSped
	toSerialize["formatName"] = o.FormatName.Get()
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsUsersUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"userName",
		"profilePicURL",
		"spedition",
		"inSped",
		"formatName",
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

	varFPHSpedVAPIObjectsUsersUser := _FPHSpedVAPIObjectsUsersUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsUsersUser)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsUsersUser(varFPHSpedVAPIObjectsUsersUser)

	return err
}

type NullableFPHSpedVAPIObjectsUsersUser struct {
	value *FPHSpedVAPIObjectsUsersUser
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsUsersUser) Get() *FPHSpedVAPIObjectsUsersUser {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsUsersUser) Set(val *FPHSpedVAPIObjectsUsersUser) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsUsersUser) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsUsersUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsUsersUser(val *FPHSpedVAPIObjectsUsersUser) *NullableFPHSpedVAPIObjectsUsersUser {
	return &NullableFPHSpedVAPIObjectsUsersUser{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsUsersUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsUsersUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

