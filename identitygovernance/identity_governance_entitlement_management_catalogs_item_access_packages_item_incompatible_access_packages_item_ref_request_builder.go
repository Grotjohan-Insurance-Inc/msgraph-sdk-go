package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilder provides operations to manage the collection of identityGovernance entities.
type IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderDeleteQueryParameters delete ref of navigation property incompatibleAccessPackages for identityGovernance
type IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderDeleteQueryParameters
}
// NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/incompatibleAccessPackages/{accessPackage%2Did1}/$ref{?%40id*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete ref of navigation property incompatibleAccessPackages for identityGovernance
func (m *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete ref of navigation property incompatibleAccessPackages for identityGovernance
func (m *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
