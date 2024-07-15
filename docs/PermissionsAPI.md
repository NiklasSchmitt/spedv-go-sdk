# \PermissionsAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PermissionsGet**](PermissionsAPI.md#V1PermissionsGet) | **Get** /v1/permissions | 



## V1PermissionsGet

> []FPHSpedVAPIObjectsSpeditionsPermission V1PermissionsGet(ctx).Execute()



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
	resp, r, err := apiClient.PermissionsAPI.V1PermissionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.V1PermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PermissionsGet`: []FPHSpedVAPIObjectsSpeditionsPermission
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.V1PermissionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1PermissionsGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsSpeditionsPermission**](FPHSpedVAPIObjectsSpeditionsPermission.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

