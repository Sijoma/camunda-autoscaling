# ExportingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exporters** | Pointer to [**[]ExporterConfig**](ExporterConfig.md) |  | [optional] 

## Methods

### NewExportingConfig

`func NewExportingConfig() *ExportingConfig`

NewExportingConfig instantiates a new ExportingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportingConfigWithDefaults

`func NewExportingConfigWithDefaults() *ExportingConfig`

NewExportingConfigWithDefaults instantiates a new ExportingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExporters

`func (o *ExportingConfig) GetExporters() []ExporterConfig`

GetExporters returns the Exporters field if non-nil, zero value otherwise.

### GetExportersOk

`func (o *ExportingConfig) GetExportersOk() (*[]ExporterConfig, bool)`

GetExportersOk returns a tuple with the Exporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporters

`func (o *ExportingConfig) SetExporters(v []ExporterConfig)`

SetExporters sets Exporters field to given value.

### HasExporters

`func (o *ExportingConfig) HasExporters() bool`

HasExporters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


