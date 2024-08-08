# PlannedOperationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeId** | Pointer to **int64** | The ID of a topology change operation | [optional] 
**CurrentTopology** | Pointer to [**[]BrokerState**](BrokerState.md) | Current topology of the cluster | [optional] 
**PlannedChanges** | Pointer to [**[]Operation**](Operation.md) | A sequence of operations that will be performed to transform the current topology into the expected topology. | [optional] 
**ExpectedTopology** | Pointer to [**[]BrokerState**](BrokerState.md) | The expected final topology when the planned changes have completed. | [optional] 

## Methods

### NewPlannedOperationsResponse

`func NewPlannedOperationsResponse() *PlannedOperationsResponse`

NewPlannedOperationsResponse instantiates a new PlannedOperationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannedOperationsResponseWithDefaults

`func NewPlannedOperationsResponseWithDefaults() *PlannedOperationsResponse`

NewPlannedOperationsResponseWithDefaults instantiates a new PlannedOperationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeId

`func (o *PlannedOperationsResponse) GetChangeId() int64`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *PlannedOperationsResponse) GetChangeIdOk() (*int64, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *PlannedOperationsResponse) SetChangeId(v int64)`

SetChangeId sets ChangeId field to given value.

### HasChangeId

`func (o *PlannedOperationsResponse) HasChangeId() bool`

HasChangeId returns a boolean if a field has been set.

### GetCurrentTopology

`func (o *PlannedOperationsResponse) GetCurrentTopology() []BrokerState`

GetCurrentTopology returns the CurrentTopology field if non-nil, zero value otherwise.

### GetCurrentTopologyOk

`func (o *PlannedOperationsResponse) GetCurrentTopologyOk() (*[]BrokerState, bool)`

GetCurrentTopologyOk returns a tuple with the CurrentTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTopology

`func (o *PlannedOperationsResponse) SetCurrentTopology(v []BrokerState)`

SetCurrentTopology sets CurrentTopology field to given value.

### HasCurrentTopology

`func (o *PlannedOperationsResponse) HasCurrentTopology() bool`

HasCurrentTopology returns a boolean if a field has been set.

### GetPlannedChanges

`func (o *PlannedOperationsResponse) GetPlannedChanges() []Operation`

GetPlannedChanges returns the PlannedChanges field if non-nil, zero value otherwise.

### GetPlannedChangesOk

`func (o *PlannedOperationsResponse) GetPlannedChangesOk() (*[]Operation, bool)`

GetPlannedChangesOk returns a tuple with the PlannedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedChanges

`func (o *PlannedOperationsResponse) SetPlannedChanges(v []Operation)`

SetPlannedChanges sets PlannedChanges field to given value.

### HasPlannedChanges

`func (o *PlannedOperationsResponse) HasPlannedChanges() bool`

HasPlannedChanges returns a boolean if a field has been set.

### GetExpectedTopology

`func (o *PlannedOperationsResponse) GetExpectedTopology() []BrokerState`

GetExpectedTopology returns the ExpectedTopology field if non-nil, zero value otherwise.

### GetExpectedTopologyOk

`func (o *PlannedOperationsResponse) GetExpectedTopologyOk() (*[]BrokerState, bool)`

GetExpectedTopologyOk returns a tuple with the ExpectedTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedTopology

`func (o *PlannedOperationsResponse) SetExpectedTopology(v []BrokerState)`

SetExpectedTopology sets ExpectedTopology field to given value.

### HasExpectedTopology

`func (o *PlannedOperationsResponse) HasExpectedTopology() bool`

HasExpectedTopology returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


