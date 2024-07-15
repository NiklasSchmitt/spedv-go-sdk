# \OmsiAPI

All URIs are relative to *https://api.sped-v.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OmsiDriventripsGet**](OmsiAPI.md#V1OmsiDriventripsGet) | **Get** /v1/omsi/driventrips | 
[**V1OmsiMapsGet**](OmsiAPI.md#V1OmsiMapsGet) | **Get** /v1/omsi/maps | 
[**V1OmsiMapsMapidBusstopsGet**](OmsiAPI.md#V1OmsiMapsMapidBusstopsGet) | **Get** /v1/omsi/maps/{mapid}/busstops | 
[**V1OmsiMapsMapidDestinationsGet**](OmsiAPI.md#V1OmsiMapsMapidDestinationsGet) | **Get** /v1/omsi/maps/{mapid}/destinations | 
[**V1OmsiMapsMapidGet**](OmsiAPI.md#V1OmsiMapsMapidGet) | **Get** /v1/omsi/maps/{mapid} | 
[**V1OmsiMapsMapidStationlinksGet**](OmsiAPI.md#V1OmsiMapsMapidStationlinksGet) | **Get** /v1/omsi/maps/{mapid}/stationlinks | 
[**V1OmsiMapsMapidToursGet**](OmsiAPI.md#V1OmsiMapsMapidToursGet) | **Get** /v1/omsi/maps/{mapid}/tours | 
[**V1OmsiMapsMapidToursTouridGet**](OmsiAPI.md#V1OmsiMapsMapidToursTouridGet) | **Get** /v1/omsi/maps/{mapid}/tours/{tourid} | 
[**V1OmsiMapsMapidToursTouridTripsGet**](OmsiAPI.md#V1OmsiMapsMapidToursTouridTripsGet) | **Get** /v1/omsi/maps/{mapid}/tours/{tourid}/trips | 
[**V1OmsiMapsMapidTourtripsDateGet**](OmsiAPI.md#V1OmsiMapsMapidTourtripsDateGet) | **Get** /v1/omsi/maps/{mapid}/tourtrips/{date} | 
[**V1OmsiMapsMapidTripsGet**](OmsiAPI.md#V1OmsiMapsMapidTripsGet) | **Get** /v1/omsi/maps/{mapid}/trips | 



## V1OmsiDriventripsGet

> []FPHSpedVAPIObjectsOMSIDrivenTrip V1OmsiDriventripsGet(ctx).Start(start).End(end).Execute()



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
	resp, r, err := apiClient.OmsiAPI.V1OmsiDriventripsGet(context.Background()).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OmsiAPI.V1OmsiDriventripsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OmsiDriventripsGet`: []FPHSpedVAPIObjectsOMSIDrivenTrip
	fmt.Fprintf(os.Stdout, "Response from `OmsiAPI.V1OmsiDriventripsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OmsiDriventripsGetRequest struct via the builder pattern


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


## V1OmsiMapsGet

> []FPHSpedVAPIObjectsOMSIMap V1OmsiMapsGet(ctx).Execute()



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
	resp, r, err := apiClient.OmsiAPI.V1OmsiMapsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OmsiAPI.V1OmsiMapsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OmsiMapsGet`: []FPHSpedVAPIObjectsOMSIMap
	fmt.Fprintf(os.Stdout, "Response from `OmsiAPI.V1OmsiMapsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1OmsiMapsGetRequest struct via the builder pattern


### Return type

[**[]FPHSpedVAPIObjectsOMSIMap**](FPHSpedVAPIObjectsOMSIMap.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OmsiMapsMapidBusstopsGet

> []FPHSpedVAPIObjectsOMSIBusStop V1OmsiMapsMapidBusstopsGet(ctx, mapid).Execute()



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
	mapid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OmsiAPI.V1OmsiMapsMapidBusstopsGet(context.Background(), mapid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OmsiAPI.V1OmsiMapsMapidBusstopsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OmsiMapsMapidBusstopsGet`: []FPHSpedVAPIObjectsOMSIBusStop
	fmt.Fprintf(os.Stdout, "Response from `OmsiAPI.V1OmsiMapsMapidBusstopsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OmsiMapsMapidBusstopsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsOMSIBusStop**](FPHSpedVAPIObjectsOMSIBusStop.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OmsiMapsMapidDestinationsGet

> []FPHSpedVAPIObjectsOMSIDestination V1OmsiMapsMapidDestinationsGet(ctx, mapid).Execute()



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
	mapid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OmsiAPI.V1OmsiMapsMapidDestinationsGet(context.Background(), mapid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OmsiAPI.V1OmsiMapsMapidDestinationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OmsiMapsMapidDestinationsGet`: []FPHSpedVAPIObjectsOMSIDestination
	fmt.Fprintf(os.Stdout, "Response from `OmsiAPI.V1OmsiMapsMapidDestinationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OmsiMapsMapidDestinationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsOMSIDestination**](FPHSpedVAPIObjectsOMSIDestination.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OmsiMapsMapidGet

> FPHSpedVAPIObjectsOMSIMap V1OmsiMapsMapidGet(ctx, mapid).Execute()



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
	mapid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OmsiAPI.V1OmsiMapsMapidGet(context.Background(), mapid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OmsiAPI.V1OmsiMapsMapidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OmsiMapsMapidGet`: FPHSpedVAPIObjectsOMSIMap
	fmt.Fprintf(os.Stdout, "Response from `OmsiAPI.V1OmsiMapsMapidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OmsiMapsMapidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FPHSpedVAPIObjectsOMSIMap**](FPHSpedVAPIObjectsOMSIMap.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OmsiMapsMapidStationlinksGet

> []FPHSpedVAPIObjectsOMSIStationLink V1OmsiMapsMapidStationlinksGet(ctx, mapid).Execute()



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
	mapid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OmsiAPI.V1OmsiMapsMapidStationlinksGet(context.Background(), mapid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OmsiAPI.V1OmsiMapsMapidStationlinksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OmsiMapsMapidStationlinksGet`: []FPHSpedVAPIObjectsOMSIStationLink
	fmt.Fprintf(os.Stdout, "Response from `OmsiAPI.V1OmsiMapsMapidStationlinksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OmsiMapsMapidStationlinksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsOMSIStationLink**](FPHSpedVAPIObjectsOMSIStationLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OmsiMapsMapidToursGet

> []FPHSpedVAPIObjectsOMSITour V1OmsiMapsMapidToursGet(ctx, mapid).Execute()



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
	mapid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OmsiAPI.V1OmsiMapsMapidToursGet(context.Background(), mapid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OmsiAPI.V1OmsiMapsMapidToursGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OmsiMapsMapidToursGet`: []FPHSpedVAPIObjectsOMSITour
	fmt.Fprintf(os.Stdout, "Response from `OmsiAPI.V1OmsiMapsMapidToursGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OmsiMapsMapidToursGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsOMSITour**](FPHSpedVAPIObjectsOMSITour.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OmsiMapsMapidToursTouridGet

> FPHSpedVAPIObjectsOMSITour V1OmsiMapsMapidToursTouridGet(ctx, mapid, tourid).Execute()



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
	mapid := int32(56) // int32 | 
	tourid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OmsiAPI.V1OmsiMapsMapidToursTouridGet(context.Background(), mapid, tourid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OmsiAPI.V1OmsiMapsMapidToursTouridGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OmsiMapsMapidToursTouridGet`: FPHSpedVAPIObjectsOMSITour
	fmt.Fprintf(os.Stdout, "Response from `OmsiAPI.V1OmsiMapsMapidToursTouridGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapid** | **int32** |  | 
**tourid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OmsiMapsMapidToursTouridGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FPHSpedVAPIObjectsOMSITour**](FPHSpedVAPIObjectsOMSITour.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OmsiMapsMapidToursTouridTripsGet

> []FPHSpedVAPIObjectsOMSITourTrip V1OmsiMapsMapidToursTouridTripsGet(ctx, mapid, tourid).BaseDate(baseDate).Execute()



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
	mapid := int32(56) // int32 | 
	tourid := int32(56) // int32 | 
	baseDate := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OmsiAPI.V1OmsiMapsMapidToursTouridTripsGet(context.Background(), mapid, tourid).BaseDate(baseDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OmsiAPI.V1OmsiMapsMapidToursTouridTripsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OmsiMapsMapidToursTouridTripsGet`: []FPHSpedVAPIObjectsOMSITourTrip
	fmt.Fprintf(os.Stdout, "Response from `OmsiAPI.V1OmsiMapsMapidToursTouridTripsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapid** | **int32** |  | 
**tourid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OmsiMapsMapidToursTouridTripsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **baseDate** | **time.Time** |  | 

### Return type

[**[]FPHSpedVAPIObjectsOMSITourTrip**](FPHSpedVAPIObjectsOMSITourTrip.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OmsiMapsMapidTourtripsDateGet

> []FPHSpedVAPIObjectsOMSITour V1OmsiMapsMapidTourtripsDateGet(ctx, mapid, date).Timediff(timediff).Execute()



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
	mapid := int32(56) // int32 | 
	date := "date_example" // string | 
	timediff := int32(56) // int32 |  (optional) (default to 2)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OmsiAPI.V1OmsiMapsMapidTourtripsDateGet(context.Background(), mapid, date).Timediff(timediff).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OmsiAPI.V1OmsiMapsMapidTourtripsDateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OmsiMapsMapidTourtripsDateGet`: []FPHSpedVAPIObjectsOMSITour
	fmt.Fprintf(os.Stdout, "Response from `OmsiAPI.V1OmsiMapsMapidTourtripsDateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapid** | **int32** |  | 
**date** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OmsiMapsMapidTourtripsDateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timediff** | **int32** |  | [default to 2]

### Return type

[**[]FPHSpedVAPIObjectsOMSITour**](FPHSpedVAPIObjectsOMSITour.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OmsiMapsMapidTripsGet

> []FPHSpedVAPIObjectsOMSITourTrip V1OmsiMapsMapidTripsGet(ctx, mapid).Execute()



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
	mapid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OmsiAPI.V1OmsiMapsMapidTripsGet(context.Background(), mapid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OmsiAPI.V1OmsiMapsMapidTripsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OmsiMapsMapidTripsGet`: []FPHSpedVAPIObjectsOMSITourTrip
	fmt.Fprintf(os.Stdout, "Response from `OmsiAPI.V1OmsiMapsMapidTripsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OmsiMapsMapidTripsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FPHSpedVAPIObjectsOMSITourTrip**](FPHSpedVAPIObjectsOMSITourTrip.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

