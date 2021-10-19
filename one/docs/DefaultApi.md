# \DefaultApi

All URIs are relative to *https://apis.one-line.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OauthAccesstokenPost**](DefaultApi.md#V1OauthAccesstokenPost) | **Post** /v1/oauth/accesstoken | 



## V1OauthAccesstokenPost

> InlineResponse200 V1OauthAccesstokenPost(ctx).GrantType(grantType).Apikey(apikey).Authorization(authorization).Execute()





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
    grantType := "client_credentials" // string | client_credentials
    apikey := "oeF32A5QrKe0k8lLQk1J9GJ2F1b5WV4H" // string | client key gotten by registering App
    authorization := "authorization_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.V1OauthAccesstokenPost(context.Background()).GrantType(grantType).Apikey(apikey).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1OauthAccesstokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OauthAccesstokenPost`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1OauthAccesstokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OauthAccesstokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | client_credentials | 
 **apikey** | **string** | client key gotten by registering App | 
 **authorization** | **string** |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

