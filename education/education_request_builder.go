package education

import (
    i11d420057e9890f34c6f87265f6025f3489c0f78de3b99ece6ba73f225d5a75c "github.com/microsoftgraph/msgraph-sdk-go/education/me"
    i82d352ad9c3de7536dcc103a1a2a276e7beba1b35548191dda9e846d8c21e897 "github.com/microsoftgraph/msgraph-sdk-go/education/users"
    i8803b2d448091478d216c51910e6ed01bbc8c76e5b76968b633118267314ebe9 "github.com/microsoftgraph/msgraph-sdk-go/education/schools"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    if6b46a923414b614133630b614b404c3005a11daa1e05f10bb834755518d40dd "github.com/microsoftgraph/msgraph-sdk-go/education/classes"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i31a8a42bc05cb93b0a8b6dc7e28fb5a481e3934df6dc18de8217d0fe268d617c "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i6fc2718a806c86250bd9f9aaba5418413286c98ed345a7dbe988c558836bde8e "github.com/microsoftgraph/msgraph-sdk-go/education/users/item"
    ibcc84b6252b9c075d39d4d8641ef3cf01a04613a5da305d7d48d21c122395065 "github.com/microsoftgraph/msgraph-sdk-go/education/schools/item"
)

type EducationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EducationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *EducationRequestBuilder) Classes()(*if6b46a923414b614133630b614b404c3005a11daa1e05f10bb834755518d40dd.ClassesRequestBuilder) {
    return if6b46a923414b614133630b614b404c3005a11daa1e05f10bb834755518d40dd.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationRequestBuilder) ClassesById(id string)(*i31a8a42bc05cb93b0a8b6dc7e28fb5a481e3934df6dc18de8217d0fe268d617c.EducationClassRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationClass_id"] = id
    }
    return i31a8a42bc05cb93b0a8b6dc7e28fb5a481e3934df6dc18de8217d0fe268d617c.NewEducationClassRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewEducationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationRequestBuilder) {
    m := &EducationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/education{?select,expand}";
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
func NewEducationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EducationRequestBuilder) CreateGetRequestInformation(q func (value *EducationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EducationRequestBuilderGetQueryParameters)
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
func (m *EducationRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Education, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationRequestBuilder) Get(q func (value *EducationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEducationRoot() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationRoot), nil
}
func (m *EducationRequestBuilder) Me()(*i11d420057e9890f34c6f87265f6025f3489c0f78de3b99ece6ba73f225d5a75c.MeRequestBuilder) {
    return i11d420057e9890f34c6f87265f6025f3489c0f78de3b99ece6ba73f225d5a75c.NewMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Education, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EducationRequestBuilder) Schools()(*i8803b2d448091478d216c51910e6ed01bbc8c76e5b76968b633118267314ebe9.SchoolsRequestBuilder) {
    return i8803b2d448091478d216c51910e6ed01bbc8c76e5b76968b633118267314ebe9.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationRequestBuilder) SchoolsById(id string)(*ibcc84b6252b9c075d39d4d8641ef3cf01a04613a5da305d7d48d21c122395065.EducationSchoolRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationSchool_id"] = id
    }
    return ibcc84b6252b9c075d39d4d8641ef3cf01a04613a5da305d7d48d21c122395065.NewEducationSchoolRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationRequestBuilder) Users()(*i82d352ad9c3de7536dcc103a1a2a276e7beba1b35548191dda9e846d8c21e897.UsersRequestBuilder) {
    return i82d352ad9c3de7536dcc103a1a2a276e7beba1b35548191dda9e846d8c21e897.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationRequestBuilder) UsersById(id string)(*i6fc2718a806c86250bd9f9aaba5418413286c98ed345a7dbe988c558836bde8e.EducationUserRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationUser_id"] = id
    }
    return i6fc2718a806c86250bd9f9aaba5418413286c98ed345a7dbe988c558836bde8e.NewEducationUserRequestBuilderInternal(urlTplParams, m.requestAdapter);
}