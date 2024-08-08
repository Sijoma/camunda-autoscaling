# BrokerState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of a broker, starting from 0 | [optional] 
**State** | Pointer to [**BrokerStateCode**](BrokerStateCode.md) |  | [optional] 
**Version** | Pointer to **int64** | The version of the broker state | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | The time when the broker state was last updated | [optional] 
**Partitions** | Pointer to [**[]PartitionState**](PartitionState.md) |  | [optional] 

## Methods

### NewBrokerState

`func NewBrokerState() *BrokerState`

NewBrokerState instantiates a new BrokerState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerStateWithDefaults

`func NewBrokerStateWithDefaults() *BrokerState`

NewBrokerStateWithDefaults instantiates a new BrokerState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BrokerState) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BrokerState) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BrokerState) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BrokerState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *BrokerState) GetState() BrokerStateCode`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BrokerState) GetStateOk() (*BrokerStateCode, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BrokerState) SetState(v BrokerStateCode)`

SetState sets State field to given value.

### HasState

`func (o *BrokerState) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVersion

`func (o *BrokerState) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BrokerState) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BrokerState) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BrokerState) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *BrokerState) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *BrokerState) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *BrokerState) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *BrokerState) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetPartitions

`func (o *BrokerState) GetPartitions() []PartitionState`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *BrokerState) GetPartitionsOk() (*[]PartitionState, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *BrokerState) SetPartitions(v []PartitionState)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *BrokerState) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


