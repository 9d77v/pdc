const getSocketURL = (path: string) => {
    let protocol = "wss:"
    if (document.location.protocol === "http:") {
        protocol = "ws:"
    }
    return `${protocol}//${document.location.host}${path}`
}

const deviceTelemetryPrefix = "device.telemetry"
const deviceHealthPrefix = "device.health"
const iotSocketURL = getSocketURL("/ws/iot")

export { deviceTelemetryPrefix, deviceHealthPrefix, iotSocketURL }