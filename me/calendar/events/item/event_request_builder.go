package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i083a9367fb14f0da75caea7bc857922fe10d91693dde0f11517ab6420e557018 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/attachments"
    i0f781326a67fbc9776be4d1ccbf43639b9bf8077faa925a7b0f90f0e03141c07 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/snoozereminder"
    i10a043ae8449d6b8061165bb67946e604bc2614783ade4a4557a2bb8d2a5a2b0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/decline"
    i339d2d1a30983a01dd7de00efdd9e7d51cf0424dd9c4c96d335f9901fcff9763 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances"
    i49ad5e4dbfe1e924deb723580385a06e87f447b9ed73e648b64a934117f1c784 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/forward"
    i62d05692265cf61165d6bb2a1bb3a36d03ea2604586e3ebd550fae433333488e "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/singlevalueextendedproperties"
    i69bb0e8065fcc365cd51c4f9cb100c9289f8c378fe585dd0126d7eaee8953622 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/calendar"
    i87334fbab17098d26d733b7c6d4d979d97f28091d5d1b3c41c2e4adcd9d2f2a3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/accept"
    id1d8fbd767f32fa77bb0a1e1f8b79c51f1a58b7dae1ccebcd3a3001af55d4f44 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/extensions"
    ie04ca0f3f9823daf845f8b3b165c34feaffc35ed4189263a607834108a606dbe "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/cancel"
    ie6641980e1225311bdf1c499d054bc5d07aa799f49902ef90e71b01800c2a555 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/dismissreminder"
    iee14c2abd2614ebab15ceaf1135456d72c4e1cd953dd5d7b3347f819c29c6002 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/tentativelyaccept"
    if78feb360a058ff87bf5a02baba6f7d4fb28a9e5f93fc081392b457ad6fae3eb "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/multivalueextendedproperties"
    i3f305339dd659581dddbb6081f2d1981394ba5f409b3fff8d5affc26648441c4 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/singlevalueextendedproperties/item"
    i658437a8657f51783d49c87367e2ae5b495c006ac48f6a7010d75098566455a0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/multivalueextendedproperties/item"
    i868bcbc69b2c3a35a368f852b47fffbb78632a8f15944de468f0e130a0242300 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/attachments/item"
    i955a97236a9f337897b15ed49defa3aa5c5d7910ae15f5b5efe5fbae5e75cb05 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/extensions/item"
    ic8bb7a27f91fd567d4417f19bf0aacb6595974ef7d4a830d5f1842796aab82f7 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item"
)

type EventRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *EventRequestBuilder) Accept()(*i87334fbab17098d26d733b7c6d4d979d97f28091d5d1b3c41c2e4adcd9d2f2a3.AcceptRequestBuilder) {
    return i87334fbab17098d26d733b7c6d4d979d97f28091d5d1b3c41c2e4adcd9d2f2a3.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i083a9367fb14f0da75caea7bc857922fe10d91693dde0f11517ab6420e557018.AttachmentsRequestBuilder) {
    return i083a9367fb14f0da75caea7bc857922fe10d91693dde0f11517ab6420e557018.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) AttachmentsById(id string)(*i868bcbc69b2c3a35a368f852b47fffbb78632a8f15944de468f0e130a0242300.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i868bcbc69b2c3a35a368f852b47fffbb78632a8f15944de468f0e130a0242300.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i69bb0e8065fcc365cd51c4f9cb100c9289f8c378fe585dd0126d7eaee8953622.CalendarRequestBuilder) {
    return i69bb0e8065fcc365cd51c4f9cb100c9289f8c378fe585dd0126d7eaee8953622.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*ie04ca0f3f9823daf845f8b3b165c34feaffc35ed4189263a607834108a606dbe.CancelRequestBuilder) {
    return ie04ca0f3f9823daf845f8b3b165c34feaffc35ed4189263a607834108a606dbe.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/calendar/events/{event_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EventRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EventRequestBuilder) CreateGetRequestInformation(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EventRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EventRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EventRequestBuilder) Decline()(*i10a043ae8449d6b8061165bb67946e604bc2614783ade4a4557a2bb8d2a5a2b0.DeclineRequestBuilder) {
    return i10a043ae8449d6b8061165bb67946e604bc2614783ade4a4557a2bb8d2a5a2b0.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) DismissReminder()(*ie6641980e1225311bdf1c499d054bc5d07aa799f49902ef90e71b01800c2a555.DismissReminderRequestBuilder) {
    return ie6641980e1225311bdf1c499d054bc5d07aa799f49902ef90e71b01800c2a555.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*id1d8fbd767f32fa77bb0a1e1f8b79c51f1a58b7dae1ccebcd3a3001af55d4f44.ExtensionsRequestBuilder) {
    return id1d8fbd767f32fa77bb0a1e1f8b79c51f1a58b7dae1ccebcd3a3001af55d4f44.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExtensionsById(id string)(*i955a97236a9f337897b15ed49defa3aa5c5d7910ae15f5b5efe5fbae5e75cb05.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i955a97236a9f337897b15ed49defa3aa5c5d7910ae15f5b5efe5fbae5e75cb05.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i49ad5e4dbfe1e924deb723580385a06e87f447b9ed73e648b64a934117f1c784.ForwardRequestBuilder) {
    return i49ad5e4dbfe1e924deb723580385a06e87f447b9ed73e648b64a934117f1c784.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Get(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEvent() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event), nil
}
func (m *EventRequestBuilder) Instances()(*i339d2d1a30983a01dd7de00efdd9e7d51cf0424dd9c4c96d335f9901fcff9763.InstancesRequestBuilder) {
    return i339d2d1a30983a01dd7de00efdd9e7d51cf0424dd9c4c96d335f9901fcff9763.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) InstancesById(id string)(*ic8bb7a27f91fd567d4417f19bf0aacb6595974ef7d4a830d5f1842796aab82f7.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ic8bb7a27f91fd567d4417f19bf0aacb6595974ef7d4a830d5f1842796aab82f7.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*if78feb360a058ff87bf5a02baba6f7d4fb28a9e5f93fc081392b457ad6fae3eb.MultiValueExtendedPropertiesRequestBuilder) {
    return if78feb360a058ff87bf5a02baba6f7d4fb28a9e5f93fc081392b457ad6fae3eb.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i658437a8657f51783d49c87367e2ae5b495c006ac48f6a7010d75098566455a0.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i658437a8657f51783d49c87367e2ae5b495c006ac48f6a7010d75098566455a0.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i62d05692265cf61165d6bb2a1bb3a36d03ea2604586e3ebd550fae433333488e.SingleValueExtendedPropertiesRequestBuilder) {
    return i62d05692265cf61165d6bb2a1bb3a36d03ea2604586e3ebd550fae433333488e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i3f305339dd659581dddbb6081f2d1981394ba5f409b3fff8d5affc26648441c4.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i3f305339dd659581dddbb6081f2d1981394ba5f409b3fff8d5affc26648441c4.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i0f781326a67fbc9776be4d1ccbf43639b9bf8077faa925a7b0f90f0e03141c07.SnoozeReminderRequestBuilder) {
    return i0f781326a67fbc9776be4d1ccbf43639b9bf8077faa925a7b0f90f0e03141c07.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*iee14c2abd2614ebab15ceaf1135456d72c4e1cd953dd5d7b3347f819c29c6002.TentativelyAcceptRequestBuilder) {
    return iee14c2abd2614ebab15ceaf1135456d72c4e1cd953dd5d7b3347f819c29c6002.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}