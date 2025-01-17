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

// checks if the FPHSpedVAPIObjectsLiveMyPlaceStructure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsLiveMyPlaceStructure{}

// FPHSpedVAPIObjectsLiveMyPlaceStructure struct for FPHSpedVAPIObjectsLiveMyPlaceStructure
type FPHSpedVAPIObjectsLiveMyPlaceStructure struct {
	SpedEURPlace int32 `json:"spedEURPlace"`
	SpedEURDiffNext float64 `json:"spedEURDiffNext"`
	SpedEURDiffLast float64 `json:"spedEURDiffLast"`
	SpedKMPlace int32 `json:"spedKMPlace"`
	SpedKMDiffNext float64 `json:"spedKMDiffNext"`
	SpedKMDiffLast float64 `json:"spedKMDiffLast"`
	SpedWeeklyEURPlace int32 `json:"spedWeeklyEURPlace"`
	SpedWeeklyEURDiffNext float64 `json:"spedWeeklyEURDiffNext"`
	SpedWeeklyEURDiffLast float64 `json:"spedWeeklyEURDiffLast"`
	SpedWeeklyKontorEURPlace int32 `json:"spedWeeklyKontorEURPlace"`
	SpedWeeklyKontorEURDiffNext float64 `json:"spedWeeklyKontorEURDiffNext"`
	SpedWeeklyKontorEURDiffLast float64 `json:"spedWeeklyKontorEURDiffLast"`
	SpedWeeklyExternalEURPlace int32 `json:"spedWeeklyExternalEURPlace"`
	SpedWeeklyExternalEURDiffNext float64 `json:"spedWeeklyExternalEURDiffNext"`
	SpedWeeklyExternalEURDiffLast float64 `json:"spedWeeklyExternalEURDiffLast"`
	SpedWeeklyKMPlace int32 `json:"spedWeeklyKMPlace"`
	SpedWeeklyKMDiffNext float64 `json:"spedWeeklyKMDiffNext"`
	SpedWeeklyKMDiffLast float64 `json:"spedWeeklyKMDiffLast"`
	SpedExternalEURPlace int32 `json:"spedExternalEURPlace"`
	SpedExternalEURDiffNext float64 `json:"spedExternalEURDiffNext"`
	SpedExternalEURDiffLast float64 `json:"spedExternalEURDiffLast"`
	UserEURPlace int32 `json:"userEURPlace"`
	UserEURDiffNext float64 `json:"userEURDiffNext"`
	UserEURDiffLast float64 `json:"userEURDiffLast"`
	UserKMPlace int32 `json:"userKMPlace"`
	UserKMDiffNext float64 `json:"userKMDiffNext"`
	UserKMDiffLast float64 `json:"userKMDiffLast"`
}

type _FPHSpedVAPIObjectsLiveMyPlaceStructure FPHSpedVAPIObjectsLiveMyPlaceStructure

// NewFPHSpedVAPIObjectsLiveMyPlaceStructure instantiates a new FPHSpedVAPIObjectsLiveMyPlaceStructure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsLiveMyPlaceStructure(spedEURPlace int32, spedEURDiffNext float64, spedEURDiffLast float64, spedKMPlace int32, spedKMDiffNext float64, spedKMDiffLast float64, spedWeeklyEURPlace int32, spedWeeklyEURDiffNext float64, spedWeeklyEURDiffLast float64, spedWeeklyKontorEURPlace int32, spedWeeklyKontorEURDiffNext float64, spedWeeklyKontorEURDiffLast float64, spedWeeklyExternalEURPlace int32, spedWeeklyExternalEURDiffNext float64, spedWeeklyExternalEURDiffLast float64, spedWeeklyKMPlace int32, spedWeeklyKMDiffNext float64, spedWeeklyKMDiffLast float64, spedExternalEURPlace int32, spedExternalEURDiffNext float64, spedExternalEURDiffLast float64, userEURPlace int32, userEURDiffNext float64, userEURDiffLast float64, userKMPlace int32, userKMDiffNext float64, userKMDiffLast float64) *FPHSpedVAPIObjectsLiveMyPlaceStructure {
	this := FPHSpedVAPIObjectsLiveMyPlaceStructure{}
	this.SpedEURPlace = spedEURPlace
	this.SpedEURDiffNext = spedEURDiffNext
	this.SpedEURDiffLast = spedEURDiffLast
	this.SpedKMPlace = spedKMPlace
	this.SpedKMDiffNext = spedKMDiffNext
	this.SpedKMDiffLast = spedKMDiffLast
	this.SpedWeeklyEURPlace = spedWeeklyEURPlace
	this.SpedWeeklyEURDiffNext = spedWeeklyEURDiffNext
	this.SpedWeeklyEURDiffLast = spedWeeklyEURDiffLast
	this.SpedWeeklyKontorEURPlace = spedWeeklyKontorEURPlace
	this.SpedWeeklyKontorEURDiffNext = spedWeeklyKontorEURDiffNext
	this.SpedWeeklyKontorEURDiffLast = spedWeeklyKontorEURDiffLast
	this.SpedWeeklyExternalEURPlace = spedWeeklyExternalEURPlace
	this.SpedWeeklyExternalEURDiffNext = spedWeeklyExternalEURDiffNext
	this.SpedWeeklyExternalEURDiffLast = spedWeeklyExternalEURDiffLast
	this.SpedWeeklyKMPlace = spedWeeklyKMPlace
	this.SpedWeeklyKMDiffNext = spedWeeklyKMDiffNext
	this.SpedWeeklyKMDiffLast = spedWeeklyKMDiffLast
	this.SpedExternalEURPlace = spedExternalEURPlace
	this.SpedExternalEURDiffNext = spedExternalEURDiffNext
	this.SpedExternalEURDiffLast = spedExternalEURDiffLast
	this.UserEURPlace = userEURPlace
	this.UserEURDiffNext = userEURDiffNext
	this.UserEURDiffLast = userEURDiffLast
	this.UserKMPlace = userKMPlace
	this.UserKMDiffNext = userKMDiffNext
	this.UserKMDiffLast = userKMDiffLast
	return &this
}

// NewFPHSpedVAPIObjectsLiveMyPlaceStructureWithDefaults instantiates a new FPHSpedVAPIObjectsLiveMyPlaceStructure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsLiveMyPlaceStructureWithDefaults() *FPHSpedVAPIObjectsLiveMyPlaceStructure {
	this := FPHSpedVAPIObjectsLiveMyPlaceStructure{}
	return &this
}

// GetSpedEURPlace returns the SpedEURPlace field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedEURPlace() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SpedEURPlace
}

// GetSpedEURPlaceOk returns a tuple with the SpedEURPlace field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedEURPlaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedEURPlace, true
}

// SetSpedEURPlace sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedEURPlace(v int32) {
	o.SpedEURPlace = v
}

// GetSpedEURDiffNext returns the SpedEURDiffNext field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedEURDiffNext() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedEURDiffNext
}

// GetSpedEURDiffNextOk returns a tuple with the SpedEURDiffNext field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedEURDiffNextOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedEURDiffNext, true
}

// SetSpedEURDiffNext sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedEURDiffNext(v float64) {
	o.SpedEURDiffNext = v
}

// GetSpedEURDiffLast returns the SpedEURDiffLast field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedEURDiffLast() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedEURDiffLast
}

// GetSpedEURDiffLastOk returns a tuple with the SpedEURDiffLast field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedEURDiffLastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedEURDiffLast, true
}

// SetSpedEURDiffLast sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedEURDiffLast(v float64) {
	o.SpedEURDiffLast = v
}

// GetSpedKMPlace returns the SpedKMPlace field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedKMPlace() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SpedKMPlace
}

// GetSpedKMPlaceOk returns a tuple with the SpedKMPlace field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedKMPlaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedKMPlace, true
}

// SetSpedKMPlace sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedKMPlace(v int32) {
	o.SpedKMPlace = v
}

// GetSpedKMDiffNext returns the SpedKMDiffNext field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedKMDiffNext() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedKMDiffNext
}

// GetSpedKMDiffNextOk returns a tuple with the SpedKMDiffNext field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedKMDiffNextOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedKMDiffNext, true
}

// SetSpedKMDiffNext sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedKMDiffNext(v float64) {
	o.SpedKMDiffNext = v
}

// GetSpedKMDiffLast returns the SpedKMDiffLast field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedKMDiffLast() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedKMDiffLast
}

// GetSpedKMDiffLastOk returns a tuple with the SpedKMDiffLast field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedKMDiffLastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedKMDiffLast, true
}

// SetSpedKMDiffLast sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedKMDiffLast(v float64) {
	o.SpedKMDiffLast = v
}

// GetSpedWeeklyEURPlace returns the SpedWeeklyEURPlace field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyEURPlace() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SpedWeeklyEURPlace
}

// GetSpedWeeklyEURPlaceOk returns a tuple with the SpedWeeklyEURPlace field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyEURPlaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedWeeklyEURPlace, true
}

// SetSpedWeeklyEURPlace sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedWeeklyEURPlace(v int32) {
	o.SpedWeeklyEURPlace = v
}

// GetSpedWeeklyEURDiffNext returns the SpedWeeklyEURDiffNext field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyEURDiffNext() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedWeeklyEURDiffNext
}

// GetSpedWeeklyEURDiffNextOk returns a tuple with the SpedWeeklyEURDiffNext field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyEURDiffNextOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedWeeklyEURDiffNext, true
}

// SetSpedWeeklyEURDiffNext sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedWeeklyEURDiffNext(v float64) {
	o.SpedWeeklyEURDiffNext = v
}

// GetSpedWeeklyEURDiffLast returns the SpedWeeklyEURDiffLast field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyEURDiffLast() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedWeeklyEURDiffLast
}

// GetSpedWeeklyEURDiffLastOk returns a tuple with the SpedWeeklyEURDiffLast field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyEURDiffLastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedWeeklyEURDiffLast, true
}

// SetSpedWeeklyEURDiffLast sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedWeeklyEURDiffLast(v float64) {
	o.SpedWeeklyEURDiffLast = v
}

// GetSpedWeeklyKontorEURPlace returns the SpedWeeklyKontorEURPlace field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyKontorEURPlace() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SpedWeeklyKontorEURPlace
}

// GetSpedWeeklyKontorEURPlaceOk returns a tuple with the SpedWeeklyKontorEURPlace field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyKontorEURPlaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedWeeklyKontorEURPlace, true
}

// SetSpedWeeklyKontorEURPlace sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedWeeklyKontorEURPlace(v int32) {
	o.SpedWeeklyKontorEURPlace = v
}

// GetSpedWeeklyKontorEURDiffNext returns the SpedWeeklyKontorEURDiffNext field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyKontorEURDiffNext() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedWeeklyKontorEURDiffNext
}

// GetSpedWeeklyKontorEURDiffNextOk returns a tuple with the SpedWeeklyKontorEURDiffNext field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyKontorEURDiffNextOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedWeeklyKontorEURDiffNext, true
}

// SetSpedWeeklyKontorEURDiffNext sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedWeeklyKontorEURDiffNext(v float64) {
	o.SpedWeeklyKontorEURDiffNext = v
}

// GetSpedWeeklyKontorEURDiffLast returns the SpedWeeklyKontorEURDiffLast field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyKontorEURDiffLast() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedWeeklyKontorEURDiffLast
}

// GetSpedWeeklyKontorEURDiffLastOk returns a tuple with the SpedWeeklyKontorEURDiffLast field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyKontorEURDiffLastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedWeeklyKontorEURDiffLast, true
}

// SetSpedWeeklyKontorEURDiffLast sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedWeeklyKontorEURDiffLast(v float64) {
	o.SpedWeeklyKontorEURDiffLast = v
}

// GetSpedWeeklyExternalEURPlace returns the SpedWeeklyExternalEURPlace field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyExternalEURPlace() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SpedWeeklyExternalEURPlace
}

// GetSpedWeeklyExternalEURPlaceOk returns a tuple with the SpedWeeklyExternalEURPlace field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyExternalEURPlaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedWeeklyExternalEURPlace, true
}

// SetSpedWeeklyExternalEURPlace sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedWeeklyExternalEURPlace(v int32) {
	o.SpedWeeklyExternalEURPlace = v
}

// GetSpedWeeklyExternalEURDiffNext returns the SpedWeeklyExternalEURDiffNext field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyExternalEURDiffNext() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedWeeklyExternalEURDiffNext
}

// GetSpedWeeklyExternalEURDiffNextOk returns a tuple with the SpedWeeklyExternalEURDiffNext field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyExternalEURDiffNextOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedWeeklyExternalEURDiffNext, true
}

// SetSpedWeeklyExternalEURDiffNext sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedWeeklyExternalEURDiffNext(v float64) {
	o.SpedWeeklyExternalEURDiffNext = v
}

// GetSpedWeeklyExternalEURDiffLast returns the SpedWeeklyExternalEURDiffLast field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyExternalEURDiffLast() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedWeeklyExternalEURDiffLast
}

// GetSpedWeeklyExternalEURDiffLastOk returns a tuple with the SpedWeeklyExternalEURDiffLast field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyExternalEURDiffLastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedWeeklyExternalEURDiffLast, true
}

// SetSpedWeeklyExternalEURDiffLast sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedWeeklyExternalEURDiffLast(v float64) {
	o.SpedWeeklyExternalEURDiffLast = v
}

// GetSpedWeeklyKMPlace returns the SpedWeeklyKMPlace field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyKMPlace() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SpedWeeklyKMPlace
}

// GetSpedWeeklyKMPlaceOk returns a tuple with the SpedWeeklyKMPlace field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyKMPlaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedWeeklyKMPlace, true
}

// SetSpedWeeklyKMPlace sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedWeeklyKMPlace(v int32) {
	o.SpedWeeklyKMPlace = v
}

// GetSpedWeeklyKMDiffNext returns the SpedWeeklyKMDiffNext field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyKMDiffNext() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedWeeklyKMDiffNext
}

// GetSpedWeeklyKMDiffNextOk returns a tuple with the SpedWeeklyKMDiffNext field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyKMDiffNextOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedWeeklyKMDiffNext, true
}

// SetSpedWeeklyKMDiffNext sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedWeeklyKMDiffNext(v float64) {
	o.SpedWeeklyKMDiffNext = v
}

// GetSpedWeeklyKMDiffLast returns the SpedWeeklyKMDiffLast field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyKMDiffLast() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedWeeklyKMDiffLast
}

// GetSpedWeeklyKMDiffLastOk returns a tuple with the SpedWeeklyKMDiffLast field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedWeeklyKMDiffLastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedWeeklyKMDiffLast, true
}

// SetSpedWeeklyKMDiffLast sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedWeeklyKMDiffLast(v float64) {
	o.SpedWeeklyKMDiffLast = v
}

// GetSpedExternalEURPlace returns the SpedExternalEURPlace field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedExternalEURPlace() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SpedExternalEURPlace
}

// GetSpedExternalEURPlaceOk returns a tuple with the SpedExternalEURPlace field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedExternalEURPlaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedExternalEURPlace, true
}

// SetSpedExternalEURPlace sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedExternalEURPlace(v int32) {
	o.SpedExternalEURPlace = v
}

// GetSpedExternalEURDiffNext returns the SpedExternalEURDiffNext field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedExternalEURDiffNext() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedExternalEURDiffNext
}

// GetSpedExternalEURDiffNextOk returns a tuple with the SpedExternalEURDiffNext field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedExternalEURDiffNextOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedExternalEURDiffNext, true
}

// SetSpedExternalEURDiffNext sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedExternalEURDiffNext(v float64) {
	o.SpedExternalEURDiffNext = v
}

// GetSpedExternalEURDiffLast returns the SpedExternalEURDiffLast field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedExternalEURDiffLast() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SpedExternalEURDiffLast
}

// GetSpedExternalEURDiffLastOk returns a tuple with the SpedExternalEURDiffLast field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetSpedExternalEURDiffLastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpedExternalEURDiffLast, true
}

// SetSpedExternalEURDiffLast sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetSpedExternalEURDiffLast(v float64) {
	o.SpedExternalEURDiffLast = v
}

// GetUserEURPlace returns the UserEURPlace field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetUserEURPlace() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserEURPlace
}

// GetUserEURPlaceOk returns a tuple with the UserEURPlace field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetUserEURPlaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEURPlace, true
}

// SetUserEURPlace sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetUserEURPlace(v int32) {
	o.UserEURPlace = v
}

// GetUserEURDiffNext returns the UserEURDiffNext field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetUserEURDiffNext() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.UserEURDiffNext
}

// GetUserEURDiffNextOk returns a tuple with the UserEURDiffNext field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetUserEURDiffNextOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEURDiffNext, true
}

// SetUserEURDiffNext sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetUserEURDiffNext(v float64) {
	o.UserEURDiffNext = v
}

// GetUserEURDiffLast returns the UserEURDiffLast field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetUserEURDiffLast() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.UserEURDiffLast
}

// GetUserEURDiffLastOk returns a tuple with the UserEURDiffLast field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetUserEURDiffLastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEURDiffLast, true
}

// SetUserEURDiffLast sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetUserEURDiffLast(v float64) {
	o.UserEURDiffLast = v
}

// GetUserKMPlace returns the UserKMPlace field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetUserKMPlace() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserKMPlace
}

// GetUserKMPlaceOk returns a tuple with the UserKMPlace field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetUserKMPlaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserKMPlace, true
}

// SetUserKMPlace sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetUserKMPlace(v int32) {
	o.UserKMPlace = v
}

// GetUserKMDiffNext returns the UserKMDiffNext field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetUserKMDiffNext() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.UserKMDiffNext
}

// GetUserKMDiffNextOk returns a tuple with the UserKMDiffNext field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetUserKMDiffNextOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserKMDiffNext, true
}

// SetUserKMDiffNext sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetUserKMDiffNext(v float64) {
	o.UserKMDiffNext = v
}

// GetUserKMDiffLast returns the UserKMDiffLast field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetUserKMDiffLast() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.UserKMDiffLast
}

// GetUserKMDiffLastOk returns a tuple with the UserKMDiffLast field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) GetUserKMDiffLastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserKMDiffLast, true
}

// SetUserKMDiffLast sets field value
func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) SetUserKMDiffLast(v float64) {
	o.UserKMDiffLast = v
}

func (o FPHSpedVAPIObjectsLiveMyPlaceStructure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsLiveMyPlaceStructure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["spedEURPlace"] = o.SpedEURPlace
	toSerialize["spedEURDiffNext"] = o.SpedEURDiffNext
	toSerialize["spedEURDiffLast"] = o.SpedEURDiffLast
	toSerialize["spedKMPlace"] = o.SpedKMPlace
	toSerialize["spedKMDiffNext"] = o.SpedKMDiffNext
	toSerialize["spedKMDiffLast"] = o.SpedKMDiffLast
	toSerialize["spedWeeklyEURPlace"] = o.SpedWeeklyEURPlace
	toSerialize["spedWeeklyEURDiffNext"] = o.SpedWeeklyEURDiffNext
	toSerialize["spedWeeklyEURDiffLast"] = o.SpedWeeklyEURDiffLast
	toSerialize["spedWeeklyKontorEURPlace"] = o.SpedWeeklyKontorEURPlace
	toSerialize["spedWeeklyKontorEURDiffNext"] = o.SpedWeeklyKontorEURDiffNext
	toSerialize["spedWeeklyKontorEURDiffLast"] = o.SpedWeeklyKontorEURDiffLast
	toSerialize["spedWeeklyExternalEURPlace"] = o.SpedWeeklyExternalEURPlace
	toSerialize["spedWeeklyExternalEURDiffNext"] = o.SpedWeeklyExternalEURDiffNext
	toSerialize["spedWeeklyExternalEURDiffLast"] = o.SpedWeeklyExternalEURDiffLast
	toSerialize["spedWeeklyKMPlace"] = o.SpedWeeklyKMPlace
	toSerialize["spedWeeklyKMDiffNext"] = o.SpedWeeklyKMDiffNext
	toSerialize["spedWeeklyKMDiffLast"] = o.SpedWeeklyKMDiffLast
	toSerialize["spedExternalEURPlace"] = o.SpedExternalEURPlace
	toSerialize["spedExternalEURDiffNext"] = o.SpedExternalEURDiffNext
	toSerialize["spedExternalEURDiffLast"] = o.SpedExternalEURDiffLast
	toSerialize["userEURPlace"] = o.UserEURPlace
	toSerialize["userEURDiffNext"] = o.UserEURDiffNext
	toSerialize["userEURDiffLast"] = o.UserEURDiffLast
	toSerialize["userKMPlace"] = o.UserKMPlace
	toSerialize["userKMDiffNext"] = o.UserKMDiffNext
	toSerialize["userKMDiffLast"] = o.UserKMDiffLast
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsLiveMyPlaceStructure) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"spedEURPlace",
		"spedEURDiffNext",
		"spedEURDiffLast",
		"spedKMPlace",
		"spedKMDiffNext",
		"spedKMDiffLast",
		"spedWeeklyEURPlace",
		"spedWeeklyEURDiffNext",
		"spedWeeklyEURDiffLast",
		"spedWeeklyKontorEURPlace",
		"spedWeeklyKontorEURDiffNext",
		"spedWeeklyKontorEURDiffLast",
		"spedWeeklyExternalEURPlace",
		"spedWeeklyExternalEURDiffNext",
		"spedWeeklyExternalEURDiffLast",
		"spedWeeklyKMPlace",
		"spedWeeklyKMDiffNext",
		"spedWeeklyKMDiffLast",
		"spedExternalEURPlace",
		"spedExternalEURDiffNext",
		"spedExternalEURDiffLast",
		"userEURPlace",
		"userEURDiffNext",
		"userEURDiffLast",
		"userKMPlace",
		"userKMDiffNext",
		"userKMDiffLast",
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

	varFPHSpedVAPIObjectsLiveMyPlaceStructure := _FPHSpedVAPIObjectsLiveMyPlaceStructure{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsLiveMyPlaceStructure)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsLiveMyPlaceStructure(varFPHSpedVAPIObjectsLiveMyPlaceStructure)

	return err
}

type NullableFPHSpedVAPIObjectsLiveMyPlaceStructure struct {
	value *FPHSpedVAPIObjectsLiveMyPlaceStructure
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsLiveMyPlaceStructure) Get() *FPHSpedVAPIObjectsLiveMyPlaceStructure {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsLiveMyPlaceStructure) Set(val *FPHSpedVAPIObjectsLiveMyPlaceStructure) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsLiveMyPlaceStructure) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsLiveMyPlaceStructure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsLiveMyPlaceStructure(val *FPHSpedVAPIObjectsLiveMyPlaceStructure) *NullableFPHSpedVAPIObjectsLiveMyPlaceStructure {
	return &NullableFPHSpedVAPIObjectsLiveMyPlaceStructure{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsLiveMyPlaceStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsLiveMyPlaceStructure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


