package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PolicyRoot struct {
    Entity
    activityBasedTimeoutPolicies []ActivityBasedTimeoutPolicy;
    adminConsentRequestPolicy *AdminConsentRequestPolicy;
    authenticationFlowsPolicy *AuthenticationFlowsPolicy;
    authenticationMethodsPolicy *AuthenticationMethodsPolicy;
    authorizationPolicy *AuthorizationPolicy;
    claimsMappingPolicies []ClaimsMappingPolicy;
    conditionalAccessPolicies []ConditionalAccessPolicy;
    featureRolloutPolicies []FeatureRolloutPolicy;
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicy;
    identitySecurityDefaultsEnforcementPolicy *IdentitySecurityDefaultsEnforcementPolicy;
    permissionGrantPolicies []PermissionGrantPolicy;
    tokenIssuancePolicies []TokenIssuancePolicy;
    tokenLifetimePolicies []TokenLifetimePolicy;
}
func NewPolicyRoot()(*PolicyRoot) {
    m := &PolicyRoot{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PolicyRoot) GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicy) {
    if m == nil {
        return nil
    } else {
        return m.activityBasedTimeoutPolicies
    }
}
func (m *PolicyRoot) GetAdminConsentRequestPolicy()(*AdminConsentRequestPolicy) {
    if m == nil {
        return nil
    } else {
        return m.adminConsentRequestPolicy
    }
}
func (m *PolicyRoot) GetAuthenticationFlowsPolicy()(*AuthenticationFlowsPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authenticationFlowsPolicy
    }
}
func (m *PolicyRoot) GetAuthenticationMethodsPolicy()(*AuthenticationMethodsPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodsPolicy
    }
}
func (m *PolicyRoot) GetAuthorizationPolicy()(*AuthorizationPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authorizationPolicy
    }
}
func (m *PolicyRoot) GetClaimsMappingPolicies()([]ClaimsMappingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.claimsMappingPolicies
    }
}
func (m *PolicyRoot) GetConditionalAccessPolicies()([]ConditionalAccessPolicy) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessPolicies
    }
}
func (m *PolicyRoot) GetFeatureRolloutPolicies()([]FeatureRolloutPolicy) {
    if m == nil {
        return nil
    } else {
        return m.featureRolloutPolicies
    }
}
func (m *PolicyRoot) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicy) {
    if m == nil {
        return nil
    } else {
        return m.homeRealmDiscoveryPolicies
    }
}
func (m *PolicyRoot) GetIdentitySecurityDefaultsEnforcementPolicy()(*IdentitySecurityDefaultsEnforcementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.identitySecurityDefaultsEnforcementPolicy
    }
}
func (m *PolicyRoot) GetPermissionGrantPolicies()([]PermissionGrantPolicy) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrantPolicies
    }
}
func (m *PolicyRoot) GetTokenIssuancePolicies()([]TokenIssuancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuancePolicies
    }
}
func (m *PolicyRoot) GetTokenLifetimePolicies()([]TokenLifetimePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenLifetimePolicies
    }
}
func (m *PolicyRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityBasedTimeoutPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewActivityBasedTimeoutPolicy() })
        if err != nil {
            return err
        }
        res := make([]ActivityBasedTimeoutPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ActivityBasedTimeoutPolicy))
        }
        m.SetActivityBasedTimeoutPolicies(res)
        return nil
    }
    res["adminConsentRequestPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdminConsentRequestPolicy() })
        if err != nil {
            return err
        }
        m.SetAdminConsentRequestPolicy(val.(*AdminConsentRequestPolicy))
        return nil
    }
    res["authenticationFlowsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationFlowsPolicy() })
        if err != nil {
            return err
        }
        m.SetAuthenticationFlowsPolicy(val.(*AuthenticationFlowsPolicy))
        return nil
    }
    res["authenticationMethodsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethodsPolicy() })
        if err != nil {
            return err
        }
        m.SetAuthenticationMethodsPolicy(val.(*AuthenticationMethodsPolicy))
        return nil
    }
    res["authorizationPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthorizationPolicy() })
        if err != nil {
            return err
        }
        m.SetAuthorizationPolicy(val.(*AuthorizationPolicy))
        return nil
    }
    res["claimsMappingPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClaimsMappingPolicy() })
        if err != nil {
            return err
        }
        res := make([]ClaimsMappingPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ClaimsMappingPolicy))
        }
        m.SetClaimsMappingPolicies(res)
        return nil
    }
    res["conditionalAccessPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessPolicy() })
        if err != nil {
            return err
        }
        res := make([]ConditionalAccessPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConditionalAccessPolicy))
        }
        m.SetConditionalAccessPolicies(res)
        return nil
    }
    res["featureRolloutPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFeatureRolloutPolicy() })
        if err != nil {
            return err
        }
        res := make([]FeatureRolloutPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*FeatureRolloutPolicy))
        }
        m.SetFeatureRolloutPolicies(res)
        return nil
    }
    res["homeRealmDiscoveryPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHomeRealmDiscoveryPolicy() })
        if err != nil {
            return err
        }
        res := make([]HomeRealmDiscoveryPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*HomeRealmDiscoveryPolicy))
        }
        m.SetHomeRealmDiscoveryPolicies(res)
        return nil
    }
    res["identitySecurityDefaultsEnforcementPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySecurityDefaultsEnforcementPolicy() })
        if err != nil {
            return err
        }
        m.SetIdentitySecurityDefaultsEnforcementPolicy(val.(*IdentitySecurityDefaultsEnforcementPolicy))
        return nil
    }
    res["permissionGrantPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermissionGrantPolicy() })
        if err != nil {
            return err
        }
        res := make([]PermissionGrantPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*PermissionGrantPolicy))
        }
        m.SetPermissionGrantPolicies(res)
        return nil
    }
    res["tokenIssuancePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTokenIssuancePolicy() })
        if err != nil {
            return err
        }
        res := make([]TokenIssuancePolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*TokenIssuancePolicy))
        }
        m.SetTokenIssuancePolicies(res)
        return nil
    }
    res["tokenLifetimePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTokenLifetimePolicy() })
        if err != nil {
            return err
        }
        res := make([]TokenLifetimePolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*TokenLifetimePolicy))
        }
        m.SetTokenLifetimePolicies(res)
        return nil
    }
    return res
}
func (m *PolicyRoot) IsNil()(bool) {
    return m == nil
}
func (m *PolicyRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetActivityBasedTimeoutPolicies()))
        for i, v := range m.GetActivityBasedTimeoutPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("activityBasedTimeoutPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("adminConsentRequestPolicy", m.GetAdminConsentRequestPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationFlowsPolicy", m.GetAuthenticationFlowsPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationMethodsPolicy", m.GetAuthenticationMethodsPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authorizationPolicy", m.GetAuthorizationPolicy())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClaimsMappingPolicies()))
        for i, v := range m.GetClaimsMappingPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("claimsMappingPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConditionalAccessPolicies()))
        for i, v := range m.GetConditionalAccessPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("conditionalAccessPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFeatureRolloutPolicies()))
        for i, v := range m.GetFeatureRolloutPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("featureRolloutPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHomeRealmDiscoveryPolicies()))
        for i, v := range m.GetHomeRealmDiscoveryPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("homeRealmDiscoveryPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identitySecurityDefaultsEnforcementPolicy", m.GetIdentitySecurityDefaultsEnforcementPolicy())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPermissionGrantPolicies()))
        for i, v := range m.GetPermissionGrantPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("permissionGrantPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTokenIssuancePolicies()))
        for i, v := range m.GetTokenIssuancePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tokenIssuancePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTokenLifetimePolicies()))
        for i, v := range m.GetTokenLifetimePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tokenLifetimePolicies", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PolicyRoot) SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicy)() {
    m.activityBasedTimeoutPolicies = value
}
func (m *PolicyRoot) SetAdminConsentRequestPolicy(value *AdminConsentRequestPolicy)() {
    m.adminConsentRequestPolicy = value
}
func (m *PolicyRoot) SetAuthenticationFlowsPolicy(value *AuthenticationFlowsPolicy)() {
    m.authenticationFlowsPolicy = value
}
func (m *PolicyRoot) SetAuthenticationMethodsPolicy(value *AuthenticationMethodsPolicy)() {
    m.authenticationMethodsPolicy = value
}
func (m *PolicyRoot) SetAuthorizationPolicy(value *AuthorizationPolicy)() {
    m.authorizationPolicy = value
}
func (m *PolicyRoot) SetClaimsMappingPolicies(value []ClaimsMappingPolicy)() {
    m.claimsMappingPolicies = value
}
func (m *PolicyRoot) SetConditionalAccessPolicies(value []ConditionalAccessPolicy)() {
    m.conditionalAccessPolicies = value
}
func (m *PolicyRoot) SetFeatureRolloutPolicies(value []FeatureRolloutPolicy)() {
    m.featureRolloutPolicies = value
}
func (m *PolicyRoot) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicy)() {
    m.homeRealmDiscoveryPolicies = value
}
func (m *PolicyRoot) SetIdentitySecurityDefaultsEnforcementPolicy(value *IdentitySecurityDefaultsEnforcementPolicy)() {
    m.identitySecurityDefaultsEnforcementPolicy = value
}
func (m *PolicyRoot) SetPermissionGrantPolicies(value []PermissionGrantPolicy)() {
    m.permissionGrantPolicies = value
}
func (m *PolicyRoot) SetTokenIssuancePolicies(value []TokenIssuancePolicy)() {
    m.tokenIssuancePolicies = value
}
func (m *PolicyRoot) SetTokenLifetimePolicies(value []TokenLifetimePolicy)() {
    m.tokenLifetimePolicies = value
}