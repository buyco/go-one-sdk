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
)

// checks if the TransportCall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportCall{}

// TransportCall struct for TransportCall
type TransportCall struct {
	// The unique identifier for a transport call
	TransportCallID string `json:"transportCallID"`
	// The code of the service for which the schedule details are published.
	CarrierServiceCode *string `json:"carrierServiceCode,omitempty"`
	// The vessel operator-specific identifier of the Voyage.  In case there are multiple voyages the export voyage is chosen.
	// Deprecated
	CarrierVoyageNumber interface{} `json:"carrierVoyageNumber,omitempty"`
	// The vessel operator-specific identifier of the export Voyage.
	ExportVoyageNumber *string `json:"exportVoyageNumber,omitempty"`
	// The vessel operator-specific identifier of the import Voyage.
	ImportVoyageNumber *string `json:"importVoyageNumber,omitempty"`
	// Transport operator&apos;s key that uniquely identifies each individual call. This key is essential to distinguish between two separate calls at the same location within one voyage.
	TransportCallSequenceNumber *int32 `json:"transportCallSequenceNumber,omitempty"`
	// The UN Location code specifying where the place is located.
	UNLocationCode *string `json:"UNLocationCode,omitempty"`
	// The code used for identifying the specific facility. This code does <b>not</b> include the UN Location Code.
	FacilityCode             *string                   `json:"facilityCode,omitempty"`
	FacilityCodeListProvider *FacilityCodeListProvider `json:"facilityCodeListProvider,omitempty"`
	FacilityTypeCode         *FacilityTypeCodeTRN      `json:"facilityTypeCode,omitempty"`
	// An alternative way to capture the facility when no standardized DCSA facility code can be found.
	OtherFacility   *string                     `json:"otherFacility,omitempty"`
	ModeOfTransport ModeOfTransport             `json:"modeOfTransport"`
	Location        *TransportCallAllOfLocation `json:"location,omitempty"`
	Vessel          *Vessel                     `json:"vessel,omitempty"`
}

type _TransportCall TransportCall

// NewTransportCall instantiates a new TransportCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportCall(transportCallID string, modeOfTransport ModeOfTransport) *TransportCall {
	this := TransportCall{}
	this.TransportCallID = transportCallID
	this.ModeOfTransport = modeOfTransport
	return &this
}

// NewTransportCallWithDefaults instantiates a new TransportCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportCallWithDefaults() *TransportCall {
	this := TransportCall{}
	return &this
}

// GetTransportCallID returns the TransportCallID field value
func (o *TransportCall) GetTransportCallID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransportCallID
}

// GetTransportCallIDOk returns a tuple with the TransportCallID field value
// and a boolean to check if the value has been set.
func (o *TransportCall) GetTransportCallIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportCallID, true
}

// SetTransportCallID sets field value
func (o *TransportCall) SetTransportCallID(v string) {
	o.TransportCallID = v
}

// GetCarrierServiceCode returns the CarrierServiceCode field value if set, zero value otherwise.
func (o *TransportCall) GetCarrierServiceCode() string {
	if o == nil || IsNil(o.CarrierServiceCode) {
		var ret string
		return ret
	}
	return *o.CarrierServiceCode
}

// GetCarrierServiceCodeOk returns a tuple with the CarrierServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCall) GetCarrierServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierServiceCode) {
		return nil, false
	}
	return o.CarrierServiceCode, true
}

// HasCarrierServiceCode returns a boolean if a field has been set.
func (o *TransportCall) HasCarrierServiceCode() bool {
	if o != nil && !IsNil(o.CarrierServiceCode) {
		return true
	}

	return false
}

// SetCarrierServiceCode gets a reference to the given string and assigns it to the CarrierServiceCode field.
func (o *TransportCall) SetCarrierServiceCode(v string) {
	o.CarrierServiceCode = &v
}

// GetCarrierVoyageNumber returns the CarrierVoyageNumber field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *TransportCall) GetCarrierVoyageNumber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CarrierVoyageNumber
}

// GetCarrierVoyageNumberOk returns a tuple with the CarrierVoyageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *TransportCall) GetCarrierVoyageNumberOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CarrierVoyageNumber) {
		return nil, false
	}
	return &o.CarrierVoyageNumber, true
}

// HasCarrierVoyageNumber returns a boolean if a field has been set.
func (o *TransportCall) HasCarrierVoyageNumber() bool {
	if o != nil && !IsNil(o.CarrierVoyageNumber) {
		return true
	}

	return false
}

// SetCarrierVoyageNumber gets a reference to the given interface{} and assigns it to the CarrierVoyageNumber field.
// Deprecated
func (o *TransportCall) SetCarrierVoyageNumber(v interface{}) {
	o.CarrierVoyageNumber = v
}

// GetExportVoyageNumber returns the ExportVoyageNumber field value if set, zero value otherwise.
func (o *TransportCall) GetExportVoyageNumber() string {
	if o == nil || IsNil(o.ExportVoyageNumber) {
		var ret string
		return ret
	}
	return *o.ExportVoyageNumber
}

// GetExportVoyageNumberOk returns a tuple with the ExportVoyageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCall) GetExportVoyageNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ExportVoyageNumber) {
		return nil, false
	}
	return o.ExportVoyageNumber, true
}

// HasExportVoyageNumber returns a boolean if a field has been set.
func (o *TransportCall) HasExportVoyageNumber() bool {
	if o != nil && !IsNil(o.ExportVoyageNumber) {
		return true
	}

	return false
}

// SetExportVoyageNumber gets a reference to the given string and assigns it to the ExportVoyageNumber field.
func (o *TransportCall) SetExportVoyageNumber(v string) {
	o.ExportVoyageNumber = &v
}

// GetImportVoyageNumber returns the ImportVoyageNumber field value if set, zero value otherwise.
func (o *TransportCall) GetImportVoyageNumber() string {
	if o == nil || IsNil(o.ImportVoyageNumber) {
		var ret string
		return ret
	}
	return *o.ImportVoyageNumber
}

// GetImportVoyageNumberOk returns a tuple with the ImportVoyageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCall) GetImportVoyageNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ImportVoyageNumber) {
		return nil, false
	}
	return o.ImportVoyageNumber, true
}

// HasImportVoyageNumber returns a boolean if a field has been set.
func (o *TransportCall) HasImportVoyageNumber() bool {
	if o != nil && !IsNil(o.ImportVoyageNumber) {
		return true
	}

	return false
}

// SetImportVoyageNumber gets a reference to the given string and assigns it to the ImportVoyageNumber field.
func (o *TransportCall) SetImportVoyageNumber(v string) {
	o.ImportVoyageNumber = &v
}

// GetTransportCallSequenceNumber returns the TransportCallSequenceNumber field value if set, zero value otherwise.
func (o *TransportCall) GetTransportCallSequenceNumber() int32 {
	if o == nil || IsNil(o.TransportCallSequenceNumber) {
		var ret int32
		return ret
	}
	return *o.TransportCallSequenceNumber
}

// GetTransportCallSequenceNumberOk returns a tuple with the TransportCallSequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCall) GetTransportCallSequenceNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.TransportCallSequenceNumber) {
		return nil, false
	}
	return o.TransportCallSequenceNumber, true
}

// HasTransportCallSequenceNumber returns a boolean if a field has been set.
func (o *TransportCall) HasTransportCallSequenceNumber() bool {
	if o != nil && !IsNil(o.TransportCallSequenceNumber) {
		return true
	}

	return false
}

// SetTransportCallSequenceNumber gets a reference to the given int32 and assigns it to the TransportCallSequenceNumber field.
func (o *TransportCall) SetTransportCallSequenceNumber(v int32) {
	o.TransportCallSequenceNumber = &v
}

// GetUNLocationCode returns the UNLocationCode field value if set, zero value otherwise.
func (o *TransportCall) GetUNLocationCode() string {
	if o == nil || IsNil(o.UNLocationCode) {
		var ret string
		return ret
	}
	return *o.UNLocationCode
}

// GetUNLocationCodeOk returns a tuple with the UNLocationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCall) GetUNLocationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UNLocationCode) {
		return nil, false
	}
	return o.UNLocationCode, true
}

// HasUNLocationCode returns a boolean if a field has been set.
func (o *TransportCall) HasUNLocationCode() bool {
	if o != nil && !IsNil(o.UNLocationCode) {
		return true
	}

	return false
}

// SetUNLocationCode gets a reference to the given string and assigns it to the UNLocationCode field.
func (o *TransportCall) SetUNLocationCode(v string) {
	o.UNLocationCode = &v
}

// GetFacilityCode returns the FacilityCode field value if set, zero value otherwise.
func (o *TransportCall) GetFacilityCode() string {
	if o == nil || IsNil(o.FacilityCode) {
		var ret string
		return ret
	}
	return *o.FacilityCode
}

// GetFacilityCodeOk returns a tuple with the FacilityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCall) GetFacilityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.FacilityCode) {
		return nil, false
	}
	return o.FacilityCode, true
}

// HasFacilityCode returns a boolean if a field has been set.
func (o *TransportCall) HasFacilityCode() bool {
	if o != nil && !IsNil(o.FacilityCode) {
		return true
	}

	return false
}

// SetFacilityCode gets a reference to the given string and assigns it to the FacilityCode field.
func (o *TransportCall) SetFacilityCode(v string) {
	o.FacilityCode = &v
}

// GetFacilityCodeListProvider returns the FacilityCodeListProvider field value if set, zero value otherwise.
func (o *TransportCall) GetFacilityCodeListProvider() FacilityCodeListProvider {
	if o == nil || IsNil(o.FacilityCodeListProvider) {
		var ret FacilityCodeListProvider
		return ret
	}
	return *o.FacilityCodeListProvider
}

// GetFacilityCodeListProviderOk returns a tuple with the FacilityCodeListProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCall) GetFacilityCodeListProviderOk() (*FacilityCodeListProvider, bool) {
	if o == nil || IsNil(o.FacilityCodeListProvider) {
		return nil, false
	}
	return o.FacilityCodeListProvider, true
}

// HasFacilityCodeListProvider returns a boolean if a field has been set.
func (o *TransportCall) HasFacilityCodeListProvider() bool {
	if o != nil && !IsNil(o.FacilityCodeListProvider) {
		return true
	}

	return false
}

// SetFacilityCodeListProvider gets a reference to the given FacilityCodeListProvider and assigns it to the FacilityCodeListProvider field.
func (o *TransportCall) SetFacilityCodeListProvider(v FacilityCodeListProvider) {
	o.FacilityCodeListProvider = &v
}

// GetFacilityTypeCode returns the FacilityTypeCode field value if set, zero value otherwise.
func (o *TransportCall) GetFacilityTypeCode() FacilityTypeCodeTRN {
	if o == nil || IsNil(o.FacilityTypeCode) {
		var ret FacilityTypeCodeTRN
		return ret
	}
	return *o.FacilityTypeCode
}

// GetFacilityTypeCodeOk returns a tuple with the FacilityTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCall) GetFacilityTypeCodeOk() (*FacilityTypeCodeTRN, bool) {
	if o == nil || IsNil(o.FacilityTypeCode) {
		return nil, false
	}
	return o.FacilityTypeCode, true
}

// HasFacilityTypeCode returns a boolean if a field has been set.
func (o *TransportCall) HasFacilityTypeCode() bool {
	if o != nil && !IsNil(o.FacilityTypeCode) {
		return true
	}

	return false
}

// SetFacilityTypeCode gets a reference to the given FacilityTypeCodeTRN and assigns it to the FacilityTypeCode field.
func (o *TransportCall) SetFacilityTypeCode(v FacilityTypeCodeTRN) {
	o.FacilityTypeCode = &v
}

// GetOtherFacility returns the OtherFacility field value if set, zero value otherwise.
func (o *TransportCall) GetOtherFacility() string {
	if o == nil || IsNil(o.OtherFacility) {
		var ret string
		return ret
	}
	return *o.OtherFacility
}

// GetOtherFacilityOk returns a tuple with the OtherFacility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCall) GetOtherFacilityOk() (*string, bool) {
	if o == nil || IsNil(o.OtherFacility) {
		return nil, false
	}
	return o.OtherFacility, true
}

// HasOtherFacility returns a boolean if a field has been set.
func (o *TransportCall) HasOtherFacility() bool {
	if o != nil && !IsNil(o.OtherFacility) {
		return true
	}

	return false
}

// SetOtherFacility gets a reference to the given string and assigns it to the OtherFacility field.
func (o *TransportCall) SetOtherFacility(v string) {
	o.OtherFacility = &v
}

// GetModeOfTransport returns the ModeOfTransport field value
func (o *TransportCall) GetModeOfTransport() ModeOfTransport {
	if o == nil {
		var ret ModeOfTransport
		return ret
	}

	return o.ModeOfTransport
}

// GetModeOfTransportOk returns a tuple with the ModeOfTransport field value
// and a boolean to check if the value has been set.
func (o *TransportCall) GetModeOfTransportOk() (*ModeOfTransport, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModeOfTransport, true
}

// SetModeOfTransport sets field value
func (o *TransportCall) SetModeOfTransport(v ModeOfTransport) {
	o.ModeOfTransport = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *TransportCall) GetLocation() TransportCallAllOfLocation {
	if o == nil || IsNil(o.Location) {
		var ret TransportCallAllOfLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCall) GetLocationOk() (*TransportCallAllOfLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *TransportCall) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given TransportCallAllOfLocation and assigns it to the Location field.
func (o *TransportCall) SetLocation(v TransportCallAllOfLocation) {
	o.Location = &v
}

// GetVessel returns the Vessel field value if set, zero value otherwise.
func (o *TransportCall) GetVessel() Vessel {
	if o == nil || IsNil(o.Vessel) {
		var ret Vessel
		return ret
	}
	return *o.Vessel
}

// GetVesselOk returns a tuple with the Vessel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCall) GetVesselOk() (*Vessel, bool) {
	if o == nil || IsNil(o.Vessel) {
		return nil, false
	}
	return o.Vessel, true
}

// HasVessel returns a boolean if a field has been set.
func (o *TransportCall) HasVessel() bool {
	if o != nil && !IsNil(o.Vessel) {
		return true
	}

	return false
}

// SetVessel gets a reference to the given Vessel and assigns it to the Vessel field.
func (o *TransportCall) SetVessel(v Vessel) {
	o.Vessel = &v
}

func (o TransportCall) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transportCallID"] = o.TransportCallID
	if !IsNil(o.CarrierServiceCode) {
		toSerialize["carrierServiceCode"] = o.CarrierServiceCode
	}
	if o.CarrierVoyageNumber != nil {
		toSerialize["carrierVoyageNumber"] = o.CarrierVoyageNumber
	}
	if !IsNil(o.ExportVoyageNumber) {
		toSerialize["exportVoyageNumber"] = o.ExportVoyageNumber
	}
	if !IsNil(o.ImportVoyageNumber) {
		toSerialize["importVoyageNumber"] = o.ImportVoyageNumber
	}
	if !IsNil(o.TransportCallSequenceNumber) {
		toSerialize["transportCallSequenceNumber"] = o.TransportCallSequenceNumber
	}
	if !IsNil(o.UNLocationCode) {
		toSerialize["UNLocationCode"] = o.UNLocationCode
	}
	if !IsNil(o.FacilityCode) {
		toSerialize["facilityCode"] = o.FacilityCode
	}
	if !IsNil(o.FacilityCodeListProvider) {
		toSerialize["facilityCodeListProvider"] = o.FacilityCodeListProvider
	}
	if !IsNil(o.FacilityTypeCode) {
		toSerialize["facilityTypeCode"] = o.FacilityTypeCode
	}
	if !IsNil(o.OtherFacility) {
		toSerialize["otherFacility"] = o.OtherFacility
	}
	toSerialize["modeOfTransport"] = o.ModeOfTransport
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Vessel) {
		toSerialize["vessel"] = o.Vessel
	}
	return toSerialize, nil
}

func (o *TransportCall) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transportCallID",
		"modeOfTransport",
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

	varTransportCall := _TransportCall{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransportCall)

	if err != nil {
		return err
	}

	*o = TransportCall(varTransportCall)

	return err
}

type NullableTransportCall struct {
	value *TransportCall
	isSet bool
}

func (v NullableTransportCall) Get() *TransportCall {
	return v.value
}

func (v *NullableTransportCall) Set(val *TransportCall) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportCall) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportCall(val *TransportCall) *NullableTransportCall {
	return &NullableTransportCall{value: val, isSet: true}
}

func (v NullableTransportCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
