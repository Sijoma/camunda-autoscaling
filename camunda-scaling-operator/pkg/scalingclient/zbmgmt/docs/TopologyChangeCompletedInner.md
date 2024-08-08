# TopologyChangeCompletedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **string** |  | [optional] 
**BrokerId** | Pointer to **int32** | The ID of a broker, starting from 0 | [optional] 
**PartitionId** | Pointer to **int32** | The ID of a partition, starting from 1 | [optional] 
**Priority** | Pointer to **int32** | The priority of the partition | [optional] 
**Brokers** | Pointer to **[]int32** |  | [optional] 
**ExporterId** | Pointer to **string** | The ID of an exporter | [optional] 
**CompletedAt** | Pointer to **time.Time** | The time when the operation was completed | [optional] 

## Methods

### NewTopologyChangeCompletedInner

`func NewTopologyChangeCompletedInner() *TopologyChangeCompletedInner`

NewTopologyChangeCompletedInner instantiates a new TopologyChangeCompletedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyChangeCompletedInnerWithDefaults

`func NewTopologyChangeCompletedInnerWithDefaults() *TopologyChangeCompletedInner`

NewTopologyChangeCompletedInnerWithDefaults instantiates a new TopologyChangeCompletedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *TopologyChangeCompletedInner) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *TopologyChangeCompletedInner) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *TopologyChangeCompletedInner) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *TopologyChangeCompletedInner) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetBrokerId

`func (o *TopologyChangeCompletedInner) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *TopologyChangeCompletedInner) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *TopologyChangeCompletedInner) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *TopologyChangeCompletedInner) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetPartitionId

`func (o *TopologyChangeCompletedInner) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *TopologyChangeCompletedInner) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *TopologyChangeCompletedInner) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.

### HasPartitionId

`func (o *TopologyChangeCompletedInner) HasPartitionId() bool`

HasPartitionId returns a boolean if a field has been set.

### GetPriority

`func (o *TopologyChangeCompletedInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TopologyChangeCompletedInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TopologyChangeCompletedInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TopologyChangeCompletedInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetBrokers

`func (o *TopologyChangeCompletedInner) GetBrokers() []int32`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *TopologyChangeCompletedInner) GetBrokersOk() (*[]int32, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *TopologyChangeCompletedInner) SetBrokers(v []int32)`

SetBrokers sets Brokers field to given value.

### HasBrokers

`func (o *TopologyChangeCompletedInner) HasBrokers() bool`

HasBrokers returns a boolean if a field has been set.

### GetExporterId

`func (o *TopologyChangeCompletedInner) GetExporterId() string`

GetExporterId returns the ExporterId field if non-nil, zero value otherwise.

### GetExporterIdOk

`func (o *TopologyChangeCompletedInner) GetExporterIdOk() (*string, bool)`

GetExporterIdOk returns a tuple with the ExporterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporterId

`func (o *TopologyChangeCompletedInner) SetExporterId(v string)`

SetExporterId sets ExporterId field to given value.

### HasExporterId

`func (o *TopologyChangeCompletedInner) HasExporterId() bool`

HasExporterId returns a boolean if a field has been set.

### GetCompletedAt

`func (o *TopologyChangeCompletedInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *TopologyChangeCompletedInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *TopologyChangeCompletedInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *TopologyChangeCompletedInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


