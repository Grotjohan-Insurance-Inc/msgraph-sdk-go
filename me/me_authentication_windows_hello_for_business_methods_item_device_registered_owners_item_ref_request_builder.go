package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder provides operations to manage the collection of user entities.
type MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteQueryParameters delete ref of navigation property registeredOwners for me
type MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteQueryParameters
}
// NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) {
    m := &MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/$ref{?%40id*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete ref of navigation property registeredOwners for me
func (m *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete ref of navigation property registeredOwners for me
func (m *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
