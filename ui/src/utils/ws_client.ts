const getSocketURL = (path: string) => {
    let protocol = "wss:"
    if (document.location.protocol === "http:") {
        protocol = "ws:"
    }
    return `${protocol}//${document.location.host}${path}`
}

const deviceTelemetryPrefix = "device.telemetry"
const deviceHealthPrefix = "device.health"
const iotTelemetrySocketURL = getSocketURL("/ws/iot/telemetry")
const iotHealthSocketURL = getSocketURL("/ws/iot/health")


export { deviceTelemetryPrefix, deviceHealthPrefix, iotTelemetrySocketURL, iotHealthSocketURL }