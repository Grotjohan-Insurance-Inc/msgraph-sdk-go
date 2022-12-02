package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
type UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetQueryParameters content in a message hosted by Microsoft Teams - for example, images or code snippets.
type UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetQueryParameters
}
// UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderInternal instantiates a new ChatMessageHostedContentItemRequestBuilder and sets the default values.
func NewUsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder) {
    m := &UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}/hostedContents/{chatMessageHostedContent%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder instantiates a new ChatMessageHostedContentItemRequestBuilder and sets the default values.
func NewUsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property hostedContents for users
func (m *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation content in a message hosted by Microsoft Teams - for example, images or code snippets.
func (m *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property hostedContents in users
func (m *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, requestConfiguration *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property hostedContents for users
func (m *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get content in a message hosted by Microsoft Teams - for example, images or code snippets.
func (m *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatMessageHostedContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable), nil
}
// Patch update the navigation property hostedContents in users
func (m *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, requestConfiguration *UsersItemJoinedTeamsItemPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatMessageHostedContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable), nil
}
