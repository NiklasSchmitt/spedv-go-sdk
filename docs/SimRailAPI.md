# \SimRailAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SimrailServerGet**](SimRailAPI.md#V1SimrailServerGet) | **Get** /v1/simrail/server | 
[**V1SimrailServerIdDispatchstationsGet**](SimRailAPI.md#V1SimrailServerIdDispatchstationsGet) | **Get** /v1/simrail/server/{id}/dispatchstations | 
[**V1SimrailServerIdDispatchstationsNameGet**](SimRailAPI.md#V1SimrailServerIdDispatchstationsNameGet) | **Get** /v1/simrail/server/{id}/dispatchstations/{name} | 
[**V1SimrailServerIdGet**](SimRailAPI.md#V1SimrailServerIdGet) | **Get** /v1/simrail/server/{id} | 
[**V1SimrailServerIdTrainsGet**](SimRailAPI.md#V1SimrailServerIdTrainsGet) | **Get** /v1/simrail/server/{id}/trains | 
[**V1SimrailServerIdTrainsNumberGet**](SimRailAPI.md#V1SimrailServerIdTrainsNumberGet) | **Get** /v1/simrail/server/{id}/trains/{number} | 



## V1SimrailServerGet

> []FPHSpedVServerObjectsSimRailSimRailServer V1SimrailServerGet(ctx).Execute()



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
	resp, r, err := apiClient.SimRailAPI.V1SimrailServerGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimRailAPI.V1SimrailServerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SimrailServerGet`: []FPHSpedVServerObjectsSimRailSimRailServer
	fmt.Fprintf(os.Stdout, "Response from `SimRailAPI.V1SimrailServerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SimrailServerGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVServerObjectsSimRailSimRailServer**](FPHSpedVServerObjectsSimRailSimRailServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SimrailServerIdDispatchstationsGet

> []FPHSpedVServerObjectsSimRailSimRailDispatchStation V1SimrailServerIdDispatchstationsGet(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimRailAPI.V1SimrailServerIdDispatchstationsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimRailAPI.V1SimrailServerIdDispatchstationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SimrailServerIdDispatchstationsGet`: []FPHSpedVServerObjectsSimRailSimRailDispatchStation
	fmt.Fprintf(os.Stdout, "Response from `SimRailAPI.V1SimrailServerIdDispatchstationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SimrailServerIdDispatchstationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVServerObjectsSimRailSimRailDispatchStation**](FPHSpedVServerObjectsSimRailSimRailDispatchStation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SimrailServerIdDispatchstationsNameGet

> FPHSpedVServerObjectsSimRailSimRailDispatchStation V1SimrailServerIdDispatchstationsNameGet(ctx, id, name).Execute()



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
	id := "id_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimRailAPI.V1SimrailServerIdDispatchstationsNameGet(context.Background(), id, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimRailAPI.V1SimrailServerIdDispatchstationsNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SimrailServerIdDispatchstationsNameGet`: FPHSpedVServerObjectsSimRailSimRailDispatchStation
	fmt.Fprintf(os.Stdout, "Response from `SimRailAPI.V1SimrailServerIdDispatchstationsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SimrailServerIdDispatchstationsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FPHSpedVServerObjectsSimRailSimRailDispatchStation**](FPHSpedVServerObjectsSimRailSimRailDispatchStation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SimrailServerIdGet

> FPHSpedVServerObjectsSimRailSimRailServer V1SimrailServerIdGet(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimRailAPI.V1SimrailServerIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimRailAPI.V1SimrailServerIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SimrailServerIdGet`: FPHSpedVServerObjectsSimRailSimRailServer
	fmt.Fprintf(os.Stdout, "Response from `SimRailAPI.V1SimrailServerIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SimrailServerIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FPHSpedVServerObjectsSimRailSimRailServer**](FPHSpedVServerObjectsSimRailSimRailServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SimrailServerIdTrainsGet

> []FPHSpedVServerObjectsSimRailSimRailActiveTrain V1SimrailServerIdTrainsGet(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimRailAPI.V1SimrailServerIdTrainsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimRailAPI.V1SimrailServerIdTrainsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SimrailServerIdTrainsGet`: []FPHSpedVServerObjectsSimRailSimRailActiveTrain
	fmt.Fprintf(os.Stdout, "Response from `SimRailAPI.V1SimrailServerIdTrainsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SimrailServerIdTrainsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVServerObjectsSimRailSimRailActiveTrain**](FPHSpedVServerObjectsSimRailSimRailActiveTrain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SimrailServerIdTrainsNumberGet

> FPHSpedVServerObjectsSimRailSimRailActiveTrain V1SimrailServerIdTrainsNumberGet(ctx, id, number).Execute()



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
	id := "id_example" // string | 
	number := "number_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimRailAPI.V1SimrailServerIdTrainsNumberGet(context.Background(), id, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimRailAPI.V1SimrailServerIdTrainsNumberGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SimrailServerIdTrainsNumberGet`: FPHSpedVServerObjectsSimRailSimRailActiveTrain
	fmt.Fprintf(os.Stdout, "Response from `SimRailAPI.V1SimrailServerIdTrainsNumberGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**number** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SimrailServerIdTrainsNumberGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FPHSpedVServerObjectsSimRailSimRailActiveTrain**](FPHSpedVServerObjectsSimRailSimRailActiveTrain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

