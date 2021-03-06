/*
ONE CARGO TRACKING API

Cargo tacking details is provided based on DCSA standards. (1.2)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package one

import (
	"encoding/json"
)

// InlineResponse400 struct for InlineResponse400
type InlineResponse400 struct {
	ErrorCode *int32 `json:"errorCode,omitempty"`
	ErrorMessages *string `json:"errorMessages,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewInlineResponse400 instantiates a new InlineResponse400 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse400() *InlineResponse400 {
	this := InlineResponse400{}
	return &this
}

// NewInlineResponse400WithDefaults instantiates a new InlineResponse400 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse400WithDefaults() *InlineResponse400 {
	this := InlineResponse400{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *InlineResponse400) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *InlineResponse400) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *InlineResponse400) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise.
func (o *InlineResponse400) GetErrorMessages() string {
	if o == nil || o.ErrorMessages == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessages
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400) GetErrorMessagesOk() (*string, bool) {
	if o == nil || o.ErrorMessages == nil {
		return nil, false
	}
	return o.ErrorMessages, true
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *InlineResponse400) HasErrorMessages() bool {
	if o != nil && o.ErrorMessages != nil {
		return true
	}

	return false
}

// SetErrorMessages gets a reference to the given string and assigns it to the ErrorMessages field.
func (o *InlineResponse400) SetErrorMessages(v string) {
	o.ErrorMessages = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *InlineResponse400) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *InlineResponse400) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *InlineResponse400) SetUuid(v string) {
	o.Uuid = &v
}

func (o InlineResponse400) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorMessages != nil {
		toSerialize["errorMessages"] = o.ErrorMessages
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse400 struct {
	value *InlineResponse400
	isSet bool
}

func (v NullableInlineResponse400) Get() *InlineResponse400 {
	return v.value
}

func (v *NullableInlineResponse400) Set(val *InlineResponse400) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse400) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse400(val *InlineResponse400) *NullableInlineResponse400 {
	return &NullableInlineResponse400{value: val, isSet: true}
}

func (v NullableInlineResponse400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


