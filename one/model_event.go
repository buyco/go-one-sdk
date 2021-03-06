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

// Event The event entity is described as a generalization of all the specific event categories. An event always takes place in relation to a shipment and can additionally be linked to a transport or an equipment
type Event struct {
	// The unique identifier for the Equipment Event ID/Transport Event ID/Shipment Event ID.
	EventID string `json:"eventID"`
	// The Event Type of the object.
	EventType string `json:"eventType"`
	// The local date and time, where the event took place, in ISO 8601 format.
	EventDateTime time.Time `json:"eventDateTime"`
	// Code for the event classifier, either PLN, ACT or EST.
	EventClassifierCode string `json:"eventClassifierCode"`
	// Unique identifier for Event Type Code.
	EventTypeCode string `json:"eventTypeCode"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(eventID string, eventType string, eventDateTime time.Time, eventClassifierCode string, eventTypeCode string) *Event {
	this := Event{}
	this.EventID = eventID
	this.EventType = eventType
	this.EventDateTime = eventDateTime
	this.EventClassifierCode = eventClassifierCode
	this.EventTypeCode = eventTypeCode
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetEventID returns the EventID field value
func (o *Event) GetEventID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value
// and a boolean to check if the value has been set.
func (o *Event) GetEventIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventID, true
}

// SetEventID sets field value
func (o *Event) SetEventID(v string) {
	o.EventID = v
}

// GetEventType returns the EventType field value
func (o *Event) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *Event) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *Event) SetEventType(v string) {
	o.EventType = v
}

// GetEventDateTime returns the EventDateTime field value
func (o *Event) GetEventDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value
// and a boolean to check if the value has been set.
func (o *Event) GetEventDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventDateTime, true
}

// SetEventDateTime sets field value
func (o *Event) SetEventDateTime(v time.Time) {
	o.EventDateTime = v
}

// GetEventClassifierCode returns the EventClassifierCode field value
func (o *Event) GetEventClassifierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value
// and a boolean to check if the value has been set.
func (o *Event) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventClassifierCode, true
}

// SetEventClassifierCode sets field value
func (o *Event) SetEventClassifierCode(v string) {
	o.EventClassifierCode = v
}

// GetEventTypeCode returns the EventTypeCode field value
func (o *Event) GetEventTypeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventTypeCode
}

// GetEventTypeCodeOk returns a tuple with the EventTypeCode field value
// and a boolean to check if the value has been set.
func (o *Event) GetEventTypeCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventTypeCode, true
}

// SetEventTypeCode sets field value
func (o *Event) SetEventTypeCode(v string) {
	o.EventTypeCode = v
}

func (o Event) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


