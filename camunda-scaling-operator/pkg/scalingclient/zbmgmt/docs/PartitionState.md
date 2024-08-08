# PartitionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of a partition, starting from 1 | [optional] 
**State** | Pointer to [**PartitionStateCode**](PartitionStateCode.md) |  | [optional] 
**Priority** | Pointer to **int32** | The priority of the partition | [optional] 
**Config** | Pointer to [**PartitionConfig**](PartitionConfig.md) |  | [optional] 

## Methods

### NewPartitionState

`func NewPartitionState() *PartitionState`

NewPartitionState instantiates a new PartitionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionStateWithDefaults

`func NewPartitionStateWithDefaults() *PartitionState`

NewPartitionStateWithDefaults instantiates a new PartitionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartitionState) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartitionState) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartitionState) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PartitionState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *PartitionState) GetState() PartitionStateCode`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PartitionState) GetStateOk() (*PartitionStateCode, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PartitionState) SetState(v PartitionStateCode)`

SetState sets State field to given value.

### HasState

`func (o *PartitionState) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPriority

`func (o *PartitionState) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PartitionState) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PartitionState) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PartitionState) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetConfig

`func (o *PartitionState) GetConfig() PartitionConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PartitionState) GetConfigOk() (*PartitionConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PartitionState) SetConfig(v PartitionConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PartitionState) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


