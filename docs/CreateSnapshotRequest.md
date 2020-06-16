# CreateSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description for the snapshot. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**FileLocation** | Pointer to **string** | The pre-signed URL of the snapshot you want to import from the OSU bucket. | [optional] 
**SnapshotSize** | Pointer to **int64** | The size of the snapshot created in your account, in bytes. This size must be exactly the same as the source snapshot one. | [optional] 
**SourceRegionName** | Pointer to **string** | The name of the source Region, which must be the same as the Region of your account. | [optional] 
**SourceSnapshotId** | Pointer to **string** | The ID of the snapshot you want to copy. | [optional] 
**VolumeId** | Pointer to **string** | The ID of the volume you want to create a snapshot of. | [optional] 

## Methods

### GetDescription

`func (o *CreateSnapshotRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSnapshotRequest) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CreateSnapshotRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CreateSnapshotRequest) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetDryRun

`func (o *CreateSnapshotRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateSnapshotRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateSnapshotRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateSnapshotRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetFileLocation

`func (o *CreateSnapshotRequest) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *CreateSnapshotRequest) GetFileLocationOk() (string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileLocation

`func (o *CreateSnapshotRequest) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### SetFileLocation

`func (o *CreateSnapshotRequest) SetFileLocation(v string)`

SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.

### GetSnapshotSize

`func (o *CreateSnapshotRequest) GetSnapshotSize() int64`

GetSnapshotSize returns the SnapshotSize field if non-nil, zero value otherwise.

### GetSnapshotSizeOk

`func (o *CreateSnapshotRequest) GetSnapshotSizeOk() (int64, bool)`

GetSnapshotSizeOk returns a tuple with the SnapshotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSnapshotSize

`func (o *CreateSnapshotRequest) HasSnapshotSize() bool`

HasSnapshotSize returns a boolean if a field has been set.

### SetSnapshotSize

`func (o *CreateSnapshotRequest) SetSnapshotSize(v int64)`

SetSnapshotSize gets a reference to the given int64 and assigns it to the SnapshotSize field.

### GetSourceRegionName

`func (o *CreateSnapshotRequest) GetSourceRegionName() string`

GetSourceRegionName returns the SourceRegionName field if non-nil, zero value otherwise.

### GetSourceRegionNameOk

`func (o *CreateSnapshotRequest) GetSourceRegionNameOk() (string, bool)`

GetSourceRegionNameOk returns a tuple with the SourceRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceRegionName

`func (o *CreateSnapshotRequest) HasSourceRegionName() bool`

HasSourceRegionName returns a boolean if a field has been set.

### SetSourceRegionName

`func (o *CreateSnapshotRequest) SetSourceRegionName(v string)`

SetSourceRegionName gets a reference to the given string and assigns it to the SourceRegionName field.

### GetSourceSnapshotId

`func (o *CreateSnapshotRequest) GetSourceSnapshotId() string`

GetSourceSnapshotId returns the SourceSnapshotId field if non-nil, zero value otherwise.

### GetSourceSnapshotIdOk

`func (o *CreateSnapshotRequest) GetSourceSnapshotIdOk() (string, bool)`

GetSourceSnapshotIdOk returns a tuple with the SourceSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceSnapshotId

`func (o *CreateSnapshotRequest) HasSourceSnapshotId() bool`

HasSourceSnapshotId returns a boolean if a field has been set.

### SetSourceSnapshotId

`func (o *CreateSnapshotRequest) SetSourceSnapshotId(v string)`

SetSourceSnapshotId gets a reference to the given string and assigns it to the SourceSnapshotId field.

### GetVolumeId

`func (o *CreateSnapshotRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *CreateSnapshotRequest) GetVolumeIdOk() (string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeId

`func (o *CreateSnapshotRequest) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeId

`func (o *CreateSnapshotRequest) SetVolumeId(v string)`

SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


