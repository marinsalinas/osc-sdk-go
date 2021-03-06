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

// CreateNicResponse struct for CreateNicResponse
type CreateNicResponse struct {
	Nic             *Nic             `json:"Nic,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetNic returns the Nic field value if set, zero value otherwise.
func (o *CreateNicResponse) GetNic() Nic {
	if o == nil || o.Nic == nil {
		var ret Nic
		return ret
	}
	return *o.Nic
}

// GetNicOk returns a tuple with the Nic field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateNicResponse) GetNicOk() (Nic, bool) {
	if o == nil || o.Nic == nil {
		var ret Nic
		return ret, false
	}
	return *o.Nic, true
}

// HasNic returns a boolean if a field has been set.
func (o *CreateNicResponse) HasNic() bool {
	if o != nil && o.Nic != nil {
		return true
	}

	return false
}

// SetNic gets a reference to the given Nic and assigns it to the Nic field.
func (o *CreateNicResponse) SetNic(v Nic) {
	o.Nic = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateNicResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateNicResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateNicResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateNicResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableCreateNicResponse struct {
	Value        CreateNicResponse
	ExplicitNull bool
}

func (v NullableCreateNicResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateNicResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
