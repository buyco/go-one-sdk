/*
ONE CARGO TRACKING API

Cargo tacking details is provided based on DCSA standards. (1.2)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package one

import (
	"encoding/json"
	"time"
)

// TransportEvent The transport event entity is a specialization of the event entity to support specification of data that only applies to a transport event.
type TransportEvent struct {
	// The unique identifier for the Equipment Event ID/Transport Event ID/Shipment Event ID.
	EventID string `json:"eventID"`
	EventType string `json:"eventType"`
	// The local date and time, where the event took place, in ISO 8601 format.
	EventDateTime time.Time `json:"eventDateTime"`
	// Code for the event classifier, either PLN, ACT or EST.
	EventClassifierCode string `json:"eventClassifierCode"`
	// Unique identifier for Event Type Code.
	EventTypeCode string `json:"eventTypeCode"`
	// The reference for the transport, e.g. when the mode of transport is a vessel, the transport reference will be the vessel IMO number.
	TransportReference string `json:"transportReference"`
	// The transport leg reference will be specific per mode of transport: - Vessel: Voyage number as specified by the vessel operator - Truck: Not yet specified - Rail: Not yet specified - Barge: Not yet specified. 
	TransportLegReference string `json:"transportLegReference"`
	// The code to identify the specific type of facility.
	FacilityTypeCode string `json:"facilityTypeCode"`
	// The UN Location Code identifies a location in the sense of a city/a town/a village, being the smaller administrative area existing as defined by the competent national authority in each country.
	UNLocationCode string `json:"UNLocationCode"`
	// The code used for identifying the specific facility.
	FacilityCode string `json:"facilityCode"`
	// An alternative way to capture the facility when no standardized DCSA facility code can be found.
	OtherFacility *string `json:"otherFacility,omitempty"`
	// A code specifying a type of transport mode.
	ModeOfTransportCode *int32 `json:"modeOfTransportCode,omitempty"`
}

// NewTransportEvent instantiates a new TransportEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportEvent(eventID string, eventType string, eventDateTime time.Time, eventClassifierCode string, eventTypeCode string, transportReference string, transportLegReference string, facilityTypeCode string, uNLocationCode string, facilityCode string) *TransportEvent {
	this := TransportEvent{}
	this.EventID = eventID
	this.EventType = eventType
	this.EventDateTime = eventDateTime
	this.EventClassifierCode = eventClassifierCode
	this.EventTypeCode = eventTypeCode
	this.TransportReference = transportReference
	this.TransportLegReference = transportLegReference
	this.FacilityTypeCode = facilityTypeCode
	this.UNLocationCode = uNLocationCode
	this.FacilityCode = facilityCode
	return &this
}

// NewTransportEventWithDefaults instantiates a new TransportEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportEventWithDefaults() *TransportEvent {
	this := TransportEvent{}
	return &this
}

// GetEventID returns the EventID field value
func (o *TransportEvent) GetEventID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetEventIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventID, true
}

// SetEventID sets field value
func (o *TransportEvent) SetEventID(v string) {
	o.EventID = v
}

// GetEventType returns the EventType field value
func (o *TransportEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *TransportEvent) SetEventType(v string) {
	o.EventType = v
}

// GetEventDateTime returns the EventDateTime field value
func (o *TransportEvent) GetEventDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetEventDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventDateTime, true
}

// SetEventDateTime sets field value
func (o *TransportEvent) SetEventDateTime(v time.Time) {
	o.EventDateTime = v
}

// GetEventClassifierCode returns the EventClassifierCode field value
func (o *TransportEvent) GetEventClassifierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventClassifierCode, true
}

// SetEventClassifierCode sets field value
func (o *TransportEvent) SetEventClassifierCode(v string) {
	o.EventClassifierCode = v
}

// GetEventTypeCode returns the EventTypeCode field value
func (o *TransportEvent) GetEventTypeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventTypeCode
}

// GetEventTypeCodeOk returns a tuple with the EventTypeCode field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetEventTypeCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventTypeCode, true
}

// SetEventTypeCode sets field value
func (o *TransportEvent) SetEventTypeCode(v string) {
	o.EventTypeCode = v
}

// GetTransportReference returns the TransportReference field value
func (o *TransportEvent) GetTransportReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransportReference
}

// GetTransportReferenceOk returns a tuple with the TransportReference field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetTransportReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransportReference, true
}

// SetTransportReference sets field value
func (o *TransportEvent) SetTransportReference(v string) {
	o.TransportReference = v
}

// GetTransportLegReference returns the TransportLegReference field value
func (o *TransportEvent) GetTransportLegReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransportLegReference
}

// GetTransportLegReferenceOk returns a tuple with the TransportLegReference field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetTransportLegReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransportLegReference, true
}

// SetTransportLegReference sets field value
func (o *TransportEvent) SetTransportLegReference(v string) {
	o.TransportLegReference = v
}

// GetFacilityTypeCode returns the FacilityTypeCode field value
func (o *TransportEvent) GetFacilityTypeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FacilityTypeCode
}

// GetFacilityTypeCodeOk returns a tuple with the FacilityTypeCode field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetFacilityTypeCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FacilityTypeCode, true
}

// SetFacilityTypeCode sets field value
func (o *TransportEvent) SetFacilityTypeCode(v string) {
	o.FacilityTypeCode = v
}

// GetUNLocationCode returns the UNLocationCode field value
func (o *TransportEvent) GetUNLocationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UNLocationCode
}

// GetUNLocationCodeOk returns a tuple with the UNLocationCode field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetUNLocationCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UNLocationCode, true
}

// SetUNLocationCode sets field value
func (o *TransportEvent) SetUNLocationCode(v string) {
	o.UNLocationCode = v
}

// GetFacilityCode returns the FacilityCode field value
func (o *TransportEvent) GetFacilityCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FacilityCode
}

// GetFacilityCodeOk returns a tuple with the FacilityCode field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetFacilityCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FacilityCode, true
}

// SetFacilityCode sets field value
func (o *TransportEvent) SetFacilityCode(v string) {
	o.FacilityCode = v
}

// GetOtherFacility returns the OtherFacility field value if set, zero value otherwise.
func (o *TransportEvent) GetOtherFacility() string {
	if o == nil || o.OtherFacility == nil {
		var ret string
		return ret
	}
	return *o.OtherFacility
}

// GetOtherFacilityOk returns a tuple with the OtherFacility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetOtherFacilityOk() (*string, bool) {
	if o == nil || o.OtherFacility == nil {
		return nil, false
	}
	return o.OtherFacility, true
}

// HasOtherFacility returns a boolean if a field has been set.
func (o *TransportEvent) HasOtherFacility() bool {
	if o != nil && o.OtherFacility != nil {
		return true
	}

	return false
}

// SetOtherFacility gets a reference to the given string and assigns it to the OtherFacility field.
func (o *TransportEvent) SetOtherFacility(v string) {
	o.OtherFacility = &v
}

// GetModeOfTransportCode returns the ModeOfTransportCode field value if set, zero value otherwise.
func (o *TransportEvent) GetModeOfTransportCode() int32 {
	if o == nil || o.ModeOfTransportCode == nil {
		var ret int32
		return ret
	}
	return *o.ModeOfTransportCode
}

// GetModeOfTransportCodeOk returns a tuple with the ModeOfTransportCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetModeOfTransportCodeOk() (*int32, bool) {
	if o == nil || o.ModeOfTransportCode == nil {
		return nil, false
	}
	return o.ModeOfTransportCode, true
}

// HasModeOfTransportCode returns a boolean if a field has been set.
func (o *TransportEvent) HasModeOfTransportCode() bool {
	if o != nil && o.ModeOfTransportCode != nil {
		return true
	}

	return false
}

// SetModeOfTransportCode gets a reference to the given int32 and assigns it to the ModeOfTransportCode field.
func (o *TransportEvent) SetModeOfTransportCode(v int32) {
	o.ModeOfTransportCode = &v
}

func (o TransportEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventID"] = o.EventID
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["eventDateTime"] = o.EventDateTime
	}
	if true {
		toSerialize["eventClassifierCode"] = o.EventClassifierCode
	}
	if true {
		toSerialize["eventTypeCode"] = o.EventTypeCode
	}
	if true {
		toSerialize["transportReference"] = o.TransportReference
	}
	if true {
		toSerialize["transportLegReference"] = o.TransportLegReference
	}
	if true {
		toSerialize["facilityTypeCode"] = o.FacilityTypeCode
	}
	if true {
		toSerialize["UNLocationCode"] = o.UNLocationCode
	}
	if true {
		toSerialize["facilityCode"] = o.FacilityCode
	}
	if o.OtherFacility != nil {
		toSerialize["otherFacility"] = o.OtherFacility
	}
	if o.ModeOfTransportCode != nil {
		toSerialize["modeOfTransportCode"] = o.ModeOfTransportCode
	}
	return json.Marshal(toSerialize)
}

type NullableTransportEvent struct {
	value *TransportEvent
	isSet bool
}

func (v NullableTransportEvent) Get() *TransportEvent {
	return v.value
}

func (v *NullableTransportEvent) Set(val *TransportEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportEvent(val *TransportEvent) *NullableTransportEvent {
	return &NullableTransportEvent{value: val, isSet: true}
}

func (v NullableTransportEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


