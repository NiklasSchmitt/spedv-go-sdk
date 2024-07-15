# \UserAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UserGet**](UserAPI.md#V1UserGet) | **Get** /v1/user | 
[**V1UserprofilesGet**](UserAPI.md#V1UserprofilesGet) | **Get** /v1/userprofiles | 
[**V1UsersGet**](UserAPI.md#V1UsersGet) | **Get** /v1/users | 
[**V1UsersUserIdChangelogGet**](UserAPI.md#V1UsersUserIdChangelogGet) | **Get** /v1/users/{userId}/changelog | 
[**V1UsersUserIdOmsiTripsGet**](UserAPI.md#V1UsersUserIdOmsiTripsGet) | **Get** /v1/users/{userId}/omsiTrips | 
[**V1UsersUserIdProfileGet**](UserAPI.md#V1UsersUserIdProfileGet) | **Get** /v1/users/{userId}/profile | 
[**V1UsersUserIdTasksGet**](UserAPI.md#V1UsersUserIdTasksGet) | **Get** /v1/users/{userId}/tasks | 
[**V1VacationsGet**](UserAPI.md#V1VacationsGet) | **Get** /v1/vacations | 



## V1UserGet

> FPHSpedVAPIObjectsUsersFullUserInfo V1UserGet(ctx).Execute()



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
	resp, r, err := apiClient.UserAPI.V1UserGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UserGet`: FPHSpedVAPIObjectsUsersFullUserInfo
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UserGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserGetRequest struct via the builder pattern


### Return type

[**FPHSpedVAPIObjectsUsersFullUserInfo**](FPHSpedVAPIObjectsUsersFullUserInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UserprofilesGet

> []FPHSpedVAPIObjectsUsersUserProfile V1UserprofilesGet(ctx).Ids(ids).Execute()



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
	ids := []int32{int32(123)} // []int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.V1UserprofilesGet(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UserprofilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UserprofilesGet`: []FPHSpedVAPIObjectsUsersUserProfile
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UserprofilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1UserprofilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int32** |  | 

### Return type

[**[]FPHSpedVAPIObjectsUsersUserProfile**](FPHSpedVAPIObjectsUsersUserProfile.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersGet

> []FPHSpedVAPIObjectsUsersUserLiteInfo V1UsersGet(ctx).IncludeKontorPartnerships(includeKontorPartnerships).IncludeTruckPartnerships(includeTruckPartnerships).Execute()



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
	includeKontorPartnerships := true // bool |  (optional) (default to false)
	includeTruckPartnerships := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.V1UsersGet(context.Background()).IncludeKontorPartnerships(includeKontorPartnerships).IncludeTruckPartnerships(includeTruckPartnerships).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UsersGet`: []FPHSpedVAPIObjectsUsersUserLiteInfo
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeKontorPartnerships** | **bool** |  | [default to false]
 **includeTruckPartnerships** | **bool** |  | [default to false]

### Return type

[**[]FPHSpedVAPIObjectsUsersUserLiteInfo**](FPHSpedVAPIObjectsUsersUserLiteInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersUserIdChangelogGet

> []FPHSpedVAPIObjectsUsersUserChangeLogEntry V1UsersUserIdChangelogGet(ctx, userId).Execute()



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
	userId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.V1UsersUserIdChangelogGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UsersUserIdChangelogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UsersUserIdChangelogGet`: []FPHSpedVAPIObjectsUsersUserChangeLogEntry
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UsersUserIdChangelogGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersUserIdChangelogGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsUsersUserChangeLogEntry**](FPHSpedVAPIObjectsUsersUserChangeLogEntry.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersUserIdOmsiTripsGet

> []FPHSpedVAPIObjectsOMSIDrivenTrip V1UsersUserIdOmsiTripsGet(ctx, userId).Start(start).End(end).Execute()



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
	userId := int32(56) // int32 | 
	start := time.Now() // time.Time |  (optional)
	end := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.V1UsersUserIdOmsiTripsGet(context.Background(), userId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UsersUserIdOmsiTripsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UsersUserIdOmsiTripsGet`: []FPHSpedVAPIObjectsOMSIDrivenTrip
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UsersUserIdOmsiTripsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersUserIdOmsiTripsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 

### Return type

[**[]FPHSpedVAPIObjectsOMSIDrivenTrip**](FPHSpedVAPIObjectsOMSIDrivenTrip.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersUserIdProfileGet

> FPHSpedVAPIObjectsUsersUserProfile V1UsersUserIdProfileGet(ctx, userId).Execute()



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
	userId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.V1UsersUserIdProfileGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UsersUserIdProfileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UsersUserIdProfileGet`: FPHSpedVAPIObjectsUsersUserProfile
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UsersUserIdProfileGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersUserIdProfileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FPHSpedVAPIObjectsUsersUserProfile**](FPHSpedVAPIObjectsUsersUserProfile.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersUserIdTasksGet

> []FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo V1UsersUserIdTasksGet(ctx, userId).Start(start).End(end).Game(game).OnlyCurrentSped(onlyCurrentSped).Execute()



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
	userId := int32(56) // int32 | 
	start := time.Now() // time.Time |  (optional)
	end := time.Now() // time.Time |  (optional)
	game := openapiclient.FPH.SpedV.API.Enums.GameEnum(0) // FPHSpedVAPIEnumsGameEnum |   0 = ETS2  1 = ATS  -1 = NotSet (optional) (default to -1)
	onlyCurrentSped := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.V1UsersUserIdTasksGet(context.Background(), userId).Start(start).End(end).Game(game).OnlyCurrentSped(onlyCurrentSped).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UsersUserIdTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UsersUserIdTasksGet`: []FPHSpedVAPIObjectsSpeditionsPartialETSTaskInfo
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UsersUserIdTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersUserIdTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 
 **game** | [**FPHSpedVAPIEnumsGameEnum**](FPHSpedVAPIEnumsGameEnum.md) |   0 &#x3D; ETS2  1 &#x3D; ATS  -1 &#x3D; NotSet | [default to -1]
 **onlyCurrentSped** | **bool** |  | [default to false]

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


## V1VacationsGet

> []FPHSpedVAPIObjectsSpeditionsVacation V1VacationsGet(ctx).Execute()



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
	resp, r, err := apiClient.UserAPI.V1VacationsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1VacationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1VacationsGet`: []FPHSpedVAPIObjectsSpeditionsVacation
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1VacationsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1VacationsGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsSpeditionsVacation**](FPHSpedVAPIObjectsSpeditionsVacation.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

