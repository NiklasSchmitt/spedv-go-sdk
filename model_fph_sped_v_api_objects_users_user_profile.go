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

// checks if the FPHSpedVAPIObjectsUsersUserProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsUsersUserProfile{}

// FPHSpedVAPIObjectsUsersUserProfile struct for FPHSpedVAPIObjectsUsersUserProfile
type FPHSpedVAPIObjectsUsersUserProfile struct {
	Id int32 `json:"id"`
	SteamID int64 `json:"steamID"`
	IncomeSum float64 `json:"incomeSum"`
	DamageSum float64 `json:"damageSum"`
	KmSum float64 `json:"kmSum"`
	RealName NullableString `json:"realName"`
	AboutMe NullableString `json:"aboutMe"`
	Birthday time.Time `json:"birthday"`
	SteamImgURL NullableString `json:"steamImgURL"`
	Residence NullableString `json:"residence"`
	UserName NullableString `json:"userName"`
	ChangeLog []FPHSpedVAPIObjectsUsersUserChangeLogEntry `json:"changeLog"`
	Spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition `json:"spedition"`
	InSped bool `json:"inSped"`
	FormatName NullableString `json:"formatName"`
	LastTaskTime time.Time `json:"lastTaskTime"`
	MainBankAccount NullableFPHSpedVAPIObjectsMoneyLiteBankAccount `json:"mainBankAccount"`
}

type _FPHSpedVAPIObjectsUsersUserProfile FPHSpedVAPIObjectsUsersUserProfile

// NewFPHSpedVAPIObjectsUsersUserProfile instantiates a new FPHSpedVAPIObjectsUsersUserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsUsersUserProfile(id int32, steamID int64, incomeSum float64, damageSum float64, kmSum float64, realName NullableString, aboutMe NullableString, birthday time.Time, steamImgURL NullableString, residence NullableString, userName NullableString, changeLog []FPHSpedVAPIObjectsUsersUserChangeLogEntry, spedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, inSped bool, formatName NullableString, lastTaskTime time.Time, mainBankAccount NullableFPHSpedVAPIObjectsMoneyLiteBankAccount) *FPHSpedVAPIObjectsUsersUserProfile {
	this := FPHSpedVAPIObjectsUsersUserProfile{}
	this.Id = id
	this.SteamID = steamID
	this.IncomeSum = incomeSum
	this.DamageSum = damageSum
	this.KmSum = kmSum
	this.RealName = realName
	this.AboutMe = aboutMe
	this.Birthday = birthday
	this.SteamImgURL = steamImgURL
	this.Residence = residence
	this.UserName = userName
	this.ChangeLog = changeLog
	this.Spedition = spedition
	this.InSped = inSped
	this.FormatName = formatName
	this.LastTaskTime = lastTaskTime
	this.MainBankAccount = mainBankAccount
	return &this
}

// NewFPHSpedVAPIObjectsUsersUserProfileWithDefaults instantiates a new FPHSpedVAPIObjectsUsersUserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsUsersUserProfileWithDefaults() *FPHSpedVAPIObjectsUsersUserProfile {
	this := FPHSpedVAPIObjectsUsersUserProfile{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetId(v int32) {
	o.Id = v
}

// GetSteamID returns the SteamID field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetSteamID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SteamID
}

// GetSteamIDOk returns a tuple with the SteamID field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetSteamIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SteamID, true
}

// SetSteamID sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetSteamID(v int64) {
	o.SteamID = v
}

// GetIncomeSum returns the IncomeSum field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetIncomeSum() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.IncomeSum
}

// GetIncomeSumOk returns a tuple with the IncomeSum field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetIncomeSumOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncomeSum, true
}

// SetIncomeSum sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetIncomeSum(v float64) {
	o.IncomeSum = v
}

// GetDamageSum returns the DamageSum field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetDamageSum() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DamageSum
}

// GetDamageSumOk returns a tuple with the DamageSum field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetDamageSumOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DamageSum, true
}

// SetDamageSum sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetDamageSum(v float64) {
	o.DamageSum = v
}

// GetKmSum returns the KmSum field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetKmSum() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.KmSum
}

// GetKmSumOk returns a tuple with the KmSum field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetKmSumOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KmSum, true
}

// SetKmSum sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetKmSum(v float64) {
	o.KmSum = v
}

// GetRealName returns the RealName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetRealName() string {
	if o == nil || o.RealName.Get() == nil {
		var ret string
		return ret
	}

	return *o.RealName.Get()
}

// GetRealNameOk returns a tuple with the RealName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetRealNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RealName.Get(), o.RealName.IsSet()
}

// SetRealName sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetRealName(v string) {
	o.RealName.Set(&v)
}

// GetAboutMe returns the AboutMe field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetAboutMe() string {
	if o == nil || o.AboutMe.Get() == nil {
		var ret string
		return ret
	}

	return *o.AboutMe.Get()
}

// GetAboutMeOk returns a tuple with the AboutMe field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetAboutMeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AboutMe.Get(), o.AboutMe.IsSet()
}

// SetAboutMe sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetAboutMe(v string) {
	o.AboutMe.Set(&v)
}

// GetBirthday returns the Birthday field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetBirthday() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Birthday
}

// GetBirthdayOk returns a tuple with the Birthday field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetBirthdayOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Birthday, true
}

// SetBirthday sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetBirthday(v time.Time) {
	o.Birthday = v
}

// GetSteamImgURL returns the SteamImgURL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetSteamImgURL() string {
	if o == nil || o.SteamImgURL.Get() == nil {
		var ret string
		return ret
	}

	return *o.SteamImgURL.Get()
}

// GetSteamImgURLOk returns a tuple with the SteamImgURL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetSteamImgURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SteamImgURL.Get(), o.SteamImgURL.IsSet()
}

// SetSteamImgURL sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetSteamImgURL(v string) {
	o.SteamImgURL.Set(&v)
}

// GetResidence returns the Residence field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetResidence() string {
	if o == nil || o.Residence.Get() == nil {
		var ret string
		return ret
	}

	return *o.Residence.Get()
}

// GetResidenceOk returns a tuple with the Residence field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetResidenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Residence.Get(), o.Residence.IsSet()
}

// SetResidence sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetResidence(v string) {
	o.Residence.Set(&v)
}

// GetUserName returns the UserName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}

	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// SetUserName sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetUserName(v string) {
	o.UserName.Set(&v)
}

// GetChangeLog returns the ChangeLog field value
// If the value is explicit nil, the zero value for []FPHSpedVAPIObjectsUsersUserChangeLogEntry will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetChangeLog() []FPHSpedVAPIObjectsUsersUserChangeLogEntry {
	if o == nil {
		var ret []FPHSpedVAPIObjectsUsersUserChangeLogEntry
		return ret
	}

	return o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetChangeLogOk() ([]FPHSpedVAPIObjectsUsersUserChangeLogEntry, bool) {
	if o == nil || IsNil(o.ChangeLog) {
		return nil, false
	}
	return o.ChangeLog, true
}

// SetChangeLog sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetChangeLog(v []FPHSpedVAPIObjectsUsersUserChangeLogEntry) {
	o.ChangeLog = v
}

// GetSpedition returns the Spedition field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsSpeditionsSpedition will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetSpedition() FPHSpedVAPIObjectsSpeditionsSpedition {
	if o == nil || o.Spedition.Get() == nil {
		var ret FPHSpedVAPIObjectsSpeditionsSpedition
		return ret
	}

	return *o.Spedition.Get()
}

// GetSpeditionOk returns a tuple with the Spedition field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Spedition.Get(), o.Spedition.IsSet()
}

// SetSpedition sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition) {
	o.Spedition.Set(&v)
}

// GetInSped returns the InSped field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetInSped() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InSped
}

// GetInSpedOk returns a tuple with the InSped field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetInSpedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InSped, true
}

// SetInSped sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetInSped(v bool) {
	o.InSped = v
}

// GetFormatName returns the FormatName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetFormatName() string {
	if o == nil || o.FormatName.Get() == nil {
		var ret string
		return ret
	}

	return *o.FormatName.Get()
}

// GetFormatNameOk returns a tuple with the FormatName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetFormatNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormatName.Get(), o.FormatName.IsSet()
}

// SetFormatName sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetFormatName(v string) {
	o.FormatName.Set(&v)
}

// GetLastTaskTime returns the LastTaskTime field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetLastTaskTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastTaskTime
}

// GetLastTaskTimeOk returns a tuple with the LastTaskTime field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetLastTaskTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastTaskTime, true
}

// SetLastTaskTime sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetLastTaskTime(v time.Time) {
	o.LastTaskTime = v
}

// GetMainBankAccount returns the MainBankAccount field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsMoneyLiteBankAccount will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetMainBankAccount() FPHSpedVAPIObjectsMoneyLiteBankAccount {
	if o == nil || o.MainBankAccount.Get() == nil {
		var ret FPHSpedVAPIObjectsMoneyLiteBankAccount
		return ret
	}

	return *o.MainBankAccount.Get()
}

// GetMainBankAccountOk returns a tuple with the MainBankAccount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsUsersUserProfile) GetMainBankAccountOk() (*FPHSpedVAPIObjectsMoneyLiteBankAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.MainBankAccount.Get(), o.MainBankAccount.IsSet()
}

// SetMainBankAccount sets field value
func (o *FPHSpedVAPIObjectsUsersUserProfile) SetMainBankAccount(v FPHSpedVAPIObjectsMoneyLiteBankAccount) {
	o.MainBankAccount.Set(&v)
}

func (o FPHSpedVAPIObjectsUsersUserProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsUsersUserProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["steamID"] = o.SteamID
	toSerialize["incomeSum"] = o.IncomeSum
	toSerialize["damageSum"] = o.DamageSum
	toSerialize["kmSum"] = o.KmSum
	toSerialize["realName"] = o.RealName.Get()
	toSerialize["aboutMe"] = o.AboutMe.Get()
	toSerialize["birthday"] = o.Birthday
	toSerialize["steamImgURL"] = o.SteamImgURL.Get()
	toSerialize["residence"] = o.Residence.Get()
	toSerialize["userName"] = o.UserName.Get()
	if o.ChangeLog != nil {
		toSerialize["changeLog"] = o.ChangeLog
	}
	toSerialize["spedition"] = o.Spedition.Get()
	toSerialize["inSped"] = o.InSped
	toSerialize["formatName"] = o.FormatName.Get()
	toSerialize["lastTaskTime"] = o.LastTaskTime
	toSerialize["mainBankAccount"] = o.MainBankAccount.Get()
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsUsersUserProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"steamID",
		"incomeSum",
		"damageSum",
		"kmSum",
		"realName",
		"aboutMe",
		"birthday",
		"steamImgURL",
		"residence",
		"userName",
		"changeLog",
		"spedition",
		"inSped",
		"formatName",
		"lastTaskTime",
		"mainBankAccount",
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

	varFPHSpedVAPIObjectsUsersUserProfile := _FPHSpedVAPIObjectsUsersUserProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsUsersUserProfile)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsUsersUserProfile(varFPHSpedVAPIObjectsUsersUserProfile)

	return err
}

type NullableFPHSpedVAPIObjectsUsersUserProfile struct {
	value *FPHSpedVAPIObjectsUsersUserProfile
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsUsersUserProfile) Get() *FPHSpedVAPIObjectsUsersUserProfile {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsUsersUserProfile) Set(val *FPHSpedVAPIObjectsUsersUserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsUsersUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsUsersUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsUsersUserProfile(val *FPHSpedVAPIObjectsUsersUserProfile) *NullableFPHSpedVAPIObjectsUsersUserProfile {
	return &NullableFPHSpedVAPIObjectsUsersUserProfile{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsUsersUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsUsersUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

