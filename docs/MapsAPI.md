# \MapsAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CompaniesCompidBranchesGet**](MapsAPI.md#V1CompaniesCompidBranchesGet) | **Get** /v1/companies/{compid}/branches | 
[**V1CompaniesCompidGet**](MapsAPI.md#V1CompaniesCompidGet) | **Get** /v1/companies/{compid} | 
[**V1CompaniesGet**](MapsAPI.md#V1CompaniesGet) | **Get** /v1/companies | 
[**V1CompanycategoriesCompidGet**](MapsAPI.md#V1CompanycategoriesCompidGet) | **Get** /v1/companycategories/{compid} | 
[**V1CompanycategoriesGet**](MapsAPI.md#V1CompanycategoriesGet) | **Get** /v1/companycategories | 
[**V1CountriesGet**](MapsAPI.md#V1CountriesGet) | **Get** /v1/countries | 
[**V1MapsGet**](MapsAPI.md#V1MapsGet) | **Get** /v1/maps | 
[**V1MapsMapIdCitiesGet**](MapsAPI.md#V1MapsMapIdCitiesGet) | **Get** /v1/maps/{mapId}/cities | 
[**V1MapsMapIdCompanycitiesGet**](MapsAPI.md#V1MapsMapIdCompanycitiesGet) | **Get** /v1/maps/{mapId}/companycities | 



## V1CompaniesCompidBranchesGet

> []FPHSpedVAPIObjectsMapsCompanyCity V1CompaniesCompidBranchesGet(ctx, compid).Execute()



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
	compid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MapsAPI.V1CompaniesCompidBranchesGet(context.Background(), compid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MapsAPI.V1CompaniesCompidBranchesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CompaniesCompidBranchesGet`: []FPHSpedVAPIObjectsMapsCompanyCity
	fmt.Fprintf(os.Stdout, "Response from `MapsAPI.V1CompaniesCompidBranchesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**compid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CompaniesCompidBranchesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsMapsCompanyCity**](FPHSpedVAPIObjectsMapsCompanyCity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CompaniesCompidGet

> FPHSpedVAPIObjectsMapsCompany V1CompaniesCompidGet(ctx, compid).Execute()



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
	compid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MapsAPI.V1CompaniesCompidGet(context.Background(), compid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MapsAPI.V1CompaniesCompidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CompaniesCompidGet`: FPHSpedVAPIObjectsMapsCompany
	fmt.Fprintf(os.Stdout, "Response from `MapsAPI.V1CompaniesCompidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**compid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CompaniesCompidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FPHSpedVAPIObjectsMapsCompany**](FPHSpedVAPIObjectsMapsCompany.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CompaniesGet

> []FPHSpedVAPIObjectsMapsCompany V1CompaniesGet(ctx).Execute()



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
	resp, r, err := apiClient.MapsAPI.V1CompaniesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MapsAPI.V1CompaniesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CompaniesGet`: []FPHSpedVAPIObjectsMapsCompany
	fmt.Fprintf(os.Stdout, "Response from `MapsAPI.V1CompaniesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CompaniesGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsMapsCompany**](FPHSpedVAPIObjectsMapsCompany.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CompanycategoriesCompidGet

> FPHSpedVAPIObjectsMapsCompanyCategory V1CompanycategoriesCompidGet(ctx, compid).Execute()



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
	compid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MapsAPI.V1CompanycategoriesCompidGet(context.Background(), compid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MapsAPI.V1CompanycategoriesCompidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CompanycategoriesCompidGet`: FPHSpedVAPIObjectsMapsCompanyCategory
	fmt.Fprintf(os.Stdout, "Response from `MapsAPI.V1CompanycategoriesCompidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**compid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CompanycategoriesCompidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FPHSpedVAPIObjectsMapsCompanyCategory**](FPHSpedVAPIObjectsMapsCompanyCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CompanycategoriesGet

> []FPHSpedVAPIObjectsMapsCompanyCategory V1CompanycategoriesGet(ctx).Execute()



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
	resp, r, err := apiClient.MapsAPI.V1CompanycategoriesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MapsAPI.V1CompanycategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CompanycategoriesGet`: []FPHSpedVAPIObjectsMapsCompanyCategory
	fmt.Fprintf(os.Stdout, "Response from `MapsAPI.V1CompanycategoriesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CompanycategoriesGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsMapsCompanyCategory**](FPHSpedVAPIObjectsMapsCompanyCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CountriesGet

> []FPHSpedVAPIObjectsMapsCountry V1CountriesGet(ctx).Execute()



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
	resp, r, err := apiClient.MapsAPI.V1CountriesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MapsAPI.V1CountriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CountriesGet`: []FPHSpedVAPIObjectsMapsCountry
	fmt.Fprintf(os.Stdout, "Response from `MapsAPI.V1CountriesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CountriesGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsMapsCountry**](FPHSpedVAPIObjectsMapsCountry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MapsGet

> []FPHSpedVAPIObjectsMapsMap V1MapsGet(ctx).Execute()



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
	resp, r, err := apiClient.MapsAPI.V1MapsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MapsAPI.V1MapsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MapsGet`: []FPHSpedVAPIObjectsMapsMap
	fmt.Fprintf(os.Stdout, "Response from `MapsAPI.V1MapsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1MapsGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsMapsMap**](FPHSpedVAPIObjectsMapsMap.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MapsMapIdCitiesGet

> []FPHSpedVAPIObjectsMapsCity V1MapsMapIdCitiesGet(ctx, mapId).Execute()



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
	mapId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MapsAPI.V1MapsMapIdCitiesGet(context.Background(), mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MapsAPI.V1MapsMapIdCitiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MapsMapIdCitiesGet`: []FPHSpedVAPIObjectsMapsCity
	fmt.Fprintf(os.Stdout, "Response from `MapsAPI.V1MapsMapIdCitiesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MapsMapIdCitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsMapsCity**](FPHSpedVAPIObjectsMapsCity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MapsMapIdCompanycitiesGet

> []FPHSpedVAPIObjectsMapsCompanyCity V1MapsMapIdCompanycitiesGet(ctx, mapId).Execute()



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
	mapId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MapsAPI.V1MapsMapIdCompanycitiesGet(context.Background(), mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MapsAPI.V1MapsMapIdCompanycitiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MapsMapIdCompanycitiesGet`: []FPHSpedVAPIObjectsMapsCompanyCity
	fmt.Fprintf(os.Stdout, "Response from `MapsAPI.V1MapsMapIdCompanycitiesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MapsMapIdCompanycitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsMapsCompanyCity**](FPHSpedVAPIObjectsMapsCompanyCity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

