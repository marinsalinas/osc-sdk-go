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

// CreateNetRequest struct for CreateNetRequest
type CreateNetRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).
	IpRange string `json:"IpRange"`
	// The tenancy options for the VMs (`default` if a VM created in a Net can be launched with any tenancy, `dedicated` if it can be launched with dedicated tenancy VMs running on single-tenant hardware).
	Tenancy *string `json:"Tenancy,omitempty"`
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateNetRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateNetRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateNetRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetIpRange returns the IpRange field value
func (o *CreateNetRequest) GetIpRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpRange
}

// SetIpRange sets field value
func (o *CreateNetRequest) SetIpRange(v string) {
	o.IpRange = v
}

// GetTenancy returns the Tenancy field value if set, zero value otherwise.
func (o *CreateNetRequest) GetTenancy() string {
	if o == nil || o.Tenancy == nil {
		var ret string
		return ret
	}
	return *o.Tenancy
}

// GetTenancyOk returns a tuple with the Tenancy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetRequest) GetTenancyOk() (string, bool) {
	if o == nil || o.Tenancy == nil {
		var ret string
		return ret, false
	}
	return *o.Tenancy, true
}

// HasTenancy returns a boolean if a field has been set.
func (o *CreateNetRequest) HasTenancy() bool {
	if o != nil && o.Tenancy != nil {
		return true
	}

	return false
}

// SetTenancy gets a reference to the given string and assigns it to the Tenancy field.
func (o *CreateNetRequest) SetTenancy(v string) {
	o.Tenancy = &v
}

type NullableCreateNetRequest struct {
	Value CreateNetRequest
	ExplicitNull bool
}

func (v NullableCreateNetRequest) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableCreateNetRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
