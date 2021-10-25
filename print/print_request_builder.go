package print

import (
    i0bda270a8db970823b8449367f0a1b6167aeb67bb1a0b116061e10c72e9e8d5d "github.com/microsoftgraph/msgraph-sdk-go/print/printers"
    i0f704d64ae3c276c8ff23cafb44315ef4067a668629360d124bdabfea38afdd3 "github.com/microsoftgraph/msgraph-sdk-go/print/connectors"
    i54feb3a1bee5e2d3891bf4b38d7ad46df95f464ec5502fc27ce12645532dcf04 "github.com/microsoftgraph/msgraph-sdk-go/print/shares"
    i802e345f3915c26a53a0e538ad96acb01b0ad515a9510581f00bcedb2e2d1266 "github.com/microsoftgraph/msgraph-sdk-go/print/services"
    ib24f76d25d0e8f20d062be56f8b94808153dc20498d344b76af7867075ea8ad9 "github.com/microsoftgraph/msgraph-sdk-go/print/operations"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ied8956e290d0c898cb2abe2b85d9373567ead7ae2e05a776597718dd3f1709e8 "github.com/microsoftgraph/msgraph-sdk-go/print/taskdefinitions"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0b3afc35ae08ae48e886b1ea1ae2fe47da46f82bd161cb48f9ff550676f78301 "github.com/microsoftgraph/msgraph-sdk-go/print/services/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7b3d94bc34459d3d2fa55e5dc4096ab7096e743ffc216c2d05b8f451f231fde1 "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item"
    i827522a32e9dfc8181a7f7793aec64fdf0814658430464c0b64af19c8e4e9a96 "github.com/microsoftgraph/msgraph-sdk-go/print/taskdefinitions/item"
    i8e18f5c596955d2371ff489e959f3488755280f5b7014ba5a0f5b95b7dc44bf2 "github.com/microsoftgraph/msgraph-sdk-go/print/shares/item"
    id0e61828b4320e56c027152ab07b04414264a89aa169cd832394397ada28defd "github.com/microsoftgraph/msgraph-sdk-go/print/operations/item"
    iebd42d165d4ff6aca1eec85e8351f67b97a4208ea00a5bc8ac67f4f1f191b677 "github.com/microsoftgraph/msgraph-sdk-go/print/connectors/item"
)

type PrintRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type PrintRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *PrintRequestBuilder) Connectors()(*i0f704d64ae3c276c8ff23cafb44315ef4067a668629360d124bdabfea38afdd3.ConnectorsRequestBuilder) {
    return i0f704d64ae3c276c8ff23cafb44315ef4067a668629360d124bdabfea38afdd3.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) ConnectorsById(id string)(*iebd42d165d4ff6aca1eec85e8351f67b97a4208ea00a5bc8ac67f4f1f191b677.PrintConnectorRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printConnector_id"] = id
    }
    return iebd42d165d4ff6aca1eec85e8351f67b97a4208ea00a5bc8ac67f4f1f191b677.NewPrintConnectorRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewPrintRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintRequestBuilder) {
    m := &PrintRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/print{?select,expand}";
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
func NewPrintRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrintRequestBuilder) CreateGetRequestInformation(q func (value *PrintRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PrintRequestBuilderGetQueryParameters)
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
func (m *PrintRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Print, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrintRequestBuilder) Get(q func (value *PrintRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Print, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPrint() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Print), nil
}
func (m *PrintRequestBuilder) Operations()(*ib24f76d25d0e8f20d062be56f8b94808153dc20498d344b76af7867075ea8ad9.OperationsRequestBuilder) {
    return ib24f76d25d0e8f20d062be56f8b94808153dc20498d344b76af7867075ea8ad9.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) OperationsById(id string)(*id0e61828b4320e56c027152ab07b04414264a89aa169cd832394397ada28defd.PrintOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printOperation_id"] = id
    }
    return id0e61828b4320e56c027152ab07b04414264a89aa169cd832394397ada28defd.NewPrintOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrintRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Print, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PrintRequestBuilder) Printers()(*i0bda270a8db970823b8449367f0a1b6167aeb67bb1a0b116061e10c72e9e8d5d.PrintersRequestBuilder) {
    return i0bda270a8db970823b8449367f0a1b6167aeb67bb1a0b116061e10c72e9e8d5d.NewPrintersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) PrintersById(id string)(*i7b3d94bc34459d3d2fa55e5dc4096ab7096e743ffc216c2d05b8f451f231fde1.PrinterRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printer_id"] = id
    }
    return i7b3d94bc34459d3d2fa55e5dc4096ab7096e743ffc216c2d05b8f451f231fde1.NewPrinterRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrintRequestBuilder) Services()(*i802e345f3915c26a53a0e538ad96acb01b0ad515a9510581f00bcedb2e2d1266.ServicesRequestBuilder) {
    return i802e345f3915c26a53a0e538ad96acb01b0ad515a9510581f00bcedb2e2d1266.NewServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) ServicesById(id string)(*i0b3afc35ae08ae48e886b1ea1ae2fe47da46f82bd161cb48f9ff550676f78301.PrintServiceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printService_id"] = id
    }
    return i0b3afc35ae08ae48e886b1ea1ae2fe47da46f82bd161cb48f9ff550676f78301.NewPrintServiceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrintRequestBuilder) Shares()(*i54feb3a1bee5e2d3891bf4b38d7ad46df95f464ec5502fc27ce12645532dcf04.SharesRequestBuilder) {
    return i54feb3a1bee5e2d3891bf4b38d7ad46df95f464ec5502fc27ce12645532dcf04.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) SharesById(id string)(*i8e18f5c596955d2371ff489e959f3488755280f5b7014ba5a0f5b95b7dc44bf2.PrinterShareRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printerShare_id"] = id
    }
    return i8e18f5c596955d2371ff489e959f3488755280f5b7014ba5a0f5b95b7dc44bf2.NewPrinterShareRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrintRequestBuilder) TaskDefinitions()(*ied8956e290d0c898cb2abe2b85d9373567ead7ae2e05a776597718dd3f1709e8.TaskDefinitionsRequestBuilder) {
    return ied8956e290d0c898cb2abe2b85d9373567ead7ae2e05a776597718dd3f1709e8.NewTaskDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) TaskDefinitionsById(id string)(*i827522a32e9dfc8181a7f7793aec64fdf0814658430464c0b64af19c8e4e9a96.PrintTaskDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printTaskDefinition_id"] = id
    }
    return i827522a32e9dfc8181a7f7793aec64fdf0814658430464c0b64af19c8e4e9a96.NewPrintTaskDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}