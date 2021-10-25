package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PermissionGrantPolicy struct {
    PolicyBase
    excludes []PermissionGrantConditionSet;
    includes []PermissionGrantConditionSet;
}
func NewPermissionGrantPolicy()(*PermissionGrantPolicy) {
    m := &PermissionGrantPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
func (m *PermissionGrantPolicy) GetExcludes()([]PermissionGrantConditionSet) {
    if m == nil {
        return nil
    } else {
        return m.excludes
    }
}
func (m *PermissionGrantPolicy) GetIncludes()([]PermissionGrantConditionSet) {
    if m == nil {
        return nil
    } else {
        return m.includes
    }
}
func (m *PermissionGrantPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["excludes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermissionGrantConditionSet() })
        if err != nil {
            return err
        }
        res := make([]PermissionGrantConditionSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*PermissionGrantConditionSet))
        }
        m.SetExcludes(res)
        return nil
    }
    res["includes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermissionGrantConditionSet() })
        if err != nil {
            return err
        }
        res := make([]PermissionGrantConditionSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*PermissionGrantConditionSet))
        }
        m.SetIncludes(res)
        return nil
    }
    return res
}
func (m *PermissionGrantPolicy) IsNil()(bool) {
    return m == nil
}
func (m *PermissionGrantPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExcludes()))
        for i, v := range m.GetExcludes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("excludes", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIncludes()))
        for i, v := range m.GetIncludes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("includes", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PermissionGrantPolicy) SetExcludes(value []PermissionGrantConditionSet)() {
    m.excludes = value
}
func (m *PermissionGrantPolicy) SetIncludes(value []PermissionGrantConditionSet)() {
    m.includes = value
}