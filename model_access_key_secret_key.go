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

// AccessKeySecretKey Information about the secret access key.
type AccessKeySecretKey struct {
	// The ID of the secret access key.
	AccessKeyId *string `json:"AccessKeyId,omitempty"`
	// The date and time of creation of the secret access key.
	CreationDate *string `json:"CreationDate,omitempty"`
	// The date and time of the last modification of the secret access key.
	LastModificationDate *string `json:"LastModificationDate,omitempty"`
	// The secret access key that enables you to send requests.
	SecretKey *string `json:"SecretKey,omitempty"`
	// The state of the secret access key (`Active` if the key is valid for API calls, or `Inactive` if not).
	State *string `json:"State,omitempty"`
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *AccessKeySecretKey) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeySecretKey) GetAccessKeyIdOk() (string, bool) {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret, false
	}
	return *o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *AccessKeySecretKey) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId != nil {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *AccessKeySecretKey) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *AccessKeySecretKey) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeySecretKey) GetCreationDateOk() (string, bool) {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret, false
	}
	return *o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *AccessKeySecretKey) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *AccessKeySecretKey) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetLastModificationDate returns the LastModificationDate field value if set, zero value otherwise.
func (o *AccessKeySecretKey) GetLastModificationDate() string {
	if o == nil || o.LastModificationDate == nil {
		var ret string
		return ret
	}
	return *o.LastModificationDate
}

// GetLastModificationDateOk returns a tuple with the LastModificationDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeySecretKey) GetLastModificationDateOk() (string, bool) {
	if o == nil || o.LastModificationDate == nil {
		var ret string
		return ret, false
	}
	return *o.LastModificationDate, true
}

// HasLastModificationDate returns a boolean if a field has been set.
func (o *AccessKeySecretKey) HasLastModificationDate() bool {
	if o != nil && o.LastModificationDate != nil {
		return true
	}

	return false
}

// SetLastModificationDate gets a reference to the given string and assigns it to the LastModificationDate field.
func (o *AccessKeySecretKey) SetLastModificationDate(v string) {
	o.LastModificationDate = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *AccessKeySecretKey) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeySecretKey) GetSecretKeyOk() (string, bool) {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret, false
	}
	return *o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *AccessKeySecretKey) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *AccessKeySecretKey) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AccessKeySecretKey) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeySecretKey) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AccessKeySecretKey) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AccessKeySecretKey) SetState(v string) {
	o.State = &v
}

type NullableAccessKeySecretKey struct {
	Value        AccessKeySecretKey
	ExplicitNull bool
}

func (v NullableAccessKeySecretKey) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAccessKeySecretKey) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}