package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilder provides operations to call the removeHold method.
type CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal instantiates a new MicrosoftGraphSecurityRemoveHoldRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    m := &CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}/microsoft.graph.security.removeHold", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilder instantiates a new MicrosoftGraphSecurityRemoveHoldRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post start the process of removing hold from eDiscovery custodians. After the operation is created, you can get the status by retrieving the Location parameter from the response headers. The location provides a URL that will return an eDiscoveryHoldOperation object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycustodian-removehold?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation start the process of removing hold from eDiscovery custodians. After the operation is created, you can get the status by retrieving the Location parameter from the response headers. The location provides a URL that will return an eDiscoveryHoldOperation object.
func (m *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
