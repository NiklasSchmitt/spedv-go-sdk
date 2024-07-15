# FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**InitSpedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**ObjectSpedition** | [**NullableFPHSpedVAPIObjectsSpeditionsSpedition**](FPHSpedVAPIObjectsSpeditionsSpedition.md) |  | [readonly] 
**Accepted** | **bool** |  | [readonly] 
**KontorTrailerUsage** | **bool** |  | [readonly] 
**KontorUserAssignment** | **bool** |  | [readonly] 
**KontorAcceptInternDeduction** | **bool** |  | [readonly] 
**KontorKMFee** | **float64** |  | [readonly] 
**AssignTruck** | **bool** |  | [readonly] 
**UseFreeTruck** | **bool** |  | [readonly] 
**UseBranches** | **bool** |  | [readonly] 
**AccessMoney** | **bool** |  | [readonly] 
**ShareConvoyInfo** | **bool** |  | [readonly] 
**FeePercent** | **float64** |  | [readonly] 
**FeeMonthly** | **float64** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry

`func NewFPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry(id int32, initSpedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, objectSpedition NullableFPHSpedVAPIObjectsSpeditionsSpedition, accepted bool, kontorTrailerUsage bool, kontorUserAssignment bool, kontorAcceptInternDeduction bool, kontorKMFee float64, assignTruck bool, useFreeTruck bool, useBranches bool, accessMoney bool, shareConvoyInfo bool, feePercent float64, feeMonthly float64, ) *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry`

NewFPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry instantiates a new FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntryWithDefaults

`func NewFPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntryWithDefaults() *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry`

NewFPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntryWithDefaults instantiates a new FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetId(v int32)`

SetId sets Id field to given value.


### GetInitSpedition

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetInitSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetInitSpedition returns the InitSpedition field if non-nil, zero value otherwise.

### GetInitSpeditionOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetInitSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetInitSpeditionOk returns a tuple with the InitSpedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitSpedition

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetInitSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetInitSpedition sets InitSpedition field to given value.


### SetInitSpeditionNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetInitSpeditionNil(b bool)`

 SetInitSpeditionNil sets the value for InitSpedition to be an explicit nil

### UnsetInitSpedition
`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) UnsetInitSpedition()`

UnsetInitSpedition ensures that no value is present for InitSpedition, not even an explicit nil
### GetObjectSpedition

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetObjectSpedition() FPHSpedVAPIObjectsSpeditionsSpedition`

GetObjectSpedition returns the ObjectSpedition field if non-nil, zero value otherwise.

### GetObjectSpeditionOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetObjectSpeditionOk() (*FPHSpedVAPIObjectsSpeditionsSpedition, bool)`

GetObjectSpeditionOk returns a tuple with the ObjectSpedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSpedition

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetObjectSpedition(v FPHSpedVAPIObjectsSpeditionsSpedition)`

SetObjectSpedition sets ObjectSpedition field to given value.


### SetObjectSpeditionNil

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetObjectSpeditionNil(b bool)`

 SetObjectSpeditionNil sets the value for ObjectSpedition to be an explicit nil

### UnsetObjectSpedition
`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) UnsetObjectSpedition()`

UnsetObjectSpedition ensures that no value is present for ObjectSpedition, not even an explicit nil
### GetAccepted

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.


### GetKontorTrailerUsage

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetKontorTrailerUsage() bool`

GetKontorTrailerUsage returns the KontorTrailerUsage field if non-nil, zero value otherwise.

### GetKontorTrailerUsageOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetKontorTrailerUsageOk() (*bool, bool)`

GetKontorTrailerUsageOk returns a tuple with the KontorTrailerUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKontorTrailerUsage

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetKontorTrailerUsage(v bool)`

SetKontorTrailerUsage sets KontorTrailerUsage field to given value.


### GetKontorUserAssignment

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetKontorUserAssignment() bool`

GetKontorUserAssignment returns the KontorUserAssignment field if non-nil, zero value otherwise.

### GetKontorUserAssignmentOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetKontorUserAssignmentOk() (*bool, bool)`

GetKontorUserAssignmentOk returns a tuple with the KontorUserAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKontorUserAssignment

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetKontorUserAssignment(v bool)`

SetKontorUserAssignment sets KontorUserAssignment field to given value.


### GetKontorAcceptInternDeduction

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetKontorAcceptInternDeduction() bool`

GetKontorAcceptInternDeduction returns the KontorAcceptInternDeduction field if non-nil, zero value otherwise.

### GetKontorAcceptInternDeductionOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetKontorAcceptInternDeductionOk() (*bool, bool)`

GetKontorAcceptInternDeductionOk returns a tuple with the KontorAcceptInternDeduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKontorAcceptInternDeduction

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetKontorAcceptInternDeduction(v bool)`

SetKontorAcceptInternDeduction sets KontorAcceptInternDeduction field to given value.


### GetKontorKMFee

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetKontorKMFee() float64`

GetKontorKMFee returns the KontorKMFee field if non-nil, zero value otherwise.

### GetKontorKMFeeOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetKontorKMFeeOk() (*float64, bool)`

GetKontorKMFeeOk returns a tuple with the KontorKMFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKontorKMFee

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetKontorKMFee(v float64)`

SetKontorKMFee sets KontorKMFee field to given value.


### GetAssignTruck

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetAssignTruck() bool`

GetAssignTruck returns the AssignTruck field if non-nil, zero value otherwise.

### GetAssignTruckOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetAssignTruckOk() (*bool, bool)`

GetAssignTruckOk returns a tuple with the AssignTruck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignTruck

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetAssignTruck(v bool)`

SetAssignTruck sets AssignTruck field to given value.


### GetUseFreeTruck

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetUseFreeTruck() bool`

GetUseFreeTruck returns the UseFreeTruck field if non-nil, zero value otherwise.

### GetUseFreeTruckOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetUseFreeTruckOk() (*bool, bool)`

GetUseFreeTruckOk returns a tuple with the UseFreeTruck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFreeTruck

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetUseFreeTruck(v bool)`

SetUseFreeTruck sets UseFreeTruck field to given value.


### GetUseBranches

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetUseBranches() bool`

GetUseBranches returns the UseBranches field if non-nil, zero value otherwise.

### GetUseBranchesOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetUseBranchesOk() (*bool, bool)`

GetUseBranchesOk returns a tuple with the UseBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBranches

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetUseBranches(v bool)`

SetUseBranches sets UseBranches field to given value.


### GetAccessMoney

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetAccessMoney() bool`

GetAccessMoney returns the AccessMoney field if non-nil, zero value otherwise.

### GetAccessMoneyOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetAccessMoneyOk() (*bool, bool)`

GetAccessMoneyOk returns a tuple with the AccessMoney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMoney

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetAccessMoney(v bool)`

SetAccessMoney sets AccessMoney field to given value.


### GetShareConvoyInfo

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetShareConvoyInfo() bool`

GetShareConvoyInfo returns the ShareConvoyInfo field if non-nil, zero value otherwise.

### GetShareConvoyInfoOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetShareConvoyInfoOk() (*bool, bool)`

GetShareConvoyInfoOk returns a tuple with the ShareConvoyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareConvoyInfo

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetShareConvoyInfo(v bool)`

SetShareConvoyInfo sets ShareConvoyInfo field to given value.


### GetFeePercent

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetFeePercent() float64`

GetFeePercent returns the FeePercent field if non-nil, zero value otherwise.

### GetFeePercentOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetFeePercentOk() (*float64, bool)`

GetFeePercentOk returns a tuple with the FeePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePercent

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetFeePercent(v float64)`

SetFeePercent sets FeePercent field to given value.


### GetFeeMonthly

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetFeeMonthly() float64`

GetFeeMonthly returns the FeeMonthly field if non-nil, zero value otherwise.

### GetFeeMonthlyOk

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) GetFeeMonthlyOk() (*float64, bool)`

GetFeeMonthlyOk returns a tuple with the FeeMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeMonthly

`func (o *FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry) SetFeeMonthly(v float64)`

SetFeeMonthly sets FeeMonthly field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


