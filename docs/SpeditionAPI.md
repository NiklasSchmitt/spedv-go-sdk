# \SpeditionAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SpeditionAccountsGet**](SpeditionAPI.md#V1SpeditionAccountsGet) | **Get** /v1/spedition/accounts | 
[**V1SpeditionBoniTypesGet**](SpeditionAPI.md#V1SpeditionBoniTypesGet) | **Get** /v1/spedition/boniTypes | 
[**V1SpeditionChangelogGet**](SpeditionAPI.md#V1SpeditionChangelogGet) | **Get** /v1/spedition/changelog | 
[**V1SpeditionGet**](SpeditionAPI.md#V1SpeditionGet) | **Get** /v1/spedition | 
[**V1SpeditionPartnershipsGet**](SpeditionAPI.md#V1SpeditionPartnershipsGet) | **Get** /v1/spedition/partnerships | 
[**V1SpeditionRanksGet**](SpeditionAPI.md#V1SpeditionRanksGet) | **Get** /v1/spedition/ranks | 
[**V1SpeditionStatsGet**](SpeditionAPI.md#V1SpeditionStatsGet) | **Get** /v1/spedition/stats | 
[**V1SpeditionStatsSystemGet**](SpeditionAPI.md#V1SpeditionStatsSystemGet) | **Get** /v1/spedition/stats/system | 
[**V1SpeditionStatsUserGet**](SpeditionAPI.md#V1SpeditionStatsUserGet) | **Get** /v1/spedition/stats/user | 
[**V1SpeditionTargetsGet**](SpeditionAPI.md#V1SpeditionTargetsGet) | **Get** /v1/spedition/targets | 



## V1SpeditionAccountsGet

> []FPHSpedVAPIObjectsMoneyBankAccount V1SpeditionAccountsGet(ctx).IncludePartnerships(includePartnerships).Execute()



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
	resp, r, err := apiClient.SpeditionAPI.V1SpeditionAccountsGet(context.Background()).IncludePartnerships(includePartnerships).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpeditionAPI.V1SpeditionAccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SpeditionAccountsGet`: []FPHSpedVAPIObjectsMoneyBankAccount
	fmt.Fprintf(os.Stdout, "Response from `SpeditionAPI.V1SpeditionAccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SpeditionAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includePartnerships** | **bool** |  | [default to false]

### Return type

[**[]FPHSpedVAPIObjectsMoneyBankAccount**](FPHSpedVAPIObjectsMoneyBankAccount.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SpeditionBoniTypesGet

> []FPHSpedVAPIObjectsSpeditionsBoniType V1SpeditionBoniTypesGet(ctx).Execute()



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
	resp, r, err := apiClient.SpeditionAPI.V1SpeditionBoniTypesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpeditionAPI.V1SpeditionBoniTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SpeditionBoniTypesGet`: []FPHSpedVAPIObjectsSpeditionsBoniType
	fmt.Fprintf(os.Stdout, "Response from `SpeditionAPI.V1SpeditionBoniTypesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SpeditionBoniTypesGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsSpeditionsBoniType**](FPHSpedVAPIObjectsSpeditionsBoniType.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SpeditionChangelogGet

> []FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry V1SpeditionChangelogGet(ctx).Execute()



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
	resp, r, err := apiClient.SpeditionAPI.V1SpeditionChangelogGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpeditionAPI.V1SpeditionChangelogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SpeditionChangelogGet`: []FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry
	fmt.Fprintf(os.Stdout, "Response from `SpeditionAPI.V1SpeditionChangelogGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SpeditionChangelogGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry**](FPHSpedVAPIObjectsSpeditionsSpedChangeLogEntry.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SpeditionGet

> FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo V1SpeditionGet(ctx).Execute()



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
	resp, r, err := apiClient.SpeditionAPI.V1SpeditionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpeditionAPI.V1SpeditionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SpeditionGet`: FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo
	fmt.Fprintf(os.Stdout, "Response from `SpeditionAPI.V1SpeditionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SpeditionGetRequest struct via the builder pattern


### Return type

[**FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo**](FPHSpedVAPIObjectsSpeditionsFullSpeditionInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SpeditionPartnershipsGet

> []FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry V1SpeditionPartnershipsGet(ctx).OnlyAccepted(onlyAccepted).Execute()



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
	onlyAccepted := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpeditionAPI.V1SpeditionPartnershipsGet(context.Background()).OnlyAccepted(onlyAccepted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpeditionAPI.V1SpeditionPartnershipsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SpeditionPartnershipsGet`: []FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry
	fmt.Fprintf(os.Stdout, "Response from `SpeditionAPI.V1SpeditionPartnershipsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SpeditionPartnershipsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyAccepted** | **bool** |  | [default to true]

### Return type

[**[]FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry**](FPHSpedVAPIObjectsSpeditionsSpeditionPartnershipEntry.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SpeditionRanksGet

> []FPHSpedVAPIObjectsSpeditionsRank V1SpeditionRanksGet(ctx).Execute()



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
	resp, r, err := apiClient.SpeditionAPI.V1SpeditionRanksGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpeditionAPI.V1SpeditionRanksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SpeditionRanksGet`: []FPHSpedVAPIObjectsSpeditionsRank
	fmt.Fprintf(os.Stdout, "Response from `SpeditionAPI.V1SpeditionRanksGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SpeditionRanksGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsSpeditionsRank**](FPHSpedVAPIObjectsSpeditionsRank.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SpeditionStatsGet

> FPHSpedVAPIServerCommunicationHelperRESTSpedStats V1SpeditionStatsGet(ctx).Execute()



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
	resp, r, err := apiClient.SpeditionAPI.V1SpeditionStatsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpeditionAPI.V1SpeditionStatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SpeditionStatsGet`: FPHSpedVAPIServerCommunicationHelperRESTSpedStats
	fmt.Fprintf(os.Stdout, "Response from `SpeditionAPI.V1SpeditionStatsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SpeditionStatsGetRequest struct via the builder pattern


### Return type

[**FPHSpedVAPIServerCommunicationHelperRESTSpedStats**](FPHSpedVAPIServerCommunicationHelperRESTSpedStats.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SpeditionStatsSystemGet

> FPHSpedVAPIObjectsLiveSystemStatistic V1SpeditionStatsSystemGet(ctx).Execute()



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
	resp, r, err := apiClient.SpeditionAPI.V1SpeditionStatsSystemGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpeditionAPI.V1SpeditionStatsSystemGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SpeditionStatsSystemGet`: FPHSpedVAPIObjectsLiveSystemStatistic
	fmt.Fprintf(os.Stdout, "Response from `SpeditionAPI.V1SpeditionStatsSystemGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SpeditionStatsSystemGetRequest struct via the builder pattern


### Return type

[**FPHSpedVAPIObjectsLiveSystemStatistic**](FPHSpedVAPIObjectsLiveSystemStatistic.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SpeditionStatsUserGet

> []FPHSpedVAPIObjectsUsersUserStatUser V1SpeditionStatsUserGet(ctx).Start(start).End(end).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpeditionAPI.V1SpeditionStatsUserGet(context.Background()).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpeditionAPI.V1SpeditionStatsUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SpeditionStatsUserGet`: []FPHSpedVAPIObjectsUsersUserStatUser
	fmt.Fprintf(os.Stdout, "Response from `SpeditionAPI.V1SpeditionStatsUserGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SpeditionStatsUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 

### Return type

[**[]FPHSpedVAPIObjectsUsersUserStatUser**](FPHSpedVAPIObjectsUsersUserStatUser.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SpeditionTargetsGet

> []FPHSpedVAPIObjectsSpeditionsSpedTarget V1SpeditionTargetsGet(ctx).Execute()



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
	resp, r, err := apiClient.SpeditionAPI.V1SpeditionTargetsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpeditionAPI.V1SpeditionTargetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SpeditionTargetsGet`: []FPHSpedVAPIObjectsSpeditionsSpedTarget
	fmt.Fprintf(os.Stdout, "Response from `SpeditionAPI.V1SpeditionTargetsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SpeditionTargetsGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsSpeditionsSpedTarget**](FPHSpedVAPIObjectsSpeditionsSpedTarget.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

