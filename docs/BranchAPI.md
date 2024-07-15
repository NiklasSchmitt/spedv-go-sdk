# \BranchAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OwnedbranchesBranchidGet**](BranchAPI.md#V1OwnedbranchesBranchidGet) | **Get** /v1/ownedbranches/{branchid} | 
[**V1OwnedbranchesBranchidStatisticsGet**](BranchAPI.md#V1OwnedbranchesBranchidStatisticsGet) | **Get** /v1/ownedbranches/{branchid}/statistics | 
[**V1OwnedbranchesGet**](BranchAPI.md#V1OwnedbranchesGet) | **Get** /v1/ownedbranches | 
[**V1SparepartsGet**](BranchAPI.md#V1SparepartsGet) | **Get** /v1/spareparts | 
[**V1SparepartsSparepartidGet**](BranchAPI.md#V1SparepartsSparepartidGet) | **Get** /v1/spareparts/{sparepartid} | 
[**V1SparepartsSparepartidProvidersGet**](BranchAPI.md#V1SparepartsSparepartidProvidersGet) | **Get** /v1/spareparts/{sparepartid}/providers | 



## V1OwnedbranchesBranchidGet

> FPHSpedVAPIObjectsSpeditionsTruckBranchStatistic V1OwnedbranchesBranchidGet(ctx, branchid).Execute()



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
	branchid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchAPI.V1OwnedbranchesBranchidGet(context.Background(), branchid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchAPI.V1OwnedbranchesBranchidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OwnedbranchesBranchidGet`: FPHSpedVAPIObjectsSpeditionsTruckBranchStatistic
	fmt.Fprintf(os.Stdout, "Response from `BranchAPI.V1OwnedbranchesBranchidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**branchid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OwnedbranchesBranchidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## V1OwnedbranchesBranchidStatisticsGet

> FPHSpedVAPIObjectsSpeditionsTruckBranchStatistic V1OwnedbranchesBranchidStatisticsGet(ctx, branchid).Startdate(startdate).Enddate(enddate).Execute()



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
	branchid := int32(56) // int32 | 
	startdate := time.Now() // time.Time |  (optional)
	enddate := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchAPI.V1OwnedbranchesBranchidStatisticsGet(context.Background(), branchid).Startdate(startdate).Enddate(enddate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchAPI.V1OwnedbranchesBranchidStatisticsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OwnedbranchesBranchidStatisticsGet`: FPHSpedVAPIObjectsSpeditionsTruckBranchStatistic
	fmt.Fprintf(os.Stdout, "Response from `BranchAPI.V1OwnedbranchesBranchidStatisticsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**branchid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OwnedbranchesBranchidStatisticsGetRequest struct via the builder pattern


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


## V1OwnedbranchesGet

> []FPHSpedVAPIObjectsSpeditionsOwnedBranch V1OwnedbranchesGet(ctx).IncludePartnerships(includePartnerships).Execute()



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
	resp, r, err := apiClient.BranchAPI.V1OwnedbranchesGet(context.Background()).IncludePartnerships(includePartnerships).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchAPI.V1OwnedbranchesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OwnedbranchesGet`: []FPHSpedVAPIObjectsSpeditionsOwnedBranch
	fmt.Fprintf(os.Stdout, "Response from `BranchAPI.V1OwnedbranchesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OwnedbranchesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includePartnerships** | **bool** |  | [default to false]

### Return type

[**[]FPHSpedVAPIObjectsSpeditionsOwnedBranch**](FPHSpedVAPIObjectsSpeditionsOwnedBranch.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SparepartsGet

> FPHSpedVAPIObjectsSpeditionsSparePart V1SparepartsGet(ctx).Execute()



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
	resp, r, err := apiClient.BranchAPI.V1SparepartsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchAPI.V1SparepartsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SparepartsGet`: FPHSpedVAPIObjectsSpeditionsSparePart
	fmt.Fprintf(os.Stdout, "Response from `BranchAPI.V1SparepartsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SparepartsGetRequest struct via the builder pattern


### Return type

[**FPHSpedVAPIObjectsSpeditionsSparePart**](FPHSpedVAPIObjectsSpeditionsSparePart.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SparepartsSparepartidGet

> FPHSpedVAPIObjectsSpeditionsSparePart V1SparepartsSparepartidGet(ctx, sparepartid).Execute()



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
	sparepartid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchAPI.V1SparepartsSparepartidGet(context.Background(), sparepartid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchAPI.V1SparepartsSparepartidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SparepartsSparepartidGet`: FPHSpedVAPIObjectsSpeditionsSparePart
	fmt.Fprintf(os.Stdout, "Response from `BranchAPI.V1SparepartsSparepartidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sparepartid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SparepartsSparepartidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FPHSpedVAPIObjectsSpeditionsSparePart**](FPHSpedVAPIObjectsSpeditionsSparePart.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SparepartsSparepartidProvidersGet

> FPHSpedVAPIObjectsSpeditionsSparePart V1SparepartsSparepartidProvidersGet(ctx, sparepartid).Execute()



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
	sparepartid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchAPI.V1SparepartsSparepartidProvidersGet(context.Background(), sparepartid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchAPI.V1SparepartsSparepartidProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SparepartsSparepartidProvidersGet`: FPHSpedVAPIObjectsSpeditionsSparePart
	fmt.Fprintf(os.Stdout, "Response from `BranchAPI.V1SparepartsSparepartidProvidersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sparepartid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SparepartsSparepartidProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FPHSpedVAPIObjectsSpeditionsSparePart**](FPHSpedVAPIObjectsSpeditionsSparePart.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

