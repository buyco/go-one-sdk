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
	"encoding/json"
	"fmt"
	"time"
)

// checks if the OperationsTransportEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationsTransportEvent{}

// OperationsTransportEvent struct for OperationsTransportEvent
type OperationsTransportEvent struct {
	// The unique identifier for the event (the message - not the source).  <b>NB</b>&#58; This field should be considered Metadata
	EventID *string `json:"eventID,omitempty"`
	// The timestamp of when the event was created.  <b>NB</b>&#58; This field should be considered Metadata
	EventCreatedDateTime time.Time `json:"eventCreatedDateTime"`
	EventType            string    `json:"eventType"`
	// Code for the event classifier can be - ACT (Actual) - PLN (Planned) - EST (Estimated)
	EventClassifierCode string `json:"eventClassifierCode"`
	// The local date and time, where the event took place or when the event will take place, in ISO 8601 format.
	EventDateTime          string                 `json:"eventDateTime"`
	TransportEventTypeCode TransportEventTypeCode `json:"transportEventTypeCode"`
	// Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/
	DelayReasonCode *string `json:"delayReasonCode,omitempty"`
	// Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage.  Deprecated - use changeRemark instead
	// Deprecated
	VesselScheduleChangeRemark *string `json:"vesselScheduleChangeRemark,omitempty"`
	// Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage.
	ChangeRemark    *string                                 `json:"changeRemark,omitempty"`
	TransportCallID *BaseTransportEventAllOfTransportCallID `json:"transportCallID,omitempty"`
	TransportCall   TransportCall                           `json:"transportCall"`
	// Unique identifier for Event Type Code, for transport events this is either - ARRI (Arrival) - DEPA (Departure)  Deprecated - use transportEventTypeCode instead
	// Deprecated
	EventTypeCode *string `json:"eventTypeCode,omitempty"`
}

type _OperationsTransportEvent OperationsTransportEvent

// NewOperationsTransportEvent instantiates a new OperationsTransportEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationsTransportEvent(eventCreatedDateTime time.Time, eventType string, eventClassifierCode string, eventDateTime string, transportEventTypeCode TransportEventTypeCode, transportCall TransportCall) *OperationsTransportEvent {
	this := OperationsTransportEvent{}
	this.EventCreatedDateTime = eventCreatedDateTime
	this.EventType = eventType
	this.EventClassifierCode = eventClassifierCode
	this.EventDateTime = eventDateTime
	this.TransportEventTypeCode = transportEventTypeCode
	this.TransportCall = transportCall
	return &this
}

// NewOperationsTransportEventWithDefaults instantiates a new OperationsTransportEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationsTransportEventWithDefaults() *OperationsTransportEvent {
	this := OperationsTransportEvent{}
	return &this
}

// GetEventID returns the EventID field value if set, zero value otherwise.
func (o *OperationsTransportEvent) GetEventID() string {
	if o == nil || IsNil(o.EventID) {
		var ret string
		return ret
	}
	return *o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetEventIDOk() (*string, bool) {
	if o == nil || IsNil(o.EventID) {
		return nil, false
	}
	return o.EventID, true
}

// HasEventID returns a boolean if a field has been set.
func (o *OperationsTransportEvent) HasEventID() bool {
	if o != nil && !IsNil(o.EventID) {
		return true
	}

	return false
}

// SetEventID gets a reference to the given string and assigns it to the EventID field.
func (o *OperationsTransportEvent) SetEventID(v string) {
	o.EventID = &v
}

// GetEventCreatedDateTime returns the EventCreatedDateTime field value
func (o *OperationsTransportEvent) GetEventCreatedDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventCreatedDateTime
}

// GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field value
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetEventCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCreatedDateTime, true
}

// SetEventCreatedDateTime sets field value
func (o *OperationsTransportEvent) SetEventCreatedDateTime(v time.Time) {
	o.EventCreatedDateTime = v
}

// GetEventType returns the EventType field value
func (o *OperationsTransportEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *OperationsTransportEvent) SetEventType(v string) {
	o.EventType = v
}

// GetEventClassifierCode returns the EventClassifierCode field value
func (o *OperationsTransportEvent) GetEventClassifierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventClassifierCode, true
}

// SetEventClassifierCode sets field value
func (o *OperationsTransportEvent) SetEventClassifierCode(v string) {
	o.EventClassifierCode = v
}

// GetEventDateTime returns the EventDateTime field value
func (o *OperationsTransportEvent) GetEventDateTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetEventDateTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventDateTime, true
}

// SetEventDateTime sets field value
func (o *OperationsTransportEvent) SetEventDateTime(v string) {
	o.EventDateTime = v
}

// GetTransportEventTypeCode returns the TransportEventTypeCode field value
func (o *OperationsTransportEvent) GetTransportEventTypeCode() TransportEventTypeCode {
	if o == nil {
		var ret TransportEventTypeCode
		return ret
	}

	return o.TransportEventTypeCode
}

// GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field value
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetTransportEventTypeCodeOk() (*TransportEventTypeCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportEventTypeCode, true
}

// SetTransportEventTypeCode sets field value
func (o *OperationsTransportEvent) SetTransportEventTypeCode(v TransportEventTypeCode) {
	o.TransportEventTypeCode = v
}

// GetDelayReasonCode returns the DelayReasonCode field value if set, zero value otherwise.
func (o *OperationsTransportEvent) GetDelayReasonCode() string {
	if o == nil || IsNil(o.DelayReasonCode) {
		var ret string
		return ret
	}
	return *o.DelayReasonCode
}

// GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetDelayReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DelayReasonCode) {
		return nil, false
	}
	return o.DelayReasonCode, true
}

// HasDelayReasonCode returns a boolean if a field has been set.
func (o *OperationsTransportEvent) HasDelayReasonCode() bool {
	if o != nil && !IsNil(o.DelayReasonCode) {
		return true
	}

	return false
}

// SetDelayReasonCode gets a reference to the given string and assigns it to the DelayReasonCode field.
func (o *OperationsTransportEvent) SetDelayReasonCode(v string) {
	o.DelayReasonCode = &v
}

// GetVesselScheduleChangeRemark returns the VesselScheduleChangeRemark field value if set, zero value otherwise.
// Deprecated
func (o *OperationsTransportEvent) GetVesselScheduleChangeRemark() string {
	if o == nil || IsNil(o.VesselScheduleChangeRemark) {
		var ret string
		return ret
	}
	return *o.VesselScheduleChangeRemark
}

// GetVesselScheduleChangeRemarkOk returns a tuple with the VesselScheduleChangeRemark field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *OperationsTransportEvent) GetVesselScheduleChangeRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.VesselScheduleChangeRemark) {
		return nil, false
	}
	return o.VesselScheduleChangeRemark, true
}

// HasVesselScheduleChangeRemark returns a boolean if a field has been set.
func (o *OperationsTransportEvent) HasVesselScheduleChangeRemark() bool {
	if o != nil && !IsNil(o.VesselScheduleChangeRemark) {
		return true
	}

	return false
}

// SetVesselScheduleChangeRemark gets a reference to the given string and assigns it to the VesselScheduleChangeRemark field.
// Deprecated
func (o *OperationsTransportEvent) SetVesselScheduleChangeRemark(v string) {
	o.VesselScheduleChangeRemark = &v
}

// GetChangeRemark returns the ChangeRemark field value if set, zero value otherwise.
func (o *OperationsTransportEvent) GetChangeRemark() string {
	if o == nil || IsNil(o.ChangeRemark) {
		var ret string
		return ret
	}
	return *o.ChangeRemark
}

// GetChangeRemarkOk returns a tuple with the ChangeRemark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetChangeRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeRemark) {
		return nil, false
	}
	return o.ChangeRemark, true
}

// HasChangeRemark returns a boolean if a field has been set.
func (o *OperationsTransportEvent) HasChangeRemark() bool {
	if o != nil && !IsNil(o.ChangeRemark) {
		return true
	}

	return false
}

// SetChangeRemark gets a reference to the given string and assigns it to the ChangeRemark field.
func (o *OperationsTransportEvent) SetChangeRemark(v string) {
	o.ChangeRemark = &v
}

// GetTransportCallID returns the TransportCallID field value if set, zero value otherwise.
func (o *OperationsTransportEvent) GetTransportCallID() BaseTransportEventAllOfTransportCallID {
	if o == nil || IsNil(o.TransportCallID) {
		var ret BaseTransportEventAllOfTransportCallID
		return ret
	}
	return *o.TransportCallID
}

// GetTransportCallIDOk returns a tuple with the TransportCallID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetTransportCallIDOk() (*BaseTransportEventAllOfTransportCallID, bool) {
	if o == nil || IsNil(o.TransportCallID) {
		return nil, false
	}
	return o.TransportCallID, true
}

// HasTransportCallID returns a boolean if a field has been set.
func (o *OperationsTransportEvent) HasTransportCallID() bool {
	if o != nil && !IsNil(o.TransportCallID) {
		return true
	}

	return false
}

// SetTransportCallID gets a reference to the given BaseTransportEventAllOfTransportCallID and assigns it to the TransportCallID field.
func (o *OperationsTransportEvent) SetTransportCallID(v BaseTransportEventAllOfTransportCallID) {
	o.TransportCallID = &v
}

// GetTransportCall returns the TransportCall field value
func (o *OperationsTransportEvent) GetTransportCall() TransportCall {
	if o == nil {
		var ret TransportCall
		return ret
	}

	return o.TransportCall
}

// GetTransportCallOk returns a tuple with the TransportCall field value
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetTransportCallOk() (*TransportCall, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportCall, true
}

// SetTransportCall sets field value
func (o *OperationsTransportEvent) SetTransportCall(v TransportCall) {
	o.TransportCall = v
}

// GetEventTypeCode returns the EventTypeCode field value if set, zero value otherwise.
// Deprecated
func (o *OperationsTransportEvent) GetEventTypeCode() string {
	if o == nil || IsNil(o.EventTypeCode) {
		var ret string
		return ret
	}
	return *o.EventTypeCode
}

// GetEventTypeCodeOk returns a tuple with the EventTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *OperationsTransportEvent) GetEventTypeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EventTypeCode) {
		return nil, false
	}
	return o.EventTypeCode, true
}

// HasEventTypeCode returns a boolean if a field has been set.
func (o *OperationsTransportEvent) HasEventTypeCode() bool {
	if o != nil && !IsNil(o.EventTypeCode) {
		return true
	}

	return false
}

// SetEventTypeCode gets a reference to the given string and assigns it to the EventTypeCode field.
// Deprecated
func (o *OperationsTransportEvent) SetEventTypeCode(v string) {
	o.EventTypeCode = &v
}

func (o OperationsTransportEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationsTransportEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventID) {
		toSerialize["eventID"] = o.EventID
	}
	toSerialize["eventCreatedDateTime"] = o.EventCreatedDateTime
	toSerialize["eventType"] = o.EventType
	toSerialize["eventClassifierCode"] = o.EventClassifierCode
	toSerialize["eventDateTime"] = o.EventDateTime
	toSerialize["transportEventTypeCode"] = o.TransportEventTypeCode
	if !IsNil(o.DelayReasonCode) {
		toSerialize["delayReasonCode"] = o.DelayReasonCode
	}
	if !IsNil(o.VesselScheduleChangeRemark) {
		toSerialize["vesselScheduleChangeRemark"] = o.VesselScheduleChangeRemark
	}
	if !IsNil(o.ChangeRemark) {
		toSerialize["changeRemark"] = o.ChangeRemark
	}
	if !IsNil(o.TransportCallID) {
		toSerialize["transportCallID"] = o.TransportCallID
	}
	toSerialize["transportCall"] = o.TransportCall
	if !IsNil(o.EventTypeCode) {
		toSerialize["eventTypeCode"] = o.EventTypeCode
	}
	return toSerialize, nil
}

func (o *OperationsTransportEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eventCreatedDateTime",
		"eventType",
		"eventClassifierCode",
		"eventDateTime",
		"transportEventTypeCode",
		"transportCall",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOperationsTransportEvent := _OperationsTransportEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOperationsTransportEvent)

	if err != nil {
		return err
	}

	*o = OperationsTransportEvent(varOperationsTransportEvent)

	return err
}

type NullableOperationsTransportEvent struct {
	value *OperationsTransportEvent
	isSet bool
}

func (v NullableOperationsTransportEvent) Get() *OperationsTransportEvent {
	return v.value
}

func (v *NullableOperationsTransportEvent) Set(val *OperationsTransportEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationsTransportEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationsTransportEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationsTransportEvent(val *OperationsTransportEvent) *NullableOperationsTransportEvent {
	return &NullableOperationsTransportEvent{value: val, isSet: true}
}

func (v NullableOperationsTransportEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationsTransportEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
