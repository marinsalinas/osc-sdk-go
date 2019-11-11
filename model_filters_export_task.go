/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"bytes"
	"encoding/json"
)

// FiltersExportTask One or more filters.
type FiltersExportTask struct {
	// The keys of the tags associated with the export tasks.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the export tasks.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the export tasks, in the following format: \"Filters\":{\"Tags\":[\"TAGKEY=TAGVALUE\"]}.
	Tags *[]string `json:"Tags,omitempty"`
	// The IDs of the export tasks.
	TaskIds *[]string `json:"TaskIds,omitempty"`
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersExportTask) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersExportTask) GetTagKeysOk() ([]string, bool) {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret, false
	}
	return *o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersExportTask) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersExportTask) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersExportTask) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersExportTask) GetTagValuesOk() ([]string, bool) {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret, false
	}
	return *o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersExportTask) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersExportTask) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersExportTask) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersExportTask) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersExportTask) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersExportTask) SetTags(v []string) {
	o.Tags = &v
}

// GetTaskIds returns the TaskIds field value if set, zero value otherwise.
func (o *FiltersExportTask) GetTaskIds() []string {
	if o == nil || o.TaskIds == nil {
		var ret []string
		return ret
	}
	return *o.TaskIds
}

// GetTaskIdsOk returns a tuple with the TaskIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersExportTask) GetTaskIdsOk() ([]string, bool) {
	if o == nil || o.TaskIds == nil {
		var ret []string
		return ret, false
	}
	return *o.TaskIds, true
}

// HasTaskIds returns a boolean if a field has been set.
func (o *FiltersExportTask) HasTaskIds() bool {
	if o != nil && o.TaskIds != nil {
		return true
	}

	return false
}

// SetTaskIds gets a reference to the given []string and assigns it to the TaskIds field.
func (o *FiltersExportTask) SetTaskIds(v []string) {
	o.TaskIds = &v
}

type NullableFiltersExportTask struct {
	Value FiltersExportTask
	ExplicitNull bool
}

func (v NullableFiltersExportTask) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableFiltersExportTask) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
