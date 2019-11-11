# PublicIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkPublicIpId** | Pointer to **string** | (Required in a Net) The ID representing the association of the EIP with the VM or the NIC. | [optional] 
**NicAccountId** | Pointer to **string** | The account ID of the owner of the NIC. | [optional] 
**NicId** | Pointer to **string** | The ID of the NIC the EIP is associated with (if any). | [optional] 
**PrivateIp** | Pointer to **string** | The private IP address associated with the EIP. | [optional] 
**PublicIp** | Pointer to **string** | The External IP address (EIP) associated with the NAT service. | [optional] 
**PublicIpId** | Pointer to **string** | The allocation ID of the EIP associated with the NAT service. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the EIP. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM the External IP (EIP) is associated with (if any). | [optional] 

## Methods

### GetLinkPublicIpId

`func (o *PublicIp) GetLinkPublicIpId() string`

GetLinkPublicIpId returns the LinkPublicIpId field if non-nil, zero value otherwise.

### GetLinkPublicIpIdOk

`func (o *PublicIp) GetLinkPublicIpIdOk() (string, bool)`

GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkPublicIpId

`func (o *PublicIp) HasLinkPublicIpId() bool`

HasLinkPublicIpId returns a boolean if a field has been set.

### SetLinkPublicIpId

`func (o *PublicIp) SetLinkPublicIpId(v string)`

SetLinkPublicIpId gets a reference to the given string and assigns it to the LinkPublicIpId field.

### GetNicAccountId

`func (o *PublicIp) GetNicAccountId() string`

GetNicAccountId returns the NicAccountId field if non-nil, zero value otherwise.

### GetNicAccountIdOk

`func (o *PublicIp) GetNicAccountIdOk() (string, bool)`

GetNicAccountIdOk returns a tuple with the NicAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNicAccountId

`func (o *PublicIp) HasNicAccountId() bool`

HasNicAccountId returns a boolean if a field has been set.

### SetNicAccountId

`func (o *PublicIp) SetNicAccountId(v string)`

SetNicAccountId gets a reference to the given string and assigns it to the NicAccountId field.

### GetNicId

`func (o *PublicIp) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *PublicIp) GetNicIdOk() (string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNicId

`func (o *PublicIp) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### SetNicId

`func (o *PublicIp) SetNicId(v string)`

SetNicId gets a reference to the given string and assigns it to the NicId field.

### GetPrivateIp

`func (o *PublicIp) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *PublicIp) GetPrivateIpOk() (string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateIp

`func (o *PublicIp) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIp

`func (o *PublicIp) SetPrivateIp(v string)`

SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.

### GetPublicIp

`func (o *PublicIp) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *PublicIp) GetPublicIpOk() (string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIp

`func (o *PublicIp) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIp

`func (o *PublicIp) SetPublicIp(v string)`

SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.

### GetPublicIpId

`func (o *PublicIp) GetPublicIpId() string`

GetPublicIpId returns the PublicIpId field if non-nil, zero value otherwise.

### GetPublicIpIdOk

`func (o *PublicIp) GetPublicIpIdOk() (string, bool)`

GetPublicIpIdOk returns a tuple with the PublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIpId

`func (o *PublicIp) HasPublicIpId() bool`

HasPublicIpId returns a boolean if a field has been set.

### SetPublicIpId

`func (o *PublicIp) SetPublicIpId(v string)`

SetPublicIpId gets a reference to the given string and assigns it to the PublicIpId field.

### GetTags

`func (o *PublicIp) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PublicIp) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *PublicIp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *PublicIp) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.

### GetVmId

`func (o *PublicIp) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *PublicIp) GetVmIdOk() (string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmId

`func (o *PublicIp) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmId

`func (o *PublicIp) SetVmId(v string)`

SetVmId gets a reference to the given string and assigns it to the VmId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


