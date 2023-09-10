package security
import (
    "errors"
)
// 
type ContainerPortProtocol int

const (
    UDP_CONTAINERPORTPROTOCOL ContainerPortProtocol = iota
    TCP_CONTAINERPORTPROTOCOL
    SCTP_CONTAINERPORTPROTOCOL
    UNKNOWNFUTUREVALUE_CONTAINERPORTPROTOCOL
)

func (i ContainerPortProtocol) String() string {
    return []string{"udp", "tcp", "sctp", "unknownFutureValue"}[i]
}
func ParseContainerPortProtocol(v string) (any, error) {
    result := UDP_CONTAINERPORTPROTOCOL
    switch v {
        case "udp":
            result = UDP_CONTAINERPORTPROTOCOL
        case "tcp":
            result = TCP_CONTAINERPORTPROTOCOL
        case "sctp":
            result = SCTP_CONTAINERPORTPROTOCOL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONTAINERPORTPROTOCOL
        default:
            return 0, errors.New("Unknown ContainerPortProtocol value: " + v)
    }
    return &result, nil
}
func SerializeContainerPortProtocol(values []ContainerPortProtocol) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
