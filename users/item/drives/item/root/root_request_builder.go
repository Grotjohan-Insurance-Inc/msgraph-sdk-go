package root

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i025b69ef2d2ddbfb7e9807a38f690bdf4cd63d6712a67206da809a0cdfc42189 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/thumbnails"
    i135e884c820565705b38a640ffdf31280e134de90528d9785ec0e75376321a61 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/analytics"
    i2ea67eb04bbbfa82d90dd88dddc4c42ccdbb62c2e552c27c0091b1608040bcf8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/listitem"
    i3d0e4f8aa6ec904f8c27942ce3497aadbb1cd99228b4c60d2a161cc1cd3859cd "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/children"
    i5bbb0ccbc4c3843f9a789010d1b81a6c7dc71a17cc52f5e0d68f386d2057dbc9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/subscriptions"
    i7b6ee7d5b6663b375ed921bc6ba9ad187a415e8a469daf7f11aa380553d07065 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/versions"
    iaa1d84731a4b8785f6dacbbb5439c2508806285148dd6a0c01c1be4f2b099002 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/content"
    ib1efd02c4c2751e576dd0de1c915b8e6db635eae24dd01117248f2e015d27ec6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/permissions"
    i67ec6eaee9045e07acc845ccdf57c93da38efc06e18c1bb8dfd55e85d7b6320d "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/subscriptions/item"
    i8dac7fcd490022c229aca597d1ef2174bb968ffc2cd99e554ec494acf608d88c "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/thumbnails/item"
    iccb7f2d5d33cfb54ca0c62427ff964d27d77254ebd63aa68529b5eed5a42266d "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/children/item"
    id0c42f467a4df86619438eccced198bfbc3edae46d9a7697c9794a1f605d4d11 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/permissions/item"
    id5e75288d7196634c44036e792f83a221d2e22ef830f4c48a8c0e09650983d14 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/root/versions/item"
)

// RootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type RootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RootRequestBuilderDeleteOptions options for Delete
type RootRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// RootRequestBuilderGetOptions options for Get
type RootRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RootRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// RootRequestBuilderGetQueryParameters the root folder of the drive. Read-only.
type RootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RootRequestBuilderPatchOptions options for Patch
type RootRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*i135e884c820565705b38a640ffdf31280e134de90528d9785ec0e75376321a61.AnalyticsRequestBuilder) {
    return i135e884c820565705b38a640ffdf31280e134de90528d9785ec0e75376321a61.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*i3d0e4f8aa6ec904f8c27942ce3497aadbb1cd99228b4c60d2a161cc1cd3859cd.ChildrenRequestBuilder) {
    return i3d0e4f8aa6ec904f8c27942ce3497aadbb1cd99228b4c60d2a161cc1cd3859cd.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*iccb7f2d5d33cfb54ca0c62427ff964d27d77254ebd63aa68529b5eed5a42266d.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return iccb7f2d5d33cfb54ca0c62427ff964d27d77254ebd63aa68529b5eed5a42266d.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/root{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRootRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *RootRequestBuilder) Content()(*iaa1d84731a4b8785f6dacbbb5439c2508806285148dd6a0c01c1be4f2b099002.ContentRequestBuilder) {
    return iaa1d84731a4b8785f6dacbbb5439c2508806285148dd6a0c01c1be4f2b099002.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for users
func (m *RootRequestBuilder) CreateDeleteRequestInformation(options *RootRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformation(options *RootRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property root in users
func (m *RootRequestBuilder) CreatePatchRequestInformation(options *RootRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property root for users
func (m *RootRequestBuilder) Delete(options *RootRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the root folder of the drive. Read-only.
func (m *RootRequestBuilder) Get(options *RootRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable), nil
}
// ListItem the listItem property
func (m *RootRequestBuilder) ListItem()(*i2ea67eb04bbbfa82d90dd88dddc4c42ccdbb62c2e552c27c0091b1608040bcf8.ListItemRequestBuilder) {
    return i2ea67eb04bbbfa82d90dd88dddc4c42ccdbb62c2e552c27c0091b1608040bcf8.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in users
func (m *RootRequestBuilder) Patch(options *RootRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Permissions the permissions property
func (m *RootRequestBuilder) Permissions()(*ib1efd02c4c2751e576dd0de1c915b8e6db635eae24dd01117248f2e015d27ec6.PermissionsRequestBuilder) {
    return ib1efd02c4c2751e576dd0de1c915b8e6db635eae24dd01117248f2e015d27ec6.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*id0c42f467a4df86619438eccced198bfbc3edae46d9a7697c9794a1f605d4d11.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return id0c42f467a4df86619438eccced198bfbc3edae46d9a7697c9794a1f605d4d11.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*i5bbb0ccbc4c3843f9a789010d1b81a6c7dc71a17cc52f5e0d68f386d2057dbc9.SubscriptionsRequestBuilder) {
    return i5bbb0ccbc4c3843f9a789010d1b81a6c7dc71a17cc52f5e0d68f386d2057dbc9.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i67ec6eaee9045e07acc845ccdf57c93da38efc06e18c1bb8dfd55e85d7b6320d.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i67ec6eaee9045e07acc845ccdf57c93da38efc06e18c1bb8dfd55e85d7b6320d.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*i025b69ef2d2ddbfb7e9807a38f690bdf4cd63d6712a67206da809a0cdfc42189.ThumbnailsRequestBuilder) {
    return i025b69ef2d2ddbfb7e9807a38f690bdf4cd63d6712a67206da809a0cdfc42189.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i8dac7fcd490022c229aca597d1ef2174bb968ffc2cd99e554ec494acf608d88c.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i8dac7fcd490022c229aca597d1ef2174bb968ffc2cd99e554ec494acf608d88c.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*i7b6ee7d5b6663b375ed921bc6ba9ad187a415e8a469daf7f11aa380553d07065.VersionsRequestBuilder) {
    return i7b6ee7d5b6663b375ed921bc6ba9ad187a415e8a469daf7f11aa380553d07065.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*id5e75288d7196634c44036e792f83a221d2e22ef830f4c48a8c0e09650983d14.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return id5e75288d7196634c44036e792f83a221d2e22ef830f4c48a8c0e09650983d14.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}