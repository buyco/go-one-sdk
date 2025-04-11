# TransportCallAllOfLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationName** | Pointer to **string** | The name of the location. | [optional] 
**Latitude** | Pointer to **string** | Geographic coordinate that specifies the north–south position of a point on the Earth&amp;apos;s surface. | [optional] 
**Longitude** | Pointer to **string** | Geographic coordinate that specifies the east–west position of a point on the Earth&amp;apos;s surface. | [optional] 
**UNLocationCode** | Pointer to **string** | The UN Location code specifying where the place is located. | [optional] 
**FacilityCode** | Pointer to **string** | The code used for identifying the specific facility. This code does &lt;b&gt;not&lt;/b&gt; include the UN Location Code.  | [optional] 
**FacilityCodeListProvider** | Pointer to [**FacilityCodeListProvider**](FacilityCodeListProvider.md) |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) | Address related information | [optional] 

## Methods

### NewTransportCallAllOfLocation

`func NewTransportCallAllOfLocation() *TransportCallAllOfLocation`

NewTransportCallAllOfLocation instantiates a new TransportCallAllOfLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportCallAllOfLocationWithDefaults

`func NewTransportCallAllOfLocationWithDefaults() *TransportCallAllOfLocation`

NewTransportCallAllOfLocationWithDefaults instantiates a new TransportCallAllOfLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationName

`func (o *TransportCallAllOfLocation) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *TransportCallAllOfLocation) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *TransportCallAllOfLocation) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *TransportCallAllOfLocation) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetLatitude

`func (o *TransportCallAllOfLocation) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *TransportCallAllOfLocation) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *TransportCallAllOfLocation) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *TransportCallAllOfLocation) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *TransportCallAllOfLocation) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *TransportCallAllOfLocation) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *TransportCallAllOfLocation) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *TransportCallAllOfLocation) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetUNLocationCode

`func (o *TransportCallAllOfLocation) GetUNLocationCode() string`

GetUNLocationCode returns the UNLocationCode field if non-nil, zero value otherwise.

### GetUNLocationCodeOk

`func (o *TransportCallAllOfLocation) GetUNLocationCodeOk() (*string, bool)`

GetUNLocationCodeOk returns a tuple with the UNLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLocationCode

`func (o *TransportCallAllOfLocation) SetUNLocationCode(v string)`

SetUNLocationCode sets UNLocationCode field to given value.

### HasUNLocationCode

`func (o *TransportCallAllOfLocation) HasUNLocationCode() bool`

HasUNLocationCode returns a boolean if a field has been set.

### GetFacilityCode

`func (o *TransportCallAllOfLocation) GetFacilityCode() string`

GetFacilityCode returns the FacilityCode field if non-nil, zero value otherwise.

### GetFacilityCodeOk

`func (o *TransportCallAllOfLocation) GetFacilityCodeOk() (*string, bool)`

GetFacilityCodeOk returns a tuple with the FacilityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCode

`func (o *TransportCallAllOfLocation) SetFacilityCode(v string)`

SetFacilityCode sets FacilityCode field to given value.

### HasFacilityCode

`func (o *TransportCallAllOfLocation) HasFacilityCode() bool`

HasFacilityCode returns a boolean if a field has been set.

### GetFacilityCodeListProvider

`func (o *TransportCallAllOfLocation) GetFacilityCodeListProvider() FacilityCodeListProvider`

GetFacilityCodeListProvider returns the FacilityCodeListProvider field if non-nil, zero value otherwise.

### GetFacilityCodeListProviderOk

`func (o *TransportCallAllOfLocation) GetFacilityCodeListProviderOk() (*FacilityCodeListProvider, bool)`

GetFacilityCodeListProviderOk returns a tuple with the FacilityCodeListProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCodeListProvider

`func (o *TransportCallAllOfLocation) SetFacilityCodeListProvider(v FacilityCodeListProvider)`

SetFacilityCodeListProvider sets FacilityCodeListProvider field to given value.

### HasFacilityCodeListProvider

`func (o *TransportCallAllOfLocation) HasFacilityCodeListProvider() bool`

HasFacilityCodeListProvider returns a boolean if a field has been set.

### GetAddress

`func (o *TransportCallAllOfLocation) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransportCallAllOfLocation) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransportCallAllOfLocation) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransportCallAllOfLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


