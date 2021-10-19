# \EventsApi

All URIs are relative to *https://apis.one-line.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CargotrackingEventsGet**](EventsApi.md#V1CargotrackingEventsGet) | **Get** /v1/cargotracking/events | Find events by type, Booking Reference, Bill of Lading or Equipment Reference.



## V1CargotrackingEventsGet

> SampleEvent V1CargotrackingEventsGet(ctx).Apikey(apikey).Authorization(authorization).EventType(eventType).BookingReference(bookingReference).BillOfLadingNumber(billOfLadingNumber).EquipmentReference(equipmentReference).Execute()

Find events by type, Booking Reference, Bill of Lading or Equipment Reference.



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
    apikey := "oeF32A5QrKe0k8lLQk1J9GJ2F1b5WV4H" // string | valid apikey
    authorization := "Bearer dTBdp1iCX2dw40xPneuWA2FSIekK" // string | valid access_token
    eventType := []string{"EventType_example"} // []string | The type of event(s) to filter by. (optional) (default to ["EQUIPMENT","SHIPMENT","TRANSPORT","TRANSPORTEQUIPMENT"])
    bookingReference := "bookingReference_example" // string | The identifier for a shipment, which is issued by and unique within each of the carriers. (optional)
    billOfLadingNumber := "billOfLadingNumber_example" // string | Bill of lading number is an identifier that links to a shipment. Bill of Lading is the legal document issued to the customer, which confirms the carrier's receipt of the cargo from the customer acknowledging goods being shipped and specifying the terms of delivery. (optional)
    equipmentReference := "equipmentReference_example" // string | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventsApi.V1CargotrackingEventsGet(context.Background()).Apikey(apikey).Authorization(authorization).EventType(eventType).BookingReference(bookingReference).BillOfLadingNumber(billOfLadingNumber).EquipmentReference(equipmentReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.V1CargotrackingEventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CargotrackingEventsGet`: SampleEvent
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.V1CargotrackingEventsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CargotrackingEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apikey** | **string** | valid apikey | 
 **authorization** | **string** | valid access_token | 
 **eventType** | **[]string** | The type of event(s) to filter by. | [default to [&quot;EQUIPMENT&quot;,&quot;SHIPMENT&quot;,&quot;TRANSPORT&quot;,&quot;TRANSPORTEQUIPMENT&quot;]]
 **bookingReference** | **string** | The identifier for a shipment, which is issued by and unique within each of the carriers. | 
 **billOfLadingNumber** | **string** | Bill of lading number is an identifier that links to a shipment. Bill of Lading is the legal document issued to the customer, which confirms the carrier&#39;s receipt of the cargo from the customer acknowledging goods being shipped and specifying the terms of delivery. | 
 **equipmentReference** | **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. | 

### Return type

[**SampleEvent**](SampleEvent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

