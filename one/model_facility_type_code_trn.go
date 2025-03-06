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

// FacilityTypeCodeTRN A specialized version of the facilityCode to be used in TransportCalls. The code to identify the specific type of facility. - BOCR (Border crossing) - CLOC (Customer location) - COFS (Container freight station) - COYA (Deprecated - use OFFD intead) - OFFD (Off dock storage) - DEPO (Depot) - INTE (Inland terminal) - POTE (Port terminal) - RAMP (Ramp)
type FacilityTypeCodeTRN string

// List of facilityTypeCodeTRN
const (
	BOCR FacilityTypeCodeTRN = "BOCR"
	CLOC FacilityTypeCodeTRN = "CLOC"
	COFS FacilityTypeCodeTRN = "COFS"
	COYA FacilityTypeCodeTRN = "COYA"
	OFFD FacilityTypeCodeTRN = "OFFD"
	DEPO FacilityTypeCodeTRN = "DEPO"
	INTE FacilityTypeCodeTRN = "INTE"
	POTE FacilityTypeCodeTRN = "POTE"
	RAMP FacilityTypeCodeTRN = "RAMP"
)

// All allowed values of FacilityTypeCodeTRN enum
var AllowedFacilityTypeCodeTRNEnumValues = []FacilityTypeCodeTRN{
	"BOCR",
	"CLOC",
	"COFS",
	"COYA",
	"OFFD",
	"DEPO",
	"INTE",
	"POTE",
	"RAMP",
}

func (v *FacilityTypeCodeTRN) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FacilityTypeCodeTRN(value)
	for _, existing := range AllowedFacilityTypeCodeTRNEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FacilityTypeCodeTRN", value)
}

// NewFacilityTypeCodeTRNFromValue returns a pointer to a valid FacilityTypeCodeTRN
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFacilityTypeCodeTRNFromValue(v string) (*FacilityTypeCodeTRN, error) {
	ev := FacilityTypeCodeTRN(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FacilityTypeCodeTRN: valid values are %v", v, AllowedFacilityTypeCodeTRNEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FacilityTypeCodeTRN) IsValid() bool {
	for _, existing := range AllowedFacilityTypeCodeTRNEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to facilityTypeCodeTRN value
func (v FacilityTypeCodeTRN) Ptr() *FacilityTypeCodeTRN {
	return &v
}

type NullableFacilityTypeCodeTRN struct {
	value *FacilityTypeCodeTRN
	isSet bool
}

func (v NullableFacilityTypeCodeTRN) Get() *FacilityTypeCodeTRN {
	return v.value
}

func (v *NullableFacilityTypeCodeTRN) Set(val *FacilityTypeCodeTRN) {
	v.value = val
	v.isSet = true
}

func (v NullableFacilityTypeCodeTRN) IsSet() bool {
	return v.isSet
}

func (v *NullableFacilityTypeCodeTRN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacilityTypeCodeTRN(val *FacilityTypeCodeTRN) *NullableFacilityTypeCodeTRN {
	return &NullableFacilityTypeCodeTRN{value: val, isSet: true}
}

func (v NullableFacilityTypeCodeTRN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacilityTypeCodeTRN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
