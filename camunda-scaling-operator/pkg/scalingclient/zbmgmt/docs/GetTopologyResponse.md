# GetTopologyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int64** | The version of the topology | [optional] 
**Brokers** | Pointer to [**[]BrokerState**](BrokerState.md) |  | [optional] 
**LastChange** | Pointer to [**CompletedChange**](CompletedChange.md) |  | [optional] 
**PendingChange** | Pointer to [**TopologyChange**](TopologyChange.md) |  | [optional] 

## Methods

### NewGetTopologyResponse

`func NewGetTopologyResponse() *GetTopologyResponse`

NewGetTopologyResponse instantiates a new GetTopologyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTopologyResponseWithDefaults

`func NewGetTopologyResponseWithDefaults() *GetTopologyResponse`

NewGetTopologyResponseWithDefaults instantiates a new GetTopologyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *GetTopologyResponse) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetTopologyResponse) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetTopologyResponse) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetTopologyResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBrokers

`func (o *GetTopologyResponse) GetBrokers() []BrokerState`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *GetTopologyResponse) GetBrokersOk() (*[]BrokerState, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *GetTopologyResponse) SetBrokers(v []BrokerState)`

SetBrokers sets Brokers field to given value.

### HasBrokers

`func (o *GetTopologyResponse) HasBrokers() bool`

HasBrokers returns a boolean if a field has been set.

### GetLastChange

`func (o *GetTopologyResponse) GetLastChange() CompletedChange`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *GetTopologyResponse) GetLastChangeOk() (*CompletedChange, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *GetTopologyResponse) SetLastChange(v CompletedChange)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *GetTopologyResponse) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetPendingChange

`func (o *GetTopologyResponse) GetPendingChange() TopologyChange`

GetPendingChange returns the PendingChange field if non-nil, zero value otherwise.

### GetPendingChangeOk

`func (o *GetTopologyResponse) GetPendingChangeOk() (*TopologyChange, bool)`

GetPendingChangeOk returns a tuple with the PendingChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingChange

`func (o *GetTopologyResponse) SetPendingChange(v TopologyChange)`

SetPendingChange sets PendingChange field to given value.

### HasPendingChange

`func (o *GetTopologyResponse) HasPendingChange() bool`

HasPendingChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


