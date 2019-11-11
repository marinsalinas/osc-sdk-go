# RouteLight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIpRange** | Pointer to **string** | The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24). | [optional] 
**RouteType** | Pointer to **string** | The type of route (always &#x60;static&#x60;). | [optional] 
**State** | Pointer to **string** | The current state of the static route (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 

## Methods

### GetDestinationIpRange

`func (o *RouteLight) GetDestinationIpRange() string`

GetDestinationIpRange returns the DestinationIpRange field if non-nil, zero value otherwise.

### GetDestinationIpRangeOk

`func (o *RouteLight) GetDestinationIpRangeOk() (string, bool)`

GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDestinationIpRange

`func (o *RouteLight) HasDestinationIpRange() bool`

HasDestinationIpRange returns a boolean if a field has been set.

### SetDestinationIpRange

`func (o *RouteLight) SetDestinationIpRange(v string)`

SetDestinationIpRange gets a reference to the given string and assigns it to the DestinationIpRange field.

### GetRouteType

`func (o *RouteLight) GetRouteType() string`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *RouteLight) GetRouteTypeOk() (string, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteType

`func (o *RouteLight) HasRouteType() bool`

HasRouteType returns a boolean if a field has been set.

### SetRouteType

`func (o *RouteLight) SetRouteType(v string)`

SetRouteType gets a reference to the given string and assigns it to the RouteType field.

### GetState

`func (o *RouteLight) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RouteLight) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *RouteLight) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *RouteLight) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


