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

// checks if the BaseEventBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseEventBody{}

// BaseEventBody The Event entity is described as a generalization of all the specific event categories. An event always takes place in relation to a shipment and can additionally be linked to a transport or an equipment
type BaseEventBody struct {
	// The Event Type of the object - to be used as a discriminator.  <b>NB</b>&#58; This field should be considered Metadata
	EventType string `json:"eventType"`
	// Code for the event classifier. Values can vary depending on eventType
	EventClassifierCode string `json:"eventClassifierCode"`
	// The local date and time, where the event took place or when the event will take place, in ISO 8601 format.
	EventDateTime time.Time `json:"eventDateTime"`
}

type _BaseEventBody BaseEventBody

// NewBaseEventBody instantiates a new BaseEventBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEventBody(eventType string, eventClassifierCode string, eventDateTime time.Time) *BaseEventBody {
	this := BaseEventBody{}
	this.EventType = eventType
	this.EventClassifierCode = eventClassifierCode
	this.EventDateTime = eventDateTime
	return &this
}

// NewBaseEventBodyWithDefaults instantiates a new BaseEventBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEventBodyWithDefaults() *BaseEventBody {
	this := BaseEventBody{}
	return &this
}

// GetEventType returns the EventType field value
func (o *BaseEventBody) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *BaseEventBody) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *BaseEventBody) SetEventType(v string) {
	o.EventType = v
}

// GetEventClassifierCode returns the EventClassifierCode field value
func (o *BaseEventBody) GetEventClassifierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value
// and a boolean to check if the value has been set.
func (o *BaseEventBody) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventClassifierCode, true
}

// SetEventClassifierCode sets field value
func (o *BaseEventBody) SetEventClassifierCode(v string) {
	o.EventClassifierCode = v
}

// GetEventDateTime returns the EventDateTime field value
func (o *BaseEventBody) GetEventDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value
// and a boolean to check if the value has been set.
func (o *BaseEventBody) GetEventDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventDateTime, true
}

// SetEventDateTime sets field value
func (o *BaseEventBody) SetEventDateTime(v time.Time) {
	o.EventDateTime = v
}

func (o BaseEventBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseEventBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventType"] = o.EventType
	toSerialize["eventClassifierCode"] = o.EventClassifierCode
	toSerialize["eventDateTime"] = o.EventDateTime
	return toSerialize, nil
}

func (o *BaseEventBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eventType",
		"eventClassifierCode",
		"eventDateTime",
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

	varBaseEventBody := _BaseEventBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseEventBody)

	if err != nil {
		return err
	}

	*o = BaseEventBody(varBaseEventBody)

	return err
}

type NullableBaseEventBody struct {
	value *BaseEventBody
	isSet bool
}

func (v NullableBaseEventBody) Get() *BaseEventBody {
	return v.value
}

func (v *NullableBaseEventBody) Set(val *BaseEventBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEventBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEventBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEventBody(val *BaseEventBody) *NullableBaseEventBody {
	return &NullableBaseEventBody{value: val, isSet: true}
}

func (v NullableBaseEventBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEventBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
