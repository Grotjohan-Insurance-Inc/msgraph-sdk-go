package getpstncallswithfromdatetimewithtodatetime

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19 "github.com/microsoftgraph/msgraph-sdk-go/models/callrecords"
)

// GetPstnCallsWithFromDateTimeWithToDateTimeResponseable 
type GetPstnCallsWithFromDateTimeWithToDateTimeResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.PstnCallLogRowable)
    SetValue(value []iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.PstnCallLogRowable)()
}
