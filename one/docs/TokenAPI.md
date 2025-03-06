# \TokenAPI

All URIs are relative to *https://apis.one-line.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authenticate**](TokenAPI.md#Authenticate) | **Post** /v1/oauth/accesstoken | 



## Authenticate

> InlineResponse200 Authenticate(ctx).GrantType(grantType).Apikey(apikey).Authorization(authorization).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/buyco/go-one-sdk/one"
)

func main() {
	grantType := "client_credentials" // string | client_credentials
	apikey := "4gH7zKl8WpQeR1v9TnM2XoF6bJsD5iYcU0ArL3" // string | Client key obtained by registering the app.
	authorization := "authorization_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.Authenticate(context.Background()).GrantType(grantType).Apikey(apikey).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.Authenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Authenticate`: InlineResponse200
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.Authenticate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | client_credentials | 
 **apikey** | **string** | Client key obtained by registering the app. | 
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

