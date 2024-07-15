# \MaintenanceAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MaintenanceTrailerGet**](MaintenanceAPI.md#V1MaintenanceTrailerGet) | **Get** /v1/maintenance/trailer | 
[**V1MaintenanceTruckGet**](MaintenanceAPI.md#V1MaintenanceTruckGet) | **Get** /v1/maintenance/truck | 



## V1MaintenanceTrailerGet

> []FPHSpedVAPIObjectsKontorTrailerMaintenance V1MaintenanceTrailerGet(ctx).ShowAll(showAll).Execute()



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
	showAll := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceAPI.V1MaintenanceTrailerGet(context.Background()).ShowAll(showAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceAPI.V1MaintenanceTrailerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MaintenanceTrailerGet`: []FPHSpedVAPIObjectsKontorTrailerMaintenance
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceAPI.V1MaintenanceTrailerGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MaintenanceTrailerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showAll** | **bool** |  | [default to false]

### Return type

[**[]FPHSpedVAPIObjectsKontorTrailerMaintenance**](FPHSpedVAPIObjectsKontorTrailerMaintenance.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MaintenanceTruckGet

> []FPHSpedVAPIObjectsSpeditionsTruckMaintenance V1MaintenanceTruckGet(ctx).ShowAll(showAll).Execute()



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
	showAll := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceAPI.V1MaintenanceTruckGet(context.Background()).ShowAll(showAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceAPI.V1MaintenanceTruckGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MaintenanceTruckGet`: []FPHSpedVAPIObjectsSpeditionsTruckMaintenance
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceAPI.V1MaintenanceTruckGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MaintenanceTruckGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showAll** | **bool** |  | [default to false]

### Return type

[**[]FPHSpedVAPIObjectsSpeditionsTruckMaintenance**](FPHSpedVAPIObjectsSpeditionsTruckMaintenance.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

