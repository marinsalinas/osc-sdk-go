# DeleteExportTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ExportTaskId** | Pointer to **string** | The ID of the export task to delete. | 

## Methods

### GetDryRun

`func (o *DeleteExportTaskRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteExportTaskRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteExportTaskRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteExportTaskRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetExportTaskId

`func (o *DeleteExportTaskRequest) GetExportTaskId() string`

GetExportTaskId returns the ExportTaskId field if non-nil, zero value otherwise.

### GetExportTaskIdOk

`func (o *DeleteExportTaskRequest) GetExportTaskIdOk() (string, bool)`

GetExportTaskIdOk returns a tuple with the ExportTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExportTaskId

`func (o *DeleteExportTaskRequest) HasExportTaskId() bool`

HasExportTaskId returns a boolean if a field has been set.

### SetExportTaskId

`func (o *DeleteExportTaskRequest) SetExportTaskId(v string)`

SetExportTaskId gets a reference to the given string and assigns it to the ExportTaskId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


