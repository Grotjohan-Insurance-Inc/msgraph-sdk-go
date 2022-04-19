package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0853bc5e78a4ac665d8aa887531e364debb1ba91c3a1bc736be9331cc8db0b52 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/singlevalueextendedproperties"
    i25a4d487ad09f2eb45ba28744d3fe919b7b5dbbf1c645ddd9776a9b357d0ff96 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview"
    i3abccbd6b1092e6246d7bc4dd7dcac1b6938ee6bf024fb4eccaaeecfb29ff1b6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events"
    i3e09978bd77ced5481016e4684368f72edfa92ef709a19504e9258cf8865bfc6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/allowedcalendarsharingroleswithuser"
    i847a76068ecb832435c51c746fba99e17f01da45f86bc0ea0f29b5a1d42ae5b3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/multivalueextendedproperties"
    ia0a09be4f5c1f5b3b99aa3064830cdc6c730ceea9bdb4dbc6934a7d5f0ffec6f "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/getschedule"
    iaac11e3c8bf2dd29e59be19f34577e748fb426a2bb9b94aae4be18a9b2f6474a "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarpermissions"
    i67198c8bd92c3d2fcf155bff41150fc55dc916147fb9b1eb9ca2ccf6b9d1ed6e "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item"
    i857bdaf62ca4819c99b59159f0888909252fbffeba2fa7050dfc75ce2c6b598d "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/singlevalueextendedproperties/item"
    i8783c6ceaf93b1654906d93bb1e4267cb1060f0bcdbc2d302e4f64bfac674b76 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item"
    i8b359363f8c83629e3e25072f613c3d3d807cc671fe5472a27f65565ce218c26 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarpermissions/item"
    idefd7aa8b66a92e862195e44b4d12afbb8cb1f08136af3a806efad5c689470d7 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/multivalueextendedproperties/item"
)

// CalendarItemRequestBuilder provides operations to manage the calendars property of the microsoft.graph.calendarGroup entity.
type CalendarItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarItemRequestBuilderDeleteOptions options for Delete
type CalendarItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CalendarItemRequestBuilderGetOptions options for Get
type CalendarItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CalendarItemRequestBuilderGetQueryParameters the calendars in the calendar group. Navigation property. Read-only. Nullable.
type CalendarItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarItemRequestBuilderPatchOptions options for Patch
type CalendarItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AllowedCalendarSharingRolesWithUser provides operations to call the allowedCalendarSharingRoles method.
func (m *CalendarItemRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i3e09978bd77ced5481016e4684368f72edfa92ef709a19504e9258cf8865bfc6.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i3e09978bd77ced5481016e4684368f72edfa92ef709a19504e9258cf8865bfc6.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
// CalendarPermissions the calendarPermissions property
func (m *CalendarItemRequestBuilder) CalendarPermissions()(*iaac11e3c8bf2dd29e59be19f34577e748fb426a2bb9b94aae4be18a9b2f6474a.CalendarPermissionsRequestBuilder) {
    return iaac11e3c8bf2dd29e59be19f34577e748fb426a2bb9b94aae4be18a9b2f6474a.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.calendarPermissions.item collection
func (m *CalendarItemRequestBuilder) CalendarPermissionsById(id string)(*i8b359363f8c83629e3e25072f613c3d3d807cc671fe5472a27f65565ce218c26.CalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission%2Did"] = id
    }
    return i8b359363f8c83629e3e25072f613c3d3d807cc671fe5472a27f65565ce218c26.NewCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView the calendarView property
func (m *CalendarItemRequestBuilder) CalendarView()(*i25a4d487ad09f2eb45ba28744d3fe919b7b5dbbf1c645ddd9776a9b357d0ff96.CalendarViewRequestBuilder) {
    return i25a4d487ad09f2eb45ba28744d3fe919b7b5dbbf1c645ddd9776a9b357d0ff96.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item collection
func (m *CalendarItemRequestBuilder) CalendarViewById(id string)(*i8783c6ceaf93b1654906d93bb1e4267cb1060f0bcdbc2d302e4f64bfac674b76.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i8783c6ceaf93b1654906d93bb1e4267cb1060f0bcdbc2d302e4f64bfac674b76.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarItemRequestBuilderInternal instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarItemRequestBuilder) {
    m := &CalendarItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarItemRequestBuilder instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property calendars for me
func (m *CalendarItemRequestBuilder) CreateDeleteRequestInformation(options *CalendarItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the calendars in the calendar group. Navigation property. Read-only. Nullable.
func (m *CalendarItemRequestBuilder) CreateGetRequestInformation(options *CalendarItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property calendars in me
func (m *CalendarItemRequestBuilder) CreatePatchRequestInformation(options *CalendarItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property calendars for me
func (m *CalendarItemRequestBuilder) Delete(options *CalendarItemRequestBuilderDeleteOptions)(error) {
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
// Events the events property
func (m *CalendarItemRequestBuilder) Events()(*i3abccbd6b1092e6246d7bc4dd7dcac1b6938ee6bf024fb4eccaaeecfb29ff1b6.EventsRequestBuilder) {
    return i3abccbd6b1092e6246d7bc4dd7dcac1b6938ee6bf024fb4eccaaeecfb29ff1b6.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.events.item collection
func (m *CalendarItemRequestBuilder) EventsById(id string)(*i67198c8bd92c3d2fcf155bff41150fc55dc916147fb9b1eb9ca2ccf6b9d1ed6e.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i67198c8bd92c3d2fcf155bff41150fc55dc916147fb9b1eb9ca2ccf6b9d1ed6e.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the calendars in the calendar group. Navigation property. Read-only. Nullable.
func (m *CalendarItemRequestBuilder) Get(options *CalendarItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCalendarFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable), nil
}
// GetSchedule the getSchedule property
func (m *CalendarItemRequestBuilder) GetSchedule()(*ia0a09be4f5c1f5b3b99aa3064830cdc6c730ceea9bdb4dbc6934a7d5f0ffec6f.GetScheduleRequestBuilder) {
    return ia0a09be4f5c1f5b3b99aa3064830cdc6c730ceea9bdb4dbc6934a7d5f0ffec6f.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *CalendarItemRequestBuilder) MultiValueExtendedProperties()(*i847a76068ecb832435c51c746fba99e17f01da45f86bc0ea0f29b5a1d42ae5b3.MultiValueExtendedPropertiesRequestBuilder) {
    return i847a76068ecb832435c51c746fba99e17f01da45f86bc0ea0f29b5a1d42ae5b3.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.multiValueExtendedProperties.item collection
func (m *CalendarItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*idefd7aa8b66a92e862195e44b4d12afbb8cb1f08136af3a806efad5c689470d7.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return idefd7aa8b66a92e862195e44b4d12afbb8cb1f08136af3a806efad5c689470d7.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendars in me
func (m *CalendarItemRequestBuilder) Patch(options *CalendarItemRequestBuilderPatchOptions)(error) {
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
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *CalendarItemRequestBuilder) SingleValueExtendedProperties()(*i0853bc5e78a4ac665d8aa887531e364debb1ba91c3a1bc736be9331cc8db0b52.SingleValueExtendedPropertiesRequestBuilder) {
    return i0853bc5e78a4ac665d8aa887531e364debb1ba91c3a1bc736be9331cc8db0b52.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.singleValueExtendedProperties.item collection
func (m *CalendarItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i857bdaf62ca4819c99b59159f0888909252fbffeba2fa7050dfc75ce2c6b598d.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i857bdaf62ca4819c99b59159f0888909252fbffeba2fa7050dfc75ce2c6b598d.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}