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

// UpdateAccountResponse struct for UpdateAccountResponse
type UpdateAccountResponse struct {
	Account         *Account         `json:"Account,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *UpdateAccountResponse) GetAccount() Account {
	if o == nil || o.Account == nil {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountResponse) GetAccountOk() (Account, bool) {
	if o == nil || o.Account == nil {
		var ret Account
		return ret, false
	}
	return *o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *UpdateAccountResponse) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *UpdateAccountResponse) SetAccount(v Account) {
	o.Account = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateAccountResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateAccountResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateAccountResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableUpdateAccountResponse struct {
	Value        UpdateAccountResponse
	ExplicitNull bool
}

func (v NullableUpdateAccountResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateAccountResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}