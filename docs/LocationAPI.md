# \LocationAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1LocationQueryGet**](LocationAPI.md#V1LocationQueryGet) | **Get** /v1/location/{query} | 
[**V1LocationRealrouteStartCityIdDestinationCityIdGet**](LocationAPI.md#V1LocationRealrouteStartCityIdDestinationCityIdGet) | **Get** /v1/location/realroute/{startCityId}/{destinationCityId} | 
[**V1LocationScsmapDistanceStartBranchIdDestinationBranchIdGet**](LocationAPI.md#V1LocationScsmapDistanceStartBranchIdDestinationBranchIdGet) | **Get** /v1/location/scsmap/distance/{startBranchId}/{destinationBranchId} | 



## V1LocationQueryGet

> BingMapsRESTToolkitLocation V1LocationQueryGet(ctx, query).Execute()



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
	query := "query_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationAPI.V1LocationQueryGet(context.Background(), query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationAPI.V1LocationQueryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LocationQueryGet`: BingMapsRESTToolkitLocation
	fmt.Fprintf(os.Stdout, "Response from `LocationAPI.V1LocationQueryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**query** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1LocationQueryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BingMapsRESTToolkitLocation**](BingMapsRESTToolkitLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LocationRealrouteStartCityIdDestinationCityIdGet

> FPHSpedVAPIObjectsMapsRealRoute V1LocationRealrouteStartCityIdDestinationCityIdGet(ctx, startCityId, destinationCityId).Execute()



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
	startCityId := int32(56) // int32 | 
	destinationCityId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationAPI.V1LocationRealrouteStartCityIdDestinationCityIdGet(context.Background(), startCityId, destinationCityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationAPI.V1LocationRealrouteStartCityIdDestinationCityIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LocationRealrouteStartCityIdDestinationCityIdGet`: FPHSpedVAPIObjectsMapsRealRoute
	fmt.Fprintf(os.Stdout, "Response from `LocationAPI.V1LocationRealrouteStartCityIdDestinationCityIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**startCityId** | **int32** |  | 
**destinationCityId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1LocationRealrouteStartCityIdDestinationCityIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FPHSpedVAPIObjectsMapsRealRoute**](FPHSpedVAPIObjectsMapsRealRoute.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LocationScsmapDistanceStartBranchIdDestinationBranchIdGet

> FPHSpedVAPIObjectsLiveSGMDistance V1LocationScsmapDistanceStartBranchIdDestinationBranchIdGet(ctx, startBranchId, destinationBranchId).Execute()



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
	startBranchId := int32(56) // int32 | 
	destinationBranchId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationAPI.V1LocationScsmapDistanceStartBranchIdDestinationBranchIdGet(context.Background(), startBranchId, destinationBranchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationAPI.V1LocationScsmapDistanceStartBranchIdDestinationBranchIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LocationScsmapDistanceStartBranchIdDestinationBranchIdGet`: FPHSpedVAPIObjectsLiveSGMDistance
	fmt.Fprintf(os.Stdout, "Response from `LocationAPI.V1LocationScsmapDistanceStartBranchIdDestinationBranchIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**startBranchId** | **int32** |  | 
**destinationBranchId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1LocationScsmapDistanceStartBranchIdDestinationBranchIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FPHSpedVAPIObjectsLiveSGMDistance**](FPHSpedVAPIObjectsLiveSGMDistance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

