# MaintenanceWindows

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMaintenanceWindow**](MaintenanceWindowsApi.md#DeleteMaintenanceWindow) | **Delete** /api/v1/clusters/{cluster_id}/maintenance-window | 
[**GetMaintenanceWindow**](MaintenanceWindowsApi.md#GetMaintenanceWindow) | **Get** /api/v1/clusters/{cluster_id}/maintenance-window | 
[**SetMaintenanceWindow**](MaintenanceWindowsApi.md#SetMaintenanceWindow) | **Put** /api/v1/clusters/{cluster_id}/maintenance-window | 



## DeleteMaintenanceWindow

> MaintenanceWindow DeleteMaintenanceWindow(ctx, clusterId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.DeleteMaintenanceWindow(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.DeleteMaintenanceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMaintenanceWindow`: MaintenanceWindow
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.DeleteMaintenanceWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMaintenanceWindow struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetMaintenanceWindow

> MaintenanceWindow GetMaintenanceWindow(ctx, clusterId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.GetMaintenanceWindow(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.GetMaintenanceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaintenanceWindow`: MaintenanceWindow
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.GetMaintenanceWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceWindow struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## SetMaintenanceWindow

> MaintenanceWindow SetMaintenanceWindow(ctx, clusterId).MaintenanceWindow(maintenanceWindow).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | 
    maintenanceWindow := *openapiclient.NewMaintenanceWindow("OffsetDuration_example", "WindowDuration_example") // MaintenanceWindow | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.SetMaintenanceWindow(context.Background(), clusterId).MaintenanceWindow(maintenanceWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.SetMaintenanceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetMaintenanceWindow`: MaintenanceWindow
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.SetMaintenanceWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaintenanceWindow struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenanceWindow** | [**MaintenanceWindow**](MaintenanceWindow.md) |  | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

