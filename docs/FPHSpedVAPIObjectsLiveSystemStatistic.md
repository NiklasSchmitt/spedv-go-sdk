# FPHSpedVAPIObjectsLiveSystemStatistic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastRestart** | **time.Time** |  | [readonly] 
**Lastrefresh** | **time.Time** |  | [readonly] 
**Loadcitys** | [**[]FPHSpedVAPIObjectsLiveSystemStatisticEntry**](FPHSpedVAPIObjectsLiveSystemStatisticEntry.md) |  | [readonly] 
**Destcitys** | [**[]FPHSpedVAPIObjectsLiveSystemStatisticEntry**](FPHSpedVAPIObjectsLiveSystemStatisticEntry.md) |  | [readonly] 
**Freights** | [**[]FPHSpedVAPIObjectsLiveSystemStatisticEntry**](FPHSpedVAPIObjectsLiveSystemStatisticEntry.md) |  | [readonly] 
**Usedmaps** | [**[]FPHSpedVAPIObjectsLiveSystemStatisticEntry**](FPHSpedVAPIObjectsLiveSystemStatisticEntry.md) |  | [readonly] 
**WeeklyStatsEUR** | [**[]FPHSpedVAPIObjectsLiveSystemStatisticEntry**](FPHSpedVAPIObjectsLiveSystemStatisticEntry.md) |  | [readonly] 
**WeeklyStatsKontorEUR** | [**[]FPHSpedVAPIObjectsLiveSystemStatisticEntry**](FPHSpedVAPIObjectsLiveSystemStatisticEntry.md) |  | [readonly] 
**WeeklyStatsExternalEUR** | [**[]FPHSpedVAPIObjectsLiveSystemStatisticEntry**](FPHSpedVAPIObjectsLiveSystemStatisticEntry.md) |  | [readonly] 
**WeeklyStatsKM** | [**[]FPHSpedVAPIObjectsLiveSystemStatisticEntry**](FPHSpedVAPIObjectsLiveSystemStatisticEntry.md) |  | [readonly] 
**Places** | [**FPHSpedVAPIObjectsLiveMyPlaceStructure**](FPHSpedVAPIObjectsLiveMyPlaceStructure.md) |  | [readonly] 
**OnlineUser** | **int32** |  | [readonly] 
**ActiveSpedLastMonth** | **int32** |  | [readonly] 
**ActiveUserLastMonth** | **int32** |  | [readonly] 
**NextTaskID** | **int32** |  | [readonly] 
**ActRAM** | **int64** |  | [readonly] 
**FormatActRAM** | **NullableString** |  | [readonly] 
**TaskLastDay** | **int32** |  | [readonly] 
**ExchangeRate** | **float64** |  | [readonly] 

## Methods

### NewFPHSpedVAPIObjectsLiveSystemStatistic

`func NewFPHSpedVAPIObjectsLiveSystemStatistic(lastRestart time.Time, lastrefresh time.Time, loadcitys []FPHSpedVAPIObjectsLiveSystemStatisticEntry, destcitys []FPHSpedVAPIObjectsLiveSystemStatisticEntry, freights []FPHSpedVAPIObjectsLiveSystemStatisticEntry, usedmaps []FPHSpedVAPIObjectsLiveSystemStatisticEntry, weeklyStatsEUR []FPHSpedVAPIObjectsLiveSystemStatisticEntry, weeklyStatsKontorEUR []FPHSpedVAPIObjectsLiveSystemStatisticEntry, weeklyStatsExternalEUR []FPHSpedVAPIObjectsLiveSystemStatisticEntry, weeklyStatsKM []FPHSpedVAPIObjectsLiveSystemStatisticEntry, places FPHSpedVAPIObjectsLiveMyPlaceStructure, onlineUser int32, activeSpedLastMonth int32, activeUserLastMonth int32, nextTaskID int32, actRAM int64, formatActRAM NullableString, taskLastDay int32, exchangeRate float64, ) *FPHSpedVAPIObjectsLiveSystemStatistic`

NewFPHSpedVAPIObjectsLiveSystemStatistic instantiates a new FPHSpedVAPIObjectsLiveSystemStatistic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPHSpedVAPIObjectsLiveSystemStatisticWithDefaults

`func NewFPHSpedVAPIObjectsLiveSystemStatisticWithDefaults() *FPHSpedVAPIObjectsLiveSystemStatistic`

NewFPHSpedVAPIObjectsLiveSystemStatisticWithDefaults instantiates a new FPHSpedVAPIObjectsLiveSystemStatistic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastRestart

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetLastRestart() time.Time`

GetLastRestart returns the LastRestart field if non-nil, zero value otherwise.

### GetLastRestartOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetLastRestartOk() (*time.Time, bool)`

GetLastRestartOk returns a tuple with the LastRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRestart

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetLastRestart(v time.Time)`

SetLastRestart sets LastRestart field to given value.


### GetLastrefresh

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetLastrefresh() time.Time`

GetLastrefresh returns the Lastrefresh field if non-nil, zero value otherwise.

### GetLastrefreshOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetLastrefreshOk() (*time.Time, bool)`

GetLastrefreshOk returns a tuple with the Lastrefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastrefresh

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetLastrefresh(v time.Time)`

SetLastrefresh sets Lastrefresh field to given value.


### GetLoadcitys

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetLoadcitys() []FPHSpedVAPIObjectsLiveSystemStatisticEntry`

GetLoadcitys returns the Loadcitys field if non-nil, zero value otherwise.

### GetLoadcitysOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetLoadcitysOk() (*[]FPHSpedVAPIObjectsLiveSystemStatisticEntry, bool)`

GetLoadcitysOk returns a tuple with the Loadcitys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadcitys

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetLoadcitys(v []FPHSpedVAPIObjectsLiveSystemStatisticEntry)`

SetLoadcitys sets Loadcitys field to given value.


### SetLoadcitysNil

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetLoadcitysNil(b bool)`

 SetLoadcitysNil sets the value for Loadcitys to be an explicit nil

### UnsetLoadcitys
`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) UnsetLoadcitys()`

UnsetLoadcitys ensures that no value is present for Loadcitys, not even an explicit nil
### GetDestcitys

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetDestcitys() []FPHSpedVAPIObjectsLiveSystemStatisticEntry`

GetDestcitys returns the Destcitys field if non-nil, zero value otherwise.

### GetDestcitysOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetDestcitysOk() (*[]FPHSpedVAPIObjectsLiveSystemStatisticEntry, bool)`

GetDestcitysOk returns a tuple with the Destcitys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestcitys

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetDestcitys(v []FPHSpedVAPIObjectsLiveSystemStatisticEntry)`

SetDestcitys sets Destcitys field to given value.


### SetDestcitysNil

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetDestcitysNil(b bool)`

 SetDestcitysNil sets the value for Destcitys to be an explicit nil

### UnsetDestcitys
`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) UnsetDestcitys()`

UnsetDestcitys ensures that no value is present for Destcitys, not even an explicit nil
### GetFreights

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetFreights() []FPHSpedVAPIObjectsLiveSystemStatisticEntry`

GetFreights returns the Freights field if non-nil, zero value otherwise.

### GetFreightsOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetFreightsOk() (*[]FPHSpedVAPIObjectsLiveSystemStatisticEntry, bool)`

GetFreightsOk returns a tuple with the Freights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreights

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetFreights(v []FPHSpedVAPIObjectsLiveSystemStatisticEntry)`

SetFreights sets Freights field to given value.


### SetFreightsNil

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetFreightsNil(b bool)`

 SetFreightsNil sets the value for Freights to be an explicit nil

### UnsetFreights
`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) UnsetFreights()`

UnsetFreights ensures that no value is present for Freights, not even an explicit nil
### GetUsedmaps

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetUsedmaps() []FPHSpedVAPIObjectsLiveSystemStatisticEntry`

GetUsedmaps returns the Usedmaps field if non-nil, zero value otherwise.

### GetUsedmapsOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetUsedmapsOk() (*[]FPHSpedVAPIObjectsLiveSystemStatisticEntry, bool)`

GetUsedmapsOk returns a tuple with the Usedmaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedmaps

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetUsedmaps(v []FPHSpedVAPIObjectsLiveSystemStatisticEntry)`

SetUsedmaps sets Usedmaps field to given value.


### SetUsedmapsNil

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetUsedmapsNil(b bool)`

 SetUsedmapsNil sets the value for Usedmaps to be an explicit nil

### UnsetUsedmaps
`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) UnsetUsedmaps()`

UnsetUsedmaps ensures that no value is present for Usedmaps, not even an explicit nil
### GetWeeklyStatsEUR

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetWeeklyStatsEUR() []FPHSpedVAPIObjectsLiveSystemStatisticEntry`

GetWeeklyStatsEUR returns the WeeklyStatsEUR field if non-nil, zero value otherwise.

### GetWeeklyStatsEUROk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetWeeklyStatsEUROk() (*[]FPHSpedVAPIObjectsLiveSystemStatisticEntry, bool)`

GetWeeklyStatsEUROk returns a tuple with the WeeklyStatsEUR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyStatsEUR

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetWeeklyStatsEUR(v []FPHSpedVAPIObjectsLiveSystemStatisticEntry)`

SetWeeklyStatsEUR sets WeeklyStatsEUR field to given value.


### SetWeeklyStatsEURNil

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetWeeklyStatsEURNil(b bool)`

 SetWeeklyStatsEURNil sets the value for WeeklyStatsEUR to be an explicit nil

### UnsetWeeklyStatsEUR
`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) UnsetWeeklyStatsEUR()`

UnsetWeeklyStatsEUR ensures that no value is present for WeeklyStatsEUR, not even an explicit nil
### GetWeeklyStatsKontorEUR

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetWeeklyStatsKontorEUR() []FPHSpedVAPIObjectsLiveSystemStatisticEntry`

GetWeeklyStatsKontorEUR returns the WeeklyStatsKontorEUR field if non-nil, zero value otherwise.

### GetWeeklyStatsKontorEUROk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetWeeklyStatsKontorEUROk() (*[]FPHSpedVAPIObjectsLiveSystemStatisticEntry, bool)`

GetWeeklyStatsKontorEUROk returns a tuple with the WeeklyStatsKontorEUR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyStatsKontorEUR

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetWeeklyStatsKontorEUR(v []FPHSpedVAPIObjectsLiveSystemStatisticEntry)`

SetWeeklyStatsKontorEUR sets WeeklyStatsKontorEUR field to given value.


### SetWeeklyStatsKontorEURNil

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetWeeklyStatsKontorEURNil(b bool)`

 SetWeeklyStatsKontorEURNil sets the value for WeeklyStatsKontorEUR to be an explicit nil

### UnsetWeeklyStatsKontorEUR
`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) UnsetWeeklyStatsKontorEUR()`

UnsetWeeklyStatsKontorEUR ensures that no value is present for WeeklyStatsKontorEUR, not even an explicit nil
### GetWeeklyStatsExternalEUR

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetWeeklyStatsExternalEUR() []FPHSpedVAPIObjectsLiveSystemStatisticEntry`

GetWeeklyStatsExternalEUR returns the WeeklyStatsExternalEUR field if non-nil, zero value otherwise.

### GetWeeklyStatsExternalEUROk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetWeeklyStatsExternalEUROk() (*[]FPHSpedVAPIObjectsLiveSystemStatisticEntry, bool)`

GetWeeklyStatsExternalEUROk returns a tuple with the WeeklyStatsExternalEUR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyStatsExternalEUR

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetWeeklyStatsExternalEUR(v []FPHSpedVAPIObjectsLiveSystemStatisticEntry)`

SetWeeklyStatsExternalEUR sets WeeklyStatsExternalEUR field to given value.


### SetWeeklyStatsExternalEURNil

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetWeeklyStatsExternalEURNil(b bool)`

 SetWeeklyStatsExternalEURNil sets the value for WeeklyStatsExternalEUR to be an explicit nil

### UnsetWeeklyStatsExternalEUR
`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) UnsetWeeklyStatsExternalEUR()`

UnsetWeeklyStatsExternalEUR ensures that no value is present for WeeklyStatsExternalEUR, not even an explicit nil
### GetWeeklyStatsKM

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetWeeklyStatsKM() []FPHSpedVAPIObjectsLiveSystemStatisticEntry`

GetWeeklyStatsKM returns the WeeklyStatsKM field if non-nil, zero value otherwise.

### GetWeeklyStatsKMOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetWeeklyStatsKMOk() (*[]FPHSpedVAPIObjectsLiveSystemStatisticEntry, bool)`

GetWeeklyStatsKMOk returns a tuple with the WeeklyStatsKM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyStatsKM

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetWeeklyStatsKM(v []FPHSpedVAPIObjectsLiveSystemStatisticEntry)`

SetWeeklyStatsKM sets WeeklyStatsKM field to given value.


### SetWeeklyStatsKMNil

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetWeeklyStatsKMNil(b bool)`

 SetWeeklyStatsKMNil sets the value for WeeklyStatsKM to be an explicit nil

### UnsetWeeklyStatsKM
`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) UnsetWeeklyStatsKM()`

UnsetWeeklyStatsKM ensures that no value is present for WeeklyStatsKM, not even an explicit nil
### GetPlaces

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetPlaces() FPHSpedVAPIObjectsLiveMyPlaceStructure`

GetPlaces returns the Places field if non-nil, zero value otherwise.

### GetPlacesOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetPlacesOk() (*FPHSpedVAPIObjectsLiveMyPlaceStructure, bool)`

GetPlacesOk returns a tuple with the Places field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaces

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetPlaces(v FPHSpedVAPIObjectsLiveMyPlaceStructure)`

SetPlaces sets Places field to given value.


### GetOnlineUser

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetOnlineUser() int32`

GetOnlineUser returns the OnlineUser field if non-nil, zero value otherwise.

### GetOnlineUserOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetOnlineUserOk() (*int32, bool)`

GetOnlineUserOk returns a tuple with the OnlineUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineUser

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetOnlineUser(v int32)`

SetOnlineUser sets OnlineUser field to given value.


### GetActiveSpedLastMonth

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetActiveSpedLastMonth() int32`

GetActiveSpedLastMonth returns the ActiveSpedLastMonth field if non-nil, zero value otherwise.

### GetActiveSpedLastMonthOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetActiveSpedLastMonthOk() (*int32, bool)`

GetActiveSpedLastMonthOk returns a tuple with the ActiveSpedLastMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSpedLastMonth

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetActiveSpedLastMonth(v int32)`

SetActiveSpedLastMonth sets ActiveSpedLastMonth field to given value.


### GetActiveUserLastMonth

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetActiveUserLastMonth() int32`

GetActiveUserLastMonth returns the ActiveUserLastMonth field if non-nil, zero value otherwise.

### GetActiveUserLastMonthOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetActiveUserLastMonthOk() (*int32, bool)`

GetActiveUserLastMonthOk returns a tuple with the ActiveUserLastMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUserLastMonth

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetActiveUserLastMonth(v int32)`

SetActiveUserLastMonth sets ActiveUserLastMonth field to given value.


### GetNextTaskID

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetNextTaskID() int32`

GetNextTaskID returns the NextTaskID field if non-nil, zero value otherwise.

### GetNextTaskIDOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetNextTaskIDOk() (*int32, bool)`

GetNextTaskIDOk returns a tuple with the NextTaskID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTaskID

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetNextTaskID(v int32)`

SetNextTaskID sets NextTaskID field to given value.


### GetActRAM

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetActRAM() int64`

GetActRAM returns the ActRAM field if non-nil, zero value otherwise.

### GetActRAMOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetActRAMOk() (*int64, bool)`

GetActRAMOk returns a tuple with the ActRAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActRAM

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetActRAM(v int64)`

SetActRAM sets ActRAM field to given value.


### GetFormatActRAM

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetFormatActRAM() string`

GetFormatActRAM returns the FormatActRAM field if non-nil, zero value otherwise.

### GetFormatActRAMOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetFormatActRAMOk() (*string, bool)`

GetFormatActRAMOk returns a tuple with the FormatActRAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatActRAM

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetFormatActRAM(v string)`

SetFormatActRAM sets FormatActRAM field to given value.


### SetFormatActRAMNil

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetFormatActRAMNil(b bool)`

 SetFormatActRAMNil sets the value for FormatActRAM to be an explicit nil

### UnsetFormatActRAM
`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) UnsetFormatActRAM()`

UnsetFormatActRAM ensures that no value is present for FormatActRAM, not even an explicit nil
### GetTaskLastDay

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetTaskLastDay() int32`

GetTaskLastDay returns the TaskLastDay field if non-nil, zero value otherwise.

### GetTaskLastDayOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetTaskLastDayOk() (*int32, bool)`

GetTaskLastDayOk returns a tuple with the TaskLastDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskLastDay

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetTaskLastDay(v int32)`

SetTaskLastDay sets TaskLastDay field to given value.


### GetExchangeRate

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *FPHSpedVAPIObjectsLiveSystemStatistic) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


