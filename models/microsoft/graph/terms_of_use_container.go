package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TermsOfUseContainer struct {
    Entity
    agreementAcceptances []AgreementAcceptance;
    agreements []Agreement;
}
func NewTermsOfUseContainer()(*TermsOfUseContainer) {
    m := &TermsOfUseContainer{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TermsOfUseContainer) GetAgreementAcceptances()([]AgreementAcceptance) {
    if m == nil {
        return nil
    } else {
        return m.agreementAcceptances
    }
}
func (m *TermsOfUseContainer) GetAgreements()([]Agreement) {
    if m == nil {
        return nil
    } else {
        return m.agreements
    }
}
func (m *TermsOfUseContainer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agreementAcceptances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementAcceptance() })
        if err != nil {
            return err
        }
        res := make([]AgreementAcceptance, len(val))
        for i, v := range val {
            res[i] = *(v.(*AgreementAcceptance))
        }
        m.SetAgreementAcceptances(res)
        return nil
    }
    res["agreements"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreement() })
        if err != nil {
            return err
        }
        res := make([]Agreement, len(val))
        for i, v := range val {
            res[i] = *(v.(*Agreement))
        }
        m.SetAgreements(res)
        return nil
    }
    return res
}
func (m *TermsOfUseContainer) IsNil()(bool) {
    return m == nil
}
func (m *TermsOfUseContainer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgreementAcceptances()))
        for i, v := range m.GetAgreementAcceptances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("agreementAcceptances", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgreements()))
        for i, v := range m.GetAgreements() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("agreements", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TermsOfUseContainer) SetAgreementAcceptances(value []AgreementAcceptance)() {
    m.agreementAcceptances = value
}
func (m *TermsOfUseContainer) SetAgreements(value []Agreement)() {
    m.agreements = value
}