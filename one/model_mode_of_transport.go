/*
ONE Specification for DCSA Track & Trace API v2.2

  # DCSA Track and Trace v2.2    **Overview:**    The **DCSA Track and Trace v2.2 API** provides a standardized and reliable mechanism for tracking container movements throughout the global shipping network. It is designed to align with the Digital Container Shipping Association (DCSA) standards, promoting **interoperability** between carriers, shippers, and third-party logistics providers.    This API allows stakeholders to monitor the **end-to-end journey** of containers, offering insights into key milestones such as:  - **Vessel arrivals**  - **Departures**  - **Equipment loading**  - **Transport events**    By delivering **uniform tracking data** across various carriers and ports, this API reduces manual processes and improves operational efficiency for managing large-scale shipping operations.    ---    ## Key Features:    ### 1. **Event Standardization**    All events related to container transport are defined according to DCSA standards, ensuring consistent terminology and reporting across the entire logistics chain.    ### 2. **On-Demand Data Retrieval**    Customers can pull **up-to-date information** about container locations, equipment status, and shipment events when needed, enabling timely and informed decision-making.    ### 3. **Seamless Integration**    Designed to integrate easily with existing systems, the API offers flexibility for users who manage large-scale shipments or have evolving logistics requirements.    ### 4. **Scalability**    Whether tracking a few containers or managing thousands, the API is optimized to handle **large datasets efficiently**, ensuring timely access to critical shipping information.    ---    ### Conclusion:    The **DCSA Track and Trace v2.2 API** empowers companies to enhance their visibility into global logistics operations, improve communication between parties, and drive better service outcomes through transparency and **real-time information sharing**.

API version: 2.2
Contact: integration.support@one-line.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package one

import (
	"encoding/json"
	"fmt"
)

// ModeOfTransport The mode of transport as defined by DCSA.
type ModeOfTransport string

// List of modeOfTransport
const (
	VESSEL ModeOfTransport = "VESSEL"
	RAIL   ModeOfTransport = "RAIL"
	TRUCK  ModeOfTransport = "TRUCK"
	BARGE  ModeOfTransport = "BARGE"
	FEEDER ModeOfTransport = "FEEDER"
)

// All allowed values of ModeOfTransport enum
var AllowedModeOfTransportEnumValues = []ModeOfTransport{
	"VESSEL",
	"RAIL",
	"TRUCK",
	"BARGE",
	"FEEDER",
}

func (v *ModeOfTransport) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModeOfTransport(value)
	for _, existing := range AllowedModeOfTransportEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModeOfTransport", value)
}

// NewModeOfTransportFromValue returns a pointer to a valid ModeOfTransport
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModeOfTransportFromValue(v string) (*ModeOfTransport, error) {
	ev := ModeOfTransport(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModeOfTransport: valid values are %v", v, AllowedModeOfTransportEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModeOfTransport) IsValid() bool {
	for _, existing := range AllowedModeOfTransportEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to modeOfTransport value
func (v ModeOfTransport) Ptr() *ModeOfTransport {
	return &v
}

type NullableModeOfTransport struct {
	value *ModeOfTransport
	isSet bool
}

func (v NullableModeOfTransport) Get() *ModeOfTransport {
	return v.value
}

func (v *NullableModeOfTransport) Set(val *ModeOfTransport) {
	v.value = val
	v.isSet = true
}

func (v NullableModeOfTransport) IsSet() bool {
	return v.isSet
}

func (v *NullableModeOfTransport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModeOfTransport(val *ModeOfTransport) *NullableModeOfTransport {
	return &NullableModeOfTransport{value: val, isSet: true}
}

func (v NullableModeOfTransport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModeOfTransport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
