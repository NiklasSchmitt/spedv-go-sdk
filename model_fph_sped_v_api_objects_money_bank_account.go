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

// checks if the FPHSpedVAPIObjectsMoneyBankAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPHSpedVAPIObjectsMoneyBankAccount{}

// FPHSpedVAPIObjectsMoneyBankAccount struct for FPHSpedVAPIObjectsMoneyBankAccount
type FPHSpedVAPIObjectsMoneyBankAccount struct {
	Id int32 `json:"id"`
	Iban NullableString `json:"iban"`
	IbanFormated NullableString `json:"ibanFormated"`
	Currency NullableString `json:"currency"`
	Description NullableString `json:"description"`
	//   0 = System  1 = Spedition  2 = User  -1 = NotSet
	Type FPHSpedVAPIEnumsBankAccountType `json:"type"`
	OwnerSpedition NullableFPHSpedVAPIObjectsSpeditionsSpedition `json:"ownerSpedition"`
	OwnerUser NullableFPHSpedVAPIObjectsUsersUser `json:"ownerUser"`
	Money float64 `json:"money"`
}

type _FPHSpedVAPIObjectsMoneyBankAccount FPHSpedVAPIObjectsMoneyBankAccount

// NewFPHSpedVAPIObjectsMoneyBankAccount instantiates a new FPHSpedVAPIObjectsMoneyBankAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPHSpedVAPIObjectsMoneyBankAccount(id int32, iban NullableString, ibanFormated NullableString, currency NullableString, description NullableString, type_ FPHSpedVAPIEnumsBankAccountType, ownerSpedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, ownerUser NullableFPHSpedVAPIObjectsUsersUser, money float64) *FPHSpedVAPIObjectsMoneyBankAccount {
	this := FPHSpedVAPIObjectsMoneyBankAccount{}
	this.Id = id
	this.Iban = iban
	this.IbanFormated = ibanFormated
	this.Currency = currency
	this.Description = description
	this.Type = type_
	this.OwnerSpedition = ownerSpedition
	this.OwnerUser = ownerUser
	this.Money = money
	return &this
}

// NewFPHSpedVAPIObjectsMoneyBankAccountWithDefaults instantiates a new FPHSpedVAPIObjectsMoneyBankAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPHSpedVAPIObjectsMoneyBankAccountWithDefaults() *FPHSpedVAPIObjectsMoneyBankAccount {
	this := FPHSpedVAPIObjectsMoneyBankAccount{}
	return &this
}

// GetId returns the Id field value
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetId(v int32) {
	o.Id = v
}

// GetIban returns the Iban field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetIban() string {
	if o == nil || o.Iban.Get() == nil {
		var ret string
		return ret
	}

	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetIbanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// SetIban sets field value
func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetIban(v string) {
	o.Iban.Set(&v)
}

// GetIbanFormated returns the IbanFormated field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetIbanFormated() string {
	if o == nil || o.IbanFormated.Get() == nil {
		var ret string
		return ret
	}

	return *o.IbanFormated.Get()
}

// GetIbanFormatedOk returns a tuple with the IbanFormated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetIbanFormatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IbanFormated.Get(), o.IbanFormated.IsSet()
}

// SetIbanFormated sets field value
func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetIbanFormated(v string) {
	o.IbanFormated.Set(&v)
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetCurrency(v string) {
	o.Currency.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetType returns the Type field value
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetType() FPHSpedVAPIEnumsBankAccountType {
	if o == nil {
		var ret FPHSpedVAPIEnumsBankAccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetTypeOk() (*FPHSpedVAPIEnumsBankAccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetType(v FPHSpedVAPIEnumsBankAccountType) {
	o.Type = v
}

// GetOwnerSpedition returns the OwnerSpedition field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsSpeditionsSpedition will be returned
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetOwnerSpedition() FPHSpedVAPIObjectsSpeditionsSpedition {
	if o == nil || o.OwnerSpedition.Get() == nil {
		var ret FPHSpedVAPIObjectsSpeditionsSpedition
		return ret
	}

	return *o.OwnerSpedition.Get()
}

// GetOwnerSpeditionOk returns a tuple with the OwnerSpedition field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetOwnerSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerSpedition.Get(), o.OwnerSpedition.IsSet()
}

// SetOwnerSpedition sets field value
func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetOwnerSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition) {
	o.OwnerSpedition.Set(&v)
}

// GetOwnerUser returns the OwnerUser field value
// If the value is explicit nil, the zero value for FPHSpedVAPIObjectsUsersUser will be returned
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetOwnerUser() FPHSpedVAPIObjectsUsersUser {
	if o == nil || o.OwnerUser.Get() == nil {
		var ret FPHSpedVAPIObjectsUsersUser
		return ret
	}

	return *o.OwnerUser.Get()
}

// GetOwnerUserOk returns a tuple with the OwnerUser field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetOwnerUserOk() (*FPHSpedVAPIObjectsUsersUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerUser.Get(), o.OwnerUser.IsSet()
}

// SetOwnerUser sets field value
func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetOwnerUser(v FPHSpedVAPIObjectsUsersUser) {
	o.OwnerUser.Set(&v)
}

// GetMoney returns the Money field value
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetMoney() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Money
}

// GetMoneyOk returns a tuple with the Money field value
// and a boolean to check if the value has been set.
func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetMoneyOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Money, true
}

// SetMoney sets field value
func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetMoney(v float64) {
	o.Money = v
}

func (o FPHSpedVAPIObjectsMoneyBankAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPHSpedVAPIObjectsMoneyBankAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["iban"] = o.Iban.Get()
	toSerialize["ibanFormated"] = o.IbanFormated.Get()
	toSerialize["currency"] = o.Currency.Get()
	toSerialize["description"] = o.Description.Get()
	toSerialize["type"] = o.Type
	toSerialize["ownerSpedition"] = o.OwnerSpedition.Get()
	toSerialize["ownerUser"] = o.OwnerUser.Get()
	toSerialize["money"] = o.Money
	return toSerialize, nil
}

func (o *FPHSpedVAPIObjectsMoneyBankAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"iban",
		"ibanFormated",
		"currency",
		"description",
		"type",
		"ownerSpedition",
		"ownerUser",
		"money",
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

	varFPHSpedVAPIObjectsMoneyBankAccount := _FPHSpedVAPIObjectsMoneyBankAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPHSpedVAPIObjectsMoneyBankAccount)

	if err != nil {
		return err
	}

	*o = FPHSpedVAPIObjectsMoneyBankAccount(varFPHSpedVAPIObjectsMoneyBankAccount)

	return err
}

type NullableFPHSpedVAPIObjectsMoneyBankAccount struct {
	value *FPHSpedVAPIObjectsMoneyBankAccount
	isSet bool
}

func (v NullableFPHSpedVAPIObjectsMoneyBankAccount) Get() *FPHSpedVAPIObjectsMoneyBankAccount {
	return v.value
}

func (v *NullableFPHSpedVAPIObjectsMoneyBankAccount) Set(val *FPHSpedVAPIObjectsMoneyBankAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableFPHSpedVAPIObjectsMoneyBankAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableFPHSpedVAPIObjectsMoneyBankAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPHSpedVAPIObjectsMoneyBankAccount(val *FPHSpedVAPIObjectsMoneyBankAccount) *NullableFPHSpedVAPIObjectsMoneyBankAccount {
	return &NullableFPHSpedVAPIObjectsMoneyBankAccount{value: val, isSet: true}
}

func (v NullableFPHSpedVAPIObjectsMoneyBankAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPHSpedVAPIObjectsMoneyBankAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


