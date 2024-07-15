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

// checks if the FPHSpedVAPIServerCommunicationHelperRESTSpedStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIServerCommunicationHelperRESTSpedStats{}

// FPHSpedVAPIServerCommunicationHelperRESTSpedStats struct for FPHSpedVAPIServerCommunicationHelperRESTSpedStats
type FPHSpedVAPIServerCommunicationHelperRESTSpedStats struct {
	DrivenTasks float64 `json:"drivenTasks"`
	DrivenKM float64 `json:"drivenKM"`
	DrivenMI float64 `json:"drivenMI"`
	DrivenWeightKG float64 `json:"drivenWeightKG"`
	DrivenWeightLBS float64 `json:"drivenWeightLBS"`
	SumBoni float64 `json:"sumBoni"`
	SumDamage float64 `json:"sumDamage"`
	SumFerry float64 `json:"sumFerry"`
	UsedFuelLi float64 `json:"usedFuelLi"`
	UsedFuelGal float64 `json:"usedFuelGal"`
	SumIncome float64 `json:"sumIncome"`
	SumMaintenance float64 `json:"sumMaintenance"`
	SumRefuel float64 `json:"sumRefuel"`
	SumTaxes float64 `json:"sumTaxes"`
	SumToll float64 `json:"sumToll"`
	SumValue float64 `json:"sumValue"`
}

type _FPHSpedVAPIServerCommunicationHelperRESTSpedStats FPHSpedVAPIServerCommunicationHelperRESTSpedStats

// NewFPHSpedVAPIServerCommunicationHelperRESTSpedStats instantiates a new FPHSpedVAPIServerCommunicationHelperRESTSpedStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIServerCommunicationHelperRESTSpedStats(drivenTasks float64, drivenKM float64, drivenMI float64, drivenWeightKG float64, drivenWeightLBS float64, sumBoni float64, sumDamage float64, sumFerry float64, usedFuelLi float64, usedFuelGal float64, sumIncome float64, sumMaintenance float64, sumRefuel float64, sumTaxes float64, sumToll float64, sumValue float64) *FPHSpedVAPIServerCommunicationHelperRESTSpedStats {
	this := FPHSpedVAPIServerCommunicationHelperRESTSpedStats{}
	this.DrivenTasks = drivenTasks
	this.DrivenKM = drivenKM
	this.DrivenMI = drivenMI
	this.DrivenWeightKG = drivenWeightKG
	this.DrivenWeightLBS = drivenWeightLBS
	this.SumBoni = sumBoni
	this.SumDamage = sumDamage
	this.SumFerry = sumFerry
	this.UsedFuelLi = usedFuelLi
	this.UsedFuelGal = usedFuelGal
	this.SumIncome = sumIncome
	this.SumMaintenance = sumMaintenance
	this.SumRefuel = sumRefuel
	this.SumTaxes = sumTaxes
	this.SumToll = sumToll
	this.SumValue = sumValue
	return &this
}

// NewFPHSpedVAPIServerCommunicationHelperRESTSpedStatsWithDefaults instantiates a new FPHSpedVAPIServerCommunicationHelperRESTSpedStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIServerCommunicationHelperRESTSpedStatsWithDefaults() *FPHSpedVAPIServerCommunicationHelperRESTSpedStats {
	this := FPHSpedVAPIServerCommunicationHelperRESTSpedStats{}
	return &this
}

// GetDrivenTasks returns the DrivenTasks field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetDrivenTasks() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DrivenTasks
}

// GetDrivenTasksOk returns a tuple with the DrivenTasks field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetDrivenTasksOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrivenTasks, true
}

// SetDrivenTasks sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetDrivenTasks(v float64) {
	o.DrivenTasks = v
}

// GetDrivenKM returns the DrivenKM field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetDrivenKM() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DrivenKM
}

// GetDrivenKMOk returns a tuple with the DrivenKM field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetDrivenKMOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrivenKM, true
}

// SetDrivenKM sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetDrivenKM(v float64) {
	o.DrivenKM = v
}

// GetDrivenMI returns the DrivenMI field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetDrivenMI() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DrivenMI
}

// GetDrivenMIOk returns a tuple with the DrivenMI field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetDrivenMIOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrivenMI, true
}

// SetDrivenMI sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetDrivenMI(v float64) {
	o.DrivenMI = v
}

// GetDrivenWeightKG returns the DrivenWeightKG field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetDrivenWeightKG() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DrivenWeightKG
}

// GetDrivenWeightKGOk returns a tuple with the DrivenWeightKG field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetDrivenWeightKGOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrivenWeightKG, true
}

// SetDrivenWeightKG sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetDrivenWeightKG(v float64) {
	o.DrivenWeightKG = v
}

// GetDrivenWeightLBS returns the DrivenWeightLBS field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetDrivenWeightLBS() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DrivenWeightLBS
}

// GetDrivenWeightLBSOk returns a tuple with the DrivenWeightLBS field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetDrivenWeightLBSOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrivenWeightLBS, true
}

// SetDrivenWeightLBS sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetDrivenWeightLBS(v float64) {
	o.DrivenWeightLBS = v
}

// GetSumBoni returns the SumBoni field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumBoni() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumBoni
}

// GetSumBoniOk returns a tuple with the SumBoni field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumBoniOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumBoni, true
}

// SetSumBoni sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetSumBoni(v float64) {
	o.SumBoni = v
}

// GetSumDamage returns the SumDamage field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumDamage() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumDamage
}

// GetSumDamageOk returns a tuple with the SumDamage field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumDamageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumDamage, true
}

// SetSumDamage sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetSumDamage(v float64) {
	o.SumDamage = v
}

// GetSumFerry returns the SumFerry field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumFerry() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumFerry
}

// GetSumFerryOk returns a tuple with the SumFerry field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumFerryOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumFerry, true
}

// SetSumFerry sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetSumFerry(v float64) {
	o.SumFerry = v
}

// GetUsedFuelLi returns the UsedFuelLi field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetUsedFuelLi() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.UsedFuelLi
}

// GetUsedFuelLiOk returns a tuple with the UsedFuelLi field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetUsedFuelLiOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedFuelLi, true
}

// SetUsedFuelLi sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetUsedFuelLi(v float64) {
	o.UsedFuelLi = v
}

// GetUsedFuelGal returns the UsedFuelGal field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetUsedFuelGal() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.UsedFuelGal
}

// GetUsedFuelGalOk returns a tuple with the UsedFuelGal field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetUsedFuelGalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedFuelGal, true
}

// SetUsedFuelGal sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetUsedFuelGal(v float64) {
	o.UsedFuelGal = v
}

// GetSumIncome returns the SumIncome field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumIncome() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumIncome
}

// GetSumIncomeOk returns a tuple with the SumIncome field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumIncomeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumIncome, true
}

// SetSumIncome sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetSumIncome(v float64) {
	o.SumIncome = v
}

// GetSumMaintenance returns the SumMaintenance field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumMaintenance() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumMaintenance
}

// GetSumMaintenanceOk returns a tuple with the SumMaintenance field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumMaintenanceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumMaintenance, true
}

// SetSumMaintenance sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetSumMaintenance(v float64) {
	o.SumMaintenance = v
}

// GetSumRefuel returns the SumRefuel field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumRefuel() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumRefuel
}

// GetSumRefuelOk returns a tuple with the SumRefuel field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumRefuelOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumRefuel, true
}

// SetSumRefuel sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetSumRefuel(v float64) {
	o.SumRefuel = v
}

// GetSumTaxes returns the SumTaxes field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumTaxes() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumTaxes
}

// GetSumTaxesOk returns a tuple with the SumTaxes field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumTaxesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumTaxes, true
}

// SetSumTaxes sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetSumTaxes(v float64) {
	o.SumTaxes = v
}

// GetSumToll returns the SumToll field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumToll() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumToll
}

// GetSumTollOk returns a tuple with the SumToll field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumTollOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumToll, true
}

// SetSumToll sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetSumToll(v float64) {
	o.SumToll = v
}

// GetSumValue returns the SumValue field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumValue
}

// GetSumValueOk returns a tuple with the SumValue field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) GetSumValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumValue, true
}

// SetSumValue sets field value
func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) SetSumValue(v float64) {
	o.SumValue = v
}

func (o FPHSpedVAPIServerCommunicationHelperRESTSpedStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIServerCommunicationHelperRESTSpedStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["drivenTasks"] = o.DrivenTasks
	toSerialize["drivenKM"] = o.DrivenKM
	toSerialize["drivenMI"] = o.DrivenMI
	toSerialize["drivenWeightKG"] = o.DrivenWeightKG
	toSerialize["drivenWeightLBS"] = o.DrivenWeightLBS
	toSerialize["sumBoni"] = o.SumBoni
	toSerialize["sumDamage"] = o.SumDamage
	toSerialize["sumFerry"] = o.SumFerry
	toSerialize["usedFuelLi"] = o.UsedFuelLi
	toSerialize["usedFuelGal"] = o.UsedFuelGal
	toSerialize["sumIncome"] = o.SumIncome
	toSerialize["sumMaintenance"] = o.SumMaintenance
	toSerialize["sumRefuel"] = o.SumRefuel
	toSerialize["sumTaxes"] = o.SumTaxes
	toSerialize["sumToll"] = o.SumToll
	toSerialize["sumValue"] = o.SumValue
	return toSerialize, nil
}

func (o *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"drivenTasks",
		"drivenKM",
		"drivenMI",
		"drivenWeightKG",
		"drivenWeightLBS",
		"sumBoni",
		"sumDamage",
		"sumFerry",
		"usedFuelLi",
		"usedFuelGal",
		"sumIncome",
		"sumMaintenance",
		"sumRefuel",
		"sumTaxes",
		"sumToll",
		"sumValue",
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

	varFPHSpedVAPIServerCommunicationHelperRESTSpedStats := _FPHSpedVAPIServerCommunicationHelperRESTSpedStats{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIServerCommunicationHelperRESTSpedStats)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIServerCommunicationHelperRESTSpedStats(varFPHSpedVAPIServerCommunicationHelperRESTSpedStats)

	return err
}

type NullableFPHSpedVAPIServerCommunicationHelperRESTSpedStats struct {
	value *FPHSpedVAPIServerCommunicationHelperRESTSpedStats
	isSet bool
}

func (v NullableFPHSpedVAPIServerCommunicationHelperRESTSpedStats) Get() *FPHSpedVAPIServerCommunicationHelperRESTSpedStats {
	return v.value
}

func (v *NullableFPHSpedVAPIServerCommunicationHelperRESTSpedStats) Set(val *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIServerCommunicationHelperRESTSpedStats) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIServerCommunicationHelperRESTSpedStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIServerCommunicationHelperRESTSpedStats(val *FPHSpedVAPIServerCommunicationHelperRESTSpedStats) *NullableFPHSpedVAPIServerCommunicationHelperRESTSpedStats {
	return &NullableFPHSpedVAPIServerCommunicationHelperRESTSpedStats{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIServerCommunicationHelperRESTSpedStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIServerCommunicationHelperRESTSpedStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

