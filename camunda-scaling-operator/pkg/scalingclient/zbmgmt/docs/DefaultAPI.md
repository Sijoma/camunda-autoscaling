# \DefaultAPI

All URIs are relative to *http://localhost:9600/actuator/cluster*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrokersBrokerIdDelete**](DefaultAPI.md#BrokersBrokerIdDelete) | **Delete** /brokers/{brokerId} | Remove a broker from the cluster.
[**BrokersBrokerIdPost**](DefaultAPI.md#BrokersBrokerIdPost) | **Post** /brokers/{brokerId} | Add a broker to the cluster
[**BrokersPost**](DefaultAPI.md#BrokersPost) | **Post** /brokers | Reconfigure the cluster with the given brokers.
[**RootGet**](DefaultAPI.md#RootGet) | **Get** / | Get current topology



## BrokersBrokerIdDelete

> PlannedOperationsResponse BrokersBrokerIdDelete(ctx, brokerId).DryRun(dryRun).Execute()

Remove a broker from the cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sijoma/camunda-autoscaling-hackday"
)

func main() {
    brokerId := int32(56) // int32 | Id of the broker
    dryRun := true // bool | If true, requested changes are only simulated and not actually applied. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BrokersBrokerIdDelete(context.Background(), brokerId).DryRun(dryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BrokersBrokerIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrokersBrokerIdDelete`: PlannedOperationsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BrokersBrokerIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brokerId** | **int32** | Id of the broker | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrokersBrokerIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **bool** | If true, requested changes are only simulated and not actually applied. | [default to false]

### Return type

[**PlannedOperationsResponse**](PlannedOperationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application.json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrokersBrokerIdPost

> PlannedOperationsResponse BrokersBrokerIdPost(ctx, brokerId).DryRun(dryRun).Execute()

Add a broker to the cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sijoma/camunda-autoscaling-hackday"
)

func main() {
    brokerId := int32(56) // int32 | Id of the broker
    dryRun := true // bool | If true, requested changes are only simulated and not actually applied. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BrokersBrokerIdPost(context.Background(), brokerId).DryRun(dryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BrokersBrokerIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrokersBrokerIdPost`: PlannedOperationsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BrokersBrokerIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brokerId** | **int32** | Id of the broker | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrokersBrokerIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **bool** | If true, requested changes are only simulated and not actually applied. | [default to false]

### Return type

[**PlannedOperationsResponse**](PlannedOperationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application.json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrokersPost

> PlannedOperationsResponse BrokersPost(ctx).RequestBody(requestBody).DryRun(dryRun).Force(force).ReplicationFactor(replicationFactor).Execute()

Reconfigure the cluster with the given brokers.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sijoma/camunda-autoscaling-hackday"
)

func main() {
    requestBody := []int32{int32(0)} // []int32 | 
    dryRun := true // bool | If true, requested changes are only simulated and not actually applied. (optional) (default to false)
    force := true // bool | If true, the operation is a force operation. This is typically used to force remove a set of brokers when they are not available. (optional) (default to false)
    replicationFactor := int32(56) // int32 | The new replication factor for the partitions. If not specified, the current replication factor is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BrokersPost(context.Background()).RequestBody(requestBody).DryRun(dryRun).Force(force).ReplicationFactor(replicationFactor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BrokersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrokersPost`: PlannedOperationsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BrokersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBrokersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]int32** |  | 
 **dryRun** | **bool** | If true, requested changes are only simulated and not actually applied. | [default to false]
 **force** | **bool** | If true, the operation is a force operation. This is typically used to force remove a set of brokers when they are not available. | [default to false]
 **replicationFactor** | **int32** | The new replication factor for the partitions. If not specified, the current replication factor is used. | 

### Return type

[**PlannedOperationsResponse**](PlannedOperationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application.json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RootGet

> GetTopologyResponse RootGet(ctx).Execute()

Get current topology



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sijoma/camunda-autoscaling-hackday"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.RootGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RootGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RootGet`: GetTopologyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RootGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRootGetRequest struct via the builder pattern


### Return type

[**GetTopologyResponse**](GetTopologyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application.json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

