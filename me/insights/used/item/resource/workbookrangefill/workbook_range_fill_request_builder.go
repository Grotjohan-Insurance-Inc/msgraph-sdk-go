package workbookrangefill

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ibf9f367d3cb8a81ddc7557acb6d52bdd9dc80236325b11ba81afdead6e53a325 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrangefill/clear"
)

type WorkbookRangeFillRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *WorkbookRangeFillRequestBuilder) Clear()(*ibf9f367d3cb8a81ddc7557acb6d52bdd9dc80236325b11ba81afdead6e53a325.ClearRequestBuilder) {
    return ibf9f367d3cb8a81ddc7557acb6d52bdd9dc80236325b11ba81afdead6e53a325.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewWorkbookRangeFillRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFillRequestBuilder) {
    m := &WorkbookRangeFillRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/used/{usedInsight_id}/resource/microsoft.graph.workbookRangeFill";
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
func NewWorkbookRangeFillRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFillRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeFillRequestBuilderInternal(urlParams, requestAdapter)
}