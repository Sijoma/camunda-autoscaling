# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **string** |  | [optional] 
**BrokerId** | Pointer to **int32** | The ID of a broker, starting from 0 | [optional] 
**PartitionId** | Pointer to **int32** | The ID of a partition, starting from 1 | [optional] 
**Priority** | Pointer to **int32** | The priority of the partition | [optional] 
**Brokers** | Pointer to **[]int32** |  | [optional] 
**ExporterId** | Pointer to **string** | The ID of an exporter | [optional] 

## Methods

### NewOperation

`func NewOperation() *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *Operation) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Operation) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Operation) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *Operation) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetBrokerId

`func (o *Operation) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *Operation) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *Operation) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *Operation) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetPartitionId

`func (o *Operation) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *Operation) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *Operation) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.

### HasPartitionId

`func (o *Operation) HasPartitionId() bool`

HasPartitionId returns a boolean if a field has been set.

### GetPriority

`func (o *Operation) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Operation) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Operation) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Operation) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetBrokers

`func (o *Operation) GetBrokers() []int32`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *Operation) GetBrokersOk() (*[]int32, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *Operation) SetBrokers(v []int32)`

SetBrokers sets Brokers field to given value.

### HasBrokers

`func (o *Operation) HasBrokers() bool`

HasBrokers returns a boolean if a field has been set.

### GetExporterId

`func (o *Operation) GetExporterId() string`

GetExporterId returns the ExporterId field if non-nil, zero value otherwise.

### GetExporterIdOk

`func (o *Operation) GetExporterIdOk() (*string, bool)`

GetExporterIdOk returns a tuple with the ExporterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporterId

`func (o *Operation) SetExporterId(v string)`

SetExporterId sets ExporterId field to given value.

### HasExporterId

`func (o *Operation) HasExporterId() bool`

HasExporterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


