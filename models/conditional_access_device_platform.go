package models
import (
    "errors"
)
// Provides operations to manage the identityContainer singleton.
type ConditionalAccessDevicePlatform int

const (
    ANDROID_CONDITIONALACCESSDEVICEPLATFORM ConditionalAccessDevicePlatform = iota
    IOS_CONDITIONALACCESSDEVICEPLATFORM
    WINDOWS_CONDITIONALACCESSDEVICEPLATFORM
    WINDOWSPHONE_CONDITIONALACCESSDEVICEPLATFORM
    MACOS_CONDITIONALACCESSDEVICEPLATFORM
    ALL_CONDITIONALACCESSDEVICEPLATFORM
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSDEVICEPLATFORM
    LINUX_CONDITIONALACCESSDEVICEPLATFORM
)

func (i ConditionalAccessDevicePlatform) String() string {
    return []string{"android", "iOS", "windows", "windowsPhone", "macOS", "all", "unknownFutureValue", "linux"}[i]
}
func ParseConditionalAccessDevicePlatform(v string) (interface{}, error) {
    result := ANDROID_CONDITIONALACCESSDEVICEPLATFORM
    switch v {
        case "android":
            result = ANDROID_CONDITIONALACCESSDEVICEPLATFORM
        case "iOS":
            result = IOS_CONDITIONALACCESSDEVICEPLATFORM
        case "windows":
            result = WINDOWS_CONDITIONALACCESSDEVICEPLATFORM
        case "windowsPhone":
            result = WINDOWSPHONE_CONDITIONALACCESSDEVICEPLATFORM
        case "macOS":
            result = MACOS_CONDITIONALACCESSDEVICEPLATFORM
        case "all":
            result = ALL_CONDITIONALACCESSDEVICEPLATFORM
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONDITIONALACCESSDEVICEPLATFORM
        case "linux":
            result = LINUX_CONDITIONALACCESSDEVICEPLATFORM
        default:
            return 0, errors.New("Unknown ConditionalAccessDevicePlatform value: " + v)
    }
    return &result, nil
}
func SerializeConditionalAccessDevicePlatform(values []ConditionalAccessDevicePlatform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
