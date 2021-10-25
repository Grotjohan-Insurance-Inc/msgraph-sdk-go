package tentativelyaccept

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

type TentativelyAcceptRequestBody struct {
    additionalData map[string]interface{};
    comment *string;
    proposedNewTime *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TimeSlot;
    sendResponse *bool;
}
func NewTentativelyAcceptRequestBody()(*TentativelyAcceptRequestBody) {
    m := &TentativelyAcceptRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TentativelyAcceptRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TentativelyAcceptRequestBody) GetComment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
func (m *TentativelyAcceptRequestBody) GetProposedNewTime()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TimeSlot) {
    if m == nil {
        return nil
    } else {
        return m.proposedNewTime
    }
}
func (m *TentativelyAcceptRequestBody) GetSendResponse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sendResponse
    }
}
func (m *TentativelyAcceptRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["comment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetComment(val)
        return nil
    }
    res["proposedNewTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewTimeSlot() })
        if err != nil {
            return err
        }
        m.SetProposedNewTime(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TimeSlot))
        return nil
    }
    res["sendResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSendResponse(val)
        return nil
    }
    return res
}
func (m *TentativelyAcceptRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *TentativelyAcceptRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("proposedNewTime", m.GetProposedNewTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendResponse", m.GetSendResponse())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TentativelyAcceptRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TentativelyAcceptRequestBody) SetComment(value *string)() {
    m.comment = value
}
func (m *TentativelyAcceptRequestBody) SetProposedNewTime(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TimeSlot)() {
    m.proposedNewTime = value
}
func (m *TentativelyAcceptRequestBody) SetSendResponse(value *bool)() {
    m.sendResponse = value
}