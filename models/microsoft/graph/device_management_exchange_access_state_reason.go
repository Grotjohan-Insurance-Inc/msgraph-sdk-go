package graph
import (
    "strings"
    "errors"
)
type DeviceManagementExchangeAccessStateReason int

const (
    NONE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON DeviceManagementExchangeAccessStateReason = iota
    UNKNOWN_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    EXCHANGEGLOBALRULE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    EXCHANGEINDIVIDUALRULE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    EXCHANGEDEVICERULE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    EXCHANGEUPGRADE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    EXCHANGEMAILBOXPOLICY_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    OTHER_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    COMPLIANT_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    NOTCOMPLIANT_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    NOTENROLLED_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    UNKNOWNLOCATION_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    MFAREQUIRED_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    AZUREADBLOCKDUETOACCESSPOLICY_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    COMPROMISEDPASSWORD_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    DEVICENOTKNOWNWITHMANAGEDAPP_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
)

func (i DeviceManagementExchangeAccessStateReason) String() string {
    return []string{"NONE", "UNKNOWN", "EXCHANGEGLOBALRULE", "EXCHANGEINDIVIDUALRULE", "EXCHANGEDEVICERULE", "EXCHANGEUPGRADE", "EXCHANGEMAILBOXPOLICY", "OTHER", "COMPLIANT", "NOTCOMPLIANT", "NOTENROLLED", "UNKNOWNLOCATION", "MFAREQUIRED", "AZUREADBLOCKDUETOACCESSPOLICY", "COMPROMISEDPASSWORD", "DEVICENOTKNOWNWITHMANAGEDAPP"}[i]
}
func ParseDeviceManagementExchangeAccessStateReason(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "UNKNOWN":
            return UNKNOWN_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "EXCHANGEGLOBALRULE":
            return EXCHANGEGLOBALRULE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "EXCHANGEINDIVIDUALRULE":
            return EXCHANGEINDIVIDUALRULE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "EXCHANGEDEVICERULE":
            return EXCHANGEDEVICERULE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "EXCHANGEUPGRADE":
            return EXCHANGEUPGRADE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "EXCHANGEMAILBOXPOLICY":
            return EXCHANGEMAILBOXPOLICY_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "OTHER":
            return OTHER_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "COMPLIANT":
            return COMPLIANT_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "NOTCOMPLIANT":
            return NOTCOMPLIANT_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "NOTENROLLED":
            return NOTENROLLED_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "UNKNOWNLOCATION":
            return UNKNOWNLOCATION_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "MFAREQUIRED":
            return MFAREQUIRED_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "AZUREADBLOCKDUETOACCESSPOLICY":
            return AZUREADBLOCKDUETOACCESSPOLICY_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "COMPROMISEDPASSWORD":
            return COMPROMISEDPASSWORD_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
        case "DEVICENOTKNOWNWITHMANAGEDAPP":
            return DEVICENOTKNOWNWITHMANAGEDAPP_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON, nil
    }
    return 0, errors.New("Unknown DeviceManagementExchangeAccessStateReason value: " + v)
}
func SerializeDeviceManagementExchangeAccessStateReason(values []DeviceManagementExchangeAccessStateReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}