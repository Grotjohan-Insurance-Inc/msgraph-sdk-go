package taughtclasses

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i9a82ba081a87a0b82fe7518f2699a91f6cdabe7560d9041a4bbd68ca17595c04 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/taughtclasses/delta"
    ibfebe8e1821c93dda8997c543dd7161494427c674556776df0eee3b6720e6f3c "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/taughtclasses/ref"
)

type TaughtClassesRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type TaughtClassesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escpaped []string;
    Skip *int32;
    Top *int32;
}
func NewTaughtClassesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TaughtClassesRequestBuilder) {
    m := &TaughtClassesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/education/users/{educationUser_id}/taughtClasses{?top,skip,search,filter,count,orderby,select,expand}";
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
func NewTaughtClassesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TaughtClassesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTaughtClassesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TaughtClassesRequestBuilder) CreateGetRequestInformation(q func (value *TaughtClassesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(TaughtClassesRequestBuilderGetQueryParameters)
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
func (m *TaughtClassesRequestBuilder) Delta()(*i9a82ba081a87a0b82fe7518f2699a91f6cdabe7560d9041a4bbd68ca17595c04.DeltaRequestBuilder) {
    return i9a82ba081a87a0b82fe7518f2699a91f6cdabe7560d9041a4bbd68ca17595c04.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TaughtClassesRequestBuilder) Get(q func (value *TaughtClassesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*TaughtClassesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTaughtClassesResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*TaughtClassesResponse), nil
}
func (m *TaughtClassesRequestBuilder) Ref()(*ibfebe8e1821c93dda8997c543dd7161494427c674556776df0eee3b6720e6f3c.RefRequestBuilder) {
    return ibfebe8e1821c93dda8997c543dd7161494427c674556776df0eee3b6720e6f3c.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}