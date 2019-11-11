# BsuToCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | Set to &#x60;true&#x60; by default, which means that the volume is deleted when the VM is terminated. If set to &#x60;false&#x60;, the volume is not deleted when the VM is terminated. | [optional] 
**Iops** | Pointer to **int32** | The number of I/O operations per second (IOPS). This parameter must be specified only if you create an &#x60;io1&#x60; volume. The maximum number of IOPS allowed for &#x60;io1&#x60; volumes is &#x60;13000&#x60;. | [optional] 
**SnapshotId** | Pointer to **string** | The ID of the snapshot used to create the volume. | [optional] 
**VolumeSize** | Pointer to **int32** | The size of the volume, in gibibytes (GiB).&lt;br /&gt; If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.&lt;br /&gt; If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one. | [optional] 
**VolumeType** | Pointer to **string** | The type of the volume (&#x60;standard&#x60; \\| &#x60;io1&#x60; \\| &#x60;gp2&#x60;). If not specified in the request, a &#x60;standard&#x60; volume is created.&lt;br /&gt; For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS). | [optional] 

## Methods

### GetDeleteOnVmDeletion

`func (o *BsuToCreate) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *BsuToCreate) GetDeleteOnVmDeletionOk() (bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleteOnVmDeletion

`func (o *BsuToCreate) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### SetDeleteOnVmDeletion

`func (o *BsuToCreate) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.

### GetIops

`func (o *BsuToCreate) GetIops() int32`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *BsuToCreate) GetIopsOk() (int32, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIops

`func (o *BsuToCreate) HasIops() bool`

HasIops returns a boolean if a field has been set.

### SetIops

`func (o *BsuToCreate) SetIops(v int32)`

SetIops gets a reference to the given int32 and assigns it to the Iops field.

### GetSnapshotId

`func (o *BsuToCreate) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *BsuToCreate) GetSnapshotIdOk() (string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSnapshotId

`func (o *BsuToCreate) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotId

`func (o *BsuToCreate) SetSnapshotId(v string)`

SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.

### GetVolumeSize

`func (o *BsuToCreate) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *BsuToCreate) GetVolumeSizeOk() (int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeSize

`func (o *BsuToCreate) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.

### SetVolumeSize

`func (o *BsuToCreate) SetVolumeSize(v int32)`

SetVolumeSize gets a reference to the given int32 and assigns it to the VolumeSize field.

### GetVolumeType

`func (o *BsuToCreate) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *BsuToCreate) GetVolumeTypeOk() (string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeType

`func (o *BsuToCreate) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeType

`func (o *BsuToCreate) SetVolumeType(v string)`

SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


