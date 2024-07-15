# \BankAccountAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1BankaccountsAccountIdGet**](BankAccountAPI.md#V1BankaccountsAccountIdGet) | **Get** /v1/bankaccounts/{accountId} | 
[**V1BankaccountsAccountIdTransfersDailyGet**](BankAccountAPI.md#V1BankaccountsAccountIdTransfersDailyGet) | **Get** /v1/bankaccounts/{accountId}/transfers/daily | 
[**V1BankaccountsAccountIdTransfersGet**](BankAccountAPI.md#V1BankaccountsAccountIdTransfersGet) | **Get** /v1/bankaccounts/{accountId}/transfers | 
[**V1BankaccountsAccountIdTransfersSummaryGet**](BankAccountAPI.md#V1BankaccountsAccountIdTransfersSummaryGet) | **Get** /v1/bankaccounts/{accountId}/transfers/summary | 
[**V1BankaccountsGet**](BankAccountAPI.md#V1BankaccountsGet) | **Get** /v1/bankaccounts | 
[**V1BankaccountsIbanGet**](BankAccountAPI.md#V1BankaccountsIbanGet) | **Get** /v1/bankaccounts/{iban} | 



## V1BankaccountsAccountIdGet

> []FPHSpedVAPIObjectsMoneyLiteBankAccount V1BankaccountsAccountIdGet(ctx, accountId).Execute()



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
	accountId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankAccountAPI.V1BankaccountsAccountIdGet(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankAccountAPI.V1BankaccountsAccountIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1BankaccountsAccountIdGet`: []FPHSpedVAPIObjectsMoneyLiteBankAccount
	fmt.Fprintf(os.Stdout, "Response from `BankAccountAPI.V1BankaccountsAccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BankaccountsAccountIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsMoneyLiteBankAccount**](FPHSpedVAPIObjectsMoneyLiteBankAccount.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BankaccountsAccountIdTransfersDailyGet

> []FPHSpedVAPIObjectsMoneyMoneyTransferDay V1BankaccountsAccountIdTransfersDailyGet(ctx, accountId).StartDate(startDate).EndDate(endDate).Execute()



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
	accountId := int32(56) // int32 | 
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankAccountAPI.V1BankaccountsAccountIdTransfersDailyGet(context.Background(), accountId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankAccountAPI.V1BankaccountsAccountIdTransfersDailyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1BankaccountsAccountIdTransfersDailyGet`: []FPHSpedVAPIObjectsMoneyMoneyTransferDay
	fmt.Fprintf(os.Stdout, "Response from `BankAccountAPI.V1BankaccountsAccountIdTransfersDailyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BankaccountsAccountIdTransfersDailyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 

### Return type

[**[]FPHSpedVAPIObjectsMoneyMoneyTransferDay**](FPHSpedVAPIObjectsMoneyMoneyTransferDay.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BankaccountsAccountIdTransfersGet

> []FPHSpedVAPIObjectsMoneyMoneyTransfer V1BankaccountsAccountIdTransfersGet(ctx, accountId).StartDate(startDate).Execute()



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
	accountId := int32(56) // int32 | 
	startDate := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankAccountAPI.V1BankaccountsAccountIdTransfersGet(context.Background(), accountId).StartDate(startDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankAccountAPI.V1BankaccountsAccountIdTransfersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1BankaccountsAccountIdTransfersGet`: []FPHSpedVAPIObjectsMoneyMoneyTransfer
	fmt.Fprintf(os.Stdout, "Response from `BankAccountAPI.V1BankaccountsAccountIdTransfersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BankaccountsAccountIdTransfersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** |  | 

### Return type

[**[]FPHSpedVAPIObjectsMoneyMoneyTransfer**](FPHSpedVAPIObjectsMoneyMoneyTransfer.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BankaccountsAccountIdTransfersSummaryGet

> []FPHSpedVAPIObjectsMoneyMoneyEURData V1BankaccountsAccountIdTransfersSummaryGet(ctx, accountId).StartDate(startDate).EndDate(endDate).Execute()



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
	accountId := int32(56) // int32 | 
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankAccountAPI.V1BankaccountsAccountIdTransfersSummaryGet(context.Background(), accountId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankAccountAPI.V1BankaccountsAccountIdTransfersSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1BankaccountsAccountIdTransfersSummaryGet`: []FPHSpedVAPIObjectsMoneyMoneyEURData
	fmt.Fprintf(os.Stdout, "Response from `BankAccountAPI.V1BankaccountsAccountIdTransfersSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BankaccountsAccountIdTransfersSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 

### Return type

[**[]FPHSpedVAPIObjectsMoneyMoneyEURData**](FPHSpedVAPIObjectsMoneyMoneyEURData.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BankaccountsGet

> []FPHSpedVAPIObjectsMoneyBankAccount V1BankaccountsGet(ctx).Execute()



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
	resp, r, err := apiClient.BankAccountAPI.V1BankaccountsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankAccountAPI.V1BankaccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1BankaccountsGet`: []FPHSpedVAPIObjectsMoneyBankAccount
	fmt.Fprintf(os.Stdout, "Response from `BankAccountAPI.V1BankaccountsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1BankaccountsGetRequest struct via the builder pattern


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


## V1BankaccountsIbanGet

> []FPHSpedVAPIObjectsMoneyLiteBankAccount V1BankaccountsIbanGet(ctx, iban).Execute()



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
	iban := "iban_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankAccountAPI.V1BankaccountsIbanGet(context.Background(), iban).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankAccountAPI.V1BankaccountsIbanGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1BankaccountsIbanGet`: []FPHSpedVAPIObjectsMoneyLiteBankAccount
	fmt.Fprintf(os.Stdout, "Response from `BankAccountAPI.V1BankaccountsIbanGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iban** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BankaccountsIbanGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsMoneyLiteBankAccount**](FPHSpedVAPIObjectsMoneyLiteBankAccount.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

