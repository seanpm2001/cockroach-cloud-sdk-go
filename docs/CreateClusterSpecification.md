# CreateClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | Pointer to [**DedicatedClusterCreateSpecification**](DedicatedClusterCreateSpecification.md) |  | [optional] 
**Serverless** | Pointer to [**ServerlessClusterCreateSpecification**](ServerlessClusterCreateSpecification.md) |  | [optional] 

## Methods

### NewCreateClusterSpecification

`func NewCreateClusterSpecification() *CreateClusterSpecification`

NewCreateClusterSpecification instantiates a new CreateClusterSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetDedicated

`func (o *CreateClusterSpecification) GetDedicated() DedicatedClusterCreateSpecification`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### SetDedicated

`func (o *CreateClusterSpecification) SetDedicated(v DedicatedClusterCreateSpecification)`

SetDedicated sets Dedicated field to given value.

### GetServerless

`func (o *CreateClusterSpecification) GetServerless() ServerlessClusterCreateSpecification`

GetServerless returns the Serverless field if non-nil, zero value otherwise.

### SetServerless

`func (o *CreateClusterSpecification) SetServerless(v ServerlessClusterCreateSpecification)`

SetServerless sets Serverless field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


