# \EventsAPI

All URIs are relative to *https://apis.one-line.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvents**](EventsAPI.md#GetEvents) | **Get** /v2/cargotracking/events | Find events by type, Booking Reference, Bill of Lading or Equipment Reference.



## GetEvents

> []Event GetEvents(ctx).Apikey(apikey).Authorization(authorization).EventType(eventType).CarrierBookingReference(carrierBookingReference).TransportDocumentReference(transportDocumentReference).EquipmentReference(equipmentReference).Limit(limit).Sort(sort).Execute()

Find events by type, Booking Reference, Bill of Lading or Equipment Reference.



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
	apikey := "H2mX7pLwQbZ9rV8TjK5oY1fNsD3u4cA" // string | Valid API key.
	authorization := "Bearer m9ZrH7yBvq3fX8UoLJ2wP5cT1nE0gKs" // string | Valid access token.
	eventType := []string{"EventType_example"} // []string | The type of event(s) to filter by.   Possible values are    - TRANSPORT (Transport events)   - EQUIPMENT (Equipment events)   (optional)
	carrierBookingReference := "carrierBookingReference_example" // string | A set of unique characters provided by carrier to identify a booking. Specifying this filter will only return events related to this particular carrierBookingReference.  (optional)
	transportDocumentReference := "transportDocumentReference_example" // string | A unique number reference allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment. Specifying this filter will only return events related to this particular transportDocumentReference.  (optional)
	equipmentReference := "equipmentReference_example" // string | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. (optional)
	limit := int32(100) // int32 | Maximum number of items to return. (optional) (default to 100)
	sort := "payload.eventDateTime:desc" // string | A comma-separated (,) list of field names is used to define the sort order. To specify the sort order for each field name required, use a colon (:) between the field name and the keyword asc (ascending) or desc (descending). If the keyword is not used, ascending (asc) order is applied by default.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetEvents(context.Background()).Apikey(apikey).Authorization(authorization).EventType(eventType).CarrierBookingReference(carrierBookingReference).TransportDocumentReference(transportDocumentReference).EquipmentReference(equipmentReference).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvents`: []Event
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apikey** | **string** | Valid API key. | 
 **authorization** | **string** | Valid access token. | 
 **eventType** | **[]string** | The type of event(s) to filter by.   Possible values are    - TRANSPORT (Transport events)   - EQUIPMENT (Equipment events)   | 
 **carrierBookingReference** | **string** | A set of unique characters provided by carrier to identify a booking. Specifying this filter will only return events related to this particular carrierBookingReference.  | 
 **transportDocumentReference** | **string** | A unique number reference allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment. Specifying this filter will only return events related to this particular transportDocumentReference.  | 
 **equipmentReference** | **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. | 
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **sort** | **string** | A comma-separated (,) list of field names is used to define the sort order. To specify the sort order for each field name required, use a colon (:) between the field name and the keyword asc (ascending) or desc (descending). If the keyword is not used, ascending (asc) order is applied by default.  | 

### Return type

[**[]Event**](Event.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

