/*
Sonatype OSS Index

This documents the available APIs into [Sonatype OSS Index](https://ossindex.sonatype.org/) - API Version: 1-SNAPSHOT (be72c8343baab38a8c598d28dafc78003dce81db).

API version: 2024.323
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ossindex

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ComponentReportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentReportRequest{}

// ComponentReportRequest Component vulnerability report request
type ComponentReportRequest struct {
	// One or more component coordinates as package-url
	Coordinates []string `json:"coordinates"`
}

type _ComponentReportRequest ComponentReportRequest

// NewComponentReportRequest instantiates a new ComponentReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentReportRequest(coordinates []string) *ComponentReportRequest {
	this := ComponentReportRequest{}
	this.Coordinates = coordinates
	return &this
}

// NewComponentReportRequestWithDefaults instantiates a new ComponentReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentReportRequestWithDefaults() *ComponentReportRequest {
	this := ComponentReportRequest{}
	return &this
}

// GetCoordinates returns the Coordinates field value
func (o *ComponentReportRequest) GetCoordinates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value
// and a boolean to check if the value has been set.
func (o *ComponentReportRequest) GetCoordinatesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Coordinates, true
}

// SetCoordinates sets field value
func (o *ComponentReportRequest) SetCoordinates(v []string) {
	o.Coordinates = v
}

func (o ComponentReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentReportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["coordinates"] = o.Coordinates
	return toSerialize, nil
}

func (o *ComponentReportRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"coordinates",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varComponentReportRequest := _ComponentReportRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varComponentReportRequest)

	if err != nil {
		return err
	}

	*o = ComponentReportRequest(varComponentReportRequest)

	return err
}

type NullableComponentReportRequest struct {
	value *ComponentReportRequest
	isSet bool
}

func (v NullableComponentReportRequest) Get() *ComponentReportRequest {
	return v.value
}

func (v *NullableComponentReportRequest) Set(val *ComponentReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentReportRequest(val *ComponentReportRequest) *NullableComponentReportRequest {
	return &NullableComponentReportRequest{value: val, isSet: true}
}

func (v NullableComponentReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


