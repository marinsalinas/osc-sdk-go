# ReadFlexibleGpusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersFlexibleGpu**](FiltersFlexibleGpu.md) |  | [optional] 

## Methods

### GetDryRun

`func (o *ReadFlexibleGpusRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadFlexibleGpusRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *ReadFlexibleGpusRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *ReadFlexibleGpusRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetFilters

`func (o *ReadFlexibleGpusRequest) GetFilters() FiltersFlexibleGpu`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadFlexibleGpusRequest) GetFiltersOk() (FiltersFlexibleGpu, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilters

`func (o *ReadFlexibleGpusRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFilters

`func (o *ReadFlexibleGpusRequest) SetFilters(v FiltersFlexibleGpu)`

SetFilters gets a reference to the given FiltersFlexibleGpu and assigns it to the Filters field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


