package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder provides operations to manage the permissionGrants property of the microsoft.graph.group entity.
type ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters the permission that has been granted for a group to a specific application. Supports $expand.
type ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters
}
// ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CheckMemberGroups()(*ItemPermissionGrantsItemCheckMemberGroupsRequestBuilder) {
    return NewItemPermissionGrantsItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CheckMemberObjects()(*ItemPermissionGrantsItemCheckMemberObjectsRequestBuilder) {
    return NewItemPermissionGrantsItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal instantiates a new ResourceSpecificPermissionGrantItemRequestBuilder and sets the default values.
func NewItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    m := &ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/permissionGrants/{resourceSpecificPermissionGrant%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder instantiates a new ResourceSpecificPermissionGrantItemRequestBuilder and sets the default values.
func NewItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property permissionGrants for groups
func (m *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the permission that has been granted for a group to a specific application. Supports $expand.
func (m *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property permissionGrants in groups
func (m *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ResourceSpecificPermissionGrantable, requestConfiguration *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property permissionGrants for groups
func (m *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the permission that has been granted for a group to a specific application. Supports $expand.
func (m *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ResourceSpecificPermissionGrantable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateResourceSpecificPermissionGrantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ResourceSpecificPermissionGrantable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) GetMemberGroups()(*ItemPermissionGrantsItemGetMemberGroupsRequestBuilder) {
    return NewItemPermissionGrantsItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) GetMemberObjects()(*ItemPermissionGrantsItemGetMemberObjectsRequestBuilder) {
    return NewItemPermissionGrantsItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property permissionGrants in groups
func (m *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ResourceSpecificPermissionGrantable, requestConfiguration *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ResourceSpecificPermissionGrantable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateResourceSpecificPermissionGrantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ResourceSpecificPermissionGrantable), nil
}
// Restore provides operations to call the restore method.
func (m *ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) Restore()(*ItemPermissionGrantsItemRestoreRequestBuilder) {
    return NewItemPermissionGrantsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}