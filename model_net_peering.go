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

// NetPeering Information about the Net peering connection.
type NetPeering struct {
	AccepterNet *AccepterNet `json:"AccepterNet,omitempty"`
	// The ID of the Net peering connection.
	NetPeeringId *string          `json:"NetPeeringId,omitempty"`
	SourceNet    *SourceNet       `json:"SourceNet,omitempty"`
	State        *NetPeeringState `json:"State,omitempty"`
	// One or more tags associated with the Net peering connection.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// GetAccepterNet returns the AccepterNet field value if set, zero value otherwise.
func (o *NetPeering) GetAccepterNet() AccepterNet {
	if o == nil || o.AccepterNet == nil {
		var ret AccepterNet
		return ret
	}
	return *o.AccepterNet
}

// GetAccepterNetOk returns a tuple with the AccepterNet field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NetPeering) GetAccepterNetOk() (AccepterNet, bool) {
	if o == nil || o.AccepterNet == nil {
		var ret AccepterNet
		return ret, false
	}
	return *o.AccepterNet, true
}

// HasAccepterNet returns a boolean if a field has been set.
func (o *NetPeering) HasAccepterNet() bool {
	if o != nil && o.AccepterNet != nil {
		return true
	}

	return false
}

// SetAccepterNet gets a reference to the given AccepterNet and assigns it to the AccepterNet field.
func (o *NetPeering) SetAccepterNet(v AccepterNet) {
	o.AccepterNet = &v
}

// GetNetPeeringId returns the NetPeeringId field value if set, zero value otherwise.
func (o *NetPeering) GetNetPeeringId() string {
	if o == nil || o.NetPeeringId == nil {
		var ret string
		return ret
	}
	return *o.NetPeeringId
}

// GetNetPeeringIdOk returns a tuple with the NetPeeringId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NetPeering) GetNetPeeringIdOk() (string, bool) {
	if o == nil || o.NetPeeringId == nil {
		var ret string
		return ret, false
	}
	return *o.NetPeeringId, true
}

// HasNetPeeringId returns a boolean if a field has been set.
func (o *NetPeering) HasNetPeeringId() bool {
	if o != nil && o.NetPeeringId != nil {
		return true
	}

	return false
}

// SetNetPeeringId gets a reference to the given string and assigns it to the NetPeeringId field.
func (o *NetPeering) SetNetPeeringId(v string) {
	o.NetPeeringId = &v
}

// GetSourceNet returns the SourceNet field value if set, zero value otherwise.
func (o *NetPeering) GetSourceNet() SourceNet {
	if o == nil || o.SourceNet == nil {
		var ret SourceNet
		return ret
	}
	return *o.SourceNet
}

// GetSourceNetOk returns a tuple with the SourceNet field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NetPeering) GetSourceNetOk() (SourceNet, bool) {
	if o == nil || o.SourceNet == nil {
		var ret SourceNet
		return ret, false
	}
	return *o.SourceNet, true
}

// HasSourceNet returns a boolean if a field has been set.
func (o *NetPeering) HasSourceNet() bool {
	if o != nil && o.SourceNet != nil {
		return true
	}

	return false
}

// SetSourceNet gets a reference to the given SourceNet and assigns it to the SourceNet field.
func (o *NetPeering) SetSourceNet(v SourceNet) {
	o.SourceNet = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NetPeering) GetState() NetPeeringState {
	if o == nil || o.State == nil {
		var ret NetPeeringState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NetPeering) GetStateOk() (NetPeeringState, bool) {
	if o == nil || o.State == nil {
		var ret NetPeeringState
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NetPeering) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given NetPeeringState and assigns it to the State field.
func (o *NetPeering) SetState(v NetPeeringState) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *NetPeering) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NetPeering) GetTagsOk() ([]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NetPeering) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *NetPeering) SetTags(v []ResourceTag) {
	o.Tags = &v
}

type NullableNetPeering struct {
	Value        NetPeering
	ExplicitNull bool
}

func (v NullableNetPeering) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNetPeering) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
