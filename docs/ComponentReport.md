# ComponentReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coordinates** | Pointer to **string** | Component coordinates as package-url | [optional] 
**Description** | Pointer to **string** | Component description | [optional] 
**Reference** | Pointer to **string** | Component details reference | [optional] 
**Vulnerabilities** | Pointer to [**[]ComponentReportVulnerability**](ComponentReportVulnerability.md) | Vulnerabilities recorded for component | [optional] 

## Methods

### NewComponentReport

`func NewComponentReport() *ComponentReport`

NewComponentReport instantiates a new ComponentReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentReportWithDefaults

`func NewComponentReportWithDefaults() *ComponentReport`

NewComponentReportWithDefaults instantiates a new ComponentReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinates

`func (o *ComponentReport) GetCoordinates() string`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *ComponentReport) GetCoordinatesOk() (*string, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *ComponentReport) SetCoordinates(v string)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *ComponentReport) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetDescription

`func (o *ComponentReport) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComponentReport) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComponentReport) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComponentReport) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReference

`func (o *ComponentReport) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ComponentReport) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ComponentReport) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ComponentReport) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *ComponentReport) GetVulnerabilities() []ComponentReportVulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *ComponentReport) GetVulnerabilitiesOk() (*[]ComponentReportVulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *ComponentReport) SetVulnerabilities(v []ComponentReportVulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *ComponentReport) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


