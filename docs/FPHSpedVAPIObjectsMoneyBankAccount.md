# FPHSpedVAPIObjectsMoneyBankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Iban** | **NullableString** |  | [readonly] 
**IbanFormated** | **NullableString** |  | [readonly] 
**Currency** | **NullableString** |  | [readonly] 
**Description** | **NullableString** |  | [readonly] 
**Type** | [**FPHSpedVAPIEnumsBankAccountType**](FPHSpedVAPIEnumsBankAccountType.md) |   0 &#x3D; System  1 &#x3D; Spedition  2 &#x3D; User  -1 &#x3D; NotSet | [readonly] 
**OwnerSpedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**OwnerUser** | [**NullableFPHSpedVAPIObjectsUsersUser**](FPHSpedVAPIObjectsUsersUser.md) |  | [readonly] 
**Money** | **float64** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsMoneyBankAccount

`func NewFPHSpedVAPIObjectsMoneyBankAccount(id int32, iban NullableString, ibanFormated NullableString, currency NullableString, description NullableString, type_ FPHSpedVAPIEnumsBankAccountType, ownerSpedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, ownerUser NullableFPHSpedVAPIObjectsUsersUser, money float64, ) *FPHSpedVAPIObjectsMoneyBankAccount`

NewFPHSpedVAPIObjectsMoneyBankAccount instantiates a new FPHSpedVAPIObjectsMoneyBankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsMoneyBankAccountWithDefaults

`func NewFPHSpedVAPIObjectsMoneyBankAccountWithDefaults() *FPHSpedVAPIObjectsMoneyBankAccount`

NewFPHSpedVAPIObjectsMoneyBankAccountWithDefaults instantiates a new FPHSpedVAPIObjectsMoneyBankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetId(v int32)`

SetId sets Id field to given value.


### GetIban

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetIban(v string)`

SetIban sets Iban field to given value.


### SetIbanNil

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *FPHSpedVAPIObjectsMoneyBankAccount) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetIbanFormated

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetIbanFormated() string`

GetIbanFormated returns the IbanFormated field if non-nil, zero value otherwise.

### GetIbanFormatedOk

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetIbanFormatedOk() (*string, bool)`

GetIbanFormatedOk returns a tuple with the IbanFormated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanFormated

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetIbanFormated(v string)`

SetIbanFormated sets IbanFormated field to given value.


### SetIbanFormatedNil

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetIbanFormatedNil(b bool)`

 SetIbanFormatedNil sets the value for IbanFormated to be an explicit nil

### UnsetIbanFormated
`func (o *FPHSpedVAPIObjectsMoneyBankAccount) UnsetIbanFormated()`

UnsetIbanFormated ensures that no value is present for IbanFormated, not even an explicit nil
### GetCurrency

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *FPHSpedVAPIObjectsMoneyBankAccount) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetDescription

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FPHSpedVAPIObjectsMoneyBankAccount) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetType() FPHSpedVAPIEnumsBankAccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetTypeOk() (*FPHSpedVAPIEnumsBankAccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetType(v FPHSpedVAPIEnumsBankAccountType)`

SetType sets Type field to given value.


### GetOwnerSpedition

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetOwnerSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetOwnerSpedition returns the OwnerSpedition field if non-nil, zero value otherwise.

### GetOwnerSpeditionOk

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetOwnerSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetOwnerSpeditionOk returns a tuple with the OwnerSpedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerSpedition

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetOwnerSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetOwnerSpedition sets OwnerSpedition field to given value.


### SetOwnerSpeditionNil

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetOwnerSpeditionNil(b bool)`

 SetOwnerSpeditionNil sets the value for OwnerSpedition to be an explicit nil

### UnsetOwnerSpedition
`func (o *FPHSpedVAPIObjectsMoneyBankAccount) UnsetOwnerSpedition()`

UnsetOwnerSpedition ensures that no value is present for OwnerSpedition, not even an explicit nil
### GetOwnerUser

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetOwnerUser() FPHSpedVAPIObjectsUsersUser`

GetOwnerUser returns the OwnerUser field if non-nil, zero value otherwise.

### GetOwnerUserOk

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetOwnerUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetOwnerUserOk returns a tuple with the OwnerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUser

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetOwnerUser(v FPHSpedVAPIObjectsUsersUser)`

SetOwnerUser sets OwnerUser field to given value.


### SetOwnerUserNil

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetOwnerUserNil(b bool)`

 SetOwnerUserNil sets the value for OwnerUser to be an explicit nil

### UnsetOwnerUser
`func (o *FPHSpedVAPIObjectsMoneyBankAccount) UnsetOwnerUser()`

UnsetOwnerUser ensures that no value is present for OwnerUser, not even an explicit nil
### GetMoney

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetMoney() float64`

GetMoney returns the Money field if non-nil, zero value otherwise.

### GetMoneyOk

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) GetMoneyOk() (*float64, bool)`

GetMoneyOk returns a tuple with the Money field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoney

`func (o *FPHSpedVAPIObjectsMoneyBankAccount) SetMoney(v float64)`

SetMoney sets Money field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


