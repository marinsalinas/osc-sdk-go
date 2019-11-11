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

// PublicIpLight Information about the public IP.
type PublicIpLight struct {
	// The External IP address (EIP) associated with the NAT service.
	PublicIp *string `json:"PublicIp,omitempty"`
	// The allocation ID of the EIP associated with the NAT service.
	PublicIpId *string `json:"PublicIpId,omitempty"`
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *PublicIpLight) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PublicIpLight) GetPublicIpOk() (string, bool) {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret, false
	}
	return *o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *PublicIpLight) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *PublicIpLight) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPublicIpId returns the PublicIpId field value if set, zero value otherwise.
func (o *PublicIpLight) GetPublicIpId() string {
	if o == nil || o.PublicIpId == nil {
		var ret string
		return ret
	}
	return *o.PublicIpId
}

// GetPublicIpIdOk returns a tuple with the PublicIpId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PublicIpLight) GetPublicIpIdOk() (string, bool) {
	if o == nil || o.PublicIpId == nil {
		var ret string
		return ret, false
	}
	return *o.PublicIpId, true
}

// HasPublicIpId returns a boolean if a field has been set.
func (o *PublicIpLight) HasPublicIpId() bool {
	if o != nil && o.PublicIpId != nil {
		return true
	}

	return false
}

// SetPublicIpId gets a reference to the given string and assigns it to the PublicIpId field.
func (o *PublicIpLight) SetPublicIpId(v string) {
	o.PublicIpId = &v
}

type NullablePublicIpLight struct {
	Value PublicIpLight
	ExplicitNull bool
}

func (v NullablePublicIpLight) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullablePublicIpLight) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
