package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder provides operations to manage the children property of the microsoft.graph.termStore.set entity.
type GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetQueryParameters children terms of set in term [store].
type GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetQueryParameters
}
// GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.termStore.term entity.
func (m *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) Children()(*GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemChildrenRequestBuilder) {
    return NewGroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.termStore.term entity.
func (m *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) ChildrenById(id string)(*GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemChildrenTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did1"] = id
    }
    return NewGroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemChildrenTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewGroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderInternal instantiates a new TermItemRequestBuilder and sets the default values.
func NewGroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) {
    m := &GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStores/{store%2Did}/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/children/{term%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder instantiates a new TermItemRequestBuilder and sets the default values.
func NewGroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property children for groups
func (m *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation children terms of set in term [store].
func (m *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property children in groups
func (m *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, requestConfiguration *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property children for groups
func (m *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get children terms of set in term [store].
func (m *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable), nil
}
// Patch update the navigation property children in groups
func (m *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) Patch(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, requestConfiguration *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable), nil
}
// Relations provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
func (m *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) Relations()(*GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRequestBuilder) {
    return NewGroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
func (m *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) RelationsById(id string)(*GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return NewGroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Set provides operations to manage the set property of the microsoft.graph.termStore.term entity.
func (m *GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) Set()(*GroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemSetRequestBuilder) {
    return NewGroupsItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
