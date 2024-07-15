# \KontorAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1KontorGameJobsAvailableGet**](KontorAPI.md#V1KontorGameJobsAvailableGet) | **Get** /v1/kontor/{game}/jobs/available | 
[**V1KontorGameJobsGet**](KontorAPI.md#V1KontorGameJobsGet) | **Get** /v1/kontor/{game}/jobs | 
[**V1KontorGameJobsJobidGet**](KontorAPI.md#V1KontorGameJobsJobidGet) | **Get** /v1/kontor/{game}/jobs/{jobid} | 
[**V1KontorGamePartsAvailableGet**](KontorAPI.md#V1KontorGamePartsAvailableGet) | **Get** /v1/kontor/{game}/parts/available | 
[**V1KontorGamePartsGet**](KontorAPI.md#V1KontorGamePartsGet) | **Get** /v1/kontor/{game}/parts | 
[**V1KontorGamePartsPartidGet**](KontorAPI.md#V1KontorGamePartsPartidGet) | **Get** /v1/kontor/{game}/parts/{partid} | 
[**V1KontorGamePartsPartidJobsGet**](KontorAPI.md#V1KontorGamePartsPartidJobsGet) | **Get** /v1/kontor/{game}/parts/{partid}/jobs | 
[**V1KontorGameTrailersGet**](KontorAPI.md#V1KontorGameTrailersGet) | **Get** /v1/kontor/{game}/trailers | 



## V1KontorGameJobsAvailableGet

> []FPHSpedVAPIObjectsKontorJobOffer V1KontorGameJobsAvailableGet(ctx, game).Execute()



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
	game := openapiclient.FPH.SpedV.API.Enums.GameEnum(0) // FPHSpedVAPIEnumsGameEnum |   0 = ETS2  1 = ATS  -1 = NotSet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KontorAPI.V1KontorGameJobsAvailableGet(context.Background(), game).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KontorAPI.V1KontorGameJobsAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1KontorGameJobsAvailableGet`: []FPHSpedVAPIObjectsKontorJobOffer
	fmt.Fprintf(os.Stdout, "Response from `KontorAPI.V1KontorGameJobsAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | [**FPHSpedVAPIEnumsGameEnum**](.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1KontorGameJobsAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsKontorJobOffer**](FPHSpedVAPIObjectsKontorJobOffer.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1KontorGameJobsGet

> []FPHSpedVAPIObjectsKontorJob V1KontorGameJobsGet(ctx, game).ShowFinished(showFinished).Execute()



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
	game := openapiclient.FPH.SpedV.API.Enums.GameEnum(0) // FPHSpedVAPIEnumsGameEnum |   0 = ETS2  1 = ATS  -1 = NotSet
	showFinished := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KontorAPI.V1KontorGameJobsGet(context.Background(), game).ShowFinished(showFinished).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KontorAPI.V1KontorGameJobsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1KontorGameJobsGet`: []FPHSpedVAPIObjectsKontorJob
	fmt.Fprintf(os.Stdout, "Response from `KontorAPI.V1KontorGameJobsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | [**FPHSpedVAPIEnumsGameEnum**](.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1KontorGameJobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showFinished** | **bool** |  | [default to true]

### Return type

[**[]FPHSpedVAPIObjectsKontorJob**](FPHSpedVAPIObjectsKontorJob.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1KontorGameJobsJobidGet

> FPHSpedVAPIObjectsKontorJob V1KontorGameJobsJobidGet(ctx, game, jobid).Execute()



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
	game := openapiclient.FPH.SpedV.API.Enums.GameEnum(0) // FPHSpedVAPIEnumsGameEnum |   0 = ETS2  1 = ATS  -1 = NotSet
	jobid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KontorAPI.V1KontorGameJobsJobidGet(context.Background(), game, jobid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KontorAPI.V1KontorGameJobsJobidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1KontorGameJobsJobidGet`: FPHSpedVAPIObjectsKontorJob
	fmt.Fprintf(os.Stdout, "Response from `KontorAPI.V1KontorGameJobsJobidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | [**FPHSpedVAPIEnumsGameEnum**](.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | 
**jobid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1KontorGameJobsJobidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FPHSpedVAPIObjectsKontorJob**](FPHSpedVAPIObjectsKontorJob.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1KontorGamePartsAvailableGet

> []FPHSpedVAPIObjectsKontorJobPart V1KontorGamePartsAvailableGet(ctx, game).IncludePartnerships(includePartnerships).Execute()



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
	game := openapiclient.FPH.SpedV.API.Enums.GameEnum(0) // FPHSpedVAPIEnumsGameEnum |   0 = ETS2  1 = ATS  -1 = NotSet
	includePartnerships := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KontorAPI.V1KontorGamePartsAvailableGet(context.Background(), game).IncludePartnerships(includePartnerships).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KontorAPI.V1KontorGamePartsAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1KontorGamePartsAvailableGet`: []FPHSpedVAPIObjectsKontorJobPart
	fmt.Fprintf(os.Stdout, "Response from `KontorAPI.V1KontorGamePartsAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | [**FPHSpedVAPIEnumsGameEnum**](.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1KontorGamePartsAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includePartnerships** | **bool** |  | [default to false]

### Return type

[**[]FPHSpedVAPIObjectsKontorJobPart**](FPHSpedVAPIObjectsKontorJobPart.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1KontorGamePartsGet

> []FPHSpedVAPIObjectsKontorJobPart V1KontorGamePartsGet(ctx, game).ShowFinished(showFinished).IncludePartnerships(includePartnerships).Execute()



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
	game := openapiclient.FPH.SpedV.API.Enums.GameEnum(0) // FPHSpedVAPIEnumsGameEnum |   0 = ETS2  1 = ATS  -1 = NotSet
	showFinished := true // bool |  (optional) (default to true)
	includePartnerships := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KontorAPI.V1KontorGamePartsGet(context.Background(), game).ShowFinished(showFinished).IncludePartnerships(includePartnerships).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KontorAPI.V1KontorGamePartsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1KontorGamePartsGet`: []FPHSpedVAPIObjectsKontorJobPart
	fmt.Fprintf(os.Stdout, "Response from `KontorAPI.V1KontorGamePartsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | [**FPHSpedVAPIEnumsGameEnum**](.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1KontorGamePartsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showFinished** | **bool** |  | [default to true]
 **includePartnerships** | **bool** |  | [default to false]

### Return type

[**[]FPHSpedVAPIObjectsKontorJobPart**](FPHSpedVAPIObjectsKontorJobPart.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1KontorGamePartsPartidGet

> FPHSpedVAPIObjectsKontorJobPart V1KontorGamePartsPartidGet(ctx, game, partid).Execute()



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
	game := openapiclient.FPH.SpedV.API.Enums.GameEnum(0) // FPHSpedVAPIEnumsGameEnum |   0 = ETS2  1 = ATS  -1 = NotSet
	partid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KontorAPI.V1KontorGamePartsPartidGet(context.Background(), game, partid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KontorAPI.V1KontorGamePartsPartidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1KontorGamePartsPartidGet`: FPHSpedVAPIObjectsKontorJobPart
	fmt.Fprintf(os.Stdout, "Response from `KontorAPI.V1KontorGamePartsPartidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | [**FPHSpedVAPIEnumsGameEnum**](.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | 
**partid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1KontorGamePartsPartidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FPHSpedVAPIObjectsKontorJobPart**](FPHSpedVAPIObjectsKontorJobPart.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1KontorGamePartsPartidJobsGet

> []FPHSpedVAPIObjectsKontorJob V1KontorGamePartsPartidJobsGet(ctx, game, partid).Execute()



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
	game := openapiclient.FPH.SpedV.API.Enums.GameEnum(0) // FPHSpedVAPIEnumsGameEnum |   0 = ETS2  1 = ATS  -1 = NotSet
	partid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KontorAPI.V1KontorGamePartsPartidJobsGet(context.Background(), game, partid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KontorAPI.V1KontorGamePartsPartidJobsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1KontorGamePartsPartidJobsGet`: []FPHSpedVAPIObjectsKontorJob
	fmt.Fprintf(os.Stdout, "Response from `KontorAPI.V1KontorGamePartsPartidJobsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | [**FPHSpedVAPIEnumsGameEnum**](.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | 
**partid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1KontorGamePartsPartidJobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]FPHSpedVAPIObjectsKontorJob**](FPHSpedVAPIObjectsKontorJob.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1KontorGameTrailersGet

> []FPHSpedVAPIObjectsKontorTrailer V1KontorGameTrailersGet(ctx, game).IncludePartnerships(includePartnerships).Execute()



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
	game := openapiclient.FPH.SpedV.API.Enums.GameEnum(0) // FPHSpedVAPIEnumsGameEnum |   0 = ETS2  1 = ATS  -1 = NotSet
	includePartnerships := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KontorAPI.V1KontorGameTrailersGet(context.Background(), game).IncludePartnerships(includePartnerships).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KontorAPI.V1KontorGameTrailersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1KontorGameTrailersGet`: []FPHSpedVAPIObjectsKontorTrailer
	fmt.Fprintf(os.Stdout, "Response from `KontorAPI.V1KontorGameTrailersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | [**FPHSpedVAPIEnumsGameEnum**](.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1KontorGameTrailersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includePartnerships** | **bool** |  | [default to false]

### Return type

[**[]FPHSpedVAPIObjectsKontorTrailer**](FPHSpedVAPIObjectsKontorTrailer.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

