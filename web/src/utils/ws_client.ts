const getSocketURL = (path: string) => {
    if (process.env.NODE_ENV==="development"){
        return `${process.env.REACT_APP_SERVER_WS_URL}${path}`
    }
    if (document.location.protocol === "http:") {
        return `ws://${document.location.host}${path}`
    }
    return `${process.env.REACT_APP_SERVER_WS_URL}${path}`
}

const iotTelemetrySocketURL = getSocketURL("/ws/iot/telemetry")
const iotHealthSocketURL = getSocketURL("/ws/iot/health")

export { iotTelemetrySocketURL, iotHealthSocketURL }
