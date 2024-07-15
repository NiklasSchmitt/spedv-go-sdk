# \ChatAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ChatChatIdGet**](ChatAPI.md#V1ChatChatIdGet) | **Get** /v1/chat/{chatId} | 
[**V1ChatChatIdMessagesGet**](ChatAPI.md#V1ChatChatIdMessagesGet) | **Get** /v1/chat/{chatId}/messages | 
[**V1ChatsGet**](ChatAPI.md#V1ChatsGet) | **Get** /v1/chats | 
[**V1SpedchatGet**](ChatAPI.md#V1SpedchatGet) | **Get** /v1/spedchat | 



## V1ChatChatIdGet

> []FPHSpedVAPIObjectsChatSystemChat V1ChatChatIdGet(ctx, chatId).Execute()



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
	chatId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.V1ChatChatIdGet(context.Background(), chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.V1ChatChatIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ChatChatIdGet`: []FPHSpedVAPIObjectsChatSystemChat
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.V1ChatChatIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ChatChatIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsChatSystemChat**](FPHSpedVAPIObjectsChatSystemChat.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ChatChatIdMessagesGet

> []FPHSpedVAPIObjectsChatSystemChatMessage V1ChatChatIdMessagesGet(ctx, chatId).LastMessage(lastMessage).Execute()



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
	chatId := int32(56) // int32 | 
	lastMessage := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.V1ChatChatIdMessagesGet(context.Background(), chatId).LastMessage(lastMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.V1ChatChatIdMessagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ChatChatIdMessagesGet`: []FPHSpedVAPIObjectsChatSystemChatMessage
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.V1ChatChatIdMessagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ChatChatIdMessagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastMessage** | **time.Time** |  | 

### Return type

[**[]FPHSpedVAPIObjectsChatSystemChatMessage**](FPHSpedVAPIObjectsChatSystemChatMessage.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ChatsGet

> []FPHSpedVAPIObjectsChatSystemChat V1ChatsGet(ctx).Execute()



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
	resp, r, err := apiClient.ChatAPI.V1ChatsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.V1ChatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ChatsGet`: []FPHSpedVAPIObjectsChatSystemChat
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.V1ChatsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ChatsGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsChatSystemChat**](FPHSpedVAPIObjectsChatSystemChat.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SpedchatGet

> FPHSpedVAPIObjectsChatSystemChat V1SpedchatGet(ctx).Execute()



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
	resp, r, err := apiClient.ChatAPI.V1SpedchatGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.V1SpedchatGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SpedchatGet`: FPHSpedVAPIObjectsChatSystemChat
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.V1SpedchatGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SpedchatGetRequest struct via the builder pattern


### Return type

[**FPHSpedVAPIObjectsChatSystemChat**](FPHSpedVAPIObjectsChatSystemChat.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

