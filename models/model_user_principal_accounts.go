/*
 * TD Ameritrade API
 *
 * TD Ameritrade API
 *
 * API version: 3.0.1
 * Contact: austin.millan@protonmail.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UserPrincipalAccounts struct for UserPrincipalAccounts
type UserPrincipalAccounts struct {
	AccountCdDomainId *string `json:"accountCdDomainId,omitempty"`
	AccountId *string `json:"accountId,omitempty"`
	Acl *string `json:"acl,omitempty"`
	Authorizations *UserPrincipalAuthorizations `json:"authorizations,omitempty"`
	Company *string `json:"company,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Preferences *UserPrincipalPreferences `json:"preferences,omitempty"`
	Segment *string `json:"segment,omitempty"`
	SurrogateIds *string `json:"surrogateIds,omitempty"`
}

// NewUserPrincipalAccounts instantiates a new UserPrincipalAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPrincipalAccounts() *UserPrincipalAccounts {
	this := UserPrincipalAccounts{}
	return &this
}

// NewUserPrincipalAccountsWithDefaults instantiates a new UserPrincipalAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPrincipalAccountsWithDefaults() *UserPrincipalAccounts {
	this := UserPrincipalAccounts{}
	return &this
}

// GetAccountCdDomainId returns the AccountCdDomainId field value if set, zero value otherwise.
func (o *UserPrincipalAccounts) GetAccountCdDomainId() string {
	if o == nil || o.AccountCdDomainId == nil {
		var ret string
		return ret
	}
	return *o.AccountCdDomainId
}

// GetAccountCdDomainIdOk returns a tuple with the AccountCdDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalAccounts) GetAccountCdDomainIdOk() (*string, bool) {
	if o == nil || o.AccountCdDomainId == nil {
		return nil, false
	}
	return o.AccountCdDomainId, true
}

// HasAccountCdDomainId returns a boolean if a field has been set.
func (o *UserPrincipalAccounts) HasAccountCdDomainId() bool {
	if o != nil && o.AccountCdDomainId != nil {
		return true
	}

	return false
}

// SetAccountCdDomainId gets a reference to the given string and assigns it to the AccountCdDomainId field.
func (o *UserPrincipalAccounts) SetAccountCdDomainId(v string) {
	o.AccountCdDomainId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *UserPrincipalAccounts) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalAccounts) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *UserPrincipalAccounts) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *UserPrincipalAccounts) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *UserPrincipalAccounts) GetAcl() string {
	if o == nil || o.Acl == nil {
		var ret string
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalAccounts) GetAclOk() (*string, bool) {
	if o == nil || o.Acl == nil {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *UserPrincipalAccounts) HasAcl() bool {
	if o != nil && o.Acl != nil {
		return true
	}

	return false
}

// SetAcl gets a reference to the given string and assigns it to the Acl field.
func (o *UserPrincipalAccounts) SetAcl(v string) {
	o.Acl = &v
}

// GetAuthorizations returns the Authorizations field value if set, zero value otherwise.
func (o *UserPrincipalAccounts) GetAuthorizations() UserPrincipalAuthorizations {
	if o == nil || o.Authorizations == nil {
		var ret UserPrincipalAuthorizations
		return ret
	}
	return *o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalAccounts) GetAuthorizationsOk() (*UserPrincipalAuthorizations, bool) {
	if o == nil || o.Authorizations == nil {
		return nil, false
	}
	return o.Authorizations, true
}

// HasAuthorizations returns a boolean if a field has been set.
func (o *UserPrincipalAccounts) HasAuthorizations() bool {
	if o != nil && o.Authorizations != nil {
		return true
	}

	return false
}

// SetAuthorizations gets a reference to the given UserPrincipalAuthorizations and assigns it to the Authorizations field.
func (o *UserPrincipalAccounts) SetAuthorizations(v UserPrincipalAuthorizations) {
	o.Authorizations = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *UserPrincipalAccounts) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalAccounts) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *UserPrincipalAccounts) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *UserPrincipalAccounts) SetCompany(v string) {
	o.Company = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserPrincipalAccounts) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalAccounts) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserPrincipalAccounts) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserPrincipalAccounts) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserPrincipalAccounts) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalAccounts) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserPrincipalAccounts) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserPrincipalAccounts) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPreferences returns the Preferences field value if set, zero value otherwise.
func (o *UserPrincipalAccounts) GetPreferences() UserPrincipalPreferences {
	if o == nil || o.Preferences == nil {
		var ret UserPrincipalPreferences
		return ret
	}
	return *o.Preferences
}

// GetPreferencesOk returns a tuple with the Preferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalAccounts) GetPreferencesOk() (*UserPrincipalPreferences, bool) {
	if o == nil || o.Preferences == nil {
		return nil, false
	}
	return o.Preferences, true
}

// HasPreferences returns a boolean if a field has been set.
func (o *UserPrincipalAccounts) HasPreferences() bool {
	if o != nil && o.Preferences != nil {
		return true
	}

	return false
}

// SetPreferences gets a reference to the given UserPrincipalPreferences and assigns it to the Preferences field.
func (o *UserPrincipalAccounts) SetPreferences(v UserPrincipalPreferences) {
	o.Preferences = &v
}

// GetSegment returns the Segment field value if set, zero value otherwise.
func (o *UserPrincipalAccounts) GetSegment() string {
	if o == nil || o.Segment == nil {
		var ret string
		return ret
	}
	return *o.Segment
}

// GetSegmentOk returns a tuple with the Segment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalAccounts) GetSegmentOk() (*string, bool) {
	if o == nil || o.Segment == nil {
		return nil, false
	}
	return o.Segment, true
}

// HasSegment returns a boolean if a field has been set.
func (o *UserPrincipalAccounts) HasSegment() bool {
	if o != nil && o.Segment != nil {
		return true
	}

	return false
}

// SetSegment gets a reference to the given string and assigns it to the Segment field.
func (o *UserPrincipalAccounts) SetSegment(v string) {
	o.Segment = &v
}

// GetSurrogateIds returns the SurrogateIds field value if set, zero value otherwise.
func (o *UserPrincipalAccounts) GetSurrogateIds() string {
	if o == nil || o.SurrogateIds == nil {
		var ret string
		return ret
	}
	return *o.SurrogateIds
}

// GetSurrogateIdsOk returns a tuple with the SurrogateIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalAccounts) GetSurrogateIdsOk() (*string, bool) {
	if o == nil || o.SurrogateIds == nil {
		return nil, false
	}
	return o.SurrogateIds, true
}

// HasSurrogateIds returns a boolean if a field has been set.
func (o *UserPrincipalAccounts) HasSurrogateIds() bool {
	if o != nil && o.SurrogateIds != nil {
		return true
	}

	return false
}

// SetSurrogateIds gets a reference to the given string and assigns it to the SurrogateIds field.
func (o *UserPrincipalAccounts) SetSurrogateIds(v string) {
	o.SurrogateIds = &v
}

func (o UserPrincipalAccounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountCdDomainId != nil {
		toSerialize["accountCdDomainId"] = o.AccountCdDomainId
	}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.Acl != nil {
		toSerialize["acl"] = o.Acl
	}
	if o.Authorizations != nil {
		toSerialize["authorizations"] = o.Authorizations
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Preferences != nil {
		toSerialize["preferences"] = o.Preferences
	}
	if o.Segment != nil {
		toSerialize["segment"] = o.Segment
	}
	if o.SurrogateIds != nil {
		toSerialize["surrogateIds"] = o.SurrogateIds
	}
	return json.Marshal(toSerialize)
}

type NullableUserPrincipalAccounts struct {
	value *UserPrincipalAccounts
	isSet bool
}

func (v NullableUserPrincipalAccounts) Get() *UserPrincipalAccounts {
	return v.value
}

func (v *NullableUserPrincipalAccounts) Set(val *UserPrincipalAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPrincipalAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPrincipalAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPrincipalAccounts(val *UserPrincipalAccounts) *NullableUserPrincipalAccounts {
	return &NullableUserPrincipalAccounts{value: val, isSet: true}
}

func (v NullableUserPrincipalAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPrincipalAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

