/*
SpedV API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry{}

// FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry struct for FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry
type FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry struct {
	Id int32 `json:"id"`
	User NullableFPHSpedVAPIObjectsUsersUser `json:"user"`
	Timestamp time.Time `json:"timestamp"`
	//   0 = TaskDeducted  1 = TaskRejected  2 = RankChanged  3 = RankDeleted  4 = UserRankChanged  5 = UserDismissed  6 = SpeditionSettingChanged  7 = ApplicationTextChanged  8 = ApplicationAcknowledged  9 = ApplicationRejected  10 = BranchBought  11 = BranchExpandedTruckParkplace  12 = BranchShrinkedTruckParkplace  13 = BranchDeleted  14 = TruckBought  15 = TruckSold  16 = TruckLicensePlateChanged  17 = BoniAdded  18 = BoniDeleted  19 = RankAddedPermission  20 = RankDeletedPermission  21 = BranchExpandedTrailerParkplace  22 = BranchShrinkedTrailerParkplace  23 = BranchExpandedMaintenancePlace  24 = BranchShrinkedMaintenancePlace  25 = TruckTransfered  26 = TruckMaintenancePlanned  27 = TruckMaintenanceAborted  -1 = NotSet
	Type FPHSpedVAPIEnumsSpeditionChangeEntryType `json:"type"`
	Comment NullableString `json:"comment"`
}

type _FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry

// NewFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry instantiates a new FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry(id int32, user NullableFPHSpedVAPIObjectsUsersUser, timestamp time.Time, type_ FPHSpedVAPIEnumsSpeditionChangeEntryType, comment NullableString) *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry {
	this := FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry{}
	this.Id = id
	this.User = user
	this.Timestamp = timestamp
	this.Type = type_
	this.Comment = comment
	return &this
}

// NewFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntryWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntryWithDefaults() *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry {
	this := FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) SetId(v int32) {
	o.Id = v
}

// GetUser returns the User field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsUsersUser will be returned
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetUser() FPHSpedVAPIObjectsUsersUser {
	if o == nil || o.User.Get() == nil {
		var ret FPHSpedVAPIObjectsUsersUser
		return ret
	}

	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetUserOk() (*FPHSpedVAPIObjectsUsersUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// SetUser sets field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) SetUser(v FPHSpedVAPIObjectsUsersUser) {
	o.User.Set(&v)
}

// GetTimestamp returns the Timestamp field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetType returns the Type field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetType() FPHSpedVAPIEnumsSpeditionChangeEntryType {
	if o == nil {
		var ret FPHSpedVAPIEnumsSpeditionChangeEntryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetTypeOk() (*FPHSpedVAPIEnumsSpeditionChangeEntryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) SetType(v FPHSpedVAPIEnumsSpeditionChangeEntryType) {
	o.Type = v
}

// GetComment returns the Comment field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}

	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// SetComment sets field value
func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) SetComment(v string) {
	o.Comment.Set(&v)
}

func (o FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user"] = o.User.Get()
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["type"] = o.Type
	toSerialize["comment"] = o.Comment.Get()
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user",
		"timestamp",
		"type",
		"comment",
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

	varFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry := _FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry(varFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry)

	return err
}

type NullableFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry struct {
	value *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) Get() *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) Set(val *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry(val *FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) *NullableFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry {
	return &NullableFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

