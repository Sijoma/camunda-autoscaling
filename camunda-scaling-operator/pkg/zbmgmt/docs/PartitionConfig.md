# PartitionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exporting** | Pointer to [**ExportingConfig**](ExportingConfig.md) |  | [optional] 

## Methods

### NewPartitionConfig

`func NewPartitionConfig() *PartitionConfig`

NewPartitionConfig instantiates a new PartitionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionConfigWithDefaults

`func NewPartitionConfigWithDefaults() *PartitionConfig`

NewPartitionConfigWithDefaults instantiates a new PartitionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExporting

`func (o *PartitionConfig) GetExporting() ExportingConfig`

GetExporting returns the Exporting field if non-nil, zero value otherwise.

### GetExportingOk

`func (o *PartitionConfig) GetExportingOk() (*ExportingConfig, bool)`

GetExportingOk returns a tuple with the Exporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporting

`func (o *PartitionConfig) SetExporting(v ExportingConfig)`

SetExporting sets Exporting field to given value.

### HasExporting

`func (o *PartitionConfig) HasExporting() bool`

HasExporting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


