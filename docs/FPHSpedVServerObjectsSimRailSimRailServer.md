# FPHSpedVServerObjectsSimRailSimRailServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerData** | [**NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo**](FPHSpedVAPIObjectsSimRailSimRailServerInfo.md) |  | 
**LastOnlineState** | **bool** |  | [readonly] 
**CurrentTrainList** | [**[]FPHSpedVServerObjectsSimRailSimRailActiveTrain**](FPHSpedVServerObjectsSimRailSimRailActiveTrain.md) |  | [readonly] 
**DispatchStationList** | [**[]FPHSpedVServerObjectsSimRailSimRailDispatchStation**](FPHSpedVServerObjectsSimRailSimRailDispatchStation.md) |  | [readonly] 

## Methods

### NewFPHSpedVServerObjectsSimRailSimRailServer

`func NewFPHSpedVServerObjectsSimRailSimRailServer(serverData NullableFPHSpedVAPIObjectsSimRailSimRailServerInfo, lastOnlineState bool, currentTrainList []FPHSpedVServerObjectsSimRailSimRailActiveTrain, dispatchStationList []FPHSpedVServerObjectsSimRailSimRailDispatchStation, ) *FPHSpedVServerObjectsSimRailSimRailServer`

NewFPHSpedVServerObjectsSimRailSimRailServer instantiates a new FPHSpedVServerObjectsSimRailSimRailServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVServerObjectsSimRailSimRailServerWithDefaults

`func NewFPHSpedVServerObjectsSimRailSimRailServerWithDefaults() *FPHSpedVServerObjectsSimRailSimRailServer`

NewFPHSpedVServerObjectsSimRailSimRailServerWithDefaults instantiates a new FPHSpedVServerObjectsSimRailSimRailServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerData

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetServerData() FPHSpedVAPIObjectsSimRailSimRailServerInfo`

GetServerData returns the ServerData field if non-nil, zero value otherwise.

### GetServerDataOk

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetServerDataOk() (*FPHSpedVAPIObjectsSimRailSimRailServerInfo, bool)`

GetServerDataOk returns a tuple with the ServerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerData

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) SetServerData(v FPHSpedVAPIObjectsSimRailSimRailServerInfo)`

SetServerData sets ServerData field to given value.


### SetServerDataNil

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) SetServerDataNil(b bool)`

 SetServerDataNil sets the value for ServerData to be an explicit nil

### UnsetServerData
`func (o *FPHSpedVServerObjectsSimRailSimRailServer) UnsetServerData()`

UnsetServerData ensures that no value is present for ServerData, not even an explicit nil
### GetLastOnlineState

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetLastOnlineState() bool`

GetLastOnlineState returns the LastOnlineState field if non-nil, zero value otherwise.

### GetLastOnlineStateOk

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetLastOnlineStateOk() (*bool, bool)`

GetLastOnlineStateOk returns a tuple with the LastOnlineState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnlineState

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) SetLastOnlineState(v bool)`

SetLastOnlineState sets LastOnlineState field to given value.


### GetCurrentTrainList

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetCurrentTrainList() []FPHSpedVServerObjectsSimRailSimRailActiveTrain`

GetCurrentTrainList returns the CurrentTrainList field if non-nil, zero value otherwise.

### GetCurrentTrainListOk

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetCurrentTrainListOk() (*[]FPHSpedVServerObjectsSimRailSimRailActiveTrain, bool)`

GetCurrentTrainListOk returns a tuple with the CurrentTrainList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTrainList

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) SetCurrentTrainList(v []FPHSpedVServerObjectsSimRailSimRailActiveTrain)`

SetCurrentTrainList sets CurrentTrainList field to given value.


### SetCurrentTrainListNil

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) SetCurrentTrainListNil(b bool)`

 SetCurrentTrainListNil sets the value for CurrentTrainList to be an explicit nil

### UnsetCurrentTrainList
`func (o *FPHSpedVServerObjectsSimRailSimRailServer) UnsetCurrentTrainList()`

UnsetCurrentTrainList ensures that no value is present for CurrentTrainList, not even an explicit nil
### GetDispatchStationList

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetDispatchStationList() []FPHSpedVServerObjectsSimRailSimRailDispatchStation`

GetDispatchStationList returns the DispatchStationList field if non-nil, zero value otherwise.

### GetDispatchStationListOk

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) GetDispatchStationListOk() (*[]FPHSpedVServerObjectsSimRailSimRailDispatchStation, bool)`

GetDispatchStationListOk returns a tuple with the DispatchStationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchStationList

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) SetDispatchStationList(v []FPHSpedVServerObjectsSimRailSimRailDispatchStation)`

SetDispatchStationList sets DispatchStationList field to given value.


### SetDispatchStationListNil

`func (o *FPHSpedVServerObjectsSimRailSimRailServer) SetDispatchStationListNil(b bool)`

 SetDispatchStationListNil sets the value for DispatchStationList to be an explicit nil

### UnsetDispatchStationList
`func (o *FPHSpedVServerObjectsSimRailSimRailServer) UnsetDispatchStationList()`

UnsetDispatchStationList ensures that no value is present for DispatchStationList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


