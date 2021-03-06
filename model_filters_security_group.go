/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.1
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// FiltersSecurityGroup One or more filters.
type FiltersSecurityGroup struct {
	// The account IDs of the owners of the security groups.
	AccountIds *[]string `json:"AccountIds,omitempty"`
	// The IDs of the security groups.
	SecurityGroupIds *[]string `json:"SecurityGroupIds,omitempty"`
	// The names of the security groups.
	SecurityGroupNames *[]string `json:"SecurityGroupNames,omitempty"`
	// The keys of the tags associated with the security groups.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the security groups.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the security groups, in the following format: \"Filters\":{\"Tags\":[\"TAGKEY=TAGVALUE\"]}.
	Tags *[]string `json:"Tags,omitempty"`
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetAccountIdsOk() ([]string, bool) {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret, false
	}
	return *o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *FiltersSecurityGroup) SetAccountIds(v []string) {
	o.AccountIds = &v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetSecurityGroupIds() []string {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetSecurityGroupIdsOk() ([]string, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret, false
	}
	return *o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *FiltersSecurityGroup) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = &v
}

// GetSecurityGroupNames returns the SecurityGroupNames field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetSecurityGroupNames() []string {
	if o == nil || o.SecurityGroupNames == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupNames
}

// GetSecurityGroupNamesOk returns a tuple with the SecurityGroupNames field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetSecurityGroupNamesOk() ([]string, bool) {
	if o == nil || o.SecurityGroupNames == nil {
		var ret []string
		return ret, false
	}
	return *o.SecurityGroupNames, true
}

// HasSecurityGroupNames returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasSecurityGroupNames() bool {
	if o != nil && o.SecurityGroupNames != nil {
		return true
	}

	return false
}

// SetSecurityGroupNames gets a reference to the given []string and assigns it to the SecurityGroupNames field.
func (o *FiltersSecurityGroup) SetSecurityGroupNames(v []string) {
	o.SecurityGroupNames = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetTagKeysOk() ([]string, bool) {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret, false
	}
	return *o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersSecurityGroup) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetTagValuesOk() ([]string, bool) {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret, false
	}
	return *o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersSecurityGroup) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersSecurityGroup) SetTags(v []string) {
	o.Tags = &v
}

type NullableFiltersSecurityGroup struct {
	Value        FiltersSecurityGroup
	ExplicitNull bool
}

func (v NullableFiltersSecurityGroup) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFiltersSecurityGroup) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
