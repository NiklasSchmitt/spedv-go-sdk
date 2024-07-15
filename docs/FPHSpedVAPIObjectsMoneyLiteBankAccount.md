# FPHSpedVAPIObjectsMoneyLiteBankAccount

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

## Methods

### NewFPHSpedVAPIObjectsMoneyLiteBankAccount

`func NewFPHSpedVAPIObjectsMoneyLiteBankAccount(id int32, iban NullableString, ibanFormated NullableString, currency NullableString, description NullableString, type_ FPHSpedVAPIEnumsBankAccountType, ownerSpedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, ownerUser NullableFPHSpedVAPIObjectsUsersUser, ) *FPHSpedVAPIObjectsMoneyLiteBankAccount`

NewFPHSpedVAPIObjectsMoneyLiteBankAccount instantiates a new FPHSpedVAPIObjectsMoneyLiteBankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsMoneyLiteBankAccountWithDefaults

`func NewFPHSpedVAPIObjectsMoneyLiteBankAccountWithDefaults() *FPHSpedVAPIObjectsMoneyLiteBankAccount`

NewFPHSpedVAPIObjectsMoneyLiteBankAccountWithDefaults instantiates a new FPHSpedVAPIObjectsMoneyLiteBankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetId(v int32)`

SetId sets Id field to given value.


### GetIban

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetIban(v string)`

SetIban sets Iban field to given value.


### SetIbanNil

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetIbanFormated

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetIbanFormated() string`

GetIbanFormated returns the IbanFormated field if non-nil, zero value otherwise.

### GetIbanFormatedOk

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetIbanFormatedOk() (*string, bool)`

GetIbanFormatedOk returns a tuple with the IbanFormated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanFormated

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetIbanFormated(v string)`

SetIbanFormated sets IbanFormated field to given value.


### SetIbanFormatedNil

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetIbanFormatedNil(b bool)`

 SetIbanFormatedNil sets the value for IbanFormated to be an explicit nil

### UnsetIbanFormated
`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) UnsetIbanFormated()`

UnsetIbanFormated ensures that no value is present for IbanFormated, not even an explicit nil
### GetCurrency

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetDescription

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetType() FPHSpedVAPIEnumsBankAccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetTypeOk() (*FPHSpedVAPIEnumsBankAccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetType(v FPHSpedVAPIEnumsBankAccountType)`

SetType sets Type field to given value.


### GetOwnerSpedition

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetOwnerSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetOwnerSpedition returns the OwnerSpedition field if non-nil, zero value otherwise.

### GetOwnerSpeditionOk

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetOwnerSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetOwnerSpeditionOk returns a tuple with the OwnerSpedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerSpedition

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetOwnerSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetOwnerSpedition sets OwnerSpedition field to given value.


### SetOwnerSpeditionNil

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetOwnerSpeditionNil(b bool)`

 SetOwnerSpeditionNil sets the value for OwnerSpedition to be an explicit nil

### UnsetOwnerSpedition
`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) UnsetOwnerSpedition()`

UnsetOwnerSpedition ensures that no value is present for OwnerSpedition, not even an explicit nil
### GetOwnerUser

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetOwnerUser() FPHSpedVAPIObjectsUsersUser`

GetOwnerUser returns the OwnerUser field if non-nil, zero value otherwise.

### GetOwnerUserOk

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) GetOwnerUserOk() (*FPHSpedVAPIObjectsUsersUser, bool)`

GetOwnerUserOk returns a tuple with the OwnerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUser

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetOwnerUser(v FPHSpedVAPIObjectsUsersUser)`

SetOwnerUser sets OwnerUser field to given value.


### SetOwnerUserNil

`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) SetOwnerUserNil(b bool)`

 SetOwnerUserNil sets the value for OwnerUser to be an explicit nil

### UnsetOwnerUser
`func (o *FPHSpedVAPIObjectsMoneyLiteBankAccount) UnsetOwnerUser()`

UnsetOwnerUser ensures that no value is present for OwnerUser, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


