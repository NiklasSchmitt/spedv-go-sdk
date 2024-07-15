# \TruckAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TrucksGet**](TruckAPI.md#V1TrucksGet) | **Get** /v1/trucks | 
[**V1TrucksTruckidStatisticsGet**](TruckAPI.md#V1TrucksTruckidStatisticsGet) | **Get** /v1/trucks/{truckid}/statistics | 
[**V1TrucktypesGet**](TruckAPI.md#V1TrucktypesGet) | **Get** /v1/trucktypes | 



## V1TrucksGet

> []FPHSpedVAPIObjectsSpeditionsTruck V1TrucksGet(ctx).IncludePartnerships(includePartnerships).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/NiklasSchmitt/spedv-go-sdk"
)

func main() {
	includePartnerships := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TruckAPI.V1TrucksGet(context.Background()).IncludePartnerships(includePartnerships).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TruckAPI.V1TrucksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TrucksGet`: []FPHSpedVAPIObjectsSpeditionsTruck
	fmt.Fprintf(os.Stdout, "Response from `TruckAPI.V1TrucksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TrucksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includePartnerships** | **bool** |  | [default to false]

### Return type

[**[]FPHSpedVAPIObjectsSpeditionsTruck**](FPHSpedVAPIObjectsSpeditionsTruck.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TrucksTruckidStatisticsGet

> FPHSpedVAPIObjectsSpeditionsTruckBranchStatistic V1TrucksTruckidStatisticsGet(ctx, truckid).Startdate(startdate).Enddate(enddate).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/NiklasSchmitt/spedv-go-sdk"
)

func main() {
	truckid := int32(56) // int32 | 
	startdate := time.Now() // time.Time |  (optional)
	enddate := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TruckAPI.V1TrucksTruckidStatisticsGet(context.Background(), truckid).Startdate(startdate).Enddate(enddate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TruckAPI.V1TrucksTruckidStatisticsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TrucksTruckidStatisticsGet`: FPHSpedVAPIObjectsSpeditionsTruckBranchStatistic
	fmt.Fprintf(os.Stdout, "Response from `TruckAPI.V1TrucksTruckidStatisticsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**truckid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TrucksTruckidStatisticsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startdate** | **time.Time** |  | 
 **enddate** | **time.Time** |  | 

### Return type

[**FPHSpedVAPIObjectsSpeditionsTruckBranchStatistic**](FPHSpedVAPIObjectsSpeditionsTruckBranchStatistic.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TrucktypesGet

> []FPHSpedVAPIObjectsSpeditionsTruckType V1TrucktypesGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/NiklasSchmitt/spedv-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TruckAPI.V1TrucktypesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TruckAPI.V1TrucktypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TrucktypesGet`: []FPHSpedVAPIObjectsSpeditionsTruckType
	fmt.Fprintf(os.Stdout, "Response from `TruckAPI.V1TrucktypesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1TrucktypesGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsSpeditionsTruckType**](FPHSpedVAPIObjectsSpeditionsTruckType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

