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

// checks if the FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance{}

// FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance struct for FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance
type FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance struct {
	//   0 = Engine  1 = OszilationDamper  2 = Stabilizer  3 = StoneChip  4 = Transmission  5 = Wishbone  6 = BrakePads  7 = BrakeDiscs  8 = EngineMaintenance  9 = TireChange  10 = MainCheck  11 = SafetyCheck  12 = SaddlePlate  13 = AirPressureUnit  14 = Alternator  15 = BrakeVentil  -1 = NotSet
	Kind FPHSpedVAPIEnumsMaintenanceKind `json:"kind"`
	ExternalTime string `json:"externalTime"`
	InternalTime string `json:"internalTime"`
	ExternalCost int32 `json:"externalCost"`
	InternalCost int32 `json:"internalCost"`
	NeededImmediately bool `json:"neededImmediately"`
	NeededTillKM NullableInt32 `json:"neededTillKM"`
	NeededTillDate NullableTime `json:"neededTillDate"`
	RepeatTimeSpan NullableString `json:"repeatTimeSpan"`
	RepeatKM NullableInt32 `json:"repeatKM"`
	CurrentKM int32 `json:"currentKM"`
	//   0 = ETS2  1 = ATS  -1 = NotSet
	Game FPHSpedVAPIEnumsGameEnum `json:"game"`
	KmSinceLast NullableInt32 `json:"kmSinceLast"`
	TimeSinceLast NullableString `json:"timeSinceLast"`
	VisibleRemaining NullableString `json:"visibleRemaining"`
	RepeatSpan float64 `json:"repeatSpan"`
	SpanSinceLast float64 `json:"spanSinceLast"`
	IsNeeded bool `json:"isNeeded"`
	IsOverdue bool `json:"isOverdue"`
}

type _FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance

// NewFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance instantiates a new FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance(kind FPHSpedVAPIEnumsMaintenanceKind, externalTime string, internalTime string, externalCost int32, internalCost int32, neededImmediately bool, neededTillKM NullableInt32, neededTillDate NullableTime, repeatTimeSpan NullableString, repeatKM NullableInt32, currentKM int32, game FPHSpedVAPIEnumsGameEnum, kmSinceLast NullableInt32, timeSinceLast NullableString, visibleRemaining NullableString, repeatSpan float64, spanSinceLast float64, isNeeded bool, isOverdue bool) *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance {
	this := FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance{}
	this.Kind = kind
	this.ExternalTime = externalTime
	this.InternalTime = internalTime
	this.ExternalCost = externalCost
	this.InternalCost = internalCost
	this.NeededImmediately = neededImmediately
	this.NeededTillKM = neededTillKM
	this.NeededTillDate = neededTillDate
	this.RepeatTimeSpan = repeatTimeSpan
	this.RepeatKM = repeatKM
	this.CurrentKM = currentKM
	this.Game = game
	this.KmSinceLast = kmSinceLast
	this.TimeSinceLast = timeSinceLast
	this.VisibleRemaining = visibleRemaining
	this.RepeatSpan = repeatSpan
	this.SpanSinceLast = spanSinceLast
	this.IsNeeded = isNeeded
	this.IsOverdue = isOverdue
	return &this
}

// NewFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenanceWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenanceWithDefaults() *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance {
	this := FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance{}
	return &this
}

// GetKind returns the Kind field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetKind() FPHSpedVAPIEnumsMaintenanceKind {
	if o == nil {
		var ret FPHSpedVAPIEnumsMaintenanceKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetKindOk() (*FPHSpedVAPIEnumsMaintenanceKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetKind(v FPHSpedVAPIEnumsMaintenanceKind) {
	o.Kind = v
}

// GetExternalTime returns the ExternalTime field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetExternalTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalTime
}

// GetExternalTimeOk returns a tuple with the ExternalTime field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetExternalTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalTime, true
}

// SetExternalTime sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetExternalTime(v string) {
	o.ExternalTime = v
}

// GetInternalTime returns the InternalTime field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetInternalTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InternalTime
}

// GetInternalTimeOk returns a tuple with the InternalTime field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetInternalTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternalTime, true
}

// SetInternalTime sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetInternalTime(v string) {
	o.InternalTime = v
}

// GetExternalCost returns the ExternalCost field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetExternalCost() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExternalCost
}

// GetExternalCostOk returns a tuple with the ExternalCost field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetExternalCostOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalCost, true
}

// SetExternalCost sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetExternalCost(v int32) {
	o.ExternalCost = v
}

// GetInternalCost returns the InternalCost field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetInternalCost() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InternalCost
}

// GetInternalCostOk returns a tuple with the InternalCost field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetInternalCostOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternalCost, true
}

// SetInternalCost sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetInternalCost(v int32) {
	o.InternalCost = v
}

// GetNeededImmediately returns the NeededImmediately field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetNeededImmediately() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NeededImmediately
}

// GetNeededImmediatelyOk returns a tuple with the NeededImmediately field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetNeededImmediatelyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeededImmediately, true
}

// SetNeededImmediately sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetNeededImmediately(v bool) {
	o.NeededImmediately = v
}

// GetNeededTillKM returns the NeededTillKM field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetNeededTillKM() int32 {
	if o == nil || o.NeededTillKM.Get() == nil {
		var ret int32
		return ret
	}

	return *o.NeededTillKM.Get()
}

// GetNeededTillKMOk returns a tuple with the NeededTillKM field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetNeededTillKMOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NeededTillKM.Get(), o.NeededTillKM.IsSet()
}

// SetNeededTillKM sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetNeededTillKM(v int32) {
	o.NeededTillKM.Set(&v)
}

// GetNeededTillDate returns the NeededTillDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetNeededTillDate() time.Time {
	if o == nil || o.NeededTillDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.NeededTillDate.Get()
}

// GetNeededTillDateOk returns a tuple with the NeededTillDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetNeededTillDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.NeededTillDate.Get(), o.NeededTillDate.IsSet()
}

// SetNeededTillDate sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetNeededTillDate(v time.Time) {
	o.NeededTillDate.Set(&v)
}

// GetRepeatTimeSpan returns the RepeatTimeSpan field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetRepeatTimeSpan() string {
	if o == nil || o.RepeatTimeSpan.Get() == nil {
		var ret string
		return ret
	}

	return *o.RepeatTimeSpan.Get()
}

// GetRepeatTimeSpanOk returns a tuple with the RepeatTimeSpan field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetRepeatTimeSpanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepeatTimeSpan.Get(), o.RepeatTimeSpan.IsSet()
}

// SetRepeatTimeSpan sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetRepeatTimeSpan(v string) {
	o.RepeatTimeSpan.Set(&v)
}

// GetRepeatKM returns the RepeatKM field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetRepeatKM() int32 {
	if o == nil || o.RepeatKM.Get() == nil {
		var ret int32
		return ret
	}

	return *o.RepeatKM.Get()
}

// GetRepeatKMOk returns a tuple with the RepeatKM field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetRepeatKMOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepeatKM.Get(), o.RepeatKM.IsSet()
}

// SetRepeatKM sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetRepeatKM(v int32) {
	o.RepeatKM.Set(&v)
}

// GetCurrentKM returns the CurrentKM field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetCurrentKM() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentKM
}

// GetCurrentKMOk returns a tuple with the CurrentKM field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetCurrentKMOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentKM, true
}

// SetCurrentKM sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetCurrentKM(v int32) {
	o.CurrentKM = v
}

// GetGame returns the Game field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetGame() FPHSpedVAPIEnumsGameEnum {
	if o == nil {
		var ret FPHSpedVAPIEnumsGameEnum
		return ret
	}

	return o.Game
}

// GetGameOk returns a tuple with the Game field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetGameOk() (*FPHSpedVAPIEnumsGameEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Game, true
}

// SetGame sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetGame(v FPHSpedVAPIEnumsGameEnum) {
	o.Game = v
}

// GetKmSinceLast returns the KmSinceLast field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetKmSinceLast() int32 {
	if o == nil || o.KmSinceLast.Get() == nil {
		var ret int32
		return ret
	}

	return *o.KmSinceLast.Get()
}

// GetKmSinceLastOk returns a tuple with the KmSinceLast field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetKmSinceLastOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.KmSinceLast.Get(), o.KmSinceLast.IsSet()
}

// SetKmSinceLast sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetKmSinceLast(v int32) {
	o.KmSinceLast.Set(&v)
}

// GetTimeSinceLast returns the TimeSinceLast field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetTimeSinceLast() string {
	if o == nil || o.TimeSinceLast.Get() == nil {
		var ret string
		return ret
	}

	return *o.TimeSinceLast.Get()
}

// GetTimeSinceLastOk returns a tuple with the TimeSinceLast field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetTimeSinceLastOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeSinceLast.Get(), o.TimeSinceLast.IsSet()
}

// SetTimeSinceLast sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetTimeSinceLast(v string) {
	o.TimeSinceLast.Set(&v)
}

// GetVisibleRemaining returns the VisibleRemaining field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetVisibleRemaining() string {
	if o == nil || o.VisibleRemaining.Get() == nil {
		var ret string
		return ret
	}

	return *o.VisibleRemaining.Get()
}

// GetVisibleRemainingOk returns a tuple with the VisibleRemaining field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetVisibleRemainingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VisibleRemaining.Get(), o.VisibleRemaining.IsSet()
}

// SetVisibleRemaining sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetVisibleRemaining(v string) {
	o.VisibleRemaining.Set(&v)
}

// GetRepeatSpan returns the RepeatSpan field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetRepeatSpan() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.RepeatSpan
}

// GetRepeatSpanOk returns a tuple with the RepeatSpan field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetRepeatSpanOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepeatSpan, true
}

// SetRepeatSpan sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetRepeatSpan(v float64) {
	o.RepeatSpan = v
}

// GetSpanSinceLast returns the SpanSinceLast field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetSpanSinceLast() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpanSinceLast
}

// GetSpanSinceLastOk returns a tuple with the SpanSinceLast field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetSpanSinceLastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanSinceLast, true
}

// SetSpanSinceLast sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetSpanSinceLast(v float64) {
	o.SpanSinceLast = v
}

// GetIsNeeded returns the IsNeeded field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetIsNeeded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsNeeded
}

// GetIsNeededOk returns a tuple with the IsNeeded field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetIsNeededOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsNeeded, true
}

// SetIsNeeded sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetIsNeeded(v bool) {
	o.IsNeeded = v
}

// GetIsOverdue returns the IsOverdue field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetIsOverdue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsOverdue
}

// GetIsOverdueOk returns a tuple with the IsOverdue field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) GetIsOverdueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsOverdue, true
}

// SetIsOverdue sets field value
func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) SetIsOverdue(v bool) {
	o.IsOverdue = v
}

func (o FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["externalTime"] = o.ExternalTime
	toSerialize["internalTime"] = o.InternalTime
	toSerialize["externalCost"] = o.ExternalCost
	toSerialize["internalCost"] = o.InternalCost
	toSerialize["neededImmediately"] = o.NeededImmediately
	toSerialize["neededTillKM"] = o.NeededTillKM.Get()
	toSerialize["neededTillDate"] = o.NeededTillDate.Get()
	toSerialize["repeatTimeSpan"] = o.RepeatTimeSpan.Get()
	toSerialize["repeatKM"] = o.RepeatKM.Get()
	toSerialize["currentKM"] = o.CurrentKM
	toSerialize["game"] = o.Game
	toSerialize["kmSinceLast"] = o.KmSinceLast.Get()
	toSerialize["timeSinceLast"] = o.TimeSinceLast.Get()
	toSerialize["visibleRemaining"] = o.VisibleRemaining.Get()
	toSerialize["repeatSpan"] = o.RepeatSpan
	toSerialize["spanSinceLast"] = o.SpanSinceLast
	toSerialize["isNeeded"] = o.IsNeeded
	toSerialize["isOverdue"] = o.IsOverdue
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
		"externalTime",
		"internalTime",
		"externalCost",
		"internalCost",
		"neededImmediately",
		"neededTillKM",
		"neededTillDate",
		"repeatTimeSpan",
		"repeatKM",
		"currentKM",
		"game",
		"kmSinceLast",
		"timeSinceLast",
		"visibleRemaining",
		"repeatSpan",
		"spanSinceLast",
		"isNeeded",
		"isOverdue",
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

	varFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance := _FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance(varFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance)

	return err
}

type NullableFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance struct {
	value *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) Get() *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) Set(val *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance(val *FPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) *NullableFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance {
	return &NullableFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsNeededTruckMaintenance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


