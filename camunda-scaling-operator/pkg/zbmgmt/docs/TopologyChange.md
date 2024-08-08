# TopologyChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of a topology change operation | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** | The time when the topology change was started | [optional] 
**CompletedAt** | Pointer to **time.Time** | The time when the topology change was completed | [optional] 
**InternalVersion** | Pointer to **int64** | The internal version of the topology change | [optional] 
**Completed** | Pointer to [**[]TopologyChangeCompletedInner**](TopologyChangeCompletedInner.md) | The list of operations that have been completed if the change status is not COMPLETED. | [optional] 
**Pending** | Pointer to [**[]Operation**](Operation.md) | The list of operations that are pending. | [optional] 

## Methods

### NewTopologyChange

`func NewTopologyChange() *TopologyChange`

NewTopologyChange instantiates a new TopologyChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyChangeWithDefaults

`func NewTopologyChangeWithDefaults() *TopologyChange`

NewTopologyChangeWithDefaults instantiates a new TopologyChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TopologyChange) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopologyChange) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopologyChange) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TopologyChange) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *TopologyChange) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TopologyChange) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TopologyChange) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TopologyChange) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *TopologyChange) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TopologyChange) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TopologyChange) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TopologyChange) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *TopologyChange) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *TopologyChange) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *TopologyChange) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *TopologyChange) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetInternalVersion

`func (o *TopologyChange) GetInternalVersion() int64`

GetInternalVersion returns the InternalVersion field if non-nil, zero value otherwise.

### GetInternalVersionOk

`func (o *TopologyChange) GetInternalVersionOk() (*int64, bool)`

GetInternalVersionOk returns a tuple with the InternalVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalVersion

`func (o *TopologyChange) SetInternalVersion(v int64)`

SetInternalVersion sets InternalVersion field to given value.

### HasInternalVersion

`func (o *TopologyChange) HasInternalVersion() bool`

HasInternalVersion returns a boolean if a field has been set.

### GetCompleted

`func (o *TopologyChange) GetCompleted() []TopologyChangeCompletedInner`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *TopologyChange) GetCompletedOk() (*[]TopologyChangeCompletedInner, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *TopologyChange) SetCompleted(v []TopologyChangeCompletedInner)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *TopologyChange) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetPending

`func (o *TopologyChange) GetPending() []Operation`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *TopologyChange) GetPendingOk() (*[]Operation, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *TopologyChange) SetPending(v []Operation)`

SetPending sets Pending field to given value.

### HasPending

`func (o *TopologyChange) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


