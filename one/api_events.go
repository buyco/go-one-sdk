/*
ONE Specification for DCSA Track & Trace API v2.2

  # DCSA Track and Trace v2.2    **Overview:**    The **DCSA Track and Trace v2.2 API** provides a standardized and reliable mechanism for tracking container movements throughout the global shipping network. It is designed to align with the Digital Container Shipping Association (DCSA) standards, promoting **interoperability** between carriers, shippers, and third-party logistics providers.    This API allows stakeholders to monitor the **end-to-end journey** of containers, offering insights into key milestones such as:  - **Vessel arrivals**  - **Departures**  - **Equipment loading**  - **Transport events**    By delivering **uniform tracking data** across various carriers and ports, this API reduces manual processes and improves operational efficiency for managing large-scale shipping operations.    ---    ## Key Features:    ### 1. **Event Standardization**    All events related to container transport are defined according to DCSA standards, ensuring consistent terminology and reporting across the entire logistics chain.    ### 2. **On-Demand Data Retrieval**    Customers can pull **up-to-date information** about container locations, equipment status, and shipment events when needed, enabling timely and informed decision-making.    ### 3. **Seamless Integration**    Designed to integrate easily with existing systems, the API offers flexibility for users who manage large-scale shipments or have evolving logistics requirements.    ### 4. **Scalability**    Whether tracking a few containers or managing thousands, the API is optimized to handle **large datasets efficiently**, ensuring timely access to critical shipping information.    ---    ### Conclusion:    The **DCSA Track and Trace v2.2 API** empowers companies to enhance their visibility into global logistics operations, improve communication between parties, and drive better service outcomes through transparency and **real-time information sharing**.

API version: 2.2
Contact: integration.support@one-line.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package one

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

type EventsAPI interface {

	/*
		GetEvents Find events by type, Booking Reference, Bill of Lading or Equipment Reference.

		# Best Practice Guidelines for API Client Configuration

	**To ensure the efficient and accurate retrieval of container tracking information via the ONE DCSA Track and Trace v2.2 API, we recommend the following configuration settings for your API client.**

	### 1. Number of Containers:
	It is anticipated that a significant number of containers, potentially reaching several thousand, may be in active transit within your system, requiring continuous tracking and status updates.

	### 2. API Call Rate: **60 API Calls per Minute**
	- Our recommendation is that your API polling client should be configured to make **60 API calls per minute**.
	- This rate limit is established to maintain the responsiveness and reliability of the ONE API service. Exceeding this rate may lead to throttling or delays in response times.

	### 3. Batch Job Execution:
	- API consumers who are batching their requests, we recommend scheduling your API client to run **batch jobs up to three times daily** to retrieve tracking information for containers in your system.
	- **Batch Job API Calls**: Depending on the volume of containers, each batch job may generate several thousand API calls to the ONE server.

	### 4. Use of `equipmentReference`:
	- Although `carrierBookingReference` and `transportDocumentReference` are available, we highly recommend using the **`equipmentReference`** (container number) in your API calls whenever possible.
	- Using the container number is more efficient in terms of latency, improving response times and reducing processing load.

	### 5. Error Handling and Retry Mechanisms:
	- Implement robust error handling and retry mechanisms within your API client to manage potential disruptions in connectivity or response errors.
	- This will help minimize data retrieval failures and ensure continuous tracking of container movements.

	### 6. Data Synchronization:
	- To avoid data discrepancies, schedule your API calls to align with the frequency of container status updates provided by the ONE system (Transport Plan).
	- This synchronization will help ensure that the tracking data in your system is up-to-date and accurate.



		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetEventsRequest
	*/
	GetEvents(ctx context.Context) ApiGetEventsRequest

	// GetEventsExecute executes the request
	//  @return []Event
	GetEventsExecute(r ApiGetEventsRequest) ([]Event, *http.Response, error)
}

// EventsAPIService EventsAPI service
type EventsAPIService service

type ApiGetEventsRequest struct {
	ctx                        context.Context
	ApiService                 EventsAPI
	apikey                     *string
	authorization              *string
	eventType                  *[]string
	carrierBookingReference    *string
	transportDocumentReference *string
	equipmentReference         *string
	limit                      *int32
	sort                       *string
}

// Valid API key.
func (r ApiGetEventsRequest) Apikey(apikey string) ApiGetEventsRequest {
	r.apikey = &apikey
	return r
}

// Valid access token.
func (r ApiGetEventsRequest) Authorization(authorization string) ApiGetEventsRequest {
	r.authorization = &authorization
	return r
}

// The type of event(s) to filter by.   Possible values are    - TRANSPORT (Transport events)   - EQUIPMENT (Equipment events)
func (r ApiGetEventsRequest) EventType(eventType []string) ApiGetEventsRequest {
	r.eventType = &eventType
	return r
}

// A set of unique characters provided by carrier to identify a booking. Specifying this filter will only return events related to this particular carrierBookingReference.
func (r ApiGetEventsRequest) CarrierBookingReference(carrierBookingReference string) ApiGetEventsRequest {
	r.carrierBookingReference = &carrierBookingReference
	return r
}

// A unique number reference allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment. Specifying this filter will only return events related to this particular transportDocumentReference.
func (r ApiGetEventsRequest) TransportDocumentReference(transportDocumentReference string) ApiGetEventsRequest {
	r.transportDocumentReference = &transportDocumentReference
	return r
}

// The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible.
func (r ApiGetEventsRequest) EquipmentReference(equipmentReference string) ApiGetEventsRequest {
	r.equipmentReference = &equipmentReference
	return r
}

// Maximum number of items to return.
func (r ApiGetEventsRequest) Limit(limit int32) ApiGetEventsRequest {
	r.limit = &limit
	return r
}

// A comma-separated (,) list of field names is used to define the sort order. To specify the sort order for each field name required, use a colon (:) between the field name and the keyword asc (ascending) or desc (descending). If the keyword is not used, ascending (asc) order is applied by default.
func (r ApiGetEventsRequest) Sort(sort string) ApiGetEventsRequest {
	r.sort = &sort
	return r
}

func (r ApiGetEventsRequest) Execute() ([]Event, *http.Response, error) {
	return r.ApiService.GetEventsExecute(r)
}

/*
GetEvents Find events by type, Booking Reference, Bill of Lading or Equipment Reference.

# Best Practice Guidelines for API Client Configuration

**To ensure the efficient and accurate retrieval of container tracking information via the ONE DCSA Track and Trace v2.2 API, we recommend the following configuration settings for your API client.**

### 1. Number of Containers:
It is anticipated that a significant number of containers, potentially reaching several thousand, may be in active transit within your system, requiring continuous tracking and status updates.

### 2. API Call Rate: **60 API Calls per Minute**
- Our recommendation is that your API polling client should be configured to make **60 API calls per minute**.
- This rate limit is established to maintain the responsiveness and reliability of the ONE API service. Exceeding this rate may lead to throttling or delays in response times.

### 3. Batch Job Execution:
- API consumers who are batching their requests, we recommend scheduling your API client to run **batch jobs up to three times daily** to retrieve tracking information for containers in your system.
- **Batch Job API Calls**: Depending on the volume of containers, each batch job may generate several thousand API calls to the ONE server.

### 4. Use of `equipmentReference`:
- Although `carrierBookingReference` and `transportDocumentReference` are available, we highly recommend using the **`equipmentReference`** (container number) in your API calls whenever possible.
- Using the container number is more efficient in terms of latency, improving response times and reducing processing load.

### 5. Error Handling and Retry Mechanisms:
- Implement robust error handling and retry mechanisms within your API client to manage potential disruptions in connectivity or response errors.
- This will help minimize data retrieval failures and ensure continuous tracking of container movements.

### 6. Data Synchronization:
- To avoid data discrepancies, schedule your API calls to align with the frequency of container status updates provided by the ONE system (Transport Plan).
- This synchronization will help ensure that the tracking data in your system is up-to-date and accurate.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetEventsRequest
*/
func (a *EventsAPIService) GetEvents(ctx context.Context) ApiGetEventsRequest {
	return ApiGetEventsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Event
func (a *EventsAPIService) GetEventsExecute(r ApiGetEventsRequest) ([]Event, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Event
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.GetEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/cargotracking/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apikey == nil {
		return localVarReturnValue, nil, reportError("apikey is required and must be specified")
	}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}

	if r.eventType != nil {
		t := *r.eventType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", t, "form", "multi")
		}
	}
	if r.carrierBookingReference != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "carrierBookingReference", r.carrierBookingReference, "form", "")
	}
	if r.transportDocumentReference != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transportDocumentReference", r.transportDocumentReference, "form", "")
	}
	if r.equipmentReference != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "equipmentReference", r.equipmentReference, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "apikey", r.apikey, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
