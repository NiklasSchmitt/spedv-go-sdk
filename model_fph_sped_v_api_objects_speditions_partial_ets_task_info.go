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

// checks if the FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo{}

// FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo struct for FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo
type FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo struct {
	Id int32 `json:"id"`
	VisibleID NullableString `json:"visibleID"`
	//   0 = InDrive  1 = Done  2 = Settled  3 = Fail  4 = AdminCheck  5 = Paused  6 = Cancelled  7 = Invalid  -1 = NotAvaliable
	State FPHSpedVAPIEnumsETSTaskState `json:"state"`
	Start NullableString `json:"start"`
	Dest NullableString `json:"dest"`
	Freight NullableString `json:"freight"`
	HasScreenshot int32 `json:"hasScreenshot"`
	IsDeductable bool `json:"isDeductable"`
	DdCleaned bool `json:"ddCleaned"`
}

type _FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo

// NewFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo instantiates a new FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo(id int32, visibleID NullableString, state FPHSpedVAPIEnumsETSTaskState, start NullableString, dest NullableString, freight NullableString, hasScreenshot int32, isDeductable bool, ddCleaned bool) *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo {
	this := FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo{}
	this.Id = id
	this.VisibleID = visibleID
	this.State = state
	this.Start = start
	this.Dest = dest
	this.Freight = freight
	this.HasScreenshot = hasScreenshot
	this.IsDeductable = isDeductable
	this.DdCleaned = ddCleaned
	return &this
}

// NewFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfoWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfoWithDefaults() *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo {
	this := FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetId(v int32) {
	o.Id = v
}

// GetVisibleID returns the VisibleID field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetVisibleID() string {
	if o == nil || o.VisibleID.Get() == nil {
		var ret string
		return ret
	}

	return *o.VisibleID.Get()
}

// GetVisibleIDOk returns a tuple with the VisibleID field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetVisibleIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VisibleID.Get(), o.VisibleID.IsSet()
}

// SetVisibleID sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetVisibleID(v string) {
	o.VisibleID.Set(&v)
}

// GetState returns the State field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetState() FPHSpedVAPIEnumsETSTaskState {
	if o == nil {
		var ret FPHSpedVAPIEnumsETSTaskState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetStateOk() (*FPHSpedVAPIEnumsETSTaskState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetState(v FPHSpedVAPIEnumsETSTaskState) {
	o.State = v
}

// GetStart returns the Start field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetStart() string {
	if o == nil || o.Start.Get() == nil {
		var ret string
		return ret
	}

	return *o.Start.Get()
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Start.Get(), o.Start.IsSet()
}

// SetStart sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetStart(v string) {
	o.Start.Set(&v)
}

// GetDest returns the Dest field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetDest() string {
	if o == nil || o.Dest.Get() == nil {
		var ret string
		return ret
	}

	return *o.Dest.Get()
}

// GetDestOk returns a tuple with the Dest field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetDestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dest.Get(), o.Dest.IsSet()
}

// SetDest sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetDest(v string) {
	o.Dest.Set(&v)
}

// GetFreight returns the Freight field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetFreight() string {
	if o == nil || o.Freight.Get() == nil {
		var ret string
		return ret
	}

	return *o.Freight.Get()
}

// GetFreightOk returns a tuple with the Freight field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetFreightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Freight.Get(), o.Freight.IsSet()
}

// SetFreight sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetFreight(v string) {
	o.Freight.Set(&v)
}

// GetHasScreenshot returns the HasScreenshot field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetHasScreenshot() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HasScreenshot
}

// GetHasScreenshotOk returns a tuple with the HasScreenshot field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetHasScreenshotOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasScreenshot, true
}

// SetHasScreenshot sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetHasScreenshot(v int32) {
	o.HasScreenshot = v
}

// GetIsDeductable returns the IsDeductable field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetIsDeductable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeductable
}

// GetIsDeductableOk returns a tuple with the IsDeductable field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetIsDeductableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeductable, true
}

// SetIsDeductable sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetIsDeductable(v bool) {
	o.IsDeductable = v
}

// GetDdCleaned returns the DdCleaned field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetDdCleaned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DdCleaned
}

// GetDdCleanedOk returns a tuple with the DdCleaned field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) GetDdCleanedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DdCleaned, true
}

// SetDdCleaned sets field value
func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) SetDdCleaned(v bool) {
	o.DdCleaned = v
}

func (o FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["visibleID"] = o.VisibleID.Get()
	toSerialize["state"] = o.State
	toSerialize["start"] = o.Start.Get()
	toSerialize["dest"] = o.Dest.Get()
	toSerialize["freight"] = o.Freight.Get()
	toSerialize["hasScreenshot"] = o.HasScreenshot
	toSerialize["isDeductable"] = o.IsDeductable
	toSerialize["ddCleaned"] = o.DdCleaned
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"visibleID",
		"state",
		"start",
		"dest",
		"freight",
		"hasScreenshot",
		"isDeductable",
		"ddCleaned",
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

	varFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo := _FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo(varFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo)

	return err
}

type NullableFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo struct {
	value *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) Get() *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) Set(val *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo(val *FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) *NullableFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo {
	return &NullableFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


