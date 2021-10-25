package graph
import (
    "strings"
    "errors"
)
type PrinterProcessingStateDetail int

const (
    PAUSED_PRINTERPROCESSINGSTATEDETAIL PrinterProcessingStateDetail = iota
    MEDIAJAM_PRINTERPROCESSINGSTATEDETAIL
    MEDIANEEDED_PRINTERPROCESSINGSTATEDETAIL
    MEDIALOW_PRINTERPROCESSINGSTATEDETAIL
    MEDIAEMPTY_PRINTERPROCESSINGSTATEDETAIL
    COVEROPEN_PRINTERPROCESSINGSTATEDETAIL
    INTERLOCKOPEN_PRINTERPROCESSINGSTATEDETAIL
    OUTPUTTRAYMISSING_PRINTERPROCESSINGSTATEDETAIL
    OUTPUTAREAFULL_PRINTERPROCESSINGSTATEDETAIL
    MARKERSUPPLYLOW_PRINTERPROCESSINGSTATEDETAIL
    MARKERSUPPLYEMPTY_PRINTERPROCESSINGSTATEDETAIL
    INPUTTRAYMISSING_PRINTERPROCESSINGSTATEDETAIL
    OUTPUTAREAALMOSTFULL_PRINTERPROCESSINGSTATEDETAIL
    MARKERWASTEALMOSTFULL_PRINTERPROCESSINGSTATEDETAIL
    MARKERWASTEFULL_PRINTERPROCESSINGSTATEDETAIL
    FUSEROVERTEMP_PRINTERPROCESSINGSTATEDETAIL
    FUSERUNDERTEMP_PRINTERPROCESSINGSTATEDETAIL
    OTHER_PRINTERPROCESSINGSTATEDETAIL
    NONE_PRINTERPROCESSINGSTATEDETAIL
    MOVINGTOPAUSED_PRINTERPROCESSINGSTATEDETAIL
    SHUTDOWN_PRINTERPROCESSINGSTATEDETAIL
    CONNECTINGTODEVICE_PRINTERPROCESSINGSTATEDETAIL
    TIMEDOUT_PRINTERPROCESSINGSTATEDETAIL
    STOPPING_PRINTERPROCESSINGSTATEDETAIL
    STOPPEDPARTIALLY_PRINTERPROCESSINGSTATEDETAIL
    TONERLOW_PRINTERPROCESSINGSTATEDETAIL
    TONEREMPTY_PRINTERPROCESSINGSTATEDETAIL
    SPOOLAREAFULL_PRINTERPROCESSINGSTATEDETAIL
    DOOROPEN_PRINTERPROCESSINGSTATEDETAIL
    OPTICALPHOTOCONDUCTORNEARENDOFLIFE_PRINTERPROCESSINGSTATEDETAIL
    OPTICALPHOTOCONDUCTORLIFEOVER_PRINTERPROCESSINGSTATEDETAIL
    DEVELOPERLOW_PRINTERPROCESSINGSTATEDETAIL
    DEVELOPEREMPTY_PRINTERPROCESSINGSTATEDETAIL
    INTERPRETERRESOURCEUNAVAILABLE_PRINTERPROCESSINGSTATEDETAIL
    UNKNOWNFUTUREVALUE_PRINTERPROCESSINGSTATEDETAIL
)

func (i PrinterProcessingStateDetail) String() string {
    return []string{"PAUSED", "MEDIAJAM", "MEDIANEEDED", "MEDIALOW", "MEDIAEMPTY", "COVEROPEN", "INTERLOCKOPEN", "OUTPUTTRAYMISSING", "OUTPUTAREAFULL", "MARKERSUPPLYLOW", "MARKERSUPPLYEMPTY", "INPUTTRAYMISSING", "OUTPUTAREAALMOSTFULL", "MARKERWASTEALMOSTFULL", "MARKERWASTEFULL", "FUSEROVERTEMP", "FUSERUNDERTEMP", "OTHER", "NONE", "MOVINGTOPAUSED", "SHUTDOWN", "CONNECTINGTODEVICE", "TIMEDOUT", "STOPPING", "STOPPEDPARTIALLY", "TONERLOW", "TONEREMPTY", "SPOOLAREAFULL", "DOOROPEN", "OPTICALPHOTOCONDUCTORNEARENDOFLIFE", "OPTICALPHOTOCONDUCTORLIFEOVER", "DEVELOPERLOW", "DEVELOPEREMPTY", "INTERPRETERRESOURCEUNAVAILABLE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePrinterProcessingStateDetail(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "PAUSED":
            return PAUSED_PRINTERPROCESSINGSTATEDETAIL, nil
        case "MEDIAJAM":
            return MEDIAJAM_PRINTERPROCESSINGSTATEDETAIL, nil
        case "MEDIANEEDED":
            return MEDIANEEDED_PRINTERPROCESSINGSTATEDETAIL, nil
        case "MEDIALOW":
            return MEDIALOW_PRINTERPROCESSINGSTATEDETAIL, nil
        case "MEDIAEMPTY":
            return MEDIAEMPTY_PRINTERPROCESSINGSTATEDETAIL, nil
        case "COVEROPEN":
            return COVEROPEN_PRINTERPROCESSINGSTATEDETAIL, nil
        case "INTERLOCKOPEN":
            return INTERLOCKOPEN_PRINTERPROCESSINGSTATEDETAIL, nil
        case "OUTPUTTRAYMISSING":
            return OUTPUTTRAYMISSING_PRINTERPROCESSINGSTATEDETAIL, nil
        case "OUTPUTAREAFULL":
            return OUTPUTAREAFULL_PRINTERPROCESSINGSTATEDETAIL, nil
        case "MARKERSUPPLYLOW":
            return MARKERSUPPLYLOW_PRINTERPROCESSINGSTATEDETAIL, nil
        case "MARKERSUPPLYEMPTY":
            return MARKERSUPPLYEMPTY_PRINTERPROCESSINGSTATEDETAIL, nil
        case "INPUTTRAYMISSING":
            return INPUTTRAYMISSING_PRINTERPROCESSINGSTATEDETAIL, nil
        case "OUTPUTAREAALMOSTFULL":
            return OUTPUTAREAALMOSTFULL_PRINTERPROCESSINGSTATEDETAIL, nil
        case "MARKERWASTEALMOSTFULL":
            return MARKERWASTEALMOSTFULL_PRINTERPROCESSINGSTATEDETAIL, nil
        case "MARKERWASTEFULL":
            return MARKERWASTEFULL_PRINTERPROCESSINGSTATEDETAIL, nil
        case "FUSEROVERTEMP":
            return FUSEROVERTEMP_PRINTERPROCESSINGSTATEDETAIL, nil
        case "FUSERUNDERTEMP":
            return FUSERUNDERTEMP_PRINTERPROCESSINGSTATEDETAIL, nil
        case "OTHER":
            return OTHER_PRINTERPROCESSINGSTATEDETAIL, nil
        case "NONE":
            return NONE_PRINTERPROCESSINGSTATEDETAIL, nil
        case "MOVINGTOPAUSED":
            return MOVINGTOPAUSED_PRINTERPROCESSINGSTATEDETAIL, nil
        case "SHUTDOWN":
            return SHUTDOWN_PRINTERPROCESSINGSTATEDETAIL, nil
        case "CONNECTINGTODEVICE":
            return CONNECTINGTODEVICE_PRINTERPROCESSINGSTATEDETAIL, nil
        case "TIMEDOUT":
            return TIMEDOUT_PRINTERPROCESSINGSTATEDETAIL, nil
        case "STOPPING":
            return STOPPING_PRINTERPROCESSINGSTATEDETAIL, nil
        case "STOPPEDPARTIALLY":
            return STOPPEDPARTIALLY_PRINTERPROCESSINGSTATEDETAIL, nil
        case "TONERLOW":
            return TONERLOW_PRINTERPROCESSINGSTATEDETAIL, nil
        case "TONEREMPTY":
            return TONEREMPTY_PRINTERPROCESSINGSTATEDETAIL, nil
        case "SPOOLAREAFULL":
            return SPOOLAREAFULL_PRINTERPROCESSINGSTATEDETAIL, nil
        case "DOOROPEN":
            return DOOROPEN_PRINTERPROCESSINGSTATEDETAIL, nil
        case "OPTICALPHOTOCONDUCTORNEARENDOFLIFE":
            return OPTICALPHOTOCONDUCTORNEARENDOFLIFE_PRINTERPROCESSINGSTATEDETAIL, nil
        case "OPTICALPHOTOCONDUCTORLIFEOVER":
            return OPTICALPHOTOCONDUCTORLIFEOVER_PRINTERPROCESSINGSTATEDETAIL, nil
        case "DEVELOPERLOW":
            return DEVELOPERLOW_PRINTERPROCESSINGSTATEDETAIL, nil
        case "DEVELOPEREMPTY":
            return DEVELOPEREMPTY_PRINTERPROCESSINGSTATEDETAIL, nil
        case "INTERPRETERRESOURCEUNAVAILABLE":
            return INTERPRETERRESOURCEUNAVAILABLE_PRINTERPROCESSINGSTATEDETAIL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PRINTERPROCESSINGSTATEDETAIL, nil
    }
    return 0, errors.New("Unknown PrinterProcessingStateDetail value: " + v)
}
func SerializePrinterProcessingStateDetail(values []PrinterProcessingStateDetail) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}