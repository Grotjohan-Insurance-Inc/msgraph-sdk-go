package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DrivesItemItemsItemPermissionsItemGrantPostRequestBody provides operations to call the grant method.
type DrivesItemItemsItemPermissionsItemGrantPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The recipients property
    recipients []DriveRecipientable
    // The roles property
    roles []string
}
// NewDrivesItemItemsItemPermissionsItemGrantPostRequestBody instantiates a new DrivesItemItemsItemPermissionsItemGrantPostRequestBody and sets the default values.
func NewDrivesItemItemsItemPermissionsItemGrantPostRequestBody()(*DrivesItemItemsItemPermissionsItemGrantPostRequestBody) {
    m := &DrivesItemItemsItemPermissionsItemGrantPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDrivesItemItemsItemPermissionsItemGrantPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDrivesItemItemsItemPermissionsItemGrantPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDrivesItemItemsItemPermissionsItemGrantPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DrivesItemItemsItemPermissionsItemGrantPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DrivesItemItemsItemPermissionsItemGrantPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["recipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveRecipientable, len(val))
            for i, v := range val {
                res[i] = v.(DriveRecipientable)
            }
            m.SetRecipients(res)
        }
        return nil
    }
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoles(res)
        }
        return nil
    }
    return res
}
// GetRecipients gets the recipients property value. The recipients property
func (m *DrivesItemItemsItemPermissionsItemGrantPostRequestBody) GetRecipients()([]DriveRecipientable) {
    return m.recipients
}
// GetRoles gets the roles property value. The roles property
func (m *DrivesItemItemsItemPermissionsItemGrantPostRequestBody) GetRoles()([]string) {
    return m.roles
}
// Serialize serializes information the current object
func (m *DrivesItemItemsItemPermissionsItemGrantPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecipients()))
        for i, v := range m.GetRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("recipients", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        err := writer.WriteCollectionOfStringValues("roles", m.GetRoles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DrivesItemItemsItemPermissionsItemGrantPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetRecipients sets the recipients property value. The recipients property
func (m *DrivesItemItemsItemPermissionsItemGrantPostRequestBody) SetRecipients(value []DriveRecipientable)() {
    m.recipients = value
}
// SetRoles sets the roles property value. The roles property
func (m *DrivesItemItemsItemPermissionsItemGrantPostRequestBody) SetRoles(value []string)() {
    m.roles = value
}
