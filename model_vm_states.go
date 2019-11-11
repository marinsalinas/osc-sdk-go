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

// VmStates Information about the states of the VMs.
type VmStates struct {
	// One or more scheduled events associated with the VM.
	MaintenanceEvents *[]MaintenanceEvent `json:"MaintenanceEvents,omitempty"`
	// The name of the Subregion of the VM.
	SubregionName *string `json:"SubregionName,omitempty"`
	// The ID of the VM.
	VmId *string `json:"VmId,omitempty"`
	// The state of the VM (`pending` \\| `running` \\| `shutting-down` \\| `terminated` \\| `stopping` \\| `stopped`).
	VmState *string `json:"VmState,omitempty"`
}

// GetMaintenanceEvents returns the MaintenanceEvents field value if set, zero value otherwise.
func (o *VmStates) GetMaintenanceEvents() []MaintenanceEvent {
	if o == nil || o.MaintenanceEvents == nil {
		var ret []MaintenanceEvent
		return ret
	}
	return *o.MaintenanceEvents
}

// GetMaintenanceEventsOk returns a tuple with the MaintenanceEvents field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *VmStates) GetMaintenanceEventsOk() ([]MaintenanceEvent, bool) {
	if o == nil || o.MaintenanceEvents == nil {
		var ret []MaintenanceEvent
		return ret, false
	}
	return *o.MaintenanceEvents, true
}

// HasMaintenanceEvents returns a boolean if a field has been set.
func (o *VmStates) HasMaintenanceEvents() bool {
	if o != nil && o.MaintenanceEvents != nil {
		return true
	}

	return false
}

// SetMaintenanceEvents gets a reference to the given []MaintenanceEvent and assigns it to the MaintenanceEvents field.
func (o *VmStates) SetMaintenanceEvents(v []MaintenanceEvent) {
	o.MaintenanceEvents = &v
}

// GetSubregionName returns the SubregionName field value if set, zero value otherwise.
func (o *VmStates) GetSubregionName() string {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret
	}
	return *o.SubregionName
}

// GetSubregionNameOk returns a tuple with the SubregionName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *VmStates) GetSubregionNameOk() (string, bool) {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret, false
	}
	return *o.SubregionName, true
}

// HasSubregionName returns a boolean if a field has been set.
func (o *VmStates) HasSubregionName() bool {
	if o != nil && o.SubregionName != nil {
		return true
	}

	return false
}

// SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.
func (o *VmStates) SetSubregionName(v string) {
	o.SubregionName = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *VmStates) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *VmStates) GetVmIdOk() (string, bool) {
	if o == nil || o.VmId == nil {
		var ret string
		return ret, false
	}
	return *o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *VmStates) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *VmStates) SetVmId(v string) {
	o.VmId = &v
}

// GetVmState returns the VmState field value if set, zero value otherwise.
func (o *VmStates) GetVmState() string {
	if o == nil || o.VmState == nil {
		var ret string
		return ret
	}
	return *o.VmState
}

// GetVmStateOk returns a tuple with the VmState field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *VmStates) GetVmStateOk() (string, bool) {
	if o == nil || o.VmState == nil {
		var ret string
		return ret, false
	}
	return *o.VmState, true
}

// HasVmState returns a boolean if a field has been set.
func (o *VmStates) HasVmState() bool {
	if o != nil && o.VmState != nil {
		return true
	}

	return false
}

// SetVmState gets a reference to the given string and assigns it to the VmState field.
func (o *VmStates) SetVmState(v string) {
	o.VmState = &v
}

type NullableVmStates struct {
	Value VmStates
	ExplicitNull bool
}

func (v NullableVmStates) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableVmStates) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
