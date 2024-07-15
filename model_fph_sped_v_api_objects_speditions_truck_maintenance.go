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

// checks if the FPHSpedVAPIObjectsSpeditionsTruckMaintenance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsSpeditionsTruckMaintenance{}

// FPHSpedVAPIObjectsSpeditionsTruckMaintenance struct for FPHSpedVAPIObjectsSpeditionsTruckMaintenance
type FPHSpedVAPIObjectsSpeditionsTruckMaintenance struct {
	Id int32 `json:"id"`
	Truck NullableFPHSpedVAPIObjectsSpeditionsTruck `json:"truck"`
	OwnedBranch NullableFPHSpedVAPIObjectsSpeditionsOwnedBranch `json:"ownedBranch"`
	InitiatedByUser NullableFPHSpedVAPIObjectsUsersUser `json:"initiatedByUser"`
	//   0 = Engine  1 = OszilationDamper  2 = Stabilizer  3 = StoneChip  4 = Transmission  5 = Wishbone  6 = BrakePads  7 = BrakeDiscs  8 = EngineMaintenance  9 = TireChange  10 = MainCheck  11 = SafetyCheck  12 = SaddlePlate  13 = AirPressureUnit  14 = Alternator  15 = BrakeVentil  -1 = NotSet
	MaintenanceKind FPHSpedVAPIEnumsMaintenanceKind `json:"maintenanceKind"`
	SparePartWarningSent bool `json:"sparePartWarningSent"`
	QueueDate time.Time `json:"queueDate"`
	StartDate NullableTime `json:"startDate"`
	EndDate NullableTime `json:"endDate"`
	//   0 = Enqueued  11 = InDrive  12 = NotInOwnedBranch  13 = OtherMaintenanceActive  21 = MissingSpareParts  22 = NotEnoughMaintenancePlaces  91 = Processing  92 = ProcessingButDelayed  93 = Finished  -1 = NotSet
	State FPHSpedVAPIEnumsMaintenanceState `json:"state"`
}

type _FPHSpedVAPIObjectsSpeditionsTruckMaintenance FPHSpedVAPIObjectsSpeditionsTruckMaintenance

// NewFPHSpedVAPIObjectsSpeditionsTruckMaintenance instantiates a new FPHSpedVAPIObjectsSpeditionsTruckMaintenance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsSpeditionsTruckMaintenance(id int32, truck NullableFPHSpedVAPIObjectsSpeditionsTruck, ownedBranch NullableFPHSpedVAPIObjectsSpeditionsOwnedBranch, initiatedByUser NullableFPHSpedVAPIObjectsUsersUser, maintenanceKind FPHSpedVAPIEnumsMaintenanceKind, sparePartWarningSent bool, queueDate time.Time, startDate NullableTime, endDate NullableTime, state FPHSpedVAPIEnumsMaintenanceState) *FPHSpedVAPIObjectsSpeditionsTruckMaintenance {
	this := FPHSpedVAPIObjectsSpeditionsTruckMaintenance{}
	this.Id = id
	this.Truck = truck
	this.OwnedBranch = ownedBranch
	this.InitiatedByUser = initiatedByUser
	this.MaintenanceKind = maintenanceKind
	this.SparePartWarningSent = sparePartWarningSent
	this.QueueDate = queueDate
	this.StartDate = startDate
	this.EndDate = endDate
	this.State = state
	return &this
}

// NewFPHSpedVAPIObjectsSpeditionsTruckMaintenanceWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsTruckMaintenance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsSpeditionsTruckMaintenanceWithDefaults() *FPHSpedVAPIObjectsSpeditionsTruckMaintenance {
	this := FPHSpedVAPIObjectsSpeditionsTruckMaintenance{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) SetId(v int32) {
	o.Id = v
}

// GetTruck returns the Truck field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsSpeditionsTruck will be returned
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetTruck() FPHSpedVAPIObjectsSpeditionsTruck {
	if o == nil || o.Truck.Get() == nil {
		var ret FPHSpedVAPIObjectsSpeditionsTruck
		return ret
	}

	return *o.Truck.Get()
}

// GetTruckOk returns a tuple with the Truck field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetTruckOk() (*FPHSpedVAPIObjectsSpeditionsTruck, bool) {
	if o == nil {
		return nil, false
	}
	return o.Truck.Get(), o.Truck.IsSet()
}

// SetTruck sets field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) SetTruck(v FPHSpedVAPIObjectsSpeditionsTruck) {
	o.Truck.Set(&v)
}

// GetOwnedBranch returns the OwnedBranch field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsSpeditionsOwnedBranch will be returned
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetOwnedBranch() FPHSpedVAPIObjectsSpeditionsOwnedBranch {
	if o == nil || o.OwnedBranch.Get() == nil {
		var ret FPHSpedVAPIObjectsSpeditionsOwnedBranch
		return ret
	}

	return *o.OwnedBranch.Get()
}

// GetOwnedBranchOk returns a tuple with the OwnedBranch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetOwnedBranchOk() (*FPHSpedVAPIObjectsSpeditionsOwnedBranch, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnedBranch.Get(), o.OwnedBranch.IsSet()
}

// SetOwnedBranch sets field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) SetOwnedBranch(v FPHSpedVAPIObjectsSpeditionsOwnedBranch) {
	o.OwnedBranch.Set(&v)
}

// GetInitiatedByUser returns the InitiatedByUser field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsUsersUser will be returned
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetInitiatedByUser() FPHSpedVAPIObjectsUsersUser {
	if o == nil || o.InitiatedByUser.Get() == nil {
		var ret FPHSpedVAPIObjectsUsersUser
		return ret
	}

	return *o.InitiatedByUser.Get()
}

// GetInitiatedByUserOk returns a tuple with the InitiatedByUser field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetInitiatedByUserOk() (*FPHSpedVAPIObjectsUsersUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitiatedByUser.Get(), o.InitiatedByUser.IsSet()
}

// SetInitiatedByUser sets field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) SetInitiatedByUser(v FPHSpedVAPIObjectsUsersUser) {
	o.InitiatedByUser.Set(&v)
}

// GetMaintenanceKind returns the MaintenanceKind field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetMaintenanceKind() FPHSpedVAPIEnumsMaintenanceKind {
	if o == nil {
		var ret FPHSpedVAPIEnumsMaintenanceKind
		return ret
	}

	return o.MaintenanceKind
}

// GetMaintenanceKindOk returns a tuple with the MaintenanceKind field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetMaintenanceKindOk() (*FPHSpedVAPIEnumsMaintenanceKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaintenanceKind, true
}

// SetMaintenanceKind sets field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) SetMaintenanceKind(v FPHSpedVAPIEnumsMaintenanceKind) {
	o.MaintenanceKind = v
}

// GetSparePartWarningSent returns the SparePartWarningSent field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetSparePartWarningSent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SparePartWarningSent
}

// GetSparePartWarningSentOk returns a tuple with the SparePartWarningSent field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetSparePartWarningSentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SparePartWarningSent, true
}

// SetSparePartWarningSent sets field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) SetSparePartWarningSent(v bool) {
	o.SparePartWarningSent = v
}

// GetQueueDate returns the QueueDate field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetQueueDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.QueueDate
}

// GetQueueDateOk returns a tuple with the QueueDate field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetQueueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueueDate, true
}

// SetQueueDate sets field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) SetQueueDate(v time.Time) {
	o.QueueDate = v
}

// GetStartDate returns the StartDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetStartDate() time.Time {
	if o == nil || o.StartDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// SetStartDate sets field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}

// GetEndDate returns the EndDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetEndDate() time.Time {
	if o == nil || o.EndDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// SetEndDate sets field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}

// GetState returns the State field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetState() FPHSpedVAPIEnumsMaintenanceState {
	if o == nil {
		var ret FPHSpedVAPIEnumsMaintenanceState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) GetStateOk() (*FPHSpedVAPIEnumsMaintenanceState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) SetState(v FPHSpedVAPIEnumsMaintenanceState) {
	o.State = v
}

func (o FPHSpedVAPIObjectsSpeditionsTruckMaintenance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsSpeditionsTruckMaintenance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["truck"] = o.Truck.Get()
	toSerialize["ownedBranch"] = o.OwnedBranch.Get()
	toSerialize["initiatedByUser"] = o.InitiatedByUser.Get()
	toSerialize["maintenanceKind"] = o.MaintenanceKind
	toSerialize["sparePartWarningSent"] = o.SparePartWarningSent
	toSerialize["queueDate"] = o.QueueDate
	toSerialize["startDate"] = o.StartDate.Get()
	toSerialize["endDate"] = o.EndDate.Get()
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"truck",
		"ownedBranch",
		"initiatedByUser",
		"maintenanceKind",
		"sparePartWarningSent",
		"queueDate",
		"startDate",
		"endDate",
		"state",
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

	varFPHSpedVAPIObjectsSpeditionsTruckMaintenance := _FPHSpedVAPIObjectsSpeditionsTruckMaintenance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsSpeditionsTruckMaintenance)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsSpeditionsTruckMaintenance(varFPHSpedVAPIObjectsSpeditionsTruckMaintenance)

	return err
}

type NullableFPHSpedVAPIObjectsSpeditionsTruckMaintenance struct {
	value *FPHSpedVAPIObjectsSpeditionsTruckMaintenance
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsSpeditionsTruckMaintenance) Get() *FPHSpedVAPIObjectsSpeditionsTruckMaintenance {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsTruckMaintenance) Set(val *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsSpeditionsTruckMaintenance) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsTruckMaintenance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsSpeditionsTruckMaintenance(val *FPHSpedVAPIObjectsSpeditionsTruckMaintenance) *NullableFPHSpedVAPIObjectsSpeditionsTruckMaintenance {
	return &NullableFPHSpedVAPIObjectsSpeditionsTruckMaintenance{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsSpeditionsTruckMaintenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsSpeditionsTruckMaintenance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


