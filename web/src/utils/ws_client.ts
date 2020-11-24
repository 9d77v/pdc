const getSocketURL = (path: string) => {
    let protocol = "wss:"
    if (document.location.protocol === "http:") {
        protocol = "ws:"
    }
    return `${protocol}//${document.location.host}${path}`
}

const iotTelemetrySocketURL = getSocketURL("/ws/iot/telemetry")
const iotHealthSocketURL = getSocketURL("/ws/iot/health")

export { iotTelemetrySocketURL, iotHealthSocketURL }
