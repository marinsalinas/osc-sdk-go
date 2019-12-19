# FiltersClientGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsns** | Pointer to **[]int64** | The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections. | [optional] 
**ClientGatewayIds** | Pointer to **[]string** | The IDs of the client gateways. | [optional] 
**ConnectionTypes** | Pointer to **[]string** | The types of communication tunnels used by the client gateways (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**PublicIps** | Pointer to **[]string** | The public IPv4 addresses of the client gateways. | [optional] 
**States** | Pointer to **[]string** | The states of the client gateways (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the client gateways. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the client gateways. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the client gateways, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 

## Methods

### GetBgpAsns

`func (o *FiltersClientGateway) GetBgpAsns() []int64`

GetBgpAsns returns the BgpAsns field if non-nil, zero value otherwise.

### GetBgpAsnsOk

`func (o *FiltersClientGateway) GetBgpAsnsOk() ([]int64, bool)`

GetBgpAsnsOk returns a tuple with the BgpAsns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBgpAsns

`func (o *FiltersClientGateway) HasBgpAsns() bool`

HasBgpAsns returns a boolean if a field has been set.

### SetBgpAsns

`func (o *FiltersClientGateway) SetBgpAsns(v []int64)`

SetBgpAsns gets a reference to the given []int64 and assigns it to the BgpAsns field.

### GetClientGatewayIds

`func (o *FiltersClientGateway) GetClientGatewayIds() []string`

GetClientGatewayIds returns the ClientGatewayIds field if non-nil, zero value otherwise.

### GetClientGatewayIdsOk

`func (o *FiltersClientGateway) GetClientGatewayIdsOk() ([]string, bool)`

GetClientGatewayIdsOk returns a tuple with the ClientGatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientGatewayIds

`func (o *FiltersClientGateway) HasClientGatewayIds() bool`

HasClientGatewayIds returns a boolean if a field has been set.

### SetClientGatewayIds

`func (o *FiltersClientGateway) SetClientGatewayIds(v []string)`

SetClientGatewayIds gets a reference to the given []string and assigns it to the ClientGatewayIds field.

### GetConnectionTypes

`func (o *FiltersClientGateway) GetConnectionTypes() []string`

GetConnectionTypes returns the ConnectionTypes field if non-nil, zero value otherwise.

### GetConnectionTypesOk

`func (o *FiltersClientGateway) GetConnectionTypesOk() ([]string, bool)`

GetConnectionTypesOk returns a tuple with the ConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectionTypes

`func (o *FiltersClientGateway) HasConnectionTypes() bool`

HasConnectionTypes returns a boolean if a field has been set.

### SetConnectionTypes

`func (o *FiltersClientGateway) SetConnectionTypes(v []string)`

SetConnectionTypes gets a reference to the given []string and assigns it to the ConnectionTypes field.

### GetPublicIps

`func (o *FiltersClientGateway) GetPublicIps() []string`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *FiltersClientGateway) GetPublicIpsOk() ([]string, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIps

`func (o *FiltersClientGateway) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### SetPublicIps

`func (o *FiltersClientGateway) SetPublicIps(v []string)`

SetPublicIps gets a reference to the given []string and assigns it to the PublicIps field.

### GetStates

`func (o *FiltersClientGateway) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersClientGateway) GetStatesOk() ([]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStates

`func (o *FiltersClientGateway) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStates

`func (o *FiltersClientGateway) SetStates(v []string)`

SetStates gets a reference to the given []string and assigns it to the States field.

### GetTagKeys

`func (o *FiltersClientGateway) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersClientGateway) GetTagKeysOk() ([]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagKeys

`func (o *FiltersClientGateway) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### SetTagKeys

`func (o *FiltersClientGateway) SetTagKeys(v []string)`

SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.

### GetTagValues

`func (o *FiltersClientGateway) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersClientGateway) GetTagValuesOk() ([]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagValues

`func (o *FiltersClientGateway) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### SetTagValues

`func (o *FiltersClientGateway) SetTagValues(v []string)`

SetTagValues gets a reference to the given []string and assigns it to the TagValues field.

### GetTags

`func (o *FiltersClientGateway) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersClientGateway) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *FiltersClientGateway) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *FiltersClientGateway) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


