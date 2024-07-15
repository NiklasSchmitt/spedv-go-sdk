# \LiveAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1LiveOnlineuserGet**](LiveAPI.md#V1LiveOnlineuserGet) | **Get** /v1/live/onlineuser | 
[**V1LiveOnlineusersGet**](LiveAPI.md#V1LiveOnlineusersGet) | **Get** /v1/live/onlineusers | 
[**V1LivePositionsGet**](LiveAPI.md#V1LivePositionsGet) | **Get** /v1/live/positions | 



## V1LiveOnlineuserGet

> FPHSpedVAPIObjectsLiveConvoyInfo V1LiveOnlineuserGet(ctx).Execute()



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
	resp, r, err := apiClient.LiveAPI.V1LiveOnlineuserGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveAPI.V1LiveOnlineuserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LiveOnlineuserGet`: FPHSpedVAPIObjectsLiveConvoyInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveAPI.V1LiveOnlineuserGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1LiveOnlineuserGetRequest struct via the builder pattern


### Return type

[**FPHSpedVAPIObjectsLiveConvoyInfo**](FPHSpedVAPIObjectsLiveConvoyInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LiveOnlineusersGet

> []FPHSpedVAPIObjectsLiveConvoyInfo V1LiveOnlineusersGet(ctx).Execute()



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
	resp, r, err := apiClient.LiveAPI.V1LiveOnlineusersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveAPI.V1LiveOnlineusersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LiveOnlineusersGet`: []FPHSpedVAPIObjectsLiveConvoyInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveAPI.V1LiveOnlineusersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1LiveOnlineusersGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsLiveConvoyInfo**](FPHSpedVAPIObjectsLiveConvoyInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LivePositionsGet

> []FPHSpedVAPIServerCommunicationHelperRESTUserPosition V1LivePositionsGet(ctx).SpedId(spedId).IncludeSinglePlayer(includeSinglePlayer).Execute()



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
	spedId := int32(56) // int32 |  (optional)
	includeSinglePlayer := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveAPI.V1LivePositionsGet(context.Background()).SpedId(spedId).IncludeSinglePlayer(includeSinglePlayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveAPI.V1LivePositionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LivePositionsGet`: []FPHSpedVAPIServerCommunicationHelperRESTUserPosition
	fmt.Fprintf(os.Stdout, "Response from `LiveAPI.V1LivePositionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1LivePositionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spedId** | **int32** |  | 
 **includeSinglePlayer** | **bool** |  | [default to false]

### Return type

[**[]FPHSpedVAPIServerCommunicationHelperRESTUserPosition**](FPHSpedVAPIServerCommunicationHelperRESTUserPosition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

