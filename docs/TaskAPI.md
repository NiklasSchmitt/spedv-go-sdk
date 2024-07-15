# \TaskAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TasksDetailsGet**](TaskAPI.md#V1TasksDetailsGet) | **Get** /v1/tasks/details | 
[**V1TasksGet**](TaskAPI.md#V1TasksGet) | **Get** /v1/tasks | 
[**V1TasksTaskIdDetailsGet**](TaskAPI.md#V1TasksTaskIdDetailsGet) | **Get** /v1/tasks/{taskId}/details | 
[**V1TasksTaskIdDrivedataGet**](TaskAPI.md#V1TasksTaskIdDrivedataGet) | **Get** /v1/tasks/{taskId}/drivedata | 



## V1TasksDetailsGet

> []FPHSpedVAPIObjectsSpeditionsTask V1TasksDetailsGet(ctx).TaskIds(taskIds).Execute()



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
	taskIds := []int32{int32(123)} // []int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.V1TasksDetailsGet(context.Background()).TaskIds(taskIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.V1TasksDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TasksDetailsGet`: []FPHSpedVAPIObjectsSpeditionsTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.V1TasksDetailsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TasksDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskIds** | **[]int32** |  | 

### Return type

[**[]FPHSpedVAPIObjectsSpeditionsTask**](FPHSpedVAPIObjectsSpeditionsTask.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TasksGet

> []FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo V1TasksGet(ctx).Start(start).End(end).Game(game).UserId(userId).Execute()



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
	start := time.Now() // time.Time |  (optional)
	end := time.Now() // time.Time |  (optional)
	game := openapiclient.FPH.SpedV.API.Enums.GameEnum(0) // FPHSpedVAPIEnumsGameEnum |   0 = ETS2  1 = ATS  -1 = NotSet (optional) (default to -1)
	userId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.V1TasksGet(context.Background()).Start(start).End(end).Game(game).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.V1TasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TasksGet`: []FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.V1TasksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 
 **game** | [**FPHSpedVAPIEnumsGameEnum**](FPHSpedVAPIEnumsGameEnum.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | [default to -1]
 **userId** | **int32** |  | 

### Return type

[**[]FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo**](FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TasksTaskIdDetailsGet

> FPHSpedVAPIObjectsSpeditionsTask V1TasksTaskIdDetailsGet(ctx, taskId).Execute()



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
	taskId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.V1TasksTaskIdDetailsGet(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.V1TasksTaskIdDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TasksTaskIdDetailsGet`: FPHSpedVAPIObjectsSpeditionsTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.V1TasksTaskIdDetailsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TasksTaskIdDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FPHSpedVAPIObjectsSpeditionsTask**](FPHSpedVAPIObjectsSpeditionsTask.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TasksTaskIdDrivedataGet

> []FPHSpedVAPIObjectsSpeditionsTaskDriveData V1TasksTaskIdDrivedataGet(ctx, taskId).Execute()



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
	taskId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.V1TasksTaskIdDrivedataGet(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.V1TasksTaskIdDrivedataGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TasksTaskIdDrivedataGet`: []FPHSpedVAPIObjectsSpeditionsTaskDriveData
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.V1TasksTaskIdDrivedataGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TasksTaskIdDrivedataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsSpeditionsTaskDriveData**](FPHSpedVAPIObjectsSpeditionsTaskDriveData.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

